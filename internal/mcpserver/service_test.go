package mcpserver

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/sairaph/apis-mcp/internal/app"
	"github.com/sairaph/apis-mcp/internal/bootstrap"
	"github.com/sairaph/apis-mcp/internal/budget"
	"github.com/sairaph/apis-mcp/internal/cache"
	"github.com/sairaph/apis-mcp/internal/config"
	"github.com/sairaph/apis-mcp/internal/httpcall"
	"github.com/sairaph/apis-mcp/internal/sessions"
	"github.com/sairaph/apis-mcp/library"
	"gopkg.in/yaml.v3"
)

func TestToolListingAndSchemas(t *testing.T) {
	runtime := testRuntime(t)
	runtime.Config.MaximumHeaderTimeout = 90
	runtime.Config.MaximumRetries = 4
	runtime.Config.AllowLargeDownload = false
	client := testClient(t, runtime)

	result, err := client.ListTools(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	var names []string
	tools := make(map[string]*mcp.Tool)
	for _, tool := range result.Tools {
		names = append(names, tool.Name)
		tools[tool.Name] = tool
	}
	slices.Sort(names)
	want := []string{"apis_call", "apis_collections", "apis_list", "apis_pages", "apis_read", "apis_search", "apis_sessions"}
	if !slices.Equal(names, want) {
		t.Fatalf("tools = %v, want %v", names, want)
	}

	search := schemaMap(t, tools["apis_search"].InputSchema)
	if got := strings.Join(stringSlice(search["required"]), ","); got != "doc_id,query" {
		t.Fatalf("search required = %q", got)
	}
	if search["additionalProperties"] != false {
		t.Fatalf("search additionalProperties = %#v", search["additionalProperties"])
	}
	searchProperties := schemaMap(t, search["properties"])
	if page := schemaMap(t, searchProperties["page"]); page["default"] != float64(1) {
		t.Fatalf("page default = %#v", page["default"])
	}
	callProperties := schemaMap(t, schemaMap(t, tools["apis_call"].InputSchema)["properties"])
	if _, exists := callProperties["allow_large_download"]; exists {
		t.Fatal("allow_large_download exposed while override policy is disabled")
	}
	timeout := schemaMap(t, callProperties["timeout"])
	if timeout["maximum"] != float64(90) {
		t.Fatalf("timeout maximum = %#v", timeout["maximum"])
	}
	if timeout["default"] != float64(30) {
		t.Fatalf("timeout default = %#v", timeout["default"])
	}
	retries := schemaMap(t, callProperties["retries"])
	if retries["maximum"] != float64(4) {
		t.Fatalf("retries maximum = %#v", retries["maximum"])
	}
	sessionsProperties := schemaMap(t, schemaMap(t, tools["apis_sessions"].InputSchema)["properties"])
	if deletion := schemaMap(t, sessionsProperties["delete"]); deletion["default"] != false {
		t.Fatalf("delete default = %#v", deletion["default"])
	}
	runtime.Config.AllowLargeDownload = true
	large := schemaMap(t, schemaMap(t, callSchema(runtime))["properties"])
	if override := schemaMap(t, large["allow_large_download"]); override["default"] != false {
		t.Fatalf("allow_large_download default = %#v", override["default"])
	}
}

func TestMCPReportsConfiguredBuildVersion(t *testing.T) {
	client := testClient(t, testRuntime(t), "2.4.6")
	result := client.InitializeResult()
	if result == nil || result.ServerInfo == nil || result.ServerInfo.Version != "2.4.6" {
		t.Fatalf("server implementation = %#v", result)
	}
}

func TestNormalTransportCloseIsNotAnError(t *testing.T) {
	if err := normalizeRunError(fmt.Errorf("server is closing: %w", io.EOF)); err != nil {
		t.Fatalf("EOF close = %v", err)
	}
	if err := normalizeRunError(context.Canceled); err != nil {
		t.Fatalf("canceled close = %v", err)
	}
}

func TestRepresentativeToolCalls(t *testing.T) {
	runtime := testRuntime(t)
	client := testClient(t, runtime)

	collections := callTool(t, client, "apis_collections", map[string]any{})
	assertTextResult(t, collections, false, "---\n", "collections:", "apis_list")

	list := callTool(t, client, "apis_list", map[string]any{"collection": "examples"})
	assertTextResult(t, list, false, "doc_id: example-api-v1", "apis_pages")

	pages := callTool(t, client, "apis_pages", map[string]any{"doc_id": "example-api-v1"})
	assertTextResult(t, pages, false, "page_id: example-api-overview", "path: items")

	search := callTool(t, client, "apis_search", map[string]any{"doc_id": "example-api-v1", "query": "POST"})
	assertTextResult(t, search, false, "matching lines", "| page_id | line |", "apis_read")

	read := callTool(t, client, "apis_read", map[string]any{"doc_id": "example-api-v1", "page_id": "create-an-item"})
	assertTextResult(t, read, false, "~~~markdown\n# Create an item", "```json", "\n~~~\n")

	sessionsResult := callTool(t, client, "apis_sessions", map[string]any{})
	assertTextResult(t, sessionsResult, false, "page: 1", "No cookie sessions found.")

	httpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusTeapot)
		fmt.Fprint(w, `{"error":"short and stout"}`)
	}))
	defer httpServer.Close()
	call := callTool(t, client, "apis_call", map[string]any{
		"method": "GET", "endpoint": httpServer.URL, "retries": 0,
	})
	assertTextResult(t, call, false, "status: 418", "state: complete", "~~~json", `"short and stout"`)

	sessionsResult = callTool(t, client, "apis_sessions", map[string]any{})
	assertTextResult(t, sessionsResult, false, "total: 1", "sessions:", "cookie_count: 0")
}

func TestErrorsRenderAsDocuments(t *testing.T) {
	client := testClient(t, testRuntime(t))
	result := callTool(t, client, "apis_pages", map[string]any{"doc_id": "does-not-exist"})
	assertTextResult(t, result, true, "---\nerror:\n", "code: not_found", "## Error", "Use apis_list and apis_pages")

	result = callTool(t, client, "apis_sessions", map[string]any{"delete": true})
	assertTextResult(t, result, true, "code: invalid_input", "id is required", "Provide a server-generated session id")

	result = callTool(t, client, "apis_search", map[string]any{"doc_id": "example-api-v1"})
	assertTextResult(t, result, true, "code: invalid_input", "## Error", "retry apis_search")
}

func TestCallErrorRendersPartialRetryDiagnostics(t *testing.T) {
	runtime := testRuntime(t)
	if err := runtime.HTTP.Shutdown(context.Background()); err != nil {
		t.Fatal(err)
	}
	var requests atomic.Int32
	httpConfig := httpcall.DefaultConfig(runtime.Cache, runtime.Sessions)
	httpConfig.BackgroundAfter = 0
	httpConfig.Transport = mcpRoundTripperFunc(func(*http.Request) (*http.Response, error) {
		requests.Add(1)
		return nil, mcpTemporaryError("injected temporary failure")
	})
	httpConfig.Sleep = func(context.Context, time.Duration) error { return errors.New("injected retry sleep failure") }
	httpService, err := httpcall.New(httpConfig)
	if err != nil {
		t.Fatal(err)
	}
	runtime.HTTP = httpService
	client := testClient(t, runtime)
	result := callTool(t, client, "apis_call", map[string]any{
		"method": "GET", "endpoint": "https://example.com/resource", "retries": 2,
	})
	assertTextResult(t, result, true,
		"code: internal_error", "request:", "attempts:", "number: 1",
		"injected temporary failure", "retry_reason: transient transport failure", "retry_delay: 5s",
	)
	if requests.Load() != 1 {
		t.Fatalf("transport requests = %d, want 1", requests.Load())
	}
}

func TestReadFenceGrowsAndPreservesContent(t *testing.T) {
	content := "before\n  ~~~~~embedded\nafter\n"
	body := readBody(library.ReadResult{
		DocID: "docs-v1", PageID: "page", Markdown: content,
		Lines: [2]int{1, 3}, TotalLines: 3,
	})
	if !strings.HasPrefix(body, "~~~~~~markdown\n") {
		t.Fatalf("opening fence did not grow: %q", body)
	}
	if !strings.Contains(body, content+"~~~~~~") {
		t.Fatalf("content was not preserved up to the closing fence: %q", body)
	}
	if strings.Contains(body, "after\n\n~~~~~~") {
		t.Fatalf("fence inserted a blank line into verbatim content: %q", body)
	}
}

func TestExactContinuationGuidance(t *testing.T) {
	collections := collectionsBody(library.CollectionsResult{
		Pagination: library.Pagination{Page: 1, Total: 2, TotalPages: 2},
	}, app.CollectionsInput{Page: 1})
	if !strings.HasSuffix(collections, "Continue with `apis_collections({\"page\":2})`.") {
		t.Fatalf("collections continuation = %q", collections)
	}
	list := listBody(library.ListResult{
		Pagination: library.Pagination{Page: 1, Total: 2, TotalPages: 2},
	}, app.ListInput{Name: "example", Version: "v1", Collection: "examples", Page: 1})
	want := "Continue with `apis_list({\"name\":\"example\",\"version\":\"v1\",\"collection\":\"examples\",\"page\":2})`."
	if !strings.HasSuffix(list, want) {
		t.Fatalf("list continuation = %q, want suffix %q", list, want)
	}
}

func TestCallPreviewReportsActualTokens(t *testing.T) {
	front := makeCallFront(httpcall.Result{Preview: &httpcall.Preview{
		Kind: "text", Content: "hello world", ApproximateTokens: 2,
	}})
	if front.Preview == nil || front.Preview.ApproximateTokens != 2 {
		t.Fatalf("preview front = %#v", front.Preview)
	}
}

func TestCallFrontUsesStableCacheCompletionAndSkippedSelection(t *testing.T) {
	completed := time.Date(2026, 7, 26, 12, 0, 0, 0, time.UTC)
	front := makeCallFront(httpcall.Result{
		Response:  httpcall.ResponseResult{State: "complete", CompletedAt: &completed},
		Selection: &httpcall.Selection{JSONPath: "$.items", State: "skipped", Error: "download continued"},
	})
	raw, err := yaml.Marshal(front)
	if err != nil {
		t.Fatal(err)
	}
	text := string(raw)
	if !strings.Contains(text, "cache:\n    completed_at:") || strings.Contains(text, "response:\n    completed_at:") {
		t.Fatalf("completion field is not under cache:\n%s", text)
	}
	if !strings.Contains(text, "state: skipped") {
		t.Fatalf("selection state missing:\n%s", text)
	}
}

func TestSessionPaginationUsesExactRenderedBoundary(t *testing.T) {
	items := []sessionInfo{{ID: "hello"}, {ID: "セッション🙂"}}
	render := func(records []sessionInfo) (string, error) {
		raw, err := yaml.Marshal(struct {
			Sessions []sessionInfo `yaml:"sessions"`
		}{Sessions: records})
		return string(raw), err
	}
	representation, err := render(items)
	if err != nil {
		t.Fatal(err)
	}
	exact, err := budget.Count(representation)
	if err != nil {
		t.Fatal(err)
	}
	window, pagination, err := paginate(items, 1, exact, render)
	if err != nil || pagination.TotalPages != 1 || len(window) != 2 {
		t.Fatalf("exact-boundary sessions = %v / %#v / %v", window, pagination, err)
	}
	window, pagination, err = paginate(items, 1, 1, render)
	if err != nil || pagination.TotalPages != 2 || len(window) != 1 || window[0].ID != "hello" {
		t.Fatalf("oversized-first session = %v / %#v / %v", window, pagination, err)
	}
}

func testRuntime(t *testing.T) *bootstrap.Runtime {
	t.Helper()
	root := t.TempDir()
	libraryRoot := filepath.Join(root, "library")
	writeTestDocumentation(t, libraryRoot)
	snapshot, err := library.Open(context.Background(), library.Options{
		UserRoot: libraryRoot, IndexPath: filepath.Join(root, "library.sqlite"), ListTokenBudget: 2_000, ReadTokenBudget: 4_000, ExcludeBuiltin: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	store, err := cache.New(cache.Config{
		Root: filepath.Join(root, "cache"), MaxDecodedBytes: 1 << 20, Retention: time.Hour,
	})
	if err != nil {
		snapshot.Close()
		t.Fatal(err)
	}
	manager, err := sessions.New(filepath.Join(root, "sessions"), time.Hour)
	if err != nil {
		snapshot.Close()
		t.Fatal(err)
	}
	httpConfig := httpcall.DefaultConfig(store, manager)
	httpConfig.BackgroundAfter = 0
	httpService, err := httpcall.New(httpConfig)
	if err != nil {
		snapshot.Close()
		t.Fatal(err)
	}
	cfg := config.Default()
	runtime := &bootstrap.Runtime{Config: cfg, Library: snapshot, Cache: store, Sessions: manager, HTTP: httpService}
	t.Cleanup(func() {
		if err := runtime.Close(context.Background()); err != nil {
			t.Errorf("close runtime: %v", err)
		}
	})
	return runtime
}

func writeTestDocumentation(t *testing.T, root string) {
	t.Helper()
	files := map[string]string{
		"_index.md":       "---\nname: Example API\nversion: v1\ndescription: Test API.\ncollections: [examples]\n---\n",
		"overview.md":     "---\ntitle: Example API overview\ndescription: Start here.\n---\n\n# Example API overview\n",
		"items/create.md": "---\ntitle: Create an item\nhttp_methods: [POST]\napi_endpoints: [/items]\noperation_ids: [createItem]\n---\n\n# Create an item\n\nSend a POST request.\n\n```json\n{\"name\":\"example\"}\n```\n",
	}
	for relative, content := range files {
		name := filepath.Join(root, "example", "v1", relative)
		if err := os.MkdirAll(filepath.Dir(name), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(name, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
}

func testClient(t *testing.T, runtime *bootstrap.Runtime, versions ...string) *mcp.ClientSession {
	t.Helper()
	service, err := New(runtime, versions...)
	if err != nil {
		t.Fatal(err)
	}
	serverTransport, clientTransport := mcp.NewInMemoryTransports()
	serverSession, err := service.Server().Connect(context.Background(), serverTransport, nil)
	if err != nil {
		t.Fatal(err)
	}
	client := mcp.NewClient(&mcp.Implementation{Name: "mcpserver-test", Version: "1"}, nil)
	clientSession, err := client.Connect(context.Background(), clientTransport, nil)
	if err != nil {
		serverSession.Close()
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := clientSession.Close(); err != nil {
			t.Errorf("close client: %v", err)
		}
		if err := serverSession.Close(); err != nil {
			t.Errorf("close server: %v", err)
		}
	})
	return clientSession
}

func callTool(t *testing.T, client *mcp.ClientSession, name string, arguments map[string]any) *mcp.CallToolResult {
	t.Helper()
	result, err := client.CallTool(context.Background(), &mcp.CallToolParams{Name: name, Arguments: arguments})
	if err != nil {
		t.Fatalf("call %s: %v", name, err)
	}
	return result
}

func assertTextResult(t *testing.T, result *mcp.CallToolResult, isError bool, fragments ...string) {
	t.Helper()
	if result.IsError != isError {
		t.Fatalf("isError = %v, want %v; result = %#v", result.IsError, isError, result)
	}
	if len(result.Content) != 1 {
		t.Fatalf("content count = %d", len(result.Content))
	}
	text, ok := result.Content[0].(*mcp.TextContent)
	if !ok {
		t.Fatalf("content type = %T", result.Content[0])
	}
	if !strings.HasPrefix(text.Text, "---\n") || !strings.Contains(text.Text, "\n---\n\n") {
		t.Fatalf("result is not YAML frontmatter plus Markdown: %q", text.Text)
	}
	for _, fragment := range fragments {
		if !strings.Contains(text.Text, fragment) {
			t.Errorf("result does not contain %q:\n%s", fragment, text.Text)
		}
	}
}

func schemaMap(t *testing.T, value any) map[string]any {
	t.Helper()
	result, ok := value.(map[string]any)
	if !ok {
		t.Fatalf("schema value has type %T", value)
	}
	return result
}

func stringSlice(value any) []string {
	items, _ := value.([]any)
	result := make([]string, 0, len(items))
	for _, item := range items {
		result = append(result, item.(string))
	}
	return result
}

type mcpRoundTripperFunc func(*http.Request) (*http.Response, error)

func (f mcpRoundTripperFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return f(request)
}

type mcpTemporaryError string

func (e mcpTemporaryError) Error() string { return string(e) }
func (mcpTemporaryError) Timeout() bool   { return false }
func (mcpTemporaryError) Temporary() bool { return true }
