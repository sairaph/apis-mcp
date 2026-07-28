package importer

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestStaticScalarConfigurationsResolveOnlyStableCredentialFreeURLs(t *testing.T) {
	base, _ := url.Parse("https://docs.example.test/reference/")
	for _, test := range []struct {
		name   string
		html   string
		proven bool
		want   string
	}{
		{name: "custom element JSON", html: `<scalar-api-reference configuration='{ "url": "../openapi.json" }'></scalar-api-reference>`, proven: true, want: "https://docs.example.test/openapi.json"},
		{name: "configuration script data URL", html: `<script id="api-reference" data-url="./openapi.yaml"></script>`, proven: true, want: "https://docs.example.test/reference/openapi.yaml"},
		{name: "inline create reference", html: `<script>Scalar.createApiReference('#app', { url: 'https://api.example.test/schema' })</script>`, proven: true, want: "https://api.example.test/schema"},
		{name: "static sources", html: `<scalar-api-reference configuration='{ "sources": [{"url":"one.json"}] }'></scalar-api-reference>`, proven: true, want: "https://docs.example.test/reference/one.json"},
		{name: "ambiguous sources", html: `<scalar-api-reference configuration='{ "sources": [{"url":"one.json"},{"url":"two.json"}] }'></scalar-api-reference>`, proven: true},
		{name: "dynamic value", html: `<script>Scalar.createApiReference('#app', { url: window.schemaURL })</script>`, proven: true},
		{name: "arbitrary concatenation", html: `<script>Scalar.createApiReference('#app', { url: "https://api.example.test/" + version })</script>`, proven: true},
		{name: "credentials", html: `<scalar-api-reference data-url="https://user:secret@api.example.test/openapi.json"></scalar-api-reference>`, proven: true},
		{name: "generic api-reference ID", html: `<script id="api-reference">{"url":"https://api.example.test/openapi.json"}</script>`},
		{name: "unrelated JSON", html: `<script type="application/json">{"url":"https://api.example.test/openapi.json"}</script>`},
	} {
		t.Run(test.name, func(t *testing.T) {
			document, err := parseHTML([]byte(`<!doctype html><html><body>` + test.html + `</body></html>`))
			if err != nil {
				t.Fatal(err)
			}
			candidates, proven := scalarHTMLSchemaCandidates(document, base)
			resolution := stableScalarSchema(candidates, proven)
			if resolution.Proven != test.proven || resolution.URL != test.want {
				t.Fatalf("resolution = %+v, candidates=%v; want proven=%t URL=%q", resolution, candidates, test.proven, test.want)
			}
		})
	}
}

func TestDetectAndImportScalarNextRouteShell(t *testing.T) {
	specRequests := 0
	specServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		specRequests++
		if request.URL.Path != "/api/v2" || request.URL.Query().Get("outputOpenapiSchema") != "true" {
			http.NotFound(writer, request)
			return
		}
		fmt.Fprint(writer, `{"openapi":"3.0.3","info":{"title":"Scalar","version":"v1"},"paths":{"/devices":{"get":{"operationId":"listDevices","responses":{"200":{"description":"ok"}}}}}}`)
	}))
	defer specServer.Close()

	landingRequests, routeRequests, globalRequests := 0, 0, 0
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/api-docs":
			landingRequests++
			fmt.Fprintf(writer, `<!doctype html><html><body><div id="__next"></div><script id="__NEXT_DATA__" type="application/json">{"props":{"pageProps":{"baseURL":%q}},"page":"/api-docs"}</script><script src="/_next/static/chunks/main-global123.js"></script><script src="/_next/static/chunks/pages/api-docs-route123.js"></script></body></html>`, specServer.URL)
		case "/_next/static/chunks/pages/api-docs-route123.js":
			routeRequests++
			fmt.Fprint(writer, `(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[1],{2:function(e,t,n){function page(e){let{baseURL:t,page:n}=e;return render({configuration:{url:`+"`"+`${t}/api/v2?outputOpenapiSchema=true`+"`"+`,darkMode:!1,forceDarkModeState:"light",hideTestRequestButton:!0}})}}}]);`)
		case "/_next/static/chunks/main-global123.js":
			globalRequests++
			fmt.Fprint(writer, `throw new Error("must not inspect global chunks")`)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	root := t.TempDir()
	options := Options{LibraryRoot: root, Rebuild: func(context.Context) error { return nil }, HTTPClient: server.Client()}
	detection, err := DetectURL(context.Background(), server.URL+"/api-docs", options)
	wantSpec := specServer.URL + "/api/v2?outputOpenapiSchema=true"
	if err != nil || detection.Engine != "openapi" || detection.Framework != "scalar" || detection.Format != "html" || detection.Source != wantSpec || detection.DownloadedBytes == 0 {
		t.Fatalf("detection = %+v, %v", detection, err)
	}
	result, err := ImportOpenAPI(context.Background(), "Scalar API", "live", server.URL+"/api-docs", options)
	if err != nil {
		t.Fatal(err)
	}
	if result.Kind != "openapi" || result.Framework != "scalar" || result.Source != wantSpec || result.Pages != 2 || result.Sources != 1 {
		t.Fatalf("result = %+v", result)
	}
	if landingRequests != 2 || routeRequests != 2 || globalRequests != 0 || specRequests != 1 {
		t.Fatalf("requests: landing=%d route=%d global=%d spec=%d", landingRequests, routeRequests, globalRequests, specRequests)
	}
	operationFiles, _ := filepath.Glob(filepath.Join(result.Destination, "operations", "untagged", "*.md"))
	if len(operationFiles) != 1 {
		t.Fatalf("operation files = %v", operationFiles)
	}
}

func TestGenericAPIReferenceScriptIDIsNotScalar(t *testing.T) {
	specRequests := 0
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path == "/openapi.json" {
			specRequests++
			fmt.Fprint(writer, `{"openapi":"3.0.3","paths":{"/wrong":{"get":{"responses":{"200":{"description":"ok"}}}}}}`)
			return
		}
		fmt.Fprint(writer, `<!doctype html><html><body><script id="api-reference">{"url":"/openapi.json"}</script></body></html>`)
	}))
	defer server.Close()
	root := t.TempDir()
	options := Options{LibraryRoot: root, Rebuild: func(context.Context) error { return nil }, HTTPClient: server.Client()}
	detection, err := DetectURL(context.Background(), server.URL+"/docs", options)
	if err != nil || detection.Engine != "html" || detection.Framework != "unknown" {
		t.Fatalf("detection = %+v, %v", detection, err)
	}
	_, err = ImportOpenAPI(context.Background(), "Generic Script", "live", server.URL+"/docs", options)
	if err == nil || !strings.Contains(err.Error(), "no discoverable specification URL") || specRequests != 0 {
		t.Fatalf("import error=%v spec requests=%d", err, specRequests)
	}
	if _, statErr := os.Stat(filepath.Join(root, "generic-script", "live")); !os.IsNotExist(statErr) {
		t.Fatalf("generic script destination exists: %v", statErr)
	}
}

func TestExplicitDirectOpenAPIPreservesCrossOriginRedirectBehavior(t *testing.T) {
	targetRequests := 0
	target := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		targetRequests++
		fmt.Fprint(writer, `{"openapi":"3.0.3","info":{"title":"Direct","version":"v1"},"paths":{"/direct":{"get":{"responses":{"200":{"description":"ok"}}}}}}`)
	}))
	defer target.Close()
	sourceRequests := 0
	source := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		sourceRequests++
		http.Redirect(writer, request, target.URL+"/openapi.json", http.StatusFound)
	}))
	defer source.Close()
	root := t.TempDir()
	result, err := ImportOpenAPI(context.Background(), "Direct API", "v1", source.URL+"/schema", Options{
		LibraryRoot: root, Rebuild: func(context.Context) error { return nil }, HTTPClient: source.Client(),
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Source != target.URL+"/openapi.json" || result.Pages != 2 || sourceRequests != 1 || targetRequests != 1 {
		t.Fatalf("direct redirect result=%+v source requests=%d target requests=%d", result, sourceRequests, targetRequests)
	}
}

func TestScalarAmbiguousAndDynamicShellsRemainUnsupportedHTML(t *testing.T) {
	for _, test := range []struct {
		name  string
		html  string
		route string
	}{
		{name: "ambiguous static", html: `<scalar-api-reference configuration='{"sources":[{"url":"one.json"},{"url":"two.json"}]}'></scalar-api-reference>`},
		{name: "dynamic inline", html: `<script>Scalar.createApiReference('#app',{url:window.schema})</script>`},
		{name: "dynamic route concatenation", html: `<script id="__NEXT_DATA__" type="application/json">{"props":{"pageProps":{"baseURL":"https://api.example.test"}},"page":"/api-docs"}</script><script src="/_next/static/chunks/pages/api-docs-dynamic1.js"></script>`, route: `function page(e){let{baseURL:t}=e;return x({configuration:{url:t+"/openapi.json",forceDarkModeState:"light",hideTestRequestButton:!0}})}`},
	} {
		t.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				if strings.Contains(request.URL.Path, "dynamic1.js") {
					fmt.Fprint(writer, test.route)
					return
				}
				fmt.Fprint(writer, `<!doctype html><html><body>`+test.html+`</body></html>`)
			}))
			defer server.Close()
			options := Options{LibraryRoot: t.TempDir(), Rebuild: func(context.Context) error { return nil }, HTTPClient: server.Client()}
			detection, err := DetectURL(context.Background(), server.URL+"/api-docs", options)
			if err != nil || detection.Engine != "html" || detection.Framework != "scalar" {
				t.Fatalf("detection = %+v, %v", detection, err)
			}
			_, err = ImportOpenAPI(context.Background(), "Unsupported Scalar", "live", server.URL+"/api-docs", options)
			if err == nil || !strings.Contains(err.Error(), "no single statically resolvable") {
				t.Fatalf("import error = %v", err)
			}
			if _, statErr := os.Stat(filepath.Join(options.LibraryRoot, "unsupported-scalar", "live")); !os.IsNotExist(statErr) {
				t.Fatalf("unsupported destination exists: %v", statErr)
			}
		})
	}
}

func TestScalarNextRouteScriptSafetyAndFalsePositives(t *testing.T) {
	t.Run("cross-origin script is never fetched", func(t *testing.T) {
		externalRequests := 0
		external := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) { externalRequests++ }))
		defer external.Close()
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
			fmt.Fprintf(writer, `<!doctype html><scalar-api-reference></scalar-api-reference><script id="__NEXT_DATA__" type="application/json">{"props":{"pageProps":{}},"page":"/api-docs"}</script><script src="%s/_next/static/chunks/pages/api-docs-unsafe1.js"></script>`, external.URL)
		}))
		defer server.Close()
		detection, err := DetectURL(context.Background(), server.URL+"/api-docs", Options{LibraryRoot: t.TempDir(), Rebuild: func(context.Context) error { return nil }, HTTPClient: server.Client()})
		if err != nil || detection.Engine != "html" || detection.Framework != "scalar" || externalRequests != 0 {
			t.Fatalf("detection=%+v err=%v external requests=%d", detection, err, externalRequests)
		}
	})

	t.Run("same-origin redirect cannot reach cross-origin target", func(t *testing.T) {
		externalRequests := 0
		external := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) { externalRequests++ }))
		defer external.Close()
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			if strings.HasSuffix(request.URL.Path, ".js") {
				http.Redirect(writer, request, external.URL+"/malicious.js", http.StatusFound)
				return
			}
			fmt.Fprint(writer, `<!doctype html><scalar-api-reference></scalar-api-reference><script id="__NEXT_DATA__" type="application/json">{"props":{"pageProps":{}},"page":"/api-docs"}</script><script src="/_next/static/chunks/pages/api-docs-redirect1.js"></script>`)
		}))
		defer server.Close()
		_, err := DetectURL(context.Background(), server.URL+"/api-docs", Options{LibraryRoot: t.TempDir(), Rebuild: func(context.Context) error { return nil }, HTTPClient: server.Client()})
		if err == nil || !strings.Contains(err.Error(), "redirect changes origin") || externalRequests != 0 {
			t.Fatalf("error=%v external requests=%d", err, externalRequests)
		}
	})

	t.Run("generic Next route is not Scalar and inspection is capped", func(t *testing.T) {
		routeRequests, globalRequests := 0, 0
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			if strings.Contains(request.URL.Path, "/pages/api-docs-") {
				routeRequests++
				fmt.Fprint(writer, `function page(){return render({configuration:{url:"/preferences"}})}`)
				return
			}
			if strings.HasSuffix(request.URL.Path, ".js") {
				globalRequests++
				return
			}
			fmt.Fprint(writer, `<!doctype html><html><body><script id="__NEXT_DATA__" type="application/json">{"props":{"pageProps":{}},"page":"/api-docs"}</script><script src="/_next/static/chunks/main-global1.js"></script><script src="/_next/static/chunks/pages/api-docs-route001.js"></script><script src="/_next/static/chunks/pages/api-docs-route002.js"></script><script src="/_next/static/chunks/pages/api-docs-route003.js"></script><script src="/_next/static/chunks/pages/api-docs-route004.js"></script></body></html>`)
		}))
		defer server.Close()
		detection, err := DetectURL(context.Background(), server.URL+"/api-docs", Options{LibraryRoot: t.TempDir(), Rebuild: func(context.Context) error { return nil }, HTTPClient: server.Client()})
		if err != nil || detection.Engine != "html" || detection.Framework != "unknown" || routeRequests != maxScalarNextRouteScripts || globalRequests != 0 {
			t.Fatalf("detection=%+v err=%v route=%d global=%d", detection, err, routeRequests, globalRequests)
		}
	})
}

func TestScalarRouteInspectionHonorsSharedByteBudgetAndCancellation(t *testing.T) {
	landing := `<!doctype html><script id="__NEXT_DATA__" type="application/json">{"props":{"pageProps":{}},"page":"/api-docs"}</script><script src="/_next/static/chunks/pages/api-docs-budget01.js"></script>`
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if strings.HasSuffix(request.URL.Path, ".js") {
			fmt.Fprint(writer, strings.Repeat("x", 128))
			return
		}
		fmt.Fprint(writer, landing)
	}))
	defer server.Close()
	limit := int64(len(landing) + 64)
	_, err := DetectURL(context.Background(), server.URL+"/api-docs", Options{LibraryRoot: t.TempDir(), Rebuild: func(context.Context) error { return nil }, HTTPClient: server.Client(), MaxSourceBytes: int64(len(landing) + 1), MaxTotalBytes: limit})
	if err == nil || !strings.Contains(err.Error(), "total source bytes") {
		t.Fatalf("shared total byte error = %v", err)
	}

	started := make(chan struct{})
	cancelServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if strings.HasSuffix(request.URL.Path, ".js") {
			close(started)
			<-request.Context().Done()
			return
		}
		fmt.Fprint(writer, landing)
	}))
	defer cancelServer.Close()
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error, 1)
	go func() {
		_, detectErr := DetectURL(ctx, cancelServer.URL+"/api-docs", Options{LibraryRoot: t.TempDir(), Rebuild: func(context.Context) error { return nil }, HTTPClient: cancelServer.Client()})
		done <- detectErr
	}()
	<-started
	cancel()
	err = <-done
	if err == nil || !strings.Contains(err.Error(), context.Canceled.Error()) {
		t.Fatalf("route cancellation error = %v", err)
	}
}
