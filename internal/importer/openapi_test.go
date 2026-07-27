package importer

import (
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestOpenAPIConfigCandidatesResolveSelectedVariable(t *testing.T) {
	base, err := url.Parse("https://docs.example.test/ui/")
	if err != nil {
		t.Fatal(err)
	}
	script := `const defaultDefinitionUrl = "./v1.json";
const selectedDefinition = "./v2.json";
const definitionURL = selectedDefinition;
SwaggerUIBundle({url: definitionURL});`
	want := []string{"https://docs.example.test/ui/v2.json"}
	if got := openAPIConfigCandidates(script, base); !reflect.DeepEqual(got, want) {
		t.Fatalf("candidates = %v, want %v", got, want)
	}
}

func TestOpenAPIConfigCandidatesResolveCurrentHostMap(t *testing.T) {
	base, err := url.Parse("https://petstore3.swagger.io/")
	if err != nil {
		t.Fatal(err)
	}
	script := "const services = `petstore.swagger.io=https://petstore.swagger.io/v2/swagger.json," +
		"petstore3.swagger.io=https://petstore3.swagger.io/api/v3/openapi.json`;\n" +
		"const selected = services.split(',').find(([host]) => window.location.host.includes(host));\n" +
		"SwaggerUIBundle({url: selected});"
	want := []string{"https://petstore3.swagger.io/api/v3/openapi.json"}
	if got := openAPIConfigCandidates(script, base); !reflect.DeepEqual(got, want) {
		t.Fatalf("candidates = %v, want %v", got, want)
	}
}

func TestCatalogOpenAPIRequiresPathsMemberButAllowsEmptyPaths(t *testing.T) {
	if _, _, err := parseCatalogAPIDescription([]byte(`{"openapi":"3.0.3","info":{"title":"Missing"}}`)); err == nil || !strings.Contains(err.Error(), "paths member") {
		t.Fatalf("expected missing paths error, got %v", err)
	}
	if _, kind, err := parseCatalogAPIDescription([]byte(`{"openapi":"3.0.3","info":{"title":"Empty"},"paths":{}}`)); err != nil || kind != "openapi" {
		t.Fatalf("empty paths catalog document: kind=%q err=%v", kind, err)
	}
}

func TestStaticRapiDocNamespacesDisambiguateOriginsAndQueries(t *testing.T) {
	sources := staticAPISpecSources([]string{
		"https://one.example/openapi.json",
		"https://two.example/openapi.json",
		"https://one.example/openapi.json?group=v2",
	})
	if len(sources) != 3 {
		t.Fatalf("sources = %+v", sources)
	}
	seen := make(map[string]bool)
	for _, source := range sources {
		if seen[source.Namespace] {
			t.Fatalf("duplicate namespace in %+v", sources)
		}
		seen[source.Namespace] = true
	}
}

func TestHTTPURLCanonicalizationNormalizesOriginsAndAliases(t *testing.T) {
	left, _ := url.Parse("https://EXAMPLE.test/openapi.json")
	right, _ := url.Parse("https://example.test:443/openapi.json")
	if !sameHTTPOrigin(left, right) {
		t.Fatalf("equivalent origins differ: %s and %s", httpOrigin(left), httpOrigin(right))
	}
	canonical, err := canonicalHTTPURL("https://EXAMPLE.test:443/openapi.json?z=2&a=1#section")
	if err != nil || canonical != "https://example.test/openapi.json?z=2&a=1" {
		t.Fatalf("canonical URL = %q, %v", canonical, err)
	}
}

func TestCatalogPathsRejectDotSegments(t *testing.T) {
	for _, value := range []string{"/../admin.json", "/v1/./admin.json"} {
		if !hasDotPathSegment(value) {
			t.Fatalf("dot path accepted: %q", value)
		}
	}
	if hasDotPathSegment("/v1/products.json") {
		t.Fatal("ordinary path rejected")
	}
}
