package httpcall

import (
	"bufio"
	"compress/flate"
	"compress/gzip"
	"compress/zlib"
	"errors"
	"fmt"
	"io"
	"mime"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/sairaph/apis-mcp/internal/cache"
)

type streamOutcome struct {
	wireBytes    int64
	decodedBytes int64
	completed    time.Time
	duration     time.Duration
	err          error
	sessionErr   error
}

type streamProgress struct {
	wire    atomic.Int64
	decoded atomic.Int64
}

type countingReader struct {
	reader   io.Reader
	count    int64
	progress *atomic.Int64
}

var errStalledDownload = errors.New("response body download timed out while stalled")

type stalledBodyReader struct {
	source     io.ReadCloser
	timeout    time.Duration
	timeoutErr error
	mu         sync.Mutex
	timer      *time.Timer
	generation uint64
	done       bool
	timedOut   bool
	closeOnce  sync.Once
	closeErr   error
}

func newStalledBodyReader(source io.ReadCloser, timeout time.Duration) *stalledBodyReader {
	r := &stalledBodyReader{source: source, timeout: timeout, timeoutErr: fmt.Errorf("%w: no body bytes received for %s", errStalledDownload, timeout)}
	r.mu.Lock()
	r.armLocked()
	r.mu.Unlock()
	return r
}

func (r *stalledBodyReader) Read(p []byte) (int, error) {
	n, err := r.source.Read(p)
	r.mu.Lock()
	switch {
	case r.timedOut:
		err = r.timeoutErr
	case err != nil:
		r.stopLocked()
	case n > 0:
		r.armLocked()
	}
	r.mu.Unlock()
	return n, err
}

func (r *stalledBodyReader) Close() error {
	r.mu.Lock()
	r.stopLocked()
	r.mu.Unlock()
	return r.closeSource()
}

func (r *stalledBodyReader) armLocked() {
	if r.done {
		return
	}
	if r.timer != nil {
		r.timer.Stop()
	}
	r.generation++
	generation := r.generation
	r.timer = time.AfterFunc(r.timeout, func() {
		r.mu.Lock()
		if r.done || r.generation != generation {
			r.mu.Unlock()
			return
		}
		r.done = true
		r.timedOut = true
		r.mu.Unlock()
		_ = r.closeSource()
	})
}

func (r *stalledBodyReader) stopLocked() {
	if r.done {
		return
	}
	r.done = true
	r.generation++
	if r.timer != nil {
		r.timer.Stop()
	}
}

func (r *stalledBodyReader) closeSource() error {
	r.closeOnce.Do(func() { r.closeErr = r.source.Close() })
	return r.closeErr
}

func (r *countingReader) Read(p []byte) (int, error) {
	n, err := r.reader.Read(p)
	r.count += int64(n)
	if r.progress != nil {
		r.progress.Add(int64(n))
	}
	return n, err
}

func (s *Service) stream(source io.ReadCloser, destination *os.File, entry *cache.Entry, allowLarge, decoded bool, encoding string, progress *streamProgress) streamOutcome {
	body := newStalledBodyReader(source, s.stalled)
	defer body.Close()
	counter := &countingReader{reader: body, progress: &progress.wire}
	reader := io.Reader(counter)
	var closers []io.Closer
	if decoded {
		var err error
		reader, closers, err = decodingReader(reader, encoding)
		if err != nil {
			destination.Close()
			return streamOutcome{wireBytes: counter.count, err: err}
		}
	}
	for i := len(closers) - 1; i >= 0; i-- {
		defer closers[i].Close()
	}
	limited := s.cache.Writer(destination, allowLarge)
	writer := progressWriter{writer: limited, count: &progress.decoded}
	_, err := io.CopyBuffer(writer, reader, make([]byte, 64<<10))
	if err != nil {
		destination.Close()
		return streamOutcome{wireBytes: counter.count, decodedBytes: limited.Written(), err: err}
	}
	if err := entry.PublishBody(destination); err != nil {
		return streamOutcome{wireBytes: counter.count, decodedBytes: limited.Written(), err: err}
	}
	return streamOutcome{wireBytes: counter.count, decodedBytes: limited.Written()}
}

type progressWriter struct {
	writer io.Writer
	count  *atomic.Int64
}

func (w progressWriter) Write(p []byte) (int, error) {
	n, err := w.writer.Write(p)
	w.count.Add(int64(n))
	return n, err
}

func supportedEncoding(value string) bool {
	for _, encoding := range encodings(value) {
		switch encoding {
		case "identity", "gzip", "deflate":
		default:
			return false
		}
	}
	return true
}

func encodings(value string) []string {
	if strings.TrimSpace(value) == "" {
		return nil
	}
	parts := strings.Split(value, ",")
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		if item := strings.ToLower(strings.TrimSpace(part)); item != "" {
			result = append(result, item)
		}
	}
	return result
}

func decodingReader(source io.Reader, value string) (io.Reader, []io.Closer, error) {
	reader := source
	var closers []io.Closer
	list := encodings(value)
	for i := len(list) - 1; i >= 0; i-- {
		switch list[i] {
		case "identity":
		case "gzip":
			decoder, err := gzip.NewReader(reader)
			if err != nil {
				return nil, closers, fmt.Errorf("decode gzip response: %w", err)
			}
			reader = decoder
			closers = append(closers, decoder)
		case "deflate":
			buffered := bufio.NewReader(reader)
			header, err := buffered.Peek(2)
			if err != nil {
				return nil, closers, fmt.Errorf("decode deflate response: %w", err)
			}
			if header[0]&0x0f == 8 && ((int(header[0])<<8)+int(header[1]))%31 == 0 {
				decoder, err := zlib.NewReader(buffered)
				if err != nil {
					return nil, closers, fmt.Errorf("decode deflate response: %w", err)
				}
				reader = decoder
				closers = append(closers, decoder)
			} else {
				decoder := flate.NewReader(buffered)
				reader = decoder
				closers = append(closers, decoder)
			}
		}
	}
	return reader, closers, nil
}

func bodyExtension(contentType string, decoded bool) string {
	if !decoded {
		return ".encoded"
	}
	mediaType, _, _ := mime.ParseMediaType(contentType)
	switch strings.ToLower(mediaType) {
	case "application/json", "application/problem+json":
		return ".json"
	case "text/html":
		return ".html"
	case "text/plain":
		return ".txt"
	case "text/csv":
		return ".csv"
	case "application/xml", "text/xml":
		return ".xml"
	case "application/pdf":
		return ".pdf"
	default:
		if strings.HasSuffix(strings.ToLower(mediaType), "+json") {
			return ".json"
		}
		if strings.HasPrefix(strings.ToLower(mediaType), "text/") {
			return ".txt"
		}
		return ".bin"
	}
}

func streamErrorCode(err error) string {
	if errors.Is(err, errStalledDownload) {
		return "download_timeout"
	}
	if errors.Is(err, cache.ErrSizeLimit) {
		return "response_too_large"
	}
	if errors.Is(err, cache.ErrDiskReserve) {
		return "disk_reserve"
	}
	return "download_error"
}

func (s *Service) metadata(result Result, started time.Time) cache.Metadata {
	return cache.Metadata{ID: result.Request.ID, State: result.Response.State, StartedAt: started.UTC(), CompletedAt: result.Response.CompletedAt, Method: result.Request.Method, Endpoint: result.Request.Endpoint, FinalURL: result.Response.FinalURL, Status: result.Response.Status, ContentType: result.Response.ContentType, ContentEncoding: result.Response.ContentEncoding, Decoded: result.Response.Decoded, WireBytes: result.Response.WireBytes, DecodedBytes: result.Response.DecodedBytes, SessionID: result.Request.SessionID, Attempts: result.Attempts, Redirects: result.Redirects}
}

func applyOutcome(result *Result, outcome streamOutcome) {
	result.Response.State = "complete"
	result.Response.WireBytes = outcome.wireBytes
	result.Response.DecodedBytes = outcome.decodedBytes
	result.Response.Duration = outcome.duration
	result.Response.CompletedAt = &outcome.completed
	result.Cache.BodyPath = result.Cache.FinalPath
	result.Cache.TempPath = ""
	result.Cache.FinalPath = ""
	result.Cache.ErrorPath = ""
}

func (s *Service) finish(result Result, jsonPath string, outcome streamOutcome) (Result, error) {
	if outcome.err != nil {
		if errors.Is(outcome.err, cache.ErrSizeLimit) {
			return result, &Error{
				Code: "response_too_large", Message: "decoded response exceeds the configured size limit",
				Hint:   "retry with allow_large_download when available, or increase the configured response size limit",
				Fields: map[string]any{"limit": s.cache.MaxDecodedBytes()}, Cause: outcome.err,
			}
		}
		return result, cacheError("cache response body", outcome.err)
	}
	applyOutcome(&result, outcome)
	if outcome.sessionErr != nil {
		return result, &Error{Code: "session_error", Message: "save cookie session", Cause: outcome.sessionErr}
	}
	preview, selection, err := makePreview(result.Cache.BodyPath, result.Response.ContentType, result.Response.Decoded, s.readBudget, jsonPath, s.cache.MaxDecodedBytes())
	if err != nil {
		return result, fmt.Errorf("construct response preview: %w", err)
	}
	result.Preview, result.Selection = preview, selection
	return result, nil
}
