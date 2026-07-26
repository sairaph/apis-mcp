// Package sessions implements persistent, process-locked HTTP cookie sessions.
package sessions

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/gofrs/flock"
	"github.com/google/uuid"
	"github.com/sairaph/apis-mcp/internal/fsx"
	"golang.org/x/net/publicsuffix"
)

var ErrNotFound = errors.New("cookie session not found")
var ErrBusy = errors.New("cookie session is busy")

type Manager struct {
	root      string
	lockRoot  string
	retention time.Duration
	now       func() time.Time

	mu    sync.Mutex
	locks map[string]*sessionLock
}

type sessionLock struct {
	token chan struct{}
	path  string
	refs  int
}

type sessionLease struct {
	manager *Manager
	id      string
	local   *sessionLock
	file    *flock.Flock
}

type sessionFile struct {
	Version    int          `json:"version"`
	ID         string       `json:"id"`
	CreatedAt  time.Time    `json:"created_at"`
	LastUsedAt time.Time    `json:"last_used_at"`
	Cookies    []CookieInfo `json:"cookies"`
}

type SessionInfo struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	LastUsedAt  time.Time `json:"last_used_at"`
	CookieCount int       `json:"cookie_count"`
	Domains     []string  `json:"domains,omitempty"`
}

type CookieInfo struct {
	Name      string        `json:"name"`
	Value     string        `json:"value"`
	Domain    string        `json:"domain"`
	Path      string        `json:"path"`
	Expires   time.Time     `json:"expires,omitempty"`
	Secure    bool          `json:"secure"`
	HTTPOnly  bool          `json:"http_only"`
	SameSite  http.SameSite `json:"same_site"`
	HostOnly  bool          `json:"host_only"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

type Inspection struct {
	Session SessionInfo  `json:"session"`
	Cookies []CookieInfo `json:"cookies"`
}

type CleanupResult struct{ Removed int }

type Handle struct {
	manager *Manager
	lease   *sessionLease
	file    sessionFile
	jar     *Jar
	once    sync.Once
	err     error
}

func New(root string, retention time.Duration) (*Manager, error) {
	return NewWithClock(root, retention, time.Now)
}

func NewWithClock(root string, retention time.Duration, now func() time.Time) (*Manager, error) {
	if root == "" {
		return nil, errors.New("session root is required")
	}
	if retention < 0 {
		return nil, errors.New("session retention cannot be negative")
	}
	if now == nil {
		now = time.Now
	}
	if err := os.MkdirAll(root, 0o700); err != nil {
		return nil, fmt.Errorf("create session directory: %w", err)
	}
	if err := os.Chmod(root, 0o700); err != nil {
		return nil, fmt.Errorf("secure session directory: %w", err)
	}
	lockRoot := filepath.Join(root, ".locks")
	if err := os.MkdirAll(lockRoot, 0o700); err != nil {
		return nil, fmt.Errorf("create session lock directory: %w", err)
	}
	if err := os.Chmod(lockRoot, 0o700); err != nil {
		return nil, fmt.Errorf("secure session lock directory: %w", err)
	}
	return &Manager{root: root, lockRoot: lockRoot, retention: retention, now: now, locks: make(map[string]*sessionLock)}, nil
}

func (m *Manager) Create(ctx context.Context) (*Handle, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, fmt.Errorf("create session ID: %w", err)
	}
	lease, err := m.acquire(ctx, id.String())
	if err != nil {
		return nil, err
	}
	now := m.now().UTC()
	f := sessionFile{Version: 1, ID: id.String(), CreatedAt: now, LastUsedAt: now, Cookies: []CookieInfo{}}
	if err := m.save(f); err != nil {
		return nil, errors.Join(err, lease.release())
	}
	return &Handle{manager: m, lease: lease, file: f, jar: newJar(f.Cookies, m.now)}, nil
}

func (m *Manager) Acquire(ctx context.Context, id string) (*Handle, error) {
	if err := validateID(id); err != nil {
		return nil, err
	}
	lease, err := m.acquire(ctx, id)
	if err != nil {
		return nil, err
	}
	f, err := m.load(id)
	if err != nil {
		return nil, errors.Join(err, lease.release())
	}
	f.LastUsedAt = m.now().UTC()
	if err := m.save(f); err != nil {
		return nil, errors.Join(err, lease.release())
	}
	return &Handle{manager: m, lease: lease, file: f, jar: newJar(f.Cookies, m.now)}, nil
}

func (h *Handle) ID() string          { return h.file.ID }
func (h *Handle) Jar() http.CookieJar { return h.jar }

func (h *Handle) Close() error {
	h.once.Do(func() {
		h.file.Cookies = h.jar.snapshot()
		h.file.LastUsedAt = h.manager.now().UTC()
		h.err = errors.Join(h.manager.save(h.file), h.lease.release())
	})
	return h.err
}

func (m *Manager) List() ([]SessionInfo, error) {
	entries, err := os.ReadDir(m.root)
	if err != nil {
		return nil, fmt.Errorf("read sessions: %w", err)
	}
	result := make([]SessionInfo, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".json" {
			continue
		}
		id := strings.TrimSuffix(entry.Name(), ".json")
		if validateID(id) != nil {
			continue
		}
		f, err := m.load(id)
		if err != nil {
			return nil, err
		}
		result = append(result, summarize(f, m.now()))
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].LastUsedAt.Equal(result[j].LastUsedAt) {
			return result[i].ID < result[j].ID
		}
		return result[i].LastUsedAt.After(result[j].LastUsedAt)
	})
	return result, nil
}

func (m *Manager) Inspect(id string) (Inspection, error) {
	if err := validateID(id); err != nil {
		return Inspection{}, err
	}
	f, err := m.load(id)
	if err != nil {
		return Inspection{}, err
	}
	cookies := activeCookies(f.Cookies, m.now())
	sort.SliceStable(cookies, func(i, j int) bool {
		if cookies[i].Domain != cookies[j].Domain {
			return cookies[i].Domain < cookies[j].Domain
		}
		if cookies[i].Path != cookies[j].Path {
			return cookies[i].Path < cookies[j].Path
		}
		return cookies[i].Name < cookies[j].Name
	})
	f.Cookies = cookies
	return Inspection{Session: summarize(f, m.now()), Cookies: cookies}, nil
}

func (m *Manager) Delete(id string) (returnErr error) {
	if err := validateID(id); err != nil {
		return err
	}
	lease, acquired, err := m.tryAcquire(id)
	if err != nil {
		return err
	}
	if !acquired {
		return ErrBusy
	}
	defer func() { returnErr = errors.Join(returnErr, lease.release()) }()
	err = os.Remove(m.path(id))
	if errors.Is(err, os.ErrNotExist) {
		return ErrNotFound
	}
	if err != nil {
		return fmt.Errorf("delete cookie session: %w", err)
	}
	return nil
}

func (m *Manager) DeleteCookie(ctx context.Context, id, domain, path, name string) error {
	h, err := m.Acquire(ctx, id)
	if err != nil {
		return err
	}
	h.jar.delete(domain, path, name)
	return h.Close()
}

func (m *Manager) Cleanup() (CleanupResult, error) {
	if m.retention == 0 {
		return CleanupResult{}, nil
	}
	entries, err := os.ReadDir(m.root)
	if err != nil {
		return CleanupResult{}, fmt.Errorf("read sessions: %w", err)
	}
	var result CleanupResult
	var errs []error
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".json" {
			continue
		}
		id := strings.TrimSuffix(entry.Name(), ".json")
		if validateID(id) != nil {
			continue
		}
		removed, err := m.removeExpired(id)
		if err != nil {
			errs = append(errs, err)
		} else if removed {
			result.Removed++
		}
	}
	return result, errors.Join(errs...)
}

func (m *Manager) retainLock(id string) *sessionLock {
	m.mu.Lock()
	defer m.mu.Unlock()
	if lock := m.locks[id]; lock != nil {
		lock.refs++
		return lock
	}
	lock := &sessionLock{token: make(chan struct{}, 1), path: filepath.Join(m.lockRoot, id+".lock"), refs: 1}
	lock.token <- struct{}{}
	m.locks[id] = lock
	return lock
}

func (m *Manager) releaseLock(id string, lock *sessionLock) {
	m.mu.Lock()
	defer m.mu.Unlock()
	lock.refs--
	if lock.refs == 0 && m.locks[id] == lock {
		delete(m.locks, id)
	}
}

func acquire(ctx context.Context, lock *sessionLock) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-lock.token:
		return nil
	}
}

func release(lock *sessionLock) { lock.token <- struct{}{} }

func (m *Manager) acquire(ctx context.Context, id string) (*sessionLease, error) {
	local := m.retainLock(id)
	if err := acquire(ctx, local); err != nil {
		m.releaseLock(id, local)
		return nil, err
	}
	file := flock.New(local.path)
	locked, err := file.TryLockContext(ctx, 10*time.Millisecond)
	if err != nil || !locked {
		_ = file.Close()
		release(local)
		m.releaseLock(id, local)
		if err != nil {
			return nil, err
		}
		return nil, ErrBusy
	}
	return &sessionLease{manager: m, id: id, local: local, file: file}, nil
}

func (m *Manager) tryAcquire(id string) (*sessionLease, bool, error) {
	local := m.retainLock(id)
	select {
	case <-local.token:
	default:
		m.releaseLock(id, local)
		return nil, false, nil
	}
	file := flock.New(local.path)
	locked, err := file.TryLock()
	if err != nil || !locked {
		_ = file.Close()
		release(local)
		m.releaseLock(id, local)
		return nil, false, err
	}
	return &sessionLease{manager: m, id: id, local: local, file: file}, true, nil
}

func (l *sessionLease) release() error {
	err := l.file.Close()
	release(l.local)
	l.manager.releaseLock(l.id, l.local)
	if err != nil {
		return fmt.Errorf("release cookie session lease: %w", err)
	}
	return nil
}

func (m *Manager) removeExpired(id string) (removed bool, returnErr error) {
	lease, acquired, err := m.tryAcquire(id)
	if err != nil || !acquired {
		return false, err
	}
	defer func() { returnErr = errors.Join(returnErr, lease.release()) }()
	f, err := m.load(id)
	if errors.Is(err, ErrNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	if f.LastUsedAt.Add(m.retention).After(m.now()) {
		return false, nil
	}
	if err := os.Remove(m.path(id)); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, fmt.Errorf("delete cookie session: %w", err)
	}
	return true, nil
}

func (m *Manager) path(id string) string { return filepath.Join(m.root, id+".json") }

func (m *Manager) load(id string) (sessionFile, error) {
	raw, err := os.ReadFile(m.path(id))
	if errors.Is(err, os.ErrNotExist) {
		return sessionFile{}, ErrNotFound
	}
	if err != nil {
		return sessionFile{}, fmt.Errorf("read cookie session: %w", err)
	}
	var f sessionFile
	if err := json.Unmarshal(raw, &f); err != nil {
		return sessionFile{}, fmt.Errorf("parse cookie session: %w", err)
	}
	if f.Version != 1 || f.ID != id {
		return sessionFile{}, errors.New("invalid cookie session file")
	}
	return f, nil
}

func (m *Manager) save(f sessionFile) error {
	raw, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return fmt.Errorf("encode cookie session: %w", err)
	}
	raw = append(raw, '\n')
	tmp, err := os.CreateTemp(m.root, ".session-*.tmp")
	if err != nil {
		return fmt.Errorf("create temporary cookie session: %w", err)
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
	if err := fsx.Replace(name, m.path(f.ID)); err != nil {
		return fmt.Errorf("publish cookie session: %w", err)
	}
	return nil
}

func validateID(id string) error {
	parsed, err := uuid.Parse(id)
	if err != nil || parsed.String() != id || parsed.Version() != 7 {
		return errors.New("session ID must be a canonical server-generated UUIDv7")
	}
	return nil
}

func summarize(f sessionFile, now time.Time) SessionInfo {
	cookies := activeCookies(f.Cookies, now)
	domainSet := make(map[string]struct{})
	for _, cookie := range cookies {
		domainSet[cookie.Domain] = struct{}{}
	}
	domains := make([]string, 0, len(domainSet))
	for domain := range domainSet {
		domains = append(domains, domain)
	}
	sort.Strings(domains)
	return SessionInfo{ID: f.ID, CreatedAt: f.CreatedAt, LastUsedAt: f.LastUsedAt, CookieCount: len(cookies), Domains: domains}
}

// Jar is a serializable CookieJar. A session Handle serializes all use of it.
type Jar struct {
	mu      sync.Mutex
	cookies []CookieInfo
	now     func() time.Time
}

func newJar(cookies []CookieInfo, now func() time.Time) *Jar {
	return &Jar{cookies: activeCookies(cookies, now()), now: now}
}

func (j *Jar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	j.mu.Lock()
	defer j.mu.Unlock()
	now := j.now().UTC()
	host := canonicalHost(u)
	if host == "" {
		return
	}
	for _, cookie := range cookies {
		if cookie == nil || cookie.Name == "" {
			continue
		}
		domain, hostOnly, ok := cookieDomain(host, cookie.Domain)
		if !ok {
			continue
		}
		path := cookie.Path
		if path == "" || path[0] != '/' {
			path = defaultPath(u.Path)
		}
		index := -1
		for i := range j.cookies {
			if j.cookies[i].Name == cookie.Name && j.cookies[i].Domain == domain && j.cookies[i].Path == path {
				index = i
				break
			}
		}
		if cookie.MaxAge < 0 || (!cookie.Expires.IsZero() && !cookie.Expires.After(now)) {
			if index >= 0 {
				j.cookies = append(j.cookies[:index], j.cookies[index+1:]...)
			}
			continue
		}
		created := now
		if index >= 0 {
			created = j.cookies[index].CreatedAt
		}
		item := CookieInfo{Name: cookie.Name, Value: cookie.Value, Domain: domain, Path: path, Expires: cookie.Expires.UTC(), Secure: cookie.Secure, HTTPOnly: cookie.HttpOnly, SameSite: cookie.SameSite, HostOnly: hostOnly, CreatedAt: created, UpdatedAt: now}
		if cookie.MaxAge > 0 {
			item.Expires = now.Add(time.Duration(cookie.MaxAge) * time.Second)
		}
		if index >= 0 {
			j.cookies[index] = item
		} else {
			j.cookies = append(j.cookies, item)
		}
	}
}

func (j *Jar) Cookies(u *url.URL) []*http.Cookie {
	j.mu.Lock()
	defer j.mu.Unlock()
	now := j.now()
	j.cookies = activeCookies(j.cookies, now)
	host := canonicalHost(u)
	secure := strings.EqualFold(u.Scheme, "https")
	path := u.EscapedPath()
	if path == "" {
		path = "/"
	}
	matched := make([]CookieInfo, 0)
	for _, cookie := range j.cookies {
		if cookie.Secure && !secure {
			continue
		}
		if cookie.HostOnly && host != cookie.Domain {
			continue
		}
		if !cookie.HostOnly && host != cookie.Domain && !strings.HasSuffix(host, "."+cookie.Domain) {
			continue
		}
		if !pathMatch(path, cookie.Path) {
			continue
		}
		matched = append(matched, cookie)
	}
	sort.SliceStable(matched, func(i, k int) bool {
		if len(matched[i].Path) != len(matched[k].Path) {
			return len(matched[i].Path) > len(matched[k].Path)
		}
		return matched[i].CreatedAt.Before(matched[k].CreatedAt)
	})
	result := make([]*http.Cookie, 0, len(matched))
	for _, item := range matched {
		result = append(result, &http.Cookie{Name: item.Name, Value: item.Value})
	}
	return result
}

func (j *Jar) snapshot() []CookieInfo {
	j.mu.Lock()
	defer j.mu.Unlock()
	j.cookies = activeCookies(j.cookies, j.now())
	result := make([]CookieInfo, len(j.cookies))
	copy(result, j.cookies)
	return result
}

func (j *Jar) delete(domain, path, name string) {
	j.mu.Lock()
	defer j.mu.Unlock()
	for i := len(j.cookies) - 1; i >= 0; i-- {
		if j.cookies[i].Domain == domain && j.cookies[i].Path == path && j.cookies[i].Name == name {
			j.cookies = append(j.cookies[:i], j.cookies[i+1:]...)
		}
	}
}

func activeCookies(cookies []CookieInfo, now time.Time) []CookieInfo {
	result := make([]CookieInfo, 0, len(cookies))
	for _, cookie := range cookies {
		if cookie.Expires.IsZero() || cookie.Expires.After(now) {
			result = append(result, cookie)
		}
	}
	return result
}

func canonicalHost(u *url.URL) string {
	host := strings.ToLower(u.Hostname())
	if ip := net.ParseIP(host); ip != nil {
		return ip.String()
	}
	return strings.TrimSuffix(host, ".")
}

func cookieDomain(host, attribute string) (string, bool, bool) {
	if attribute == "" {
		return host, true, true
	}
	domain := strings.ToLower(strings.TrimPrefix(attribute, "."))
	domain = strings.TrimSuffix(domain, ".")
	if domain == "" {
		return "", false, false
	}
	if net.ParseIP(host) != nil {
		return domain, false, host == domain
	}
	if host != domain && !strings.HasSuffix(host, "."+domain) {
		return "", false, false
	}
	if suffix, _ := publicsuffix.PublicSuffix(domain); suffix == domain {
		return "", false, false
	}
	return domain, false, true
}

func defaultPath(path string) string {
	if path == "" || path[0] != '/' {
		return "/"
	}
	last := strings.LastIndex(path, "/")
	if last == 0 {
		return "/"
	}
	return path[:last]
}

func pathMatch(requestPath, cookiePath string) bool {
	if requestPath == cookiePath {
		return true
	}
	if !strings.HasPrefix(requestPath, cookiePath) {
		return false
	}
	return strings.HasSuffix(cookiePath, "/") || (len(requestPath) > len(cookiePath) && requestPath[len(cookiePath)] == '/')
}
