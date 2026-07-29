package bootstrap

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/sairaph/apis-mcp/internal/cache"
	"github.com/sairaph/apis-mcp/internal/config"
	"github.com/sairaph/apis-mcp/internal/docpacks"
	"github.com/sairaph/apis-mcp/internal/httpcall"
	"github.com/sairaph/apis-mcp/internal/sessions"
	"github.com/sairaph/apis-mcp/library"
)

const (
	runtimeCleanupInterval = time.Hour
	maxRuntimeDiagnostics  = 64
)

type Runtime struct {
	Paths    config.Paths
	Config   config.Config
	Library  *library.Snapshot
	Cache    *cache.Store
	Sessions *sessions.Manager
	HTTP     *httpcall.Service
	Packs    *docpacks.Manager

	cleanupCancel context.CancelFunc
	cleanupDone   chan struct{}
	libraryMu     sync.RWMutex
	packArchives  []string
	diagnosticMu  sync.Mutex
	diagnostics   []error
	closeOnce     sync.Once
	closeErr      error
}

func Open(ctx context.Context) (*Runtime, error) {
	paths, err := config.DefaultPaths()
	if err != nil {
		return nil, err
	}
	for _, dir := range []string{paths.Root, paths.Library, paths.Packs, paths.Index, paths.Cache, paths.Sessions} {
		if err := os.MkdirAll(dir, 0o700); err != nil {
			return nil, fmt.Errorf("create application directory %s: %w", dir, err)
		}
	}
	cfg, err := config.Load(paths)
	if err != nil {
		return nil, err
	}
	packManager, err := docpacks.Open(paths.Packs, docpacks.Options{})
	if err != nil {
		return nil, fmt.Errorf("open documentation packs: %w", err)
	}
	packArchives, err := packManager.ActiveArchives()
	if err != nil {
		return nil, fmt.Errorf("open active documentation packs: %w", err)
	}
	snapshot, err := library.Open(ctx, library.Options{
		UserRoot:        paths.Library,
		IndexPath:       filepath.Join(paths.Index, "library.sqlite"),
		ListTokenBudget: cfg.ListTokenBudget,
		ReadTokenBudget: cfg.ReadTokenBudget,
		PackArchives:    packArchives,
	})
	if err != nil {
		return nil, fmt.Errorf("open documentation library: %w", err)
	}
	cacheStore, err := cache.New(cache.Config{
		Root:            paths.Cache,
		MaxDecodedBytes: cfg.ResponseSizeLimit,
		FreeDiskReserve: cfg.FreeDiskReserve,
		Retention:       cfg.Retention,
	})
	if err != nil {
		snapshot.Close()
		return nil, err
	}
	sessionManager, err := sessions.New(paths.Sessions, cfg.Retention)
	if err != nil {
		snapshot.Close()
		return nil, err
	}
	runtime := &Runtime{
		Paths: paths, Config: cfg, Library: snapshot, Cache: cacheStore,
		Sessions: sessionManager, Packs: packManager, packArchives: append([]string(nil), packArchives...),
	}
	_ = runtime.cleanup("initial")

	httpConfig := httpcall.DefaultConfig(cacheStore, sessionManager)
	httpConfig.MaximumHeaderTimeout = cfg.MaximumHeaderTimeout
	httpConfig.MaximumRetries = cfg.MaximumRetries
	httpConfig.MaximumRedirects = cfg.MaximumRedirects
	httpConfig.BackgroundAfter = cfg.BackgroundAfter
	httpConfig.StalledDownloadAfter = cfg.StalledDownloadAfter
	httpConfig.ReadTokenBudget = cfg.ReadTokenBudget
	if !cfg.TLSVerify {
		transport := http.DefaultTransport.(*http.Transport).Clone()
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // User-controlled local setting.
		httpConfig.Transport = transport
	}
	httpService, err := httpcall.New(httpConfig)
	if err != nil {
		snapshot.Close()
		return nil, err
	}
	runtime.HTTP = httpService
	runtime.startCleanup(runtimeCleanupInterval)
	return runtime, nil
}

func (r *Runtime) Close(ctx context.Context) error {
	if r == nil {
		return nil
	}
	r.closeOnce.Do(func() {
		r.closeErr = r.close(ctx)
	})
	return r.closeErr
}

func (r *Runtime) close(ctx context.Context) error {
	var errs []error
	if r.cleanupCancel != nil {
		r.cleanupCancel()
	}
	if r.cleanupDone != nil {
		select {
		case <-r.cleanupDone:
		case <-ctx.Done():
			errs = append(errs, ctx.Err())
		}
	}
	if r.HTTP != nil {
		shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		if err := r.HTTP.Shutdown(shutdownCtx); err != nil {
			errs = append(errs, err)
		}
		cancel()
	}
	if r.Library != nil {
		if err := r.Library.Close(); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

// Diagnostics returns cleanup failures observed by this runtime.
func (r *Runtime) Diagnostics() []error {
	if r == nil {
		return nil
	}
	r.diagnosticMu.Lock()
	defer r.diagnosticMu.Unlock()
	result := make([]error, len(r.diagnostics))
	copy(result, r.diagnostics)
	return result
}

func (r *Runtime) startCleanup(interval time.Duration) {
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	r.cleanupCancel = cancel
	r.cleanupDone = done
	go func() {
		defer close(done)
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				_ = r.cleanup("periodic")
			}
		}
	}()
}

func (r *Runtime) cleanup(phase string) error {
	var errs []error
	if r.Cache != nil {
		if _, err := r.Cache.Cleanup(); err != nil {
			errs = append(errs, fmt.Errorf("cache: %w", err))
		}
	}
	if r.Sessions != nil {
		if _, err := r.Sessions.Cleanup(); err != nil {
			errs = append(errs, fmt.Errorf("sessions: %w", err))
		}
	}
	err := errors.Join(errs...)
	if err == nil {
		return nil
	}
	err = fmt.Errorf("%s cleanup: %w", phase, err)
	r.diagnosticMu.Lock()
	defer r.diagnosticMu.Unlock()
	if writeErr := appendDiagnostic(r.Paths.Diagnostics, err); writeErr != nil {
		err = errors.Join(err, fmt.Errorf("write cleanup diagnostic: %w", writeErr))
	}
	if len(r.diagnostics) == maxRuntimeDiagnostics {
		r.diagnostics = r.diagnostics[1:]
	}
	r.diagnostics = append(r.diagnostics, err)
	return err
}

func appendDiagnostic(path string, diagnostic error) error {
	if path == "" {
		return fmt.Errorf("diagnostics path is empty")
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o600)
	if err != nil {
		return err
	}
	_, writeErr := fmt.Fprintf(file, "%s %v\n", time.Now().UTC().Format(time.RFC3339Nano), diagnostic)
	return errors.Join(writeErr, file.Close())
}

func (r *Runtime) RebuildLibrary(ctx context.Context) error {
	archives := []string(nil)
	if r.Packs != nil {
		var err error
		archives, err = r.Packs.ActiveArchives()
		if err != nil {
			return fmt.Errorf("open active documentation packs: %w", err)
		}
	}
	options := library.Options{
		UserRoot:        r.Paths.Library,
		IndexPath:       filepath.Join(r.Paths.Index, "library.sqlite"),
		ListTokenBudget: r.Config.ListTokenBudget,
		ReadTokenBudget: r.Config.ReadTokenBudget,
		PackArchives:    archives,
	}
	if err := library.Rebuild(ctx, options); err != nil {
		return err
	}
	r.libraryMu.Lock()
	r.packArchives = append(r.packArchives[:0], archives...)
	r.libraryMu.Unlock()
	return nil
}

// LibraryOptions returns the current pack and user library configuration for
// opening another snapshot.
func (r *Runtime) LibraryOptions() library.Options {
	r.libraryMu.RLock()
	archives := append([]string(nil), r.packArchives...)
	r.libraryMu.RUnlock()
	return library.Options{
		UserRoot:        r.Paths.Library,
		IndexPath:       filepath.Join(r.Paths.Index, "library.sqlite"),
		ListTokenBudget: r.Config.ListTokenBudget,
		ReadTokenBudget: r.Config.ReadTokenBudget,
		PackArchives:    archives,
	}
}
