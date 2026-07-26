package importer

import (
	"net/url"
	"reflect"
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
