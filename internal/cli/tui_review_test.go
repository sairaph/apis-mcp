package cli

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sairaph/apis-mcp/internal/bootstrap"
	"github.com/sairaph/apis-mcp/internal/config"
	"github.com/sairaph/apis-mcp/internal/httpcall"
	"github.com/sairaph/apis-mcp/internal/install"
	"github.com/sairaph/apis-mcp/internal/sessions"
	"github.com/sairaph/apis-mcp/library"
)

const hostileTerminalText = "visible\x1b[2J\x1b]2;owned-title\x07\x1b]52;c;c3RvbGVu\x07\x00\x08\x9bhidden"

func TestUntrustedTerminalTextIsSanitizedBeforeStyling(t *testing.T) {
	clean := safeMultiline(hostileTerminalText + "\n\tkept")
	for _, forbidden := range []string{"\x1b", "]2;", "]52;", "c3RvbGVu", "\x00", "\x08", "\x9b"} {
		if strings.Contains(clean, forbidden) {
			t.Fatalf("sanitized text retained %q: %q", forbidden, clean)
		}
	}
	if !strings.Contains(clean, "\n\tkept") {
		t.Fatalf("multiline sanitization removed safe layout: %q", clean)
	}

	m := testRoot(t, contextDocumentation, nil)
	m.docs.collections = []library.Collection{{Name: hostileTerminalText, Collection: "bad", APICount: 1}}
	m.docs.apis = []library.API{{Name: hostileTerminalText, Description: hostileTerminalText, Versions: []library.APIVersion{{Version: hostileTerminalText, DocID: "bad", Pages: 1}}}}
	assertSafeRenderedView(t, m.View())

	m.setContext(contextSessions)
	m.sessions.setItems([]sessions.SessionInfo{{ID: hostileTerminalText, CookieCount: 1, Domains: []string{hostileTerminalText}}})
	m.sessions.setInspection(sessions.Inspection{Session: m.sessions.items[0], Cookies: []sessions.CookieInfo{{Name: hostileTerminalText, Domain: hostileTerminalText, Path: hostileTerminalText}}})
	m.stack = []screenID{screenSessionList, screenSessionDetail}
	assertSafeRenderedView(t, m.View())

	m.setContext(contextRequest)
	m.request.applyResult(httpcall.Result{
		Request:   httpcall.RequestResult{ID: hostileTerminalText, Method: hostileTerminalText, Endpoint: hostileTerminalText, SessionID: hostileTerminalText, AutomaticHeaders: map[string]string{hostileTerminalText: hostileTerminalText}},
		Response:  httpcall.ResponseResult{Status: 200, StatusText: hostileTerminalText, State: hostileTerminalText, ContentType: hostileTerminalText, Headers: map[string][]string{hostileTerminalText: {hostileTerminalText}}},
		Cache:     httpcall.CacheResult{Directory: hostileTerminalText, MetadataPath: hostileTerminalText},
		Attempts:  []httpcall.Attempt{{Number: 1, Error: hostileTerminalText, RetryReason: hostileTerminalText}},
		Redirects: []httpcall.Redirect{{FromURL: hostileTerminalText, ToURL: hostileTerminalText}},
		Preview:   &httpcall.Preview{Kind: hostileTerminalText, Content: hostileTerminalText},
	})
	assertSafeRenderedView(t, m.View())
}

func assertSafeRenderedView(t *testing.T, view string) {
	t.Helper()
	for _, forbidden := range []string{"\x1b[2J", "\x1b]2;", "\x1b]52;", "c3RvbGVu", "\x00", "\x08", "\x9b"} {
		if strings.Contains(view, forbidden) {
			t.Fatalf("view retained hostile terminal fragment %q", forbidden)
		}
	}
	styled := "\x1b[31mapp-owned\x1b[0m"
	if !strings.Contains(fixed(styled, 20), styled) {
		t.Fatal("layout helpers stripped application-owned styles")
	}
}

func TestNonCancellableWriteWaitsAndBlocksOverlap(t *testing.T) {
	m := testRoot(t, contextSettings, nil)
	started := make(chan struct{})
	release := make(chan struct{})
	cmd := m.startOperation("settings-save", "durable write", false, func(context.Context, int64) tea.Msg {
		close(started)
		<-release
		return asyncMsg{id: 1, kind: "cache-clean", value: "finished"}
	})
	operation := operationCommand(t, cmd)
	result := make(chan tea.Msg, 1)
	go func() { result <- operation() }()
	<-started
	m.Update(key("ctrl+c"))
	if m.active == nil || !strings.Contains(m.status, "Cannot cancel") {
		t.Fatal("ctrl+c claimed to cancel a durable write")
	}
	if overlap := m.startOperation("restore", "overlap", false, func(context.Context, int64) tea.Msg { return nil }); overlap != nil {
		t.Fatal("overlapping durable operation was launched")
	}
	closed := make(chan struct{})
	go func() { m.close(); close(closed) }()
	select {
	case <-closed:
		t.Fatal("close returned before durable write completed")
	case <-time.After(30 * time.Millisecond):
	}
	close(release)
	<-result
	select {
	case <-closed:
	case <-time.After(time.Second):
		t.Fatal("close did not wait for and finish durable write")
	}
}

func TestCloseDoesNotWaitForUnscheduledCommand(t *testing.T) {
	m := testRoot(t, contextSettings, nil)
	ran := false
	cmd := m.startOperation("settings-save", "queued write", false, func(context.Context, int64) tea.Msg {
		ran = true
		return nil
	})
	closed := make(chan struct{})
	go func() { m.close(); close(closed) }()
	select {
	case <-closed:
	case <-time.After(time.Second):
		t.Fatal("close waited forever for a command Bubble Tea never invoked")
	}
	operation := operationCommand(t, cmd)
	if message := operation(); message != nil || ran {
		t.Fatalf("command ran after shutdown: message=%T ran=%t", message, ran)
	}
}

func TestCancellableCloseAndPagesRollback(t *testing.T) {
	m := testRoot(t, contextDocumentation, nil)
	m.docs.frames = append(m.docs.frames, documentFrame{loading: true})
	m.stack = []screenID{screenDocumentation, screenDocumentBrowser}
	started := make(chan struct{})
	cmd := m.startOperation("pages", "pages", true, func(ctx context.Context, id int64) tea.Msg {
		close(started)
		<-ctx.Done()
		return asyncMsg{id: id, kind: "pages", err: ctx.Err()}
	})
	operation := operationCommand(t, cmd)
	result := make(chan tea.Msg, 1)
	go func() { result <- operation() }()
	<-started
	m.Update(key("ctrl+c"))
	if len(m.docs.frames) != 0 || len(m.stack) != 1 {
		t.Fatalf("cancelled pages load left loading state: frames=%d stack=%v", len(m.docs.frames), m.stack)
	}
	<-result
	m.close()
}

func TestSnapshotSwapStaleAndPendingTeardownCloseImmediately(t *testing.T) {
	m, closeFixture := documentationFixture(t)
	defer closeFixture()
	first, err := library.Open(context.Background(), fixtureLibraryOptions(m.runtime))
	if err != nil {
		t.Fatal(err)
	}
	collections, apis, err := loadCatalog(context.Background(), first)
	if err != nil {
		t.Fatal(err)
	}
	m.applyReload(reloadPayload{snapshot: first, collections: collections, apis: apis})
	second, err := library.Open(context.Background(), fixtureLibraryOptions(m.runtime))
	if err != nil {
		t.Fatal(err)
	}
	m.applyReload(reloadPayload{snapshot: second, collections: collections, apis: apis})
	assertSnapshotClosed(t, first)

	stale, err := library.Open(context.Background(), fixtureLibraryOptions(m.runtime))
	if err != nil {
		t.Fatal(err)
	}
	m.accept(asyncMsg{id: 999, value: reloadPayload{snapshot: stale}})
	assertSnapshotClosed(t, stale)

	pending, err := library.Open(context.Background(), fixtureLibraryOptions(m.runtime))
	if err != nil {
		t.Fatal(err)
	}
	m.registerSnapshot(42, pending)
	m.close()
	assertSnapshotClosed(t, pending)
	assertSnapshotClosed(t, second)
}

func fixtureLibraryOptions(runtime *bootstrap.Runtime) library.Options {
	return libraryOptions(runtime)
}

func assertSnapshotClosed(t *testing.T, snapshot *library.Snapshot) {
	t.Helper()
	closed := false
	func() {
		defer func() { closed = closed || recover() != nil }()
		_, err := snapshot.List(context.Background(), library.ListRequest{})
		closed = err != nil
	}()
	if !closed {
		t.Fatal("snapshot remained usable after ownership close")
	}
}

func TestLargeDownloadPolicyCannotBeBypassed(t *testing.T) {
	cfg := config.Default()
	cfg.AllowLargeDownload = false
	runtime := &bootstrap.Runtime{Config: cfg}
	m := testRoot(t, contextRequest, runtime)
	m.request.allowLarge = true
	m.request.setLargeDownloadPolicy(false)
	if m.request.allowLarge || m.request.largeDownloadsAllowed {
		t.Fatal("disabled policy retained large-download override")
	}
	m.request.focus = 8
	m.Update(key("enter"))
	m.request.url.SetValue("https://example.test")
	input, err := m.request.input()
	if err != nil {
		t.Fatal(err)
	}
	if input.AllowLargeDownload {
		t.Fatal("request input bypassed disabled large-download policy")
	}

	m.settings.cfg = config.Default()
	m.settings.cfg.AllowLargeDownload = false
	m.active = &operation{id: 7, cancel: func() {}}
	m.Update(asyncMsg{id: 7, kind: "settings-save", value: settingsSavePayload{message: "saved"}})
	if m.request.allowLarge || m.request.largeDownloadsAllowed {
		t.Fatal("settings save did not reset large-download override")
	}
}

func TestConfigureOnlyPreselectionRestrictionAndResult(t *testing.T) {
	clients := []install.Status{
		{Client: install.Client{ID: "detected", Name: "Detected"}, Detected: true},
		{Client: install.Client{ID: "unsafe", Name: "Unsafe"}, Detected: true, Err: errors.New("bad config")},
		{Client: install.Client{ID: "configured", Name: "Configured"}, Configured: true},
	}
	m := testRoot(t, contextSettings, nil)
	m.configureOnly = true
	m.settings.setClients(clients, true)
	if !m.settings.selected[0] || m.settings.selected[1] || !m.settings.selected[2] {
		t.Fatalf("configure preselection is unsafe or incomplete: %v", m.settings.selected)
	}
	m.settings.setClients(clients, false)
	if m.settings.selected[0] || m.settings.selected[1] || !m.settings.selected[2] {
		t.Fatalf("main Settings did not reflect configured clients only: %v", m.settings.selected)
	}
	m.settings.setClients(clients, true)
	m.setContext(contextDocumentation)
	if m.context != contextSettings || strings.Contains(m.contextBar(), "Documentation") {
		t.Fatal("configure-only mode exposed unavailable contexts")
	}
	if err := configureRunResult(context.Background(), m, nil); err == nil {
		t.Fatal("configure cancellation was reported as success")
	}
	m.configureSaved = true
	m.settings.dirty = false
	if err := configureRunResult(context.Background(), m, nil); err != nil {
		t.Fatalf("successful save was not reported: %v", err)
	}
	saveFailure := errors.New("save failed")
	m.active = &operation{id: 9, cancel: func() {}}
	m.Update(asyncMsg{id: 9, kind: "settings-save", err: saveFailure})
	if err := configureRunResult(context.Background(), m, nil); !errors.Is(err, saveFailure) {
		t.Fatalf("save failure was not returned: %v", err)
	}
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	if err := configureRunResult(cancelled, m, nil); !errors.Is(err, context.Canceled) {
		t.Fatalf("context cancellation was not returned: %v", err)
	}
}

func TestFailedClientActionsRemainSelectedForRetry(t *testing.T) {
	s := newSettingsState(config.Default())
	clients := []install.Status{
		{Client: install.Client{ID: "register", Name: "Register"}, Configured: false},
		{Client: install.Client{ID: "unregister", Name: "Unregister"}, Configured: true},
	}
	s.applySave(settingsSavePayload{
		clients: clients,
		failed:  2,
		results: []clientApplyResult{
			{client: clients[0].Client, action: "registered", err: errors.New("register failed")},
			{client: clients[1].Client, action: "unregistered", err: errors.New("unregister failed")},
		},
	})
	if !s.selected[0] || s.selected[1] || !s.dirty {
		t.Fatalf("failed actions were not preserved for retry: selected=%v dirty=%t", s.selected, s.dirty)
	}
}

func TestRestorePreservesPendingClientsAndReportsFailure(t *testing.T) {
	m := testRoot(t, contextSettings, nil)
	m.configureOnly = true
	m.settings.clients = []install.Status{{Client: install.Client{ID: "pending"}}}
	m.settings.selected = []bool{true}
	m.settings.original = []bool{false}
	m.settings.dirty = true
	m.active = &operation{id: 12, cancel: func() {}}
	m.Update(asyncMsg{id: 12, kind: "restore", value: config.Default()})
	if !m.settings.dirty || m.configureSaved {
		t.Fatalf("restore discarded pending client registration: dirty=%t saved=%t", m.settings.dirty, m.configureSaved)
	}

	restoreErr := errors.New("restore failed")
	m.active = &operation{id: 13, cancel: func() {}}
	m.Update(asyncMsg{id: 13, kind: "restore", err: restoreErr})
	if !errors.Is(m.runErr, restoreErr) {
		t.Fatalf("restore failure was not retained: %v", m.runErr)
	}
}

func TestPartialRuntimeBootstrapAndDiagnosticsAreNilSafe(t *testing.T) {
	m, closeFixture := documentationFixture(t)
	defer closeFixture()
	m.runtime.Sessions, m.runtime.Cache, m.runtime.HTTP = nil, nil, nil
	m.docs.apis = nil
	runTeaCommand(t, m, m.Init())
	if len(m.docs.apis) == 0 || m.severity != severityWarning || !strings.Contains(m.status, "session service") {
		t.Fatalf("partial bootstrap discarded success or warning: apis=%d status=%q", len(m.docs.apis), m.status)
	}
	checks := diagnose(m.runtime, Options{})
	if len(checks) == 0 {
		t.Fatal("partial diagnostics returned no checks")
	}
	_ = m.View()
	if got := diagnose(nil, Options{}); len(got) != 1 || got[0].OK {
		t.Fatalf("nil runtime diagnostics = %#v", got)
	}
}

func TestCookieDeleteRefreshesAllSessionModelsAndClamps(t *testing.T) {
	manager, err := sessions.New(t.TempDir(), time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	handle, err := manager.Create(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	parsed, _ := url.Parse("https://example.test/path")
	handle.Jar().SetCookies(parsed, []*http.Cookie{{Name: "token", Value: "secret", Path: "/"}})
	if err := handle.Close(); err != nil {
		t.Fatal(err)
	}
	runtime := &bootstrap.Runtime{Config: config.Default(), Sessions: manager}
	m := testRoot(t, contextSessions, runtime)
	items, _ := manager.List()
	m.sessions.setItems(items)
	m.request.setSessions(items, items[0].ID)
	runTeaCommand(t, m, m.inspectSessionCmd(items[0].ID))
	m.sessions.cookieCursor = 0
	m.confirmCookieDelete()
	_, cmd := m.Update(key("y"))
	runTeaCommand(t, m, cmd)
	if len(m.sessions.inspection.Cookies) != 0 || len(m.sessions.items) != 1 || m.sessions.items[0].CookieCount != 0 || len(m.request.sessions) != 1 {
		t.Fatalf("cookie deletion did not refresh all models: inspection=%#v list=%#v request=%#v", m.sessions.inspection, m.sessions.items, m.request.sessions)
	}
	if m.sessions.cookieCursor != 0 || m.sessions.cursor != 0 || m.request.sessionCursor != 1 {
		t.Fatalf("refreshed cursors were not clamped: session=%d cookie=%d request=%d", m.sessions.cursor, m.sessions.cookieCursor, m.request.sessionCursor)
	}
}

func TestSearchHitPositioningCollectionsWindowAndRepaint(t *testing.T) {
	m := testRoot(t, contextDocumentation, nil)
	var markdown strings.Builder
	for line := 1; line <= 80; line++ {
		markdown.WriteString("line content\n")
	}
	m.docs.openReader(library.ReadResult{PageID: "page", Title: "Page", Markdown: markdown.String(), TotalLines: 80, Lines: [2]int{1, 80}}, 70, 40, 8)
	if m.docs.reader.viewport.YOffset < 30 {
		t.Fatalf("search hit line was not positioned in reader: offset=%d", m.docs.reader.viewport.YOffset)
	}

	for index := 0; index < 30; index++ {
		m.docs.collections = append(m.docs.collections, library.Collection{Name: fmt.Sprintf("collection-%02d", index), APICount: 1})
	}
	m.docs.collectionCursor = 29
	m.Update(tea.WindowSizeMsg{Width: 65, Height: 20})
	bottom := m.View()
	if !strings.Contains(bottom, "collection-28") || strings.Contains(bottom, "All documentation") {
		t.Fatal("collection pane did not scroll to cursor")
	}
	m.docs.collectionCursor = 0
	top := m.View()
	if !strings.Contains(top, "All documentation") || strings.Contains(top, "collection-28") {
		t.Fatal("collection repaint retained stale rows")
	}
}

func TestMutationSummaryDuplicateHeadersAndResponseMetadata(t *testing.T) {
	parsed, err := parseHeaders("X-Test: one\nX-Test: two")
	if err != nil {
		t.Fatal(err)
	}
	headers, ok := parsed.(http.Header)
	if !ok || len(headers.Values("X-Test")) != 2 {
		t.Fatalf("duplicate headers were collapsed: %#v", parsed)
	}
	m := testRoot(t, contextRequest, nil)
	m.request.method.SetValue("POST")
	m.request.url.SetValue("https://example.test/\x1b]2;bad\x07")
	m.request.headers.SetValue("Authorization: secret-value")
	m.request.body.SetValue(`{"secret":"body-value"}`)
	m.request.sessions = []sessions.SessionInfo{{ID: "session-secret"}}
	m.request.sessionCursor = 1
	m.prepareRequest()
	modal := strings.Join(m.modalView(80, 12), "\n")
	for _, wanted := range []string{"headers: yes", "body: yes", "session: yes"} {
		if !strings.Contains(modal, wanted) {
			t.Fatalf("mutation summary missing %q: %s", wanted, modal)
		}
	}
	for _, secret := range []string{"secret-value", "body-value", "session-secret", "\x1b]2;bad"} {
		if strings.Contains(modal, secret) {
			t.Fatalf("mutation summary leaked %q", secret)
		}
	}

	m.modal = nil
	m.request.response.Height = 100
	m.request.applyResult(httpcall.Result{
		Request:   httpcall.RequestResult{ID: "request-7", Method: "GET", Endpoint: "https://example.test", AutomaticHeaders: map[string]string{"Accept": "application/json"}},
		Response:  httpcall.ResponseResult{Status: 200, State: "complete"},
		Attempts:  []httpcall.Attempt{{Number: 1, Error: "reset", RetryReason: "transient", RetryDelay: time.Second}},
		Redirects: []httpcall.Redirect{{Status: 302, FromURL: "https://one", ToURL: "https://two", MethodBefore: "POST", MethodAfter: "GET"}},
		Cache:     httpcall.CacheResult{Directory: "/cache", BodyPath: "/cache/body", HeadersPath: "/cache/headers", MetadataPath: "/cache/meta", ErrorPath: "/cache/error"},
	})
	content := m.request.response.View()
	for _, wanted := range []string{"request-7", "Automatic headers", "transient", "1s", "Redirects", "/cache/meta", "/cache/error"} {
		if !strings.Contains(content, wanted) {
			t.Fatalf("response detail missing %q:\n%s", wanted, content)
		}
	}
}

func TestWindowsHeaderPathAndFailedRequestOutcome(t *testing.T) {
	for _, path := range []string{`C:\headers.json`, `C:headers.json`, `\\?\C:\headers.json`, `\\server\share\headers.json`} {
		parsed, err := parseHeaders(path)
		if err != nil || parsed != path {
			t.Fatalf("Windows header path was parsed as inline headers: %#v (%v)", parsed, err)
		}
	}
	parsed, err := parseHeaders("X: value")
	headers, ok := parsed.(http.Header)
	if err != nil || !ok || headers.Get("X") != "value" {
		t.Fatalf("single-letter inline header was parsed as a path: %#v (%v)", parsed, err)
	}

	m := testRoot(t, contextRequest, nil)
	m.request.applyResult(httpcall.Result{Request: httpcall.RequestResult{ID: "old"}, Response: httpcall.ResponseResult{Status: 200}})
	m.active = &operation{id: 11, cancel: func() {}}
	callErr := errors.New("transport failed")
	partial := httpcall.Result{Request: httpcall.RequestResult{ID: "new", SessionID: "session-new", Method: "GET", Endpoint: "https://example.test"}, Attempts: []httpcall.Attempt{{Number: 1, Error: "reset"}}}
	items := []sessions.SessionInfo{{ID: "session-new"}}
	m.Update(asyncMsg{id: 11, kind: "request", value: requestPayload{result: partial, sessions: items, err: callErr}})
	if m.request.result == nil || m.request.result.Request.ID != "new" || !errors.Is(m.request.resultErr, callErr) {
		t.Fatalf("partial failed result was not displayed: result=%+v err=%v", m.request.result, m.request.resultErr)
	}
	if len(m.sessions.items) != 1 || m.request.sessionCursor != 1 || m.severity != severityError {
		t.Fatalf("failed request did not refresh session state: sessions=%v cursor=%d severity=%v", m.sessions.items, m.request.sessionCursor, m.severity)
	}
}

func TestSessionCookiesMatchCurrentSelection(t *testing.T) {
	s := newSessionState()
	s.setItems([]sessions.SessionInfo{{ID: "session-a"}, {ID: "session-b"}})
	s.setInspection(sessions.Inspection{Session: sessions.SessionInfo{ID: "session-a"}, Cookies: []sessions.CookieInfo{{Name: "cookie-a"}}})
	s.cursor = 1
	rows := strings.Join(s.cookieRows(10), "\n")
	if strings.Contains(rows, "cookie-a") || !strings.Contains(rows, "Press enter") {
		t.Fatalf("stale cookies were rendered for another selection: %s", rows)
	}
	s.cursor = 0
	s.setInspection(sessions.Inspection{Session: sessions.SessionInfo{ID: "session-a"}, Cookies: []sessions.CookieInfo{{Name: "stale"}}})
	s.setItems([]sessions.SessionInfo{{ID: "session-a", CookieCount: 2}})
	if s.inspection.Session.ID != "" || len(s.inspection.Cookies) != 0 {
		t.Fatalf("session refresh retained stale inspection: %+v", s.inspection)
	}
}

func operationCommand(t *testing.T, command tea.Cmd) tea.Cmd {
	t.Helper()
	batch, ok := command().(tea.BatchMsg)
	if !ok || len(batch) == 0 {
		t.Fatalf("operation command returned %T", command())
	}
	return batch[0]
}
