// Package library provides immutable, transport-neutral documentation
// snapshots built from canonical Markdown.
package library

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	internal "github.com/sairaph/apis-mcp/internal/library"
)

//go:embed all:builtin
var builtinFiles embed.FS

var (
	ErrInvalidArgument = internal.ErrInvalidArgument
	ErrNotFound        = internal.ErrNotFound
)

type (
	Options            = internal.Options
	CollectionsRequest = internal.CollectionsRequest
	ListRequest        = internal.ListRequest
	PagesRequest       = internal.PagesRequest
	SearchRequest      = internal.SearchRequest
	ReadRequest        = internal.ReadRequest
	Pagination         = internal.Pagination
	Collection         = internal.Collection
	CollectionsResult  = internal.CollectionsResult
	APIVersion         = internal.APIVersion
	API                = internal.API
	ListResult         = internal.ListResult
	Path               = internal.Path
	Page               = internal.Page
	PagesResult        = internal.PagesResult
	SearchHit          = internal.SearchHit
	SearchResult       = internal.SearchResult
	ReadResult         = internal.ReadResult
	ValidationError    = internal.ValidationError
	Snapshot           = internal.Snapshot
)

// Open validates canonical sources, rebuilds a changed index, and pins the
// resulting generation until the returned snapshot is closed.
func Open(ctx context.Context, options Options) (*Snapshot, error) {
	sources, err := canonicalSources(options.UserRoot, options.ExcludeBuiltin)
	if err != nil {
		return nil, err
	}
	return internal.Open(ctx, options, sources)
}

// Rebuild validates and publishes a fingerprint-named SQLite generation. Open
// snapshots continue to observe their original generation.
func Rebuild(ctx context.Context, options Options) error {
	sources, err := canonicalSources(options.UserRoot, options.ExcludeBuiltin)
	if err != nil {
		return err
	}
	return internal.Rebuild(ctx, options, sources)
}

func canonicalSources(userRoot string, excludeBuiltin bool) ([]internal.Source, error) {
	sources := make([]internal.Source, 0, 2)
	if !excludeBuiltin {
		sources = append(sources, internal.Source{Name: "builtin", FS: builtinFiles})
	}
	if userRoot == "" {
		return sources, nil
	}
	absolute, err := filepath.Abs(userRoot)
	if err != nil {
		return nil, fmt.Errorf("resolve user library root: %w", err)
	}
	info, err := os.Stat(absolute)
	if os.IsNotExist(err) {
		return sources, nil
	}
	if err != nil {
		return nil, fmt.Errorf("inspect user library root: %w", err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("user library root is not a directory: %s", absolute)
	}
	sources = append(sources, internal.Source{Name: absolute, FS: fs.FS(os.DirFS(absolute))})
	return sources, nil
}
