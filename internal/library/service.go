package library

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"
	"unicode"

	"github.com/sairaph/apis-mcp/internal/budget"
	"gopkg.in/yaml.v3"
)

type documentRow struct {
	docID       string
	familyID    string
	name        string
	version     string
	description string
	collections []string
	pageCount   int
}

type pageRow struct {
	pageID       string
	title        string
	description  string
	path         string
	body         string
	apiEndpoints []string
	operationIDs []string
}

func (s *Snapshot) Collections(ctx context.Context, request CollectionsRequest) (CollectionsResult, error) {
	page, err := validPage(request.Page)
	if err != nil {
		return CollectionsResult{}, err
	}
	rows, err := s.db.QueryContext(ctx, "SELECT family_id, collections FROM families ORDER BY lower(name), family_id")
	if err != nil {
		return CollectionsResult{}, err
	}
	defer rows.Close()
	counts := make(map[string]int)
	for rows.Next() {
		var familyID, encoded string
		if err := rows.Scan(&familyID, &encoded); err != nil {
			return CollectionsResult{}, err
		}
		collections, err := decodeStrings(encoded)
		if err != nil {
			return CollectionsResult{}, err
		}
		for _, collection := range collections {
			counts[collection]++
		}
	}
	if err := rows.Err(); err != nil {
		return CollectionsResult{}, err
	}
	collections := make([]Collection, 0, len(counts))
	for id, count := range counts {
		collections = append(collections, Collection{Collection: id, Name: displayName(id), APICount: count})
	}
	sort.Slice(collections, func(i, j int) bool { return collections[i].Collection < collections[j].Collection })
	window, pagination, err := paginate(collections, page, s.listTokenBudget, func(records []Collection) (string, error) {
		return yamlRepresentation(struct {
			Collections []Collection `yaml:"collections"`
		}{Collections: records})
	})
	if err != nil {
		return CollectionsResult{}, err
	}
	return CollectionsResult{Pagination: pagination, Collections: window}, nil
}

func (s *Snapshot) List(ctx context.Context, request ListRequest) (ListResult, error) {
	page, err := validPage(request.Page)
	if err != nil {
		return ListResult{}, err
	}
	documents, err := s.documents(ctx)
	if err != nil {
		return ListResult{}, err
	}
	type grouped struct {
		api      API
		familyID string
		matched  bool
	}
	byFamily := make(map[string]*grouped)
	for _, document := range documents {
		group := byFamily[document.familyID]
		if group == nil {
			group = &grouped{
				familyID: document.familyID,
				api: API{
					Name: document.name, Description: document.description,
					Collections: append([]string(nil), document.collections...),
				},
			}
			byFamily[document.familyID] = group
		}
		group.api.Versions = append(group.api.Versions, APIVersion{
			Version: document.version, DocID: document.docID, Pages: document.pageCount,
		})
		if request.Version == "" || strings.EqualFold(request.Version, document.version) {
			group.matched = true
		}
	}
	nameFilter := strings.ToLower(strings.TrimSpace(request.Name))
	collectionFilter := strings.TrimSpace(request.Collection)
	apis := make([]API, 0, len(byFamily))
	for _, group := range byFamily {
		if nameFilter != "" && !strings.Contains(strings.ToLower(group.api.Name), nameFilter) {
			continue
		}
		if request.Version != "" && !group.matched {
			continue
		}
		if collectionFilter != "" && !containsFold(group.api.Collections, collectionFilter) {
			continue
		}
		sort.Slice(group.api.Versions, func(i, j int) bool {
			if group.api.Versions[i].Version == group.api.Versions[j].Version {
				return group.api.Versions[i].DocID < group.api.Versions[j].DocID
			}
			return group.api.Versions[i].Version > group.api.Versions[j].Version
		})
		apis = append(apis, group.api)
	}
	sort.Slice(apis, func(i, j int) bool {
		left, right := strings.ToLower(apis[i].Name), strings.ToLower(apis[j].Name)
		if left == right {
			return apis[i].Name < apis[j].Name
		}
		return left < right
	})
	window, pagination, err := paginate(apis, page, s.listTokenBudget, func(records []API) (string, error) {
		return yamlRepresentation(struct {
			APIs []API `yaml:"apis"`
		}{APIs: records})
	})
	if err != nil {
		return ListResult{}, err
	}
	return ListResult{Pagination: pagination, APIs: window}, nil
}

func (s *Snapshot) Pages(ctx context.Context, request PagesRequest) (PagesResult, error) {
	pageNumber, err := validPage(request.Page)
	if err != nil {
		return PagesResult{}, err
	}
	if request.DocID == "" {
		return PagesResult{}, fmt.Errorf("%w: doc_id is required", ErrInvalidArgument)
	}
	selectedPath, err := canonicalPath(request.Path)
	if err != nil {
		return PagesResult{}, fmt.Errorf("%w: invalid path: %v", ErrInvalidArgument, err)
	}
	pages, err := s.pageMetadataRows(ctx, request.DocID)
	if err != nil {
		return PagesResult{}, err
	}
	if selectedPath != "" {
		found := false
		for _, page := range pages {
			if err := ctx.Err(); err != nil {
				return PagesResult{}, err
			}
			if page.path == selectedPath || strings.HasPrefix(page.path, selectedPath+"/") {
				found = true
				break
			}
		}
		if !found {
			return PagesResult{}, fmt.Errorf("%w: path %q in %s", ErrNotFound, selectedPath, request.DocID)
		}
	}

	type navigationRecord struct {
		path *Path
		page *Page
	}
	childCounts := make(map[string]int)
	var directPages []Page
	for _, candidate := range pages {
		if err := ctx.Err(); err != nil {
			return PagesResult{}, err
		}
		if candidate.path == selectedPath {
			directPages = append(directPages, Page{
				PageID: candidate.pageID, Title: candidate.title, Description: candidate.description,
			})
			continue
		}
		prefix := ""
		if selectedPath != "" {
			prefix = selectedPath + "/"
		}
		if !strings.HasPrefix(candidate.path, prefix) {
			continue
		}
		remainder := strings.TrimPrefix(candidate.path, prefix)
		if remainder == "" {
			continue
		}
		segment := strings.SplitN(remainder, "/", 2)[0]
		child := segment
		if selectedPath != "" {
			child = selectedPath + "/" + segment
		}
		childCounts[child]++
	}
	childPaths := make([]string, 0, len(childCounts))
	for child := range childCounts {
		if err := ctx.Err(); err != nil {
			return PagesResult{}, err
		}
		childPaths = append(childPaths, child)
	}
	sort.Strings(childPaths)
	sort.Slice(directPages, func(i, j int) bool {
		left, right := strings.ToLower(directPages[i].Title), strings.ToLower(directPages[j].Title)
		if left == right {
			return directPages[i].PageID < directPages[j].PageID
		}
		return left < right
	})
	records := make([]navigationRecord, 0, len(childPaths)+len(directPages))
	for _, child := range childPaths {
		item := Path{Path: child, NestedPages: childCounts[child]}
		records = append(records, navigationRecord{path: &item})
	}
	for i := range directPages {
		item := directPages[i]
		records = append(records, navigationRecord{page: &item})
	}
	window, pagination, err := paginate(records, pageNumber, s.listTokenBudget, func(records []navigationRecord) (string, error) {
		if err := ctx.Err(); err != nil {
			return "", err
		}
		value := struct {
			Paths []Path `yaml:"paths,omitempty"`
			Pages []Page `yaml:"pages,omitempty"`
		}{}
		for _, record := range records {
			if record.path != nil {
				value.Paths = append(value.Paths, *record.path)
			} else {
				value.Pages = append(value.Pages, *record.page)
			}
		}
		return yamlRepresentation(value)
	})
	if err != nil {
		return PagesResult{}, err
	}
	result := PagesResult{DocID: request.DocID, Path: selectedPath, Pagination: pagination}
	for _, record := range window {
		if record.path != nil {
			result.Paths = append(result.Paths, *record.path)
		} else {
			result.Pages = append(result.Pages, *record.page)
		}
	}
	return result, nil
}

// Browse returns one fixed-size hierarchy window. Its queries read only page
// navigation metadata; page bodies and API operation metadata are left on disk.
func (s *Snapshot) Browse(ctx context.Context, request BrowseRequest) (BrowseResult, error) {
	if request.DocID == "" {
		return BrowseResult{}, fmt.Errorf("%w: doc_id is required", ErrInvalidArgument)
	}
	selectedPath, err := canonicalPath(request.Path)
	if err != nil {
		return BrowseResult{}, fmt.Errorf("%w: invalid path: %v", ErrInvalidArgument, err)
	}
	if request.Offset < 0 {
		return BrowseResult{}, fmt.Errorf("%w: offset must not be negative", ErrInvalidArgument)
	}
	if request.Limit == 0 {
		request.Limit = DefaultBrowseLimit
	}
	if request.Limit < 1 || request.Limit > MaxBrowseLimit {
		return BrowseResult{}, fmt.Errorf("%w: limit must be between 1 and %d", ErrInvalidArgument, MaxBrowseLimit)
	}
	var exists int
	if err := s.db.QueryRowContext(ctx, "SELECT 1 FROM documents WHERE doc_id = ?", request.DocID).Scan(&exists); err != nil {
		return BrowseResult{}, notFound(err, "document "+request.DocID)
	}
	if selectedPath != "" {
		if err := s.db.QueryRowContext(ctx, `
SELECT EXISTS(
    SELECT 1 FROM pages
    WHERE doc_id = @doc_id
      AND (path = @path OR substr(path, 1, length(@path) + 1) = @path || '/')
)`, sql.Named("doc_id", request.DocID), sql.Named("path", selectedPath)).Scan(&exists); err != nil {
			return BrowseResult{}, err
		}
		if exists == 0 {
			return BrowseResult{}, fmt.Errorf("%w: path %q in %s", ErrNotFound, selectedPath, request.DocID)
		}
	}

	rows, err := s.db.QueryContext(ctx, `
WITH descendants AS (
    SELECT CASE
        WHEN @path = '' THEN path
        ELSE substr(path, length(@path) + 2)
    END AS remainder
    FROM pages
    WHERE doc_id = @doc_id
      AND path <> @path
      AND (@path = '' OR substr(path, 1, length(@path) + 1) = @path || '/')
), children AS (
    SELECT CASE
        WHEN @path = '' THEN segment
        ELSE @path || '/' || segment
    END AS child
    FROM (
        SELECT CASE
            WHEN instr(remainder, '/') = 0 THEN remainder
            ELSE substr(remainder, 1, instr(remainder, '/') - 1)
        END AS segment
        FROM descendants
        WHERE remainder <> ''
    )
), navigation AS (
    SELECT 0 AS kind, child AS nav_path, count(*) AS nested_pages,
           NULL AS page_id, NULL AS title, NULL AS description
    FROM children
    GROUP BY child
    UNION ALL
    SELECT 1 AS kind, '' AS nav_path, 0 AS nested_pages,
           page_id, title, description
    FROM pages
    WHERE doc_id = @doc_id AND path = @path
), totals AS (
    SELECT count(*) AS total FROM navigation
), window AS (
    SELECT kind, nav_path, nested_pages, page_id, title, description
    FROM navigation
    ORDER BY kind, nav_path, lower(title), page_id
    LIMIT @limit OFFSET @offset
)
SELECT window.kind, window.nav_path, window.nested_pages,
       window.page_id, window.title, window.description, totals.total
FROM totals
LEFT JOIN window ON 1
ORDER BY window.kind, window.nav_path, lower(window.title), window.page_id`,
		sql.Named("doc_id", request.DocID), sql.Named("path", selectedPath),
		sql.Named("limit", request.Limit), sql.Named("offset", request.Offset))
	if err != nil {
		return BrowseResult{}, err
	}
	defer rows.Close()
	result := BrowseResult{DocID: request.DocID, Path: selectedPath, Offset: request.Offset}
	for rows.Next() {
		if err := ctx.Err(); err != nil {
			return BrowseResult{}, err
		}
		var kind, nested sql.NullInt64
		var navigationPath, pageID, title, description sql.NullString
		if err := rows.Scan(&kind, &navigationPath, &nested, &pageID, &title, &description, &result.Total); err != nil {
			return BrowseResult{}, err
		}
		if !kind.Valid {
			continue
		}
		if kind.Int64 == 0 {
			result.Paths = append(result.Paths, Path{Path: navigationPath.String, NestedPages: int(nested.Int64)})
		} else {
			result.Pages = append(result.Pages, Page{PageID: pageID.String, Title: title.String, Description: description.String})
		}
	}
	if err := rows.Err(); err != nil {
		return BrowseResult{}, err
	}
	return result, nil
}

func (s *Snapshot) Read(ctx context.Context, request ReadRequest) (ReadResult, error) {
	if request.DocID == "" || request.PageID == "" {
		return ReadResult{}, fmt.Errorf("%w: doc_id and page_id are required", ErrInvalidArgument)
	}
	var page pageRow
	if err := s.db.QueryRowContext(ctx, `
SELECT page_id, title, description, path, body
FROM pages WHERE doc_id = ? AND page_id = ?`, request.DocID, request.PageID).Scan(
		&page.pageID, &page.title, &page.description, &page.path, &page.body,
	); err != nil {
		return ReadResult{}, notFound(err, fmt.Sprintf("page %s/%s", request.DocID, request.PageID))
	}
	lines := physicalLines(page.body)
	if len(lines) == 0 {
		return ReadResult{
			DocID: request.DocID, Title: page.title, PageID: page.pageID, Path: page.path,
			Lines: [2]int{0, 0}, TotalLines: 0,
		}, nil
	}
	start, targetEnd := 1, len(lines)
	if len(request.Lines) != 0 {
		if len(request.Lines) != 2 || request.Lines[0] < 1 || request.Lines[1] < request.Lines[0] {
			return ReadResult{}, fmt.Errorf("%w: lines must be an inclusive [start, end] range", ErrInvalidArgument)
		}
		start = request.Lines[0]
		if start > len(lines) {
			return ReadResult{}, fmt.Errorf("%w: line %d exceeds page length %d", ErrInvalidArgument, start, len(lines))
		}
		targetEnd = min(request.Lines[1], len(lines))
	}
	end := start - 1
	var content strings.Builder
	for i := start - 1; i < targetEnd; i++ {
		candidate := content.String() + lines[i].raw
		used, err := budget.Count(candidate)
		if err != nil {
			return ReadResult{}, fmt.Errorf("count read tokens: %w", err)
		}
		if end >= start && used > s.readTokenBudget {
			break
		}
		content.WriteString(lines[i].raw)
		end = i + 1
	}
	return ReadResult{
		DocID: request.DocID, Title: page.title, PageID: page.pageID, Path: page.path,
		Lines: [2]int{start, end}, TotalLines: len(lines), Truncated: end < targetEnd,
		Markdown: content.String(),
	}, nil
}

func (s *Snapshot) documents(ctx context.Context) ([]documentRow, error) {
	rows, err := s.db.QueryContext(ctx, `
SELECT doc_id, family_id, name, version, description, collections, page_count
FROM documents ORDER BY lower(name), version DESC, doc_id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var documents []documentRow
	for rows.Next() {
		var document documentRow
		var collections string
		if err := rows.Scan(
			&document.docID, &document.familyID, &document.name, &document.version,
			&document.description, &collections, &document.pageCount,
		); err != nil {
			return nil, err
		}
		document.collections, err = decodeStrings(collections)
		if err != nil {
			return nil, err
		}
		documents = append(documents, document)
	}
	return documents, rows.Err()
}

func (s *Snapshot) pageMetadataRows(ctx context.Context, docID string) ([]pageRow, error) {
	var exists int
	if err := s.db.QueryRowContext(ctx, "SELECT 1 FROM documents WHERE doc_id = ?", docID).Scan(&exists); err != nil {
		return nil, notFound(err, "document "+docID)
	}
	rows, err := s.db.QueryContext(ctx, `
SELECT page_id, title, description, path
FROM pages WHERE doc_id = ? ORDER BY page_id`, docID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var pages []pageRow
	for rows.Next() {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		var page pageRow
		if err := rows.Scan(
			&page.pageID, &page.title, &page.description, &page.path,
		); err != nil {
			return nil, err
		}
		pages = append(pages, page)
	}
	return pages, rows.Err()
}

func validPage(page int) (int, error) {
	if page == 0 {
		return 1, nil
	}
	if page < 1 {
		return 0, fmt.Errorf("%w: page must be one-based", ErrInvalidArgument)
	}
	return page, nil
}

func paginate[T any](records []T, requestedPage, tokenBudget int, render func([]T) (string, error)) ([]T, Pagination, error) {
	window, totalPages, err := budget.Paginate(records, requestedPage, tokenBudget, render)
	if totalPages == 0 {
		totalPages = 1
	}
	pagination := Pagination{Page: requestedPage, Total: len(records), TotalPages: totalPages}
	if err != nil {
		return nil, pagination, fmt.Errorf("paginate library records: %w", err)
	}
	return window, pagination, nil
}

func yamlRepresentation(value any) (string, error) {
	raw, err := yaml.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(raw), nil
}

type physicalLine struct {
	raw  string
	text string
}

func physicalLines(content string) []physicalLine {
	if content == "" {
		return nil
	}
	parts := strings.SplitAfter(content, "\n")
	if parts[len(parts)-1] == "" {
		parts = parts[:len(parts)-1]
	}
	lines := make([]physicalLine, 0, len(parts))
	for _, part := range parts {
		lines = append(lines, physicalLine{raw: part, text: strings.TrimSuffix(part, "\n")})
	}
	return lines
}

func decodeStrings(encoded string) ([]string, error) {
	var values []string
	if err := json.Unmarshal([]byte(encoded), &values); err != nil {
		return nil, fmt.Errorf("decode library metadata: %w", err)
	}
	return values, nil
}

func containsFold(values []string, wanted string) bool {
	for _, value := range values {
		if strings.EqualFold(value, wanted) {
			return true
		}
	}
	return false
}

func displayName(slug string) string {
	words := strings.Fields(strings.ReplaceAll(slug, "_", " "))
	for i, word := range words {
		runes := []rune(strings.ToLower(word))
		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
		}
		words[i] = string(runes)
	}
	return strings.Join(words, " ")
}

func notFound(err error, subject string) error {
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("%w: %s", ErrNotFound, subject)
	}
	return err
}
