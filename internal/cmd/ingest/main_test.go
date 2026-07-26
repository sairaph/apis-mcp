//go:build dev

package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/sairaph/apis-mcp/library"
)

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
	if err != nil || recovered.State != jobSucceeded || recovered.Result == nil || recovered.Result.Kind != "openapi" {
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
