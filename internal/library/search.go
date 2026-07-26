package library

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type queryClause struct {
	tokens []string
}

type textToken struct {
	text       string
	start, end int
}

type rankedHit struct {
	hit   SearchHit
	score float64
}

func (s *Snapshot) Search(ctx context.Context, request SearchRequest) (SearchResult, error) {
	pageNumber, err := validPage(request.Page)
	if err != nil {
		return SearchResult{}, err
	}
	if request.DocID == "" {
		return SearchResult{}, fmt.Errorf("%w: doc_id is required", ErrInvalidArgument)
	}
	clauses, err := parseQuery(request.Query)
	if err != nil {
		return SearchResult{}, err
	}
	selectedPath, err := canonicalPath(request.Path)
	if err != nil {
		return SearchResult{}, fmt.Errorf("%w: invalid path: %v", ErrInvalidArgument, err)
	}
	var exists int
	if err := s.db.QueryRowContext(ctx, "SELECT 1 FROM documents WHERE doc_id = ?", request.DocID).Scan(&exists); err != nil {
		return SearchResult{}, notFound(err, "document "+request.DocID)
	}

	rows, err := s.db.QueryContext(ctx, `
SELECT p.page_id, p.title, p.description, p.path, p.body, p.api_endpoints, p.operation_ids,
       bm25(page_search, 0.0, 12.0, 12.0, 6.0, 4.0, 1.0, 12.0, 12.0)
FROM page_search
JOIN pages p ON p.doc_id = page_search.doc_id AND p.page_id = page_search.page_id
WHERE page_search MATCH ? AND page_search.doc_id = ?`, ftsQuery(clauses), request.DocID)
	if err != nil {
		return SearchResult{}, fmt.Errorf("search library index: %w", err)
	}
	defer rows.Close()
	var hits []rankedHit
	for rows.Next() {
		var candidate pageRow
		var endpoints, operations string
		var pageRank float64
		if err := rows.Scan(
			&candidate.pageID, &candidate.title, &candidate.description, &candidate.path,
			&candidate.body, &endpoints, &operations, &pageRank,
		); err != nil {
			return SearchResult{}, err
		}
		if selectedPath != "" && candidate.path != selectedPath && !strings.HasPrefix(candidate.path, selectedPath+"/") {
			continue
		}
		candidate.apiEndpoints, err = decodeStrings(endpoints)
		if err != nil {
			return SearchResult{}, err
		}
		candidate.operationIDs, err = decodeStrings(operations)
		if err != nil {
			return SearchResult{}, err
		}
		hits = append(hits, hitsForPage(candidate, clauses, pageRank)...)
	}
	if err := rows.Err(); err != nil {
		return SearchResult{}, err
	}
	sort.Slice(hits, func(i, j int) bool {
		if hits[i].score != hits[j].score {
			return hits[i].score > hits[j].score
		}
		if hits[i].hit.PageID != hits[j].hit.PageID {
			return hits[i].hit.PageID < hits[j].hit.PageID
		}
		if hits[i].hit.Line != hits[j].hit.Line {
			return hits[i].hit.Line < hits[j].hit.Line
		}
		return hits[i].hit.Match < hits[j].hit.Match
	})
	plain := make([]SearchHit, len(hits))
	for i := range hits {
		plain[i] = hits[i].hit
	}
	window, pagination, err := paginate(plain, pageNumber, s.listTokenBudget, searchRepresentation)
	if err != nil {
		return SearchResult{}, err
	}
	return SearchResult{
		DocID: request.DocID, Query: request.Query, Path: selectedPath,
		Pagination: pagination, Hits: window,
	}, nil
}

func searchRepresentation(hits []SearchHit) (string, error) {
	front, err := yamlRepresentation(struct {
		Hits []SearchHit `yaml:"hits"`
	}{Hits: hits})
	if err != nil {
		return "", err
	}
	var representation strings.Builder
	representation.WriteString(front)
	for _, hit := range hits {
		line := ""
		if hit.Line > 0 {
			line = strconv.Itoa(hit.Line)
		}
		fmt.Fprintf(&representation, "| %s | %s | %s | %s | %s | %s |\n",
			tableValue(hit.PageID), line, tableValue(hit.Title), tableValue(hit.Path), tableValue(hit.Match), tableValue(hit.Snippet))
	}
	return representation.String(), nil
}

func tableValue(value string) string {
	value = strings.ReplaceAll(value, "|", "\\|")
	return strings.Join(strings.Fields(value), " ")
}

func hitsForPage(page pageRow, clauses []queryClause, pageRank float64) []rankedHit {
	var hits []rankedHit
	for index, line := range physicalLines(page.body) {
		text := strings.TrimSuffix(line.text, "\r")
		matched, all := clauseMatches(text, clauses)
		if matched == 0 {
			continue
		}
		hits = append(hits, rankedHit{
			hit: SearchHit{
				PageID: page.pageID, Line: index + 1, Title: page.title, Path: page.path,
				Match: "body", Snippet: snippet(text, clauses),
			},
			score: relevance(matched, all, 0, false, pageRank),
		})
	}
	bodyTokens := tokenize(page.body)
	metadata := []struct {
		name   string
		values []string
		weight float64
	}{
		{name: "title", values: []string{page.title}, weight: 50},
		{name: "page_id", values: []string{page.pageID}, weight: 50},
		{name: "path", values: []string{page.path}, weight: 30},
		{name: "api_endpoint", values: page.apiEndpoints, weight: 50},
		{name: "operation_id", values: page.operationIDs, weight: 50},
	}
	for _, field := range metadata {
		for _, value := range field.values {
			if value == "" || containsTokenSequence(bodyTokens, tokenTexts(tokenize(value))) {
				continue
			}
			matched, all := clauseMatches(value, clauses)
			if matched == 0 {
				continue
			}
			exact := exactClause(value, clauses)
			hits = append(hits, rankedHit{
				hit: SearchHit{
					PageID: page.pageID, Title: page.title, Path: page.path,
					Match: field.name, Snippet: snippet(value, clauses),
				},
				score: relevance(matched, all, field.weight, exact, pageRank),
			})
			break
		}
	}
	return hits
}

func relevance(matched int, all bool, fieldWeight float64, exact bool, pageRank float64) float64 {
	score := float64(matched)*100 + fieldWeight - pageRank
	if all {
		score += 100
	}
	if exact {
		score += 25
	}
	return score
}

func parseQuery(query string) ([]queryClause, error) {
	runes := []rune(strings.TrimSpace(query))
	if len(runes) == 0 {
		return nil, fmt.Errorf("%w: query is required", ErrInvalidArgument)
	}
	var clauses []queryClause
	for index := 0; index < len(runes); {
		for index < len(runes) && unicode.IsSpace(runes[index]) {
			index++
		}
		if index == len(runes) {
			break
		}
		if runes[index] == '"' {
			index++
			start := index
			for index < len(runes) && runes[index] != '"' {
				index++
			}
			if index == len(runes) {
				return nil, fmt.Errorf("%w: query has an unmatched quote", ErrInvalidArgument)
			}
			tokens := tokenTexts(tokenize(string(runes[start:index])))
			if len(tokens) == 0 {
				return nil, fmt.Errorf("%w: quoted phrases may not be empty", ErrInvalidArgument)
			}
			clauses = append(clauses, queryClause{tokens: tokens})
			index++
			continue
		}
		start := index
		for index < len(runes) && !unicode.IsSpace(runes[index]) && runes[index] != '"' {
			index++
		}
		for _, token := range tokenize(string(runes[start:index])) {
			clauses = append(clauses, queryClause{tokens: []string{token.text}})
		}
	}
	if len(clauses) == 0 {
		return nil, fmt.Errorf("%w: query must contain a searchable word", ErrInvalidArgument)
	}
	return clauses, nil
}

func ftsQuery(clauses []queryClause) string {
	parts := make([]string, len(clauses))
	for i, clause := range clauses {
		parts[i] = `"` + strings.ReplaceAll(strings.Join(clause.tokens, " "), `"`, `""`) + `"`
	}
	return strings.Join(parts, " OR ")
}

func clauseMatches(text string, clauses []queryClause) (int, bool) {
	tokens := tokenize(text)
	matched := 0
	for _, clause := range clauses {
		if containsTokenSequence(tokens, clause.tokens) {
			matched++
		}
	}
	return matched, matched == len(clauses)
}

func exactClause(text string, clauses []queryClause) bool {
	textTokens := tokenTexts(tokenize(text))
	for _, clause := range clauses {
		if equalStrings(textTokens, clause.tokens) {
			return true
		}
	}
	return false
}

func containsTokenSequence(tokens []textToken, wanted []string) bool {
	if len(wanted) == 0 || len(wanted) > len(tokens) {
		return false
	}
	for start := 0; start+len(wanted) <= len(tokens); start++ {
		matched := true
		for offset := range wanted {
			if tokens[start+offset].text != wanted[offset] {
				matched = false
				break
			}
		}
		if matched {
			return true
		}
	}
	return false
}

func tokenize(text string) []textToken {
	runes := []rune(strings.ToLower(text))
	var tokens []textToken
	start := -1
	for index, r := range runes {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			if start < 0 {
				start = index
			}
			continue
		}
		if start >= 0 {
			tokens = append(tokens, textToken{text: string(runes[start:index]), start: start, end: index})
			start = -1
		}
	}
	if start >= 0 {
		tokens = append(tokens, textToken{text: string(runes[start:]), start: start, end: len(runes)})
	}
	return tokens
}

func tokenTexts(tokens []textToken) []string {
	values := make([]string, len(tokens))
	for i := range tokens {
		values[i] = tokens[i].text
	}
	return values
}

func snippet(text string, clauses []queryClause) string {
	normalized := strings.Join(strings.Fields(text), " ")
	runes := []rune(normalized)
	tokens := tokenize(normalized)
	matchStart, matchEnd := -1, -1
	for _, clause := range clauses {
		for start := 0; start+len(clause.tokens) <= len(tokens); start++ {
			matched := true
			for offset := range clause.tokens {
				if tokens[start+offset].text != clause.tokens[offset] {
					matched = false
					break
				}
			}
			if matched {
				matchStart = tokens[start].start
				matchEnd = tokens[start+len(clause.tokens)-1].end
				break
			}
		}
		if matchStart >= 0 {
			break
		}
	}
	if matchStart < 0 {
		end := min(160, len(runes))
		value := string(runes[:end])
		if end < len(runes) {
			value += "..."
		}
		return value
	}
	start := max(0, matchStart-80)
	end := min(len(runes), matchEnd+80)
	value := string(runes[start:end])
	if start > 0 {
		value = "..." + value
	}
	if end < len(runes) {
		value += "..."
	}
	return value
}
