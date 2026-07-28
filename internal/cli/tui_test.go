package cli

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/sairaph/apis-mcp/internal/bootstrap"
	"github.com/sairaph/apis-mcp/internal/config"
	"github.com/sairaph/apis-mcp/internal/httpcall"
	"github.com/sairaph/apis-mcp/internal/install"
	"github.com/sairaph/apis-mcp/internal/sessions"
	"github.com/sairaph/apis-mcp/library"
)

func TestDocumentationHierarchySearchReadAndBackContext(t *testing.T) {
	m, closeSnapshot := documentationFixture(t)
	defer closeSnapshot()
	m.docs.focus = 1

	_, cmd := m.Update(key("enter"))
	runTeaCommand(t, m, cmd)
	if m.currentScreen() != screenDocumentBrowser || len(m.docs.frames) != 1 {
		t.Fatalf("catalog did not open browser: stack=%v frames=%d", m.stack, len(m.docs.frames))
	}
	rootFrame := m.docs.currentFrame()
	if len(rootFrame.entries) != 2 || rootFrame.entries[0].path == nil {
		t.Fatalf("unexpected root hierarchy: %#v", rootFrame.entries)
	}
	rootFrame.cursor = 0
	_, cmd = m.Update(key("enter"))
	runTeaCommand(t, m, cmd)
	if got := m.docs.currentFrame().path; got != "guides" {
		t.Fatalf("opened path %q, want guides", got)
	}
	m.setContext(contextRequest)
	m.setContext(contextDocumentation)
	if m.currentScreen() != screenDocumentBrowser || m.docs.currentFrame().path != "guides" {
		t.Fatal("switching contexts lost the documentation screen stack")
	}

	frame := m.docs.currentFrame()
	frame.search.SetValue("needle")
	runTeaCommand(t, m, m.searchDocumentCmd("needle", frame.path))
	if len(frame.results) != 1 || frame.results[0].Title != "Getting Started" {
		t.Fatalf("unexpected typed search results: %#v", frame.results)
	}
	runTeaCommand(t, m, m.readDocumentCmd(frame.results[0].PageID))
	if m.currentScreen() != screenDocumentReader || !strings.Contains(m.docs.reader.result.Markdown, "needle") {
		t.Fatalf("reader did not open search result: %#v", m.docs.reader.result)
	}
	m.Update(key("esc"))
	if m.currentScreen() != screenDocumentBrowser || m.docs.currentFrame().path != "guides" {
		t.Fatal("reader back-navigation lost the selected path")
	}
	m.Update(key("esc"))
	if len(m.docs.frames) != 1 || m.docs.currentFrame().path != "" || m.docs.currentFrame().cursor != 0 {
		t.Fatal("path back-navigation did not preserve the parent frame and cursor")
	}
}

func TestResponsiveViewsFitAndSwitchPaneCounts(t *testing.T) {
	m, closeSnapshot := documentationFixture(t)
	defer closeSnapshot()
	for _, size := range []struct {
		width, height int
		dividers      int
	}{{120, 30, 2}, {90, 25, 1}, {65, 20, 0}} {
		m.Update(tea.WindowSizeMsg{Width: size.width, Height: size.height})
		view := m.View()
		lines := strings.Split(view, "\n")
		if len(lines) != size.height {
			t.Fatalf("%dx%d view has %d rows", size.width, size.height, len(lines))
		}
		for row, line := range lines {
			if got := lipgloss.Width(line); got != size.width {
				t.Fatalf("%dx%d row %d has width %d: %q", size.width, size.height, row, got, line)
			}
		}
		bodyLine := lines[3]
		if got := strings.Count(bodyLine, "│") - 2; got != size.dividers {
			t.Fatalf("%d-column layout has %d interior dividers, want %d", size.width, got, size.dividers)
		}
	}
	m.Update(tea.WindowSizeMsg{Width: 59, Height: 19})
	if view := m.View(); !strings.Contains(view, "needs at least 60x20") {
		t.Fatalf("small terminal did not show size warning: %q", view)
	}
}

func TestRequestInputsAreCursorAwareAndResponseScrolls(t *testing.T) {
	m := testRoot(t, contextRequest, nil)
	r := m.request
	r.url.SetValue("ac")
	r.focusField(1)
	m.Update(key("left"))
	m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'b'}})
	if got := r.url.Value(); got != "abc" {
		t.Fatalf("cursor insertion produced %q, want abc", got)
	}
	m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'界'}})
	if got := r.url.Value(); got != "ab界c" {
		t.Fatalf("unicode insertion produced %q", got)
	}
	r.focusField(4)
	m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'{'}})
	m.Update(key("enter"))
	m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'}'}})
	if got := r.body.Value(); got != "{\n}" {
		t.Fatalf("multiline body editing produced %q", got)
	}
	r.method.SetValue("POST")
	r.url.SetValue("https://example.test")
	if cmd := m.prepareRequest(); cmd != nil || m.modal == nil || m.modal.yes {
		t.Fatal("mutating request did not open a safe-default confirmation")
	}
	m.Update(key("enter"))

	var preview strings.Builder
	for index := 0; index < 100; index++ {
		fmt.Fprintf(&preview, "line %03d payload\n", index)
	}
	r.applyResult(httpcall.Result{
		Request:  httpcall.RequestResult{Method: "GET", Endpoint: "https://example.test", SessionID: "session"},
		Response: httpcall.ResponseResult{Status: 200, State: "complete"},
		Preview:  &httpcall.Preview{Kind: "text", Content: preview.String()},
	})
	r.response.Width, r.response.Height = 30, 5
	before := r.response.YOffset
	m.Update(key("down"))
	if r.response.YOffset <= before {
		t.Fatal("response viewport did not scroll")
	}
}

func TestAsyncOperationIDsAndCancellation(t *testing.T) {
	m := testRoot(t, contextDocumentation, nil)
	first := m.startOperation("first", "first", true, func(context.Context, int64) tea.Msg { return nil })
	if first == nil {
		t.Fatal("first operation returned no command")
	}
	_ = operationCommand(t, first)()
	firstID := m.active.id
	m.cancelActive("cancel")
	second := m.startOperation("second", "second", true, func(context.Context, int64) tea.Msg { return nil })
	_ = operationCommand(t, second)()
	secondID := m.active.id
	m.Update(asyncMsg{id: firstID, kind: "cache-clean", value: "stale"})
	if m.active == nil || m.active.id != secondID || m.status == "stale" {
		t.Fatal("stale operation result was accepted")
	}
	m.Update(key("ctrl+c"))
	if m.active != nil || m.quit {
		t.Fatal("ctrl+c did not cancel active work before quitting")
	}
}

func TestSessionsCreateInspectAndSafeDelete(t *testing.T) {
	manager, err := sessions.New(t.TempDir(), time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	runtime := &bootstrap.Runtime{Config: config.Default(), Sessions: manager}
	m := testRoot(t, contextSessions, runtime)
	runTeaCommand(t, m, m.createSessionCmd())
	if len(m.sessions.items) != 1 {
		t.Fatalf("create did not refresh typed session list: %#v", m.sessions.items)
	}
	id := m.sessions.items[0].ID
	runTeaCommand(t, m, m.inspectSessionCmd(id))
	if m.currentScreen() != screenSessionDetail || m.sessions.inspection.Session.ID != id {
		t.Fatal("session inspection did not open detail screen")
	}
	m.Update(key("d"))
	if m.modal == nil || m.modal.yes {
		t.Fatal("delete confirmation did not default to cancel")
	}
	m.Update(key("enter"))
	if _, err := manager.Inspect(id); err != nil {
		t.Fatalf("default confirmation deleted session: %v", err)
	}
	m.Update(key("d"))
	_, cmd := m.Update(key("y"))
	runTeaCommand(t, m, cmd)
	if len(m.sessions.items) != 0 {
		t.Fatalf("confirmed deletion did not refresh list: %#v", m.sessions.items)
	}
}

func TestSettingsDeselectionCallsUninstall(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	if _, err := install.Configure(install.Options{Home: home, Executable: "/bin/apis-mcp", ClientIDs: []string{"cursor"}}); err != nil {
		t.Fatal(err)
	}
	paths := config.Paths{Root: filepath.Join(home, ".apis-mcp"), Config: filepath.Join(home, ".apis-mcp", "config.toml")}
	runtime := &bootstrap.Runtime{Paths: paths, Config: config.Default()}
	m := testRoot(t, contextSettings, runtime)
	m.settings.setClients(install.Detect(home, ""), false)
	index := -1
	for candidate, status := range m.settings.clients {
		if status.Client.ID == "cursor" {
			index = candidate
			break
		}
	}
	if index < 0 || !m.settings.clients[index].Configured {
		t.Fatal("cursor fixture was not detected as configured")
	}
	m.settings.selected[index] = false
	runTeaCommand(t, m, m.saveSettingsCmd())
	for _, status := range install.Detect(home, "") {
		if status.Client.ID == "cursor" && status.Configured {
			t.Fatal("deselecting a configured client did not uninstall apis-mcp")
		}
	}
	found := false
	for _, result := range m.settings.results {
		found = found || result.client.ID == "cursor" && result.action == "unregistered" && result.changed && result.err == nil
	}
	if !found {
		t.Fatalf("per-client uninstall result missing: %#v", m.settings.results)
	}
}

func TestModalOwnsInputAndConfirmationDefaultsToCancel(t *testing.T) {
	m, closeSnapshot := documentationFixture(t)
	defer closeSnapshot()
	before := m.docs.collectionCursor
	m.openHelp()
	m.Update(key("down"))
	if m.docs.collectionCursor != before || m.modal == nil {
		t.Fatal("help modal leaked navigation to underlying context")
	}
	m.Update(key("esc"))
	called := false
	m.confirm("Danger", nil, "Proceed", func() tea.Cmd {
		called = true
		return nil
	})
	m.Update(key("enter"))
	if called || m.modal != nil {
		t.Fatal("confirmation did not default to cancel")
	}
}

func TestProgramUsesSingleSharedOptionAssembly(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	reader, writer := io.Pipe()
	defer reader.Close()
	go func() {
		for range 80 {
			time.Sleep(100 * time.Millisecond)
			if _, err := writer.Write([]byte("q")); err != nil {
				break
			}
		}
		_ = writer.Close()
	}()
	var output bytes.Buffer
	options := normalizeOptions(Options{Stdin: reader, Stdout: &output})
	m := newRootModel(ctx, nil, options, contextSettings)
	if got := len(teaProgramOptions(m, options)); got != 4 {
		t.Fatalf("program option count = %d, want context, alternate screen, input, output", got)
	}
	if _, err := newTeaProgram(m, options).Run(); err != nil {
		t.Fatal(err)
	}
	rendered := output.String()
	if !strings.Contains(rendered, "\x1b[?1049h") || strings.Contains(rendered, "\x1b[?1002h") || strings.Contains(rendered, "\x1b[?1003h") {
		t.Fatalf("program alternate-screen/mouse options are wrong: %q", rendered)
	}
}

func documentationFixture(t *testing.T) (*model, func()) {
	t.Helper()
	root := t.TempDir()
	writeTUIFixture(t, root, "test/v1/_index.md", "---\nname: Test API\nversion: v1\ndescription: Test documentation\ncollections: [developer_tools]\n---\n")
	writeTUIFixture(t, root, "test/v1/overview.md", "---\ntitle: Overview\ndescription: Root page\n---\n\n# Overview\n\nRoot content.\n")
	writeTUIFixture(t, root, "test/v1/guides/start.md", "---\ntitle: Getting Started\ndescription: Guide\n---\n\n# Getting Started\n\nA searchable needle appears here.\n")
	paths := config.Paths{Library: root, Index: t.TempDir(), Config: filepath.Join(t.TempDir(), "config.toml")}
	snapshot, err := library.Open(context.Background(), library.Options{
		UserRoot: root, IndexPath: filepath.Join(paths.Index, "library.sqlite"),
	})
	if err != nil {
		t.Fatal(err)
	}
	runtime := &bootstrap.Runtime{Paths: paths, Config: config.Default(), Library: snapshot}
	m := testRoot(t, contextDocumentation, runtime)
	m.docs.collections, m.docs.apis, err = loadCatalog(context.Background(), snapshot)
	if err != nil {
		snapshot.Close()
		t.Fatal(err)
	}
	for index, choice := range m.docs.choices() {
		if choice.name == "Test API" {
			m.docs.catalogCursor = index
			break
		}
	}
	return m, func() { _ = snapshot.Close() }
}

func testRoot(t *testing.T, initial humanContext, runtime *bootstrap.Runtime) *model {
	t.Helper()
	if runtime == nil {
		runtime = &bootstrap.Runtime{Paths: config.Paths{Config: filepath.Join(t.TempDir(), "config.toml")}, Config: config.Default()}
	}
	return newRootModel(context.Background(), runtime, Options{Executable: "/bin/apis-mcp"}, initial)
}

func writeTUIFixture(t *testing.T, root, relative, content string) {
	t.Helper()
	name := filepath.Join(root, filepath.FromSlash(relative))
	if err := os.MkdirAll(filepath.Dir(name), 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(name, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
}

func key(value string) tea.KeyMsg {
	switch value {
	case "enter":
		return tea.KeyMsg{Type: tea.KeyEnter}
	case "esc":
		return tea.KeyMsg{Type: tea.KeyEsc}
	case "left":
		return tea.KeyMsg{Type: tea.KeyLeft}
	case "down":
		return tea.KeyMsg{Type: tea.KeyDown}
	case "ctrl+c":
		return tea.KeyMsg{Type: tea.KeyCtrlC}
	default:
		return tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(value)}
	}
}

func runTeaCommand(t *testing.T, m *model, cmd tea.Cmd) {
	t.Helper()
	if cmd == nil {
		return
	}
	message := cmd()
	switch message := message.(type) {
	case tea.BatchMsg:
		for _, command := range message {
			runTeaCommand(t, m, command)
		}
	case asyncMsg:
		_, next := m.Update(message)
		runTeaCommand(t, m, next)
	case nil:
		return
	}
}
