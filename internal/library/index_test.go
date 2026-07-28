package library

import (
	"context"
	"path/filepath"
	"strings"
	"testing"
)

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
