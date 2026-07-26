package library_test

import (
	"bytes"
	"context"
	"errors"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/sairaph/apis-mcp/internal/budget"
	"github.com/sairaph/apis-mcp/library"
)

func TestManifestCatalogAndFilters(t *testing.T) {
	root, index := fixture(t)
	snapshot := open(t, root, index, 2_000, 4_000)
	defer snapshot.Close()

	collections, err := snapshot.Collections(context.Background(), library.CollectionsRequest{})
	if err != nil {
		t.Fatal(err)
	}
	collection := findCollection(t, collections.Collections, "developer_tools")
	if collection.Name != "Developer Tools" || collection.APICount != 1 {
		t.Fatalf("unexpected collection: %+v", collection)
	}

	listed, err := snapshot.List(context.Background(), library.ListRequest{
		Name: "widget", Version: "V2", Collection: "DEVELOPER_TOOLS",
	})
	if err != nil {
		t.Fatal(err)
	}
	if listed.Total != 1 || len(listed.APIs) != 1 {
		t.Fatalf("unexpected API list: %+v", listed)
	}
	api := listed.APIs[0]
	if api.Name != "Widget API" || len(api.Versions) != 2 {
		t.Fatalf("versions were not grouped by family: %+v", api)
	}
	if api.Versions[0].Version != "v2" || api.Versions[0].DocID != "widget-api-v2" {
		t.Fatalf("versions are not deterministically ordered: %+v", api.Versions)
	}
	if api.Versions[1].Pages != 5 {
		t.Fatalf("wrong v1 page count: %+v", api.Versions[1])
	}
}

func TestManifestValidationAndFamilyAgreement(t *testing.T) {
	t.Run("required fields", func(t *testing.T) {
		root := t.TempDir()
		write(t, root, "broken/_index.md", "---\nname: Broken\n---\n")
		_, err := library.Open(context.Background(), library.Options{
			UserRoot: root, IndexPath: filepath.Join(t.TempDir(), "index.sqlite"),
		})
		if err == nil || !strings.Contains(err.Error(), "requires name and version") {
			t.Fatalf("expected required manifest field error, got %v", err)
		}
	})

	t.Run("versions agree", func(t *testing.T) {
		root := t.TempDir()
		writeManifest(t, root, "one", "Same API", "v1", "first")
		writeManifest(t, root, "two", "Same API", "v2", "different")
		_, err := library.Open(context.Background(), library.Options{
			UserRoot: root, IndexPath: filepath.Join(t.TempDir(), "index.sqlite"),
		})
		if err == nil || !strings.Contains(err.Error(), "disagree on family metadata") {
			t.Fatalf("expected family consistency error, got %v", err)
		}
	})
}

func TestDuplicateDocIDReportsBothSources(t *testing.T) {
	root := t.TempDir()
	writeManifest(t, root, "duplicate", "Example API", "v1", "user copy")
	_, err := library.Open(context.Background(), library.Options{
		UserRoot: root, IndexPath: filepath.Join(t.TempDir(), "index.sqlite"),
	})
	if err == nil {
		t.Fatal("expected duplicate doc_id rejection")
	}
	var validation *library.ValidationError
	if !errors.As(err, &validation) {
		t.Fatalf("expected ValidationError, got %T: %v", err, err)
	}
	if len(validation.Locations) != 2 ||
		!containsSubstring(validation.Locations, "builtin:builtin/example/v1/_index.md") ||
		!containsSubstring(validation.Locations, filepath.Join(root, "duplicate", "_index.md")) {
		t.Fatalf("duplicate locations missing: %+v", validation.Locations)
	}
}

func TestPageIDsAndNavigation(t *testing.T) {
	root, index := fixture(t)
	snapshot := open(t, root, index, 2_000, 4_000)
	defer snapshot.Close()

	top, err := snapshot.Pages(context.Background(), library.PagesRequest{DocID: "widget-api-v1"})
	if err != nil {
		t.Fatal(err)
	}
	if top.Total != 3 || len(top.Paths) != 2 || len(top.Pages) != 1 {
		t.Fatalf("root navigation should contain paths before its direct page: %+v", top)
	}
	if top.Paths[0].Path != "guide" || top.Paths[0].NestedPages != 2 || top.Paths[1].Path != "reference" {
		t.Fatalf("unexpected root paths: %+v", top.Paths)
	}

	guide, err := snapshot.Pages(context.Background(), library.PagesRequest{DocID: "widget-api-v1", Path: "guide"})
	if err != nil {
		t.Fatal(err)
	}
	if len(guide.Paths) != 1 || guide.Paths[0].Path != "guide/deep" || len(guide.Pages) != 1 {
		t.Fatalf("navigation leaked siblings or descendants: %+v", guide)
	}
	firstID := guide.Pages[0].PageID
	deep, err := snapshot.Pages(context.Background(), library.PagesRequest{DocID: "widget-api-v1", Path: "guide/deep"})
	if err != nil {
		t.Fatal(err)
	}
	secondID := deep.Pages[0].PageID
	if firstID == secondID || !strings.HasPrefix(firstID, "getting-started-") || !strings.HasPrefix(secondID, "getting-started-") {
		t.Fatalf("default title collision was not deterministically suffixed: %q %q", firstID, secondID)
	}

	writePage(t, root, "widget/v1/guide/unrelated.md", "Unrelated", "", "# Unrelated\n")
	if err := library.Rebuild(context.Background(), library.Options{UserRoot: root, IndexPath: index}); err != nil {
		t.Fatal(err)
	}
	newSnapshot := open(t, root, index, 2_000, 4_000)
	defer newSnapshot.Close()
	again, err := newSnapshot.Pages(context.Background(), library.PagesRequest{DocID: "widget-api-v1", Path: "guide"})
	if err != nil {
		t.Fatal(err)
	}
	var ids []string
	for _, page := range again.Pages {
		ids = append(ids, page.PageID)
	}
	if !contains(ids, firstID) {
		t.Fatalf("adding a page changed an existing deterministic ID: %q not in %v", firstID, ids)
	}
}

func TestDuplicateExplicitPageIDIsRejected(t *testing.T) {
	root := t.TempDir()
	writeManifest(t, root, "api/v1", "Collision API", "v1", "Collisions.")
	writePage(t, root, "api/v1/one.md", "One", "page_id: same", "# One\n")
	writePage(t, root, "api/v1/two.md", "Two", "page_id: same", "# Two\n")
	_, err := library.Open(context.Background(), library.Options{
		UserRoot: root, IndexPath: filepath.Join(t.TempDir(), "index.sqlite"),
	})
	var validation *library.ValidationError
	if !errors.As(err, &validation) || len(validation.Locations) != 2 ||
		!strings.Contains(err.Error(), `duplicate explicit page_id "same"`) {
		t.Fatalf("expected both explicit collision locations, got %v", err)
	}
}

func TestSearchTermsPhrasesRankingMetadataAndPath(t *testing.T) {
	root, index := fixture(t)
	snapshot := open(t, root, index, 2_000, 4_000)
	defer snapshot.Close()
	ctx := context.Background()

	whole, err := snapshot.Search(ctx, library.SearchRequest{DocID: "widget-api-v1", Query: "post"})
	if err != nil {
		t.Fatal(err)
	}
	if whole.Total != 0 {
		t.Fatalf("whole-token search matched 'posting': %+v", whole.Hits)
	}

	phrase, err := snapshot.Search(ctx, library.SearchRequest{DocID: "widget-api-v1", Query: `"create widget"`})
	if err != nil {
		t.Fatal(err)
	}
	if phrase.Total != 2 {
		t.Fatalf("phrase search should find both matching body lines: %+v", phrase.Hits)
	}
	if !hasHit(phrase.Hits, "body", 1) || !hasHit(phrase.Hits, "body", 3) {
		t.Fatalf("missing phrase hits with exact source line: %+v", phrase.Hits)
	}

	ranked, err := snapshot.Search(ctx, library.SearchRequest{DocID: "widget-api-v1", Query: "widget later"})
	if err != nil {
		t.Fatal(err)
	}
	if len(ranked.Hits) < 2 || ranked.Hits[0].Line != 4 {
		t.Fatalf("all-clause coordination did not rank first: %+v", ranked.Hits)
	}

	metadata, err := snapshot.Search(ctx, library.SearchRequest{DocID: "widget-api-v1", Query: "createWidget"})
	if err != nil {
		t.Fatal(err)
	}
	if metadata.Total != 1 || metadata.Hits[0].Match != "operation_id" || metadata.Hits[0].Line != 0 {
		t.Fatalf("expected metadata-only operation hit: %+v", metadata.Hits)
	}

	filtered, err := snapshot.Search(ctx, library.SearchRequest{
		DocID: "widget-api-v1", Query: "getting", Path: "guide/deep",
	})
	if err != nil {
		t.Fatal(err)
	}
	if filtered.Total == 0 {
		t.Fatal("expected a search hit below the selected path")
	}
	for _, hit := range filtered.Hits {
		if hit.Path != "guide/deep" {
			t.Fatalf("path filter leaked hit: %+v", hit)
		}
	}
}

func TestReadRangesAndBudget(t *testing.T) {
	root, index := fixture(t)
	snapshot := open(t, root, index, 2_000, 4_000)
	defer snapshot.Close()

	read, err := snapshot.Read(context.Background(), library.ReadRequest{
		DocID: "widget-api-v1", PageID: "create-widget", Lines: []int{2, 4},
	})
	if err != nil {
		t.Fatal(err)
	}
	if read.Lines != [2]int{2, 4} || read.TotalLines != 7 || read.Markdown != "\nCreate widget now.\nA widget can be created later.\n" || read.Truncated {
		t.Fatalf("line range was not verbatim: %+v markdown=%q", read, read.Markdown)
	}

	small := open(t, root, filepath.Join(t.TempDir(), "small.sqlite"), 2_000, 3)
	defer small.Close()
	cropped, err := small.Read(context.Background(), library.ReadRequest{DocID: "widget-api-v1", PageID: "create-widget"})
	if err != nil {
		t.Fatal(err)
	}
	if !cropped.Truncated || cropped.Lines != [2]int{1, 1} || cropped.Markdown != "# Create Widget\n" {
		t.Fatalf("read budget did not stop on a deterministic line boundary: %+v %q", cropped, cropped.Markdown)
	}
	if _, err := snapshot.Read(context.Background(), library.ReadRequest{
		DocID: "widget-api-v1", PageID: "create-widget", Lines: []int{99, 100},
	}); !errors.Is(err, library.ErrInvalidArgument) {
		t.Fatalf("expected invalid range error, got %v", err)
	}
}

func TestReadExactTokenBoundaryUnicodeAndOversizedFirstLine(t *testing.T) {
	root := t.TempDir()
	writeManifest(t, root, "unicode/v1", "Unicode API", "v1", "Unicode token boundaries.")
	body := "hello world\nお誕生日おめでとう\nlast\n"
	writePage(t, root, "unicode/v1/page.md", "Unicode", "page_id: unicode", body)

	firstTwo := "hello world\nお誕生日おめでとう\n"
	exact, err := budget.Count(firstTwo)
	if err != nil {
		t.Fatal(err)
	}
	snapshot := open(t, root, filepath.Join(t.TempDir(), "exact.sqlite"), 2_000, exact)
	defer snapshot.Close()
	read, err := snapshot.Read(context.Background(), library.ReadRequest{DocID: "unicode-api-v1", PageID: "unicode"})
	if err != nil {
		t.Fatal(err)
	}
	if read.Lines != [2]int{1, 2} || read.Markdown != firstTwo || !read.Truncated {
		t.Fatalf("exact Unicode boundary = %+v, markdown %q", read, read.Markdown)
	}

	oversized := open(t, root, filepath.Join(t.TempDir(), "oversized.sqlite"), 2_000, 1)
	defer oversized.Close()
	read, err = oversized.Read(context.Background(), library.ReadRequest{DocID: "unicode-api-v1", PageID: "unicode"})
	if err != nil {
		t.Fatal(err)
	}
	if read.Lines != [2]int{1, 1} || read.Markdown != "hello world\n" || !read.Truncated {
		t.Fatalf("oversized first line = %+v, markdown %q", read, read.Markdown)
	}
}

func TestFailedRebuildPreservesPublishedIndexAndOpenSnapshot(t *testing.T) {
	root, index := fixture(t)
	old := open(t, root, index, 2_000, 4_000)
	defer old.Close()
	oldFingerprint := old.Fingerprint()

	writePage(t, root, "widget/v1/reference/create.md", "Create Widget", `page_id: create-widget
description: Create one widget.
api_endpoints: [/v1/widgets]
operation_ids: [createWidget]`, "# Changed\n\nChanged body.\n")
	if err := library.Rebuild(context.Background(), library.Options{UserRoot: root, IndexPath: index}); err != nil {
		t.Fatal(err)
	}
	current := open(t, root, index, 2_000, 4_000)
	defer current.Close()
	if current.Fingerprint() == oldFingerprint {
		t.Fatal("successful rebuild did not publish a new generation")
	}
	oldGeneration := generationFile(t, index, oldFingerprint)
	currentGeneration := generationFile(t, index, current.Fingerprint())
	if oldGeneration == currentGeneration {
		t.Fatal("rebuild reused a generation filename")
	}
	if _, err := os.Stat(index); !os.IsNotExist(err) {
		t.Fatalf("mutable base index should not be published: %v", err)
	}
	if _, err := os.Stat(oldGeneration); err != nil {
		t.Fatalf("open snapshot generation was removed during rebuild: %v", err)
	}
	oldRead, err := old.Read(context.Background(), library.ReadRequest{DocID: "widget-api-v1", PageID: "create-widget"})
	if err != nil || !strings.Contains(oldRead.Markdown, "Create widget now.") {
		t.Fatalf("old process snapshot changed after publication: %q, %v", oldRead.Markdown, err)
	}
	currentRead, err := current.Read(context.Background(), library.ReadRequest{DocID: "widget-api-v1", PageID: "create-widget"})
	if err != nil || !strings.Contains(currentRead.Markdown, "Changed body.") {
		t.Fatalf("new snapshot did not see rebuilt generation: %q, %v", currentRead.Markdown, err)
	}

	published, err := os.ReadFile(currentGeneration)
	if err != nil {
		t.Fatal(err)
	}
	writeManifest(t, root, "duplicate", "Example API", "v1", "duplicate")
	if err := library.Rebuild(context.Background(), library.Options{UserRoot: root, IndexPath: index}); err == nil {
		t.Fatal("invalid rebuild unexpectedly succeeded")
	}
	after, err := os.ReadFile(currentGeneration)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(published, after) {
		t.Fatal("failed rebuild modified the published SQLite generation")
	}
	afterFailure, err := current.Read(context.Background(), library.ReadRequest{DocID: "widget-api-v1", PageID: "create-widget"})
	if err != nil || !strings.Contains(afterFailure.Markdown, "Changed body.") {
		t.Fatalf("failed rebuild damaged active generation: %q, %v", afterFailure.Markdown, err)
	}
	if err := old.Close(); err != nil {
		t.Fatal(err)
	}
	if err := os.RemoveAll(filepath.Join(root, "duplicate")); err != nil {
		t.Fatal(err)
	}
	if err := library.Rebuild(context.Background(), library.Options{UserRoot: root, IndexPath: index}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(oldGeneration); !os.IsNotExist(err) {
		t.Fatalf("unused old generation was not cleaned up: %v", err)
	}
}

func TestTinyBudgetPaginationIsStable(t *testing.T) {
	root, index := fixture(t)
	snapshot := open(t, root, index, 1, 4_000)
	defer snapshot.Close()

	first, err := snapshot.Pages(context.Background(), library.PagesRequest{DocID: "widget-api-v1", Page: 1})
	if err != nil {
		t.Fatal(err)
	}
	second, err := snapshot.Pages(context.Background(), library.PagesRequest{DocID: "widget-api-v1", Page: 2})
	if err != nil {
		t.Fatal(err)
	}
	if first.TotalPages != 3 || len(first.Paths) != 1 || first.Paths[0].Path != "guide" ||
		len(second.Paths) != 1 || second.Paths[0].Path != "reference" {
		t.Fatalf("deterministic paths-first pagination failed: first=%+v second=%+v", first, second)
	}
}

func fixture(t *testing.T) (string, string) {
	t.Helper()
	root := t.TempDir()
	writeManifest(t, root, "widget/v1", "Widget API", "v1", "Manage widgets.")
	writeManifest(t, root, "widget/v2", "Widget API", "v2", "Manage widgets.")
	writePage(t, root, "widget/v1/home.md", "Home", "", "# Home\n\nWidget API home.\n")
	writePage(t, root, "widget/v1/guide/intro.md", "Getting Started", "", "# Getting Started\n\nRead the widget introduction.\n")
	writePage(t, root, "widget/v1/guide/deep/auth.md", "Getting Started", "", "# Deep Getting Started\n\nConfigure authentication.\n")
	writePage(t, root, "widget/v1/reference/create.md", "Create Widget", `page_id: create-widget
description: Create one widget.
api_endpoints: [/v1/widgets]
operation_ids: [createWidget]`, "# Create Widget\n\nCreate widget now.\nA widget can be created later.\nPosting is unrelated.\nAnother line.\nFinal line.\n")
	writePage(t, root, "widget/v1/reference/list.md", "List Widgets", "", "# List Widgets\n\nList all widgets.\n")
	writePage(t, root, "widget/v2/overview.md", "Version Two", "", "# Version Two\n")
	return root, filepath.Join(t.TempDir(), "index.sqlite")
}

func writeManifest(t *testing.T, root, relative, name, version, description string) {
	t.Helper()
	write(t, root, filepath.Join(relative, "_index.md"), "---\nname: "+name+"\nversion: "+version+"\ndescription: "+description+"\ncollections:\n  - developer_tools\n---\n")
}

func writePage(t *testing.T, root, relative, title, extra, body string) {
	t.Helper()
	front := "---\ntitle: " + title + "\n"
	if extra != "" {
		front += extra + "\n"
	}
	write(t, root, relative, front+"---\n\n"+body)
}

func write(t *testing.T, root, relative, content string) {
	t.Helper()
	name := filepath.Join(root, relative)
	if err := os.MkdirAll(filepath.Dir(name), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(name, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func open(t *testing.T, root, index string, listBudget, readBudget int) *library.Snapshot {
	t.Helper()
	snapshot, err := library.Open(context.Background(), library.Options{
		UserRoot: root, IndexPath: index, ListTokenBudget: listBudget, ReadTokenBudget: readBudget,
	})
	if err != nil {
		t.Fatal(err)
	}
	return snapshot
}

func generationFile(t *testing.T, index, fingerprint string) string {
	t.Helper()
	extension := filepath.Ext(index)
	name := strings.TrimSuffix(index, extension) + "-" + fingerprint + extension
	if _, err := os.Stat(name); err != nil {
		t.Fatalf("generation %s: %v", name, err)
	}
	return name
}

func findCollection(t *testing.T, values []library.Collection, id string) library.Collection {
	t.Helper()
	for _, value := range values {
		if value.Collection == id {
			return value
		}
	}
	t.Fatalf("collection %q not found in %+v", id, values)
	return library.Collection{}
}

func containsSubstring(values []string, wanted string) bool {
	for _, value := range values {
		if strings.Contains(value, wanted) {
			return true
		}
	}
	return false
}

func contains(values []string, wanted string) bool {
	sort.Strings(values)
	index := sort.SearchStrings(values, wanted)
	return index < len(values) && values[index] == wanted
}

func hasHit(hits []library.SearchHit, match string, line int) bool {
	for _, hit := range hits {
		if hit.Match == match && hit.Line == line {
			return true
		}
	}
	return false
}
