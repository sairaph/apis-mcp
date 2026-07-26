package sessions

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestSessionPersistsCookies(t *testing.T) {
	root := t.TempDir()
	manager, err := New(root, 24*time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	handle, err := manager.Create(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	id := handle.ID()
	u, _ := url.Parse("https://api.example.test/v1/items")
	handle.Jar().SetCookies(u, []*http.Cookie{{Name: "token", Value: "secret", Path: "/v1", Secure: true, HttpOnly: true}})
	if err := handle.Close(); err != nil {
		t.Fatal(err)
	}

	manager, err = New(root, 24*time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	reopened, err := manager.Acquire(context.Background(), id)
	if err != nil {
		t.Fatal(err)
	}
	cookies := reopened.Jar().Cookies(u)
	if len(cookies) != 1 || cookies[0].Value != "secret" {
		t.Fatalf("cookies = %#v", cookies)
	}
	if err := reopened.Close(); err != nil {
		t.Fatal(err)
	}
	inspection, err := manager.Inspect(id)
	if err != nil {
		t.Fatal(err)
	}
	if inspection.Session.CookieCount != 1 || !inspection.Cookies[0].HTTPOnly {
		t.Fatalf("inspection = %#v", inspection)
	}
	if info, err := os.Stat(filepath.Join(root, id+".json")); err != nil || runtime.GOOS != "windows" && info.Mode().Perm() != 0o600 {
		t.Fatalf("session mode: %v, %v", info, err)
	}
}

func TestCookieDomainRulesAcrossOrigins(t *testing.T) {
	now := time.Date(2026, 7, 26, 12, 0, 0, 0, time.UTC)
	jar := newJar(nil, func() time.Time { return now })
	origin, _ := url.Parse("https://api.example.com/v1")
	sibling, _ := url.Parse("https://static.example.com/v1")
	unrelated, _ := url.Parse("https://attacker.test/v1")

	jar.SetCookies(origin, []*http.Cookie{
		{Name: "shared", Value: "yes", Domain: "example.com", Path: "/"},
		{Name: "host", Value: "yes", Path: "/"},
		{Name: "unrelated", Value: "no", Domain: "attacker.test", Path: "/"},
	})

	if got := cookieNames(jar.Cookies(sibling)); !equalStrings(got, []string{"shared"}) {
		t.Fatalf("cookies on sibling origin = %v", got)
	}
	if got := cookieNames(jar.Cookies(unrelated)); len(got) != 0 {
		t.Fatalf("cookies on unrelated origin = %v", got)
	}
	if got := cookieNames(jar.Cookies(origin)); !equalStrings(got, []string{"shared", "host"}) {
		t.Fatalf("cookies on source origin = %v", got)
	}
}

func TestCookieRejectsPublicSuffixDomains(t *testing.T) {
	tests := []struct {
		origin string
		domain string
	}{
		{origin: "https://shop.co.uk/", domain: "co.uk"},
		{origin: "https://tenant.appspot.com/", domain: "appspot.com"},
		{origin: "https://example.com/", domain: "com"},
	}
	for _, test := range tests {
		t.Run(test.domain, func(t *testing.T) {
			jar := newJar(nil, time.Now)
			u, _ := url.Parse(test.origin)
			jar.SetCookies(u, []*http.Cookie{{Name: "blocked", Value: "value", Domain: test.domain, Path: "/"}})
			if got := jar.snapshot(); len(got) != 0 {
				t.Fatalf("stored public-suffix cookie = %#v", got)
			}
		})
	}
}

func cookieNames(cookies []*http.Cookie) []string {
	result := make([]string, len(cookies))
	for i, cookie := range cookies {
		result[i] = cookie.Name
	}
	return result
}

func equalStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestDeleteBusySession(t *testing.T) {
	root := t.TempDir()
	manager, err := New(root, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	other, err := New(root, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	handle, err := manager.Create(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if err := other.Delete(handle.ID()); !errors.Is(err, ErrBusy) {
		t.Fatalf("delete error = %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel()
	if _, err := other.Acquire(ctx, handle.ID()); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("acquire error = %v", err)
	}
	if err := handle.Close(); err != nil {
		t.Fatal(err)
	}
	if err := other.Delete(handle.ID()); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(root, ".locks", handle.ID()+".lock")); err != nil {
		t.Fatalf("persistent lock file: %v", err)
	}
}

func TestSessionRetentionCleanup(t *testing.T) {
	now := time.Date(2026, 7, 26, 12, 0, 0, 0, time.UTC)
	manager, err := NewWithClock(t.TempDir(), time.Hour, func() time.Time { return now })
	if err != nil {
		t.Fatal(err)
	}
	handle, err := manager.Create(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if err := handle.Close(); err != nil {
		t.Fatal(err)
	}
	now = now.Add(2 * time.Hour)
	result, err := manager.Cleanup()
	if err != nil {
		t.Fatal(err)
	}
	if result.Removed != 1 {
		t.Fatalf("removed = %d", result.Removed)
	}
}

func TestSessionLockBookkeepingIsBounded(t *testing.T) {
	manager, err := New(t.TempDir(), time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 100; i++ {
		handle, err := manager.Create(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if len(manager.locks) != 1 {
			t.Fatalf("active lock count = %d, want 1", len(manager.locks))
		}
		if err := handle.Close(); err != nil {
			t.Fatal(err)
		}
		if len(manager.locks) != 0 {
			t.Fatalf("released lock count = %d, want 0", len(manager.locks))
		}
	}
}

func TestCleanupSkipsSessionHeldByOtherManager(t *testing.T) {
	now := time.Date(2026, 7, 26, 12, 0, 0, 0, time.UTC)
	root := t.TempDir()
	owner, err := NewWithClock(root, time.Hour, func() time.Time { return now })
	if err != nil {
		t.Fatal(err)
	}
	cleaner, err := NewWithClock(root, time.Hour, func() time.Time { return now })
	if err != nil {
		t.Fatal(err)
	}
	handle, err := owner.Create(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	id := handle.ID()
	now = now.Add(2 * time.Hour)
	result, err := cleaner.Cleanup()
	if err != nil {
		t.Fatal(err)
	}
	if result.Removed != 0 {
		t.Fatalf("removed busy sessions = %d", result.Removed)
	}
	if _, err := os.Stat(filepath.Join(root, id+".json")); err != nil {
		t.Fatalf("busy session was removed: %v", err)
	}
	if err := handle.Close(); err != nil {
		t.Fatal(err)
	}
	now = now.Add(2 * time.Hour)
	result, err = cleaner.Cleanup()
	if err != nil {
		t.Fatal(err)
	}
	if result.Removed != 1 {
		t.Fatalf("removed expired sessions = %d", result.Removed)
	}
}

func TestMultipleManagersSerializeUpdates(t *testing.T) {
	root := t.TempDir()
	const managerCount = 4
	managers := make([]*Manager, managerCount)
	for i := range managers {
		var err error
		managers[i], err = New(root, time.Hour)
		if err != nil {
			t.Fatal(err)
		}
	}
	handle, err := managers[0].Create(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	id := handle.ID()
	if err := handle.Close(); err != nil {
		t.Fatal(err)
	}

	u, _ := url.Parse("https://api.example.test/")
	const updates = 24
	errs := make(chan error, updates)
	var wg sync.WaitGroup
	for i := 0; i < updates; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			h, err := managers[i%len(managers)].Acquire(context.Background(), id)
			if err != nil {
				errs <- err
				return
			}
			h.Jar().SetCookies(u, []*http.Cookie{{Name: fmt.Sprintf("cookie-%d", i), Value: "value", Path: "/"}})
			errs <- h.Close()
		}(i)
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		if err != nil {
			t.Fatal(err)
		}
	}
	inspection, err := managers[0].Inspect(id)
	if err != nil {
		t.Fatal(err)
	}
	if inspection.Session.CookieCount != updates {
		t.Fatalf("cookie count = %d, want %d", inspection.Session.CookieCount, updates)
	}
}

func TestSessionLeaseReleasedWhenProcessExits(t *testing.T) {
	if os.Getenv("APIS_MCP_SESSION_LEASE_HELPER") == "1" {
		manager, err := New(os.Getenv("APIS_MCP_SESSION_ROOT"), time.Hour)
		if err != nil {
			t.Fatal(err)
		}
		handle, err := manager.Acquire(context.Background(), os.Getenv("APIS_MCP_SESSION_ID"))
		if err != nil {
			t.Fatal(err)
		}
		fmt.Fprintln(os.Stdout, "ready")
		time.Sleep(time.Minute)
		runtime.KeepAlive(handle)
		return
	}

	root := t.TempDir()
	manager, err := New(root, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	handle, err := manager.Create(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	id := handle.ID()
	if err := handle.Close(); err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command(os.Args[0], "-test.run=^TestSessionLeaseReleasedWhenProcessExits$")
	cmd.Env = append(os.Environ(), "APIS_MCP_SESSION_LEASE_HELPER=1", "APIS_MCP_SESSION_ROOT="+root, "APIS_MCP_SESSION_ID="+id)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}
	waited := false
	defer func() {
		if !waited {
			_ = cmd.Process.Kill()
			_ = cmd.Wait()
		}
	}()
	line, err := bufio.NewReader(stdout).ReadString('\n')
	if err != nil || strings.TrimSpace(line) != "ready" {
		t.Fatalf("helper readiness = %q, %v", line, err)
	}
	if err := manager.Delete(id); !errors.Is(err, ErrBusy) {
		t.Fatalf("delete while helper owns lease = %v", err)
	}
	if err := cmd.Process.Kill(); err != nil {
		t.Fatal(err)
	}
	_ = cmd.Wait()
	waited = true
	if err := manager.Delete(id); err != nil {
		t.Fatalf("delete after helper exit: %v", err)
	}
}
