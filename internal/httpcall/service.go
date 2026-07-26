package httpcall

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"github.com/sairaph/apis-mcp/internal/cache"
	"github.com/sairaph/apis-mcp/internal/sessions"
)

func New(cfg Config) (*Service, error) {
	if cfg.Cache == nil || cfg.Sessions == nil {
		return nil, errors.New("cache and sessions services are required")
	}
	if cfg.MaximumHeaderTimeout == 0 {
		cfg.MaximumHeaderTimeout = 600
	}
	if cfg.ReadTokenBudget == 0 {
		cfg.ReadTokenBudget = 4_000
	}
	if cfg.StalledDownloadAfter == 0 {
		cfg.StalledDownloadAfter = time.Hour
	}
	if cfg.MaximumHeaderTimeout < 30 || cfg.MaximumHeaderTimeout > 600 || cfg.MaximumRetries < 0 || cfg.MaximumRetries > 30 || cfg.MaximumRedirects < 0 || cfg.MaximumRedirects > 20 || cfg.StalledDownloadAfter < 0 || cfg.ReadTokenBudget < 1 {
		return nil, errors.New("invalid HTTP call service limits")
	}
	if cfg.Sleep == nil {
		cfg.Sleep = func(ctx context.Context, delay time.Duration) error {
			timer := time.NewTimer(delay)
			defer timer.Stop()
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-timer.C:
				return nil
			}
		}
	}
	if cfg.Now == nil {
		cfg.Now = time.Now
	}
	transport := cfg.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}
	if base, ok := transport.(*http.Transport); ok {
		clone := base.Clone()
		clone.DisableCompression = true
		transport = clone
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &Service{
		cache: cfg.Cache, sessions: cfg.Sessions, transport: transport,
		maxTimeout: cfg.MaximumHeaderTimeout, maxRetries: cfg.MaximumRetries,
		redirects: cfg.MaximumRedirects, background: cfg.BackgroundAfter, stalled: cfg.StalledDownloadAfter,
		readBudget: cfg.ReadTokenBudget, sleep: cfg.Sleep, now: cfg.Now, ctx: ctx, cancel: cancel,
	}, nil
}

// Shutdown cancels process-owned downloads and waits for their bounded cleanup.
func (s *Service) Shutdown(ctx context.Context) error {
	s.lifeMu.Lock()
	if !s.closing {
		s.closing = true
		s.cancel()
	}
	s.lifeMu.Unlock()
	done := make(chan struct{})
	go func() { s.wg.Wait(); close(done) }()
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (s *Service) Call(ctx context.Context, input Input) (Result, error) {
	s.lifeMu.Lock()
	if s.closing {
		s.lifeMu.Unlock()
		return Result{}, &Error{Code: "shutting_down", Message: "HTTP call service is shutting down"}
	}
	s.wg.Add(1)
	s.lifeMu.Unlock()
	defer s.wg.Done()

	prepared, err := s.prepare(input)
	if err != nil {
		return Result{}, err
	}

	var session *sessions.Handle
	if prepared.session == "" {
		session, err = s.sessions.Create(ctx)
	} else {
		session, err = s.sessions.Acquire(ctx, prepared.session)
	}
	if err != nil {
		return Result{}, &Error{Code: "session_error", Message: "open cookie session", Cause: err}
	}
	owned := true
	defer func() {
		if owned {
			_ = session.Close()
		}
	}()

	callID, err := uuid.NewV7()
	if err != nil {
		return Result{}, &Error{Code: "internal_error", Message: "create call ID", Cause: err}
	}
	started := s.now()
	result := Result{Request: RequestResult{ID: callID.String(), SessionID: session.ID(), Method: prepared.method, Endpoint: prepared.endpoint.String(), AutomaticHeaders: prepared.automatic}}

	response, cancelRequest, attempts, redirects, err := s.do(ctx, prepared, session)
	result.Attempts, result.Redirects = attempts, redirects
	if err != nil {
		return result, err
	}
	if response == nil {
		return result, &Error{Code: "transport_error", Message: "request completed without a response"}
	}

	contentType := response.Header.Get("Content-Type")
	encoding := response.Header.Get("Content-Encoding")
	decoded := supportedEncoding(encoding)
	extension := bodyExtension(contentType, decoded)
	if !prepared.allowLarge && decoded && strings.TrimSpace(encoding) == "" && response.ContentLength > s.cache.MaxDecodedBytes() {
		response.Body.Close()
		cancelRequest()
		return result, &Error{Code: "response_too_large", Message: "declared response size exceeds the configured decoded-size limit", Hint: "retry with allow_large_download when policy permits", Cause: cache.ErrSizeLimit, Fields: map[string]any{"content_length": response.ContentLength, "limit": s.cache.MaxDecodedBytes()}}
	}
	reserveSize := int64(0)
	if decoded && strings.TrimSpace(encoding) == "" && response.ContentLength > 0 {
		reserveSize = response.ContentLength
	}
	if err := s.cache.CheckSpace(reserveSize); err != nil {
		response.Body.Close()
		cancelRequest()
		return result, cacheError("reserve response cache space", err)
	}
	entry, err := s.cache.Begin(callID.String(), extension)
	if err != nil {
		response.Body.Close()
		cancelRequest()
		return result, cacheError("prepare response cache", err)
	}
	if err := entry.SaveHeaders(map[string][]string(response.Header.Clone())); err != nil {
		response.Body.Close()
		cancelRequest()
		entry.Abort()
		return result, cacheError("save response headers", err)
	}
	body, err := entry.CreateBody()
	if err != nil {
		response.Body.Close()
		cancelRequest()
		entry.Abort()
		return result, cacheError("create cached response body", err)
	}

	result.Response = ResponseResult{
		State: "downloading", Status: response.StatusCode, StatusText: response.Status,
		FinalURL: response.Request.URL.String(), ContentType: contentType, ContentEncoding: encoding,
		Headers: displayHeaders(response.Header), Decoded: decoded, Duration: s.now().Sub(started),
	}
	if response.ContentLength > 0 {
		result.Response.DeclaredBytes = response.ContentLength
	}
	result.Cache = CacheResult{Directory: entry.Directory, TempPath: entry.TempPath, FinalPath: entry.BodyPath, HeadersPath: entry.HeadersPath, MetadataPath: entry.MetadataPath, ErrorPath: entry.ErrorPath}
	metadata := s.metadata(result, started)
	if err := entry.SaveMetadata(metadata); err != nil {
		body.Close()
		response.Body.Close()
		cancelRequest()
		entry.Abort()
		return result, cacheError("save response metadata", err)
	}

	outcomes := make(chan streamOutcome, 1)
	progress := &streamProgress{}
	s.wg.Add(1)
	go func(base Result) {
		defer s.wg.Done()
		outcome := s.stream(response.Body, body, entry, prepared.allowLarge, decoded, encoding, progress)
		outcome.completed = s.now().UTC()
		outcome.duration = s.now().Sub(started)
		if outcome.err == nil {
			complete := base
			applyOutcome(&complete, outcome)
			m := s.metadata(complete, started)
			// The atomically published body is authoritative. A metadata refresh
			// failure must not turn that successful publication into an error.
			_ = entry.SaveMetadata(m)
		}
		if outcome.err != nil {
			failed := base
			failed.Response.State = "failed"
			failed.Response.WireBytes = outcome.wireBytes
			failed.Response.DecodedBytes = outcome.decodedBytes
			failed.Response.Duration = outcome.duration
			failed.Response.CompletedAt = &outcome.completed
			_ = entry.SaveMetadata(s.metadata(failed, started))
			_ = entry.PublishError(cache.ErrorRecord{Code: streamErrorCode(outcome.err), Message: outcome.err.Error(), FailedAt: outcome.completed, BytesRead: outcome.decodedBytes})
		}
		entry.Close()
		cancelRequest()
		if closeErr := session.Close(); closeErr != nil {
			outcome.sessionErr = closeErr
		}
		outcomes <- outcome
	}(result)
	owned = false

	if s.background <= 0 {
		select {
		case outcome := <-outcomes:
			return s.finish(result, prepared.jsonPath, outcome)
		case <-ctx.Done():
			cancelRequest()
			outcome := <-outcomes
			_ = outcome
			return result, ctx.Err()
		}
	}
	timer := time.NewTimer(s.background)
	defer timer.Stop()
	select {
	case outcome := <-outcomes:
		return s.finish(result, prepared.jsonPath, outcome)
	case <-timer.C:
		result.Response.WireBytes = progress.wire.Load()
		result.Response.DecodedBytes = progress.decoded.Load()
		result.Response.Duration = s.now().Sub(started)
		if seconds := result.Response.Duration.Seconds(); seconds > 0 {
			result.Response.BytesPerSecond = float64(result.Response.DecodedBytes) / seconds
		}
		if result.Response.DeclaredBytes > result.Response.WireBytes && result.Response.BytesPerSecond > 0 && strings.TrimSpace(encoding) == "" {
			remaining := time.Duration(float64(result.Response.DeclaredBytes-result.Response.WireBytes) / result.Response.BytesPerSecond * float64(time.Second))
			result.Response.EstimatedRemaining = &remaining
		}
		if prepared.jsonPath != "" {
			result.Selection = &Selection{JSONPath: prepared.jsonPath, State: "skipped", Error: "selection skipped while the response continues downloading"}
		}
		return result, nil
	case <-ctx.Done():
		cancelRequest()
		outcome := <-outcomes
		_ = outcome
		return result, ctx.Err()
	}
}

func (s *Service) do(ctx context.Context, input preparedInput, session *sessions.Handle) (*http.Response, context.CancelFunc, []Attempt, []Redirect, error) {
	var attempts []Attempt
	var redirects []Redirect
	for number := 1; number <= input.retries+1; number++ {
		if err := s.operationErr(ctx); err != nil {
			return nil, nil, attempts, redirects, err
		}
		requestContext, cancel := context.WithCancel(s.ctx)
		headersDone := make(chan struct{})
		go func() {
			select {
			case <-ctx.Done():
				cancel()
			case <-headersDone:
			case <-requestContext.Done():
			}
		}()
		var body io.Reader
		if input.payload != nil {
			body = bytes.NewReader(input.payload)
		}
		request, err := http.NewRequestWithContext(requestContext, input.method, input.endpoint.String(), body)
		if err != nil {
			close(headersDone)
			cancel()
			return nil, nil, attempts, redirects, validation("request", err.Error())
		}
		request.Header = input.headers.Clone()
		if host := request.Header.Get("Host"); host != "" {
			request.Host = host
			request.Header.Del("Host")
		}
		jar := overlayJar{base: session.Jar(), host: input.endpoint.Hostname(), excluded: input.explicitCookies}
		client := &http.Client{Transport: s.transport, Jar: jar, CheckRedirect: s.redirectPolicy(&redirects)}
		var timedOut atomic.Bool
		timer := time.AfterFunc(time.Duration(input.timeout)*time.Second, func() { timedOut.Store(true); cancel() })
		response, requestErr := client.Do(request)
		if !timer.Stop() && response != nil {
			timedOut.Store(true)
		}
		close(headersDone)
		if requestErr != nil {
			cancel()
			attempt := Attempt{Number: number, Error: requestErr.Error()}
			if number <= input.retries && (timedOut.Load() || transientTransport(requestErr)) && s.operationErr(ctx) == nil {
				delay := retryDelay(number)
				attempt.RetryReason, attempt.RetryDelay = transportReason(timedOut.Load()), delay
				attempts = append(attempts, attempt)
				if err := s.sleepRetry(ctx, delay); err != nil {
					return nil, nil, attempts, redirects, err
				}
				continue
			}
			attempts = append(attempts, attempt)
			if err := s.operationErr(ctx); err != nil {
				return nil, nil, attempts, redirects, err
			}
			if timedOut.Load() {
				return nil, nil, attempts, redirects, &Error{Code: "response_header_timeout", Message: fmt.Sprintf("response headers were not received within %d seconds", input.timeout), Hint: "retry with a larger timeout", Cause: requestErr, Fields: map[string]any{"attempts": number, "timeout": input.timeout}}
			}
			return nil, nil, attempts, redirects, &Error{Code: "transport_error", Message: "HTTP request failed before a response was received", Hint: "change retries to control automatic retry attempts", Cause: requestErr, Fields: map[string]any{"attempts": number}}
		}
		attempt := Attempt{Number: number, Status: response.StatusCode}
		if transientStatus(response.StatusCode) && number <= input.retries {
			delay, tooLong := retryAfter(response.Header.Get("Retry-After"), s.now(), retryDelay(number))
			if !tooLong {
				attempt.RetryReason = fmt.Sprintf("transient HTTP status %d", response.StatusCode)
				attempt.RetryDelay = delay
				attempts = append(attempts, attempt)
				response.Body.Close()
				cancel()
				if err := s.sleepRetry(ctx, delay); err != nil {
					return nil, nil, attempts, redirects, err
				}
				continue
			}
		}
		attempts = append(attempts, attempt)
		return response, cancel, attempts, redirects, nil
	}
	return nil, nil, attempts, redirects, &Error{Code: "transport_error", Message: "HTTP attempts exhausted"}
}

func (s *Service) operationErr(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return s.ctx.Err()
}

func (s *Service) sleepRetry(ctx context.Context, delay time.Duration) error {
	if err := s.operationErr(ctx); err != nil {
		return err
	}
	sleepCtx, cancel := context.WithCancel(ctx)
	stop := context.AfterFunc(s.ctx, cancel)
	defer func() {
		stop()
		cancel()
	}()
	if err := s.sleep(sleepCtx, delay); err != nil {
		if operationErr := s.operationErr(ctx); operationErr != nil {
			return operationErr
		}
		return err
	}
	return s.operationErr(ctx)
}

var errRedirectPolicy = errors.New("redirect rejected by policy")

func (s *Service) redirectPolicy(history *[]Redirect) func(*http.Request, []*http.Request) error {
	return func(next *http.Request, via []*http.Request) error {
		if s.redirects == 0 || len(via) > s.redirects {
			return http.ErrUseLastResponse
		}
		if _, err := parseEndpoint(next.URL.String()); err != nil {
			return fmt.Errorf("%w: %v", errRedirectPolicy, err)
		}
		previous := via[len(via)-1]
		status := 0
		if next.Response != nil {
			status = next.Response.StatusCode
		}
		*history = append(*history, Redirect{FromURL: previous.URL.String(), ToURL: next.URL.String(), Status: status, MethodBefore: previous.Method, MethodAfter: next.Method, BodyRetained: next.Body != nil})
		return nil
	}
}

type overlayJar struct {
	base     http.CookieJar
	host     string
	excluded map[string]struct{}
}

func (j overlayJar) SetCookies(u *url.URL, cookies []*http.Cookie) { j.base.SetCookies(u, cookies) }
func (j overlayJar) Cookies(u *url.URL) []*http.Cookie {
	cookies := j.base.Cookies(u)
	if !strings.EqualFold(u.Hostname(), j.host) || len(j.excluded) == 0 {
		return cookies
	}
	result := cookies[:0]
	for _, cookie := range cookies {
		if _, exists := j.excluded[cookie.Name]; !exists {
			result = append(result, cookie)
		}
	}
	return result
}

func retryDelay(retry int) time.Duration {
	switch retry {
	case 1:
		return 5 * time.Second
	case 2:
		return 15 * time.Second
	case 3:
		return 30 * time.Second
	case 4:
		return 60 * time.Second
	default:
		return 120 * time.Second
	}
}

func retryAfter(value string, now time.Time, fallback time.Duration) (time.Duration, bool) {
	if value == "" {
		return fallback, false
	}
	var delay time.Duration
	if seconds, err := strconv.Atoi(strings.TrimSpace(value)); err == nil && seconds >= 0 {
		delay = time.Duration(seconds) * time.Second
	} else if at, err := http.ParseTime(value); err == nil {
		delay = at.Sub(now)
		if delay < 0 {
			delay = 0
		}
	} else {
		return fallback, false
	}
	return delay, delay > 120*time.Second
}

func transientStatus(status int) bool {
	switch status {
	case 408, 425, 429, 500, 502, 503, 504:
		return true
	default:
		return false
	}
}

func transientTransport(err error) bool {
	if errors.Is(err, errRedirectPolicy) {
		return false
	}
	var verification *tls.CertificateVerificationError
	if errors.As(err, &verification) {
		return false
	}
	var dns *net.DNSError
	if errors.As(err, &dns) {
		return dns.IsTimeout || dns.IsTemporary
	}
	var unknown x509.UnknownAuthorityError
	if errors.As(err, &unknown) {
		return false
	}
	var hostname x509.HostnameError
	if errors.As(err, &hostname) {
		return false
	}
	var invalid x509.CertificateInvalidError
	if errors.As(err, &invalid) {
		return false
	}
	var roots x509.SystemRootsError
	if errors.As(err, &roots) {
		return false
	}
	var record tls.RecordHeaderError
	if errors.As(err, &record) {
		return false
	}
	var network net.Error
	if errors.As(err, &network) {
		return network.Timeout() || network.Temporary() || !errors.As(err, &dns)
	}
	return errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF)
}

func transportReason(timeout bool) string {
	if timeout {
		return "response-header timeout"
	}
	return "transient transport failure"
}

func cacheError(message string, err error) *Error {
	return &Error{Code: "cache_error", Message: message, Cause: err}
}

func displayHeaders(headers http.Header) map[string][]string {
	allowed := map[string]bool{"Content-Type": true, "Content-Length": true, "Content-Encoding": true, "Content-Disposition": true, "Location": true, "Link": true, "Retry-After": true, "ETag": true, "Last-Modified": true, "X-Request-Id": true, "X-Trace-Id": true, "RateLimit-Limit": true, "RateLimit-Remaining": true, "RateLimit-Reset": true}
	result := make(map[string][]string)
	for name, values := range headers {
		canonical := http.CanonicalHeaderKey(name)
		if !allowed[canonical] && !strings.HasPrefix(canonical, "X-RateLimit-") {
			continue
		}
		length := 0
		for _, value := range values {
			length += len(value)
		}
		if length <= 1024 {
			result[canonical] = append([]string(nil), values...)
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}
