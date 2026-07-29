package library

import (
	"context"
	"io/fs"
	"path/filepath"
	"strings"
	"testing"
	"testing/fstest"
)

type closeTrackingFS struct {
	fs.FS
	closed bool
}

func (f *closeTrackingFS) Close() error {
	f.closed = true
	return nil
}

func TestPublishCatalogFlushesTemporaryIndex(t *testing.T) {
	fingerprint := strings.Repeat("a", 64)
	generation := filepath.Join(t.TempDir(), "library-"+fingerprint+".sqlite")
	if err := publishCatalog(context.Background(), generation, &catalog{fingerprint: fingerprint}); err != nil {
		t.Fatal(err)
	}
	got, err := indexFingerprint(context.Background(), generation)
	if err != nil {
		t.Fatal(err)
	}
	if got != fingerprint {
		t.Fatalf("fingerprint = %q, want %q", got, fingerprint)
	}
}

func TestLoadCatalogClosesArchiveSources(t *testing.T) {
	source := &closeTrackingFS{FS: fstest.MapFS{
		"example/v1/_index.md":   {Data: []byte("---\nname: Example\nversion: v1\n---\n")},
		"example/v1/overview.md": {Data: []byte("---\ntitle: Overview\n---\n")},
	}}
	if _, err := loadCatalog([]Source{{Name: "example.zip", FS: source, Official: true}}); err != nil {
		t.Fatal(err)
	}
	if !source.closed {
		t.Fatal("catalog materialization did not close archive source")
	}
}

func TestInvalidOptionsCloseArchiveSources(t *testing.T) {
	source := &closeTrackingFS{FS: fstest.MapFS{}}
	if _, err := Open(context.Background(), Options{}, []Source{{Name: "example.zip", FS: source, Official: true}}); err == nil {
		t.Fatal("Open accepted a missing index path")
	}
	if !source.closed {
		t.Fatal("invalid options left archive source open")
	}
}
