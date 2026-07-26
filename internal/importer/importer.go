// Package importer converts external documentation into canonical Markdown and
// publishes one document set atomically into the user library.
package importer

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/gofrs/flock"
	"gopkg.in/yaml.v3"
)

const (
	DefaultMaxSourceBytes = int64(16 << 20)
	DefaultMaxTotalBytes  = int64(64 << 20)
	defaultMaxFiles       = 2_000
)

// Options controls source access and publication. Rebuild must validate and
// atomically publish the library index rooted at LibraryRoot.
type Options struct {
	LibraryRoot    string
	Rebuild        func(context.Context) error
	HTTPClient     *http.Client
	Collections    []string
	JobID          string
	HTMLScope      string
	HTMLLimitsSet  bool
	Progress       func(Progress)
	MaxSourceBytes int64
	MaxTotalBytes  int64
	MaxHTMLPages   int
	MaxHTMLDepth   int
}

// Progress describes a durable ingestion progress update.
type Progress struct {
	Stage     string `json:"stage"`
	Message   string `json:"message,omitempty"`
	URL       string `json:"url,omitempty"`
	Framework string `json:"framework,omitempty"`
	Pages     int    `json:"pages,omitempty"`
	Queued    int    `json:"queued,omitempty"`
	Truncated bool   `json:"truncated,omitempty"`
}

// RollbackError reports a failed import together with a failure to restore the
// previous library state.
type RollbackError struct {
	Cause    error
	Rollback error
}

func (e *RollbackError) Error() string {
	return fmt.Sprintf("validate imported documentation: %v (rollback failed: %v)", e.Cause, e.Rollback)
}

func (e *RollbackError) Unwrap() []error { return []error{e.Cause, e.Rollback} }

// Result describes one successfully published document set.
type Result struct {
	Kind        string `json:"kind"`
	Framework   string `json:"framework,omitempty"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Pages       int    `json:"pages"`
	Truncated   bool   `json:"truncated,omitempty"`
}

// CollisionError reports all known paths participating in an import identity
// collision.
type CollisionError struct {
	DocID string
	Paths []string
}

func (e *CollisionError) Error() string {
	return fmt.Sprintf("import doc_id %q already exists: %v", e.DocID, e.Paths)
}

type manifest struct {
	Name         string   `yaml:"name"`
	Version      string   `yaml:"version"`
	Description  string   `yaml:"description,omitempty"`
	Collections  []string `yaml:"collections,omitempty"`
	SourceRoot   string   `yaml:"source_root,omitempty"`
	SourceType   string   `yaml:"source_type,omitempty"`
	ImportedFrom string   `yaml:"imported_from,omitempty"`
}

type pageFront struct {
	Title        string   `yaml:"title"`
	PageID       string   `yaml:"page_id,omitempty"`
	Path         string   `yaml:"path,omitempty"`
	Description  string   `yaml:"description,omitempty"`
	Source       string   `yaml:"source,omitempty"`
	HTTPMethods  []string `yaml:"http_methods,omitempty"`
	APIEndpoints []string `yaml:"api_endpoints,omitempty"`
	OperationIDs []string `yaml:"operation_ids,omitempty"`
	SourceType   string   `yaml:"source_type,omitempty"`
	ImportedFrom string   `yaml:"imported_from,omitempty"`
}

type sourceReader struct {
	client    *http.Client
	perSource int64
	total     int64
	used      int64
}

// ImportMarkdown imports a canonical document-set directory whose root
// contains _index.md. Only Markdown files are copied.
func ImportMarkdown(ctx context.Context, sourceDirectory string, options Options) (Result, error) {
	options, err := normalizeOptions(options)
	if err != nil {
		return Result{}, err
	}
	absolute, err := filepath.Abs(sourceDirectory)
	if err != nil {
		return Result{}, fmt.Errorf("resolve Markdown source: %w", err)
	}
	info, err := os.Stat(absolute)
	if err != nil {
		return Result{}, fmt.Errorf("inspect Markdown source: %w", err)
	}
	if !info.IsDir() {
		return Result{}, fmt.Errorf("Markdown source is not a directory: %s", absolute)
	}
	raw, err := readLimitedFile(filepath.Join(absolute, "_index.md"), options.MaxSourceBytes)
	if err != nil {
		return Result{}, fmt.Errorf("read canonical manifest: %w", err)
	}
	var metadata manifest
	if _, err := splitFrontmatter(raw, &metadata); err != nil {
		return Result{}, fmt.Errorf("parse canonical manifest: %w", err)
	}
	metadata.Name = strings.TrimSpace(metadata.Name)
	metadata.Version = strings.TrimSpace(metadata.Version)
	if metadata.Name == "" || metadata.Version == "" {
		return Result{}, errors.New("canonical manifest requires name and version")
	}

	pages := 0
	result, err := publish(ctx, options, metadata.Name, metadata.Version, func(stage string) error {
		var total int64
		var files int
		return filepath.WalkDir(absolute, func(name string, entry fs.DirEntry, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}
			relative, err := filepath.Rel(absolute, name)
			if err != nil {
				return err
			}
			if relative == "." {
				return nil
			}
			if entry.Type()&os.ModeSymlink != 0 {
				return fmt.Errorf("symbolic links are not supported: %s", name)
			}
			if entry.IsDir() {
				return nil
			}
			if !strings.EqualFold(filepath.Ext(entry.Name()), ".md") {
				return nil
			}
			if filepath.Base(relative) == "_index.md" && filepath.Dir(relative) != "." {
				return fmt.Errorf("nested canonical manifest is not supported: %s", relative)
			}
			files++
			if files > defaultMaxFiles {
				return fmt.Errorf("Markdown source exceeds %d files", defaultMaxFiles)
			}
			content, err := readLimitedFile(name, options.MaxSourceBytes)
			if err != nil {
				return err
			}
			total += int64(len(content))
			if total > options.MaxTotalBytes {
				return fmt.Errorf("Markdown source exceeds %d bytes", options.MaxTotalBytes)
			}
			if relative != "_index.md" {
				pages++
			}
			return writeFile(stage, relative, content)
		})
	})
	if err != nil {
		return Result{}, err
	}
	result.Kind, result.Source, result.Pages = "markdown", absolute, pages
	return result, nil
}

func normalizeOptions(options Options) (Options, error) {
	if strings.TrimSpace(options.LibraryRoot) == "" {
		return Options{}, errors.New("library root is required")
	}
	if options.Rebuild == nil {
		return Options{}, errors.New("library rebuild function is required")
	}
	if options.MaxSourceBytes == 0 {
		options.MaxSourceBytes = DefaultMaxSourceBytes
	}
	if options.MaxTotalBytes == 0 {
		options.MaxTotalBytes = DefaultMaxTotalBytes
	}
	if options.MaxSourceBytes < 1 || options.MaxTotalBytes < options.MaxSourceBytes {
		return Options{}, errors.New("download limits must be positive and total must cover one source")
	}
	if options.HTTPClient == nil {
		options.HTTPClient = &http.Client{Timeout: 30 * time.Second}
	}
	if !options.HTMLLimitsSet {
		if options.MaxHTMLPages == 0 {
			options.MaxHTMLPages = DefaultMaxHTMLPages
		}
		if options.MaxHTMLDepth == 0 {
			options.MaxHTMLDepth = DefaultMaxHTMLDepth
		}
	}
	if options.MaxHTMLPages == 0 || options.MaxHTMLPages < -1 {
		return Options{}, errors.New("HTML page limit must be -1 (unlimited) or positive")
	}
	if options.MaxHTMLDepth < -1 {
		return Options{}, errors.New("HTML depth limit must be -1 (unlimited) or non-negative")
	}
	if options.HTMLScope == "" {
		options.HTMLScope = "domain"
	}
	if options.HTMLScope != "domain" && options.HTMLScope != "path" {
		return Options{}, errors.New("HTML scope must be path or domain")
	}
	return options, nil
}

func reportProgress(options Options, progress Progress) {
	if options.Progress != nil {
		options.Progress(progress)
	}
}

func newSourceReader(options Options) *sourceReader {
	return &sourceReader{client: options.HTTPClient, perSource: options.MaxSourceBytes, total: options.MaxTotalBytes}
}

func (reader *sourceReader) read(ctx context.Context, source string, base *url.URL) ([]byte, string, error) {
	resolved, isHTTP, err := resolveSource(source, base)
	if err != nil {
		return nil, "", err
	}
	var raw []byte
	if isHTTP {
		request, err := http.NewRequestWithContext(ctx, http.MethodGet, resolved, nil)
		if err != nil {
			return nil, "", err
		}
		request.Header.Set("User-Agent", "apis-mcp-documentation-importer")
		response, err := reader.client.Do(request)
		if err != nil {
			return nil, "", fmt.Errorf("fetch %s: %w", resolved, err)
		}
		defer response.Body.Close()
		if response.StatusCode < 200 || response.StatusCode >= 300 {
			return nil, "", fmt.Errorf("fetch %s: HTTP %d", resolved, response.StatusCode)
		}
		raw, err = readLimited(response.Body, reader.perSource)
		if err != nil {
			return nil, "", fmt.Errorf("fetch %s: %w", resolved, err)
		}
		if response.Request != nil && response.Request.URL != nil {
			resolved = response.Request.URL.String()
		}
	} else {
		raw, err = readLimitedFile(resolved, reader.perSource)
		if err != nil {
			return nil, "", fmt.Errorf("read %s: %w", resolved, err)
		}
	}
	reader.used += int64(len(raw))
	if reader.used > reader.total {
		return nil, "", fmt.Errorf("documentation import exceeds %d total source bytes", reader.total)
	}
	return raw, resolved, nil
}

func resolveSource(source string, base *url.URL) (string, bool, error) {
	source = strings.TrimSpace(source)
	if source == "" {
		return "", false, errors.New("source is required")
	}
	parsed, err := url.Parse(source)
	if err != nil {
		return "", false, fmt.Errorf("invalid source %q: %w", source, err)
	}
	if base != nil {
		parsed = base.ResolveReference(parsed)
		parsed.Scheme = strings.ToLower(parsed.Scheme)
		if (base.Scheme == "http" || base.Scheme == "https") && parsed.Scheme != "http" && parsed.Scheme != "https" {
			return "", false, errors.New("remote indexes may only link to HTTP(S) sources")
		}
	}
	parsed.Scheme = strings.ToLower(parsed.Scheme)
	parsed.Fragment = ""
	if parsed.Scheme != "" {
		if parsed.Scheme == "file" {
			if parsed.Host != "" && parsed.Host != "localhost" {
				return "", false, errors.New("file source may not specify a remote host")
			}
			name, err := url.PathUnescape(parsed.Path)
			if err != nil {
				return "", false, fmt.Errorf("invalid file source: %w", err)
			}
			return filepath.Clean(filepath.FromSlash(name)), false, nil
		}
		if parsed.Scheme != "http" && parsed.Scheme != "https" {
			return "", false, fmt.Errorf("unsupported source scheme %q", parsed.Scheme)
		}
		if parsed.Host == "" {
			return "", false, errors.New("HTTP source requires a host")
		}
		return parsed.String(), true, nil
	}
	absolute, err := filepath.Abs(source)
	if err != nil {
		return "", false, err
	}
	return absolute, false, nil
}

func publish(ctx context.Context, options Options, name, version string, populate func(string) error) (result Result, returnErr error) {
	name, version = strings.TrimSpace(name), strings.TrimSpace(version)
	family, release := SafeSlug(name), SafeSlug(version)
	if family == "" || release == "" {
		return Result{}, errors.New("API name and version must contain ASCII letters or digits")
	}
	libraryRoot, err := filepath.Abs(options.LibraryRoot)
	if err != nil {
		return Result{}, fmt.Errorf("resolve library root: %w", err)
	}
	if err := os.MkdirAll(libraryRoot, 0o700); err != nil {
		return Result{}, fmt.Errorf("create library root: %w", err)
	}
	temporary, err := os.MkdirTemp(filepath.Dir(libraryRoot), ".apis-mcp-import-*")
	if err != nil {
		return Result{}, fmt.Errorf("create import staging directory: %w", err)
	}
	defer os.RemoveAll(temporary)
	stage := filepath.Join(temporary, "document")
	if err := os.Mkdir(stage, 0o700); err != nil {
		return Result{}, err
	}
	if err := populate(stage); err != nil {
		return Result{}, fmt.Errorf("stage documentation: %w", err)
	}
	if options.JobID != "" {
		if err := writeFile(stage, ".apis-mcp-ingest-job", []byte(options.JobID+"\n")); err != nil {
			return Result{}, fmt.Errorf("stage ingestion ownership: %w", err)
		}
	}
	if _, err := os.Stat(filepath.Join(stage, "_index.md")); err != nil {
		return Result{}, fmt.Errorf("staged documentation has no _index.md: %w", err)
	}
	if err := syncTree(stage); err != nil {
		return Result{}, fmt.Errorf("sync staged documentation: %w", err)
	}

	destination := filepath.Join(libraryRoot, family, release)
	docID := family + "-" + release
	transactionLock := flock.New(filepath.Join(libraryRoot, ".import.lock"))
	locked, err := transactionLock.TryLockContext(ctx, 10*time.Millisecond)
	if err != nil {
		return Result{}, fmt.Errorf("lock import transaction: %w", err)
	}
	if !locked {
		return Result{}, errors.New("import transaction lock was not acquired")
	}
	defer func() {
		returnErr = errors.Join(returnErr, transactionLock.Unlock())
	}()

	if info, err := os.Stat(destination); err == nil {
		paths := []string{destination}
		if info.IsDir() {
			if _, manifestErr := os.Stat(filepath.Join(destination, "_index.md")); manifestErr == nil {
				paths = append(paths, filepath.Join(destination, "_index.md"))
			}
		}
		return Result{}, newCollisionError(docID, paths...)
	} else if !errors.Is(err, os.ErrNotExist) {
		return Result{}, fmt.Errorf("inspect import destination: %w", err)
	}
	collisions, err := docIDCollisionPaths(libraryRoot, docID, options.MaxSourceBytes)
	if err != nil {
		return Result{}, err
	}
	if len(collisions) > 0 {
		collisions = append(collisions, filepath.Join(destination, "_index.md"))
		return Result{}, newCollisionError(docID, collisions...)
	}

	if err := os.MkdirAll(filepath.Dir(destination), 0o700); err != nil {
		return Result{}, err
	}
	if err := syncDirectory(libraryRoot); err != nil {
		return Result{}, fmt.Errorf("sync library root: %w", err)
	}
	if err := os.Rename(stage, destination); err != nil {
		_ = os.Remove(filepath.Dir(destination))
		return Result{}, fmt.Errorf("publish documentation: %w", err)
	}
	if err := syncDirectory(filepath.Dir(destination)); err != nil {
		_ = os.RemoveAll(destination)
		_ = syncPublicationParents(destination, libraryRoot)
		return Result{}, fmt.Errorf("sync published documentation: %w", err)
	}
	if err := syncDirectory(libraryRoot); err != nil {
		_ = os.RemoveAll(destination)
		_ = syncPublicationParents(destination, libraryRoot)
		return Result{}, fmt.Errorf("sync library root: %w", err)
	}
	if err := options.Rebuild(ctx); err != nil {
		removeErr := os.RemoveAll(destination)
		syncErr := syncPublicationParents(destination, libraryRoot)
		rebuildErr := options.Rebuild(context.WithoutCancel(ctx))
		if rollbackErr := errors.Join(removeErr, syncErr, rebuildErr); rollbackErr != nil {
			return Result{}, &RollbackError{Cause: err, Rollback: rollbackErr}
		}
		return Result{}, fmt.Errorf("validate imported documentation: %w", err)
	}
	return Result{Name: name, Version: version, Destination: destination}, nil
}

func syncPublicationParents(destination, libraryRoot string) error {
	parent := filepath.Dir(destination)
	removeErr := os.Remove(parent)
	if removeErr != nil && !errors.Is(removeErr, os.ErrNotExist) {
		if err := syncDirectory(parent); err != nil {
			return errors.Join(removeErr, err)
		}
	}
	return syncDirectory(libraryRoot)
}

func docIDCollisionPaths(libraryRoot, docID string, limit int64) ([]string, error) {
	var paths []string
	err := filepath.WalkDir(libraryRoot, func(name string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || entry.Name() != "_index.md" {
			return nil
		}
		raw, err := readLimitedFile(name, limit)
		if err != nil {
			return fmt.Errorf("inspect existing manifest %s: %w", name, err)
		}
		var metadata manifest
		if _, err := splitFrontmatter(raw, &metadata); err != nil {
			return nil
		}
		if SafeSlug(metadata.Name)+"-"+SafeSlug(metadata.Version) == docID {
			paths = append(paths, name)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("inspect existing documentation: %w", err)
	}
	sort.Strings(paths)
	return paths, nil
}

func newCollisionError(docID string, paths ...string) error {
	unique := make(map[string]bool)
	cleaned := make([]string, 0, len(paths))
	for _, name := range paths {
		if name != "" && !unique[name] {
			unique[name] = true
			cleaned = append(cleaned, name)
		}
	}
	sort.Strings(cleaned)
	return &CollisionError{DocID: docID, Paths: cleaned}
}

func writeCanonicalFile(root, relative string, front any, body string) error {
	raw, err := yaml.Marshal(front)
	if err != nil {
		return err
	}
	content := append([]byte("---\n"), raw...)
	content = append(content, []byte("---\n\n")...)
	content = append(content, strings.TrimSpace(strings.ReplaceAll(body, "\r\n", "\n"))...)
	content = append(content, '\n')
	return writeFile(root, relative, content)
}

func writeFile(root, relative string, content []byte) error {
	relative = filepath.Clean(relative)
	if relative == "." || filepath.IsAbs(relative) || relative == ".." || strings.HasPrefix(relative, ".."+string(filepath.Separator)) {
		return fmt.Errorf("unsafe output path %q", relative)
	}
	name := filepath.Join(root, relative)
	if err := os.MkdirAll(filepath.Dir(name), 0o700); err != nil {
		return err
	}
	if err := os.WriteFile(name, content, 0o600); err != nil {
		return fmt.Errorf("write %s: %w", relative, err)
	}
	return nil
}

func splitFrontmatter(raw []byte, target any) (string, error) {
	text := strings.ReplaceAll(strings.ReplaceAll(string(raw), "\r\n", "\n"), "\r", "\n")
	lines := strings.Split(text, "\n")
	if len(lines) == 0 || strings.TrimSpace(lines[0]) != "---" {
		return "", errors.New("YAML frontmatter must start on the first line")
	}
	for index := 1; index < len(lines); index++ {
		if strings.TrimSpace(lines[index]) != "---" {
			continue
		}
		if err := yaml.Unmarshal([]byte(strings.Join(lines[1:index], "\n")), target); err != nil {
			return "", err
		}
		body := lines[index+1:]
		if len(body) > 0 && body[0] == "" {
			body = body[1:]
		}
		return strings.Join(body, "\n"), nil
	}
	return "", errors.New("YAML frontmatter is missing its closing delimiter")
}

func readLimitedFile(name string, limit int64) ([]byte, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return readLimited(file, limit)
}

func readLimited(reader io.Reader, limit int64) ([]byte, error) {
	raw, err := io.ReadAll(io.LimitReader(reader, limit+1))
	if err != nil {
		return nil, err
	}
	if int64(len(raw)) > limit {
		return nil, fmt.Errorf("source exceeds %d bytes", limit)
	}
	return raw, nil
}

// SafeSlug returns a portable, deterministic path component.
func SafeSlug(value string) string {
	var output strings.Builder
	dash := false
	for _, character := range strings.ToLower(strings.TrimSpace(value)) {
		if character >= 'a' && character <= 'z' || character >= '0' && character <= '9' {
			if dash && output.Len() > 0 {
				output.WriteByte('-')
			}
			output.WriteRune(character)
			dash = false
			continue
		}
		// IDs intentionally remain portable ASCII; every other rune separates words.
		dash = true
	}
	return strings.Trim(output.String(), "-")
}

func stableID(prefix, value string) string {
	slug := SafeSlug(value)
	if slug == "" {
		slug = "document"
	}
	if len(slug) > 72 {
		slug = strings.Trim(slug[:72], "-")
	}
	sum := sha256.Sum256([]byte(value))
	return prefix + slug + "-" + hex.EncodeToString(sum[:4])
}

func sortedKeys[V any](values map[string]V) []string {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func syncTree(root string) error {
	var directories []string
	if err := filepath.WalkDir(root, func(name string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			directories = append(directories, name)
			return nil
		}
		file, err := os.Open(name)
		if err != nil {
			return err
		}
		syncErr := file.Sync()
		closeErr := file.Close()
		return errors.Join(syncErr, closeErr)
	}); err != nil {
		return err
	}
	for index := len(directories) - 1; index >= 0; index-- {
		if err := syncDirectory(directories[index]); err != nil {
			return err
		}
	}
	return nil
}

func syncDirectory(name string) error {
	directory, err := os.Open(name)
	if err != nil {
		return err
	}
	syncErr := directory.Sync()
	closeErr := directory.Close()
	if runtime.GOOS == "windows" {
		syncErr = nil
	}
	return errors.Join(syncErr, closeErr)
}
