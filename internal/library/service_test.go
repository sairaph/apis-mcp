package library

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
)

func TestHierarchyQueriesUseOnlyMetadataColumns(t *testing.T) {
	snapshot := metadataOnlySnapshot(t, 250)

	first, err := snapshot.Browse(context.Background(), BrowseRequest{
		DocID: "cloudflare-v1", Path: "schemas", Limit: 75,
	})
	if err != nil {
		t.Fatal(err)
	}
	if first.Offset != 0 || first.Total != 250 || len(first.Pages) != 75 || len(first.Paths) != 0 {
		t.Fatalf("unexpected first hierarchy window: %+v", first)
	}
	if first.Pages[0].PageID != "schema-0000" || first.Pages[74].PageID != "schema-0074" {
		t.Fatalf("unexpected first window bounds: %q through %q", first.Pages[0].PageID, first.Pages[74].PageID)
	}

	last, err := snapshot.Browse(context.Background(), BrowseRequest{
		DocID: "cloudflare-v1", Path: "schemas", Offset: 225, Limit: 75,
	})
	if err != nil {
		t.Fatal(err)
	}
	if last.Offset != 225 || last.Total != 250 || len(last.Pages) != 25 || last.Pages[0].PageID != "schema-0225" {
		t.Fatalf("unexpected last hierarchy window: %+v", last)
	}

	// The fixture omits body, api_endpoints, and operation_ids entirely. Pages
	// succeeding against it ensures the MCP hierarchy query is metadata-only too.
	pages, err := snapshot.Pages(context.Background(), PagesRequest{DocID: "cloudflare-v1", Path: "schemas"})
	if err != nil {
		t.Fatal(err)
	}
	if pages.Total != 250 || len(pages.Pages) == 0 {
		t.Fatalf("unexpected token-paginated hierarchy: %+v", pages)
	}
}

func TestBrowseValidationAndCancellation(t *testing.T) {
	snapshot := metadataOnlySnapshot(t, 10)
	for _, request := range []BrowseRequest{
		{},
		{DocID: "cloudflare-v1", Offset: -1},
		{DocID: "cloudflare-v1", Limit: MaxBrowseLimit + 1},
	} {
		if _, err := snapshot.Browse(context.Background(), request); err == nil {
			t.Fatalf("Browse accepted invalid request: %+v", request)
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := snapshot.Browse(ctx, BrowseRequest{DocID: "cloudflare-v1"}); err == nil {
		t.Fatal("Browse ignored a cancelled context")
	}
}

func metadataOnlySnapshot(t *testing.T, pageCount int) *Snapshot {
	t.Helper()
	db, err := sql.Open("sqlite", "file:"+t.Name()+"?mode=memory&cache=shared")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = db.Close() })
	if _, err := db.Exec(`
CREATE TABLE documents (doc_id TEXT PRIMARY KEY);
CREATE TABLE pages (
    doc_id TEXT NOT NULL,
    page_id TEXT NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    path TEXT NOT NULL,
    PRIMARY KEY (doc_id, page_id)
);
CREATE INDEX pages_navigation ON pages(doc_id, path, title, page_id);
INSERT INTO documents(doc_id) VALUES ('cloudflare-v1');`); err != nil {
		t.Fatal(err)
	}
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	statement, err := tx.Prepare("INSERT INTO pages(doc_id, page_id, title, description, path) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		t.Fatal(err)
	}
	for index := 0; index < pageCount; index++ {
		id := fmt.Sprintf("schema-%04d", index)
		if _, err := statement.Exec("cloudflare-v1", id, fmt.Sprintf("Schema %04d", index), "Schema metadata", "schemas"); err != nil {
			t.Fatal(err)
		}
	}
	if err := statement.Close(); err != nil {
		t.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
	return &Snapshot{db: db, listTokenBudget: defaultListTokenBudget, readTokenBudget: defaultReadTokenBudget}
}
