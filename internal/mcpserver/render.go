package mcpserver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/sairaph/apis-mcp/internal/app"
	"github.com/sairaph/apis-mcp/internal/httpcall"
	"github.com/sairaph/apis-mcp/internal/render"
	"github.com/sairaph/apis-mcp/internal/sessions"
	"github.com/sairaph/apis-mcp/library"
)

func successResult(front any, body string) *mcp.CallToolResult {
	return documentResult(render.Document{Front: front, Body: body})
}

func errorResult(tool string, err error) *mcp.CallToolResult {
	code, message, hint, fields := errorDetails(tool, err)
	front := struct {
		Error callError `yaml:"error"`
	}{}
	front.Error = callError{Code: code, Message: message, Hint: hint, Fields: fields}
	result := documentResult(render.Document{Front: front, Body: errorBody(message, hint), IsError: true})
	result.IsError = true
	return result
}

func callErrorResult(result httpcall.Result, err error) *mcp.CallToolResult {
	code, message, hint, fields := errorDetails("apis_call", err)
	call := makeCallFront(result)
	front := struct {
		Error     callError      `yaml:"error"`
		Request   *callRequest   `yaml:"request,omitempty"`
		Response  *callResponse  `yaml:"response,omitempty"`
		Cache     *callCache     `yaml:"cache,omitempty"`
		Attempts  []callAttempt  `yaml:"attempts,omitempty"`
		Redirects []callRedirect `yaml:"redirects,omitempty"`
	}{
		Error:    callError{Code: code, Message: message, Hint: hint, Fields: fields},
		Attempts: call.Attempts, Redirects: call.Redirects,
	}
	if result.Request.ID != "" {
		front.Request = &call.Request
	}
	if result.Response.State != "" {
		front.Response = &call.Response
	}
	if result.Cache.Directory != "" {
		front.Cache = &call.Cache
	}
	rendered := documentResult(render.Document{Front: front, Body: errorBody(message, hint), IsError: true})
	rendered.IsError = true
	return rendered
}

func errorBody(message, hint string) string {
	body := "## Error\n\n" + message
	if hint != "" {
		body += "\n\n" + hint
	}
	return body
}

func documentResult(document render.Document) *mcp.CallToolResult {
	text, err := document.String()
	if err != nil {
		text = "---\nerror:\n  code: render_error\n  message: failed to render tool result\n---\n\n## Error\n\nFailed to render the tool result.\n"
		document.IsError = true
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: text}},
		IsError: document.IsError,
	}
}

func errorDetails(tool string, err error) (string, string, string, map[string]any) {
	var applicationError *app.Error
	if errors.As(err, &applicationError) {
		return applicationError.Code, applicationError.Message, applicationError.Hint, applicationError.Fields
	}
	var callError *httpcall.Error
	if errors.As(err, &callError) {
		return callError.Code, callError.Message, callError.Hint, callError.Fields
	}
	switch {
	case errors.Is(err, library.ErrInvalidArgument):
		return "invalid_input", err.Error(), "Correct the arguments and retry " + tool + ".", nil
	case errors.Is(err, library.ErrNotFound):
		return "not_found", err.Error(), "Use apis_list and apis_pages to select existing documentation identifiers.", nil
	case errors.Is(err, sessions.ErrNotFound):
		return "session_not_found", err.Error(), "Call apis_sessions without an id to list available sessions.", nil
	case errors.Is(err, sessions.ErrBusy):
		return "session_busy", err.Error(), "Wait for the active apis_call to finish, then retry.", nil
	case errors.Is(err, context.Canceled), errors.Is(err, context.DeadlineExceeded):
		return "cancelled", err.Error(), "Retry the tool call when it can run to completion.", nil
	default:
		return "internal_error", err.Error(), "Retry the tool call; if it continues to fail, inspect local diagnostics.", nil
	}
}

func collectionsBody(result library.CollectionsResult, input app.CollectionsInput) string {
	var lines []string
	if len(result.Collections) == 0 {
		lines = append(lines, "No documentation collections found.")
	} else {
		lines = append(lines, "Use a returned `collection` with `apis_list` to filter API documentation sets.")
	}
	if result.Page < result.TotalPages {
		input.Page = result.Page + 1
		lines = append(lines, continuation("apis_collections", input))
	}
	return strings.Join(lines, "\n\n")
}

func listBody(result library.ListResult, input app.ListInput) string {
	var lines []string
	if len(result.APIs) == 0 {
		lines = append(lines, "No API documentation sets match these filters.")
	} else {
		lines = append(lines, "Use a returned `doc_id` with `apis_pages` or `apis_search`.")
	}
	if result.Page < result.TotalPages {
		input.Page = result.Page + 1
		lines = append(lines, continuation("apis_list", input))
	}
	return strings.Join(lines, "\n\n")
}

func pagesBody(result library.PagesResult, input app.PagesInput) string {
	var lines []string
	if len(result.Paths) == 0 && len(result.Pages) == 0 {
		lines = append(lines, "No documentation pages or child paths were found here.")
	}
	if len(result.Paths) > 0 {
		next := app.PagesInput{DocID: result.DocID, Path: result.Paths[0].Path}
		lines = append(lines, "Browse a returned child path with `"+toolCall("apis_pages", next)+"`.")
	}
	if len(result.Pages) > 0 {
		read := app.ReadInput{DocID: result.DocID, PageID: result.Pages[0].PageID}
		lines = append(lines, "Read a returned page with `"+toolCall("apis_read", read)+"`.")
	}
	if result.Page < result.TotalPages {
		input.Page = result.Page + 1
		lines = append(lines, continuation("apis_pages", input))
	}
	return strings.Join(lines, "\n\n")
}

func searchBody(result library.SearchResult, input app.SearchInput) string {
	var body strings.Builder
	shown := len(result.Hits)
	switch {
	case result.Total == 0:
		fmt.Fprintf(&body, "No documentation lines match %s.", strconv.Quote(result.Query))
	case shown == 0:
		fmt.Fprintf(&body, "No results on page %d (%d %s; %d %s).", result.Page, result.Total, plural(result.Total, "match", "matches"), result.TotalPages, plural(result.TotalPages, "page", "pages"))
	case result.TotalPages <= 1:
		fmt.Fprintf(&body, "Showing %d matching lines for %s (keyword, ranked by relevance).", result.Total, strconv.Quote(result.Query))
	default:
		fmt.Fprintf(&body, "Showing %d out of %d matching lines for %s (keyword, ranked by relevance).", shown, result.Total, strconv.Quote(result.Query))
	}
	if shown > 0 {
		body.WriteString("\n\n| page_id | line | title | path | match | snippet |\n|---|---:|---|---|---|---|\n")
		for _, hit := range result.Hits {
			line := ""
			if hit.Line > 0 {
				line = strconv.Itoa(hit.Line)
			}
			fmt.Fprintf(&body, "| %s | %s | %s | %s | %s | %s |\n", tableCell(hit.PageID), line, tableCell(hit.Title), tableCell(hit.Path), tableCell(hit.Match), tableCell(hit.Snippet))
		}
	}
	for _, hit := range result.Hits {
		if hit.Line > 0 {
			read := app.ReadInput{DocID: result.DocID, PageID: hit.PageID, Lines: []int{max(1, hit.Line-2), hit.Line + 2}}
			body.WriteString("\nRead around a hit with " + toolCall("apis_read", read) + ".")
			break
		}
	}
	for _, hit := range result.Hits {
		if hit.Line == 0 {
			read := app.ReadInput{DocID: result.DocID, PageID: hit.PageID}
			body.WriteString("\nRead a metadata-only hit with " + toolCall("apis_read", read) + ".")
			break
		}
	}
	if result.Page < result.TotalPages {
		input.Page = result.Page + 1
		body.WriteString("\n\n" + continuation("apis_search", input))
	}
	return body.String()
}

type readMetadata struct {
	DocID      string `yaml:"doc_id"`
	Title      string `yaml:"title"`
	PageID     string `yaml:"page_id"`
	Path       string `yaml:"path,omitempty"`
	Lines      [2]int `yaml:"lines,flow"`
	TotalLines int    `yaml:"total_lines"`
	Truncated  bool   `yaml:"truncated"`
}

func readFront(result library.ReadResult) readMetadata {
	return readMetadata{
		DocID: result.DocID, Title: result.Title, PageID: result.PageID, Path: result.Path,
		Lines: result.Lines, TotalLines: result.TotalLines, Truncated: result.Truncated,
	}
}

func readBody(result library.ReadResult) string {
	body := render.Fence(result.Markdown, "markdown")
	if result.Truncated {
		width := max(1, result.Lines[1]-result.Lines[0]+1)
		start := result.Lines[1] + 1
		end := min(result.TotalLines, start+width-1)
		next := app.ReadInput{DocID: result.DocID, PageID: result.PageID, Lines: []int{start, end}}
		body += "\n\nOutput was cropped. Continue with `" + toolCall("apis_read", next) + "`."
	}
	return body
}

func continuation(name string, input any) string {
	return "Continue with `" + toolCall(name, input) + "`."
}

func toolCall(name string, input any) string {
	raw, err := json.Marshal(input)
	if err != nil {
		return name + "({})"
	}
	return name + "(" + string(raw) + ")"
}

func tableCell(value string) string {
	value = strings.ReplaceAll(value, "|", "\\|")
	return strings.Join(strings.Fields(value), " ")
}

func plural(count int, singular, plural string) string {
	if count == 1 {
		return singular
	}
	return plural
}

type callFront struct {
	Request   callRequest    `yaml:"request"`
	Response  callResponse   `yaml:"response"`
	Cache     callCache      `yaml:"cache"`
	Preview   *callPreview   `yaml:"preview,omitempty"`
	Selection *callSelection `yaml:"selection,omitempty"`
	Attempts  []callAttempt  `yaml:"attempts,omitempty"`
	Redirects []callRedirect `yaml:"redirects,omitempty"`
}

type callError struct {
	Code    string         `yaml:"code"`
	Message string         `yaml:"message"`
	Hint    string         `yaml:"hint,omitempty"`
	Fields  map[string]any `yaml:"fields,omitempty"`
}

type callRequest struct {
	ID               string            `yaml:"id"`
	SessionID        string            `yaml:"session_id"`
	Method           string            `yaml:"method"`
	Endpoint         string            `yaml:"endpoint"`
	AutomaticHeaders map[string]string `yaml:"automatic_headers,omitempty"`
}

type callResponse struct {
	State              string              `yaml:"state"`
	Status             int                 `yaml:"status"`
	StatusText         string              `yaml:"status_text,omitempty"`
	FinalURL           string              `yaml:"final_url,omitempty"`
	ContentType        string              `yaml:"content_type,omitempty"`
	ContentEncoding    string              `yaml:"content_encoding,omitempty"`
	Headers            map[string][]string `yaml:"headers,omitempty"`
	Decoded            bool                `yaml:"decoded"`
	WireBytes          int64               `yaml:"wire_bytes"`
	BytesReceived      int64               `yaml:"bytes_received"`
	DeclaredBytes      int64               `yaml:"declared_bytes,omitempty"`
	BytesPerSecond     float64             `yaml:"bytes_per_second,omitempty"`
	EstimatedRemaining string              `yaml:"estimated_remaining,omitempty"`
	DurationMS         int64               `yaml:"duration_ms"`
}

type callCache struct {
	Directory    string     `yaml:"directory,omitempty"`
	BodyPath     string     `yaml:"body_path,omitempty"`
	TempPath     string     `yaml:"temp_path,omitempty"`
	FinalPath    string     `yaml:"final_path,omitempty"`
	HeadersPath  string     `yaml:"headers_path,omitempty"`
	MetadataPath string     `yaml:"metadata_path,omitempty"`
	ErrorPath    string     `yaml:"error_path,omitempty"`
	CompletedAt  *time.Time `yaml:"completed_at,omitempty"`
}

type callPreview struct {
	Kind              string `yaml:"kind"`
	Language          string `yaml:"language,omitempty"`
	Truncated         bool   `yaml:"truncated"`
	ApproximateTokens int    `yaml:"approximate_tokens,omitempty"`
}

type callSelection struct {
	JSONPath string `yaml:"json_path"`
	State    string `yaml:"state,omitempty"`
	Matched  bool   `yaml:"matched"`
	Error    string `yaml:"selection_error,omitempty"`
}

type callAttempt struct {
	Number      int    `yaml:"number"`
	Status      int    `yaml:"status,omitempty"`
	Error       string `yaml:"error,omitempty"`
	RetryReason string `yaml:"retry_reason,omitempty"`
	RetryDelay  string `yaml:"retry_delay,omitempty"`
}

type callRedirect struct {
	FromURL      string `yaml:"from_url"`
	ToURL        string `yaml:"to_url"`
	Status       int    `yaml:"status"`
	MethodBefore string `yaml:"method_before"`
	MethodAfter  string `yaml:"method_after"`
	BodyRetained bool   `yaml:"body_retained"`
}

func makeCallFront(result httpcall.Result) callFront {
	front := callFront{
		Request: callRequest{
			ID: result.Request.ID, SessionID: result.Request.SessionID, Method: result.Request.Method,
			Endpoint: result.Request.Endpoint, AutomaticHeaders: result.Request.AutomaticHeaders,
		},
		Response: callResponse{
			State: result.Response.State, Status: result.Response.Status, StatusText: result.Response.StatusText,
			FinalURL: result.Response.FinalURL, ContentType: result.Response.ContentType,
			ContentEncoding: result.Response.ContentEncoding, Headers: result.Response.Headers,
			Decoded: result.Response.Decoded, WireBytes: result.Response.WireBytes,
			BytesReceived: result.Response.DecodedBytes, DeclaredBytes: result.Response.DeclaredBytes,
			BytesPerSecond: result.Response.BytesPerSecond, DurationMS: result.Response.Duration.Milliseconds(),
		},
		Cache: callCache{
			Directory: result.Cache.Directory, BodyPath: result.Cache.BodyPath, TempPath: result.Cache.TempPath,
			FinalPath: result.Cache.FinalPath, HeadersPath: result.Cache.HeadersPath,
			MetadataPath: result.Cache.MetadataPath, ErrorPath: result.Cache.ErrorPath,
			CompletedAt: result.Response.CompletedAt,
		},
	}
	if result.Response.EstimatedRemaining != nil {
		front.Response.EstimatedRemaining = result.Response.EstimatedRemaining.Round(time.Second).String()
	}
	if result.Preview != nil {
		front.Preview = &callPreview{
			Kind: result.Preview.Kind, Language: result.Preview.Language,
			Truncated: result.Preview.Truncated, ApproximateTokens: result.Preview.ApproximateTokens,
		}
	}
	if result.Selection != nil {
		front.Selection = &callSelection{
			JSONPath: result.Selection.JSONPath, State: result.Selection.State,
			Matched: result.Selection.Matched, Error: result.Selection.Error,
		}
	}
	for _, attempt := range result.Attempts {
		item := callAttempt{
			Number: attempt.Number, Status: attempt.Status, Error: attempt.Error,
			RetryReason: attempt.RetryReason,
		}
		if attempt.RetryDelay > 0 {
			item.RetryDelay = attempt.RetryDelay.String()
		}
		front.Attempts = append(front.Attempts, item)
	}
	for _, redirect := range result.Redirects {
		front.Redirects = append(front.Redirects, callRedirect{
			FromURL: redirect.FromURL, ToURL: redirect.ToURL, Status: redirect.Status,
			MethodBefore: redirect.MethodBefore, MethodAfter: redirect.MethodAfter,
			BodyRetained: redirect.BodyRetained,
		})
	}
	return front
}

func callBody(result httpcall.Result) string {
	var parts []string
	if result.Preview != nil {
		parts = append(parts, render.Fence(result.Preview.Content, result.Preview.Language))
		if result.Preview.Truncated {
			parts = append(parts, "Preview was truncated. Read the complete response at `"+result.Cache.BodyPath+"`.")
		}
	} else if result.Response.State == "complete" && result.Cache.BodyPath != "" {
		parts = append(parts, "No text preview is available. Read the complete response at `"+result.Cache.BodyPath+"`.")
	}
	if len(result.Request.AutomaticHeaders) > 0 {
		parts = append(parts, "The request added the headers shown in `request.automatic_headers` automatically.")
	}
	if result.Selection != nil && result.Selection.Error != "" {
		parts = append(parts, "JSONPath selection was not applied: "+result.Selection.Error+". The complete response remains at `"+firstNonempty(result.Cache.BodyPath, result.Cache.FinalPath)+"`.")
	}
	if result.Response.State == "downloading" {
		wait := "Completion time cannot be estimated."
		if result.Response.EstimatedRemaining != nil {
			wait = "Wait approximately " + result.Response.EstimatedRemaining.Round(time.Second).String() + "."
		}
		parts = append(parts,
			"The response is still downloading. `"+result.Cache.TempPath+"` is partial and must not be treated as complete. "+wait+
				" Then check for `"+result.Cache.FinalPath+"`; if it is absent, check `"+result.Cache.ErrorPath+"`.")
	}
	if len(parts) == 0 {
		parts = append(parts, "The HTTP response was received and cached.")
	}
	return strings.Join(parts, "\n\n")
}

func firstNonempty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return "the cache path in frontmatter"
}
