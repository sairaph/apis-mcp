package docpacks

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"sync"
	"time"

	"github.com/gofrs/flock"
	"github.com/sairaph/apis-mcp/internal/fsx"
)

const packLockRetryDelay = 10 * time.Millisecond

const (
	progressUpdateLimit   int64 = 100
	progressUpdateMinimum int64 = 64 << 10
)

// ApplyStage identifies one transport-neutral step of preparing and publishing
// an API pack selection.
type ApplyStage string

const (
	ApplyStageWaiting       ApplyStage = "waiting"
	ApplyStageCheckingCache ApplyStage = "checking-cache"
	ApplyStageDownloading   ApplyStage = "downloading"
	ApplyStageVerifying     ApplyStage = "verifying"
	ApplyStageReady         ApplyStage = "ready"
	ApplyStageIndexing      ApplyStage = "indexing"
	ApplyStageApplying      ApplyStage = "applying"
)

// ApplyEvent is an immutable cumulative snapshot of apply progress. Pack fields
// are empty for selection-wide stages such as indexing and applying.
type ApplyEvent struct {
	Stage              ApplyStage
	PackID             string
	PackName           string
	PackBytesDone      int64
	PackBytesTotal     int64
	PreparedBytesDone  int64
	PreparedBytesTotal int64
	Cached             bool
}

// ApplyReporter receives cumulative progress snapshots. It is called
// synchronously by Apply and must not retain mutable state owned by Manager.
type ApplyReporter func(ApplyEvent)

// Options configures a Manager. A nil HTTP client and empty catalog URL use
// the production defaults.
type Options struct {
	HTTPClient *http.Client
	CatalogURL string
}

// ActiveState is the atomically published set of enabled packs, keyed by
// stable pack ID.
type ActiveState struct {
	SchemaVersion int             `json:"schema_version"`
	Packs         map[string]Pack `json:"packs"`
}

// RebuildFunc rebuilds the documentation library from the next archive list.
type RebuildFunc func(context.Context, []string) error

// Manager refreshes, validates, downloads, and activates documentation packs.
type Manager struct {
	root       string
	blobs      string
	catalogURL *url.URL
	client     *http.Client
	mu         sync.Mutex
}

// Open creates the local pack layout and opens a manager without network I/O.
func Open(root string, options Options) (*Manager, error) {
	if root == "" {
		return nil, errors.New("pack root is required")
	}
	absolute, err := filepath.Abs(root)
	if err != nil {
		return nil, fmt.Errorf("resolve pack root: %w", err)
	}
	catalogURL := options.CatalogURL
	if catalogURL == "" {
		catalogURL = ProductionCatalogURL
	}
	parsedURL, err := validateCatalogURL(catalogURL)
	if err != nil {
		return nil, err
	}
	blobs := filepath.Join(absolute, "blobs")
	if err := os.MkdirAll(blobs, 0o700); err != nil {
		return nil, fmt.Errorf("create pack directory: %w", err)
	}
	return &Manager{
		root: absolute, blobs: blobs, catalogURL: parsedURL,
		client: clientForCatalog(options.HTTPClient, parsedURL),
	}, nil
}

// Refresh fetches and verifies the configured catalog, then atomically caches
// it. A verified cached catalog is returned when the remote refresh fails.
func (m *Manager) Refresh(ctx context.Context) (Catalog, error) {
	var result Catalog
	err := m.withLock(ctx, func() error {
		catalog, err := m.fetchCatalog(ctx)
		if err != nil {
			if ctx.Err() != nil {
				return ctx.Err()
			}
			cached, cacheErr := m.cachedCatalog()
			if cacheErr == nil {
				result = cached
				return nil
			}
			return errors.Join(err, fmt.Errorf("load cached pack catalog: %w", cacheErr))
		}
		if err := m.writeJSON("catalog.json", catalog); err != nil {
			return err
		}
		result = catalog
		return nil
	})
	return result, err
}

// CachedCatalog returns the last verified catalog without network I/O.
func (m *Manager) CachedCatalog() (Catalog, error) {
	return m.cachedCatalog()
}

func (m *Manager) cachedCatalog() (Catalog, error) {
	raw, err := readCapped(filepath.Join(m.root, "catalog.json"), maxCatalogBytes)
	if err != nil {
		return Catalog{}, err
	}
	return decodeCatalog(raw)
}

// Active returns the validated active pack state. A new installation has an
// empty, valid state.
func (m *Manager) Active() (ActiveState, error) {
	state, _, err := m.readActive()
	return state, err
}

// ActiveArchives returns verified active blob paths in stable pack ID order.
// It performs no network I/O.
func (m *Manager) ActiveArchives() ([]string, error) {
	state, _, err := m.readActive()
	if err != nil {
		return nil, err
	}
	return m.archivesForState(state)
}

// Apply validates a desired selection, prepares its blobs, rebuilds the next
// library generation, and then atomically publishes the active state. A failed
// rebuild leaves the prior active state unchanged.
func (m *Manager) Apply(ctx context.Context, catalog Catalog, desiredIDs []string, rebuild RebuildFunc, reporter ApplyReporter) error {
	if rebuild == nil {
		return errors.New("pack apply requires a rebuild callback")
	}
	if err := validateCatalog(catalog); err != nil {
		return err
	}
	available := catalogPackMap(catalog)
	next := ActiveState{SchemaVersion: catalogSchemaVersion, Packs: make(map[string]Pack, len(desiredIDs))}
	for _, id := range desiredIDs {
		if next.Packs[id].ID != "" {
			return fmt.Errorf("duplicate desired pack ID %q", id)
		}
		pack, ok := available[id]
		if !ok {
			return fmt.Errorf("unknown desired pack ID %q", id)
		}
		next.Packs[id] = pack
	}
	ids := sortedPackIDs(next.Packs)
	var preparedTotal int64
	for _, id := range ids {
		preparedTotal += next.Packs[id].Bytes
	}
	report := func(event ApplyEvent) {
		if reporter != nil {
			reporter(event)
		}
	}
	for _, id := range ids {
		pack := next.Packs[id]
		report(ApplyEvent{
			Stage: ApplyStageWaiting, PackID: pack.ID, PackName: pack.Name,
			PackBytesTotal: pack.Bytes, PreparedBytesTotal: preparedTotal,
		})
	}

	return m.withLock(ctx, func() error {
		prior, _, err := m.readActive()
		if err != nil {
			return err
		}
		nextArchives := make([]string, 0, len(next.Packs))
		var prepared int64
		for _, id := range ids {
			pack := next.Packs[id]
			packBase := prepared
			reportPack := func(stage ApplyStage, done int64, cached bool) {
				report(ApplyEvent{
					Stage: stage, PackID: pack.ID, PackName: pack.Name,
					PackBytesDone: done, PackBytesTotal: pack.Bytes,
					PreparedBytesDone: packBase + done, PreparedBytesTotal: preparedTotal,
					Cached: cached,
				})
			}
			archive, cached, err := m.ensureBlob(ctx, pack, reportPack)
			if err != nil {
				return fmt.Errorf("prepare pack %s: %w", id, err)
			}
			nextArchives = append(nextArchives, archive)
			prepared += pack.Bytes
			reportPack(ApplyStageReady, pack.Bytes, cached)
		}
		if reflect.DeepEqual(prior, next) {
			return nil
		}
		report(ApplyEvent{Stage: ApplyStageIndexing, PreparedBytesDone: prepared, PreparedBytesTotal: preparedTotal})
		if err := rebuild(ctx, append([]string(nil), nextArchives...)); err != nil {
			return fmt.Errorf("rebuild library with active packs: %w", err)
		}
		if err := ctx.Err(); err != nil {
			return err
		}
		report(ApplyEvent{Stage: ApplyStageApplying, PreparedBytesDone: prepared, PreparedBytesTotal: preparedTotal})
		return m.writeJSON("active.json", next)
	})
}

func (m *Manager) fetchCatalog(ctx context.Context) (Catalog, error) {
	raw, err := m.fetch(ctx, m.catalogURL.String(), maxCatalogBytes)
	if err != nil {
		return Catalog{}, fmt.Errorf("fetch pack catalog: %w", err)
	}
	return decodeCatalog(raw)
}

func (m *Manager) fetch(ctx context.Context, target string, limit int64) ([]byte, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}
	response, err := m.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected HTTP status %s", response.Status)
	}
	if response.ContentLength > limit {
		return nil, fmt.Errorf("response exceeds %d bytes", limit)
	}
	raw, err := io.ReadAll(io.LimitReader(response.Body, limit+1))
	if err != nil {
		return nil, err
	}
	if int64(len(raw)) > limit {
		return nil, fmt.Errorf("response exceeds %d bytes", limit)
	}
	return raw, nil
}

func (m *Manager) ensureBlob(ctx context.Context, pack Pack, report func(ApplyStage, int64, bool)) (string, bool, error) {
	destination := m.blobPath(pack)
	report(ApplyStageCheckingCache, 0, false)
	if info, err := os.Lstat(destination); err == nil && info.Mode().IsRegular() && info.Size() == pack.Bytes {
		report(ApplyStageVerifying, 0, false)
		if err := verifyBlobContext(ctx, destination, pack, func(done int64) {
			report(ApplyStageVerifying, done, false)
		}); err == nil {
			return destination, true, nil
		} else if ctx.Err() != nil {
			return "", false, ctx.Err()
		}
		report(ApplyStageDownloading, 0, false)
	} else {
		report(ApplyStageDownloading, 0, false)
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, siblingURL(m.catalogURL, pack.Asset), nil)
	if err != nil {
		return "", false, err
	}
	response, err := m.client.Do(request)
	if err != nil {
		return "", false, fmt.Errorf("download pack asset: %w", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return "", false, fmt.Errorf("download pack asset: unexpected HTTP status %s", response.Status)
	}
	if response.ContentLength >= 0 && response.ContentLength != pack.Bytes {
		return "", false, fmt.Errorf("download pack asset: content length is %d, expected %d", response.ContentLength, pack.Bytes)
	}

	temporary, err := os.CreateTemp(m.blobs, ".blob-*.tmp")
	if err != nil {
		return "", false, fmt.Errorf("create temporary pack blob: %w", err)
	}
	temporaryName := temporary.Name()
	defer os.Remove(temporaryName)
	if err := temporary.Chmod(0o600); err != nil {
		temporary.Close()
		return "", false, err
	}
	hash := sha256.New()
	progress := newProgressWriter(pack.Bytes, func(done int64) { report(ApplyStageDownloading, done, false) })
	written, copyErr := io.Copy(io.MultiWriter(temporary, hash, progress), io.LimitReader(&contextReader{ctx: ctx, reader: response.Body}, pack.Bytes+1))
	if copyErr == nil && written != pack.Bytes {
		copyErr = fmt.Errorf("downloaded %d bytes, expected %d", written, pack.Bytes)
	}
	if copyErr == nil {
		actual := hex.EncodeToString(hash.Sum(nil))
		if actual != pack.SHA256 {
			copyErr = fmt.Errorf("downloaded SHA-256 is %s, expected %s", actual, pack.SHA256)
		}
	}
	if copyErr == nil {
		copyErr = temporary.Sync()
	}
	closeErr := temporary.Close()
	if err := errors.Join(copyErr, closeErr); err != nil {
		return "", false, err
	}
	report(ApplyStageVerifying, pack.Bytes, false)
	if err := validateArchiveContext(ctx, temporaryName, pack); err != nil {
		return "", false, err
	}
	if err := fsx.Replace(temporaryName, destination); err != nil {
		return "", false, fmt.Errorf("publish pack blob: %w", err)
	}
	if err := syncDirectory(m.blobs); err != nil {
		return "", false, err
	}
	return destination, false, nil
}

type contextReader struct {
	ctx    context.Context
	reader io.Reader
}

func (r *contextReader) Read(buffer []byte) (int, error) {
	if err := r.ctx.Err(); err != nil {
		return 0, err
	}
	return r.reader.Read(buffer)
}

type progressWriter struct {
	total     int64
	threshold int64
	written   int64
	reported  int64
	report    func(int64)
}

func newProgressWriter(total int64, report func(int64)) *progressWriter {
	threshold := max(progressUpdateMinimum, (total+progressUpdateLimit-1)/progressUpdateLimit)
	return &progressWriter{total: total, threshold: threshold, report: report}
}

func (w *progressWriter) Write(buffer []byte) (int, error) {
	w.written += int64(len(buffer))
	w.Update(w.written)
	return len(buffer), nil
}

func (w *progressWriter) Update(done int64) {
	done = min(done, w.total)
	if done < w.total && done-w.reported < w.threshold {
		return
	}
	if done > w.reported {
		w.reported = done
		w.report(done)
	}
}

func (m *Manager) readActive() (ActiveState, bool, error) {
	name := filepath.Join(m.root, "active.json")
	raw, err := readCapped(name, maxCatalogBytes)
	if errors.Is(err, os.ErrNotExist) {
		return emptyActive(), false, nil
	}
	if err != nil {
		return ActiveState{}, false, fmt.Errorf("read active packs: %w", err)
	}
	var state ActiveState
	if err := decodeStrict(raw, &state); err != nil {
		return ActiveState{}, true, fmt.Errorf("parse active packs: %w", err)
	}
	if err := validateActive(state); err != nil {
		return ActiveState{}, true, err
	}
	return state, true, nil
}

func emptyActive() ActiveState {
	return ActiveState{SchemaVersion: catalogSchemaVersion, Packs: make(map[string]Pack)}
}

func validateActive(state ActiveState) error {
	if state.SchemaVersion != catalogSchemaVersion {
		return fmt.Errorf("unsupported active pack schema %d", state.SchemaVersion)
	}
	if state.Packs == nil {
		return errors.New("active pack map is missing")
	}
	if len(state.Packs) > maxCatalogPacks {
		return errors.New("too many active packs")
	}
	assets := make(map[string]bool, len(state.Packs))
	for id, pack := range state.Packs {
		if id != pack.ID {
			return fmt.Errorf("active pack key %q disagrees with ID %q", id, pack.ID)
		}
		if err := validatePack(pack); err != nil {
			return fmt.Errorf("active pack %s: %w", id, err)
		}
		if assets[pack.Asset] {
			return fmt.Errorf("duplicate active pack asset %q", pack.Asset)
		}
		assets[pack.Asset] = true
	}
	return nil
}

func (m *Manager) archivesForState(state ActiveState) ([]string, error) {
	archives := make([]string, 0, len(state.Packs))
	for _, id := range sortedPackIDs(state.Packs) {
		pack := state.Packs[id]
		name := m.blobPath(pack)
		if err := verifyBlob(name, pack); err != nil {
			return nil, fmt.Errorf("verify active pack %s: %w", id, err)
		}
		archives = append(archives, name)
	}
	return archives, nil
}

func (m *Manager) blobPath(pack Pack) string {
	return filepath.Join(m.blobs, pack.SHA256+".zip")
}

func (m *Manager) writeJSON(base string, value any) error {
	raw, err := json.Marshal(value)
	if err != nil {
		return err
	}
	raw = append(raw, '\n')
	if int64(len(raw)) > maxCatalogBytes {
		return fmt.Errorf("encoded %s exceeds %d bytes", base, maxCatalogBytes)
	}
	name := filepath.Join(m.root, base)
	temporary, err := os.CreateTemp(m.root, ".pack-state-*.tmp")
	if err != nil {
		return fmt.Errorf("create temporary pack state: %w", err)
	}
	temporaryName := temporary.Name()
	defer os.Remove(temporaryName)
	if err := temporary.Chmod(0o600); err != nil {
		temporary.Close()
		return err
	}
	if _, err := temporary.Write(raw); err != nil {
		temporary.Close()
		return err
	}
	if err := temporary.Sync(); err != nil {
		temporary.Close()
		return err
	}
	if err := temporary.Close(); err != nil {
		return err
	}
	if err := fsx.Replace(temporaryName, name); err != nil {
		return fmt.Errorf("publish %s: %w", base, err)
	}
	return syncDirectory(m.root)
}

func (m *Manager) withLock(ctx context.Context, action func() error) (returnErr error) {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("lock pack publication: %w", err)
	}
	for !m.mu.TryLock() {
		timer := time.NewTimer(packLockRetryDelay)
		select {
		case <-ctx.Done():
			timer.Stop()
			return fmt.Errorf("lock pack publication: %w", ctx.Err())
		case <-timer.C:
		}
	}
	defer m.mu.Unlock()
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("lock pack publication: %w", err)
	}
	file := flock.New(filepath.Join(m.root, ".publish.lock"))
	locked, err := file.TryLockContext(ctx, packLockRetryDelay)
	if err != nil {
		_ = file.Close()
		return fmt.Errorf("lock pack publication: %w", err)
	}
	if !locked {
		_ = file.Close()
		return errors.New("pack publication lock was not acquired")
	}
	defer func() { returnErr = errors.Join(returnErr, file.Close()) }()
	return action()
}

func readCapped(name string, limit int64) ([]byte, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		return nil, err
	}
	if !info.Mode().IsRegular() || info.Size() > limit {
		return nil, fmt.Errorf("file is not a regular file of at most %d bytes", limit)
	}
	raw, err := io.ReadAll(io.LimitReader(file, limit+1))
	if err != nil {
		return nil, err
	}
	if int64(len(raw)) > limit {
		return nil, fmt.Errorf("file exceeds %d bytes", limit)
	}
	return raw, nil
}

func syncDirectory(name string) error {
	directory, err := os.Open(name)
	if err != nil {
		return err
	}
	err = directory.Sync()
	closeErr := directory.Close()
	if runtime.GOOS == "windows" {
		err = nil
	}
	return errors.Join(err, closeErr)
}
