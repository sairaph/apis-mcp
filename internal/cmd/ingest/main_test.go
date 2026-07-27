//go:build dev

package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

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
	if job.State != jobSucceeded || job.Result == nil || job.Result.Pages != 2 || job.Result.Truncated || blogRequests != 0 {
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
