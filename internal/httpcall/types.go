// Package httpcall executes and persistently caches transport-neutral HTTP calls.
package httpcall

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/sairaph/apis-mcp/internal/app"
	"github.com/sairaph/apis-mcp/internal/cache"
	"github.com/sairaph/apis-mcp/internal/sessions"
)

type Input = app.CallInput

type Sleeper func(context.Context, time.Duration) error

type Config struct {
	Cache                *cache.Store
	Sessions             *sessions.Manager
	Transport            http.RoundTripper
	MaximumHeaderTimeout int
	MaximumRetries       int
	MaximumRedirects     int
	BackgroundAfter      time.Duration
	StalledDownloadAfter time.Duration
	ReadTokenBudget      int
	Sleep                Sleeper
	Now                  func() time.Time
}

func DefaultConfig(cacheStore *cache.Store, sessionManager *sessions.Manager) Config {
	return Config{Cache: cacheStore, Sessions: sessionManager, MaximumHeaderTimeout: 600, MaximumRetries: 30, MaximumRedirects: 5, BackgroundAfter: 60 * time.Second, StalledDownloadAfter: time.Hour, ReadTokenBudget: 4_000}
}

type Service struct {
	cache      *cache.Store
	sessions   *sessions.Manager
	transport  http.RoundTripper
	maxTimeout int
	maxRetries int
	redirects  int
	background time.Duration
	stalled    time.Duration
	readBudget int
	sleep      Sleeper
	now        func() time.Time

	ctx     context.Context
	cancel  context.CancelFunc
	wg      sync.WaitGroup
	lifeMu  sync.Mutex
	closing bool
}

type Result struct {
	Request   RequestResult  `json:"request"`
	Response  ResponseResult `json:"response"`
	Cache     CacheResult    `json:"cache"`
	Preview   *Preview       `json:"preview,omitempty"`
	Selection *Selection     `json:"selection,omitempty"`
	Attempts  []Attempt      `json:"attempts"`
	Redirects []Redirect     `json:"redirects,omitempty"`
}

type RequestResult struct {
	ID               string            `json:"id"`
	SessionID        string            `json:"session_id"`
	Method           string            `json:"method"`
	Endpoint         string            `json:"endpoint"`
	AutomaticHeaders map[string]string `json:"automatic_headers,omitempty"`
}

type ResponseResult struct {
	State              string              `json:"state"`
	Status             int                 `json:"status"`
	StatusText         string              `json:"status_text"`
	FinalURL           string              `json:"final_url"`
	ContentType        string              `json:"content_type,omitempty"`
	ContentEncoding    string              `json:"content_encoding,omitempty"`
	Headers            map[string][]string `json:"headers,omitempty"`
	Decoded            bool                `json:"decoded"`
	WireBytes          int64               `json:"wire_bytes"`
	DecodedBytes       int64               `json:"decoded_bytes"`
	DeclaredBytes      int64               `json:"declared_bytes,omitempty"`
	BytesPerSecond     float64             `json:"bytes_per_second,omitempty"`
	EstimatedRemaining *time.Duration      `json:"estimated_remaining,omitempty"`
	Duration           time.Duration       `json:"duration"`
	CompletedAt        *time.Time          `json:"completed_at,omitempty"`
}

type CacheResult struct {
	Directory    string `json:"directory"`
	BodyPath     string `json:"body_path,omitempty"`
	TempPath     string `json:"temp_path,omitempty"`
	FinalPath    string `json:"final_path,omitempty"`
	HeadersPath  string `json:"headers_path"`
	MetadataPath string `json:"metadata_path"`
	ErrorPath    string `json:"error_path,omitempty"`
}

type Preview struct {
	Kind              string `json:"kind"`
	Language          string `json:"language,omitempty"`
	Content           string `json:"content"`
	Truncated         bool   `json:"truncated"`
	ApproximateTokens int    `json:"approximate_tokens"`
}

type Selection struct {
	JSONPath string `json:"json_path"`
	State    string `json:"state,omitempty"`
	Matched  bool   `json:"matched"`
	Error    string `json:"selection_error,omitempty"`
}

type Attempt struct {
	Number      int           `json:"number"`
	Status      int           `json:"status,omitempty"`
	Error       string        `json:"error,omitempty"`
	RetryReason string        `json:"retry_reason,omitempty"`
	RetryDelay  time.Duration `json:"retry_delay,omitempty"`
}

type Redirect struct {
	FromURL      string `json:"from_url"`
	ToURL        string `json:"to_url"`
	Status       int    `json:"status"`
	MethodBefore string `json:"method_before"`
	MethodAfter  string `json:"method_after"`
	BodyRetained bool   `json:"body_retained"`
}

type Error struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Hint    string         `json:"hint,omitempty"`
	Fields  map[string]any `json:"fields,omitempty"`
	Cause   error          `json:"-"`
}

func (e *Error) Error() string {
	if e.Cause != nil {
		return e.Message + ": " + e.Cause.Error()
	}
	return e.Message
}

func (e *Error) Unwrap() error { return e.Cause }
