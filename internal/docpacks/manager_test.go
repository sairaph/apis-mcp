package docpacks

import (
	"archive/zip"
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func TestRefreshCachesAndFallsBack(t *testing.T) {
	pack, _ := testPack(t, "alpha", "Alpha", "v1")
	catalog := Catalog{SchemaVersion: 1, Packs: []Pack{pack}}
	var unavailable atomic.Bool
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/packs/catalog.json" {
			http.NotFound(writer, request)
			return
		}
		if unavailable.Load() {
			http.Error(writer, "offline", http.StatusServiceUnavailable)
			return
		}
		_ = json.NewEncoder(writer).Encode(catalog)
	}))
	defer server.Close()

	manager := openTestManager(t, server.URL+"/packs/catalog.json")
	refreshed, err := manager.Refresh(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	unavailable.Store(true)
	fallback, err := manager.Refresh(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	cached, err := manager.CachedCatalog()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(refreshed, catalog) || !reflect.DeepEqual(fallback, catalog) || !reflect.DeepEqual(cached, catalog) {
		t.Fatalf("catalog refresh/fallback mismatch: refreshed=%+v fallback=%+v cached=%+v", refreshed, fallback, cached)
	}
	raw, err := os.ReadFile(filepath.Join(manager.root, "catalog.json"))
	if err != nil {
		t.Fatal(err)
	}
	var compact bytes.Buffer
	if err := json.Compact(&compact, bytes.TrimSpace(raw)); err != nil || !bytes.Equal(compact.Bytes(), bytes.TrimSpace(raw)) {
		t.Fatalf("cached catalog is not compact JSON: %v", err)
	}
}

func TestApplyDownloadsAndRemovesWithoutDeletingBlobs(t *testing.T) {
	pack, archive := testPack(t, "alpha", "Alpha", "v1")
	catalog := Catalog{SchemaVersion: 1, Packs: []Pack{pack}}
	var assetRequests atomic.Int32
	server := packServer(t, catalog, map[string][]byte{pack.Asset: archive}, &assetRequests)
	defer server.Close()
	manager := openTestManager(t, server.URL+"/packs/catalog.json")
	refreshed, err := manager.Refresh(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	var rebuilt [][]string
	var visible []ActiveState
	rebuild := func(_ context.Context, archives []string) error {
		rebuilt = append(rebuilt, append([]string(nil), archives...))
		active, err := manager.Active()
		if err != nil {
			return err
		}
		visible = append(visible, active)
		return nil
	}
	if err := manager.Apply(context.Background(), refreshed, []string{pack.ID}, rebuild); err != nil {
		t.Fatal(err)
	}
	active, err := manager.Active()
	if err != nil {
		t.Fatal(err)
	}
	archives, err := manager.ActiveArchives()
	if err != nil {
		t.Fatal(err)
	}
	if len(active.Packs) != 1 || active.Packs[pack.ID].SHA256 != pack.SHA256 || len(archives) != 1 || len(rebuilt) != 1 || !reflect.DeepEqual(rebuilt[0], archives) || len(visible[0].Packs) != 0 {
		t.Fatalf("unexpected active state: active=%+v archives=%v rebuilt=%v", active, archives, rebuilt)
	}
	blob := archives[0]
	if pack.Asset != pack.ID+"-"+pack.SHA256+".zip" || filepath.Base(blob) != pack.SHA256+".zip" || assetRequests.Load() != 1 {
		t.Fatalf("unexpected blob/download: %s requests=%d", blob, assetRequests.Load())
	}
	activePath := filepath.Join(manager.root, "active.json")
	oldTime := time.Now().Add(-time.Hour).Truncate(time.Second)
	if err := os.Chtimes(activePath, oldTime, oldTime); err != nil {
		t.Fatal(err)
	}
	activeInfo, err := os.Stat(activePath)
	if err != nil {
		t.Fatal(err)
	}
	if err := manager.Apply(context.Background(), refreshed, []string{pack.ID}, rebuild); err != nil {
		t.Fatal(err)
	}
	afterNoop, err := os.Stat(activePath)
	if err != nil {
		t.Fatal(err)
	}
	if len(rebuilt) != 1 || assetRequests.Load() != 1 || !afterNoop.ModTime().Equal(activeInfo.ModTime()) {
		t.Fatalf("no-op apply rebuilt or rewrote state: rebuilds=%d requests=%d before=%v after=%v", len(rebuilt), assetRequests.Load(), activeInfo.ModTime(), afterNoop.ModTime())
	}

	userFile := filepath.Join(filepath.Dir(manager.root), "library", "user.md")
	if err := os.MkdirAll(filepath.Dir(userFile), 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(userFile, []byte("user"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := manager.Apply(context.Background(), refreshed, nil, rebuild); err != nil {
		t.Fatal(err)
	}
	active, err = manager.Active()
	if err != nil {
		t.Fatal(err)
	}
	if len(active.Packs) != 0 || len(rebuilt) != 2 || len(rebuilt[1]) != 0 || len(visible) != 2 || len(visible[1].Packs) != 1 {
		t.Fatalf("removal did not publish an empty state: active=%+v rebuilt=%v", active, rebuilt)
	}
	if _, err := os.Stat(blob); err != nil {
		t.Fatalf("removal deleted verified blob: %v", err)
	}
	if raw, err := os.ReadFile(userFile); err != nil || string(raw) != "user" {
		t.Fatalf("removal changed user library: %q, %v", raw, err)
	}
}

func TestApplyRejectsSHAMismatch(t *testing.T) {
	pack, archive := testPack(t, "alpha", "Alpha", "v1")
	pack.SHA256 = strings.Repeat("0", 64)
	pack.Asset = pack.ID + "-" + pack.SHA256 + ".zip"
	catalog := Catalog{SchemaVersion: 1, Packs: []Pack{pack}}
	server := packServer(t, catalog, map[string][]byte{pack.Asset: archive}, nil)
	defer server.Close()
	manager := openTestManager(t, server.URL+"/packs/catalog.json")
	err := manager.Apply(context.Background(), catalog, []string{pack.ID}, func(context.Context, []string) error {
		t.Fatal("rebuild called for invalid blob")
		return nil
	})
	if err == nil || !strings.Contains(err.Error(), "SHA-256") {
		t.Fatalf("SHA mismatch error = %v", err)
	}
	active, activeErr := manager.Active()
	if activeErr != nil || len(active.Packs) != 0 {
		t.Fatalf("invalid blob changed active state: %+v, %v", active, activeErr)
	}
}

func TestApplyRejectsMalformedAndTraversalZIPs(t *testing.T) {
	tests := []struct {
		name    string
		archive []byte
		files   int
		pages   int
		bytes   int64
	}{
		{name: "malformed", archive: []byte("not a zip"), files: 1, pages: 0, bytes: 1},
		{name: "traversal", archive: traversalArchive(t), files: 2, pages: 1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			pack := packForArchive("alpha", "Alpha", "v1", test.archive, test.files, test.pages, test.bytes)
			catalog := Catalog{SchemaVersion: 1, Packs: []Pack{pack}}
			server := packServer(t, catalog, map[string][]byte{pack.Asset: test.archive}, nil)
			defer server.Close()
			manager := openTestManager(t, server.URL+"/packs/catalog.json")
			err := manager.Apply(context.Background(), catalog, []string{pack.ID}, func(context.Context, []string) error { return nil })
			if err == nil {
				t.Fatal("invalid ZIP was accepted")
			}
		})
	}
}

func TestApplyRebuildFailureLeavesPriorStateUnchanged(t *testing.T) {
	alpha, alphaArchive := testPack(t, "alpha", "Alpha", "v1")
	beta, betaArchive := testPack(t, "beta", "Beta", "v1")
	catalog := Catalog{SchemaVersion: 1, Packs: []Pack{alpha, beta}}
	server := packServer(t, catalog, map[string][]byte{alpha.Asset: alphaArchive, beta.Asset: betaArchive}, nil)
	defer server.Close()
	manager := openTestManager(t, server.URL+"/packs/catalog.json")
	if err := manager.Apply(context.Background(), catalog, []string{alpha.ID}, func(context.Context, []string) error { return nil }); err != nil {
		t.Fatal(err)
	}
	activePath := filepath.Join(manager.root, "active.json")
	priorRaw, err := os.ReadFile(activePath)
	if err != nil {
		t.Fatal(err)
	}

	var calls [][]string
	rebuildFailure := errors.New("rebuild failed")
	err = manager.Apply(context.Background(), catalog, []string{beta.ID}, func(_ context.Context, archives []string) error {
		calls = append(calls, append([]string(nil), archives...))
		visible, activeErr := manager.Active()
		if activeErr != nil {
			return activeErr
		}
		if len(visible.Packs) != 1 || visible.Packs[alpha.ID].ID != alpha.ID {
			return fmt.Errorf("active state changed before rebuild: %+v", visible)
		}
		return rebuildFailure
	})
	if !errors.Is(err, rebuildFailure) {
		t.Fatalf("Apply error = %v", err)
	}
	active, activeErr := manager.Active()
	if activeErr != nil {
		t.Fatal(activeErr)
	}
	afterRaw, err := os.ReadFile(activePath)
	if err != nil {
		t.Fatal(err)
	}
	if len(active.Packs) != 1 || active.Packs[alpha.ID].ID != alpha.ID || len(calls) != 1 || filepath.Base(calls[0][0]) != beta.SHA256+".zip" || !bytes.Equal(priorRaw, afterRaw) {
		t.Fatalf("failed transaction changed state: active=%+v calls=%v", active, calls)
	}
}

func TestApplyRepairsOrRemovesCorruptPriorBlob(t *testing.T) {
	pack, archive := testPack(t, "alpha", "Alpha", "v1")
	catalog := Catalog{SchemaVersion: 1, Packs: []Pack{pack}}

	t.Run("repair selected pack", func(t *testing.T) {
		var requests atomic.Int32
		server := packServer(t, catalog, map[string][]byte{pack.Asset: archive}, &requests)
		defer server.Close()
		manager := openTestManager(t, server.URL+"/packs/catalog.json")
		if err := manager.Apply(context.Background(), catalog, []string{pack.ID}, func(context.Context, []string) error { return nil }); err != nil {
			t.Fatal(err)
		}
		activePath := filepath.Join(manager.root, "active.json")
		oldTime := time.Now().Add(-time.Hour).Truncate(time.Second)
		if err := os.Chtimes(activePath, oldTime, oldTime); err != nil {
			t.Fatal(err)
		}
		before, err := os.Stat(activePath)
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(manager.blobPath(pack), make([]byte, len(archive)), 0o600); err != nil {
			t.Fatal(err)
		}
		callbacks := 0
		if err := manager.Apply(context.Background(), catalog, []string{pack.ID}, func(context.Context, []string) error {
			callbacks++
			return nil
		}); err != nil {
			t.Fatal(err)
		}
		after, err := os.Stat(activePath)
		if err != nil {
			t.Fatal(err)
		}
		if callbacks != 0 || requests.Load() != 2 || !after.ModTime().Equal(before.ModTime()) {
			t.Fatalf("repair was not a metadata no-op: callbacks=%d requests=%d", callbacks, requests.Load())
		}
		if _, err := manager.ActiveArchives(); err != nil {
			t.Fatalf("repaired pack is invalid: %v", err)
		}
	})

	t.Run("remove corrupt pack", func(t *testing.T) {
		server := packServer(t, catalog, map[string][]byte{pack.Asset: archive}, nil)
		defer server.Close()
		manager := openTestManager(t, server.URL+"/packs/catalog.json")
		if err := manager.Apply(context.Background(), catalog, []string{pack.ID}, func(context.Context, []string) error { return nil }); err != nil {
			t.Fatal(err)
		}
		blob := manager.blobPath(pack)
		if err := os.WriteFile(blob, []byte("corrupt"), 0o600); err != nil {
			t.Fatal(err)
		}
		callbacks := 0
		if err := manager.Apply(context.Background(), catalog, nil, func(_ context.Context, archives []string) error {
			callbacks++
			if len(archives) != 0 {
				t.Fatalf("removal archives = %v", archives)
			}
			visible, err := manager.Active()
			if err != nil || len(visible.Packs) != 1 {
				t.Fatalf("prior state was not visible during rebuild: %+v, %v", visible, err)
			}
			return nil
		}); err != nil {
			t.Fatal(err)
		}
		active, err := manager.Active()
		if err != nil || len(active.Packs) != 0 || callbacks != 1 {
			t.Fatalf("corrupt pack removal failed: active=%+v callbacks=%d err=%v", active, callbacks, err)
		}
	})
}

func TestRefreshRejectsInvalidCatalogs(t *testing.T) {
	pack, _ := testPack(t, "alpha", "Alpha", "v1")
	valid, err := json.Marshal(Catalog{SchemaVersion: 1, Packs: []Pack{pack}})
	if err != nil {
		t.Fatal(err)
	}
	duplicate, err := json.Marshal(Catalog{SchemaVersion: 1, Packs: []Pack{pack, pack}})
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name string
		raw  []byte
	}{
		{name: "unknown field", raw: []byte(`{"schema_version":1,"packs":[],"url":"https://example.com"}`)},
		{name: "unsupported schema", raw: []byte(`{"schema_version":2,"packs":[]}`)},
		{name: "missing packs", raw: []byte(`{"schema_version":1}`)},
		{name: "arbitrary asset URL", raw: replaceJSONField(t, valid, "asset", "https://example.com/alpha.zip")},
		{name: "non-content-addressed asset", raw: replaceJSONField(t, valid, "asset", "alpha.zip")},
		{name: "invalid ID", raw: replaceJSONField(t, valid, "id", "../alpha")},
		{name: "invalid hash", raw: replaceJSONField(t, valid, "sha256", "nope")},
		{name: "duplicate ID and asset", raw: duplicate},
		{name: "inconsistent counts", raw: replaceJSONField(t, valid, "files", float64(3))},
		{name: "inconsistent name", raw: replaceJSONField(t, valid, "name", "Different API")},
		{name: "oversized", raw: bytes.Repeat([]byte(" "), maxCatalogBytes+1)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
				_, _ = writer.Write(test.raw)
			}))
			defer server.Close()
			manager := openTestManager(t, server.URL+"/catalog.json")
			if _, err := manager.Refresh(context.Background()); err == nil {
				t.Fatal("invalid catalog was accepted")
			}
		})
	}
}

func TestOpenUsesTimeoutOnlyForDefaultClient(t *testing.T) {
	manager, err := Open(t.TempDir(), Options{})
	if err != nil {
		t.Fatal(err)
	}
	if manager.client.Timeout != defaultHTTPTimeout {
		t.Fatalf("default HTTP timeout = %v", manager.client.Timeout)
	}
	injected := &http.Client{Timeout: 7 * time.Second}
	manager, err = Open(t.TempDir(), Options{HTTPClient: injected})
	if err != nil {
		t.Fatal(err)
	}
	if manager.client.Timeout != injected.Timeout {
		t.Fatalf("injected HTTP timeout changed: got %v want %v", manager.client.Timeout, injected.Timeout)
	}
}

func TestWriteJSONRejectsOversizedEncoding(t *testing.T) {
	manager, err := Open(t.TempDir(), Options{})
	if err != nil {
		t.Fatal(err)
	}
	if err := manager.writeJSON("catalog.json", strings.Repeat("x", maxCatalogBytes)); err == nil {
		t.Fatal("oversized encoded state was written")
	}
	if _, err := os.Stat(filepath.Join(manager.root, "catalog.json")); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("oversized state left a cache file: %v", err)
	}
}

func openTestManager(t *testing.T, catalogURL string) *Manager {
	t.Helper()
	manager, err := Open(filepath.Join(t.TempDir(), "packs"), Options{HTTPClient: http.DefaultClient, CatalogURL: catalogURL})
	if err != nil {
		t.Fatal(err)
	}
	return manager
}

func packServer(t *testing.T, catalog Catalog, assets map[string][]byte, requests *atomic.Int32) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path == "/packs/catalog.json" {
			_ = json.NewEncoder(writer).Encode(catalog)
			return
		}
		name := strings.TrimPrefix(request.URL.Path, "/packs/")
		raw, ok := assets[name]
		if !ok || name == request.URL.Path {
			http.NotFound(writer, request)
			return
		}
		if requests != nil {
			requests.Add(1)
		}
		_, _ = writer.Write(raw)
	}))
}

func testPack(t *testing.T, id, name, version string) (Pack, []byte) {
	t.Helper()
	description := name + " docs."
	manifest := fmt.Sprintf("---\nname: %s\nversion: %s\ndescription: %s\ncollections:\n  - examples\n---\n\nManifest.\n", name, version, description)
	page := "---\ntitle: Overview\n---\n\n# Overview\n"
	files := map[string]string{
		id + "/" + version + "/_index.md":   manifest,
		id + "/" + version + "/overview.md": page,
	}
	var buffer bytes.Buffer
	writer := zip.NewWriter(&buffer)
	for _, name := range []string{id + "/" + version + "/_index.md", id + "/" + version + "/overview.md"} {
		entry, err := writer.Create(name)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := entry.Write([]byte(files[name])); err != nil {
			t.Fatal(err)
		}
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	return packForArchive(id, name, version, buffer.Bytes(), 2, 1, int64(len(manifest)+len(page))), buffer.Bytes()
}

func traversalArchive(t *testing.T) []byte {
	t.Helper()
	manifest := "---\nname: Alpha\nversion: v1\ndescription: Alpha docs.\ncollections: [examples]\n---\n"
	var buffer bytes.Buffer
	writer := zip.NewWriter(&buffer)
	for name, content := range map[string]string{
		"alpha/v1/_index.md":       manifest,
		"alpha/v1/../../escape.md": "escape",
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
	return buffer.Bytes()
}

func packForArchive(id, name, version string, archive []byte, files, pages int, uncompressed int64) Pack {
	if uncompressed == 0 {
		reader, err := zip.NewReader(bytes.NewReader(archive), int64(len(archive)))
		if err == nil {
			for _, file := range reader.File {
				uncompressed += int64(file.UncompressedSize64)
			}
		}
	}
	hash := sha256.Sum256(archive)
	return Pack{
		ID: id, Name: name, Description: name + " docs.", Asset: id + "-" + hex.EncodeToString(hash[:]) + ".zip",
		SHA256: hex.EncodeToString(hash[:]), Bytes: int64(len(archive)), UncompressedBytes: uncompressed,
		Files: files, Pages: pages, Versions: []string{version}, Collections: []string{"examples"},
	}
}

func replaceJSONField(t *testing.T, raw []byte, field string, value any) []byte {
	t.Helper()
	var document map[string]any
	if err := json.Unmarshal(raw, &document); err != nil {
		t.Fatal(err)
	}
	packs := document["packs"].([]any)
	packs[0].(map[string]any)[field] = value
	result, err := json.Marshal(document)
	if err != nil {
		t.Fatal(err)
	}
	return result
}
