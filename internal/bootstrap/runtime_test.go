package bootstrap

import (
	"archive/zip"
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/sairaph/apis-mcp/internal/cache"
	"github.com/sairaph/apis-mcp/internal/config"
	"github.com/sairaph/apis-mcp/internal/docpacks"
	"github.com/sairaph/apis-mcp/internal/httpcall"
	"github.com/sairaph/apis-mcp/internal/sessions"
	"github.com/sairaph/apis-mcp/library"
)

func TestRuntimePeriodicCleanupAndShutdown(t *testing.T) {
	root := t.TempDir()
	cacheRoot := filepath.Join(root, "cache")
	sessionRoot := filepath.Join(root, "sessions")
	store, err := cache.New(cache.Config{Root: cacheRoot, Retention: time.Millisecond})
	if err != nil {
		t.Fatal(err)
	}
	manager, err := sessions.New(sessionRoot, time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	handle, err := manager.Create(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	sessionPath := filepath.Join(sessionRoot, handle.ID()+".json")
	if err := handle.Close(); err != nil {
		t.Fatal(err)
	}
	orphan := filepath.Join(cacheRoot, "orphan")
	if err := os.Mkdir(orphan, 0o700); err != nil {
		t.Fatal(err)
	}
	time.Sleep(2 * time.Millisecond)

	runtime := &Runtime{
		Paths: config.Paths{Diagnostics: filepath.Join(root, "diagnostics.log")},
		Cache: store, Sessions: manager,
	}
	runtime.startCleanup(2 * time.Millisecond)
	waitFor(t, time.Second, func() bool {
		_, sessionErr := os.Stat(sessionPath)
		_, cacheErr := os.Stat(orphan)
		return os.IsNotExist(sessionErr) && os.IsNotExist(cacheErr)
	})

	if err := runtime.Close(context.Background()); err != nil {
		t.Fatal(err)
	}
	select {
	case <-runtime.cleanupDone:
	default:
		t.Fatal("cleanup loop was not stopped before Close returned")
	}
	afterClose := filepath.Join(cacheRoot, "after-close")
	if err := os.Mkdir(afterClose, 0o700); err != nil {
		t.Fatal(err)
	}
	time.Sleep(10 * time.Millisecond)
	if _, err := os.Stat(afterClose); err != nil {
		t.Fatalf("cleanup ran after Close: %v", err)
	}
	if err := runtime.Close(context.Background()); err != nil {
		t.Fatalf("second Close: %v", err)
	}
}

func TestRuntimePeriodicCleanupReportsFailures(t *testing.T) {
	root := t.TempDir()
	store, err := cache.New(cache.Config{Root: filepath.Join(root, "cache"), Retention: time.Hour})
	if err != nil {
		t.Fatal(err)
	}
	sessionRoot := filepath.Join(root, "sessions")
	manager, err := sessions.New(sessionRoot, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.RemoveAll(sessionRoot); err != nil {
		t.Fatal(err)
	}
	diagnosticPath := filepath.Join(root, "diagnostics.log")
	runtime := &Runtime{
		Paths: config.Paths{Diagnostics: diagnosticPath}, Cache: store, Sessions: manager,
	}
	runtime.startCleanup(time.Millisecond)
	waitFor(t, time.Second, func() bool { return len(runtime.Diagnostics()) > 0 })
	if err := runtime.Close(context.Background()); err != nil {
		t.Fatal(err)
	}
	raw, err := os.ReadFile(diagnosticPath)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(raw), "periodic cleanup") || !strings.Contains(string(raw), "sessions") {
		t.Fatalf("diagnostic = %q", raw)
	}
}

func TestOpenReportsInitialCleanupFailureWithoutFailingStartup(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	sessionRoot := filepath.Join(home, ".apis-mcp", "sessions")
	manager, err := sessions.New(sessionRoot, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	handle, err := manager.Create(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(sessionRoot, handle.ID()+".json")
	if err := handle.Close(); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte("{"), 0o600); err != nil {
		t.Fatal(err)
	}

	runtime, err := Open(context.Background())
	if err != nil {
		t.Fatalf("Open failed for a cleanup diagnostic: %v", err)
	}
	defer runtime.Close(context.Background())
	if diagnostics := runtime.Diagnostics(); len(diagnostics) == 0 || !strings.Contains(diagnostics[0].Error(), "initial cleanup") {
		t.Fatalf("runtime diagnostics = %v", diagnostics)
	}
	raw, readErr := os.ReadFile(filepath.Join(home, ".apis-mcp", "diagnostics.log"))
	if readErr != nil {
		t.Fatal(readErr)
	}
	if !strings.Contains(string(raw), "initial cleanup") {
		t.Fatalf("diagnostic = %q", raw)
	}
}

func TestCloseBoundsCleanupWaitWithContext(t *testing.T) {
	runtime := &Runtime{
		cleanupCancel: func() {},
		cleanupDone:   make(chan struct{}),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	started := time.Now()
	err := runtime.Close(ctx)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("Close error = %v", err)
	}
	if elapsed := time.Since(started); elapsed > time.Second {
		t.Fatalf("Close was not bounded: %v", elapsed)
	}
	close(runtime.cleanupDone)
}

func TestOpenWiresReadTokenBudgetToHTTPPreview(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	paths, err := config.DefaultPaths()
	if err != nil {
		t.Fatal(err)
	}
	cfg := config.Default()
	cfg.ReadTokenBudget = 2
	cfg.FreeDiskReserve = 0
	if err := config.Save(paths, cfg); err != nil {
		t.Fatal(err)
	}
	runtime, err := Open(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	defer runtime.Close(context.Background())

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "hello world again")
	}))
	defer server.Close()
	zero := 0
	result, err := runtime.HTTP.Call(context.Background(), httpcall.Input{
		Method: "GET", Endpoint: server.URL, Retries: &zero,
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Preview == nil || !result.Preview.Truncated || result.Preview.ApproximateTokens > cfg.ReadTokenBudget {
		t.Fatalf("preview did not use configured read budget: %#v", result.Preview)
	}
}

func TestOpenLoadsActivePacksWithoutRefreshing(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	paths, err := config.DefaultPaths()
	if err != nil {
		t.Fatal(err)
	}
	pack, archive := runtimeTestPack(t)
	catalog := docpacks.Catalog{SchemaVersion: 1, Packs: []docpacks.Pack{pack}}
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/packs/"+pack.Asset {
			http.NotFound(writer, request)
			return
		}
		_, _ = writer.Write(archive)
	}))
	manager, err := docpacks.Open(paths.Packs, docpacks.Options{CatalogURL: server.URL + "/packs/catalog.json"})
	if err != nil {
		t.Fatal(err)
	}
	if err := manager.Apply(context.Background(), catalog, []string{pack.ID}, func(context.Context, []string) error { return nil }, nil); err != nil {
		t.Fatal(err)
	}
	server.Close()

	runtime, err := Open(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	defer runtime.Close(context.Background())
	listed, err := runtime.Library.List(context.Background(), library.ListRequest{Name: "Startup"})
	if err != nil {
		t.Fatal(err)
	}
	if runtime.Packs == nil || listed.Total != 1 || len(listed.APIs) != 1 || listed.APIs[0].Name != pack.Name {
		t.Fatalf("runtime did not load active pack: packs=%v list=%+v", runtime.Packs, listed)
	}
	if options := runtime.LibraryOptions(); len(options.PackArchives) != 1 || filepath.Base(options.PackArchives[0]) != pack.SHA256+".zip" {
		t.Fatalf("runtime library options = %+v", options)
	}
	if err := manager.Apply(context.Background(), catalog, nil, func(context.Context, []string) error { return nil }, nil); err != nil {
		t.Fatal(err)
	}
	if err := runtime.RebuildLibrary(context.Background()); err != nil {
		t.Fatal(err)
	}
	if options := runtime.LibraryOptions(); len(options.PackArchives) != 0 {
		t.Fatalf("rebuild did not reread active packs: %+v", options)
	}
}

func waitFor(t *testing.T, timeout time.Duration, condition func() bool) {
	t.Helper()
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if condition() {
			return
		}
		time.Sleep(time.Millisecond)
	}
	t.Fatal("condition was not met before timeout")
}

func runtimeTestPack(t *testing.T) (docpacks.Pack, []byte) {
	t.Helper()
	manifest := "---\nname: Startup API\nversion: v1\ndescription: Startup docs.\ncollections: [examples]\n---\n"
	page := "---\ntitle: Overview\n---\n\n# Startup\n"
	var buffer bytes.Buffer
	writer := zip.NewWriter(&buffer)
	for name, content := range map[string]string{
		"startup-api/v1/_index.md":   manifest,
		"startup-api/v1/overview.md": page,
	} {
		entry, err := writer.Create(name)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := entry.Write([]byte(content)); err != nil {
			t.Fatal(err)
		}
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	archive := buffer.Bytes()
	hash := sha256.Sum256(archive)
	return docpacks.Pack{
		ID: "startup-api", Name: "Startup API", Description: "Startup docs.", Asset: "startup-api-" + hex.EncodeToString(hash[:]) + ".zip",
		SHA256: hex.EncodeToString(hash[:]), Bytes: int64(len(archive)), UncompressedBytes: int64(len(manifest) + len(page)),
		Files: 2, Pages: 1, Versions: []string{"v1"}, Collections: []string{"examples"},
	}, archive
}
