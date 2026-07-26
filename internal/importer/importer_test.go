package importer_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/sairaph/apis-mcp/internal/importer"
	"github.com/sairaph/apis-mcp/library"
)

func TestImportMarkdownRejectsExistingDestinationAndDocID(t *testing.T) {
	ctx := context.Background()
	root := t.TempDir()
	index := filepath.Join(t.TempDir(), "library.sqlite")
	write(t, root, "widget-service/v1/_index.md", manifest("Widget Service", "v1"))
	write(t, root, "widget-service/v1/old.md", page("Old Page", "# Old\n"))
	rebuild := rebuildFunc(root, index)
	if err := rebuild(ctx); err != nil {
		t.Fatal(err)
	}

	source := t.TempDir()
	write(t, source, "_index.md", manifest("Widget Service", "v1"))
	write(t, source, "guide/new.md", page("New Page", "# New\n"))
	_, err := importer.ImportMarkdown(ctx, source, importer.Options{LibraryRoot: root, Rebuild: rebuild})
	var collision *importer.CollisionError
	if !errors.As(err, &collision) || collision.DocID != "widget-service-v1" ||
		!containsPath(collision.Paths, filepath.Join(root, "widget-service", "v1")) {
		t.Fatalf("expected destination collision paths, got %v", err)
	}
	if _, err := os.Stat(filepath.Join(root, "widget-service", "v1", "old.md")); err != nil {
		t.Fatalf("existing destination changed: %v", err)
	}

	write(t, root, "legacy/location/_index.md", manifest("Alias Service", "v2"))
	alias := t.TempDir()
	write(t, alias, "_index.md", manifest("Alias Service", "v2"))
	write(t, alias, "new.md", page("New", "# New\n"))
	_, err = importer.ImportMarkdown(ctx, alias, importer.Options{LibraryRoot: root, Rebuild: rebuild})
	if !errors.As(err, &collision) ||
		!containsPath(collision.Paths, filepath.Join(root, "legacy", "location", "_index.md")) ||
		!containsPath(collision.Paths, filepath.Join(root, "alias-service", "v2", "_index.md")) {
		t.Fatalf("expected doc_id collision paths, got %v", err)
	}
}

func TestImportMarkdownRollsBackInvalidNewDocument(t *testing.T) {
	ctx := context.Background()
	root := t.TempDir()
	index := filepath.Join(t.TempDir(), "library.sqlite")
	rebuild := rebuildFunc(root, index)
	if err := rebuild(ctx); err != nil {
		t.Fatal(err)
	}
	before := generationBytes(t, index)
	bad := t.TempDir()
	write(t, bad, "_index.md", manifest("Broken Service", "v1"))
	write(t, bad, "broken.md", "not canonical\n")
	if _, err := importer.ImportMarkdown(ctx, bad, importer.Options{LibraryRoot: root, Rebuild: rebuild}); err == nil || !strings.Contains(err.Error(), "validate imported documentation") {
		t.Fatalf("expected validation failure, got %v", err)
	}
	if _, err := os.Stat(filepath.Join(root, "broken-service", "v1")); !os.IsNotExist(err) {
		t.Fatalf("invalid tree remained after rollback: %v", err)
	}
	if after := generationBytes(t, index); !bytes.Equal(before, after) {
		t.Fatal("failed import changed the published index")
	}
}

func TestImportMarkdownRejectsSymlinksAndNestedManifests(t *testing.T) {
	for _, test := range []struct {
		name  string
		setup func(*testing.T, string)
		want  string
	}{
		{"symlink", func(t *testing.T, root string) {
			if err := os.Symlink(filepath.Join(root, "_index.md"), filepath.Join(root, "linked.md")); err != nil {
				t.Fatal(err)
			}
		}, "symbolic links"},
		{"nested manifest", func(t *testing.T, root string) {
			write(t, root, "nested/_index.md", manifest("Nested", "v1"))
		}, "nested canonical manifest"},
	} {
		t.Run(test.name, func(t *testing.T) {
			source := t.TempDir()
			write(t, source, "_index.md", manifest("Safe API", "v1"))
			test.setup(t, source)
			_, err := importer.ImportMarkdown(context.Background(), source, importer.Options{
				LibraryRoot: t.TempDir(), Rebuild: func(context.Context) error { return nil },
			})
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("expected %q error, got %v", test.want, err)
			}
		})
	}
}

func TestImportOpenAPIFromHTTPGeneratesOperationsSchemasAndProvenance(t *testing.T) {
	const specification = `openapi: 3.1.0
info:
  title: Upstream title
  version: ignored
  description: Manage pets.
servers:
  - url: https://api.example.test/v1
paths:
  /pets/{petId}:
    get:
      summary: Get a pet
      description: Finds one pet.
      operationId: getPet
      tags: [Pets]
      responses:
        "200":
          description: Found
components:
  schemas:
    Pet:
      description: A pet record.
      type: object
      properties:
        id: {type: string}
`
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/openapi.yaml" {
			http.NotFound(writer, request)
			return
		}
		writer.Header().Set("Content-Type", "application/yaml")
		_, _ = writer.Write([]byte(specification))
	}))
	defer server.Close()

	root := t.TempDir()
	index := filepath.Join(t.TempDir(), "library.sqlite")
	result, err := importer.ImportOpenAPI(context.Background(), "Pet API", "2026-01", server.URL+"/openapi.yaml", importer.Options{
		LibraryRoot: root, Rebuild: rebuildFunc(root, index), HTTPClient: server.Client(),
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Kind != "openapi" || result.Pages != 3 || result.Destination != filepath.Join(root, "pet-api", "2026-01") {
		t.Fatalf("unexpected result: %+v", result)
	}
	operationFiles, err := filepath.Glob(filepath.Join(result.Destination, "operations", "pets", "*.md"))
	if err != nil || len(operationFiles) != 1 {
		t.Fatalf("operation files: %v, %v", operationFiles, err)
	}
	operation, err := os.ReadFile(operationFiles[0])
	if err != nil {
		t.Fatal(err)
	}
	for _, wanted := range []string{"http_methods:\n    - GET", "api_endpoints:\n    - /pets/{petId}", "operation_ids:\n    - getPet", "source: " + server.URL + "/openapi.yaml", "`GET /pets/{petId}`"} {
		if !strings.Contains(string(operation), wanted) {
			t.Errorf("operation page missing %q:\n%s", wanted, operation)
		}
	}
	schemaFiles, _ := filepath.Glob(filepath.Join(result.Destination, "schemas", "*.md"))
	if len(schemaFiles) != 1 {
		t.Fatalf("schema files: %v", schemaFiles)
	}

	snapshot, err := library.Open(context.Background(), library.Options{UserRoot: root, IndexPath: index})
	if err != nil {
		t.Fatal(err)
	}
	defer snapshot.Close()
	search, err := snapshot.Search(context.Background(), library.SearchRequest{DocID: "pet-api-2026-01", Query: "getPet"})
	if err != nil || search.Total == 0 {
		t.Fatalf("operation metadata was not indexed: %+v, %v", search, err)
	}

	firstNames := relativeFiles(t, result.Destination)
	if _, err := importer.ImportOpenAPI(context.Background(), "Pet API", "2026-01", server.URL+"/openapi.yaml", importer.Options{
		LibraryRoot: root, Rebuild: rebuildFunc(root, index), HTTPClient: server.Client(),
	}); err == nil || !strings.Contains(err.Error(), "pet-api-2026-01") {
		t.Fatalf("expected duplicate import rejection, got %v", err)
	}
	if secondNames := relativeFiles(t, result.Destination); strings.Join(firstNames, "\n") != strings.Join(secondNames, "\n") {
		t.Fatalf("generated filenames changed: %v != %v", firstNames, secondNames)
	}
}

func TestConcurrentImportsSerializePublicationAndRebuild(t *testing.T) {
	root := t.TempDir()
	var active atomic.Int32
	var maximum atomic.Int32
	rebuild := func(context.Context) error {
		current := active.Add(1)
		for {
			previous := maximum.Load()
			if current <= previous || maximum.CompareAndSwap(previous, current) {
				break
			}
		}
		time.Sleep(40 * time.Millisecond)
		active.Add(-1)
		return nil
	}
	sources := []string{t.TempDir(), t.TempDir()}
	write(t, sources[0], "_index.md", manifest("Concurrent One", "v1"))
	write(t, sources[1], "_index.md", manifest("Concurrent Two", "v1"))
	for _, source := range sources {
		write(t, source, "page.md", page("Page", "# Page\n"))
	}
	var group sync.WaitGroup
	errorsFound := make(chan error, len(sources))
	for _, source := range sources {
		group.Add(1)
		go func() {
			defer group.Done()
			_, err := importer.ImportMarkdown(context.Background(), source, importer.Options{LibraryRoot: root, Rebuild: rebuild})
			errorsFound <- err
		}()
	}
	group.Wait()
	close(errorsFound)
	for err := range errorsFound {
		if err != nil {
			t.Fatal(err)
		}
	}
	if maximum.Load() != 1 {
		t.Fatalf("import rebuilds overlapped: maximum concurrency %d", maximum.Load())
	}
}

func TestImportTransactionLockAcrossProcesses(t *testing.T) {
	if os.Getenv("APIS_MCP_IMPORT_LOCK_CHILD") == "1" {
		rebuild := func(context.Context) error {
			marker := os.Getenv("APIS_MCP_IMPORT_LOCK_MARKER")
			file, err := os.OpenFile(marker, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o600)
			if err != nil {
				return fmt.Errorf("concurrent rebuild: %w", err)
			}
			if err := file.Close(); err != nil {
				return err
			}
			defer os.Remove(marker)
			time.Sleep(200 * time.Millisecond)
			return nil
		}
		_, err := importer.ImportMarkdown(context.Background(), os.Getenv("APIS_MCP_IMPORT_LOCK_SOURCE"), importer.Options{
			LibraryRoot: os.Getenv("APIS_MCP_IMPORT_LOCK_ROOT"), Rebuild: rebuild,
		})
		if err != nil {
			t.Fatal(err)
		}
		return
	}

	root := t.TempDir()
	marker := filepath.Join(t.TempDir(), "rebuild-active")
	sources := []string{t.TempDir(), t.TempDir()}
	write(t, sources[0], "_index.md", manifest("Process One", "v1"))
	write(t, sources[1], "_index.md", manifest("Process Two", "v1"))
	for _, source := range sources {
		write(t, source, "page.md", page("Page", "# Page\n"))
	}
	commands := make([]*exec.Cmd, 0, len(sources))
	var outputs [2]bytes.Buffer
	for index, source := range sources {
		command := exec.Command(os.Args[0], "-test.run=^TestImportTransactionLockAcrossProcesses$", "-test.count=1")
		command.Env = append(os.Environ(),
			"APIS_MCP_IMPORT_LOCK_CHILD=1",
			"APIS_MCP_IMPORT_LOCK_ROOT="+root,
			"APIS_MCP_IMPORT_LOCK_MARKER="+marker,
			"APIS_MCP_IMPORT_LOCK_SOURCE="+source,
		)
		command.Stdout, command.Stderr = &outputs[index], &outputs[index]
		if err := command.Start(); err != nil {
			t.Fatal(err)
		}
		commands = append(commands, command)
	}
	for index, command := range commands {
		if err := command.Wait(); err != nil {
			t.Fatalf("import process %d: %v\n%s", index, err, outputs[index].String())
		}
	}
}

func TestImportSwaggerJSONAndDownloadLimit(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path == "/large" {
			_, _ = writer.Write([]byte(strings.Repeat("x", 64)))
			return
		}
		fmt.Fprint(writer, `{"swagger":"2.0","info":{"title":"Old","version":"1"},"host":"api.example.test","basePath":"/v1","paths":{"/items":{"post":{"operationId":"createItem","responses":{"200":{"description":"ok"}}}}},"definitions":{"Item":{"type":"object"}}}`)
	}))
	defer server.Close()
	root := t.TempDir()
	result, err := importer.ImportOpenAPI(context.Background(), "Legacy API", "v1", server.URL+"/swagger.json", importer.Options{
		LibraryRoot: root, Rebuild: rebuildFunc(root, filepath.Join(t.TempDir(), "index.sqlite")), HTTPClient: server.Client(),
	})
	if err != nil || result.Kind != "swagger" || result.Pages != 3 {
		t.Fatalf("Swagger import: %+v, %v", result, err)
	}
	operationFiles, _ := filepath.Glob(filepath.Join(result.Destination, "operations", "untagged", "*.md"))
	if len(operationFiles) != 1 {
		t.Fatalf("Swagger operation files: %v", operationFiles)
	}
	operation, readErr := os.ReadFile(operationFiles[0])
	if readErr != nil {
		t.Fatal(readErr)
	}
	if strings.Contains(string(operation), "description:") {
		t.Fatalf("generated fallback was stored as metadata:\n%s", operation)
	}
	_, err = importer.ImportOpenAPI(context.Background(), "Too Large", "v1", server.URL+"/large", importer.Options{
		LibraryRoot: root, Rebuild: func(context.Context) error { return nil }, HTTPClient: server.Client(), MaxSourceBytes: 16, MaxTotalBytes: 16,
	})
	if err == nil || !strings.Contains(err.Error(), "exceeds 16 bytes") {
		t.Fatalf("expected bounded download error, got %v", err)
	}
}

func TestImportOpenAPIDiscoversSwaggerUIAndRedocSpecifications(t *testing.T) {
	tests := []struct {
		name       string
		landing    string
		landingURL string
		specPath   string
		spec       string
		kind       string
	}{
		{
			name:       "Swagger UI urls array",
			landingURL: "/swagger/ui/index.html",
			specPath:   "/specs/swagger.json",
			landing: `<html><head><title>Swagger</title></head><body><div id="swagger-ui"></div><script>
window.ui = SwaggerUIBundle({
  "urls": [{"url": "../../specs/swagger.json", "name": "current"}]
});
</script></body></html>`,
			spec: `{"swagger":"2.0","info":{"title":"Legacy","version":"1"},"paths":{"/legacy":{"get":{"responses":{"200":{"description":"ok"}}}}}}`,
			kind: "swagger",
		},
		{
			name:       "Redoc spec-url",
			landingURL: "/redoc/index.html",
			specPath:   "/openapi.yaml",
			landing:    `<html><body><redoc spec-url="../openapi.yaml"></redoc><script src="redoc.js"></script></body></html>`,
			spec:       "openapi: 3.0.3\ninfo: {title: Current, version: v1}\npaths:\n  /current:\n    post:\n      responses:\n        '200': {description: ok}\n",
			kind:       "openapi",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var server *httptest.Server
			server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				switch request.URL.Path {
				case test.landingURL:
					writer.Header().Set("Content-Type", "text/html")
					fmt.Fprint(writer, test.landing)
				case test.specPath:
					fmt.Fprint(writer, test.spec)
				default:
					http.NotFound(writer, request)
				}
			}))
			defer server.Close()

			root := t.TempDir()
			result, err := importer.ImportOpenAPI(context.Background(), "Discovered API", "v1", server.URL+test.landingURL, importer.Options{
				LibraryRoot: root, Rebuild: rebuildFunc(root, filepath.Join(t.TempDir(), "index.sqlite")), HTTPClient: server.Client(),
			})
			if err != nil {
				t.Fatal(err)
			}
			if result.Kind != test.kind || result.Source != server.URL+test.specPath || result.Pages != 2 {
				t.Fatalf("unexpected result: %+v", result)
			}
			overview, err := os.ReadFile(filepath.Join(result.Destination, "overview.md"))
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(overview), "source: "+server.URL+test.specPath) || !strings.Contains(string(overview), "imported_from: "+server.URL+test.specPath) {
				t.Fatalf("resolved specification provenance missing:\n%s", overview)
			}
		})
	}
}

func TestImportHTMLCrawlsSameOriginWithinLimitsAndGeneratesMarkdown(t *testing.T) {
	externalRequests := 0
	external := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		externalRequests++
	}))
	defer external.Close()

	deepRequests := 0
	assetRequests := 0
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		switch request.URL.Path {
		case "/docs/index.html":
			fmt.Fprintf(writer, `<!doctype html><html><head><title>Acme Docs</title><meta name="description" content="Explicit Acme metadata."></head><body>
<nav><a href="%s/outside.html">Outside</a></nav><main>
<h1>Acme API</h1><p>Use <strong>widgets</strong> with <code>client.create</code>. Read the <a href="guide.html">guide</a>.</p>
<h2>Features</h2><ul><li>Fast</li><li>Safe</li></ul>
<table><thead><tr><th>Field</th><th>Meaning</th></tr></thead><tbody><tr><td>id</td><td>Widget ID</td></tr></tbody></table>
<pre><code class="language-json">{"ok": true}</code></pre>
<a href="schema.json">Not HTML</a></main></body></html>`, external.URL)
		case "/docs/guide.html":
			fmt.Fprint(writer, `<html><head><title>Guide</title></head><body><article><h1>Guide</h1><ol><li>Create a widget</li><li>Check status</li></ol><p><a href="deep/page.html">Deep page</a></p></article></body></html>`)
		case "/docs/deep/page.html":
			deepRequests++
			fmt.Fprint(writer, `<html><body><h1>Too deep</h1></body></html>`)
		case "/docs/schema.json":
			assetRequests++
			fmt.Fprint(writer, `{}`)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	root := t.TempDir()
	result, err := importer.ImportHTML(context.Background(), "Acme Static", "v1", server.URL+"/docs/index.html", importer.Options{
		LibraryRoot: root, Rebuild: rebuildFunc(root, filepath.Join(t.TempDir(), "index.sqlite")), HTTPClient: server.Client(),
		MaxHTMLPages: 5, MaxHTMLDepth: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Kind != "html" || result.Pages != 2 || result.Source != server.URL+"/docs/index.html" {
		t.Fatalf("unexpected result: %+v", result)
	}
	if externalRequests != 0 || deepRequests != 0 || assetRequests != 0 {
		t.Fatalf("crawl escaped boundaries: external=%d deep=%d asset=%d", externalRequests, deepRequests, assetRequests)
	}

	pageNames, err := filepath.Glob(filepath.Join(result.Destination, "documentation", "*.md"))
	if err != nil || len(pageNames) != 2 {
		t.Fatalf("generated pages: %v, %v", pageNames, err)
	}
	var generated string
	for _, pageName := range pageNames {
		raw, readErr := os.ReadFile(pageName)
		if readErr != nil {
			t.Fatal(readErr)
		}
		generated += string(raw) + "\n"
	}
	for _, wanted := range []string{
		"title: Acme Docs", "source: " + server.URL + "/docs/index.html", "imported_from: " + server.URL + "/docs/index.html",
		"description: Explicit Acme metadata.",
		"# Acme API", "Use **widgets** with `client.create`.", "[guide](" + server.URL + "/docs/guide.html)",
		"- Fast", "| Field | Meaning |", "| --- | --- |", "```json\n{\"ok\": true}\n```", "1. Create a widget",
	} {
		if !strings.Contains(generated, wanted) {
			t.Errorf("generated Markdown missing %q:\n%s", wanted, generated)
		}
	}
	if count := strings.Count(generated, "description:"); count != 1 {
		t.Fatalf("generated page descriptions = %d, want only explicit metadata:\n%s", count, generated)
	}
}

func TestImportHTMLRejectsExcessiveTreeDepthAndNodes(t *testing.T) {
	tests := []struct {
		name string
		body string
		want string
	}{
		{name: "depth", body: "<html><body>" + strings.Repeat("<div>", 1_000) + "bounded", want: "maximum tree depth"},
		{name: "nodes", body: "<html><body>" + strings.Repeat("<br>", 50_001), want: "maximum node count"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
				fmt.Fprint(writer, test.body)
			}))
			defer server.Close()
			root := t.TempDir()
			_, err := importer.ImportHTML(context.Background(), "Unsafe HTML", test.name, server.URL, importer.Options{
				LibraryRoot: root, Rebuild: func(context.Context) error { return nil }, HTTPClient: server.Client(),
			})
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("expected bounded HTML rejection %q, got %v", test.want, err)
			}
			if _, statErr := os.Stat(filepath.Join(root, "unsafe-html", test.name)); !os.IsNotExist(statErr) {
				t.Fatalf("unsafe HTML was published: %v", statErr)
			}
		})
	}
}

func TestImportLLMSTxtFollowsRelativeHTTPLinksAndPreservesFrontmatter(t *testing.T) {
	guideRequests := 0
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/docs/llms.txt":
			fmt.Fprint(writer, "# Acme Docs\n\nPractical Acme documentation.\n\n- [Guide](guide.md)\n- [Reference](reference.md): Explicit reference metadata.\n- [Duplicate](guide.md)\n")
		case "/docs/guide.md":
			guideRequests++
			fmt.Fprint(writer, "---\ntitle: Original Guide\nsource: https://author.example/guide\ncustom: retained\n---\n\n# Better Guide\n\nUse Acme.\n")
		case "/docs/reference.md":
			fmt.Fprint(writer, "# Reference\n\nDetails.\n")
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()
	root := t.TempDir()
	result, err := importer.ImportLLMSTxt(context.Background(), "Acme API", "v2", server.URL+"/docs/llms.txt", importer.Options{
		LibraryRoot: root, Rebuild: rebuildFunc(root, filepath.Join(t.TempDir(), "index.sqlite")), HTTPClient: server.Client(),
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Pages != 2 || result.Kind != "llms.txt" {
		t.Fatalf("unexpected result: %+v", result)
	}
	if guideRequests != 1 {
		t.Fatalf("duplicate link fetched %d times", guideRequests)
	}
	pages, _ := filepath.Glob(filepath.Join(result.Destination, "documentation", "*.md"))
	if len(pages) != 2 {
		t.Fatalf("linked pages: %v", pages)
	}
	var guide, reference string
	for _, name := range pages {
		raw, readErr := os.ReadFile(name)
		if readErr != nil {
			t.Fatal(readErr)
		}
		if strings.Contains(string(raw), "Better Guide") {
			guide = string(raw)
		} else if strings.Contains(string(raw), "# Reference") {
			reference = string(raw)
		}
	}
	if !strings.Contains(reference, "description: Explicit reference metadata.") {
		t.Fatalf("reference link description was not preserved:\n%s", reference)
	}
	for _, wanted := range []string{"custom: retained", "source: " + server.URL + "/docs/guide.md", "upstream_source: https://author.example/guide", "imported_from: " + server.URL + "/docs/llms.txt"} {
		if !strings.Contains(guide, wanted) {
			t.Errorf("guide missing %q:\n%s", wanted, guide)
		}
	}
}

func TestRemoteLLMSTxtCannotReadLocalFiles(t *testing.T) {
	local := filepath.Join(t.TempDir(), "private.md")
	if err := os.WriteFile(local, []byte("# Private\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(writer, "[private](file://%s)\n", filepath.ToSlash(local))
	}))
	defer server.Close()
	_, err := importer.ImportLLMSTxt(context.Background(), "Unsafe", "v1", server.URL, importer.Options{
		LibraryRoot: t.TempDir(), Rebuild: func(context.Context) error { return nil }, HTTPClient: server.Client(),
	})
	if err == nil || !strings.Contains(err.Error(), "only link to HTTP(S)") {
		t.Fatalf("expected remote-to-local source rejection, got %v", err)
	}
}

func TestSafeSlugMatchesCanonicalPortableIDs(t *testing.T) {
	for input, wanted := range map[string]string{"Pet's API": "pet-s-api", "A@B/C": "a-b-c", " API v2 ": "api-v2"} {
		if got := importer.SafeSlug(input); got != wanted {
			t.Errorf("SafeSlug(%q) = %q, want %q", input, got, wanted)
		}
	}
}

func TestImportLLMSTxtReadsRelativeLocalFiles(t *testing.T) {
	source := t.TempDir()
	write(t, source, "llms.txt", "# Local\n\n- [One](pages/one.md)\n")
	write(t, source, "pages/one.md", "# Local Page\n\nLocal content.\n")
	root := t.TempDir()
	result, err := importer.ImportLLMSTxt(context.Background(), "Local API", "v1", filepath.Join(source, "llms.txt"), importer.Options{
		LibraryRoot: root, Rebuild: rebuildFunc(root, filepath.Join(t.TempDir(), "index.sqlite")),
	})
	if err != nil || result.Pages != 1 {
		t.Fatalf("local llms.txt import: %+v, %v", result, err)
	}
}

func rebuildFunc(root, index string) func(context.Context) error {
	return func(ctx context.Context) error {
		return library.Rebuild(ctx, library.Options{UserRoot: root, IndexPath: index})
	}
}

func manifest(name, version string) string {
	return fmt.Sprintf("---\nname: %s\nversion: %s\ndescription: Test documentation.\n---\n", name, version)
}

func page(title, body string) string {
	return "---\ntitle: " + title + "\n---\n\n" + body
}

func write(t *testing.T, root, relative, content string) {
	t.Helper()
	name := filepath.Join(root, relative)
	if err := os.MkdirAll(filepath.Dir(name), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(name, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func relativeFiles(t *testing.T, root string) []string {
	t.Helper()
	var names []string
	if err := filepath.WalkDir(root, func(name string, entry os.DirEntry, err error) error {
		if err != nil || entry.IsDir() {
			return err
		}
		relative, err := filepath.Rel(root, name)
		if err == nil {
			names = append(names, relative)
		}
		return err
	}); err != nil {
		t.Fatal(err)
	}
	sort.Strings(names)
	return names
}

func containsPath(paths []string, wanted string) bool {
	for _, name := range paths {
		if name == wanted {
			return true
		}
	}
	return false
}

func generationBytes(t *testing.T, index string) []byte {
	t.Helper()
	extension := filepath.Ext(index)
	names, err := filepath.Glob(strings.TrimSuffix(index, extension) + "-*" + extension)
	if err != nil || len(names) != 1 {
		t.Fatalf("published generations: %v, %v", names, err)
	}
	raw, err := os.ReadFile(names[0])
	if err != nil {
		t.Fatal(err)
	}
	return raw
}
