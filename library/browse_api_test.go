package library_test

import (
	"context"
	"testing"

	"github.com/sairaph/apis-mcp/library"
)

func TestBrowseAPIIsPubliclyNameable(t *testing.T) {
	var browse func(*library.Snapshot, context.Context, library.BrowseRequest) (library.BrowseResult, error)
	browse = (*library.Snapshot).Browse
	if browse == nil {
		t.Fatal("Snapshot.Browse method expression is nil")
	}
	if library.DefaultBrowseLimit < 1 || library.MaxBrowseLimit < library.DefaultBrowseLimit {
		t.Fatalf("invalid public browse limits: default=%d max=%d", library.DefaultBrowseLimit, library.MaxBrowseLimit)
	}
}
