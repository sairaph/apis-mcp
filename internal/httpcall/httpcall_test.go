package httpcall

import (
	"compress/flate"
	"compress/gzip"
	"compress/zlib"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync/atomic"
	"testing"
	"time"
	"unicode/utf8"

	"github.com/sairaph/apis-mcp/internal/budget"
	"github.com/sairaph/apis-mcp/internal/cache"
	"github.com/sairaph/apis-mcp/internal/sessions"
)

func testService(t *testing.T, max int64, configure func(*Config)) *Service {
	t.Helper()
	root := t.TempDir()
	store, err := cache.New(cache.Config{Root: filepath.Join(root, "cache"), MaxDecodedBytes: max, Retention: time.Hour})
	if err != nil {
		t.Fatal(err)
	}
	manager, err := sessions.New(filepath.Join(root, "sessions"), time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	cfg := DefaultConfig(store, manager)
	cfg.BackgroundAfter = 0
	cfg.Sleep = func(context.Context, time.Duration) error { return nil }
	if configure != nil {
		configure(&cfg)
	}
	service, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = service.Shutdown(context.Background()) })
	return service
}

func testServiceWithStore(t *testing.T, store *cache.Store, configure func(*Config)) *Service {
	t.Helper()
	manager, err := sessions.New(filepath.Join(t.TempDir(), "sessions"), time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	cfg := DefaultConfig(store, manager)
	cfg.BackgroundAfter = 0
	cfg.Sleep = func(context.Context, time.Duration) error { return nil }
	if configure != nil {
		configure(&cfg)
	}
	service, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = service.Shutdown(context.Background()) })
	return service
}

func TestStatusSemanticsRetriesAndPublication(t *testing.T) {
	var requests atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempt := requests.Add(1)
		w.Header().Set("Content-Type", "application/json")
		if attempt < 3 {
			w.WriteHeader(http.StatusServiceUnavailable)
		} else {
			fmt.Fprint(w, `{"ok":true}`)
		}
	}))
	defer server.Close()
	service := testService(t, 1<<20, nil)
	retries := 2
	result, err := service.Call(context.Background(), Input{Method: "GET", Endpoint: server.URL, Retries: &retries})
	if err != nil {
		t.Fatal(err)
	}
	if result.Response.Status != 200 || len(result.Attempts) != 3 {
		t.Fatalf("result = %#v", result)
	}
	if result.Attempts[0].RetryDelay != 5*time.Second || result.Attempts[1].RetryDelay != 15*time.Second {
		t.Fatalf("retry delays = %#v", result.Attempts)
	}
	if result.Preview == nil || result.Preview.Kind != "json" {
		t.Fatalf("preview = %#v", result.Preview)
	}
	if _, err := os.Stat(result.Cache.BodyPath); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(result.Cache.BodyPath + ".temp"); !os.IsNotExist(err) {
		t.Fatalf("temporary file: %v", err)
	}

	zero := 0
	result, err = service.Call(context.Background(), Input{Method: "GET", Endpoint: server.URL + "/error", Retries: &zero})
	if err != nil {
		t.Fatal(err)
	}
	if result.Response.Status != 200 {
		t.Fatalf("status = %d", result.Response.Status)
	}
}

func TestErrorStatusIsResult(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { http.Error(w, "bad", http.StatusBadRequest) }))
	defer server.Close()
	service := testService(t, 1<<20, nil)
	zero := 0
	result, err := service.Call(context.Background(), Input{Method: "POST", Endpoint: server.URL, Retries: &zero})
	if err != nil {
		t.Fatal(err)
	}
	if result.Response.Status != 400 || result.Response.State != "complete" {
		t.Fatalf("response = %#v", result.Response)
	}
}

func TestTransientStatusAfterRetriesIsResult(t *testing.T) {
	var requests atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requests.Add(1)
		http.Error(w, "later", http.StatusServiceUnavailable)
	}))
	defer server.Close()
	service := testService(t, 1<<20, nil)
	result, err := service.Call(context.Background(), Input{Method: "GET", Endpoint: server.URL, Retries: intPtr(1)})
	if err != nil {
		t.Fatal(err)
	}
	if result.Response.Status != 503 || requests.Load() != 2 {
		t.Fatalf("result/requests = %#v / %d", result, requests.Load())
	}
}

func TestPermanentTLSCertificateFailureIsNotRetried(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	defer server.Close()
	service := testService(t, 1<<20, nil)
	result, err := service.Call(context.Background(), Input{Method: "GET", Endpoint: server.URL, Retries: intPtr(3)})
	var callErr *Error
	if !errors.As(err, &callErr) || callErr.Code != "transport_error" {
		t.Fatalf("error = %#v", err)
	}
	if len(result.Attempts) != 1 || result.Attempts[0].RetryReason != "" {
		t.Fatalf("TLS attempts = %#v", result.Attempts)
	}
}

func TestRedirectPolicyFailureIsNotRetried(t *testing.T) {
	var requests atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests.Add(1)
		http.Redirect(w, r, "ftp://example.com/file", http.StatusFound)
	}))
	defer server.Close()
	service := testService(t, 1<<20, nil)
	result, err := service.Call(context.Background(), Input{Method: "GET", Endpoint: server.URL, Retries: intPtr(3)})
	var callErr *Error
	if !errors.As(err, &callErr) || callErr.Code != "transport_error" {
		t.Fatalf("error = %#v", err)
	}
	if requests.Load() != 1 || len(result.Attempts) != 1 || result.Attempts[0].RetryReason != "" {
		t.Fatalf("requests/attempts = %d / %#v", requests.Load(), result.Attempts)
	}
}

func TestShutdownCancelsRetrySleepAndPreventsAnotherAttempt(t *testing.T) {
	var requests atomic.Int32
	sleeping := make(chan struct{})
	service := testService(t, 1<<20, func(cfg *Config) {
		cfg.Transport = roundTripperFunc(func(*http.Request) (*http.Response, error) {
			requests.Add(1)
			return nil, temporaryError("temporary transport failure")
		})
		cfg.Sleep = func(ctx context.Context, _ time.Duration) error {
			close(sleeping)
			<-ctx.Done()
			return ctx.Err()
		}
	})
	type callResult struct {
		result Result
		err    error
	}
	completed := make(chan callResult, 1)
	go func() {
		result, err := service.Call(context.Background(), Input{Method: "GET", Endpoint: "https://example.com", Retries: intPtr(2)})
		completed <- callResult{result: result, err: err}
	}()
	select {
	case <-sleeping:
	case <-time.After(time.Second):
		t.Fatal("call did not enter retry sleep")
	}
	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := service.Shutdown(shutdownCtx); err != nil {
		t.Fatalf("shutdown: %v", err)
	}
	outcome := <-completed
	if !errors.Is(outcome.err, context.Canceled) {
		t.Fatalf("call error = %v", outcome.err)
	}
	if requests.Load() != 1 || len(outcome.result.Attempts) != 1 || outcome.result.Attempts[0].RetryReason == "" {
		t.Fatalf("requests/result = %d / %#v", requests.Load(), outcome.result)
	}
}

func TestCookiesPersistAcrossCalls(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/set" {
			http.SetCookie(w, &http.Cookie{Name: "sid", Value: "abc", Path: "/"})
			return
		}
		cookie, err := r.Cookie("sid")
		if err != nil {
			http.Error(w, "missing", 400)
			return
		}
		fmt.Fprint(w, cookie.Value)
	}))
	defer server.Close()
	service := testService(t, 1<<20, nil)
	first, err := service.Call(context.Background(), Input{Method: "GET", Endpoint: server.URL + "/set", Retries: intPtr(0)})
	if err != nil {
		t.Fatal(err)
	}
	second, err := service.Call(context.Background(), Input{Method: "GET", Endpoint: server.URL + "/get", Session: first.Request.SessionID, Retries: intPtr(0)})
	if err != nil {
		t.Fatal(err)
	}
	if second.Preview == nil || second.Preview.Content != "abc" {
		t.Fatalf("preview = %#v", second.Preview)
	}
}

func TestRedirectLimitUsesGoBehavior(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			http.Redirect(w, r, "/one", 302)
		case "/one":
			http.Redirect(w, r, "/two", 302)
		default:
			fmt.Fprint(w, "done")
		}
	}))
	defer server.Close()
	service := testService(t, 1<<20, func(cfg *Config) { cfg.MaximumRedirects = 1 })
	result, err := service.Call(context.Background(), Input{Method: "GET", Endpoint: server.URL, Retries: intPtr(0)})
	if err != nil {
		t.Fatal(err)
	}
	if result.Response.Status != 302 || len(result.Redirects) != 1 {
		t.Fatalf("result = %#v", result)
	}
}

func TestLocalHeaderAndPayloadFiles(t *testing.T) {
	dir := t.TempDir()
	headers := filepath.Join(dir, "headers.json")
	payload := filepath.Join(dir, "payload.json")
	if err := os.WriteFile(headers, []byte(`{"X-Test":["one","two"]}`), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(payload, []byte(`{"name":"value"}`), 0o600); err != nil {
		t.Fatal(err)
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		raw, _ := io.ReadAll(r.Body)
		json.NewEncoder(w).Encode(map[string]any{"headers": r.Header.Values("X-Test"), "body": string(raw), "content_type": r.Header.Get("Content-Type")})
	}))
	defer server.Close()
	service := testService(t, 1<<20, nil)
	result, err := service.Call(context.Background(), Input{Method: "POST", Endpoint: server.URL, Headers: headers, Payload: payload, Retries: intPtr(0)})
	if err != nil {
		t.Fatal(err)
	}
	if result.Preview == nil || !strings.Contains(result.Preview.Content, `"one"`) {
		t.Fatalf("preview = %#v", result.Preview)
	}
	if strings.Contains(result.Preview.Content, "application/json") {
		t.Fatal("content type was inferred despite supplied headers")
	}
}

func TestDecodedSizeLimit(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("too large"))
	}))
	defer server.Close()
	service := testService(t, 4, nil)
	_, err := service.Call(context.Background(), Input{Method: "GET", Endpoint: server.URL, Retries: intPtr(0)})
	if !errors.Is(err, cache.ErrSizeLimit) {
		t.Fatalf("error = %v", err)
	}
	var callErr *Error
	if !errors.As(err, &callErr) || callErr.Code != "response_too_large" || callErr.Hint == "" {
		t.Fatalf("size-limit classification = %#v", callErr)
	}
	result, err := service.Call(context.Background(), Input{Method: "GET", Endpoint: server.URL, Retries: intPtr(0), AllowLargeDownload: true})
	if err != nil || result.Response.DecodedBytes != 9 {
		t.Fatalf("result/error = %#v / %v", result, err)
	}
}

func TestCompressionDecoding(t *testing.T) {
	for _, encoding := range []string{"gzip", "deflate", "raw-deflate"} {
		t.Run(encoding, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("Content-Encoding", strings.TrimPrefix(encoding, "raw-"))
				w.Header().Set("Content-Type", "application/json")
				var writer io.WriteCloser
				switch encoding {
				case "gzip":
					writer = gzip.NewWriter(w)
				case "deflate":
					writer = zlib.NewWriter(w)
				default:
					writer, _ = flate.NewWriter(w, flate.DefaultCompression)
				}
				writer.Write([]byte(`{"compressed":true}`))
				writer.Close()
			}))
			defer server.Close()
			service := testService(t, 1<<20, nil)
			result, err := service.Call(context.Background(), Input{Method: "GET", Endpoint: server.URL, Retries: intPtr(0)})
			if err != nil {
				t.Fatal(err)
			}
			raw, err := os.ReadFile(result.Cache.BodyPath)
			if err != nil {
				t.Fatal(err)
			}
			if string(raw) != `{"compressed":true}` || !result.Response.Decoded {
				t.Fatalf("body/result = %q / %#v", raw, result.Response)
			}
		})
	}
}

func TestUnsupportedEncodingCachesWireBodyWithoutPreview(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Encoding", "br")
		w.Write([]byte{0, 1, 2})
	}))
	defer server.Close()
	service := testService(t, 1<<20, nil)
	result, err := service.Call(context.Background(), Input{Method: "GET", Endpoint: server.URL, Retries: intPtr(0)})
	if err != nil {
		t.Fatal(err)
	}
	if result.Response.Decoded || result.Preview != nil || filepath.Ext(result.Cache.BodyPath) != ".encoded" {
		t.Fatalf("result = %#v", result)
	}
}

func TestJSONPathRFC9535Selectors(t *testing.T) {
	root := map[string]any{"items": []any{
		map[string]any{"name": "one", "price": 5.0},
		map[string]any{"name": "two", "price": 12.0},
		map[string]any{"name": "three", "price": 8.0},
	}}
	tests := []struct {
		name       string
		expression string
		want       any
	}{
		{name: "filter", expression: `$.items[?@.price < 10].name`, want: []any{"one", "three"}},
		{name: "slice", expression: `$.items[0:3:2].name`, want: []any{"one", "three"}},
		{name: "recursive", expression: `$..name`, want: []any{"one", "two", "three"}},
		{name: "singular", expression: `$.items[0].name`, want: "one"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := selectJSON(root, tc.expression)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("selection = %#v, want %#v", got, tc.want)
			}
		})
	}
}

func TestJSONPathInvalidNoMatchAndBoundedPreview(t *testing.T) {
	path := filepath.Join(t.TempDir(), "body.json")
	if err := os.WriteFile(path, []byte(`{"items":[{"name":"one"},{"name":"two"}]}`), 0o600); err != nil {
		t.Fatal(err)
	}
	for _, expression := range []string{`$.items[?`, `$.missing`} {
		preview, selection, err := makePreview(path, "application/json", true, 8, expression, 1<<20)
		if err != nil {
			t.Fatal(err)
		}
		if preview != nil || selection == nil || selection.Matched || selection.Error == "" {
			t.Fatalf("%q preview/selection = %#v / %#v", expression, preview, selection)
		}
	}
	preview, selection, err := makePreview(path, "application/json", true, 8, `$.items[*].name`, 1<<20)
	if err != nil {
		t.Fatal(err)
	}
	if preview == nil || !preview.Truncated || preview.ApproximateTokens > 8 || selection == nil || !selection.Matched {
		t.Fatalf("bounded preview/selection = %#v / %#v", preview, selection)
	}
}

func TestJSONPathUsesCompleteLargeBodyUpToResponseCap(t *testing.T) {
	body := `{"padding":"` + strings.Repeat("x", 9<<20) + `","tail":{"value":"found"}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, body)
	}))
	defer server.Close()
	service := testService(t, 12<<20, func(cfg *Config) { cfg.ReadTokenBudget = 32 })
	result, err := service.Call(context.Background(), Input{Method: "GET", Endpoint: server.URL, Retries: intPtr(0), JSONPath: `$.tail.value`})
	if err != nil {
		t.Fatal(err)
	}
	if result.Selection == nil || !result.Selection.Matched || result.Preview == nil || result.Preview.Content != `"found"` {
		t.Fatalf("selection/preview = %#v / %#v", result.Selection, result.Preview)
	}
	if result.Preview.ApproximateTokens > 32 {
		t.Fatalf("embedded preview exceeded token bound: %d", result.Preview.ApproximateTokens)
	}
}

func TestPreviewTokenBoundaryUnicodeAndAccounting(t *testing.T) {
	path := filepath.Join(t.TempDir(), "body.txt")
	body := "hello world " + strings.Repeat("🙂", 20)
	if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}

	full, _, err := makePreview(path, "text/plain; charset=utf-8", true, 100, "", 1<<20)
	if err != nil || full == nil || full.Truncated || full.Content != body {
		t.Fatalf("full preview = %#v, %v", full, err)
	}
	actual, err := budget.Count(full.Content)
	if err != nil || actual != full.ApproximateTokens {
		t.Fatalf("full accounting = %d, %v; preview = %#v", actual, err, full)
	}

	limited, _, err := makePreview(path, "text/plain; charset=utf-8", true, 5, "", 1<<20)
	if err != nil || limited == nil || !limited.Truncated || !utf8.ValidString(limited.Content) {
		t.Fatalf("limited preview = %#v, %v", limited, err)
	}
	actual, err = budget.Count(limited.Content)
	if err != nil || actual != limited.ApproximateTokens || actual > 5 {
		t.Fatalf("limited accounting = %d, %v; preview = %#v", actual, err, limited)
	}
}

func TestPreviewRetainsByteHardCeiling(t *testing.T) {
	path := filepath.Join(t.TempDir(), "body.txt")
	body := strings.Repeat("word ", maxPreviewBytes/5+100)
	if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
	preview, _, err := makePreview(path, "text/plain", true, maxPreviewBytes, "", 1<<20)
	if err != nil {
		t.Fatal(err)
	}
	if preview == nil || !preview.Truncated || len(preview.Content) > maxPreviewBytes {
		t.Fatalf("hard-ceiling preview = %#v, bytes = %d", preview, len(preview.Content))
	}
}

func TestPostPublicationMetadataFailureKeepsBodyAuthoritative(t *testing.T) {
	store, err := cache.New(cache.Config{
		Root: filepath.Join(t.TempDir(), "cache"), MaxDecodedBytes: 1 << 20, Retention: time.Hour,
		Faults: &cache.Faults{SaveMetadata: func(metadata cache.Metadata) error {
			if metadata.State == "complete" {
				return errors.New("injected completed metadata failure")
			}
			return nil
		}},
	})
	if err != nil {
		t.Fatal(err)
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { _, _ = io.WriteString(w, "complete") }))
	defer server.Close()
	service := testServiceWithStore(t, store, nil)
	result, err := service.Call(context.Background(), Input{Method: "GET", Endpoint: server.URL, Retries: intPtr(0)})
	if err != nil || result.Response.State != "complete" {
		t.Fatalf("result/error = %#v / %v", result, err)
	}
	if raw, err := os.ReadFile(result.Cache.BodyPath); err != nil || string(raw) != "complete" {
		t.Fatalf("published body = %q, %v", raw, err)
	}
	if _, err := os.Stat(result.Cache.BodyPath + ".error"); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("error artifact exists beside final body: %v", err)
	}
	if cleanup, err := store.Cleanup(); err != nil || cleanup.RemovedOrphans != 0 {
		t.Fatalf("cleanup treated final body as orphan: %#v, %v", cleanup, err)
	}
}

func TestBodyPublicationFailurePublishesOnlyErrorArtifact(t *testing.T) {
	store, err := cache.New(cache.Config{
		Root: filepath.Join(t.TempDir(), "cache"), MaxDecodedBytes: 1 << 20, Retention: time.Hour,
		Faults: &cache.Faults{PublishBody: func() error { return errors.New("injected rename failure") }},
	})
	if err != nil {
		t.Fatal(err)
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { _, _ = io.WriteString(w, "complete") }))
	defer server.Close()
	service := testServiceWithStore(t, store, nil)
	result, err := service.Call(context.Background(), Input{Method: "GET", Endpoint: server.URL, Retries: intPtr(0)})
	var callErr *Error
	if !errors.As(err, &callErr) || callErr.Code != "cache_error" {
		t.Fatalf("error = %#v", err)
	}
	if _, err := os.Stat(result.Cache.FinalPath); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("final body exists after publication failure: %v", err)
	}
	if _, err := os.Stat(result.Cache.TempPath); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("temporary body remains after publication failure: %v", err)
	}
	if _, err := os.Stat(result.Cache.ErrorPath); err != nil {
		t.Fatalf("error artifact was not published: %v", err)
	}
}

func TestResponseHeaderTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { time.Sleep(1200 * time.Millisecond); w.WriteHeader(200) }))
	defer server.Close()
	service := testService(t, 1<<20, nil)
	_, err := service.Call(context.Background(), Input{Method: "GET", Endpoint: server.URL, Timeout: 1, Retries: intPtr(0)})
	var callErr *Error
	if !errors.As(err, &callErr) || callErr.Code != "response_header_timeout" {
		t.Fatalf("error = %#v", err)
	}
}

func TestBackgroundTransitionPublishesFinalBody(t *testing.T) {
	release := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("first"))
		w.(http.Flusher).Flush()
		<-release
		w.Write([]byte("second"))
	}))
	defer server.Close()
	service := testService(t, 1<<20, func(cfg *Config) { cfg.BackgroundAfter = 20 * time.Millisecond })
	result, err := service.Call(context.Background(), Input{Method: "GET", Endpoint: server.URL, Retries: intPtr(0)})
	if err != nil {
		t.Fatal(err)
	}
	if result.Response.State != "downloading" {
		t.Fatalf("state = %s", result.Response.State)
	}
	if _, err := os.Stat(result.Cache.TempPath); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(result.Cache.FinalPath); !os.IsNotExist(err) {
		t.Fatalf("final path exists early: %v", err)
	}
	close(release)
	deadline := time.Now().Add(2 * time.Second)
	for {
		if raw, err := os.ReadFile(result.Cache.FinalPath); err == nil {
			if string(raw) != "firstsecond" {
				t.Fatalf("body = %q", raw)
			}
			break
		}
		if time.Now().After(deadline) {
			t.Fatal("background body was not published")
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func TestStalledDownloadPublishesTimeout(t *testing.T) {
	for _, tc := range []struct {
		name       string
		background time.Duration
	}{
		{name: "foreground"},
		{name: "background", background: 5 * time.Millisecond},
	} {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "text/plain")
				_, _ = w.Write([]byte("first"))
				w.(http.Flusher).Flush()
				<-r.Context().Done()
			}))
			defer server.Close()
			service := testService(t, 1<<20, func(cfg *Config) {
				cfg.BackgroundAfter = tc.background
				cfg.StalledDownloadAfter = 30 * time.Millisecond
			})
			result, err := service.Call(context.Background(), Input{Method: "GET", Endpoint: server.URL, Retries: intPtr(0)})
			if tc.background == 0 {
				var callErr *Error
				if !errors.As(err, &callErr) || callErr.Code != "cache_error" {
					t.Fatalf("error = %#v", err)
				}
			} else if err != nil || result.Response.State != "downloading" {
				t.Fatalf("result/error = %#v / %v", result, err)
			}

			deadline := time.Now().Add(2 * time.Second)
			var record cache.ErrorRecord
			for {
				raw, readErr := os.ReadFile(result.Cache.ErrorPath)
				if readErr == nil {
					if err := json.Unmarshal(raw, &record); err != nil {
						t.Fatal(err)
					}
					break
				}
				if time.Now().After(deadline) {
					t.Fatalf("timeout artifact was not published: %v", readErr)
				}
				time.Sleep(5 * time.Millisecond)
			}
			if record.Code != "download_timeout" || record.BytesRead != int64(len("first")) || !strings.Contains(record.Message, "no body bytes received") {
				t.Fatalf("error record = %#v", record)
			}
			if _, err := os.Stat(result.Cache.TempPath); !os.IsNotExist(err) {
				t.Fatalf("temporary body still exists: %v", err)
			}
		})
	}
}

func TestStalledDownloadTimerResetsAfterBodyReads(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		for _, chunk := range []string{"one", "two", "three", "four", "five"} {
			_, _ = w.Write([]byte(chunk))
			w.(http.Flusher).Flush()
			time.Sleep(20 * time.Millisecond)
		}
	}))
	defer server.Close()
	service := testService(t, 1<<20, func(cfg *Config) { cfg.StalledDownloadAfter = 60 * time.Millisecond })
	result, err := service.Call(context.Background(), Input{Method: "GET", Endpoint: server.URL, Retries: intPtr(0)})
	if err != nil {
		t.Fatal(err)
	}
	if result.Response.State != "complete" || result.Response.DecodedBytes != int64(len("onetwothreefourfive")) {
		t.Fatalf("response = %#v", result.Response)
	}
}

func intPtr(value int) *int { return &value }

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return f(request)
}

type temporaryError string

func (e temporaryError) Error() string { return string(e) }
func (temporaryError) Timeout() bool   { return false }
func (temporaryError) Temporary() bool { return true }
