// Package mcpserver exposes the application runtime through MCP.
package mcpserver

import (
	"context"
	"errors"
	"io"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/sairaph/apis-mcp/internal/app"
	"github.com/sairaph/apis-mcp/internal/bootstrap"
	"github.com/sairaph/apis-mcp/library"
)

const (
	serverName         = "apis-mcp"
	developmentVersion = "dev"
)

// Service binds one process-pinned runtime to an MCP server.
type Service struct {
	runtime *bootstrap.Runtime
	server  *mcp.Server
}

// New constructs the complete seven-tool MCP service.
func New(runtime *bootstrap.Runtime, versions ...string) (*Service, error) {
	if runtime == nil {
		return nil, errors.New("runtime is required")
	}
	if runtime.Library == nil {
		return nil, errors.New("runtime library is required")
	}
	if runtime.HTTP == nil {
		return nil, errors.New("runtime HTTP service is required")
	}
	if runtime.Sessions == nil {
		return nil, errors.New("runtime session manager is required")
	}

	server := mcp.NewServer(
		&mcp.Implementation{Name: serverName, Version: resolvedVersion(versions)},
		&mcp.ServerOptions{Capabilities: &mcp.ServerCapabilities{}},
	)
	service := &Service{runtime: runtime, server: server}
	service.registerTools()
	server.AddReceivingMiddleware(renderSDKToolErrors)
	return service, nil
}

// NewServer constructs the underlying SDK server for callers that do not need
// the Service wrapper.
func NewServer(runtime *bootstrap.Runtime, versions ...string) (*mcp.Server, error) {
	service, err := New(runtime, versions...)
	if err != nil {
		return nil, err
	}
	return service.server, nil
}

// Server returns the underlying official SDK server.
func (s *Service) Server() *mcp.Server { return s.server }

// Run serves one MCP connection over transport. In-memory transports can be
// supplied by tests and embedders.
func (s *Service) Run(ctx context.Context, transport mcp.Transport) error {
	if transport == nil {
		return errors.New("MCP transport is required")
	}
	return s.server.Run(ctx, transport)
}

// RunStdio serves MCP framing on stdin/stdout. The SDK is the only component
// that writes stdout.
func (s *Service) RunStdio(ctx context.Context) error {
	return normalizeRunError(s.Run(ctx, &mcp.StdioTransport{}))
}

// ServeStdio constructs and runs a service over stdin/stdout.
func ServeStdio(ctx context.Context, runtime *bootstrap.Runtime, versions ...string) error {
	service, err := New(runtime, versions...)
	if err != nil {
		return err
	}
	return service.RunStdio(ctx)
}

func resolvedVersion(versions []string) string {
	if len(versions) > 0 && strings.TrimSpace(versions[0]) != "" {
		return strings.TrimSpace(versions[0])
	}
	return developmentVersion
}

func normalizeRunError(err error) error {
	if errors.Is(err, io.EOF) || errors.Is(err, context.Canceled) {
		return nil
	}
	return err
}

func (s *Service) registerTools() {
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "apis_collections",
		Description: "List documentation collections and their stable identifiers.",
		InputSchema: collectionsSchema(),
	}, s.collections)
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "apis_list",
		Description: "List and filter available API documentation sets and versions.",
		InputSchema: listSchema(),
	}, s.list)
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "apis_pages",
		Description: "Browse pages and direct child paths in one documentation set.",
		InputSchema: pagesSchema(),
	}, s.pages)
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "apis_search",
		Description: "Search one API documentation set using terms and quoted phrases.",
		InputSchema: searchSchema(),
	}, s.search)
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "apis_read",
		Description: "Read a documentation page or an inclusive source-line range.",
		InputSchema: readSchema(),
	}, s.read)
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "apis_call",
		Description: "Execute an HTTP request, cache the complete response, and return a bounded preview. Large retry values can keep a call active for a long time.",
		InputSchema: callSchema(s.runtime),
	}, s.call)
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "apis_sessions",
		Description: "List, inspect, or delete persistent HTTP cookie sessions.",
		InputSchema: sessionsSchema(),
	}, s.sessions)
}

func (s *Service) collections(ctx context.Context, _ *mcp.CallToolRequest, input app.CollectionsInput) (*mcp.CallToolResult, any, error) {
	result, err := s.runtime.Library.Collections(ctx, library.CollectionsRequest(input))
	if err != nil {
		return errorResult("apis_collections", err), nil, nil
	}
	return successResult(result, collectionsBody(result, input)), nil, nil
}

func (s *Service) list(ctx context.Context, _ *mcp.CallToolRequest, input app.ListInput) (*mcp.CallToolResult, any, error) {
	result, err := s.runtime.Library.List(ctx, library.ListRequest(input))
	if err != nil {
		return errorResult("apis_list", err), nil, nil
	}
	return successResult(result, listBody(result, input)), nil, nil
}

func (s *Service) pages(ctx context.Context, _ *mcp.CallToolRequest, input app.PagesInput) (*mcp.CallToolResult, any, error) {
	result, err := s.runtime.Library.Pages(ctx, library.PagesRequest(input))
	if err != nil {
		return errorResult("apis_pages", err), nil, nil
	}
	return successResult(result, pagesBody(result, input)), nil, nil
}

func (s *Service) search(ctx context.Context, _ *mcp.CallToolRequest, input app.SearchInput) (*mcp.CallToolResult, any, error) {
	result, err := s.runtime.Library.Search(ctx, library.SearchRequest(input))
	if err != nil {
		return errorResult("apis_search", err), nil, nil
	}
	return successResult(result, searchBody(result, input)), nil, nil
}

func (s *Service) read(ctx context.Context, _ *mcp.CallToolRequest, input app.ReadInput) (*mcp.CallToolResult, any, error) {
	result, err := s.runtime.Library.Read(ctx, library.ReadRequest(input))
	if err != nil {
		return errorResult("apis_read", err), nil, nil
	}
	return successResult(readFront(result), readBody(result)), nil, nil
}

func (s *Service) call(ctx context.Context, _ *mcp.CallToolRequest, input app.CallInput) (*mcp.CallToolResult, any, error) {
	result, err := s.runtime.HTTP.Call(ctx, input)
	if err != nil {
		if result.Request.ID != "" || len(result.Attempts) > 0 || len(result.Redirects) > 0 || result.Cache.Directory != "" {
			return callErrorResult(result, err), nil, nil
		}
		return errorResult("apis_call", err), nil, nil
	}
	front := makeCallFront(result)
	return successResult(front, callBody(result)), nil, nil
}

func (s *Service) sessions(ctx context.Context, _ *mcp.CallToolRequest, input app.SessionsInput) (*mcp.CallToolResult, any, error) {
	front, body, err := sessionResult(ctx, s.runtime, input)
	if err != nil {
		return errorResult("apis_sessions", err), nil, nil
	}
	return successResult(front, body), nil, nil
}

func objectSchema(properties map[string]any, required ...string) map[string]any {
	schema := map[string]any{
		"type":                 "object",
		"properties":           properties,
		"additionalProperties": false,
	}
	if len(required) > 0 {
		schema["required"] = required
	}
	return schema
}

func stringProperty(description string) map[string]any {
	return map[string]any{"type": "string", "description": description}
}

func pageProperty() map[string]any {
	return map[string]any{"type": "integer", "minimum": 1, "default": 1, "description": "One-based result page; defaults to 1"}
}

func collectionsSchema() map[string]any {
	return objectSchema(map[string]any{"page": pageProperty()})
}

func listSchema() map[string]any {
	return objectSchema(map[string]any{
		"name":       stringProperty("Case-insensitive substring matched against API names"),
		"version":    stringProperty("Case-insensitive exact version match"),
		"collection": stringProperty("Case-insensitive exact collection identifier"),
		"page":       pageProperty(),
	})
}

func pagesSchema() map[string]any {
	return objectSchema(map[string]any{
		"doc_id": stringProperty("Documentation version returned by apis_list"),
		"path":   stringProperty("Exact documentation category path"),
		"page":   pageProperty(),
	}, "doc_id")
}

func searchSchema() map[string]any {
	return objectSchema(map[string]any{
		"doc_id": stringProperty("Documentation version to search"),
		"query":  stringProperty("Terms and quoted phrases to find"),
		"path":   stringProperty("Restrict results to this path and descendants"),
		"page":   pageProperty(),
	}, "doc_id", "query")
}

func readSchema() map[string]any {
	return objectSchema(map[string]any{
		"doc_id":  stringProperty("Documentation version containing the page"),
		"page_id": stringProperty("Page identifier returned by apis_pages or apis_search"),
		"lines": map[string]any{
			"type": "array", "items": map[string]any{"type": "integer", "minimum": 1},
			"minItems": 2, "maxItems": 2, "description": "Inclusive one-based start and end line",
		},
	}, "doc_id", "page_id")
}

func callSchema(runtime *bootstrap.Runtime) map[string]any {
	maximumTimeout := runtime.Config.MaximumHeaderTimeout
	if maximumTimeout < 30 {
		maximumTimeout = 600
	}
	maximumRetries := runtime.Config.MaximumRetries
	if maximumRetries < 0 {
		maximumRetries = 0
	}
	headerValue := map[string]any{"oneOf": []any{
		map[string]any{"type": "string"},
		map[string]any{"type": "array", "items": map[string]any{"type": "string"}},
	}}
	properties := map[string]any{
		"method":   stringProperty("Case-sensitive HTTP method token other than CONNECT"),
		"endpoint": stringProperty("Complete absolute HTTP or HTTPS URL"),
		"headers": map[string]any{"oneOf": []any{
			map[string]any{"type": "string"},
			map[string]any{"type": "object", "additionalProperties": headerValue},
		}, "description": "Inline header object or local JSON file path"},
		"payload": map[string]any{"oneOf": []any{
			map[string]any{"type": "string"}, map[string]any{"type": "object"}, map[string]any{"type": "array"},
		}, "description": "Inline JSON object or array, or local JSON file path"},
		"timeout": map[string]any{
			"type": "integer", "minimum": 1, "maximum": maximumTimeout, "default": 30,
			"description": "Response-header timeout in seconds; defaults to 30",
		},
		"retries": map[string]any{
			"type": "integer", "minimum": 0, "maximum": maximumRetries,
			"description": "Retries after the initial attempt",
		},
		"json_path": stringProperty("RFC 9535 JSONPath for the embedded preview"),
		"session":   stringProperty("Existing server-generated cookie session ID"),
	}
	if runtime.Config.AllowLargeDownload {
		properties["allow_large_download"] = map[string]any{
			"type":        "boolean",
			"default":     false,
			"description": "Bypass the configured response-size cap",
		}
	}
	return objectSchema(properties, "method", "endpoint")
}

func sessionsSchema() map[string]any {
	return objectSchema(map[string]any{
		"id":     stringProperty("Server-generated cookie session ID"),
		"delete": map[string]any{"type": "boolean", "default": false, "description": "Delete the identified session"},
		"page":   pageProperty(),
	})
}

func invalidInput(message, hint string, fields map[string]any) error {
	return &app.Error{Code: "invalid_input", Message: message, Hint: hint, Fields: fields}
}

func renderSDKToolErrors(next mcp.MethodHandler) mcp.MethodHandler {
	return func(ctx context.Context, method string, request mcp.Request) (mcp.Result, error) {
		result, err := next(ctx, method, request)
		if err != nil || method != "tools/call" {
			return result, err
		}
		callResult, ok := result.(*mcp.CallToolResult)
		if !ok || !callResult.IsError || resultHasDocument(callResult) {
			return result, nil
		}
		name := "tool"
		if call, ok := request.(*mcp.CallToolRequest); ok {
			name = call.Params.Name
		}
		message := "invalid tool arguments"
		if len(callResult.Content) > 0 {
			if text, ok := callResult.Content[0].(*mcp.TextContent); ok && text.Text != "" {
				message = text.Text
			}
		}
		return errorResult(name, invalidInput(message, "Correct the arguments and retry "+name+".", nil)), nil
	}
}

func resultHasDocument(result *mcp.CallToolResult) bool {
	if len(result.Content) != 1 {
		return false
	}
	text, ok := result.Content[0].(*mcp.TextContent)
	return ok && len(text.Text) >= 4 && text.Text[:4] == "---\n"
}
