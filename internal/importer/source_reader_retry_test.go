package importer

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"syscall"
	"testing"
	"time"
)

type sourceRetryRoundTripFunc func(*http.Request) (*http.Response, error)

func (function sourceRetryRoundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return function(request)
}

type sourceRetryBody struct {
	*strings.Reader
	closed bool
}

type sourceRetryFailingBody struct {
	content []byte
	err     error
	closed  bool
	read    bool
}

func (body *sourceRetryFailingBody) Read(destination []byte) (int, error) {
	if body.read {
		return 0, io.EOF
	}
	body.read = true
	return copy(destination, body.content), body.err
}

func (body *sourceRetryFailingBody) Close() error {
	body.closed = true
	return nil
}

func (body *sourceRetryBody) Close() error {
	body.closed = true
	return nil
}

func sourceRetryResponse(request *http.Request, status int, body io.ReadCloser) *http.Response {
	return &http.Response{
		StatusCode: status,
		Status:     fmt.Sprintf("%d %s", status, http.StatusText(status)),
		Header:     make(http.Header),
		Body:       body,
		Request:    request,
	}
}

func newRetryTestReader(client *http.Client) *sourceReader {
	return &sourceReader{client: client, perSource: DefaultMaxSourceBytes, total: DefaultMaxTotalBytes}
}

func TestSourceReaderRetries503ThenSucceeds(t *testing.T) {
	t.Setenv("GITHUB_TOKEN", "retry-token")
	attempts := 0
	transientBody := &sourceRetryBody{Reader: strings.NewReader("temporarily unavailable")}
	client := &http.Client{Transport: sourceRetryRoundTripFunc(func(request *http.Request) (*http.Response, error) {
		attempts++
		if request.Header.Get("User-Agent") != "apis-mcp-documentation-importer" {
			t.Errorf("attempt %d user agent = %q", attempts, request.Header.Get("User-Agent"))
		}
		if request.Header.Get("Authorization") != "Bearer retry-token" {
			t.Errorf("attempt %d authorization = %q", attempts, request.Header.Get("Authorization"))
		}
		if attempts == 1 {
			return sourceRetryResponse(request, http.StatusServiceUnavailable, transientBody), nil
		}
		return sourceRetryResponse(request, http.StatusOK, io.NopCloser(strings.NewReader("accepted"))), nil
	})}
	reader := newSourceReader(Options{HTTPClient: client, MaxSourceBytes: DefaultMaxSourceBytes, MaxTotalBytes: DefaultMaxTotalBytes})
	var delays []time.Duration
	reader.retryWait = func(_ context.Context, delay time.Duration) error {
		delays = append(delays, delay)
		return nil
	}

	raw, resolved, err := reader.readFromOrigin(context.Background(), "https://api.github.com/repos/example/docs", nil, "https://api.github.com")
	if err != nil {
		t.Fatal(err)
	}
	if string(raw) != "accepted" || resolved != "https://api.github.com/repos/example/docs" {
		t.Fatalf("read = %q, %q", raw, resolved)
	}
	if attempts != 2 || len(delays) != 1 || delays[0] != sourceReadRetryDelay {
		t.Fatalf("attempts = %d, delays = %v", attempts, delays)
	}
	if !transientBody.closed || transientBody.Len() != 0 {
		t.Fatalf("transient body closed = %v, unread = %d", transientBody.closed, transientBody.Len())
	}
	if reader.used != int64(len("accepted")) {
		t.Fatalf("used bytes = %d", reader.used)
	}
}

func TestSourceReaderAllowedOriginRejectsUserinfoRedirect(t *testing.T) {
	targetRequests := 0
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path == "/target" {
			targetRequests++
			fmt.Fprint(writer, "must not be reached")
			return
		}
		location := strings.Replace(server.URL, "://", "://user:secret@", 1) + "/target"
		http.Redirect(writer, request, location, http.StatusFound)
	}))
	defer server.Close()
	parsed, err := url.Parse(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	reader := newRetryTestReader(server.Client())
	_, _, err = reader.readFromOrigin(context.Background(), server.URL+"/start", nil, httpOrigin(parsed))
	if err == nil || !strings.Contains(err.Error(), "redirect URL must not contain credentials") {
		t.Fatalf("userinfo redirect error = %v", err)
	}
	if targetRequests != 0 || reader.used != 0 {
		t.Fatalf("target requests=%d used bytes=%d", targetRequests, reader.used)
	}
}

func TestSourceReaderReturnsExactErrorAfterRetries(t *testing.T) {
	attempts := 0
	client := &http.Client{Transport: sourceRetryRoundTripFunc(func(request *http.Request) (*http.Response, error) {
		attempts++
		return sourceRetryResponse(request, http.StatusServiceUnavailable, io.NopCloser(strings.NewReader("unavailable"))), nil
	})}
	reader := newRetryTestReader(client)
	reader.retryWait = func(context.Context, time.Duration) error { return nil }

	_, _, err := reader.read(context.Background(), "https://example.test/source", nil)
	if err == nil || err.Error() != "fetch https://example.test/source: HTTP 503" {
		t.Fatalf("error = %v", err)
	}
	if attempts != sourceReadAttempts || reader.used != 0 {
		t.Fatalf("attempts = %d, used bytes = %d", attempts, reader.used)
	}
}

func TestSourceReaderRetryAfterHonorsCancellation(t *testing.T) {
	if delay, ok := parseSourceRetryAfter("120", time.Unix(0, 0)); !ok || delay != sourceReadMaxRetry {
		t.Fatalf("bounded Retry-After = %v, %v", delay, ok)
	}

	ctx, cancel := context.WithCancel(context.Background())
	attempts := 0
	client := &http.Client{Transport: sourceRetryRoundTripFunc(func(request *http.Request) (*http.Response, error) {
		attempts++
		cancel()
		response := sourceRetryResponse(request, http.StatusTooManyRequests, io.NopCloser(strings.NewReader("slow down")))
		response.Header.Set("Retry-After", "120")
		return response, nil
	})}
	reader := newRetryTestReader(client)

	_, _, err := reader.read(ctx, "https://example.test/source", nil)
	if err == nil || err.Error() != "fetch https://example.test/source: context canceled" || !errors.Is(err, context.Canceled) {
		t.Fatalf("error = %v", err)
	}
	if attempts != 1 {
		t.Fatalf("attempts after cancellation = %d", attempts)
	}
}

func TestSourceReaderDoesNotRetryPermanent404(t *testing.T) {
	attempts := 0
	client := &http.Client{Transport: sourceRetryRoundTripFunc(func(request *http.Request) (*http.Response, error) {
		attempts++
		return sourceRetryResponse(request, http.StatusNotFound, io.NopCloser(strings.NewReader("missing"))), nil
	})}
	reader := newRetryTestReader(client)
	reader.retryWait = func(context.Context, time.Duration) error {
		t.Fatal("permanent response waited for retry")
		return nil
	}

	_, _, err := reader.read(context.Background(), "https://example.test/missing", nil)
	if err == nil || err.Error() != "fetch https://example.test/missing: HTTP 404" {
		t.Fatalf("error = %v", err)
	}
	if attempts != 1 {
		t.Fatalf("attempts = %d", attempts)
	}
}

func TestSourceReaderRetriesTransportError(t *testing.T) {
	attempts := 0
	client := &http.Client{Transport: sourceRetryRoundTripFunc(func(request *http.Request) (*http.Response, error) {
		attempts++
		if attempts == 1 {
			return nil, errors.New("connection reset")
		}
		return sourceRetryResponse(request, http.StatusOK, io.NopCloser(strings.NewReader("recovered"))), nil
	})}
	reader := newRetryTestReader(client)
	reader.retryWait = func(context.Context, time.Duration) error { return nil }

	raw, _, err := reader.read(context.Background(), "https://example.test/source", nil)
	if err != nil || string(raw) != "recovered" || attempts != 2 {
		t.Fatalf("read = %q, attempts = %d, error = %v", raw, attempts, err)
	}
}

func TestSourceReaderRetriesTransientBodyErrorsThenSucceeds(t *testing.T) {
	attempts := 0
	var failedBodies []*sourceRetryFailingBody
	completeBody := &sourceRetryBody{Reader: strings.NewReader("complete")}
	client := &http.Client{Transport: sourceRetryRoundTripFunc(func(request *http.Request) (*http.Response, error) {
		attempts++
		if attempts < sourceReadAttempts {
			readErr := error(io.ErrUnexpectedEOF)
			if attempts == 2 {
				readErr = fmt.Errorf("read response: %w", syscall.ECONNRESET)
			}
			body := &sourceRetryFailingBody{content: []byte("partial"), err: readErr}
			failedBodies = append(failedBodies, body)
			return sourceRetryResponse(request, http.StatusOK, body), nil
		}
		return sourceRetryResponse(request, http.StatusOK, completeBody), nil
	})}
	reader := newRetryTestReader(client)
	var delays []time.Duration
	reader.retryWait = func(_ context.Context, delay time.Duration) error {
		delays = append(delays, delay)
		return nil
	}

	raw, _, err := reader.read(context.Background(), "https://example.test/source", nil)
	if err != nil || string(raw) != "complete" {
		t.Fatalf("read = %q, error = %v", raw, err)
	}
	if attempts != sourceReadAttempts || len(delays) != 2 || delays[0] != sourceReadRetryDelay || delays[1] != 2*sourceReadRetryDelay {
		t.Fatalf("attempts = %d, delays = %v", attempts, delays)
	}
	for index, body := range failedBodies {
		if !body.closed {
			t.Errorf("failed body %d was not closed", index)
		}
	}
	if !completeBody.closed {
		t.Error("successful body was not closed")
	}
	if reader.used != int64(len("complete")) {
		t.Fatalf("used bytes = %d", reader.used)
	}
}

func TestSourceReaderReturnsExactErrorAfterTransientBodyRetries(t *testing.T) {
	attempts := 0
	const streamError = "stream error: stream ID 7; INTERNAL_ERROR; received from peer"
	var bodies []*sourceRetryFailingBody
	client := &http.Client{Transport: sourceRetryRoundTripFunc(func(request *http.Request) (*http.Response, error) {
		attempts++
		body := &sourceRetryFailingBody{content: []byte("partial"), err: errors.New(streamError)}
		bodies = append(bodies, body)
		return sourceRetryResponse(request, http.StatusOK, body), nil
	})}
	reader := newRetryTestReader(client)
	reader.retryWait = func(context.Context, time.Duration) error { return nil }

	_, _, err := reader.read(context.Background(), "https://example.test/source", nil)
	if err == nil || err.Error() != "fetch https://example.test/source: "+streamError {
		t.Fatalf("error = %v", err)
	}
	if attempts != sourceReadAttempts || reader.used != 0 {
		t.Fatalf("attempts = %d, used bytes = %d", attempts, reader.used)
	}
	for index, body := range bodies {
		if !body.closed {
			t.Errorf("failed body %d was not closed", index)
		}
	}
}

func TestSourceReaderDoesNotRetryCanceledBodyRead(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	attempts := 0
	body := &sourceRetryFailingBody{content: []byte("partial"), err: context.Canceled}
	client := &http.Client{Transport: sourceRetryRoundTripFunc(func(request *http.Request) (*http.Response, error) {
		attempts++
		cancel()
		return sourceRetryResponse(request, http.StatusOK, body), nil
	})}
	reader := newRetryTestReader(client)
	reader.retryWait = func(context.Context, time.Duration) error {
		t.Fatal("canceled body read waited for retry")
		return nil
	}

	_, _, err := reader.read(ctx, "https://example.test/source", nil)
	if err == nil || !errors.Is(err, context.Canceled) {
		t.Fatalf("error = %v", err)
	}
	if attempts != 1 || reader.used != 0 {
		t.Fatalf("attempts = %d, used bytes = %d", attempts, reader.used)
	}
	if !body.closed {
		t.Error("canceled body was not closed")
	}
}

func TestSourceReaderDoesNotRetrySourceTooLarge(t *testing.T) {
	attempts := 0
	body := &sourceRetryBody{Reader: strings.NewReader("oversized")}
	client := &http.Client{Transport: sourceRetryRoundTripFunc(func(request *http.Request) (*http.Response, error) {
		attempts++
		return sourceRetryResponse(request, http.StatusOK, body), nil
	})}
	reader := &sourceReader{client: client, perSource: 4, total: DefaultMaxTotalBytes}
	reader.retryWait = func(context.Context, time.Duration) error {
		t.Fatal("size-limit error waited for retry")
		return nil
	}

	_, _, err := reader.read(context.Background(), "https://example.test/source", nil)
	if err == nil || err.Error() != "fetch https://example.test/source: source exceeds 4 bytes" || !errors.Is(err, ErrSourceTooLarge) {
		t.Fatalf("error = %v", err)
	}
	if attempts != 1 || reader.used != 0 {
		t.Fatalf("attempts = %d, used bytes = %d", attempts, reader.used)
	}
	if !body.closed {
		t.Error("oversized body was not closed")
	}
}
