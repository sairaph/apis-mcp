//go:build dev

package main

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/gofrs/flock"
	"github.com/google/uuid"
	"github.com/sairaph/apis-mcp/internal/fsx"
	"github.com/sairaph/apis-mcp/internal/importer"
	"github.com/sairaph/apis-mcp/library"
	"gopkg.in/yaml.v3"
)

type jobState string

const (
	jobQueued    jobState = "queued"
	jobRunning   jobState = "running"
	jobSucceeded jobState = "succeeded"
	jobFailed    jobState = "failed"
	jobCanceled  jobState = "canceled"
)

type ingestRequest struct {
	Output         string   `json:"output"`
	Source         string   `json:"source"`
	Name           string   `json:"name"`
	Version        string   `json:"version"`
	Collections    []string `json:"collections,omitempty"`
	Scope          string   `json:"scope"`
	MaxPages       int      `json:"max_pages"`
	MaxDepth       int      `json:"max_depth"`
	MaxSourceBytes int64    `json:"max_source_bytes,omitempty"`
	MaxTotalBytes  int64    `json:"max_total_bytes,omitempty"`
}

type ingestJob struct {
	ID              string              `json:"id"`
	State           jobState            `json:"state"`
	Revision        uint64              `json:"revision"`
	Request         ingestRequest       `json:"request"`
	Detection       *importer.Detection `json:"detection,omitempty"`
	Result          *importer.Result    `json:"result,omitempty"`
	Error           string              `json:"error,omitempty"`
	WorkerPID       int                 `json:"worker_pid,omitempty"`
	CurrentStage    string              `json:"current_stage,omitempty"`
	CurrentURL      string              `json:"current_url,omitempty"`
	Pages           int                 `json:"pages"`
	Queued          int                 `json:"queued"`
	Truncated       bool                `json:"truncated,omitempty"`
	CancelRequested bool                `json:"cancel_requested,omitempty"`
	CreatedAt       time.Time           `json:"created_at"`
	StartedAt       *time.Time          `json:"started_at,omitempty"`
	FinishedAt      *time.Time          `json:"finished_at,omitempty"`
}

type jobEvent struct {
	JobID     string    `json:"job_id"`
	Revision  uint64    `json:"revision"`
	Time      time.Time `json:"time"`
	Type      string    `json:"type"`
	State     jobState  `json:"state"`
	Stage     string    `json:"stage,omitempty"`
	Message   string    `json:"message,omitempty"`
	URL       string    `json:"url,omitempty"`
	Framework string    `json:"framework,omitempty"`
	Pages     int       `json:"pages"`
	Queued    int       `json:"queued"`
	Truncated bool      `json:"truncated,omitempty"`
}

type jobStore struct {
	output string
	root   string
}

var excludeBuiltinFromJobIndex bool

func openJobStore(output string, create bool) (*jobStore, error) {
	absolute, err := filepath.Abs(strings.TrimSpace(output))
	if err != nil || strings.TrimSpace(output) == "" {
		return nil, errors.New("output root is required")
	}
	root := filepath.Join(absolute, ".ingest", "jobs")
	if create {
		existing := root
		for {
			if _, statErr := os.Stat(existing); statErr == nil {
				break
			} else if !errors.Is(statErr, os.ErrNotExist) {
				return nil, statErr
			}
			parent := filepath.Dir(existing)
			if parent == existing {
				return nil, errors.New("no existing parent for job store")
			}
			existing = parent
		}
		if err := os.MkdirAll(root, 0o700); err != nil {
			return nil, err
		}
		if err := os.Chmod(filepath.Dir(root), 0o700); err != nil {
			return nil, err
		}
		if err := os.Chmod(root, 0o700); err != nil {
			return nil, err
		}
		if err := syncJobDirectoryChain(existing, root); err != nil {
			return nil, err
		}
	}
	return &jobStore{output: absolute, root: root}, nil
}

func syncJobDirectoryChain(existing, target string) error {
	directories := []string{target}
	for current := target; current != existing; {
		current = filepath.Dir(current)
		directories = append(directories, current)
	}
	for index := len(directories) - 1; index >= 0; index-- {
		if err := syncJobDirectory(directories[index]); err != nil {
			return err
		}
	}
	return nil
}

func (store *jobStore) jobPath(id string) (string, error) {
	if err := validateJobID(id); err != nil {
		return "", err
	}
	return filepath.Join(store.root, id+".json"), nil
}

func validateJobID(id string) error {
	parsed, err := uuid.Parse(id)
	if err != nil || parsed.String() != id {
		return errors.New("job ID must be a canonical UUID")
	}
	return nil
}

func (store *jobStore) create(request ingestRequest) (ingestJob, error) {
	source, err := url.Parse(strings.TrimSpace(request.Source))
	if err != nil || source.Host == "" || source.User != nil || source.Scheme != "http" && source.Scheme != "https" {
		return ingestJob{}, errors.New("job source must be a credential-free HTTP(S) URL")
	}
	id, err := uuid.NewV7()
	if err != nil {
		return ingestJob{}, err
	}
	job := ingestJob{ID: id.String(), State: jobQueued, Revision: 1, Request: request, CreatedAt: time.Now().UTC(), CurrentStage: "queued"}
	if err := writeJSONAtomic(filepath.Join(store.root, job.ID+".json"), job); err != nil {
		return ingestJob{}, err
	}
	if err := store.appendEvent(job, "queued", "job queued", importer.Progress{Stage: "queued"}); err != nil {
		return ingestJob{}, err
	}
	return job, nil
}

func (store *jobStore) get(id string) (ingestJob, error) {
	path, err := store.jobPath(id)
	if err != nil {
		return ingestJob{}, err
	}
	var job ingestJob
	if err := readJSON(path, &job); err != nil {
		return ingestJob{}, err
	}
	if job.ID != id {
		return ingestJob{}, errors.New("job file ID mismatch")
	}
	if job.State == jobRunning || job.State == jobQueued && time.Since(job.CreatedAt) > 10*time.Second {
		worker := flock.New(filepath.Join(store.root, id+".worker.lock"))
		locked, lockErr := worker.TryLock()
		if lockErr == nil && locked {
			defer worker.Unlock()
			if job.State == jobRunning {
				if recovered, ok := store.recoverPublished(job); ok {
					return recovered, nil
				}
				if _, rollbackErr := os.Stat(filepath.Join(store.root, id+".rollback")); rollbackErr == nil {
					if err := store.rollbackOwnedPublication(job); err != nil {
						return ingestJob{}, err
					}
					return store.finishCanceled(id, context.Canceled)
				}
				if store.ownsPublishedDestination(job) {
					if _, cancelErr := os.Stat(filepath.Join(store.root, id+".cancel")); cancelErr == nil {
						if rollbackErr := store.rollbackOwnedPublication(job); rollbackErr != nil {
							return ingestJob{}, rollbackErr
						}
						return store.finishCanceled(id, context.Canceled)
					}
					if recovered, ok := store.recoverOwnedPublication(job); ok {
						return recovered, nil
					}
					if _, cancelErr := os.Stat(filepath.Join(store.root, id+".cancel")); cancelErr == nil {
						if rollbackErr := store.rollbackOwnedPublication(job); rollbackErr != nil {
							return ingestJob{}, rollbackErr
						}
						return store.finishCanceled(id, context.Canceled)
					}
					if rollbackErr := store.rollbackOwnedPublication(job); rollbackErr != nil {
						return store.update(id, "failed", "owned publication recovery failed: "+rollbackErr.Error(), nil, func(current *ingestJob) {
							current.State = jobFailed
							current.Error = rollbackErr.Error()
							current.CurrentStage = "failed"
							now := time.Now().UTC()
							current.FinishedAt = &now
						})
					}
				}
			}
			if _, cancelErr := os.Stat(filepath.Join(store.root, id+".cancel")); cancelErr == nil {
				return store.finishCanceled(id, context.Canceled)
			}
			return store.update(id, "failed", "worker exited without a terminal state", nil, func(current *ingestJob) {
				if current.CancelRequested {
					current.State = jobCanceled
					current.Error = context.Canceled.Error()
					current.CurrentStage = "canceled"
				} else {
					current.State = jobFailed
					current.Error = "worker exited without a terminal state"
					current.CurrentStage = "failed"
				}
				now := time.Now().UTC()
				current.FinishedAt = &now
			})
		}
	}
	return job, nil
}

func (store *jobStore) list() ([]ingestJob, error) {
	entries, err := os.ReadDir(store.root)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []ingestJob{}, nil
		}
		return nil, err
	}
	var jobs []ingestJob
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".json" {
			continue
		}
		id := strings.TrimSuffix(entry.Name(), ".json")
		if validateJobID(id) != nil {
			continue
		}
		job, getErr := store.get(id)
		if getErr != nil {
			return nil, fmt.Errorf("read job %s: %w", id, getErr)
		}
		jobs = append(jobs, job)
	}
	sort.Slice(jobs, func(i, j int) bool { return jobs[i].CreatedAt.After(jobs[j].CreatedAt) })
	return jobs, nil
}

func (store *jobStore) update(id, eventType, message string, progress *importer.Progress, mutate func(*ingestJob)) (ingestJob, error) {
	path, err := store.jobPath(id)
	if err != nil {
		return ingestJob{}, err
	}
	lock := flock.New(filepath.Join(store.root, id+".state.lock"))
	if err := lock.Lock(); err != nil {
		return ingestJob{}, err
	}
	defer lock.Unlock()
	var job ingestJob
	if err := readJSON(path, &job); err != nil {
		return ingestJob{}, err
	}
	mutate(&job)
	job.Revision++
	if progress == nil {
		progress = &importer.Progress{Stage: job.CurrentStage, URL: job.CurrentURL, Pages: job.Pages, Queued: job.Queued}
	}
	if err := store.appendEvent(job, eventType, message, *progress); err != nil {
		return ingestJob{}, err
	}
	if err := writeJSONAtomic(path, job); err != nil {
		return ingestJob{}, err
	}
	return job, nil
}

func (store *jobStore) appendEvent(job ingestJob, eventType, message string, progress importer.Progress) error {
	if len(message) > 32<<10 {
		message = message[:32<<10] + "..."
	}
	event := jobEvent{
		JobID: job.ID, Revision: job.Revision, Time: time.Now().UTC(), Type: eventType, State: job.State,
		Stage: progress.Stage, Message: message, URL: progress.URL, Framework: progress.Framework,
		Pages: progress.Pages, Queued: progress.Queued, Truncated: progress.Truncated,
	}
	raw, err := json.Marshal(event)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(filepath.Join(store.root, job.ID+".events.jsonl"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o600)
	if err != nil {
		return err
	}
	_, writeErr := file.Write(append(raw, '\n'))
	syncErr := file.Sync()
	closeErr := file.Close()
	return errors.Join(writeErr, syncErr, closeErr, syncJobDirectory(store.root))
}

func (store *jobStore) cancel(id string) (ingestJob, error) {
	job, err := store.get(id)
	if err != nil || terminalJob(job.State) {
		return job, err
	}
	marker := filepath.Join(store.root, id+".cancel")
	if err := os.WriteFile(marker, []byte("cancel\n"), 0o600); err != nil {
		return ingestJob{}, err
	}
	return store.update(id, "cancel_requested", "cancellation requested", nil, func(current *ingestJob) {
		if terminalJob(current.State) {
			return
		}
		current.CancelRequested = true
		if current.State == jobQueued {
			current.State = jobCanceled
			current.CurrentStage = "canceled"
			now := time.Now().UTC()
			current.FinishedAt = &now
		} else {
			current.CurrentStage = "canceling"
		}
	})
}

func terminalJob(state jobState) bool {
	return state == jobSucceeded || state == jobFailed || state == jobCanceled
}

func (store *jobStore) recoverPublished(job ingestJob) (ingestJob, bool) {
	var result importer.Result
	if err := readJSON(filepath.Join(store.root, job.ID+".published.json"), &result); err != nil {
		return ingestJob{}, false
	}
	destination := filepath.Join(job.Request.Output, importer.SafeSlug(job.Request.Name), importer.SafeSlug(job.Request.Version))
	if result.Name != job.Request.Name || result.Version != job.Request.Version || result.Destination != destination {
		return ingestJob{}, false
	}
	if _, err := os.Stat(filepath.Join(destination, "_index.md")); err != nil {
		return ingestJob{}, false
	}
	snapshot, err := library.Open(context.Background(), library.Options{
		UserRoot: job.Request.Output, IndexPath: filepath.Join(job.Request.Output, "library.sqlite"), ExcludeBuiltin: excludeBuiltinFromJobIndex,
	})
	if err != nil {
		return ingestJob{}, false
	}
	defer snapshot.Close()
	listed, err := snapshot.List(context.Background(), library.ListRequest{Name: job.Request.Name, Version: job.Request.Version})
	if err != nil || listed.Total == 0 {
		return ingestJob{}, false
	}
	recovered, err := store.update(job.ID, "succeeded", "recovered published ingestion", &importer.Progress{
		Stage: "succeeded", URL: result.Source, Framework: result.Framework, Pages: result.Pages,
	}, func(current *ingestJob) {
		current.State = jobSucceeded
		current.Result = &result
		current.Pages = result.Pages
		current.Truncated = result.Truncated
		current.Queued = 0
		current.CurrentStage = "succeeded"
		now := time.Now().UTC()
		current.FinishedAt = &now
	})
	return recovered, err == nil
}

func (store *jobStore) ownsPublishedDestination(job ingestJob) bool {
	destination := filepath.Join(job.Request.Output, importer.SafeSlug(job.Request.Name), importer.SafeSlug(job.Request.Version))
	raw, err := os.ReadFile(filepath.Join(destination, ".apis-mcp-ingest-job"))
	return err == nil && strings.TrimSpace(string(raw)) == job.ID
}

func (store *jobStore) recoverOwnedPublication(job ingestJob) (ingestJob, bool) {
	destination := filepath.Join(job.Request.Output, importer.SafeSlug(job.Request.Name), importer.SafeSlug(job.Request.Version))
	raw, err := os.ReadFile(filepath.Join(destination, "_index.md"))
	if err != nil {
		return ingestJob{}, false
	}
	type manifest struct {
		SourceRoot string `yaml:"source_root"`
		SourceType string `yaml:"source_type"`
		Sources    int    `yaml:"sources"`
	}
	var metadata manifest
	text := strings.ReplaceAll(string(raw), "\r\n", "\n")
	if !strings.HasPrefix(text, "---\n") {
		return ingestJob{}, false
	}
	closing := strings.Index(text[4:], "\n---")
	if closing < 0 || yaml.Unmarshal([]byte(text[4:4+closing]), &metadata) != nil {
		return ingestJob{}, false
	}
	pages := 0
	if err := filepath.WalkDir(destination, func(name string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if !entry.IsDir() && strings.EqualFold(filepath.Ext(name), ".md") && filepath.Base(name) != "_index.md" {
			pages++
		}
		return nil
	}); err != nil {
		return ingestJob{}, false
	}
	snapshot, err := library.Open(context.Background(), library.Options{
		UserRoot: job.Request.Output, IndexPath: filepath.Join(job.Request.Output, "library.sqlite"), ExcludeBuiltin: excludeBuiltinFromJobIndex,
	})
	if err != nil {
		return ingestJob{}, false
	}
	defer snapshot.Close()
	listed, err := snapshot.List(context.Background(), library.ListRequest{Name: job.Request.Name, Version: job.Request.Version})
	if err != nil || listed.Total == 0 {
		return ingestJob{}, false
	}
	framework := ""
	if job.Detection != nil {
		framework = job.Detection.Framework
	}
	result := importer.Result{
		Kind: metadata.SourceType, Framework: framework, Name: job.Request.Name, Version: job.Request.Version,
		Source: metadata.SourceRoot, Destination: destination, Pages: pages, Sources: metadata.Sources, Truncated: job.Truncated,
	}
	recovered, err := store.update(job.ID, "succeeded", "recovered owned publication", &importer.Progress{
		Stage: "succeeded", URL: result.Source, Framework: framework, Pages: pages,
	}, func(current *ingestJob) {
		if current.CancelRequested {
			return
		}
		current.State = jobSucceeded
		current.Result = &result
		current.Pages = pages
		current.Truncated = result.Truncated
		current.Queued = 0
		current.CurrentStage = "succeeded"
		now := time.Now().UTC()
		current.FinishedAt = &now
	})
	return recovered, err == nil && recovered.State == jobSucceeded
}

func (store *jobStore) rollbackOwnedPublication(job ingestJob) error {
	rollbackMarker := filepath.Join(store.root, job.ID+".rollback")
	owned := store.ownsPublishedDestination(job)
	if !owned {
		if _, err := os.Stat(rollbackMarker); errors.Is(err, os.ErrNotExist) {
			return nil
		} else if err != nil {
			return err
		}
	}
	if owned {
		if err := os.WriteFile(rollbackMarker, []byte("rollback\n"), 0o600); err != nil {
			return err
		}
	}
	destination := filepath.Join(job.Request.Output, importer.SafeSlug(job.Request.Name), importer.SafeSlug(job.Request.Version))
	if owned {
		if err := os.RemoveAll(destination); err != nil {
			return err
		}
		if err := syncJobDirectory(filepath.Dir(destination)); err != nil {
			return err
		}
	}
	if err := library.Rebuild(context.Background(), library.Options{
		UserRoot: job.Request.Output, IndexPath: filepath.Join(job.Request.Output, "library.sqlite"), ExcludeBuiltin: excludeBuiltinFromJobIndex,
	}); err != nil {
		return err
	}
	if err := os.Remove(rollbackMarker); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	if err := syncJobDirectory(store.root); err != nil {
		return err
	}
	return nil
}

func startDetachedJob(store *jobStore, request ingestRequest) (ingestJob, error) {
	job, err := store.create(request)
	if err != nil {
		return ingestJob{}, err
	}
	executable, err := os.Executable()
	if err != nil {
		return store.failLaunch(job.ID, err)
	}
	logFile, err := os.OpenFile(filepath.Join(store.root, job.ID+".worker.log"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o600)
	if err != nil {
		return store.failLaunch(job.ID, err)
	}
	defer logFile.Close()
	devNull, err := os.Open(os.DevNull)
	if err != nil {
		return store.failLaunch(job.ID, err)
	}
	defer devNull.Close()
	command := exec.Command(executable, "__worker", "-out", store.output, "-id", job.ID)
	command.Stdin, command.Stdout, command.Stderr = devNull, logFile, logFile
	configureDetached(command)
	if err := command.Start(); err != nil {
		return store.failLaunch(job.ID, err)
	}
	if err := command.Process.Release(); err != nil {
		return store.failLaunch(job.ID, err)
	}
	return store.get(job.ID)
}

func (store *jobStore) failLaunch(id string, launchErr error) (ingestJob, error) {
	job, updateErr := store.update(id, "failed", launchErr.Error(), nil, func(current *ingestJob) {
		current.State = jobFailed
		current.Error = launchErr.Error()
		now := time.Now().UTC()
		current.FinishedAt = &now
	})
	return job, errors.Join(launchErr, updateErr)
}

func runWorker(ctx context.Context, store *jobStore, id string, client *http.Client) error {
	workerLock := flock.New(filepath.Join(store.root, id+".worker.lock"))
	locked, err := workerLock.TryLock()
	if err != nil || !locked {
		return errors.New("job already has an active worker")
	}
	defer workerLock.Unlock()
	job, err := store.get(id)
	if err != nil {
		return err
	}
	if terminalJob(job.State) {
		return nil
	}
	if _, err := os.Stat(filepath.Join(store.root, id+".cancel")); err == nil {
		_, err = store.finishCanceled(id, context.Canceled)
		return err
	}
	now := time.Now().UTC()
	job, err = store.update(id, "running", "worker started", nil, func(current *ingestJob) {
		if terminalJob(current.State) {
			return
		}
		if _, markerErr := os.Stat(filepath.Join(store.root, id+".cancel")); markerErr == nil {
			current.State = jobCanceled
			current.CancelRequested = true
			current.CurrentStage = "canceled"
			current.FinishedAt = &now
			return
		}
		current.State = jobRunning
		current.WorkerPID = os.Getpid()
		current.StartedAt = &now
		current.CurrentStage = "detecting"
	})
	if err != nil {
		return err
	}
	if job.State != jobRunning {
		return nil
	}
	workerCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	done := make(chan struct{})
	go func() {
		defer close(done)
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()
		for {
			select {
			case <-workerCtx.Done():
				return
			case <-ticker.C:
				if _, statErr := os.Stat(filepath.Join(store.root, id+".cancel")); statErr == nil {
					cancel()
					return
				}
			}
		}
	}()
	latestPages := 0
	progress := func(update importer.Progress) {
		if update.Pages > latestPages {
			latestPages = update.Pages
		}
		_, _ = store.update(id, "progress", update.Message, &update, func(current *ingestJob) {
			current.CurrentStage, current.CurrentURL = update.Stage, update.URL
			current.Pages, current.Queued = update.Pages, update.Queued
			current.Truncated = current.Truncated || update.Truncated
		})
	}
	indexPath := filepath.Join(job.Request.Output, "library.sqlite")
	maxSourceBytes := job.Request.MaxSourceBytes
	if maxSourceBytes == 0 {
		maxSourceBytes = importer.DefaultMaxSourceBytes
	}
	maxTotalBytes := job.Request.MaxTotalBytes
	if maxTotalBytes == 0 {
		maxTotalBytes = importer.DefaultMaxTotalBytes
	}
	options := importer.Options{
		LibraryRoot: job.Request.Output, HTTPClient: client, Collections: job.Request.Collections, JobID: id,
		HTMLScope: job.Request.Scope, HTMLLimitsSet: true, MaxHTMLPages: job.Request.MaxPages, MaxHTMLDepth: job.Request.MaxDepth, Progress: progress,
		MaxSourceBytes: maxSourceBytes, MaxTotalBytes: maxTotalBytes,
		Rebuild: func(rebuildCtx context.Context) error {
			progress(importer.Progress{Stage: "indexing", Pages: latestPages})
			return library.Rebuild(rebuildCtx, library.Options{UserRoot: job.Request.Output, IndexPath: indexPath, ExcludeBuiltin: excludeBuiltinFromJobIndex})
		},
	}
	detection, ingestErr := importer.DetectURL(workerCtx, job.Request.Source, options)
	if ingestErr == nil {
		_, ingestErr = store.update(id, "detected", "source detected", &importer.Progress{Stage: "detected", URL: detection.Source, Framework: detection.Framework}, func(current *ingestJob) {
			current.Detection = &detection
			current.CurrentStage = "detected"
			current.CurrentURL = detection.Source
		})
	}
	if ingestErr == nil && detection.Engine == "html" && detection.Framework == "unknown" {
		ingestErr = errors.New("automatic ingestion refused HTML with an unknown framework: generic anchor crawling has no finite completeness inventory; explicit `apis-mcp import html` remains available for intentional best-effort imports")
	}
	if ingestErr == nil && detection.Engine == "html" && !supportedHTMLFramework(detection.Framework) {
		ingestErr = fmt.Errorf("automatic ingestion detected unsupported documentation framework %s; no complete importer is available", detection.Framework)
	}
	if ingestErr == nil {
		remainingBytes := maxTotalBytes - detection.DownloadedBytes
		if remainingBytes < 1 {
			ingestErr = fmt.Errorf("detection consumed the %d-byte aggregate download limit", maxTotalBytes)
		} else {
			options.MaxTotalBytes = remainingBytes
			options.MaxSourceBytes = min(maxSourceBytes, remainingBytes)
		}
	}
	var result importer.Result
	if ingestErr == nil {
		importSource := job.Request.Source
		if detection.Engine == "openapi" && detection.Framework == "scalar" {
			importSource = detection.Source
			schemaOrigin, originErr := importer.NormalizedHTTPOrigin(detection.Source)
			if originErr != nil {
				ingestErr = errors.New("detected Scalar schema has an invalid HTTP(S) URL")
			} else {
				options.OpenAPIInitialOrigin = schemaOrigin
			}
		}
		switch {
		case ingestErr != nil:
		case detection.Engine == "openapi":
			result, ingestErr = importer.ImportOpenAPI(workerCtx, job.Request.Name, job.Request.Version, importSource, options)
		case detection.Engine == "html":
			result, ingestErr = importer.ImportHTML(workerCtx, job.Request.Name, job.Request.Version, job.Request.Source, options)
		case detection.Engine == "docsify":
			result, ingestErr = importer.ImportDocsify(workerCtx, job.Request.Name, job.Request.Version, job.Request.Source, options)
		default:
			ingestErr = fmt.Errorf("unsupported ingestion engine %q", detection.Engine)
		}
		if result.Framework == "" {
			result.Framework = detection.Framework
		}
	}
	cancel()
	<-done
	if ingestErr != nil {
		var rollbackErr *importer.RollbackError
		if errors.As(ingestErr, &rollbackErr) {
			_, err = store.update(id, "failed", ingestErr.Error(), nil, func(current *ingestJob) {
				current.State = jobFailed
				current.Error = ingestErr.Error()
				current.CurrentStage = "failed"
				finished := time.Now().UTC()
				current.FinishedAt = &finished
			})
			return err
		}
		if errors.Is(ingestErr, context.Canceled) {
			_, err = store.finishCanceled(id, ingestErr)
			return err
		}
		_, err = store.update(id, "failed", ingestErr.Error(), nil, func(current *ingestJob) {
			current.State = jobFailed
			current.Error = ingestErr.Error()
			current.CurrentStage = "failed"
			finished := time.Now().UTC()
			current.FinishedAt = &finished
		})
		return err
	}
	_, completed, err := store.finishSuccess(id, result)
	if err != nil {
		return err
	}
	if completed {
		return nil
	}
	if rollbackErr := store.rollbackOwnedPublication(job); rollbackErr != nil {
		_, _ = store.update(id, "failed", "cancellation rollback failed: "+rollbackErr.Error(), nil, func(current *ingestJob) {
			current.State = jobFailed
			current.Error = rollbackErr.Error()
			current.CurrentStage = "failed"
			finished := time.Now().UTC()
			current.FinishedAt = &finished
		})
		return rollbackErr
	}
	_, err = store.finishCanceled(id, context.Canceled)
	return err
}

func supportedHTMLFramework(framework string) bool {
	switch framework {
	case "docusaurus", "mkdocs-material", "mkdocs", "sphinx", "vitepress", "nextra", "astro-starlight", "mdbook":
		return true
	default:
		return false
	}
}

func (store *jobStore) finishSuccess(id string, result importer.Result) (ingestJob, bool, error) {
	path, err := store.jobPath(id)
	if err != nil {
		return ingestJob{}, false, err
	}
	lock := flock.New(filepath.Join(store.root, id+".state.lock"))
	if err := lock.Lock(); err != nil {
		return ingestJob{}, false, err
	}
	defer lock.Unlock()
	var job ingestJob
	if err := readJSON(path, &job); err != nil {
		return ingestJob{}, false, err
	}
	if job.CancelRequested {
		return job, false, nil
	}
	if _, err := os.Stat(filepath.Join(store.root, id+".cancel")); err == nil {
		return job, false, nil
	} else if !errors.Is(err, os.ErrNotExist) {
		return ingestJob{}, false, err
	}
	if err := writeJSONAtomic(filepath.Join(store.root, id+".published.json"), result); err != nil {
		return ingestJob{}, false, err
	}
	job.State = jobSucceeded
	job.Result = &result
	job.Pages = result.Pages
	job.Truncated = result.Truncated
	job.Queued = 0
	job.CurrentStage = "succeeded"
	finished := time.Now().UTC()
	job.FinishedAt = &finished
	job.Revision++
	progress := importer.Progress{Stage: "succeeded", URL: result.Source, Framework: result.Framework, Pages: result.Pages}
	if err := store.appendEvent(job, "succeeded", "ingestion completed", progress); err != nil {
		return ingestJob{}, false, err
	}
	if err := writeJSONAtomic(path, job); err != nil {
		return ingestJob{}, false, err
	}
	return job, true, nil
}

func (store *jobStore) finishCanceled(id string, cause error) (ingestJob, error) {
	return store.update(id, "canceled", "ingestion canceled", nil, func(current *ingestJob) {
		current.State = jobCanceled
		current.Error = cause.Error()
		current.CurrentStage = "canceled"
		current.CancelRequested = true
		finished := time.Now().UTC()
		current.FinishedAt = &finished
	})
}

func watchJob(ctx context.Context, store *jobStore, id string, output io.Writer) (ingestJob, error) {
	seen := uint64(0)
	for {
		events, err := store.events(id, seen)
		if err != nil {
			return ingestJob{}, err
		}
		for _, event := range events {
			raw, marshalErr := json.Marshal(event)
			if marshalErr != nil {
				return ingestJob{}, marshalErr
			}
			if _, err := fmt.Fprintln(output, string(raw)); err != nil {
				return ingestJob{}, err
			}
			seen++
		}
		job, err := store.get(id)
		if err != nil {
			return job, err
		}
		if terminalJob(job.State) {
			finalEvents, finalErr := store.events(id, seen)
			if finalErr != nil {
				return job, finalErr
			}
			for _, event := range finalEvents {
				raw, marshalErr := json.Marshal(event)
				if marshalErr != nil {
					return job, marshalErr
				}
				if _, writeErr := fmt.Fprintln(output, string(raw)); writeErr != nil {
					return job, writeErr
				}
			}
			return job, nil
		}
		select {
		case <-ctx.Done():
			return job, ctx.Err()
		case <-time.After(250 * time.Millisecond):
		}
	}
}

func (store *jobStore) events(id string, after uint64) ([]jobEvent, error) {
	if err := validateJobID(id); err != nil {
		return nil, err
	}
	file, err := os.Open(filepath.Join(store.root, id+".events.jsonl"))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var events []jobEvent
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 64<<10), 1<<20)
	line := uint64(0)
	for scanner.Scan() {
		line++
		var event jobEvent
		if err := json.Unmarshal(scanner.Bytes(), &event); err != nil {
			return nil, err
		}
		if line > after {
			events = append(events, event)
		}
	}
	return events, scanner.Err()
}

func writeJSONAtomic(name string, value any) error {
	raw, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	raw = append(raw, '\n')
	temporary, err := os.CreateTemp(filepath.Dir(name), ".job-*.tmp")
	if err != nil {
		return err
	}
	temporaryName := temporary.Name()
	defer os.Remove(temporaryName)
	if err := temporary.Chmod(0o600); err != nil {
		temporary.Close()
		return err
	}
	_, writeErr := temporary.Write(raw)
	syncErr := temporary.Sync()
	closeErr := temporary.Close()
	if err := errors.Join(writeErr, syncErr, closeErr); err != nil {
		return err
	}
	if err := fsx.Replace(temporaryName, name); err != nil {
		return err
	}
	return syncJobDirectory(filepath.Dir(name))
}

func syncJobDirectory(name string) error {
	if runtime.GOOS == "windows" {
		return nil
	}
	directory, err := os.Open(name)
	if err != nil {
		return err
	}
	syncErr := directory.Sync()
	closeErr := directory.Close()
	return errors.Join(syncErr, closeErr)
}

func readJSON(name string, value any) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(io.LimitReader(file, 1<<20))
	decoder.DisallowUnknownFields()
	return decoder.Decode(value)
}
