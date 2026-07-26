// Package cache provides private, atomically published HTTP response storage.
package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/sairaph/apis-mcp/internal/fsx"

	"github.com/gofrs/flock"
)

const DefaultMaxDecodedBytes int64 = 50 << 20

var ErrSizeLimit = errors.New("decoded response exceeds the configured size limit")
var ErrDiskReserve = errors.New("response would cross the configured free disk reserve")

type Config struct {
	Root            string
	MaxDecodedBytes int64
	FreeDiskReserve int64
	Retention       time.Duration
	Now             func() time.Time
	Faults          *Faults
	FreeBytes       func(string) (uint64, error)
}

// Faults provides deterministic storage failure injection for package tests.
type Faults struct {
	SaveMetadata func(Metadata) error
	PublishBody  func() error
}

type Store struct {
	root      string
	lockRoot  string
	max       int64
	reserve   int64
	retention time.Duration
	now       func() time.Time
	faults    *Faults
	freeBytes func(string) (uint64, error)

	mu     sync.Mutex
	active map[string]struct{}
}

type Entry struct {
	ID           string `json:"id"`
	Directory    string `json:"directory"`
	TempPath     string `json:"temp_path"`
	BodyPath     string `json:"body_path"`
	HeadersPath  string `json:"headers_path"`
	MetadataPath string `json:"metadata_path"`
	ErrorPath    string `json:"error_path"`

	store *Store
	lease *entryLease
	once  sync.Once
}

type entryLease struct {
	store *Store
	id    string
	file  *flock.Flock
}

type Metadata struct {
	ID              string         `json:"id"`
	State           string         `json:"state"`
	StartedAt       time.Time      `json:"started_at"`
	CompletedAt     *time.Time     `json:"completed_at,omitempty"`
	Method          string         `json:"method,omitempty"`
	Endpoint        string         `json:"endpoint,omitempty"`
	FinalURL        string         `json:"final_url,omitempty"`
	Status          int            `json:"status,omitempty"`
	ContentType     string         `json:"content_type,omitempty"`
	ContentEncoding string         `json:"content_encoding,omitempty"`
	Decoded         bool           `json:"decoded"`
	WireBytes       int64          `json:"wire_bytes"`
	DecodedBytes    int64          `json:"decoded_bytes"`
	SessionID       string         `json:"session_id,omitempty"`
	Attempts        any            `json:"attempts,omitempty"`
	Redirects       any            `json:"redirects,omitempty"`
	Extra           map[string]any `json:"extra,omitempty"`
}

type ErrorRecord struct {
	Code      string    `json:"code"`
	Message   string    `json:"message"`
	FailedAt  time.Time `json:"failed_at"`
	BytesRead int64     `json:"bytes_received"`
}

type CleanupResult struct {
	RemovedEntries int
	RemovedOrphans int
}

func New(cfg Config) (*Store, error) {
	if cfg.Root == "" {
		return nil, errors.New("cache root is required")
	}
	if cfg.MaxDecodedBytes == 0 {
		cfg.MaxDecodedBytes = DefaultMaxDecodedBytes
	}
	if cfg.MaxDecodedBytes < 0 || cfg.FreeDiskReserve < 0 || cfg.Retention < 0 {
		return nil, errors.New("cache sizes and retention cannot be negative")
	}
	if cfg.Now == nil {
		cfg.Now = time.Now
	}
	if cfg.FreeBytes == nil {
		cfg.FreeBytes = freeBytes
	}
	if err := os.MkdirAll(cfg.Root, 0o700); err != nil {
		return nil, fmt.Errorf("create call cache: %w", err)
	}
	if err := os.Chmod(cfg.Root, 0o700); err != nil {
		return nil, fmt.Errorf("secure call cache: %w", err)
	}
	lockRoot := filepath.Join(cfg.Root, ".locks")
	if err := os.MkdirAll(lockRoot, 0o700); err != nil {
		return nil, fmt.Errorf("create cache lock directory: %w", err)
	}
	if err := os.Chmod(lockRoot, 0o700); err != nil {
		return nil, fmt.Errorf("secure cache lock directory: %w", err)
	}
	return &Store{root: cfg.Root, lockRoot: lockRoot, max: cfg.MaxDecodedBytes, reserve: cfg.FreeDiskReserve, retention: cfg.Retention, now: cfg.Now, faults: cfg.Faults, freeBytes: cfg.FreeBytes, active: make(map[string]struct{})}, nil
}

func (s *Store) Root() string           { return s.root }
func (s *Store) MaxDecodedBytes() int64 { return s.max }

// CheckSpace verifies the configured reserve before a response body is opened.
// A zero expected size checks only the current free-space floor.
func (s *Store) CheckSpace(expected int64) error {
	if s.reserve == 0 {
		return nil
	}
	ok, err := s.hasSpace(expected)
	if err != nil || ok {
		return err
	}
	// Expired entries and abandoned downloads may make the request feasible.
	// Cleanup is best effort; the authoritative second space check decides.
	_, _ = s.Cleanup()
	ok, err = s.hasSpace(expected)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}
	return ErrDiskReserve
}

func (s *Store) hasSpace(expected int64) (bool, error) {
	free, err := s.freeBytes(s.root)
	if err != nil {
		return false, fmt.Errorf("check free cache space: %w", err)
	}
	return free > uint64(s.reserve) && (expected <= 0 || uint64(expected) <= free-uint64(s.reserve)), nil
}

func (s *Store) Begin(id, extension string) (*Entry, error) {
	if id == "" || strings.ContainsAny(id, `/\\`) || id == "." || id == ".." {
		return nil, errors.New("invalid cache entry ID")
	}
	if extension == "" || extension[0] != '.' || strings.ContainsAny(extension, `/\\`) {
		return nil, errors.New("invalid cache body extension")
	}
	lease, acquired, err := s.tryAcquire(id)
	if err != nil {
		return nil, fmt.Errorf("lock cache entry: %w", err)
	}
	if !acquired {
		return nil, errors.New("cache entry is already active")
	}

	dir := filepath.Join(s.root, id)
	if err := os.Mkdir(dir, 0o700); err != nil {
		return nil, errors.Join(fmt.Errorf("create cache entry: %w", err), lease.release())
	}
	body := filepath.Join(dir, "body"+extension)
	e := &Entry{
		ID: id, Directory: dir, TempPath: body + ".temp", BodyPath: body,
		HeadersPath:  filepath.Join(dir, "headers.json"),
		MetadataPath: filepath.Join(dir, "metadata.yaml"), ErrorPath: body + ".error", store: s, lease: lease,
	}
	return e, nil
}

func (e *Entry) CreateBody() (*os.File, error) {
	f, err := os.OpenFile(e.TempPath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o600)
	if err != nil {
		return nil, fmt.Errorf("create temporary response body: %w", err)
	}
	return f, nil
}

func (e *Entry) PublishBody(f *os.File) error {
	if f != nil {
		if err := f.Sync(); err != nil {
			f.Close()
			return fmt.Errorf("flush response body: %w", err)
		}
		if err := f.Close(); err != nil {
			return fmt.Errorf("close response body: %w", err)
		}
	}
	if e.store.faults != nil && e.store.faults.PublishBody != nil {
		if err := e.store.faults.PublishBody(); err != nil {
			return fmt.Errorf("publish response body: %w", err)
		}
	}
	if err := os.Remove(e.ErrorPath); err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("remove response error before publication: %w", err)
	}
	if err := fsx.Replace(e.TempPath, e.BodyPath); err != nil {
		return fmt.Errorf("publish response body: %w", err)
	}
	return nil
}

func (e *Entry) SaveHeaders(headers map[string][]string) error {
	return writeJSONAtomic(e.HeadersPath, headers)
}

func (e *Entry) SaveMetadata(metadata Metadata) error {
	if e.store.faults != nil && e.store.faults.SaveMetadata != nil {
		if err := e.store.faults.SaveMetadata(metadata); err != nil {
			return err
		}
	}
	return writeJSONAtomic(e.MetadataPath, metadata)
}

func (e *Entry) PublishError(record ErrorRecord) error {
	_ = os.Remove(e.TempPath)
	if _, err := os.Stat(e.BodyPath); err == nil {
		return nil
	} else if !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("inspect published response body: %w", err)
	}
	return writeJSONAtomic(e.ErrorPath, record)
}

func (e *Entry) Abort() {
	_ = os.RemoveAll(e.Directory)
	e.Close()
}

func (e *Entry) Close() {
	e.once.Do(func() { _ = e.lease.release() })
}

// LimitWriter enforces the decoded-byte cap and the non-bypassable disk reserve.
type LimitWriter struct {
	store      *Store
	w          io.Writer
	limit      int64
	written    int64
	checkAfter int64
}

func (s *Store) Writer(w io.Writer, allowLarge bool) *LimitWriter {
	limit := s.max
	if allowLarge {
		limit = 0
	}
	return &LimitWriter{store: s, w: w, limit: limit}
}

func (w *LimitWriter) Write(p []byte) (int, error) {
	if w.limit > 0 && int64(len(p)) > w.limit-w.written {
		return 0, ErrSizeLimit
	}
	if w.store.reserve > 0 && (w.written == 0 || w.written >= w.checkAfter) {
		if err := w.store.CheckSpace(int64(len(p))); err != nil {
			return 0, err
		}
		w.checkAfter = w.written + (1 << 20)
	}
	n, err := w.w.Write(p)
	w.written += int64(n)
	return n, err
}

func (w *LimitWriter) Written() int64 { return w.written }

func (s *Store) Cleanup() (CleanupResult, error) {
	entries, err := os.ReadDir(s.root)
	if err != nil {
		return CleanupResult{}, fmt.Errorf("read cache root: %w", err)
	}
	var result CleanupResult
	var errs []error
	now := s.now()
	for _, item := range entries {
		if !item.IsDir() || item.Name() == filepath.Base(s.lockRoot) {
			continue
		}
		removed, orphan, err := s.cleanupEntry(item.Name(), now)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		if !removed {
			continue
		}
		if orphan {
			result.RemovedOrphans++
		} else {
			result.RemovedEntries++
		}
	}
	return result, errors.Join(errs...)
}

func (s *Store) tryAcquire(id string) (*entryLease, bool, error) {
	s.mu.Lock()
	if _, exists := s.active[id]; exists {
		s.mu.Unlock()
		return nil, false, nil
	}
	s.active[id] = struct{}{}
	s.mu.Unlock()

	file := flock.New(filepath.Join(s.lockRoot, id+".lock"))
	locked, err := file.TryLock()
	if err != nil || !locked {
		_ = file.Close()
		s.mu.Lock()
		delete(s.active, id)
		s.mu.Unlock()
		return nil, false, err
	}
	return &entryLease{store: s, id: id, file: file}, true, nil
}

func (l *entryLease) release() error {
	err := l.file.Close()
	l.store.mu.Lock()
	delete(l.store.active, l.id)
	l.store.mu.Unlock()
	if err != nil {
		return fmt.Errorf("release cache entry lease: %w", err)
	}
	return nil
}

func (s *Store) cleanupEntry(id string, now time.Time) (removed, orphan bool, returnErr error) {
	lease, acquired, err := s.tryAcquire(id)
	if err != nil || !acquired {
		return false, false, err
	}
	defer func() { returnErr = errors.Join(returnErr, lease.release()) }()

	dir := filepath.Join(s.root, id)
	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		return false, false, nil
	} else if err != nil {
		return false, false, fmt.Errorf("inspect cache entry %s: %w", id, err)
	}
	var metadata Metadata
	raw, readErr := os.ReadFile(filepath.Join(dir, "metadata.yaml"))
	if readErr == nil {
		readErr = json.Unmarshal(raw, &metadata)
	}
	bodyInfo, bodyErr := publishedBody(dir)
	if bodyErr != nil {
		return false, false, fmt.Errorf("inspect cache entry %s body: %w", id, bodyErr)
	}
	orphan = bodyInfo == nil && (readErr != nil || metadata.CompletedAt == nil)
	eligible := orphan
	completedAt := metadata.CompletedAt
	if completedAt == nil && bodyInfo != nil {
		modified := bodyInfo.ModTime()
		completedAt = &modified
	}
	if !orphan && s.retention > 0 && completedAt != nil {
		eligible = !completedAt.Add(s.retention).After(now)
	}
	if !eligible {
		return false, orphan, nil
	}
	if err := os.RemoveAll(dir); err != nil {
		return false, orphan, fmt.Errorf("remove cache entry %s: %w", id, err)
	}
	return true, orphan, nil
}

func publishedBody(dir string) (os.FileInfo, error) {
	items, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		name := item.Name()
		if item.IsDir() || !strings.HasPrefix(name, "body.") || strings.HasSuffix(name, ".temp") || strings.HasSuffix(name, ".error") {
			continue
		}
		return item.Info()
	}
	return nil, nil
}

func writeJSONAtomic(path string, value any) error {
	raw, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return fmt.Errorf("encode %s: %w", filepath.Base(path), err)
	}
	raw = append(raw, '\n')
	tmp, err := os.CreateTemp(filepath.Dir(path), ".write-*.tmp")
	if err != nil {
		return err
	}
	name := tmp.Name()
	defer os.Remove(name)
	if err := tmp.Chmod(0o600); err != nil {
		tmp.Close()
		return err
	}
	if _, err := tmp.Write(raw); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Sync(); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	if err := fsx.Replace(name, path); err != nil {
		return err
	}
	return nil
}
