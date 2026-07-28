//go:build dev

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/sairaph/apis-mcp/internal/importer"
	"github.com/sairaph/apis-mcp/library"
)

type workerRoundTripFunc func(*http.Request) (*http.Response, error)

func (function workerRoundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return function(request)
}

func workerResponse(request *http.Request, status int, body string) *http.Response {
	return &http.Response{StatusCode: status, Status: fmt.Sprintf("%d %s", status, http.StatusText(status)), Header: make(http.Header), Body: io.NopCloser(strings.NewReader(body)), Request: request}
}

func TestWorkerIngestsAndIndexesOpenAPI(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, `{"openapi":"3.0.3","info":{"title":"Pets","version":"v1"},"paths":{"/pets":{"get":{"responses":{"200":{"description":"ok"}}}}}}`)
	}))
	defer server.Close()

	store, err := openJobStore(t.TempDir(), true)
	if err != nil {
		t.Fatal(err)
	}
	job, err := store.create(ingestRequest{
		Output: store.output, Source: server.URL + "/openapi.json", Name: "Pet API", Version: "v1",
		Collections: []string{"payments", "developer_tools"}, Scope: "path", MaxPages: -1, MaxDepth: -1,
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := runWorker(context.Background(), store, job.ID, server.Client()); err != nil {
		t.Fatal(err)
	}
	job, err = store.get(job.ID)
	if err != nil {
		t.Fatal(err)
	}
	if job.State != jobSucceeded || job.Result == nil || job.Result.Pages != 2 {
		t.Fatalf("unexpected job: %+v", job)
	}
	manifest, err := os.ReadFile(filepath.Join(store.output, "pet-api", "v1", "_index.md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(manifest), "    - payments") || !strings.Contains(string(manifest), "    - developer_tools") {
		t.Fatalf("manifest collections missing:\n%s", manifest)
	}
	generations, _ := filepath.Glob(filepath.Join(store.output, "library-*.sqlite"))
	if len(generations) != 1 {
		t.Fatalf("SQLite generations: %v", generations)
	}
	snapshot, err := library.Open(context.Background(), library.Options{UserRoot: store.output, IndexPath: filepath.Join(store.output, "library.sqlite")})
	if err != nil {
		t.Fatal(err)
	}
	defer snapshot.Close()
	search, err := snapshot.Search(context.Background(), library.SearchRequest{DocID: "pet-api-v1", Query: "pets"})
	if err != nil || search.Total == 0 {
		t.Fatalf("indexed search: %+v, %v", search, err)
	}
	events, err := store.events(job.ID, 0)
	if err != nil || len(events) < 4 {
		t.Fatalf("events: %+v, %v", events, err)
	}
	job.State = jobRunning
	job.Result = nil
	job.FinishedAt = nil
	job.CurrentStage = "publishing"
	if err := os.Remove(filepath.Join(store.root, job.ID+".published.json")); err != nil {
		t.Fatal(err)
	}
	if err := writeJSONAtomic(filepath.Join(store.root, job.ID+".json"), job); err != nil {
		t.Fatal(err)
	}
	recovered, err := store.get(job.ID)
	if err != nil || recovered.State != jobSucceeded || recovered.Result == nil || recovered.Result.Kind != "openapi" || recovered.Result.Sources != 1 {
		t.Fatalf("published job recovery: %+v, %v", recovered, err)
	}
}

func TestWorkerDiscoversScalarShellAndPublishes129IndexedPages(t *testing.T) {
	landingRequests, routeRequests, schemaRequests, globalRequests := 0, 0, 0, 0
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/api-docs":
			landingRequests++
			fmt.Fprintf(writer, `<!doctype html><html><body><div id="__next"></div><script id="__NEXT_DATA__" type="application/json">{"props":{"pageProps":{"baseURL":%q}},"page":"/api-docs"}</script><script src="/_next/static/chunks/framework-global1.js"></script><script src="/_next/static/chunks/pages/api-docs-worker01.js"></script></body></html>`, server.URL)
		case "/_next/static/chunks/pages/api-docs-worker01.js":
			routeRequests++
			fmt.Fprint(writer, `function page(e){let{baseURL:t,page:n}=e;return render({configuration:{url:`+"`"+`${t}/api/v2?outputOpenapiSchema=true`+"`"+`,darkMode:!1,forceDarkModeState:"light",hideTestRequestButton:!0}})}`)
		case "/api/v2":
			if request.URL.Query().Get("outputOpenapiSchema") != "true" {
				http.Error(writer, "missing schema query", http.StatusBadRequest)
				return
			}
			schemaRequests++
			fmt.Fprint(writer, scalarWorkerSchema(128))
		case "/_next/static/chunks/framework-global1.js":
			globalRequests++
			http.Error(writer, "must not inspect global chunks", http.StatusInternalServerError)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	store, err := openJobStore(t.TempDir(), true)
	if err != nil {
		t.Fatal(err)
	}
	job, err := store.create(ingestRequest{
		Output: store.output, Source: server.URL + "/api-docs", Name: "Scalar Worker API", Version: "live",
		Scope: "path", MaxPages: -1, MaxDepth: -1,
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := runWorker(context.Background(), store, job.ID, server.Client()); err != nil {
		t.Fatal(err)
	}
	job, err = store.get(job.ID)
	if err != nil {
		t.Fatal(err)
	}
	wantSchema := server.URL + "/api/v2?outputOpenapiSchema=true"
	if job.State != jobSucceeded || job.Detection == nil || job.Detection.Engine != "openapi" || job.Detection.Framework != "scalar" || job.Result == nil || job.Result.Kind != "openapi" || job.Result.Framework != "scalar" || job.Result.Source != wantSchema || job.Result.Pages != 129 || job.Result.Sources != 1 || job.Pages != 129 || job.Truncated {
		t.Fatalf("unexpected Scalar worker job: %+v", job)
	}
	if landingRequests != 1 || routeRequests != 1 || schemaRequests != 1 || globalRequests != 0 {
		t.Fatalf("worker requests: landing=%d route=%d schema=%d global=%d", landingRequests, routeRequests, schemaRequests, globalRequests)
	}
	pages, err := filepath.Glob(filepath.Join(job.Result.Destination, "operations", "worker", "*.md"))
	if err != nil || len(pages) != 128 {
		t.Fatalf("published operation pages = %d, %v", len(pages), err)
	}
	if _, err := os.Stat(filepath.Join(store.root, job.ID+".published.json")); err != nil {
		t.Fatalf("publication receipt: %v", err)
	}
	generations, err := filepath.Glob(filepath.Join(store.output, "library-*.sqlite"))
	if err != nil || len(generations) != 1 {
		t.Fatalf("SQLite generations: %v, %v", generations, err)
	}
	snapshot, err := library.Open(context.Background(), library.Options{UserRoot: store.output, IndexPath: filepath.Join(store.output, "library.sqlite")})
	if err != nil {
		t.Fatal(err)
	}
	defer snapshot.Close()
	search, err := snapshot.Search(context.Background(), library.SearchRequest{DocID: "scalar-worker-api-live", Query: "scalarNeedle127"})
	if err != nil || search.Total == 0 || len(search.Hits) == 0 || search.Hits[0].PageID == "" {
		t.Fatalf("Scalar FTS result: %+v, %v", search, err)
	}
}

func TestWorkerEnforcesAggregateBytesAcrossScalarDetectionAndImport(t *testing.T) {
	landingRequests, routeRequests, schemaRequests := 0, 0, 0
	route := `function page(e){let{baseURL:t}=e;return render({configuration:{url:` + "`" + `${t}/openapi.json` + "`" + `,forceDarkModeState:"light",hideTestRequestButton:!0}})}`
	schema := `{"openapi":"3.0.3","info":{"title":"Budget","version":"v1"},"paths":{"/budget":{"get":{"responses":{"200":{"description":"ok"}}}}}}`
	var landing string
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/api-docs":
			landingRequests++
			fmt.Fprint(writer, landing)
		case "/_next/static/chunks/pages/api-docs-budget01.js":
			routeRequests++
			fmt.Fprint(writer, route)
		case "/openapi.json":
			schemaRequests++
			fmt.Fprint(writer, schema)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()
	landing = fmt.Sprintf(`<!doctype html><html><body><script id="__NEXT_DATA__" type="application/json">{"props":{"pageProps":{"baseURL":%q}},"page":"/api-docs"}</script><script src="/_next/static/chunks/pages/api-docs-budget01.js"></script></body></html>`, server.URL)
	detectionBytes := int64(len(landing) + len(route))
	totalBytes := detectionBytes + int64(len(schema)) - 1
	maxSourceBytes := int64(max(len(landing), len(route), len(schema)))

	store, err := openJobStore(t.TempDir(), true)
	if err != nil {
		t.Fatal(err)
	}
	job, err := store.create(ingestRequest{
		Output: store.output, Source: server.URL + "/api-docs", Name: "Scalar Budget", Version: "live",
		Scope: "path", MaxPages: -1, MaxDepth: -1, MaxSourceBytes: maxSourceBytes, MaxTotalBytes: totalBytes,
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := runWorker(context.Background(), store, job.ID, server.Client()); err != nil {
		t.Fatal(err)
	}
	job, err = store.get(job.ID)
	if err != nil {
		t.Fatal(err)
	}
	wantBudgetError := fmt.Sprintf("source exceeds %d bytes", len(schema)-1)
	if job.State != jobFailed || job.Detection == nil || job.Detection.Engine != "openapi" || job.Detection.Framework != "scalar" || job.Detection.DownloadedBytes != detectionBytes || job.Result != nil || !strings.Contains(job.Error, wantBudgetError) {
		t.Fatalf("aggregate-budget job: %+v", job)
	}
	if landingRequests != 1 || routeRequests != 1 || schemaRequests != 1 {
		t.Fatalf("aggregate-budget requests: landing=%d route=%d schema=%d", landingRequests, routeRequests, schemaRequests)
	}
	if _, err := os.Stat(filepath.Join(store.output, "scalar-budget", "live")); !os.IsNotExist(err) {
		t.Fatalf("aggregate-budget destination exists: %v", err)
	}
	indexes, err := filepath.Glob(filepath.Join(store.output, "library*.sqlite"))
	if err != nil || len(indexes) != 0 {
		t.Fatalf("aggregate-budget indexes = %v, %v", indexes, err)
	}
}

func TestWorkerRestrictsDiscoveredScalarSchemaRedirects(t *testing.T) {
	t.Run("cross origin", func(t *testing.T) {
		targetRequests := 0
		target := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
			targetRequests++
			fmt.Fprint(writer, scalarWorkerSchema(1))
		}))
		defer target.Close()
		schemaRequests := 0
		schema := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			schemaRequests++
			http.Redirect(writer, request, target.URL+"/openapi.json", http.StatusFound)
		}))
		defer schema.Close()

		store, job, docsRequests := runDiscoveredScalarWorker(t, schema.Client(), schema.URL+"/openapi.json")
		if job.State != jobFailed || job.Result != nil || !strings.Contains(job.Error, "redirect changes origin") {
			t.Fatalf("cross-origin redirect job: %+v", job)
		}
		if docsRequests != 1 || schemaRequests != 1 || targetRequests != 0 {
			t.Fatalf("cross-origin requests: docs=%d schema=%d target=%d", docsRequests, schemaRequests, targetRequests)
		}
		assertWorkerPublishedNothing(t, store, job)
	})

	t.Run("same host userinfo", func(t *testing.T) {
		schemaRequests, targetRequests := 0, 0
		var schema *httptest.Server
		schema = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			if request.URL.Path == "/final.json" {
				targetRequests++
				fmt.Fprint(writer, scalarWorkerSchema(1))
				return
			}
			schemaRequests++
			location := strings.Replace(schema.URL, "://", "://user:secret@", 1) + "/final.json"
			http.Redirect(writer, request, location, http.StatusFound)
		}))
		defer schema.Close()

		store, job, docsRequests := runDiscoveredScalarWorker(t, schema.Client(), schema.URL+"/openapi.json")
		if job.State != jobFailed || job.Result != nil || !strings.Contains(job.Error, "redirect URL must not contain credentials") {
			t.Fatalf("userinfo redirect job: %+v", job)
		}
		if docsRequests != 1 || schemaRequests != 1 || targetRequests != 0 {
			t.Fatalf("userinfo requests: docs=%d schema=%d target=%d", docsRequests, schemaRequests, targetRequests)
		}
		assertWorkerPublishedNothing(t, store, job)
	})

	t.Run("same origin", func(t *testing.T) {
		schemaRequests, finalRequests, clientRedirects := 0, 0, 0
		schema := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			if request.URL.Path == "/final.json" {
				finalRequests++
				fmt.Fprint(writer, scalarWorkerSchema(1))
				return
			}
			schemaRequests++
			http.Redirect(writer, request, "/final.json", http.StatusFound)
		}))
		defer schema.Close()
		client := schema.Client()
		client.CheckRedirect = func(*http.Request, []*http.Request) error {
			clientRedirects++
			return nil
		}

		store, job, docsRequests := runDiscoveredScalarWorker(t, client, schema.URL+"/openapi.json")
		if job.State != jobSucceeded || job.Result == nil || job.Result.Framework != "scalar" || job.Result.Source != schema.URL+"/final.json" || job.Result.Pages != 2 {
			t.Fatalf("same-origin redirect job: %+v", job)
		}
		if docsRequests != 1 || schemaRequests != 1 || finalRequests != 1 || clientRedirects != 1 {
			t.Fatalf("same-origin requests: docs=%d schema=%d final=%d redirects=%d", docsRequests, schemaRequests, finalRequests, clientRedirects)
		}
		if _, err := os.Stat(filepath.Join(store.root, job.ID+".published.json")); err != nil {
			t.Fatalf("same-origin publication receipt: %v", err)
		}
	})
}

func TestWorkerScopesDocusaurusToStartingPath(t *testing.T) {
	blogRequests := 0
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/docs/start":
			fmt.Fprint(writer, docusaurusPage("Start", `<a href="/docs/second">Second</a><a href="/blog">Blog</a>`))
		case "/docs/second":
			fmt.Fprint(writer, docusaurusPage("Second", `<a href="/docs/start">Start</a>`))
		case "/blog":
			blogRequests++
			fmt.Fprint(writer, `<html><body><h1>Blog</h1></body></html>`)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	store, err := openJobStore(t.TempDir(), true)
	if err != nil {
		t.Fatal(err)
	}
	job, err := store.create(ingestRequest{
		Output: store.output, Source: server.URL + "/docs/start", Name: "Docs", Version: "v1",
		Scope: "path", MaxPages: -1, MaxDepth: -1,
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := runWorker(context.Background(), store, job.ID, server.Client()); err != nil {
		t.Fatal(err)
	}
	job, _ = store.get(job.ID)
	if job.State != jobSucceeded || job.Detection == nil || job.Detection.Engine != "html" || job.Detection.Framework != "docusaurus" || job.Result == nil || job.Result.Pages != 2 || job.Result.Truncated || blogRequests != 0 {
		t.Fatalf("unexpected scoped job: %+v, blog requests=%d", job, blogRequests)
	}
}

func TestWorkerExhaustsMDBookTOC(t *testing.T) {
	firstChapterRequests := 0
	secondRequests := 0
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/book/":
			fmt.Fprint(writer, `<!doctype html><html><head><!-- Book generated using mdBook --><title>Worker Book</title><script src="toc-fixture.js"></script></head><body><nav id="mdbook-sidebar"><noscript><iframe class="sidebar-iframe-outer" src="toc.html"></iframe></noscript></nav><div id="mdbook-content"><main><h1><a class="header" href="#book">Worker Book</a></h1><p>Durable mdBook dispatch.</p></main></div></body></html>`)
		case "/book/toc.html":
			fmt.Fprint(writer, `<html><body class="sidebar-iframe-inner"><ol class="chapter"><li><a href="index.html">Worker Book</a></li><li><a href="second.html">Second</a></li></ol></body></html>`)
		case "/book/index.html":
			firstChapterRequests++
			fmt.Fprint(writer, `<!doctype html><html><head><!-- Book generated using mdBook --><title>Worker Book</title><script src="toc-fixture.js"></script></head><body><nav id="mdbook-sidebar"></nav><div id="mdbook-content"><main><h1><a class="header" href="#book">Worker Book</a></h1><p>Durable mdBook dispatch.</p></main></div></body></html>`)
		case "/book/second.html":
			secondRequests++
			fmt.Fprint(writer, `<!doctype html><html><head><!-- Book generated using mdBook --><title>Second - Worker Book</title><script src="toc-fixture.js"></script></head><body><nav id="mdbook-sidebar"></nav><div id="mdbook-content"><main><h1><a class="header" href="#second">Second</a></h1><p>TOC-only worker page.</p></main></div></body></html>`)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	store, err := openJobStore(t.TempDir(), true)
	if err != nil {
		t.Fatal(err)
	}
	job, err := store.create(ingestRequest{
		Output: store.output, Source: server.URL + "/book/", Name: "Worker mdBook", Version: "v1",
		Scope: "path", MaxPages: -1, MaxDepth: -1,
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := runWorker(context.Background(), store, job.ID, server.Client()); err != nil {
		t.Fatal(err)
	}
	job, err = store.get(job.ID)
	if err != nil {
		t.Fatal(err)
	}
	if job.State != jobSucceeded || job.Detection == nil || job.Detection.Engine != "html" || job.Detection.Framework != "mdbook" || job.Result == nil || job.Result.Framework != "mdbook" || job.Result.Pages != 2 || job.Result.Truncated || firstChapterRequests != 1 || secondRequests != 1 {
		t.Fatalf("unexpected mdBook job: %+v, first requests=%d second requests=%d", job, firstChapterRequests, secondRequests)
	}
}

func TestWorkerDispatchesDocsifyImporter(t *testing.T) {
	const commit = "0123456789abcdef0123456789abcdef01234567"
	const tree = "abcdef0123456789abcdef0123456789abcdef01"
	client := &http.Client{Transport: workerRoundTripFunc(func(request *http.Request) (*http.Response, error) {
		switch {
		case request.URL.Host == "docs.example":
			return workerResponse(request, http.StatusOK, `<!doctype html><html><body><div id="app"></div><script>window.$docsify = {}; const edit = 'https://github.com/example/docs/tree/main/docs/' + route;</script><script src="https://cdn.jsdelivr.net/npm/docsify@5"></script></body></html>`), nil
		case request.URL.Host == "api.github.com" && strings.HasSuffix(request.URL.Path, "/commits/main/docs"):
			return workerResponse(request, http.StatusNotFound, `{}`), nil
		case request.URL.Host == "api.github.com" && strings.HasSuffix(request.URL.Path, "/commits/main"):
			return workerResponse(request, http.StatusOK, `{"sha":"`+commit+`","commit":{"tree":{"sha":"`+tree+`"}}}`), nil
		case request.URL.Host == "api.github.com" && strings.HasSuffix(request.URL.Path, "/git/trees/"+tree):
			return workerResponse(request, http.StatusOK, `{"sha":"`+tree+`","truncated":false,"tree":[{"path":"docs/README.md","type":"blob","size":25}]}`), nil
		case request.URL.Host == "raw.githubusercontent.com" && strings.HasSuffix(request.URL.Path, "/docs/README.md"):
			return workerResponse(request, http.StatusOK, "# Worker Docsify\n\nDispatch works.\n"), nil
		default:
			return workerResponse(request, http.StatusNotFound, `{}`), nil
		}
	})}
	store, err := openJobStore(t.TempDir(), true)
	if err != nil {
		t.Fatal(err)
	}
	job, err := store.create(ingestRequest{
		Output: store.output, Source: "https://docs.example/", Name: "Worker Docs", Version: "v1",
		Scope: "path", MaxPages: -1, MaxDepth: -1,
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := runWorker(context.Background(), store, job.ID, client); err != nil {
		t.Fatal(err)
	}
	job, err = store.get(job.ID)
	if err != nil {
		t.Fatal(err)
	}
	if job.State != jobSucceeded || job.Detection == nil || job.Detection.Engine != "docsify" || job.Result == nil || job.Result.Framework != "docsify" || job.Result.Pages != 1 || job.Result.Sources != 1 {
		t.Fatalf("unexpected Docsify job: %+v", job)
	}
	generated := filepath.Join(store.output, "worker-docs", "v1", "documentation", "example", "docs", commit, "docs", "README.md")
	if raw, err := os.ReadFile(generated); err != nil || !strings.Contains(string(raw), "Dispatch works.") {
		t.Fatalf("Docsify worker output: %q, %v", raw, err)
	}
}

func TestWorkerRefusesUnsupportedAndUnknownHTMLWithoutImporting(t *testing.T) {
	const unknownRefusal = "automatic ingestion refused HTML with an unknown framework: generic anchor crawling has no finite completeness inventory; explicit `apis-mcp import html` remains available for intentional best-effort imports"
	for _, test := range []struct {
		name      string
		html      string
		framework string
		refusal   string
	}{
		{name: "mintlify", html: `<meta name="generator" content="Mintlify"><main>Docs</main>`, framework: "mintlify", refusal: "automatic ingestion detected unsupported documentation framework mintlify; no complete importer is available"},
		{name: "mintlify with scalar", html: `<meta name="generator" content="Mintlify"><scalar-api-reference configuration='{"url":"/openapi.json"}'></scalar-api-reference>`, framework: "mintlify", refusal: "automatic ingestion detected unsupported documentation framework mintlify; no complete importer is available"},
		{name: "scalar", html: `<scalar-api-reference configuration="{}"></scalar-api-reference>`, framework: "scalar", refusal: "automatic ingestion detected unsupported documentation framework scalar; no complete importer is available"},
		{name: "astro", html: `<meta name="generator" content="Astro v5.13.5"><main>Docs</main>`, framework: "astro", refusal: "automatic ingestion detected unsupported documentation framework astro; no complete importer is available"},
		{name: "sveltekit", html: `<body data-sveltekit-preload-data="hover"><script>window.__sveltekit_app = {}; import("/_app/immutable/entry/start.js")</script></body>`, framework: "sveltekit", refusal: "automatic ingestion detected unsupported documentation framework sveltekit; no complete importer is available"},
		{name: "stripe docs", html: `<link rel="stylesheet" href="https://b.stripecdn.com/docs-statics-srv/assets/sail.205757ec.css"><main class="Article--ApiReference">Docs</main>`, framework: "stripe-docs", refusal: "automatic ingestion detected unsupported documentation framework stripe-docs; no complete importer is available"},
		{name: "stripe docs with scalar", html: `<link rel="stylesheet" href="https://b.stripecdn.com/docs-statics-srv/assets/sail.205757ec.css"><main class="Article--ApiReference"><scalar-api-reference configuration='{"url":"/openapi.json"}'></scalar-api-reference></main>`, framework: "stripe-docs", refusal: "automatic ingestion detected unsupported documentation framework stripe-docs; no complete importer is available"},
		{name: "unknown static", html: `<main><h1>Generic Docs</h1><a href="/second">Second</a></main>`, framework: "unknown", refusal: unknownRefusal},
		{name: "unknown client shell", html: `<div id="app">Loading documentation</div><script src="/assets/app.js"></script>`, framework: "unknown", refusal: unknownRefusal},
	} {
		t.Run(test.name, func(t *testing.T) {
			requests := 0
			server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
				requests++
				fmt.Fprint(writer, `<!doctype html><html>`+test.html+`</html>`)
			}))
			defer server.Close()

			store, err := openJobStore(t.TempDir(), true)
			if err != nil {
				t.Fatal(err)
			}
			job, err := store.create(ingestRequest{
				Output: store.output, Source: server.URL + "/docs", Name: "Unknown " + test.name, Version: "v1",
				Scope: "path", MaxPages: -1, MaxDepth: -1,
			})
			if err != nil {
				t.Fatal(err)
			}
			if err := runWorker(context.Background(), store, job.ID, server.Client()); err != nil {
				t.Fatal(err)
			}
			job, err = store.get(job.ID)
			if err != nil {
				t.Fatal(err)
			}
			if job.State != jobFailed || job.Detection == nil || job.Detection.Engine != "html" || job.Detection.Framework != test.framework || job.Result != nil || job.Error != test.refusal {
				t.Fatalf("unexpected refused job: %+v", job)
			}
			if requests != 1 {
				t.Fatalf("requests = %d, want detection request only", requests)
			}
			destination := filepath.Join(store.output, importer.SafeSlug(job.Request.Name), importer.SafeSlug(job.Request.Version))
			if _, err := os.Stat(destination); !os.IsNotExist(err) {
				t.Fatalf("refused destination exists: %v", err)
			}
			for _, name := range []string{"library.sqlite", filepath.Join(".ingest", "jobs", job.ID+".published.json")} {
				if _, err := os.Stat(filepath.Join(store.output, name)); !os.IsNotExist(err) {
					t.Fatalf("refused ingestion artifact %s exists: %v", name, err)
				}
			}
			indexes, err := filepath.Glob(filepath.Join(store.output, "library*.sqlite"))
			if err != nil || len(indexes) != 0 {
				t.Fatalf("refused ingestion indexes = %v, %v", indexes, err)
			}
		})
	}
}

func TestSupportedHTMLFrameworkAllowlist(t *testing.T) {
	for _, framework := range []string{"docusaurus", "mkdocs-material", "mkdocs", "sphinx", "vitepress", "nextra", "astro-starlight", "mdbook"} {
		if !supportedHTMLFramework(framework) {
			t.Errorf("supported framework %q refused", framework)
		}
	}
	for _, framework := range []string{"", "unknown", "mintlify", "scalar", "astro", "sveltekit", "stripe-docs"} {
		if supportedHTMLFramework(framework) {
			t.Errorf("unsupported framework %q allowed", framework)
		}
	}
}

func TestWorkerCancellationIsPersisted(t *testing.T) {
	started := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		close(started)
		<-request.Context().Done()
	}))
	defer server.Close()

	store, err := openJobStore(t.TempDir(), true)
	if err != nil {
		t.Fatal(err)
	}
	job, err := store.create(ingestRequest{
		Output: store.output, Source: server.URL + "/docs", Name: "Canceled", Version: "v1",
		Scope: "path", MaxPages: -1, MaxDepth: -1,
	})
	if err != nil {
		t.Fatal(err)
	}
	workerDone := make(chan error, 1)
	go func() { workerDone <- runWorker(context.Background(), store, job.ID, server.Client()) }()
	select {
	case <-started:
	case <-time.After(5 * time.Second):
		t.Fatal("worker did not start request")
	}
	if _, err := store.cancel(job.ID); err != nil {
		t.Fatal(err)
	}
	select {
	case err := <-workerDone:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("worker did not stop")
	}
	job, err = store.get(job.ID)
	if err != nil {
		t.Fatal(err)
	}
	if job.State != jobCanceled || !job.CancelRequested {
		t.Fatalf("unexpected canceled job: %+v", job)
	}
	if _, err := os.Stat(filepath.Join(store.output, "canceled", "v1")); !os.IsNotExist(err) {
		t.Fatalf("canceled destination exists: %v", err)
	}
}

func TestQueuedJobCanBeCanceledWithoutWorker(t *testing.T) {
	store, err := openJobStore(t.TempDir(), true)
	if err != nil {
		t.Fatal(err)
	}
	job, err := store.create(ingestRequest{
		Output: store.output, Source: "https://example.test/docs", Name: "Queued", Version: "v1",
		Scope: "path", MaxPages: -1, MaxDepth: -1,
	})
	if err != nil {
		t.Fatal(err)
	}
	job, err = store.cancel(job.ID)
	if err != nil || job.State != jobCanceled || !job.CancelRequested {
		t.Fatalf("queued cancellation: %+v, %v", job, err)
	}
}

func TestJobRejectsCredentialBearingSourceBeforePersistence(t *testing.T) {
	store, err := openJobStore(t.TempDir(), true)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := store.create(ingestRequest{Output: store.output, Source: "https://user:secret@docs.example/book/", Name: "Private", Version: "v1", Scope: "path", MaxPages: -1, MaxDepth: -1}); err == nil {
		t.Fatal("credential-bearing job source accepted")
	}
	files, err := filepath.Glob(filepath.Join(store.root, "*.json"))
	if err != nil || len(files) != 0 {
		t.Fatalf("credential-bearing job persisted files: %v, %v", files, err)
	}
}

func TestListIgnoresSucceededPublicationReceipt(t *testing.T) {
	store, err := openJobStore(t.TempDir(), true)
	if err != nil {
		t.Fatal(err)
	}
	job, err := store.create(ingestRequest{Output: store.output, Source: "https://example.test/docs", Name: "Listed", Version: "v1"})
	if err != nil {
		t.Fatal(err)
	}
	job, completed, err := store.finishSuccess(job.ID, importer.Result{Name: "Listed", Version: "v1", Source: job.Request.Source})
	if err != nil || !completed {
		t.Fatalf("finish success: %+v, %t, %v", job, completed, err)
	}
	if _, err := os.Stat(filepath.Join(store.root, job.ID+".published.json")); err != nil {
		t.Fatal(err)
	}

	var stdout, stderr bytes.Buffer
	if status := runList([]string{"-out", store.output}, &stdout, &stderr); status != 0 {
		t.Fatalf("list status = %d, stderr = %q", status, stderr.String())
	}
	var jobs []ingestJob
	if err := json.Unmarshal(stdout.Bytes(), &jobs); err != nil {
		t.Fatal(err)
	}
	if len(jobs) != 1 || jobs[0].ID != job.ID || jobs[0].State != jobSucceeded {
		t.Fatalf("listed jobs: %+v", jobs)
	}
}

func TestListReturnsMalformedCanonicalStateError(t *testing.T) {
	store, err := openJobStore(t.TempDir(), true)
	if err != nil {
		t.Fatal(err)
	}
	job, err := store.create(ingestRequest{Output: store.output, Source: "https://example.test/docs", Name: "Malformed", Version: "v1"})
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(store.root, job.ID+".json"), []byte("not JSON\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := store.list(); err == nil || !strings.Contains(err.Error(), "read job "+job.ID) {
		t.Fatalf("malformed canonical state error: %v", err)
	}
}

func TestRunValidatesCommands(t *testing.T) {
	for _, args := range [][]string{nil, {"start"}, {"unknown"}, {"status", "-out", t.TempDir(), "invalid"}} {
		var stdout, stderr bytes.Buffer
		if status := run(context.Background(), args, &stdout, &stderr, nil); status == 0 {
			t.Fatalf("args %v unexpectedly succeeded", args)
		}
	}
}

func docusaurusPage(title, links string) string {
	return `<!doctype html><html><head><meta name="generator" content="Docusaurus v3"></head><body>` +
		`<ul class="theme-doc-sidebar-menu">` + links + `</ul>` +
		`<div class="theme-doc-markdown"><header><h1>` + title + `</h1></header><p>Content for ` + title + `.</p></div></body></html>`
}

func scalarWorkerSchema(operations int) string {
	var schema strings.Builder
	schema.WriteString(`{"openapi":"3.0.3","info":{"title":"Scalar Worker","version":"live"},"paths":{`)
	for index := range operations {
		if index > 0 {
			schema.WriteByte(',')
		}
		fmt.Fprintf(&schema, `"/worker/%03d":{"get":{"operationId":"workerOperation%03d","summary":"Worker operation %03d","description":"scalarNeedle%03d","tags":["worker"],"responses":{"200":{"description":"ok"}}}}`, index, index, index, index)
	}
	schema.WriteString(`}}`)
	return schema.String()
}

func runDiscoveredScalarWorker(t *testing.T, client *http.Client, schemaURL string) (*jobStore, ingestJob, int) {
	t.Helper()
	docsRequests := 0
	docs := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		docsRequests++
		fmt.Fprintf(writer, `<!doctype html><html><body><scalar-api-reference data-url=%q></scalar-api-reference></body></html>`, schemaURL)
	}))
	defer docs.Close()
	store, err := openJobStore(t.TempDir(), true)
	if err != nil {
		t.Fatal(err)
	}
	job, err := store.create(ingestRequest{
		Output: store.output, Source: docs.URL + "/api-docs", Name: "Scalar Redirect", Version: "live",
		Scope: "path", MaxPages: -1, MaxDepth: -1,
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := runWorker(context.Background(), store, job.ID, client); err != nil {
		t.Fatal(err)
	}
	job, err = store.get(job.ID)
	if err != nil {
		t.Fatal(err)
	}
	return store, job, docsRequests
}

func assertWorkerPublishedNothing(t *testing.T, store *jobStore, job ingestJob) {
	t.Helper()
	if _, err := os.Stat(filepath.Join(store.output, "scalar-redirect", "live")); !os.IsNotExist(err) {
		t.Fatalf("failed redirect destination exists: %v", err)
	}
	if _, err := os.Stat(filepath.Join(store.root, job.ID+".published.json")); !os.IsNotExist(err) {
		t.Fatalf("failed redirect publication receipt exists: %v", err)
	}
	indexes, err := filepath.Glob(filepath.Join(store.output, "library*.sqlite"))
	if err != nil || len(indexes) != 0 {
		t.Fatalf("failed redirect indexes = %v, %v", indexes, err)
	}
}
