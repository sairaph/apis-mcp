package cache

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestCheckSpaceCleansRetainedEntriesBeforeRejectingReserve(t *testing.T) {
	root := t.TempDir()
	now := time.Date(2026, 7, 26, 12, 0, 0, 0, time.UTC)
	var checks int
	store, err := New(Config{
		Root: root, FreeDiskReserve: 100, Retention: time.Hour, Now: func() time.Time { return now },
		FreeBytes: func(string) (uint64, error) {
			checks++
			if _, err := os.Stat(filepath.Join(root, "orphan")); err == nil {
				return 100, nil
			}
			return 1_000, nil
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	entry, err := store.Begin("orphan", ".bin")
	if err != nil {
		t.Fatal(err)
	}
	body, err := entry.CreateBody()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := body.Write([]byte("expired")); err != nil {
		t.Fatal(err)
	}
	if err := entry.PublishBody(body); err != nil {
		t.Fatal(err)
	}
	completed := now.Add(-2 * time.Hour)
	if err := entry.SaveMetadata(Metadata{ID: entry.ID, State: "complete", CompletedAt: &completed}); err != nil {
		t.Fatal(err)
	}
	entry.Close()
	if err := store.CheckSpace(500); err != nil {
		t.Fatalf("space check after cleanup: %v", err)
	}
	if checks != 2 {
		t.Fatalf("free-space checks = %d, want 2", checks)
	}
	if _, err := os.Stat(entry.Directory); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("expired retained entry was not cleaned: %v", err)
	}
}

func TestPublishedBodySuppressesErrorArtifactAndIsNotOrphaned(t *testing.T) {
	store, err := New(Config{Root: t.TempDir(), Retention: time.Hour})
	if err != nil {
		t.Fatal(err)
	}
	entry, err := store.Begin("complete", ".txt")
	if err != nil {
		t.Fatal(err)
	}
	body, err := entry.CreateBody()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := body.Write([]byte("authoritative")); err != nil {
		t.Fatal(err)
	}
	if err := entry.PublishBody(body); err != nil {
		t.Fatal(err)
	}
	if err := entry.PublishError(ErrorRecord{Code: "late_failure"}); err != nil {
		t.Fatal(err)
	}
	entry.Close()
	if _, err := os.Stat(entry.ErrorPath); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("error artifact exists beside final body: %v", err)
	}
	result, err := store.Cleanup()
	if err != nil {
		t.Fatal(err)
	}
	if result.RemovedOrphans != 0 {
		t.Fatalf("published body treated as orphan: %#v", result)
	}
	if _, err := os.Stat(entry.BodyPath); err != nil {
		t.Fatalf("authoritative body was removed: %v", err)
	}
}

func TestPublishAndCleanup(t *testing.T) {
	now := time.Date(2026, 7, 26, 12, 0, 0, 0, time.UTC)
	store, err := New(Config{Root: t.TempDir(), MaxDecodedBytes: 32, Retention: time.Hour, Now: func() time.Time { return now }})
	if err != nil {
		t.Fatal(err)
	}
	entry, err := store.Begin("call", ".txt")
	if err != nil {
		t.Fatal(err)
	}
	body, err := entry.CreateBody()
	if err != nil {
		t.Fatal(err)
	}
	writer := store.Writer(body, false)
	if _, err := writer.Write([]byte("complete")); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(entry.BodyPath); !os.IsNotExist(err) {
		t.Fatalf("final body exists before publication: %v", err)
	}
	if err := entry.PublishBody(body); err != nil {
		t.Fatal(err)
	}
	completed := now
	if err := entry.SaveMetadata(Metadata{ID: "call", State: "complete", CompletedAt: &completed}); err != nil {
		t.Fatal(err)
	}
	entry.Close()
	if _, err := os.Stat(entry.TempPath); !os.IsNotExist(err) {
		t.Fatalf("temporary body remains: %v", err)
	}
	if info, err := os.Stat(entry.BodyPath); err != nil || runtime.GOOS != "windows" && info.Mode().Perm() != 0o600 {
		t.Fatalf("body mode: %v, %v", info, err)
	}

	now = now.Add(2 * time.Hour)
	result, err := store.Cleanup()
	if err != nil {
		t.Fatal(err)
	}
	if result.RemovedEntries != 1 {
		t.Fatalf("removed entries = %d", result.RemovedEntries)
	}
}

func TestDecodedLimit(t *testing.T) {
	store, err := New(Config{Root: t.TempDir(), MaxDecodedBytes: 3})
	if err != nil {
		t.Fatal(err)
	}
	entry, err := store.Begin("limited", ".bin")
	if err != nil {
		t.Fatal(err)
	}
	defer entry.Abort()
	body, err := entry.CreateBody()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := store.Writer(body, false).Write([]byte("four")); err != ErrSizeLimit {
		t.Fatalf("error = %v", err)
	}
	if err := body.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestCleanupRemovesOrphanedTemporaryEntry(t *testing.T) {
	store, err := New(Config{Root: t.TempDir(), Retention: time.Hour})
	if err != nil {
		t.Fatal(err)
	}
	entry, err := store.Begin("orphan", ".bin")
	if err != nil {
		t.Fatal(err)
	}
	body, err := entry.CreateBody()
	if err != nil {
		t.Fatal(err)
	}
	if err := body.Close(); err != nil {
		t.Fatal(err)
	}
	entry.Close()
	result, err := store.Cleanup()
	if err != nil {
		t.Fatal(err)
	}
	if result.RemovedOrphans != 1 {
		t.Fatalf("removed orphans = %d", result.RemovedOrphans)
	}
	if _, err := os.Stat(entry.Directory); !os.IsNotExist(err) {
		t.Fatalf("orphan remains: %v", err)
	}
}

func TestCleanupSkipsEntryHeldByOtherStore(t *testing.T) {
	root := t.TempDir()
	owner, err := New(Config{Root: root, Retention: time.Hour})
	if err != nil {
		t.Fatal(err)
	}
	cleaner, err := New(Config{Root: root, Retention: time.Hour})
	if err != nil {
		t.Fatal(err)
	}
	entry, err := owner.Begin("active", ".bin")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := cleaner.Begin("active", ".bin"); err == nil || !strings.Contains(err.Error(), "already active") {
		t.Fatalf("second begin error = %v", err)
	}
	result, err := cleaner.Cleanup()
	if err != nil {
		t.Fatal(err)
	}
	if result.RemovedEntries != 0 || result.RemovedOrphans != 0 {
		t.Fatalf("cleanup while active = %#v", result)
	}
	if _, err := os.Stat(entry.Directory); err != nil {
		t.Fatalf("active entry was removed: %v", err)
	}
	entry.Close()
	result, err = cleaner.Cleanup()
	if err != nil {
		t.Fatal(err)
	}
	if result.RemovedOrphans != 1 {
		t.Fatalf("removed orphans = %d", result.RemovedOrphans)
	}
	if _, err := os.Stat(filepath.Join(root, ".locks", "active.lock")); err != nil {
		t.Fatalf("persistent lock file: %v", err)
	}
}

func TestCleanupRecoversDownloadAfterProcessExit(t *testing.T) {
	if os.Getenv("APIS_MCP_CACHE_LEASE_HELPER") == "1" {
		store, err := New(Config{Root: os.Getenv("APIS_MCP_CACHE_ROOT"), Retention: time.Hour})
		if err != nil {
			t.Fatal(err)
		}
		entry, err := store.Begin("download", ".bin")
		if err != nil {
			t.Fatal(err)
		}
		body, err := entry.CreateBody()
		if err != nil {
			t.Fatal(err)
		}
		if _, err := body.Write([]byte("partial")); err != nil {
			t.Fatal(err)
		}
		fmt.Fprintln(os.Stdout, "ready")
		time.Sleep(time.Minute)
		runtime.KeepAlive(body)
		runtime.KeepAlive(entry)
		return
	}

	root := t.TempDir()
	store, err := New(Config{Root: root, Retention: time.Hour})
	if err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestCleanupRecoversDownloadAfterProcessExit$")
	cmd.Env = append(os.Environ(), "APIS_MCP_CACHE_LEASE_HELPER=1", "APIS_MCP_CACHE_ROOT="+root)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}
	waited := false
	defer func() {
		if !waited {
			_ = cmd.Process.Kill()
			_ = cmd.Wait()
		}
	}()
	line, err := bufio.NewReader(stdout).ReadString('\n')
	if err != nil || strings.TrimSpace(line) != "ready" {
		t.Fatalf("helper readiness = %q, %v", line, err)
	}
	result, err := store.Cleanup()
	if err != nil {
		t.Fatal(err)
	}
	if result.RemovedOrphans != 0 {
		t.Fatalf("removed live download = %d", result.RemovedOrphans)
	}
	if _, err := os.Stat(filepath.Join(root, "download")); err != nil {
		t.Fatalf("live download was removed: %v", err)
	}
	if err := cmd.Process.Kill(); err != nil {
		t.Fatal(err)
	}
	_ = cmd.Wait()
	waited = true
	result, err = store.Cleanup()
	if err != nil {
		t.Fatal(err)
	}
	if result.RemovedOrphans != 1 {
		t.Fatalf("removed crashed download orphans = %d", result.RemovedOrphans)
	}
}
