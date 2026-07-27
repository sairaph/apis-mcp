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

func TestImportOpenAPIDiscoversAllStaticRapiDocSpecifications(t *testing.T) {
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/docs/":
			fmt.Fprint(writer, `<html><body><rapi-doc spec-url="./pets.json"></rapi-doc><rapi-doc-mini spec-url="./orders.json"></rapi-doc-mini></body></html>`)
		case "/docs/pets.json":
			fmt.Fprint(writer, `{"openapi":"3.0.3","info":{"title":"Pets","version":"v1"},"paths":{"/pets":{"get":{"responses":{"200":{"description":"ok"}}}}}}`)
		case "/docs/orders.json":
			fmt.Fprint(writer, `{"swagger":"2.0","info":{"title":"Orders","version":"v1"},"paths":{"/orders":{"post":{"responses":{"200":{"description":"ok"}}}}}}`)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	root := t.TempDir()
	index := filepath.Join(t.TempDir(), "index.sqlite")
	result, err := importer.ImportOpenAPI(context.Background(), "RapiDoc APIs", "live", server.URL+"/docs/", importer.Options{
		LibraryRoot: root, Rebuild: rebuildFunc(root, index), HTTPClient: server.Client(),
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Kind != "openapi-catalog" || result.Framework != "rapidoc" || result.Sources != 2 || result.Pages != 5 || result.Source != server.URL+"/docs/" {
		t.Fatalf("unexpected result: %+v", result)
	}
	for _, relative := range []string{
		"apis/docs/orders/operations/untagged", "apis/docs/pets/operations/untagged",
	} {
		files, globErr := filepath.Glob(filepath.Join(result.Destination, relative, "*.md"))
		if globErr != nil || len(files) != 1 {
			t.Fatalf("generated %s: %v, %v", relative, files, globErr)
		}
	}
	snapshot, err := library.Open(context.Background(), library.Options{UserRoot: root, IndexPath: index})
	if err != nil {
		t.Fatal(err)
	}
	defer snapshot.Close()
	search, err := snapshot.Search(context.Background(), library.SearchRequest{DocID: "rapidoc-apis-live", Query: "orders"})
	if err != nil || search.Total == 0 {
		t.Fatalf("indexed catalog search: %+v, %v", search, err)
	}
}

func TestImportOpenAPIExpandsOVHRapiDocCatalogs(t *testing.T) {
	requested := make(map[string]int)
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		requested[request.URL.Path+"?"+request.URL.RawQuery]++
		switch request.URL.Path {
		case "/console/":
			fmt.Fprintf(writer, `<html><body><rapi-doc spec-url="" spec-roots="%s/v1,%s/v2"></rapi-doc></body></html>`, server.URL, server.URL)
		case "/v1":
			fmt.Fprintf(writer, `{"apis":[{"path":"/pets","schema":"{path}.{format}","format":["json","yaml"]},{"path":"/products","schema":"{path}.{format}","format":["json"]}],"basePath":%q}`, server.URL+"/v1")
		case "/v2":
			fmt.Fprintf(writer, `{"apis":[{"path":"/pets","schema":"{path}.{format}","format":["json"]}],"basePath":%q}`, server.URL+"/v2")
		case "/v1/pets.json", "/v2/pets.json":
			if request.URL.Query().Get("format") != "openapi3" {
				http.Error(writer, "missing OpenAPI format", http.StatusBadRequest)
				return
			}
			fmt.Fprintf(writer, `{"openapi":"3.0.3","info":{"title":"OVH Pets","version":"1.0"},"paths":{"/pets":{"get":{"summary":%q,"responses":{"200":{"description":"ok"}}}}},"components":{"schemas":{"Pet":{"type":"object"}}}}`, request.URL.Path)
		case "/v1/products.json":
			fmt.Fprint(writer, `{"openapi":"3.0.3","info":{"title":"OVH Products","version":"1.0"},"paths":{}}`)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	root := t.TempDir()
	index := filepath.Join(t.TempDir(), "index.sqlite")
	detection, err := importer.DetectURL(context.Background(), server.URL+"/console/", importer.Options{
		LibraryRoot: root, Rebuild: rebuildFunc(root, index), HTTPClient: server.Client(),
	})
	if err != nil || detection.Engine != "openapi" || detection.Framework != "rapidoc" {
		t.Fatalf("catalog detection: %+v, %v", detection, err)
	}
	result, err := importer.ImportOpenAPI(context.Background(), "OVHcloud API", "live", server.URL+"/console/", importer.Options{
		LibraryRoot: root, Rebuild: rebuildFunc(root, index), HTTPClient: server.Client(),
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Kind != "openapi-catalog" || result.Sources != 3 || result.Pages != 8 {
		t.Fatalf("unexpected result: %+v", result)
	}
	for _, relative := range []string{
		"apis/v1/pets/operations/untagged", "apis/v1/pets/schemas",
		"apis/v2/pets/operations/untagged", "apis/v2/pets/schemas",
	} {
		files, globErr := filepath.Glob(filepath.Join(result.Destination, relative, "*.md"))
		if globErr != nil || len(files) != 1 {
			t.Fatalf("generated %s: %v, %v", relative, files, globErr)
		}
	}
	if requested["/v1/pets.json?format=openapi3"] != 1 || requested["/v1/products.json?format=openapi3"] != 1 || requested["/v2/pets.json?format=openapi3"] != 1 {
		t.Fatalf("catalog schema requests: %+v", requested)
	}
}

func TestImportOpenAPICatalogFailurePublishesNothing(t *testing.T) {
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/console/":
			fmt.Fprintf(writer, `<html><body><rapi-doc spec-roots="%s/v1"></rapi-doc></body></html>`, server.URL)
		case "/v1":
			fmt.Fprintf(writer, `{"apis":[{"path":"/first","schema":"{path}.{format}","format":["json"]},{"path":"/missing","schema":"{path}.{format}","format":["json"]}],"basePath":%q}`, server.URL+"/v1")
		case "/v1/first.json":
			fmt.Fprint(writer, `{"openapi":"3.0.3","info":{"title":"First","version":"v1"},"paths":{"/first":{"get":{"responses":{"200":{"description":"ok"}}}}}}`)
		case "/v1/missing.json":
			http.Error(writer, "unavailable", http.StatusServiceUnavailable)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	root := t.TempDir()
	rebuilds := 0
	_, err := importer.ImportOpenAPI(context.Background(), "Incomplete Catalog", "live", server.URL+"/console/", importer.Options{
		LibraryRoot: root, HTTPClient: server.Client(), Rebuild: func(context.Context) error { rebuilds++; return nil },
	})
	if err == nil || !strings.Contains(err.Error(), "missing") || !strings.Contains(err.Error(), "HTTP 503") {
		t.Fatalf("expected catalog failure, got %v", err)
	}
	if rebuilds != 0 {
		t.Fatalf("failed catalog rebuilt index %d times", rebuilds)
	}
	if _, statErr := os.Stat(filepath.Join(root, "incomplete-catalog", "live")); !os.IsNotExist(statErr) {
		t.Fatalf("failed catalog destination exists: %v", statErr)
	}
}

func TestImportOpenAPICatalogRejectsCrossOriginRedirect(t *testing.T) {
	externalRequests := 0
	external := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		externalRequests++
		fmt.Fprint(writer, `{"openapi":"3.0.3","paths":{}}`)
	}))
	defer external.Close()
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/console/":
			fmt.Fprintf(writer, `<html><body><rapi-doc spec-roots="%s/v1"></rapi-doc></body></html>`, server.URL)
		case "/v1":
			fmt.Fprintf(writer, `{"apis":[{"path":"/redirect","schema":"{path}.{format}","format":["json"]}],"basePath":%q}`, server.URL+"/v1")
		case "/v1/redirect.json":
			http.Redirect(writer, request, external.URL+"/openapi.json", http.StatusFound)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	root := t.TempDir()
	_, err := importer.ImportOpenAPI(context.Background(), "Redirect Catalog", "live", server.URL+"/console/", importer.Options{
		LibraryRoot: root, HTTPClient: server.Client(), Rebuild: func(context.Context) error { return nil },
	})
	if err == nil || !strings.Contains(err.Error(), "redirect changes origin") {
		t.Fatalf("expected cross-origin redirect error, got %v", err)
	}
	if externalRequests != 0 {
		t.Fatalf("cross-origin redirect target received %d requests", externalRequests)
	}
	if _, statErr := os.Stat(filepath.Join(root, "redirect-catalog", "live")); !os.IsNotExist(statErr) {
		t.Fatalf("redirect catalog destination exists: %v", statErr)
	}
}

func TestRuntimeOnlyRapiDocIsDetectedButNotClaimedComplete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(writer, `<html><body><rapi-doc id="docs"></rapi-doc><script>document.getElementById("docs").loadSpec({openapi:"3.0.3",paths:{}})</script></body></html>`)
	}))
	defer server.Close()
	root := t.TempDir()
	options := importer.Options{LibraryRoot: root, Rebuild: func(context.Context) error { return nil }, HTTPClient: server.Client()}
	detection, err := importer.DetectURL(context.Background(), server.URL, options)
	if err != nil || detection.Engine != "openapi" || detection.Framework != "rapidoc" {
		t.Fatalf("runtime RapiDoc detection: %+v, %v", detection, err)
	}
	_, err = importer.ImportOpenAPI(context.Background(), "Runtime API", "live", server.URL, options)
	if err == nil || !strings.Contains(err.Error(), "no discoverable specification URL") {
		t.Fatalf("expected explicit nondiscoverable error, got %v", err)
	}
	if _, statErr := os.Stat(filepath.Join(root, "runtime-api", "live")); !os.IsNotExist(statErr) {
		t.Fatalf("runtime-only RapiDoc destination exists: %v", statErr)
	}
}

func TestImportOpenAPIDiscoversExternalSwaggerInitializer(t *testing.T) {
	validatorRequests := 0
	wrongRequests := 0
	wrong := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		wrongRequests++
		fmt.Fprint(writer, `{"openapi":"3.0.3","info":{"title":"Wrong","version":"v1"},"paths":{"/wrong":{"get":{"responses":{"200":{"description":"wrong"}}}}}}`)
	}))
	defer wrong.Close()
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/docs/":
			fmt.Fprint(writer, `<html><body><div id="swagger-ui"></div><script src="./swagger-initializer.js"></script></body></html>`)
		case "/docs/swagger-initializer.js":
			fmt.Fprintf(writer, `const defaultDefinitionUrl = %q;
const serviceDefinitions = "./openapi.json";
const definitionURL = serviceDefinitions;
window.ui = SwaggerUIBundle({url: definitionURL, validatorUrl: "./validator"});`, wrong.URL+"/wrong.json")
		case "/docs/openapi.json":
			fmt.Fprint(writer, `{"openapi":"3.0.3","info":{"title":"External","version":"v1"},"paths":{"/external":{"get":{"responses":{"200":{"description":"ok"}}}}}}`)
		case "/docs/validator":
			validatorRequests++
			http.Error(writer, "not a specification", http.StatusBadRequest)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	root := t.TempDir()
	result, err := importer.ImportOpenAPI(context.Background(), "External Config", "v1", server.URL+"/docs/", importer.Options{
		LibraryRoot: root, Rebuild: rebuildFunc(root, filepath.Join(t.TempDir(), "index.sqlite")), HTTPClient: server.Client(),
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Kind != "openapi" || result.Source != server.URL+"/docs/openapi.json" || result.Pages != 2 {
		t.Fatalf("unexpected result: %+v", result)
	}
	if validatorRequests != 0 || wrongRequests != 0 {
		t.Fatalf("non-specification candidates were fetched: validator=%d wrong_default=%d", validatorRequests, wrongRequests)
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

func TestImportHTMLDetectsAndScrapesDocusaurus(t *testing.T) {
	blogRequests := 0
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/html")
		switch request.URL.Path {
		case "/docs":
			fmt.Fprint(writer, `<!doctype html><html class="plugin-docs docs-doc-page"><head>
<meta name=generator content="Docusaurus v3.10.1"><title>Site chrome title</title></head><body>
<nav><a href=/blog>Blog</a></nav>
<aside><ul class=theme-doc-sidebar-menu><li><a href=/docs/second?view=full>Second guide</a><li><a href=/docs>First guide</a></ul></aside>
<div class="theme-doc-markdown markdown"><header><h1>First Guide<a class=hash-link href=#first>​</a></h1></header>
<p>Framework-specific content.</p><pre class=language-go><code><span>first line</span><br><span>second line</span></code></pre>
<div hidden><p>Hidden duplicate content.</p></div></div>
<aside><p>On this page duplicate.</p></aside><nav class=pagination-nav><a href=/docs/second?view=full>Next</a></nav></body></html>`)
		case "/docs/second":
			if request.URL.Query().Get("view") != "full" {
				http.NotFound(writer, request)
				return
			}
			fmt.Fprint(writer, `<!doctype html><html><head><meta name=generator content="Docusaurus v3.10.1"></head><body>
<div role=banner>Announcement chrome.</div>
<ul class=theme-doc-sidebar-menu><li><a href=/docs>First guide</a></ul>
<div class=generatedIndexPage_fixture><header><h1>Second Guide</h1></header><p>Independent second page.</p></div></body></html>`)
		case "/blog":
			blogRequests++
			fmt.Fprint(writer, `<html><body><h1>Blog</h1></body></html>`)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	root := t.TempDir()
	options := importer.Options{
		LibraryRoot: root, Rebuild: rebuildFunc(root, filepath.Join(t.TempDir(), "index.sqlite")), HTTPClient: server.Client(),
		MaxHTMLPages: 5, MaxHTMLDepth: 2,
	}
	detection, err := importer.DetectURL(context.Background(), server.URL+"/docs", options)
	if err != nil {
		t.Fatal(err)
	}
	if detection.Engine != "html" || detection.Framework != "docusaurus" {
		t.Fatalf("unexpected detection: %+v", detection)
	}
	result, err := importer.ImportHTML(context.Background(), "Docusaurus Fixture", "v1", server.URL+"/docs", options)
	if err != nil {
		t.Fatal(err)
	}
	if result.Framework != "docusaurus" || result.Pages != 2 || blogRequests != 0 {
		t.Fatalf("unexpected import: result=%+v blog_requests=%d", result, blogRequests)
	}
	var generated string
	for _, name := range relativeFiles(t, result.Destination) {
		if name == "_index.md" {
			continue
		}
		raw, readErr := os.ReadFile(filepath.Join(result.Destination, name))
		if readErr != nil {
			t.Fatal(readErr)
		}
		generated += string(raw)
	}
	for _, wanted := range []string{"# First Guide", "Framework-specific content.", "```go\nfirst line\nsecond line\n```", "# Second Guide", "Independent second page."} {
		if !strings.Contains(generated, wanted) {
			t.Errorf("generated Docusaurus Markdown missing %q:\n%s", wanted, generated)
		}
	}
	for _, unwanted := range []string{"Site chrome title", "Blog", "On this page duplicate.", "Hidden duplicate content.", "Announcement chrome.", "hash-link"} {
		if strings.Contains(generated, unwanted) {
			t.Errorf("generated Docusaurus Markdown contains %q:\n%s", unwanted, generated)
		}
	}
}

func TestImportHTMLDetectsAndScrapesMkDocsMaterialSitemap(t *testing.T) {
	secondRequests := 0
	aliasRequests := 0
	excludedRequests := 0
	var server *httptest.Server
	materialPage := func(title, content string) string {
		return `<!doctype html><html><head><meta name="generator" content="mkdocs-1.6.1, mkdocs-material-9.7.6"><title>Chrome</title></head><body>` +
			`<header><a href="/blog/">Blog chrome</a></header>` +
			`<nav class="md-nav md-nav--primary" data-md-level="0"><a href="/docs/getting-started/">Start</a><a href="/docs/second/">Second</a>` +
			`<nav class="md-nav md-nav--secondary"><a href="/docs/not-in-sitemap/">TOC</a></nav></nav>` +
			`<article class="md-content__inner md-typeset language-french"><h1>` + title + `<a class="headerlink" href="#title">¶</a></h1>` + content + `</article>` +
			`<footer>Footer chrome</footer></body></html>`
	}
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/docs/getting-started/":
			fmt.Fprint(writer, materialPage("Getting Started", `<script>const example = '<meta http-equiv="refresh" content="0;url=/docs/second/">'</script><!-- <meta http-equiv="refresh" content="0;url=/docs/second/"> --><p>Material content.</p><div class="language-python highlight"><pre><code>print("material")</code></pre></div><pre><code>plain()</code></pre><p><a href="/docs/not-in-sitemap/">Excluded content link</a></p>`))
		case "/docs/second/":
			secondRequests++
			fmt.Fprint(writer, materialPage(`<a class="toclink" href="#second">Second Guide</a>`, `<p>Independent second page.</p>`))
		case "/docs/alias/":
			aliasRequests++
			fmt.Fprint(writer, `<!doctype html><html><head><noscript><meta http-equiv="refresh" content="0;url=../second/"></noscript></head><body></body></html>`)
		case "/docs/sitemap.xml":
			writer.Header().Set("Content-Type", "application/xml")
			fmt.Fprintf(writer, `<?xml version="1.0"?><urlset><url><loc>%s/docs/getting-started/</loc></url><url><loc>%s/docs/alias/</loc></url><url><loc>%s/docs/second/</loc></url></urlset>`, server.URL, server.URL, server.URL)
		case "/docs/not-in-sitemap/", "/blog/":
			excludedRequests++
			http.Error(writer, "not documentation inventory", http.StatusInternalServerError)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	root := t.TempDir()
	index := filepath.Join(t.TempDir(), "index.sqlite")
	options := importer.Options{
		LibraryRoot: root, Rebuild: rebuildFunc(root, index), HTTPClient: server.Client(),
		HTMLScope: "path", MaxHTMLPages: -1, MaxHTMLDepth: -1,
	}
	detection, err := importer.DetectURL(context.Background(), server.URL+"/docs/getting-started/", options)
	if err != nil || detection.Engine != "html" || detection.Framework != "mkdocs-material" {
		t.Fatalf("unexpected detection: %+v, %v", detection, err)
	}
	result, err := importer.ImportHTML(context.Background(), "Material Fixture", "v1", server.URL+"/docs/getting-started/", options)
	if err != nil {
		t.Fatal(err)
	}
	if result.Framework != "mkdocs-material" || result.Pages != 2 || result.Truncated || secondRequests != 1 || aliasRequests != 1 || excludedRequests != 0 {
		t.Fatalf("unexpected import: result=%+v second=%d alias=%d excluded=%d", result, secondRequests, aliasRequests, excludedRequests)
	}
	var generated string
	for _, name := range relativeFiles(t, result.Destination) {
		if filepath.Ext(name) != ".md" || name == "_index.md" {
			continue
		}
		raw, readErr := os.ReadFile(filepath.Join(result.Destination, name))
		if readErr != nil {
			t.Fatal(readErr)
		}
		generated += string(raw)
	}
	for _, wanted := range []string{"# Getting Started", "Material content.", "```python\nprint(\"material\")\n```", "```\nplain()\n```", "# Second Guide", "Independent second page."} {
		if !strings.Contains(generated, wanted) {
			t.Errorf("generated Material Markdown missing %q:\n%s", wanted, generated)
		}
	}
	for _, unwanted := range []string{"¶", "Blog chrome", "TOC", "Footer chrome", "[Second Guide](", "```french"} {
		if strings.Contains(generated, unwanted) {
			t.Errorf("generated Material Markdown contains %q:\n%s", unwanted, generated)
		}
	}
	snapshot, err := library.Open(context.Background(), library.Options{UserRoot: root, IndexPath: index})
	if err != nil {
		t.Fatal(err)
	}
	defer snapshot.Close()
	search, err := snapshot.Search(context.Background(), library.SearchRequest{DocID: "material-fixture-v1", Query: "material"})
	if err != nil || search.Total == 0 {
		t.Fatalf("indexed Material search: %+v, %v", search, err)
	}
}

func TestImportHTMLMkDocsMaterialFailsIncompleteSitemap(t *testing.T) {
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/docs/start/":
			fmt.Fprint(writer, `<!doctype html><html><head><meta name="generator" content="mkdocs-1.6.1, mkdocs-material-9.7.6"></head><body><nav class="md-nav--primary"><a href="/docs/">Docs</a></nav><article class="md-content__inner"><h1>Start</h1></article></body></html>`)
		case "/docs/sitemap.xml":
			fmt.Fprintf(writer, `<urlset><url><loc>%s/docs/start/</loc></url><url><loc>%s/docs/missing/</loc></url></urlset>`, server.URL, server.URL)
		case "/docs/missing/":
			http.Error(writer, "missing", http.StatusServiceUnavailable)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	root := t.TempDir()
	_, err := importer.ImportHTML(context.Background(), "Incomplete Material", "v1", server.URL+"/docs/start/", importer.Options{
		LibraryRoot: root, Rebuild: func(context.Context) error { return nil }, HTTPClient: server.Client(),
		HTMLScope: "path", MaxHTMLPages: -1, MaxHTMLDepth: -1,
	})
	if err == nil || !strings.Contains(err.Error(), "mkdocs-material crawl did not complete") || !strings.Contains(err.Error(), "HTTP 503") {
		t.Fatalf("expected incomplete Material failure, got %v", err)
	}
	if _, statErr := os.Stat(filepath.Join(root, "incomplete-material", "v1")); !os.IsNotExist(statErr) {
		t.Fatalf("incomplete Material destination exists: %v", statErr)
	}
}

func TestImportHTMLMkDocsMaterialBoundsRefreshAliases(t *testing.T) {
	aliasRequests := 0
	targetRequests := 0
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/docs/start/":
			fmt.Fprint(writer, `<!doctype html><html><head><meta name="generator" content="mkdocs-1.6.1, mkdocs-material-9.7.6"></head><body><nav class="md-nav--primary"><a href="/docs/start/">Start</a><a href="/docs/target/">Target</a></nav><article class="md-content__inner"><h1>Start</h1></article></body></html>`)
		case "/docs/sitemap.xml":
			fmt.Fprintf(writer, `<urlset><url><loc>%s/docs/start/</loc></url><url><loc>%s/docs/alias/</loc></url><url><loc>%s/docs/target/</loc></url></urlset>`, server.URL, server.URL, server.URL)
		case "/docs/alias/":
			aliasRequests++
			fmt.Fprint(writer, `<html><head><meta http-equiv="refresh" content="0;url=../target/"></head></html>`)
		case "/docs/target/":
			targetRequests++
			fmt.Fprint(writer, `<!doctype html><html><head><meta name="generator" content="mkdocs-1.6.1, mkdocs-material-9.7.6"></head><body><nav class="md-nav--primary"><a href="/docs/start/">Start</a><a href="/docs/target/">Target</a></nav><article class="md-content__inner"><h1>Target</h1></article></body></html>`)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	root := t.TempDir()
	result, err := importer.ImportHTML(context.Background(), "Bounded Material", "v1", server.URL+"/docs/start/", importer.Options{
		LibraryRoot: root, Rebuild: rebuildFunc(root, filepath.Join(t.TempDir(), "index.sqlite")), HTTPClient: server.Client(),
		HTMLScope: "path", MaxHTMLPages: 2, MaxHTMLDepth: -1,
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Pages != 1 || !result.Truncated || aliasRequests != 1 || targetRequests != 0 {
		t.Fatalf("alias limit bypass: result=%+v alias=%d target=%d", result, aliasRequests, targetRequests)
	}
}

func TestImportHTMLMkDocsMaterialRejectsHTTPRedirectOutsideInventory(t *testing.T) {
	unlistedRequests := 0
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/docs/start/":
			fmt.Fprint(writer, `<!doctype html><html><head><meta name="generator" content="mkdocs-1.6.1, mkdocs-material-9.7.6"></head><body><nav class="md-nav--primary"><a href="/docs/start/">Start</a><a href="/docs/advertised/">Advertised</a></nav><article class="md-content__inner"><h1>Start</h1></article></body></html>`)
		case "/docs/sitemap.xml":
			fmt.Fprintf(writer, `<urlset><url><loc>%s/docs/start/</loc></url><url><loc>%s/docs/advertised/</loc></url><url><loc>%s/docs/final/</loc></url></urlset>`, server.URL, server.URL, server.URL)
		case "/docs/advertised/":
			http.Redirect(writer, request, "/docs/unlisted/", http.StatusFound)
		case "/docs/unlisted/":
			unlistedRequests++
			http.Redirect(writer, request, "/docs/final/", http.StatusFound)
		case "/docs/final/":
			fmt.Fprint(writer, `<!doctype html><html><head><meta name="generator" content="mkdocs-1.6.1, mkdocs-material-9.7.6"></head><body><nav class="md-nav--primary"><a href="/docs/start/">Start</a><a href="/docs/final/">Final</a></nav><article class="md-content__inner"><h1>Final</h1></article></body></html>`)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	root := t.TempDir()
	_, err := importer.ImportHTML(context.Background(), "Redirected Material", "v1", server.URL+"/docs/start/", importer.Options{
		LibraryRoot: root, Rebuild: func(context.Context) error { return nil }, HTTPClient: server.Client(),
		HTMLScope: "path", MaxHTMLPages: -1, MaxHTMLDepth: -1,
	})
	if err == nil || !strings.Contains(err.Error(), "redirect crosses the finite inventory") {
		t.Fatalf("expected finite inventory redirect rejection, got %v", err)
	}
	if unlistedRequests != 0 {
		t.Fatalf("unlisted intermediate redirect received %d requests", unlistedRequests)
	}
	if _, statErr := os.Stat(filepath.Join(root, "redirected-material", "v1")); !os.IsNotExist(statErr) {
		t.Fatalf("redirected Material destination exists: %v", statErr)
	}
}

func TestImportHTMLMkDocsMaterialRejectsRefreshCycle(t *testing.T) {
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/docs/start/":
			fmt.Fprint(writer, `<!doctype html><html><head><meta name="generator" content="mkdocs-1.6.1, mkdocs-material-9.7.6"></head><body><nav class="md-nav--primary"><a href="/docs/start/">Start</a><a href="/docs/a/">A</a><a href="/docs/b/">B</a></nav><article class="md-content__inner"><h1>Start</h1></article></body></html>`)
		case "/docs/sitemap.xml":
			fmt.Fprintf(writer, `<urlset><url><loc>%s/docs/start/</loc></url><url><loc>%s/docs/a/</loc></url><url><loc>%s/docs/b/</loc></url></urlset>`, server.URL, server.URL, server.URL)
		case "/docs/a/":
			fmt.Fprint(writer, `<html><head><meta http-equiv="refresh" content="0;url=../b/"></head></html>`)
		case "/docs/b/":
			fmt.Fprint(writer, `<html><head><meta http-equiv="refresh" content="0;url=../a/"></head></html>`)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	root := t.TempDir()
	_, err := importer.ImportHTML(context.Background(), "Cyclic Material", "v1", server.URL+"/docs/start/", importer.Options{
		LibraryRoot: root, Rebuild: func(context.Context) error { return nil }, HTTPClient: server.Client(),
		HTMLScope: "path", MaxHTMLPages: -1, MaxHTMLDepth: -1,
	})
	if err == nil || !strings.Contains(err.Error(), "inventory alias cycle") {
		t.Fatalf("expected refresh cycle rejection, got %v", err)
	}
	if _, statErr := os.Stat(filepath.Join(root, "cyclic-material", "v1")); !os.IsNotExist(statErr) {
		t.Fatalf("cyclic Material destination exists: %v", statErr)
	}
}

func TestImportHTMLMkDocsMaterialUsesScopeScript(t *testing.T) {
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/nested/guide/start/":
			fmt.Fprint(writer, `<!doctype html><html><head><meta name="generator" content="mkdocs-1.6.1, mkdocs-material-9.7.6"><script>__md_scope=new URL("../..",location)</script></head><body><nav class="md-nav--primary"><a href="/nested/guide/start/">Start</a></nav><article class="md-content__inner"><h1>Nested Start</h1></article></body></html>`)
		case "/nested/other/":
			fmt.Fprint(writer, `<!doctype html><html><head><meta name="generator" content="mkdocs-1.6.1, mkdocs-material-9.7.6"></head><body><nav class="md-nav--primary"><a href="/nested/guide/start/">Start</a></nav><article class="md-content__inner"><h1>Other</h1></article></body></html>`)
		case "/nested/sitemap.xml":
			fmt.Fprintf(writer, `<urlset><url><loc>%s/nested/guide/start/</loc></url><url><loc>%s/nested/other/</loc></url></urlset>`, server.URL, server.URL)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	root := t.TempDir()
	result, err := importer.ImportHTML(context.Background(), "Scoped Material", "v1", server.URL+"/nested/guide/start/", importer.Options{
		LibraryRoot: root, Rebuild: rebuildFunc(root, filepath.Join(t.TempDir(), "index.sqlite")), HTTPClient: server.Client(),
		HTMLScope: "path", MaxHTMLPages: -1, MaxHTMLDepth: -1,
	})
	if err != nil || result.Pages != 2 || result.Truncated {
		t.Fatalf("scope script import: %+v, %v", result, err)
	}
}

func TestDetectHTMLDoesNotMatchMaterialSubstring(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(writer, `<!doctype html><html><head><meta name="generator" content="not-mkdocs-materialized"></head><body><main><h1>Generic</h1></main></body></html>`)
	}))
	defer server.Close()
	detection, err := importer.DetectURL(context.Background(), server.URL, importer.Options{
		LibraryRoot: t.TempDir(), Rebuild: func(context.Context) error { return nil }, HTTPClient: server.Client(),
	})
	if err != nil || detection.Framework != "unknown" {
		t.Fatalf("unexpected substring detection: %+v, %v", detection, err)
	}
}

func TestImportHTMLDetectsMkDocsBuiltInThemes(t *testing.T) {
	for _, test := range []struct {
		name        string
		startPath   string
		secondPath  string
		sitemapPath string
		page        func(string, string) string
		sitemap     func(string) string
	}{
		{
			name: "default theme", startPath: "/docs/start/", secondPath: "/docs/second/", sitemapPath: "/docs/sitemap.xml",
			page: func(title, content string) string {
				return `<!doctype html><html><head><script>var base_url = "/";</script><link href="../css/base.css" rel="stylesheet"><script>var base_url = "..";</script><script src="../js/base.js"></script></head><body><div id="navbar-collapse"><a href="/docs/start/">Start</a><a href="/docs/second/">Second</a></div><div class="col-md-9" role="main"><h1>` + title + `<a class="headerlink" href="#title">¶</a></h1>` + content + `</div></body></html>`
			},
			sitemap: func(serverURL string) string {
				return fmt.Sprintf(`<urlset><url><loc>%s/docs/start/</loc></url><url><loc>%s/docs/second/</loc></url></urlset>`, serverURL, serverURL)
			},
		},
		{
			name: "readthedocs theme", startPath: "/", secondPath: "/second/", sitemapPath: "/sitemap.xml",
			page: func(title, content string) string {
				return `<!doctype html><html><head><script>var mkdocs_page_name = "Page"; var base_url = ".";</script><script src="js/theme.js"></script></head><body><div class="wy-menu wy-menu-vertical"><a href=".">Home</a><a href="second/">Second</a></div><div class="wy-nav-content"><div class="rst-content"><div role="main" class="document"><div class="section"><h1>` + title + `<a class="headerlink" href="#title">¶</a></h1>` + content + `</div></div></div></div></body></html>`
			},
			sitemap: func(serverURL string) string {
				return fmt.Sprintf(`<urlset><url><loc>%s/</loc></url><url><loc>%s/second/</loc></url></urlset>`, serverURL, serverURL)
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			secondRequests := 0
			var server *httptest.Server
			server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				switch request.URL.Path {
				case test.startPath:
					fmt.Fprint(writer, test.page("MkDocs Start", `<p>Built-in theme content.</p><pre><code class="language-python">print("mkdocs")</code></pre>`))
				case test.secondPath:
					secondRequests++
					fmt.Fprint(writer, test.page("MkDocs Second", `<p>Independent navigation page.</p>`))
				case test.sitemapPath:
					fmt.Fprint(writer, test.sitemap(server.URL))
				default:
					http.NotFound(writer, request)
				}
			}))
			defer server.Close()

			root := t.TempDir()
			index := filepath.Join(t.TempDir(), "index.sqlite")
			options := importer.Options{
				LibraryRoot: root, Rebuild: rebuildFunc(root, index), HTTPClient: server.Client(),
				HTMLScope: "path", MaxHTMLPages: -1, MaxHTMLDepth: -1,
			}
			source := server.URL + test.startPath + "?utm_source=fixture"
			detection, err := importer.DetectURL(context.Background(), source, options)
			if err != nil || detection.Framework != "mkdocs" {
				t.Fatalf("MkDocs detection: %+v, %v", detection, err)
			}
			result, err := importer.ImportHTML(context.Background(), "MkDocs Fixture", "v1-"+strings.ReplaceAll(test.name, " ", "-"), source, options)
			if err != nil {
				t.Fatal(err)
			}
			if result.Framework != "mkdocs" || result.Pages != 2 || result.Truncated || secondRequests != 1 {
				t.Fatalf("MkDocs import: %+v, second=%d", result, secondRequests)
			}
			var generated string
			for _, name := range relativeFiles(t, result.Destination) {
				if filepath.Ext(name) == ".md" && name != "_index.md" {
					raw, readErr := os.ReadFile(filepath.Join(result.Destination, name))
					if readErr != nil {
						t.Fatal(readErr)
					}
					generated += string(raw)
				}
			}
			for _, wanted := range []string{"# MkDocs Start", "Built-in theme content.", "```python\nprint(\"mkdocs\")\n```", "# MkDocs Second"} {
				if !strings.Contains(generated, wanted) {
					t.Errorf("generated MkDocs Markdown missing %q:\n%s", wanted, generated)
				}
			}
			if strings.Contains(generated, "¶") {
				t.Fatalf("heading permalink leaked:\n%s", generated)
			}
		})
	}
}

func TestImportHTMLRejectsEmptyMkDocsSitemap(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/":
			fmt.Fprint(writer, `<!doctype html><html><head><link href="css/base.css" rel="stylesheet"><script>var base_url = ".";</script><script src="js/base.js"></script></head><body><div id="navbar-collapse"><a href="/">Home</a></div><main role="main"><h1>Home</h1></main></body></html>`)
		case "/sitemap.xml":
			fmt.Fprint(writer, `<urlset></urlset>`)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	_, err := importer.ImportHTML(context.Background(), "Incomplete MkDocs", "v1", server.URL, importer.Options{
		LibraryRoot: t.TempDir(), Rebuild: func(context.Context) error { return nil }, HTTPClient: server.Client(),
		HTMLScope: "path", MaxHTMLPages: -1, MaxHTMLDepth: -1,
	})
	if err == nil || !strings.Contains(err.Error(), "contains no URLs") {
		t.Fatalf("expected empty sitemap rejection, got %v", err)
	}
}

func TestDetectHTMLRecognizesMkDocsMetadata(t *testing.T) {
	for _, test := range []struct {
		name      string
		html      string
		framework string
	}{
		{name: "generator", html: `<meta name="generator" content="mkdocs-1.6.1"><main><h1>Guide</h1></main>`, framework: "mkdocs"},
		{name: "build comment", html: `<main role="main"><h1>Guide</h1></main><!-- MkDocs version : 1.6.1
Docs Build Date UTC : 2026-07-27 00:00:00 -->`, framework: "mkdocs"},
		{name: "quoted comment", html: `<main role="main"><h1>Guide</h1></main><!-- MkDocs version : 1.6.1 -->`, framework: "unknown"},
		{name: "malformed generator", html: `<meta name="generator" content="mkdocs-1anything"><main><h1>Guide</h1></main>`, framework: "unknown"},
	} {
		t.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
				fmt.Fprint(writer, `<!doctype html><html><body>`+test.html+`</body></html>`)
			}))
			defer server.Close()
			detection, err := importer.DetectURL(context.Background(), server.URL, importer.Options{
				LibraryRoot: t.TempDir(), Rebuild: func(context.Context) error { return nil }, HTTPClient: server.Client(),
			})
			if err != nil || detection.Framework != test.framework {
				t.Fatalf("detection: %+v, %v", detection, err)
			}
		})
	}
}

func TestImportHTMLExhaustsSphinxSearchIndex(t *testing.T) {
	for _, test := range []struct {
		name       string
		root       string
		startPath  string
		aliasPath  string
		builder    string
		fileSuffix string
		guidePath  string
		orphanPath string
		theme      string
	}{
		{name: "html builder", root: "/docs/", startPath: "/docs/", aliasPath: "/docs/index.xhtml", builder: "html", fileSuffix: ".xhtml", guidePath: "/docs/guide.xhtml", orphanPath: "/docs/orphan.xhtml", theme: "basic"},
		{name: "dirhtml builder", root: "/manual/", startPath: "/manual/index.html", aliasPath: "/manual/", builder: "dirhtml", fileSuffix: ".html", guidePath: "/manual/guide/", orphanPath: "/manual/orphan/", theme: "readthedocs"},
	} {
		t.Run(test.name, func(t *testing.T) {
			orphanRequests := 0
			indexAliasRequests := 0
			page := func(title, body, contentRoot string) string {
				content := `<div class="body" role="main"><h1>` + title + `<a class="headerlink" href="#title">¶</a></h1>` + body + `</div>`
				navigation := `<div class="sphinxsidebar"><a href="guide">Guide</a><span>Sidebar chrome</span></div>`
				if test.theme == "readthedocs" {
					content = `<div class="rst-content"><div class="document" role="main"><h1>` + title + `<a class="headerlink" href="#title">¶</a></h1>` + body + `</div></div>`
					navigation = `<div class="wy-menu-vertical"><a href="guide">Guide</a><span>RTD chrome</span></div>`
				}
				return `<!doctype html><html data-content_root="` + contentRoot + `"><head><script src="` + test.root + `_static/documentation_options.js?v=1"></script><script src="` + test.root + `_static/doctools.js?v=1"></script></head><body>` + navigation + content + `</body></html>`
			}

			server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				if request.URL.Path == test.startPath {
					fmt.Fprint(writer, page("Sphinx Home", `<p>Root content.</p>`, "./"))
					return
				}
				switch request.URL.Path {
				case test.guidePath:
					fmt.Fprint(writer, page("Sphinx Guide", `<div class="highlight-python"><pre><code>print("sphinx")</code></pre></div>`, "../"))
				case test.orphanPath:
					orphanRequests++
					fmt.Fprint(writer, page("Indexed Orphan", `<dl><dt><code>request()</code></dt><dd><p>Orphan API description.</p></dd><dt class="sig">request(*args,<br> **kwargs)<a class="viewcode-link" href="source.html">[source]</a><a class="headerlink" href="#request">¶</a></dt><dd><p>Signature description.</p></dd></dl><aside class="footnote-list"><aside class="footnote" role="doc-footnote"><span class="label">[1]</span><p>Footnote definition.</p></aside></aside>`, "../"))
				case test.aliasPath:
					indexAliasRequests++
					http.Error(writer, "duplicate index fetch", http.StatusInternalServerError)
				case test.root + "_static/documentation_options.js":
					fmt.Fprintf(writer, `const DOCUMENTATION_OPTIONS = {BUILDER: '%s', FILE_SUFFIX: '%s', LINK_SUFFIX: ''};`, test.builder, test.fileSuffix)
				case test.root + "_static/doctools.js":
					fmt.Fprint(writer, `const sphinx = true;`)
				case test.root + "searchindex.js":
					fmt.Fprint(writer, `Search.setIndex({"docnames":["orphan","index","guide"],"titles":["Orphan","Home","Guide"]});`)
				default:
					http.NotFound(writer, request)
				}
			}))
			defer server.Close()

			root := t.TempDir()
			index := filepath.Join(t.TempDir(), "index.sqlite")
			options := importer.Options{
				LibraryRoot: root, Rebuild: rebuildFunc(root, index), HTTPClient: server.Client(),
				HTMLScope: "path", MaxHTMLPages: -1, MaxHTMLDepth: -1,
			}
			source := server.URL + test.startPath + "?utm_source=fixture"
			detection, err := importer.DetectURL(context.Background(), source, options)
			if err != nil || detection.Framework != "sphinx" {
				t.Fatalf("Sphinx detection: %+v, %v", detection, err)
			}
			result, err := importer.ImportHTML(context.Background(), "Sphinx Fixture", "v1-"+strings.ReplaceAll(test.name, " ", "-"), source, options)
			if err != nil {
				t.Fatal(err)
			}
			if result.Framework != "sphinx" || result.Pages != 3 || result.Truncated || orphanRequests != 1 || indexAliasRequests != 0 {
				t.Fatalf("Sphinx import: %+v, orphan=%d, index_alias=%d", result, orphanRequests, indexAliasRequests)
			}
			var generated string
			for _, name := range relativeFiles(t, result.Destination) {
				if filepath.Ext(name) != ".md" || name == "_index.md" {
					continue
				}
				raw, readErr := os.ReadFile(filepath.Join(result.Destination, name))
				if readErr != nil {
					t.Fatal(readErr)
				}
				generated += string(raw)
			}
			for _, wanted := range []string{"# Sphinx Home", "# Sphinx Guide", "# Indexed Orphan", "```python\nprint(\"sphinx\")\n```", "**`request()`**", "`request(*args, **kwargs)`", "Orphan API description.", "**[1]**", "Footnote definition."} {
				if !strings.Contains(generated, wanted) {
					t.Errorf("generated Sphinx Markdown missing %q:\n%s", wanted, generated)
				}
			}
			for _, excluded := range []string{"Sidebar chrome", "RTD chrome", "¶"} {
				if strings.Contains(generated, excluded) {
					t.Errorf("generated Sphinx Markdown contains chrome %q", excluded)
				}
			}
		})
	}
}

func TestImportHTMLRejectsInvalidSphinxSearchIndex(t *testing.T) {
	for index, payload := range []string{
		`Search.setIndex({"docnames":[],"titles":[]});`,
		`Search.setIndex({objects:{},filenames:["index"]});`,
		`Search.setIndex({"docnames":["index","../escape"],"titles":["Home","Escape"]});`,
		`Search.setIndex({"docnames":["index","guide","guide/index"],"titles":["Home","Guide","Guide Index"]});`,
	} {
		t.Run(fmt.Sprintf("case-%d", index), func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				switch request.URL.Path {
				case "/docs/":
					fmt.Fprint(writer, `<!doctype html><html data-content_root="./"><head><script src="_static/documentation_options.js"></script><script src="_static/doctools.js"></script></head><body><main role="main"><h1>Docs</h1></main></body></html>`)
				case "/docs/_static/documentation_options.js":
					fmt.Fprint(writer, `const DOCUMENTATION_OPTIONS = {BUILDER: 'dirhtml', FILE_SUFFIX: '.html', LINK_SUFFIX: '.html'};`)
				case "/docs/searchindex.js":
					fmt.Fprint(writer, payload)
				default:
					http.NotFound(writer, request)
				}
			}))
			defer server.Close()

			_, err := importer.ImportHTML(context.Background(), "Invalid Sphinx", "v1", server.URL+"/docs/", importer.Options{
				LibraryRoot: t.TempDir(), Rebuild: func(context.Context) error { return nil }, HTTPClient: server.Client(),
				HTMLScope: "path", MaxHTMLPages: -1, MaxHTMLDepth: -1,
			})
			if err == nil || !strings.Contains(err.Error(), "sphinx crawl did not complete") {
				t.Fatalf("expected Sphinx inventory rejection, got %v", err)
			}
		})
	}
}

func TestImportHTMLFollowsInitialSphinxRefresh(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/docs/start.html":
			fmt.Fprint(writer, `<!doctype html><html><head><meta http-equiv="refresh" content="0;url=/docs/"></head></html>`)
		case "/docs/":
			fmt.Fprint(writer, `<!doctype html><html data-content_root="./"><head><script src="_static/documentation_options.js"></script><script src="_static/doctools.js"></script></head><body><main role="main"><h1>Redirected Sphinx</h1></main></body></html>`)
		case "/docs/_static/documentation_options.js":
			fmt.Fprint(writer, `const DOCUMENTATION_OPTIONS = {BUILDER: 'html', FILE_SUFFIX: '.html', LINK_SUFFIX: '.html'};`)
		case "/docs/_static/doctools.js":
			fmt.Fprint(writer, `const sphinx = true;`)
		case "/docs/searchindex.js":
			fmt.Fprint(writer, `Search.setIndex({"docnames":["index"],"titles":["Redirected Sphinx"]});`)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	root := t.TempDir()
	result, err := importer.ImportHTML(context.Background(), "Redirected Sphinx", "v1", server.URL+"/docs/start.html", importer.Options{
		LibraryRoot: root, Rebuild: rebuildFunc(root, filepath.Join(t.TempDir(), "index.sqlite")), HTTPClient: server.Client(),
		HTMLScope: "path", MaxHTMLPages: -1, MaxHTMLDepth: -1,
	})
	if err != nil || result.Framework != "sphinx" || result.Pages != 1 || result.Truncated {
		t.Fatalf("initial Sphinx refresh: %+v, %v", result, err)
	}
}

func TestImportHTMLMapsSphinxLinkSuffixAlias(t *testing.T) {
	nativeRequests := 0
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/docs/guide":
			fmt.Fprint(writer, `<!doctype html><html data-content_root="./"><head><script src="_static/documentation_options.js"></script><script src="_static/doctools.js"></script></head><body><main role="main"><h1>Extensionless Guide</h1></main></body></html>`)
		case "/docs/guide.html":
			nativeRequests++
			http.Error(writer, "duplicate native fetch", http.StatusInternalServerError)
		case "/docs/_static/documentation_options.js":
			fmt.Fprint(writer, `const DOCUMENTATION_OPTIONS = {BUILDER: 'html', FILE_SUFFIX: '.html', LINK_SUFFIX: ''};`)
		case "/docs/_static/doctools.js":
			fmt.Fprint(writer, `const sphinx = true;`)
		case "/docs/searchindex.js":
			fmt.Fprint(writer, `Search.setIndex({"docnames":["guide"],"titles":["Extensionless Guide"]});`)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	root := t.TempDir()
	result, err := importer.ImportHTML(context.Background(), "Extensionless Sphinx", "v1", server.URL+"/docs/guide?source=test", importer.Options{
		LibraryRoot: root, Rebuild: rebuildFunc(root, filepath.Join(t.TempDir(), "index.sqlite")), HTTPClient: server.Client(),
		HTMLScope: "path", MaxHTMLPages: -1, MaxHTMLDepth: -1,
	})
	if err != nil || result.Pages != 1 || result.Truncated || nativeRequests != 0 {
		t.Fatalf("Sphinx link suffix alias: %+v, native=%d, %v", result, nativeRequests, err)
	}
}

func TestImportHTMLExhaustsVitePressSitemap(t *testing.T) {
	cleanStartRequests := 0
	page := func(layout, title, body string) string {
		content := `<div class="VPDoc"><div class="vp-doc"><h1>` + title + `<a class="header-anchor" href="#title">​</a></h1>` + body + `</div></div>`
		if layout == "home" {
			content = `<div class="VPHome"><h1>` + title + `</h1>` + body + `</div>`
		} else if layout == "page" {
			content = `<div class="VPPage"><h1>` + title + `</h1>` + body + `</div>`
		}
		return `<!doctype html><html><head><meta name="generator" content="VitePress v2.0.0-alpha.18"><link rel="preload stylesheet" href="/project/docs/vp-icons.css"></head><body><header class="VPNav">Navigation chrome</header><aside class="VPSidebar">Sidebar chrome</aside>` + content + `<footer class="VPFooter">Footer chrome</footer></body></html>`
	}
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/project/docs/guide/start.html":
			fmt.Fprint(writer, page("doc", "VitePress Start", `<p>Document content.</p><div class="language-go"><pre><code>fmt.Println("vitepress")</code></pre></div>`))
		case "/project/docs/guide/start":
			cleanStartRequests++
			http.Error(writer, "duplicate clean alias", http.StatusInternalServerError)
		case "/project/docs/second.html":
			fmt.Fprint(writer, page("page", "VitePress Page", `<p>Page layout content.</p>`))
		case "/project/docs/custom.html":
			fmt.Fprint(writer, `<!doctype html><html><head><meta name="generator" content="VitePress v2.0.0-alpha.18"><link rel="stylesheet" href="/project/docs/vp-icons.css"></head><body><div class="VPContentDoc"><main><div class="vt-doc guide"><h1>Custom VitePress</h1><p>Scoped custom-theme content.</p></div></main></div></body></html>`)
		case "/project/docs/custom-home.html":
			fmt.Fprint(writer, `<!doctype html><html><head><meta name="generator" content="VitePress v2.0.0-alpha.18"><link rel="stylesheet" href="/project/docs/vp-icons.css"></head><body><div class="VPContentPage"><main><h1>Custom Home</h1><p>Scoped homepage content.</p></main><footer class="VPFooter">Custom footer chrome</footer></div></body></html>`)
		case "/project/docs/marketing.html":
			fmt.Fprint(writer, `<!doctype html><html><head><meta name="generator" content="VitePress v2.0.0-alpha.18"><link rel="stylesheet" href="/project/docs/vp-icons.css"></head><body><div class="marketing-layout"><header>Marketing nav</header><div><div><h1>Marketing Home</h1><p>Scoped marketing content.</p><footer>Marketing footer</footer></div></div></div></body></html>`)
		case "/project/docs/":
			fmt.Fprint(writer, page("home", "VitePress Home", `<p>Hero and feature content.</p>`))
		case "/project/docs/sitemap.xml":
			fmt.Fprintf(writer, `<urlset><url><loc>%s/project/docs/guide/start</loc></url><url><loc>%s/project/docs/second.html</loc></url><url><loc>%s/project/docs/custom.html</loc></url><url><loc>%s/project/docs/custom-home.html</loc></url><url><loc>%s/project/docs/marketing.html</loc></url><url><loc>%s/project/docs/</loc></url></urlset>`, server.URL, server.URL, server.URL, server.URL, server.URL, server.URL)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	root := t.TempDir()
	options := importer.Options{
		LibraryRoot: root, Rebuild: rebuildFunc(root, filepath.Join(t.TempDir(), "index.sqlite")), HTTPClient: server.Client(),
		HTMLScope: "path", MaxHTMLPages: -1, MaxHTMLDepth: -1,
	}
	source := server.URL + "/project/docs/guide/start.html?source=fixture"
	detection, err := importer.DetectURL(context.Background(), source, options)
	if err != nil || detection.Framework != "vitepress" {
		t.Fatalf("VitePress detection: %+v, %v", detection, err)
	}
	result, err := importer.ImportHTML(context.Background(), "VitePress Fixture", "v1", source, options)
	if err != nil || result.Framework != "vitepress" || result.Pages != 6 || result.Truncated || cleanStartRequests != 0 {
		t.Fatalf("VitePress import: %+v, clean_start=%d, %v", result, cleanStartRequests, err)
	}
	var generated string
	for _, name := range relativeFiles(t, result.Destination) {
		if filepath.Ext(name) == ".md" && name != "_index.md" {
			raw, readErr := os.ReadFile(filepath.Join(result.Destination, name))
			if readErr != nil {
				t.Fatal(readErr)
			}
			generated += string(raw)
		}
	}
	for _, wanted := range []string{"# VitePress Start", "```go\nfmt.Println(\"vitepress\")\n```", "# VitePress Page", "# VitePress Home", "Hero and feature content.", "# Custom VitePress", "Scoped custom-theme content.", "# Custom Home", "Scoped homepage content.", "# Marketing Home", "Scoped marketing content."} {
		if !strings.Contains(generated, wanted) {
			t.Errorf("generated VitePress Markdown missing %q:\n%s", wanted, generated)
		}
	}
	for _, excluded := range []string{"Navigation chrome", "Sidebar chrome", "Footer chrome", "Custom footer chrome", "Marketing nav", "Marketing footer", "​"} {
		if strings.Contains(generated, excluded) {
			t.Errorf("generated VitePress Markdown contains chrome %q", excluded)
		}
	}
}

func TestImportHTMLExhaustsScopedNextraSitemap(t *testing.T) {
	for _, profile := range []string{"v4", "classic"} {
		t.Run(profile, func(t *testing.T) {
			outsideRequests := 0
			page := func(title, body string) string {
				content := `<a id="nextra-skip-nav"></a><aside class="nextra-sidebar"><a href="/docs/start/">Start</a><a href="/docs/second/">Second</a><a href="/about">About</a></aside><main data-pagefind-body="true"><h1>` + title + `</h1>` + body + `</main>`
				if profile == "classic" {
					content = `<nav class="nextra-sidebar-container"><a href="/docs/start/">Start</a><a href="/docs/second/">Second</a><a href="/about">About</a></nav><article class="nextra-content"><main><h1>` + title + `</h1>` + body + `</main></article>`
				}
				return `<!doctype html><html><head><meta name="generator" content="Next.js"><script src="/_next/static/chunks/app.js"></script></head><body><header>Nextra chrome</header>` + content + `<aside class="nextra-toc">TOC chrome</aside></body></html>`
			}
			var server *httptest.Server
			server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				switch request.URL.Path {
				case "/docs/start/":
					fmt.Fprint(writer, page("Nextra Start", `<p>Scoped docs content.</p><div class="language-ts"><pre><code>const docs = true</code></pre></div>`))
				case "/docs/second/":
					fmt.Fprint(writer, page("Nextra Second", `<p>Second docs page.</p>`))
				case "/about":
					outsideRequests++
					http.Error(writer, "outside docs scope", http.StatusInternalServerError)
				case "/sitemap.xml":
					fmt.Fprintf(writer, `<urlset><url><loc>%s/about</loc></url><url><loc>%s/docs/start/</loc></url><url><loc>%s/docs/second/</loc></url></urlset>`, server.URL, server.URL, server.URL)
				default:
					http.NotFound(writer, request)
				}
			}))
			defer server.Close()

			root := t.TempDir()
			options := importer.Options{
				LibraryRoot: root, Rebuild: rebuildFunc(root, filepath.Join(t.TempDir(), "index.sqlite")), HTTPClient: server.Client(),
				HTMLScope: "path", MaxHTMLPages: -1, MaxHTMLDepth: -1,
			}
			detection, err := importer.DetectURL(context.Background(), server.URL+"/docs/start/", options)
			if err != nil || detection.Framework != "nextra" {
				t.Fatalf("Nextra detection: %+v, %v", detection, err)
			}
			result, err := importer.ImportHTML(context.Background(), "Nextra Fixture", "v1-"+profile, server.URL+"/docs/start/", options)
			if err != nil || result.Framework != "nextra" || result.Pages != 2 || result.Truncated || outsideRequests != 0 {
				t.Fatalf("Nextra import: %+v, outside=%d, %v", result, outsideRequests, err)
			}
			var generated string
			for _, name := range relativeFiles(t, result.Destination) {
				if filepath.Ext(name) == ".md" && name != "_index.md" {
					raw, readErr := os.ReadFile(filepath.Join(result.Destination, name))
					if readErr != nil {
						t.Fatal(readErr)
					}
					generated += string(raw)
				}
			}
			for _, wanted := range []string{"# Nextra Start", "Scoped docs content.", "```ts\nconst docs = true\n```", "# Nextra Second"} {
				if !strings.Contains(generated, wanted) {
					t.Errorf("generated Nextra Markdown missing %q:\n%s", wanted, generated)
				}
			}
			for _, excluded := range []string{"Nextra chrome", "TOC chrome"} {
				if strings.Contains(generated, excluded) {
					t.Errorf("generated Nextra Markdown contains chrome %q", excluded)
				}
			}
		})
	}
}

func TestImportHTMLRejectsNextraPageWithoutStaticContent(t *testing.T) {
	page := func(content string) string {
		return `<!doctype html><html><head><script src="/_next/static/chunks/app.js"></script></head><body><a id="nextra-skip-nav"></a><aside class="nextra-sidebar"><a href="/docs/start">Start</a><a href="/docs/empty">Empty</a></aside><main data-pagefind-body="true">` + content + `</main></body></html>`
	}
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/docs/start":
			fmt.Fprint(writer, page("<h1>Start</h1>"))
		case "/docs/empty":
			fmt.Fprint(writer, page(""))
		case "/sitemap.xml":
			fmt.Fprintf(writer, `<urlset><url><loc>%s/docs/start</loc></url><url><loc>%s/docs/empty</loc></url></urlset>`, server.URL, server.URL)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	root := t.TempDir()
	options := importer.Options{
		LibraryRoot: root, Rebuild: rebuildFunc(root, filepath.Join(t.TempDir(), "index.sqlite")), HTTPClient: server.Client(),
		HTMLScope: "path", MaxHTMLPages: -1, MaxHTMLDepth: -1,
	}
	_, err := importer.ImportHTML(context.Background(), "Nextra Empty", "v1", server.URL+"/docs/start", options)
	if err == nil || !strings.Contains(err.Error(), "page has no statically rendered content") {
		t.Fatalf("expected empty Nextra page rejection, got %v", err)
	}
}

func TestDetectAndImportHTMLRejectPlainText(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		writer.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(writer, "# Plain Markdown\n\nThis is not an HTML document.\n")
	}))
	defer server.Close()

	options := importer.Options{LibraryRoot: t.TempDir(), Rebuild: func(context.Context) error { return nil }, HTTPClient: server.Client()}
	if _, err := importer.DetectURL(context.Background(), server.URL, options); err == nil || !strings.Contains(err.Error(), "neither OpenAPI nor HTML") {
		t.Fatalf("expected detection rejection, got %v", err)
	}
	if _, err := importer.ImportHTML(context.Background(), "Plain", "v1", server.URL, options); err == nil || !strings.Contains(err.Error(), "not a static HTML document") {
		t.Fatalf("expected HTML rejection, got %v", err)
	}
}

func TestImportHTMLAcceptsUTF8BOM(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		_, _ = writer.Write(append([]byte{0xef, 0xbb, 0xbf}, []byte(`<!doctype html><html><body><main><h1>BOM Guide</h1><p>Valid content.</p></main></body></html>`)...))
	}))
	defer server.Close()

	root := t.TempDir()
	result, err := importer.ImportHTML(context.Background(), "BOM", "v1", server.URL, importer.Options{
		LibraryRoot: root, Rebuild: rebuildFunc(root, filepath.Join(t.TempDir(), "index.sqlite")), HTTPClient: server.Client(), MaxHTMLPages: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Pages != 1 {
		t.Fatalf("unexpected result: %+v", result)
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
		{name: "prefixed nodes", body: "prefix<html><body>" + strings.Repeat("<br>", 50_001), want: "maximum node count"},
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
