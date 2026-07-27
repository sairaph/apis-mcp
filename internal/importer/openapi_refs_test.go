package importer

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/sairaph/apis-mcp/library"
	"gopkg.in/yaml.v3"
)

func TestRapiDocExternalRefGraphGeneratesCompleteAggregate(t *testing.T) {
	requests := make(map[string]int)
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		requests[request.URL.Path]++
		switch request.URL.Path {
		case "/docs/":
			fmt.Fprint(writer, `<html><body><rapi-doc spec-url="specs/main.yaml"></rapi-doc></body></html>`)
		case "/docs/specs/main.yaml":
			fmt.Fprint(writer, `openapi: 3.0.3
info: {title: Complete, version: v1}
paths:
  /pets:
    $ref: child.yaml#/components/paths/pets
components:
  schemas:
    Root:
      type: object
      properties:
        common: {$ref: common.yaml#/components/schemas/Common}
`)
		case "/docs/specs/child.yaml":
			fmt.Fprint(writer, `components:
  paths:
    pets:
      get:
        summary: List pets
        operationId: listPets
        tags: [Pets]
        responses:
          "200":
            description: OK
            content:
              application/json:
                schema: {$ref: common.yaml#/components/schemas/Common}
  schemas:
    Pet:
      type: object
      properties:
        root: {$ref: main.yaml#/components/schemas/Root}
`)
		case "/docs/specs/common.yaml":
			fmt.Fprint(writer, `components:
  schemas:
    Common:
      type: object
      properties:
        pet: {$ref: child.yaml#/components/schemas/Pet}
`)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	root := t.TempDir()
	result, err := ImportOpenAPI(context.Background(), "Complete API", "live", server.URL+"/docs/", Options{
		LibraryRoot: root, Rebuild: func(context.Context) error { return nil }, HTTPClient: server.Client(),
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Kind != "openapi" || result.Framework != "" || result.Sources != 3 || result.Pages != 5 {
		t.Fatalf("unexpected result: %+v", result)
	}
	for _, source := range []string{"/docs/", "/docs/specs/main.yaml", "/docs/specs/child.yaml", "/docs/specs/common.yaml"} {
		if requests[source] != 1 {
			t.Errorf("requests for %s = %d, want 1", source, requests[source])
		}
	}

	operationNames, err := filepath.Glob(filepath.Join(result.Destination, "operations", "pets", "*.md"))
	if err != nil || len(operationNames) != 1 {
		t.Fatalf("operation pages: %v, %v", operationNames, err)
	}
	operation, err := os.ReadFile(operationNames[0])
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(operation), "List pets") || !strings.Contains(string(operation), "$ref: '#/components/schemas/Common'") || strings.Contains(string(operation), ".yaml#/") {
		t.Fatalf("operation was not bundled:\n%s", operation)
	}

	schemaNames, err := filepath.Glob(filepath.Join(result.Destination, "schemas", "*.md"))
	if err != nil || len(schemaNames) != 3 {
		t.Fatalf("schema pages: %v, %v", schemaNames, err)
	}
	var schemas strings.Builder
	for _, name := range schemaNames {
		raw, readErr := os.ReadFile(name)
		if readErr != nil {
			t.Fatal(readErr)
		}
		schemas.Write(raw)
	}
	if strings.Contains(schemas.String(), ".yaml#/") {
		t.Fatalf("external schema reference remained in generated pages:\n%s", schemas.String())
	}
	manifest, err := os.ReadFile(filepath.Join(result.Destination, "_index.md"))
	if err != nil || !strings.Contains(string(manifest), "sources: 3") {
		t.Fatalf("graph source count missing from manifest: %v\n%s", err, manifest)
	}
}

func TestOpenAPIExternalRefsFailClosedBeforeUnsafeRequests(t *testing.T) {
	externalRequests := 0
	external := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		externalRequests++
	}))
	defer external.Close()

	tests := []struct {
		name string
		ref  func(string) string
		want string
	}{
		{name: "origin", ref: func(string) string { return external.URL + "/child.yaml#/components/paths/pets" }, want: "changes origin"},
		{name: "scope", ref: func(string) string { return "../private.yaml#/components/paths/pets" }, want: "outside specification scope"},
		{name: "credentials", ref: func(serverURL string) string {
			return strings.Replace(serverURL, "://", "://user:secret@", 1) + "/docs/specs/child.yaml#/components/paths/pets"
		}, want: "credentials"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			unsafeRequests := 0
			var server *httptest.Server
			server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				if request.URL.Path != "/docs/specs/main.yaml" {
					unsafeRequests++
					http.Error(writer, "unsafe", http.StatusInternalServerError)
					return
				}
				fmt.Fprintf(writer, "openapi: 3.0.3\ninfo: {title: Unsafe, version: v1}\npaths:\n  /pets:\n    $ref: %q\n", test.ref(server.URL))
			}))
			defer server.Close()

			root := t.TempDir()
			rebuilds := 0
			_, err := ImportOpenAPI(context.Background(), "Unsafe "+test.name, "v1", server.URL+"/docs/specs/main.yaml", Options{
				LibraryRoot: root, HTTPClient: server.Client(), Rebuild: func(context.Context) error { rebuilds++; return nil },
			})
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("expected %q error, got %v", test.want, err)
			}
			if unsafeRequests != 0 || externalRequests != 0 || rebuilds != 0 {
				t.Fatalf("failed closed too late: local=%d external=%d rebuilds=%d", unsafeRequests, externalRequests, rebuilds)
			}
		})
	}
}

func TestOpenAPIExternalRefsRejectMissingMalformedAndOversizedSources(t *testing.T) {
	tests := []struct {
		name      string
		child     string
		pointer   string
		maxSource int64
		maxTotal  int64
		want      string
	}{
		{name: "missing pointer", child: "components: {paths: {}}\n", pointer: "/components/paths/missing", want: "does not exist"},
		{name: "malformed document", child: "components: [\n", pointer: "/components/paths/pets", want: "parse referenced OpenAPI document"},
		{name: "source limit", child: "components: {paths: {pets: {get: {responses: {\"200\": {description: ok}}}}}}\n#" + strings.Repeat("x", 600), pointer: "/components/paths/pets", maxSource: 512, maxTotal: 2_048, want: "exceeds 512 bytes"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				switch request.URL.Path {
				case "/specs/main.yaml":
					fmt.Fprintf(writer, "openapi: 3.0.3\ninfo: {title: Invalid, version: v1}\npaths:\n  /pets:\n    $ref: child.yaml#%s\n", test.pointer)
				case "/specs/child.yaml":
					fmt.Fprint(writer, test.child)
				default:
					http.NotFound(writer, request)
				}
			}))
			defer server.Close()
			options := Options{LibraryRoot: t.TempDir(), HTTPClient: server.Client(), Rebuild: func(context.Context) error { return nil }}
			if test.maxSource != 0 {
				options.MaxSourceBytes, options.MaxTotalBytes = test.maxSource, test.maxTotal
			}
			_, err := ImportOpenAPI(context.Background(), "Invalid "+test.name, "v1", server.URL+"/specs/main.yaml", options)
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("expected %q error, got %v", test.want, err)
			}
		})
	}
}

func TestOpenAPIExternalRefGraphHonorsAggregateByteLimit(t *testing.T) {
	padding := strings.Repeat("x", 300)
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/specs/main.yaml":
			fmt.Fprint(writer, "openapi: 3.0.3\ninfo: {title: Bounded, version: v1}\npaths:\n  /pets:\n    $ref: child.yaml#/components/paths/pets\n")
		case "/specs/child.yaml":
			fmt.Fprintf(writer, "components:\n  paths:\n    pets:\n      get:\n        responses:\n          \"200\": {$ref: common.yaml#/components/responses/ok}\n# %s\n", padding)
		case "/specs/common.yaml":
			fmt.Fprintf(writer, "components:\n  responses:\n    ok: {description: OK}\n# %s\n", padding)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	_, err := ImportOpenAPI(context.Background(), "Bounded graph", "v1", server.URL+"/specs/main.yaml", Options{
		LibraryRoot: t.TempDir(), HTTPClient: server.Client(), Rebuild: func(context.Context) error { return nil },
		MaxSourceBytes: 512, MaxTotalBytes: 700,
	})
	if err == nil || !strings.Contains(err.Error(), "exceeds 700 total source bytes") {
		t.Fatalf("expected aggregate byte limit error, got %v", err)
	}
}

func TestOpenAPILiteralDataRefsAreNotFetchedOrRewritten(t *testing.T) {
	tests := []struct {
		name string
		root string
		real string
	}{
		{
			name: "Swagger 2",
			root: `swagger: "2.0"
info: {title: Literal data, version: v1}
paths:
  /items:
    get:
      x-literal:
        $ref: sample.yaml
      responses:
        "200":
          description: OK
          schema: {$ref: real.yaml#/definitions/Real}
          examples:
            application/json:
              $ref: sample.yaml
`,
			real: "definitions:\n  Real: {type: object}\n",
		},
		{
			name: "OpenAPI 3.0",
			root: `openapi: 3.0.3
info: {title: Literal data, version: v1}
paths:
  /items:
    get:
      x-literal:
        $ref: sample.yaml
      responses:
        "200":
          description: OK
          content:
            application/json:
              schema: {$ref: real.yaml#/components/schemas/Real}
              example: {$ref: sample.yaml}
              examples:
                literal:
                  value: {$ref: sample.yaml}
components:
  examples:
    Payload:
      value: {$ref: sample.yaml}
`,
			real: "components:\n  schemas:\n    Real: {type: object}\n",
		},
		{
			name: "OpenAPI 3.1",
			root: `openapi: 3.1.0
info: {title: Literal data, version: v1}
paths:
  /items:
    get:
      x-literal:
        $ref: sample.yaml
      responses:
        "200":
          description: OK
          content:
            application/json:
              schema: {$ref: real.yaml#/components/schemas/Real}
              example: {$ref: sample.yaml}
              examples:
                literal:
                  value: {$ref: sample.yaml}
`,
			real: "components:\n  schemas:\n    Real: {type: object}\n",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			requests := make(map[string]int)
			server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				requests[request.URL.Path]++
				switch request.URL.Path {
				case "/specs/main.yaml":
					fmt.Fprint(writer, test.root)
				case "/specs/real.yaml":
					fmt.Fprint(writer, test.real)
				default:
					http.NotFound(writer, request)
				}
			}))
			defer server.Close()

			result, err := ImportOpenAPI(context.Background(), "Literal "+test.name, "v1", server.URL+"/specs/main.yaml", Options{
				LibraryRoot: t.TempDir(), HTTPClient: server.Client(), Rebuild: func(context.Context) error { return nil },
			})
			if err != nil {
				t.Fatal(err)
			}
			if result.Sources != 2 || requests["/specs/real.yaml"] != 1 || requests["/specs/sample.yaml"] != 0 {
				t.Fatalf("semantic requests=%v result=%+v", requests, result)
			}
			operations, _ := filepath.Glob(filepath.Join(result.Destination, "operations", "untagged", "*.md"))
			if len(operations) != 1 {
				t.Fatalf("operation pages: %v", operations)
			}
			raw, readErr := os.ReadFile(operations[0])
			if readErr != nil {
				t.Fatal(readErr)
			}
			generated := string(raw)
			if !strings.Contains(generated, "$ref: sample.yaml") || strings.Contains(generated, "real.yaml#/") || !strings.Contains(generated, "#/"+map[bool]string{true: "definitions", false: "components/schemas"}[test.name == "Swagger 2"]+"/Real") {
				t.Fatalf("literal or real references were handled incorrectly:\n%s", generated)
			}
		})
	}
}

func TestOpenAPIYAMLAliasesAreBoundedInLiteralData(t *testing.T) {
	t.Run("cycle", func(t *testing.T) {
		const specification = `openapi: 3.0.3
info: {title: Alias cycle, version: v1}
paths:
  /items:
    get:
      x-data: &cycle
        self: *cycle
      responses:
        "200":
          description: OK
          content:
            application/json:
              schema: {$ref: real.yaml#/components/schemas/Real}
`
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			if request.URL.Path == "/specs/main.yaml" {
				fmt.Fprint(writer, specification)
				return
			}
			if request.URL.Path == "/specs/real.yaml" {
				fmt.Fprint(writer, "components:\n  schemas:\n    Real: {type: object}\n")
				return
			}
			http.NotFound(writer, request)
		}))
		defer server.Close()

		root := t.TempDir()
		_, err := ImportOpenAPI(context.Background(), "Alias cycle", "v1", server.URL+"/specs/main.yaml", Options{
			LibraryRoot: root, HTTPClient: server.Client(), Rebuild: func(context.Context) error { return nil },
		})
		if err == nil || !strings.Contains(err.Error(), "cyclic YAML alias") {
			t.Fatalf("expected clean alias cycle error, got %v", err)
		}
		if _, statErr := os.Stat(filepath.Join(root, "alias-cycle", "v1")); !os.IsNotExist(statErr) {
			t.Fatalf("cyclic alias import published output: %v", statErr)
		}
	})

	t.Run("DAG node budget", func(t *testing.T) {
		document, _, err := parseOpenAPIRefDocument([]byte(`payload:
  leaf: &leaf {value: true}
  copies: [*leaf, *leaf, *leaf, *leaf]
`), "fixture.yaml")
		if err != nil {
			t.Fatal(err)
		}
		payload, _, err := openAPIJSONPointer(document, "/payload")
		if err != nil {
			t.Fatal(err)
		}
		graph := &openAPIRefGraph{bundleNodes: maxOpenAPIBundleNodes - 6}
		_, err = graph.cloneYAMLNodeBounded(payload, 0, true, make(map[*yaml.Node]bool))
		if err == nil || !strings.Contains(err.Error(), fmt.Sprintf("exceeds %d nodes", maxOpenAPIBundleNodes)) {
			t.Fatalf("alias DAG bypassed global node budget: %v", err)
		}
	})
}

func TestOpenAPI31ExternalGraphsRejectSchemaResourceSemantics(t *testing.T) {
	for _, keyword := range []string{"$id", "$anchor", "$dynamicAnchor", "$dynamicRef", "$recursiveAnchor", "$recursiveRef"} {
		t.Run(keyword, func(t *testing.T) {
			realRequests := 0
			server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				switch request.URL.Path {
				case "/specs/main.yaml":
					fmt.Fprintf(writer, `openapi: 3.1.0
info: {title: Resources, version: v1}
paths:
  /items:
    get:
      responses:
        "200":
          description: OK
          content:
            application/json:
              schema: {$ref: real.yaml#/components/schemas/Real}
components:
  schemas:
    Local:
      type: object
      %s: resource-value
`, keyword)
				case "/specs/real.yaml":
					realRequests++
					fmt.Fprint(writer, "components:\n  schemas:\n    Real: {type: object}\n")
				default:
					http.NotFound(writer, request)
				}
			}))
			defer server.Close()

			_, err := ImportOpenAPI(context.Background(), "Schema resources", "v1", server.URL+"/specs/main.yaml", Options{
				LibraryRoot: t.TempDir(), HTTPClient: server.Client(), Rebuild: func(context.Context) error { return nil },
			})
			wanted := "schema resource keyword"
			if keyword == "$dynamicRef" || keyword == "$recursiveRef" {
				wanted = "external schema resource reference"
			}
			if err == nil || !strings.Contains(err.Error(), wanted) || !strings.Contains(err.Error(), keyword) {
				t.Fatalf("expected precise %s rejection, got %v", keyword, err)
			}
			if realRequests != 0 {
				t.Fatalf("%s graph fetched with an ambiguous schema base", keyword)
			}
		})
	}

	t.Run("non-pointer fragment", func(t *testing.T) {
		requests := 0
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			requests++
			fmt.Fprint(writer, `openapi: 3.1.0
info: {title: Anchor, version: v1}
paths:
  /items:
    get:
      responses:
        "200":
          description: OK
          content:
            application/json:
              schema: {$ref: anchored.yaml#Thing}
`)
		}))
		defer server.Close()
		_, err := ImportOpenAPI(context.Background(), "Anchor", "v1", server.URL+"/specs/main.yaml", Options{
			LibraryRoot: t.TempDir(), HTTPClient: server.Client(), Rebuild: func(context.Context) error { return nil },
		})
		if err == nil || !strings.Contains(err.Error(), "non-pointer schema fragment") || !strings.Contains(err.Error(), "Thing") {
			t.Fatalf("expected non-pointer fragment rejection, got %v", err)
		}
		if requests != 1 {
			t.Fatalf("anchor target was fetched: requests=%d", requests)
		}
	})

	t.Run("single file resource remains single-file behavior", func(t *testing.T) {
		source := filepath.Join(t.TempDir(), "openapi.yaml")
		if err := os.WriteFile(source, []byte(`openapi: 3.1.0
info: {title: Local resource, version: v1}
paths:
  /items:
    get:
      responses:
        "200": {description: OK}
components:
  schemas:
    Local: {$id: urn:local, type: object}
`), 0o600); err != nil {
			t.Fatal(err)
		}
		result, err := ImportOpenAPI(context.Background(), "Local resource", "v1", source, Options{
			LibraryRoot: t.TempDir(), Rebuild: func(context.Context) error { return nil },
		})
		if err != nil || result.Sources != 1 {
			t.Fatalf("single-file resource behavior changed: %+v, %v", result, err)
		}
	})
}

func TestOpenAPI31DynamicAndRecursiveOnlyExternalRefsFailBeforeFastPath(t *testing.T) {
	for _, keyword := range []string{"$dynamicRef", "$recursiveRef"} {
		t.Run(keyword, func(t *testing.T) {
			targetRequests := 0
			var server *httptest.Server
			server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				switch request.URL.Path {
				case "/specs/main.yaml":
					fmt.Fprintf(writer, `openapi: 3.1.0
info: {title: Resource reference, version: v1}
paths:
  /items:
    get:
      responses:
        "200": {description: OK}
components:
  schemas:
    Node:
      type: object
      %s: %q
`, keyword, server.URL+"/specs/resource.yaml#Node")
				case "/specs/resource.yaml":
					targetRequests++
					fmt.Fprint(writer, "type: object\n")
				default:
					http.NotFound(writer, request)
				}
			}))
			defer server.Close()

			root := t.TempDir()
			rebuilds := 0
			_, err := ImportOpenAPI(context.Background(), "External "+keyword, "v1", server.URL+"/specs/main.yaml", Options{
				LibraryRoot: root, HTTPClient: server.Client(), Rebuild: func(context.Context) error { rebuilds++; return nil },
			})
			if err == nil || !strings.Contains(err.Error(), "external schema resource reference") || !strings.Contains(err.Error(), keyword) {
				t.Fatalf("expected preflight %s rejection, got %v", keyword, err)
			}
			if targetRequests != 0 || rebuilds != 0 {
				t.Fatalf("%s failed after side effects: target_requests=%d rebuilds=%d", keyword, targetRequests, rebuilds)
			}
			if _, statErr := os.Stat(filepath.Join(root, SafeSlug("External "+keyword), "v1")); !os.IsNotExist(statErr) {
				t.Fatalf("%s failure published output: %v", keyword, statErr)
			}
		})
	}
}

func TestOpenAPI31InternalDynamicAndRecursiveRefsRemainSingleFile(t *testing.T) {
	tests := []struct {
		name   string
		schema string
	}{
		{name: "dynamic", schema: "$dynamicAnchor: node\n      $dynamicRef: '#node'"},
		{name: "recursive", schema: "$recursiveAnchor: true\n      $recursiveRef: '#'"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			source := filepath.Join(t.TempDir(), "openapi.yaml")
			specification := `openapi: 3.1.0
info: {title: Internal resource, version: v1}
paths:
  /items:
    get:
      responses:
        "200": {description: OK}
components:
  schemas:
    Node:
      type: object
      ` + test.schema + "\n"
			if err := os.WriteFile(source, []byte(specification), 0o600); err != nil {
				t.Fatal(err)
			}
			result, err := ImportOpenAPI(context.Background(), "Internal "+test.name, "v1", source, Options{
				LibraryRoot: t.TempDir(), Rebuild: func(context.Context) error { return nil },
			})
			if err != nil || result.Sources != 1 || result.Pages != 3 {
				t.Fatalf("internal %s reference changed single-file behavior: %+v, %v", test.name, result, err)
			}
		})
	}
}

func TestOpenAPIExternalRootSchemaPreservesRecursiveReference(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/specs/main.yaml":
			fmt.Fprint(writer, `openapi: 3.0.3
info: {title: Recursive, version: v1}
paths:
  /nodes:
    get:
      responses:
        "200": {description: OK}
components:
  schemas:
    Node: {$ref: node.yaml}
`)
		case "/specs/node.yaml":
			fmt.Fprint(writer, `type: object
properties:
  value: {type: string}
  child: {$ref: "#"}
`)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	result, err := ImportOpenAPI(context.Background(), "Recursive API", "v1", server.URL+"/specs/main.yaml", Options{
		LibraryRoot: t.TempDir(), HTTPClient: server.Client(), Rebuild: func(context.Context) error { return nil },
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Sources != 2 || result.Pages != 3 {
		t.Fatalf("unexpected result: %+v", result)
	}
	schemas, err := filepath.Glob(filepath.Join(result.Destination, "schemas", "*.md"))
	if err != nil || len(schemas) != 1 {
		t.Fatalf("recursive schema pages: %v, %v", schemas, err)
	}
	raw, err := os.ReadFile(schemas[0])
	if err != nil {
		t.Fatal(err)
	}
	generated := string(raw)
	if !strings.Contains(generated, "# Node") || !strings.Contains(generated, "child:") || !strings.Contains(generated, "#/components/schemas/Node") || strings.Contains(generated, "node.yaml") {
		t.Fatalf("recursive root schema was not bundled once:\n%s", generated)
	}
}

func TestOpenAPIExternalRootSchemaReferencedByOperationPreservesRecursion(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/specs/main.yaml":
			fmt.Fprint(writer, `openapi: 3.0.3
info: {title: Recursive, version: v1}
paths:
  /nodes:
    get:
      responses:
        "200":
          description: OK
          content:
            application/json:
              schema: {$ref: node.yaml}
`)
		case "/specs/node.yaml":
			fmt.Fprint(writer, `type: object
properties:
  value: {type: string}
  child: {$ref: "#"}
`)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	root := t.TempDir()
	index := filepath.Join(t.TempDir(), "library.sqlite")
	result, err := ImportOpenAPI(context.Background(), "Recursive response API", "v1", server.URL+"/specs/main.yaml", Options{
		LibraryRoot: root, HTTPClient: server.Client(), Rebuild: func(ctx context.Context) error {
			return library.Rebuild(ctx, library.Options{UserRoot: root, IndexPath: index})
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Sources != 2 || result.Pages != 3 {
		t.Fatalf("unexpected result: %+v", result)
	}
	schemas, err := filepath.Glob(filepath.Join(result.Destination, "schemas", "*.md"))
	if err != nil || len(schemas) != 1 {
		t.Fatalf("recursive response schemas: %v, %v", schemas, err)
	}
	raw, err := os.ReadFile(schemas[0])
	if err != nil {
		t.Fatal(err)
	}
	if generated := string(raw); !strings.Contains(generated, "child:") || !strings.Contains(generated, "#/components/schemas/node") || strings.Contains(generated, "node.yaml") {
		t.Fatalf("operation schema root was not bundled once:\n%s", generated)
	}
	snapshot, err := library.Open(context.Background(), library.Options{UserRoot: root, IndexPath: index})
	if err != nil {
		t.Fatal(err)
	}
	defer snapshot.Close()
	search, err := snapshot.Search(context.Background(), library.SearchRequest{DocID: "recursive-response-api-v1", Query: "child"})
	if err != nil || search.Total == 0 {
		t.Fatalf("recursive response schema was not indexed: %+v, %v", search, err)
	}
}

func TestOpenAPI30PathItemRefMergesDistinctSiblingOperations(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/specs/main.yaml":
			fmt.Fprint(writer, `openapi: 3.0.3
info: {title: Siblings, version: v1}
paths:
  /items:
    $ref: path.yaml#/components/paths/items
    get:
      summary: Local operation
      responses:
        "200": {description: wrong}
`)
		case "/specs/path.yaml":
			fmt.Fprint(writer, `components:
  paths:
    items:
      post:
        summary: Referenced operation
        responses:
          "200": {description: OK}
`)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	root := t.TempDir()
	index := filepath.Join(t.TempDir(), "library.sqlite")
	result, err := ImportOpenAPI(context.Background(), "Sibling API", "v1", server.URL+"/specs/main.yaml", Options{
		LibraryRoot: root, HTTPClient: server.Client(), Rebuild: func(ctx context.Context) error {
			return library.Rebuild(ctx, library.Options{UserRoot: root, IndexPath: index})
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	operations, err := filepath.Glob(filepath.Join(result.Destination, "operations", "untagged", "*.md"))
	if err != nil || len(operations) != 2 || result.Pages != 3 {
		t.Fatalf("operations=%v result=%+v err=%v", operations, result, err)
	}
	var generated string
	for _, operation := range operations {
		raw, readErr := os.ReadFile(operation)
		if readErr != nil {
			t.Fatal(readErr)
		}
		generated += string(raw)
	}
	if !strings.Contains(generated, "Referenced operation") || !strings.Contains(generated, "`POST /items`") || !strings.Contains(generated, "Local operation") || !strings.Contains(generated, "`GET /items`") {
		t.Fatalf("OpenAPI 3.0 Path Item siblings were not merged:\n%s", generated)
	}
	snapshot, err := library.Open(context.Background(), library.Options{UserRoot: root, IndexPath: index})
	if err != nil {
		t.Fatal(err)
	}
	defer snapshot.Close()
	search, err := snapshot.Search(context.Background(), library.SearchRequest{DocID: "sibling-api-v1", Query: "Local operation"})
	if err != nil || search.Total == 0 {
		t.Fatalf("merged Path Item operation was not indexed: %+v, %v", search, err)
	}
}

func TestOpenAPIPathItemContextsBundleCallbacksAndWebhooks(t *testing.T) {
	const mainSpec = `openapi: 3.1.0
info: {title: Path contexts, version: v1}
paths:
  /register:
    post:
      callbacks:
        event:
          '{$request.body#/callbackUrl}':
            $ref: callback.yaml#/components/pathItems/Event
            post:
              summary: Local callback operation
              responses:
                "200": {description: OK}
            arbitrary:
              summary: Must not override the Path Item
      responses:
        "202": {description: Accepted}
webhooks:
  incoming:
    $ref: webhook.yaml#/components/pathItems/Incoming
    delete:
      summary: Local webhook operation
      responses:
        "204": {description: Deleted}
    arbitrary:
      summary: Must not override the webhook Path Item
`
	const callbackSpec = `components:
  pathItems:
    Event:
      get:
        summary: Referenced callback operation
        responses:
          "200": {description: OK}
`
	const webhookSpec = `components:
  pathItems:
    Incoming:
      put:
        summary: Referenced webhook operation
        responses:
          "200": {description: OK}
`
	requests := make(map[string]int)
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		requests[request.URL.Path]++
		switch request.URL.Path {
		case "/specs/main.yaml":
			fmt.Fprint(writer, mainSpec)
		case "/specs/callback.yaml":
			fmt.Fprint(writer, callbackSpec)
		case "/specs/webhook.yaml":
			fmt.Fprint(writer, webhookSpec)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	root := t.TempDir()
	index := filepath.Join(t.TempDir(), "library.sqlite")
	result, err := ImportOpenAPI(context.Background(), "Path contexts", "v1", server.URL+"/specs/main.yaml", Options{
		LibraryRoot: root, HTTPClient: server.Client(), Rebuild: func(ctx context.Context) error {
			return library.Rebuild(ctx, library.Options{UserRoot: root, IndexPath: index})
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Sources != 3 || result.Pages != 4 || requests["/specs/callback.yaml"] != 1 || requests["/specs/webhook.yaml"] != 1 {
		t.Fatalf("path context sources: result=%+v requests=%v", result, requests)
	}
	operations, _ := filepath.Glob(filepath.Join(result.Destination, "operations", "untagged", "*.md"))
	if len(operations) != 1 {
		t.Fatalf("operation pages: %v", operations)
	}
	operation, err := os.ReadFile(operations[0])
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(operation), "Referenced callback operation") || !strings.Contains(string(operation), "Local callback operation") || strings.Contains(string(operation), "Must not override the Path Item") {
		t.Fatalf("callback Path Item was not merged semantically:\n%s", operation)
	}
	webhookPages, err := filepath.Glob(filepath.Join(result.Destination, "webhooks", "untagged", "*.md"))
	if err != nil || len(webhookPages) != 2 {
		t.Fatalf("webhook pages: %v, %v", webhookPages, err)
	}
	var generatedWebhooks string
	for _, name := range webhookPages {
		raw, readErr := os.ReadFile(name)
		if readErr != nil {
			t.Fatal(readErr)
		}
		generatedWebhooks += string(raw)
	}
	for _, wanted := range []string{"Referenced webhook operation", "Local webhook operation", "api_endpoints:\n    - webhook:incoming", "source: " + server.URL + "/specs/main.yaml"} {
		if !strings.Contains(generatedWebhooks, wanted) {
			t.Errorf("generated webhook pages missing %q:\n%s", wanted, generatedWebhooks)
		}
	}
	if strings.Count(generatedWebhooks, "page_id: webhook-operation-") != 2 {
		t.Fatalf("webhook page IDs are not distinct and deterministic:\n%s", generatedWebhooks)
	}
	snapshot, err := library.Open(context.Background(), library.Options{UserRoot: root, IndexPath: index})
	if err != nil {
		t.Fatal(err)
	}
	defer snapshot.Close()
	search, err := snapshot.Search(context.Background(), library.SearchRequest{DocID: "path-contexts-v1", Query: "Referenced webhook operation"})
	if err != nil || search.Total == 0 {
		t.Fatalf("webhook operation was not indexed: %+v, %v", search, err)
	}

	mainNode, _, err := parseOpenAPIRefDocument([]byte(mainSpec), server.URL+"/specs/main.yaml")
	if err != nil {
		t.Fatal(err)
	}
	callbackNode, _, _ := parseOpenAPIRefDocument([]byte(callbackSpec), server.URL+"/specs/callback.yaml")
	webhookNode, _, _ := parseOpenAPIRefDocument([]byte(webhookSpec), server.URL+"/specs/webhook.yaml")
	scope, mainSource, err := newOpenAPIRefScope(server.URL + "/specs/main.yaml")
	if err != nil {
		t.Fatal(err)
	}
	callbackSource, _ := scope.canonical(server.URL + "/specs/callback.yaml")
	webhookSource, _ := scope.canonical(server.URL + "/specs/webhook.yaml")
	graph := &openAPIRefGraph{
		documents: map[string]*openAPIRefDocument{
			mainSource:     {source: mainSource, root: mainNode},
			callbackSource: {source: callbackSource, root: callbackNode},
			webhookSource:  {source: webhookSource, root: webhookNode},
		},
		aliases: map[string]string{mainSource: mainSource, callbackSource: callbackSource, webhookSource: webhookSource},
		scope:   scope, kind: "openapi", openAPI31: true,
	}
	bundled, err := graph.bundleNode(mainSource, mainNode, make(map[string]bool), openAPIDocumentContext)
	if err != nil {
		t.Fatal(err)
	}
	var structure map[string]any
	if err := bundled.Decode(&structure); err != nil {
		t.Fatal(err)
	}
	webhooks := structure["webhooks"].(map[string]any)
	incoming := webhooks["incoming"].(map[string]any)
	if incoming["put"] == nil || incoming["delete"] == nil || incoming["arbitrary"] != nil || incoming["$ref"] != nil {
		t.Fatalf("webhook Path Item structure: %#v", incoming)
	}
}

func TestOpenAPI31WebhookOnlySpecificationPublishesOperation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/openapi.yaml" {
			http.NotFound(writer, request)
			return
		}
		fmt.Fprint(writer, `openapi: 3.1.0
info: {title: Webhook only, version: v1}
webhooks:
  event:
    post:
      summary: Receive event
      operationId: receiveEvent
      responses:
        "204": {description: Accepted}
`)
	}))
	defer server.Close()

	root := t.TempDir()
	index := filepath.Join(t.TempDir(), "library.sqlite")
	result, err := ImportOpenAPI(context.Background(), "Webhook only", "v1", server.URL+"/openapi.yaml", Options{
		LibraryRoot: root, HTTPClient: server.Client(), Rebuild: func(ctx context.Context) error {
			return library.Rebuild(ctx, library.Options{UserRoot: root, IndexPath: index})
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Pages != 2 || result.Sources != 1 {
		t.Fatalf("unexpected webhook-only result: %+v", result)
	}
	pages, _ := filepath.Glob(filepath.Join(result.Destination, "webhooks", "untagged", "*.md"))
	if len(pages) != 1 {
		t.Fatalf("webhook-only pages: %v", pages)
	}
	raw, err := os.ReadFile(pages[0])
	if err != nil {
		t.Fatal(err)
	}
	for _, wanted := range []string{"Receive event", "operation_ids:\n    - receiveEvent", "api_endpoints:\n    - webhook:event", "`POST webhook:event`"} {
		if !strings.Contains(string(raw), wanted) {
			t.Errorf("webhook-only page missing %q:\n%s", wanted, raw)
		}
	}
}

func TestOpenAPI31ReferenceAndSchemaSiblingSemantics(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/specs/main.yaml":
			fmt.Fprint(writer, `openapi: 3.1.0
info: {title: Siblings, version: v1}
paths:
  /items:
    get:
      parameters:
        - $ref: common.yaml#/components/parameters/Limit
          summary: Ignored parameter summary
          description: Overridden parameter description
          required: true
          x-ignored: value
      responses:
        "200": {description: OK}
        "400":
          $ref: common.yaml#/components/responses/Error
          summary: Ignored response summary
          description: Overridden response description
components:
  schemas:
    Extended:
      $ref: common.yaml#/components/schemas/Base
      description: Extended schema
      maxLength: 12
    Constrained:
      $ref: common.yaml#/schemaBank/Text
      maxLength: 8
      unevaluatedProperties: false
      unevaluatedItems: false
      examples:
        - {$ref: literal-example.yaml}
`)
		case "/specs/common.yaml":
			fmt.Fprint(writer, `components:
  responses:
    Error:
      description: Base response description
  parameters:
    Limit:
      name: limit
      in: query
      description: Base parameter description
      required: false
      schema: {type: integer}
  schemas:
    Base:
      type: string
      minLength: 2
schemaBank:
  Text:
    type: string
    minLength: 3
`)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	result, err := ImportOpenAPI(context.Background(), "OpenAPI 31", "v1", server.URL+"/specs/main.yaml", Options{
		LibraryRoot: t.TempDir(), HTTPClient: server.Client(), Rebuild: func(context.Context) error { return nil },
	})
	if err != nil {
		t.Fatal(err)
	}
	operations, err := filepath.Glob(filepath.Join(result.Destination, "operations", "untagged", "*.md"))
	if err != nil || len(operations) != 1 {
		t.Fatalf("operation pages: %v, %v", operations, err)
	}
	operation, err := os.ReadFile(operations[0])
	if err != nil {
		t.Fatal(err)
	}
	generatedOperation := string(operation)
	if !strings.Contains(generatedOperation, "description: Overridden parameter description") || !strings.Contains(generatedOperation, "description: Overridden response description") || !strings.Contains(generatedOperation, "required: false") || strings.Contains(generatedOperation, "required: true") || strings.Contains(generatedOperation, "x-ignored") || strings.Contains(generatedOperation, "Ignored parameter summary") || strings.Contains(generatedOperation, "Ignored response summary") {
		t.Fatalf("OpenAPI 3.1 Reference Object siblings were handled incorrectly:\n%s", generatedOperation)
	}

	schemaFiles, err := filepath.Glob(filepath.Join(result.Destination, "schemas", "*.md"))
	if err != nil || len(schemaFiles) != 4 {
		t.Fatalf("schema pages: %v, %v", schemaFiles, err)
	}
	var constrained, extended string
	for _, name := range schemaFiles {
		raw, readErr := os.ReadFile(name)
		if readErr != nil {
			t.Fatal(readErr)
		}
		if strings.Contains(string(raw), "# Extended") {
			extended = string(raw)
		}
		if strings.Contains(string(raw), "# Constrained") {
			constrained = string(raw)
		}
	}
	if !strings.Contains(extended, "$ref: '#/components/schemas/Base'") || !strings.Contains(extended, "description: Extended schema") || !strings.Contains(extended, "maxLength: 12") {
		t.Fatalf("OpenAPI 3.1 Schema Object siblings were not preserved:\n%s", extended)
	}
	definition := generatedSchemaDefinition(t, constrained)
	if definition["$ref"] != "#/components/schemas/common.schemaBank.Text" || definition["maxLength"] != 8 || definition["unevaluatedProperties"] != false || definition["unevaluatedItems"] != false || definition["allOf"] != nil {
		t.Fatalf("OpenAPI 3.1 Schema siblings changed structure: %#v\n%s", definition, constrained)
	}
	examples, ok := definition["examples"].([]any)
	if !ok || len(examples) != 1 {
		t.Fatalf("schema examples changed structure: %#v", definition["examples"])
	}
	literal, ok := examples[0].(map[string]any)
	if !ok || literal["$ref"] != "literal-example.yaml" {
		t.Fatalf("literal schema example $ref changed: %#v", examples[0])
	}
}

func generatedSchemaDefinition(t *testing.T, page string) map[string]any {
	t.Helper()
	start := strings.Index(page, "```yaml\n")
	if start < 0 {
		t.Fatalf("schema page has no YAML block:\n%s", page)
	}
	definition := page[start+len("```yaml\n"):]
	end := strings.Index(definition, "```")
	if end < 0 {
		t.Fatalf("schema page has unterminated YAML block:\n%s", page)
	}
	var decoded map[string]any
	if err := yaml.Unmarshal([]byte(definition[:end]), &decoded); err != nil {
		t.Fatalf("decode generated schema: %v\n%s", err, definition[:end])
	}
	return decoded
}

func TestOpenAPI31ReferenceAnnotationsAreContextSpecific(t *testing.T) {
	graph := &openAPIRefGraph{kind: "openapi", openAPI31: true}
	tests := []struct {
		name    string
		context openAPIBundleContext
		summary bool
		desc    bool
	}{
		{name: "Path Item", context: openAPIPathItemContext, summary: true, desc: true},
		{name: "Parameter", context: openAPIParameterContext, desc: true},
		{name: "Header", context: openAPIHeaderContext, desc: true},
		{name: "Request Body", context: openAPIRequestBodyContext, desc: true},
		{name: "Response", context: openAPIResponseContext, desc: true},
		{name: "Example", context: openAPIExampleContext, summary: true, desc: true},
		{name: "Link", context: openAPILinkContext, desc: true},
		{name: "Callback", context: openAPICallbackContext},
		{name: "Security Scheme", context: openAPISecuritySchemeContext, desc: true},
		{name: "Unknown Reference", context: openAPIReferenceContext},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := graph.referenceSiblingAllowed(test.context, "summary"); got != test.summary {
				t.Errorf("summary allowed = %t, want %t", got, test.summary)
			}
			if got := graph.referenceSiblingAllowed(test.context, "description"); got != test.desc {
				t.Errorf("description allowed = %t, want %t", got, test.desc)
			}
			if graph.referenceSiblingAllowed(test.context, "required") {
				t.Error("unsupported Reference Object field was allowed")
			}
		})
	}
}

func TestOpenAPISchemaInventoryNamespacesReachabilityAndAliases(t *testing.T) {
	const mainSpec = `openapi: 3.0.3
info: {title: Inventory, version: v1}
paths:
  /inventory:
    get:
      responses:
        "200":
          description: OK
          content:
            application/json:
              schema:
                allOf:
                  - {$ref: common.yaml#/components/schemas/clients/get}
                  - {$ref: common.yaml#/components/schemas/twins/one}
                  - {$ref: common.yaml#/components/schemas/twins/two}
                  - {$ref: common.yaml#/components/errors/bad_request}
                  - {$ref: common.yaml#/components/schemas/Alias}
`
	const commonSpec = `components:
  schemas:
    Empty: {}
    Real: {type: object}
    Alias: {$ref: '#/components/schemas/Real'}
    clients:
      get: {type: object, description: Reachable client schema}
      dead: {type: object, description: Dead client schema}
    twins:
      one: {type: object, description: Identical content}
      two: {type: object, description: Identical content}
    dead_namespace:
      ghost: {type: object, description: Dead namespace child}
  errors:
    bad_request: {type: object, description: Reachable error schema}
`
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/specs/main.yaml":
			fmt.Fprint(writer, mainSpec)
		case "/specs/common.yaml":
			fmt.Fprint(writer, commonSpec)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	result, err := ImportOpenAPI(context.Background(), "Inventory", "v1", server.URL+"/specs/main.yaml", Options{
		LibraryRoot: t.TempDir(), HTTPClient: server.Client(), Rebuild: func(context.Context) error { return nil },
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Sources != 2 || result.Pages != 8 {
		t.Fatalf("inventory result: %+v", result)
	}
	schemaFiles, err := filepath.Glob(filepath.Join(result.Destination, "schemas", "*.md"))
	if err != nil || len(schemaFiles) != 6 {
		t.Fatalf("schema inventory: %v, %v", schemaFiles, err)
	}
	var titles []string
	var generated string
	for _, name := range schemaFiles {
		raw, readErr := os.ReadFile(name)
		if readErr != nil {
			t.Fatal(readErr)
		}
		page := string(raw)
		generated += page
		for _, line := range strings.Split(page, "\n") {
			if strings.HasPrefix(line, "# ") {
				titles = append(titles, strings.TrimPrefix(line, "# "))
				break
			}
		}
	}
	sort.Strings(titles)
	want := []string{"Empty", "Real", "clients.get", "common.errors.bad_request", "twins.one", "twins.two"}
	if strings.Join(titles, "\n") != strings.Join(want, "\n") {
		t.Fatalf("schema titles = %v, want %v", titles, want)
	}
	for _, omitted := range []string{"# Alias\n", "# clients\n", "# twins\n", "# dead_namespace\n", "# clients.dead\n", "# dead_namespace.ghost\n"} {
		if strings.Contains(generated, omitted) {
			t.Errorf("inventory promoted or retained %q:\n%s", omitted, generated)
		}
	}
}
