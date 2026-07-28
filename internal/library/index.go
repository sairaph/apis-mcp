package library

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/gofrs/flock"
	_ "modernc.org/sqlite"
)

const schemaVersion = "1"

const lockRetryDelay = 10 * time.Millisecond

// Snapshot pins one immutable index generation until it is closed.
type Snapshot struct {
	db              *sql.DB
	listTokenBudget int
	readTokenBudget int
	fingerprint     string
	generationLock  *flock.Flock
}

func Open(ctx context.Context, options Options, sources []Source) (*Snapshot, error) {
	options, err := normalizeOptions(options)
	if err != nil {
		return nil, err
	}
	catalog, err := loadCatalog(sources)
	if err != nil {
		return nil, err
	}
	var snapshot *Snapshot
	err = withPublicationLock(ctx, options.IndexPath, func() error {
		generation, err := ensureCatalog(ctx, options.IndexPath, catalog)
		if err != nil {
			return err
		}
		snapshot, err = openSnapshot(ctx, options, generation, catalog.fingerprint)
		if err != nil {
			return err
		}
		cleanupGenerations(options.IndexPath, generation)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return snapshot, nil
}

func Rebuild(ctx context.Context, options Options, sources []Source) error {
	options, err := normalizeOptions(options)
	if err != nil {
		return err
	}
	catalog, err := loadCatalog(sources)
	if err != nil {
		return err
	}
	return withPublicationLock(ctx, options.IndexPath, func() error {
		generation, err := ensureCatalog(ctx, options.IndexPath, catalog)
		if err != nil {
			return err
		}
		cleanupGenerations(options.IndexPath, generation)
		return nil
	})
}

func (s *Snapshot) Close() error {
	if s == nil || s.db == nil {
		return nil
	}
	err := s.db.Close()
	s.db = nil
	if s.generationLock != nil {
		err = errors.Join(err, s.generationLock.Unlock())
		s.generationLock = nil
	}
	return err
}

func (s *Snapshot) Fingerprint() string {
	if s == nil {
		return ""
	}
	return s.fingerprint
}

func indexFingerprint(ctx context.Context, indexPath string) (string, error) {
	db, err := openReadOnly(indexPath)
	if err != nil {
		return "", err
	}
	defer db.Close()
	var version, fingerprint string
	if err := db.QueryRowContext(ctx, "SELECT value FROM metadata WHERE key = 'schema_version'").Scan(&version); err != nil {
		return "", err
	}
	if version != schemaVersion {
		return "", fmt.Errorf("unsupported library index schema %s", version)
	}
	if err := db.QueryRowContext(ctx, "SELECT value FROM metadata WHERE key = 'fingerprint'").Scan(&fingerprint); err != nil {
		return "", err
	}
	return fingerprint, nil
}

func openSnapshot(ctx context.Context, options Options, generation, expectedFingerprint string) (*Snapshot, error) {
	generationLock := flock.New(generation + ".lock")
	locked, err := generationLock.TryRLockContext(ctx, lockRetryDelay)
	if err != nil {
		return nil, fmt.Errorf("lock library generation: %w", err)
	}
	if !locked {
		return nil, errors.New("library generation lock was not acquired")
	}
	db, err := openReadOnly(generation)
	if err != nil {
		generationLock.Unlock()
		return nil, fmt.Errorf("open library index: %w", err)
	}
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	var fingerprint string
	if err := db.QueryRowContext(ctx, "SELECT value FROM metadata WHERE key = 'fingerprint'").Scan(&fingerprint); err != nil {
		db.Close()
		generationLock.Unlock()
		return nil, fmt.Errorf("read library generation: %w", err)
	}
	if fingerprint != expectedFingerprint {
		db.Close()
		generationLock.Unlock()
		return nil, errors.New("published library generation changed while opening")
	}
	return &Snapshot{
		db:              db,
		listTokenBudget: options.ListTokenBudget,
		readTokenBudget: options.ReadTokenBudget,
		fingerprint:     fingerprint,
		generationLock:  generationLock,
	}, nil
}

func openReadOnly(indexPath string) (*sql.DB, error) {
	absolute, err := filepath.Abs(indexPath)
	if err != nil {
		return nil, err
	}
	uriPath := filepath.ToSlash(absolute)
	if runtime.GOOS == "windows" && filepath.VolumeName(absolute) != "" && !strings.HasPrefix(uriPath, "//") {
		// A drive-letter path is a file URL path, not an authority. The leading
		// slash produces file:///C:/... instead of the invalid file://C:/....
		uriPath = "/" + uriPath
	}
	uri := (&url.URL{Scheme: "file", Path: uriPath}).String()
	db, err := sql.Open("sqlite", uri+"?mode=ro&immutable=1")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func withPublicationLock(ctx context.Context, indexPath string, action func() error) error {
	parent := filepath.Dir(indexPath)
	if err := os.MkdirAll(parent, 0o700); err != nil {
		return fmt.Errorf("create library index directory: %w", err)
	}
	publicationLock := flock.New(indexPath + ".publish.lock")
	locked, err := publicationLock.TryLockContext(ctx, lockRetryDelay)
	if err != nil {
		return fmt.Errorf("lock library publication: %w", err)
	}
	if !locked {
		return errors.New("library publication lock was not acquired")
	}
	defer publicationLock.Unlock()
	return action()
}

func ensureCatalog(ctx context.Context, indexPath string, catalog *catalog) (string, error) {
	generation := generationPath(indexPath, catalog.fingerprint)
	if fingerprint, err := indexFingerprint(ctx, generation); err == nil && fingerprint == catalog.fingerprint {
		return generation, nil
	}
	generationLock := flock.New(generation + ".lock")
	locked, err := generationLock.TryLockContext(ctx, lockRetryDelay)
	if err != nil {
		return "", fmt.Errorf("lock invalid library generation: %w", err)
	}
	if !locked {
		return "", errors.New("invalid library generation lock was not acquired")
	}
	defer generationLock.Unlock()
	if err := os.Remove(generation); err != nil && !errors.Is(err, os.ErrNotExist) {
		return "", fmt.Errorf("remove invalid library generation: %w", err)
	}
	if err := publishCatalog(ctx, generation, catalog); err != nil {
		return "", err
	}
	return generation, nil
}

func generationPath(indexPath, fingerprint string) string {
	extension := filepath.Ext(indexPath)
	return strings.TrimSuffix(indexPath, extension) + "-" + fingerprint + extension
}

func publishCatalog(ctx context.Context, generation string, catalog *catalog) (returnErr error) {
	parent := filepath.Dir(generation)
	temporary, err := os.CreateTemp(parent, ".library-*.sqlite")
	if err != nil {
		return fmt.Errorf("create temporary library index: %w", err)
	}
	temporaryPath := temporary.Name()
	if err := temporary.Close(); err != nil {
		return err
	}
	if err := os.Remove(temporaryPath); err != nil {
		return err
	}
	defer func() {
		os.Remove(temporaryPath)
		os.Remove(temporaryPath + "-journal")
	}()

	db, err := sql.Open("sqlite", temporaryPath)
	if err != nil {
		return fmt.Errorf("open temporary library index: %w", err)
	}
	defer func() {
		if db == nil {
			return
		}
		if err := db.Close(); returnErr == nil && err != nil {
			returnErr = err
		}
	}()
	if _, err := db.ExecContext(ctx, "PRAGMA journal_mode=DELETE; PRAGMA synchronous=FULL; PRAGMA foreign_keys=ON"); err != nil {
		return fmt.Errorf("configure library index: %w", err)
	}
	if err := createSchema(ctx, db); err != nil {
		return err
	}
	if err := insertCatalog(ctx, db, catalog); err != nil {
		return err
	}
	if _, err := db.ExecContext(ctx, "PRAGMA optimize"); err != nil {
		return fmt.Errorf("optimize library index: %w", err)
	}
	if err := db.Close(); err != nil {
		return fmt.Errorf("close temporary library index: %w", err)
	}
	db = nil
	if err := os.Chmod(temporaryPath, 0o600); err != nil {
		return fmt.Errorf("secure temporary library index: %w", err)
	}
	// FlushFileBuffers requires a write-capable handle on Windows. os.Open uses
	// a read-only handle there and returns ERROR_ACCESS_DENIED from Sync.
	file, err := os.OpenFile(temporaryPath, os.O_RDWR, 0)
	if err != nil {
		return err
	}
	if err := file.Sync(); err != nil {
		file.Close()
		return fmt.Errorf("sync temporary library index: %w", err)
	}
	if err := file.Close(); err != nil {
		return err
	}
	if err := os.Rename(temporaryPath, generation); err != nil {
		return fmt.Errorf("publish library index: %w", err)
	}
	directory, err := os.Open(parent)
	if err != nil {
		return fmt.Errorf("open library index directory for sync: %w", err)
	}
	defer directory.Close()
	if err := directory.Sync(); err != nil {
		// Windows does not expose a directory fsync equivalent. The generation
		// file itself has already been flushed before its atomic rename.
		if runtime.GOOS == "windows" {
			return nil
		}
		return fmt.Errorf("sync library index directory: %w", err)
	}
	return nil
}

func cleanupGenerations(indexPath, keep string) {
	parent := filepath.Dir(indexPath)
	extension := filepath.Ext(indexPath)
	prefix := strings.TrimSuffix(filepath.Base(indexPath), extension) + "-"
	entries, err := os.ReadDir(parent)
	if err != nil {
		return
	}
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasPrefix(name, prefix) || !strings.HasSuffix(name, extension) {
			continue
		}
		fingerprint := strings.TrimSuffix(strings.TrimPrefix(name, prefix), extension)
		if len(fingerprint) != sha256.Size*2 || strings.Trim(fingerprint, "0123456789abcdef") != "" {
			continue
		}
		generation := filepath.Join(parent, name)
		if generation == keep {
			continue
		}
		generationLock := flock.New(generation + ".lock")
		locked, lockErr := generationLock.TryLock()
		if lockErr != nil || !locked {
			continue
		}
		if err := os.Remove(generation); err == nil || errors.Is(err, os.ErrNotExist) {
			generationLock.Unlock()
			_ = os.Remove(generation + ".lock")
			continue
		}
		generationLock.Unlock()
	}
}

func createSchema(ctx context.Context, db *sql.DB) error {
	const schema = `
CREATE TABLE metadata (
    key TEXT PRIMARY KEY,
    value TEXT NOT NULL
);
CREATE TABLE families (
    family_id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    collections TEXT NOT NULL
);
CREATE TABLE documents (
    doc_id TEXT PRIMARY KEY,
    family_id TEXT NOT NULL REFERENCES families(family_id),
    name TEXT NOT NULL,
    version TEXT NOT NULL,
    description TEXT NOT NULL,
    collections TEXT NOT NULL,
    source_root TEXT NOT NULL,
    source_location TEXT NOT NULL,
    page_count INTEGER NOT NULL
);
CREATE TABLE pages (
    doc_id TEXT NOT NULL REFERENCES documents(doc_id),
    page_id TEXT NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    path TEXT NOT NULL,
    source TEXT NOT NULL,
    http_methods TEXT NOT NULL,
    api_endpoints TEXT NOT NULL,
    operation_ids TEXT NOT NULL,
    body TEXT NOT NULL,
    relative_file TEXT NOT NULL,
    PRIMARY KEY (doc_id, page_id)
);
CREATE INDEX pages_navigation ON pages(doc_id, path, title, page_id);
CREATE VIRTUAL TABLE page_search USING fts5(
    doc_id UNINDEXED,
    page_id,
    title,
    path,
    headings,
    body,
    api_endpoints,
    operation_ids,
    tokenize = 'unicode61'
);`
	if _, err := db.ExecContext(ctx, schema); err != nil {
		return fmt.Errorf("create library index schema: %w", err)
	}
	return nil
}

func insertCatalog(ctx context.Context, db *sql.DB, catalog *catalog) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if _, err := tx.ExecContext(ctx,
		"INSERT INTO metadata(key, value) VALUES ('schema_version', ?), ('fingerprint', ?)",
		schemaVersion, catalog.fingerprint,
	); err != nil {
		return err
	}
	for _, family := range catalog.families {
		collections, _ := json.Marshal(family.Collections)
		if _, err := tx.ExecContext(ctx,
			"INSERT INTO families(family_id, name, description, collections) VALUES (?, ?, ?, ?)",
			family.ID, family.Name, family.Description, string(collections),
		); err != nil {
			return fmt.Errorf("insert API family %s: %w", family.ID, err)
		}
	}
	for _, document := range catalog.documents {
		collections, _ := json.Marshal(document.Collections)
		if _, err := tx.ExecContext(ctx, `
INSERT INTO documents(doc_id, family_id, name, version, description, collections, source_root, source_location, page_count)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			document.DocID, document.FamilyID, document.Name, document.Version,
			document.Description, string(collections), document.SourceRoot, document.Location, len(document.Pages),
		); err != nil {
			return fmt.Errorf("insert document %s: %w", document.DocID, err)
		}
		for _, page := range document.Pages {
			httpMethods, _ := json.Marshal(page.HTTPMethods)
			apiEndpoints, _ := json.Marshal(page.APIEndpoints)
			operationIDs, _ := json.Marshal(page.OperationIDs)
			if _, err := tx.ExecContext(ctx, `
INSERT INTO pages(doc_id, page_id, title, description, path, source, http_methods, api_endpoints, operation_ids, body, relative_file)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
				document.DocID, page.PageID, page.Title, page.Description, page.Path, page.Source,
				string(httpMethods), string(apiEndpoints), string(operationIDs), page.Body, page.RelativeFile,
			); err != nil {
				return fmt.Errorf("insert page %s/%s: %w", document.DocID, page.PageID, err)
			}
			if _, err := tx.ExecContext(ctx, `
INSERT INTO page_search(doc_id, page_id, title, path, headings, body, api_endpoints, operation_ids)
VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
				document.DocID, page.PageID, page.Title, page.Path, headings(page.Body), page.Body,
				strings.Join(page.APIEndpoints, "\n"), strings.Join(page.OperationIDs, "\n"),
			); err != nil {
				return fmt.Errorf("index page %s/%s: %w", document.DocID, page.PageID, err)
			}
		}
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit library index: %w", err)
	}
	return nil
}

func headings(markdown string) string {
	var found []string
	for _, line := range physicalLines(markdown) {
		trimmed := strings.TrimSpace(line.text)
		if strings.HasPrefix(trimmed, "#") {
			trimmed = strings.TrimSpace(strings.TrimLeft(trimmed, "#"))
			if trimmed != "" {
				found = append(found, trimmed)
			}
		}
	}
	return strings.Join(found, "\n")
}
