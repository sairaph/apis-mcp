package cli

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/sairaph/apis-mcp/internal/bootstrap"
	"github.com/sairaph/apis-mcp/internal/config"
	"github.com/sairaph/apis-mcp/internal/install"
	"github.com/sairaph/apis-mcp/internal/sessions"
	"github.com/sairaph/apis-mcp/library"
)

type humanContext int

const (
	contextDocumentation humanContext = iota
	contextRequest
	contextSessions
	contextSettings
)

var contextNames = [...]string{"Documentation", "Request", "Sessions", "Settings"}

type screenID int

const (
	screenDocumentation screenID = iota
	screenDocumentBrowser
	screenDocumentReader
	screenDocumentImport
	screenRequest
	screenSessionList
	screenSessionDetail
	screenSettings
)

type severity int

const (
	severityInfo severity = iota
	severitySuccess
	severityWarning
	severityError
)

type operation struct {
	id          int64
	kind        string
	label       string
	cancellable bool
	cancel      context.CancelFunc
}

type commandTracker struct {
	mu      sync.Mutex
	running int
	closing bool
	idle    chan struct{}
}

func (t *commandTracker) begin() bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.closing {
		return false
	}
	if t.running == 0 {
		t.idle = make(chan struct{})
	}
	t.running++
	return true
}

func (t *commandTracker) done() {
	t.mu.Lock()
	t.running--
	if t.running == 0 {
		close(t.idle)
	}
	t.mu.Unlock()
}

func (t *commandTracker) closeAndWait() {
	t.mu.Lock()
	t.closing = true
	if t.running == 0 {
		t.mu.Unlock()
		return
	}
	idle := t.idle
	t.mu.Unlock()
	<-idle
}

type asyncMsg struct {
	id    int64
	kind  string
	value any
	err   error
}

type modalState struct {
	help    bool
	title   string
	details []string
	cancel  string
	confirm string
	yes     bool
	action  func() tea.Cmd
}

type bootstrapData struct {
	collections []library.Collection
	apis        []library.API
	sessions    []sessions.SessionInfo
	clients     []install.Status
	warnings    []error
}

// model is the single root pointer model for every human workflow.
type model struct {
	ctx     context.Context
	runtime *bootstrap.Runtime
	options Options
	paths   config.Paths

	context humanContext
	stack   []screenID
	stacks  [4][]screenID
	width   int
	height  int

	docs     *documentationState
	request  *requestState
	sessions *sessionState
	settings *settingsState

	modal    *modalState
	status   string
	severity severity
	spinner  spinner.Model
	nextID   int64
	active   *operation
	quit     bool

	configureOnly  bool
	configureSaved bool
	runErr         error

	commands         commandTracker
	snapMu           sync.Mutex
	pendingSnapshots map[int64]*library.Snapshot
}

func newModel(ctx context.Context, runtime *bootstrap.Runtime, options Options) *model {
	return newRootModel(ctx, runtime, options, contextDocumentation)
}

func newRootModel(ctx context.Context, runtime *bootstrap.Runtime, options Options, initial humanContext) *model {
	options = normalizeOptions(options)
	paths := config.Paths{}
	cfg := config.Default()
	if runtime != nil {
		paths, cfg = runtime.Paths, runtime.Config
	} else if resolved, err := config.DefaultPaths(); err == nil {
		paths = resolved
		if loaded, loadErr := config.Load(paths); loadErr == nil {
			cfg = loaded
		}
	}
	spin := spinner.New(spinner.WithSpinner(spinner.MiniDot), spinner.WithStyle(styleTitle))
	m := &model{
		ctx: ctx, runtime: runtime, options: options, paths: paths,
		width: 100, height: 30, spinner: spin,
		docs: newDocumentationState(runtime), request: newRequestState(),
		sessions: newSessionState(), settings: newSettingsState(cfg),
		status: "Loading workspace", severity: severityInfo,
		pendingSnapshots: make(map[int64]*library.Snapshot),
	}
	m.request.setLargeDownloadPolicy(cfg.AllowLargeDownload)
	m.setContext(initial)
	return m
}

func (m *model) Init() tea.Cmd {
	return m.startOperation("bootstrap", "Loading workspace", false, func(ctx context.Context, id int64) tea.Msg {
		data := bootstrapData{clients: install.Detect("", "")}
		if m.runtime != nil {
			if m.runtime.Library != nil {
				var err error
				data.collections, data.apis, err = loadCatalog(ctx, m.docs.snapshot)
				if err != nil {
					data.warnings = append(data.warnings, fmt.Errorf("documentation: %w", err))
				}
			} else {
				data.warnings = append(data.warnings, errors.New("documentation service is unavailable"))
			}
			if m.runtime.Sessions != nil {
				var err error
				data.sessions, err = m.runtime.Sessions.List()
				if err != nil {
					data.warnings = append(data.warnings, fmt.Errorf("sessions: %w", err))
				}
			} else {
				data.warnings = append(data.warnings, errors.New("session service is unavailable"))
			}
		}
		return asyncMsg{id: id, kind: "bootstrap", value: data}
	})
}

func (m *model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch message := message.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = message.Width, message.Height
		m.resize()
		return m, nil
	case spinner.TickMsg:
		if m.active == nil {
			return m, nil
		}
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(message)
		return m, cmd
	case asyncMsg:
		return m, m.handleAsync(message)
	case tea.KeyMsg:
		return m.handleKey(message)
	}
	return m, m.updateFocusedComponent(message)
}

func (m *model) handleKey(key tea.KeyMsg) (tea.Model, tea.Cmd) {
	if m.modal != nil {
		return m.handleModalKey(key)
	}
	if key.String() == "ctrl+c" {
		if m.active != nil {
			if !m.active.cancellable {
				m.status, m.severity = "Cannot cancel "+m.active.label+"; waiting for it to finish", severityWarning
				return m, nil
			}
			m.cancelActive("Operation cancelled")
			return m, nil
		}
		m.quit = true
		return m, tea.Quit
	}
	if m.active != nil {
		return m, nil
	}
	if !m.editing() {
		switch key.String() {
		case "?":
			m.openHelp()
			return m, nil
		case "1", "2", "3", "4":
			if m.configureOnly {
				return m, nil
			}
			m.setContext(humanContext(int(key.Runes[0] - '1')))
			return m, nil
		case "q":
			if len(m.stack) == 1 {
				m.quit = true
				return m, tea.Quit
			}
		}
	}
	switch m.context {
	case contextDocumentation:
		return m.updateDocumentation(key)
	case contextRequest:
		return m.updateRequest(key)
	case contextSessions:
		return m.updateSessions(key)
	case contextSettings:
		return m.updateSettings(key)
	}
	return m, nil
}

func (m *model) View() string {
	if m.width < minimumWidth || m.height < minimumHeight {
		return fmt.Sprintf("\n  Terminal is %dx%d.\n  apis-mcp needs at least %dx%d.\n", m.width, m.height, minimumWidth, minimumHeight)
	}
	interiorWidth := m.width - 2
	bodyHeight := m.height - 5
	body := m.bodyView(interiorWidth, bodyHeight)
	footer := m.footer()
	if m.modal != nil {
		body = m.modalView(interiorWidth, bodyHeight)
		footer = "arrows/tab choose  enter confirm  esc close"
	}
	lines := make([]string, 0, m.height-2)
	lines = append(lines, fixed(m.contextBar(), interiorWidth))
	lines = append(lines, body...)
	lines = append(lines, fixed(m.statusView(), interiorWidth))
	lines = append(lines, fixed(" "+styleDim.Render(footer), interiorWidth))
	for len(lines) < m.height-2 {
		lines = append(lines, strings.Repeat(" ", interiorWidth))
	}
	if len(lines) > m.height-2 {
		lines = lines[:m.height-2]
	}
	return fullFrame(m.width, m.height, "apis-mcp", lines)
}

func (m *model) bodyView(width, height int) []string {
	switch m.context {
	case contextDocumentation:
		return m.viewDocumentation(width, height)
	case contextRequest:
		return m.viewRequest(width, height)
	case contextSessions:
		return m.viewSessions(width, height)
	case contextSettings:
		return m.viewSettings(width, height)
	default:
		return make([]string, height)
	}
}

func (m *model) contextBar() string {
	if m.configureOnly {
		return styleSelect.Render(" Settings ")
	}
	var tabs []string
	for index, name := range contextNames {
		label := fmt.Sprintf(" %d %s ", index+1, name)
		if humanContext(index) == m.context {
			label = styleSelect.Render(label)
		} else {
			label = styleDim.Render(label)
		}
		tabs = append(tabs, label)
	}
	return strings.Join(tabs, " ")
}

func (m *model) statusView() string {
	text := m.status
	if m.active != nil {
		hint := "ctrl+c cancel"
		if !m.active.cancellable {
			hint = "cannot cancel; please wait"
		}
		text = m.spinner.View() + " " + m.active.label + "  " + styleDim.Render(hint)
	}
	if text == "" {
		text = "Ready"
	}
	style := styleDim
	switch m.severity {
	case severitySuccess:
		style = styleSuccess
	case severityWarning:
		style = styleWarning
	case severityError:
		style = styleError
	}
	return " " + style.Render(safeLine(text))
}

func (m *model) footer() string {
	switch m.context {
	case contextDocumentation:
		return m.documentationFooter()
	case contextRequest:
		return m.requestFooter()
	case contextSessions:
		return m.sessionsFooter()
	case contextSettings:
		return m.settingsFooter()
	default:
		return "1-4 contexts  ? help  q quit"
	}
}

func (m *model) setContext(next humanContext) {
	if m.configureOnly && next != contextSettings {
		return
	}
	if next < contextDocumentation || next > contextSettings {
		return
	}
	if len(m.stack) > 0 {
		m.stacks[m.context] = append([]screenID(nil), m.stack...)
	}
	m.context = next
	if saved := m.stacks[next]; len(saved) > 0 {
		m.stack = append([]screenID(nil), saved...)
	} else {
		switch next {
		case contextDocumentation:
			m.stack = []screenID{screenDocumentation}
		case contextRequest:
			m.stack = []screenID{screenRequest}
		case contextSessions:
			m.stack = []screenID{screenSessionList}
		case contextSettings:
			m.stack = []screenID{screenSettings}
		}
	}
	m.status, m.severity = "Ready", severityInfo
}

func (m *model) currentScreen() screenID { return m.stack[len(m.stack)-1] }

func (m *model) push(screen screenID) { m.stack = append(m.stack, screen) }

func (m *model) pop() bool {
	if len(m.stack) <= 1 {
		return false
	}
	m.stack = m.stack[:len(m.stack)-1]
	return true
}

func (m *model) startOperation(kind, label string, cancellable bool, run func(context.Context, int64) tea.Msg) tea.Cmd {
	if m.active != nil {
		m.status, m.severity = "Waiting for "+m.active.label+" to finish", severityWarning
		return nil
	}
	m.nextID++
	id := m.nextID
	ctx, cancel := context.WithCancel(m.ctx)
	m.active = &operation{id: id, kind: kind, label: label, cancellable: cancellable, cancel: cancel}
	m.status, m.severity = label, severityInfo
	command := func() tea.Msg {
		if !m.commands.begin() {
			return nil
		}
		defer m.commands.done()
		return run(ctx, id)
	}
	return tea.Batch(command, m.spinner.Tick)
}

func (m *model) accept(message asyncMsg) bool {
	if m.active == nil || message.id != m.active.id {
		if payload, ok := message.value.(reloadPayload); ok && payload.snapshot != nil {
			m.releaseSnapshot(message.id, payload.snapshot, true)
		}
		return false
	}
	m.active.cancel()
	m.active = nil
	return true
}

func (m *model) cancelActive(status string) {
	if m.active != nil {
		if !m.active.cancellable {
			m.status, m.severity = "Cannot cancel "+m.active.label+"; waiting for it to finish", severityWarning
			return
		}
		if m.active.kind == "pages" {
			m.rollbackLoadingFrame()
		}
		m.active.cancel()
		m.active = nil
	}
	m.status, m.severity = status, severityWarning
}

func (m *model) handleAsync(message asyncMsg) tea.Cmd {
	if !m.accept(message) {
		return nil
	}
	if message.err != nil {
		m.status, m.severity = message.err.Error(), severityError
		m.handleAsyncFailure(message)
		return nil
	}
	switch message.kind {
	case "bootstrap":
		data := message.value.(bootstrapData)
		m.docs.collections, m.docs.apis = data.collections, data.apis
		m.sessions.setItems(data.sessions)
		m.request.setSessions(data.sessions, "")
		m.settings.setClients(data.clients, m.configureOnly)
		if len(data.warnings) > 0 {
			m.status, m.severity = errors.Join(data.warnings...).Error(), severityWarning
		} else {
			m.status, m.severity = "Workspace ready", severitySuccess
		}
	case "pages":
		payload := message.value.(pagesPayload)
		m.docs.applyPages(payload)
		m.status, m.severity = fmt.Sprintf("%d items in %s", len(payload.entries), displayPath(payload.path)), severitySuccess
	case "search":
		result := message.value.(library.SearchResult)
		frame := m.docs.currentFrame()
		frame.searching = false
		frame.results = result.Hits
		frame.resultCursor = 0
		frame.focus = 1
		m.status, m.severity = fmt.Sprintf("%d search matches", result.Total), severitySuccess
	case "read":
		payload := message.value.(readPayload)
		m.docs.openReader(payload.result, payload.targetLine, m.width-4, m.height-7)
		m.push(screenDocumentReader)
		m.status, m.severity = "Documentation page opened", severitySuccess
	case "reload", "import":
		payload := message.value.(reloadPayload)
		m.releaseSnapshot(message.id, payload.snapshot, false)
		m.applyReload(payload)
	case "request":
		payload := message.value.(requestPayload)
		m.request.applyOutcome(payload.result, payload.err)
		if payload.sessionsErr == nil {
			m.request.setSessions(payload.sessions, payload.result.Request.SessionID)
			m.sessions.setItems(payload.sessions)
		}
		m.status, m.severity = m.request.resultStatus(payload.err)
		if payload.sessionsErr != nil {
			m.status += "; session refresh failed: " + payload.sessionsErr.Error()
			if payload.err == nil {
				m.severity = severityWarning
			}
		}
	case "sessions-list":
		items := message.value.([]sessions.SessionInfo)
		m.sessions.setItems(items)
		m.request.setSessions(items, "")
		m.sessions.clamp()
		m.status, m.severity = fmt.Sprintf("%d cookie sessions", len(items)), severitySuccess
	case "session-inspect":
		payload := message.value.(sessionInspectionPayload)
		if payload.warning == nil {
			m.sessions.setItems(payload.items)
			m.request.setSessions(payload.items, "")
		}
		m.sessions.setInspection(payload.inspection)
		if m.currentScreen() != screenSessionDetail {
			m.push(screenSessionDetail)
		}
		m.status, m.severity = "Session details loaded", severitySuccess
		if payload.warning != nil {
			m.status, m.severity = "Session loaded; refresh warning: "+payload.warning.Error(), severityWarning
		}
	case "session-action":
		payload := message.value.(sessionActionPayload)
		if payload.warning == nil {
			m.sessions.setItems(payload.items)
			m.request.setSessions(payload.items, "")
		}
		m.status, m.severity = payload.message, severitySuccess
		if payload.resetDetail {
			m.sessions.inspection = sessions.Inspection{}
			m.stacks[contextSessions] = []screenID{screenSessionList}
			if m.context == contextSessions {
				m.stack = []screenID{screenSessionList}
			}
		}
		if payload.warning != nil {
			m.status, m.severity = payload.message+"; refresh warning: "+payload.warning.Error(), severityWarning
		}
	case "settings-save":
		payload := message.value.(settingsSavePayload)
		m.settings.applySave(payload)
		if m.runtime != nil {
			m.runtime.Config = m.settings.cfg
		}
		m.request.setLargeDownloadPolicy(m.settings.cfg.AllowLargeDownload)
		m.status, m.severity = payload.message, severitySuccess
		if payload.failed > 0 {
			m.severity = severityWarning
		}
		m.configureSaved = payload.failed == 0
		m.runErr = payload.err
	case "settings-clients":
		m.settings.setClients(message.value.([]install.Status), m.configureOnly)
		m.status, m.severity = "Client registrations refreshed", severitySuccess
	case "restore":
		m.settings.cfg = message.value.(config.Config)
		m.settings.dirty = m.settings.clientsDirty()
		m.configureSaved = !m.settings.dirty
		m.runErr = nil
		if m.runtime != nil {
			m.runtime.Config = m.settings.cfg
		}
		m.request.setLargeDownloadPolicy(m.settings.cfg.AllowLargeDownload)
		m.status, m.severity = "Recommended defaults restored; restart client connections to apply runtime limits", severitySuccess
	case "cache-clean":
		m.status, m.severity = message.value.(string), severitySuccess
	case "doctor":
		m.settings.diagnostics = message.value.([]Diagnostic)
		m.status, m.severity = "Diagnostics refreshed", severitySuccess
	}
	return nil
}

func (m *model) handleAsyncFailure(message asyncMsg) {
	if m.configureOnly && (message.kind == "settings-save" || message.kind == "restore") {
		m.configureSaved = false
		m.runErr = message.err
	}
	if message.kind == "pages" && len(m.docs.frames) > 0 && m.docs.currentFrame().loading {
		m.docs.frames = m.docs.frames[:len(m.docs.frames)-1]
		m.pop()
	}
	if message.kind == "search" && len(m.docs.frames) > 0 {
		m.docs.currentFrame().searching = false
	}
}

func (m *model) rollbackLoadingFrame() {
	if len(m.docs.frames) > 0 && m.docs.currentFrame().loading {
		m.docs.frames = m.docs.frames[:len(m.docs.frames)-1]
		m.pop()
	}
}

func (m *model) loadSessionsCmd() tea.Cmd {
	if m.runtime == nil || m.runtime.Sessions == nil {
		m.status, m.severity = "Session service is unavailable", severityWarning
		return nil
	}
	manager := m.runtime.Sessions
	return m.startOperation("sessions-list", "Refreshing sessions", false, func(_ context.Context, id int64) tea.Msg {
		items, err := manager.List()
		return asyncMsg{id: id, kind: "sessions-list", value: items, err: err}
	})
}

func (m *model) openHelp() {
	details := []string{
		"tab / shift+tab changes focus; arrows or j/k navigate lists.",
		"enter opens or confirms; esc returns without discarding parent context.",
		"While typing, digits and ? are text; press esc first to use navigation keys.",
		"ctrl+c cancels cancellable work; durable writes must finish safely.",
		"q quits only at a context root; ? opens and closes this help.",
	}
	if !m.configureOnly {
		details = append([]string{"1-4 switch human contexts; each context keeps its working state."}, details...)
	}
	m.modal = &modalState{help: true, title: "Keyboard help", details: details}
}

func (m *model) confirm(title string, details []string, confirm string, action func() tea.Cmd) {
	m.modal = &modalState{title: title, details: details, cancel: "Cancel", confirm: confirm, action: action}
}

func (m *model) handleModalKey(key tea.KeyMsg) (tea.Model, tea.Cmd) {
	if m.modal.help {
		switch key.String() {
		case "?", "esc", "q", "enter":
			m.modal = nil
		}
		return m, nil
	}
	switch key.String() {
	case "left", "right", "up", "down", "h", "j", "k", "l", "tab", "shift+tab":
		m.modal.yes = !m.modal.yes
	case "y":
		action := m.modal.action
		m.modal = nil
		return m, action()
	case "n", "esc", "q", "ctrl+c":
		m.modal = nil
		m.status, m.severity = "Cancelled", severityWarning
	case "enter":
		if !m.modal.yes {
			m.modal = nil
			m.status, m.severity = "Cancelled", severityWarning
			return m, nil
		}
		action := m.modal.action
		m.modal = nil
		return m, action()
	}
	return m, nil
}

func (m *model) modalView(width, height int) []string {
	rows := make([]string, 0, height)
	rows = append(rows, "", "  "+styleWarning.Render(safeLine(m.modal.title)), "")
	for _, detail := range m.modal.details {
		for _, line := range wrapLines(safeMultiline(detail), max(10, width-6)) {
			rows = append(rows, "  "+styleDim.Render(line))
		}
	}
	if !m.modal.help {
		rows = append(rows, "")
		cancel, confirm := "  "+safeLine(m.modal.cancel), "  "+styleError.Render(safeLine(m.modal.confirm))
		if m.modal.yes {
			confirm = styleSelect.Render(fixed(" "+m.modal.confirm, max(10, len(m.modal.confirm)+4)))
		} else {
			cancel = styleSelect.Render(fixed(" "+m.modal.cancel, max(10, len(m.modal.cancel)+4)))
		}
		rows = append(rows, "  "+cancel, "  "+confirm)
	}
	for len(rows) < height {
		rows = append(rows, "")
	}
	return rows[:height]
}

func (m *model) resize() {
	m.docs.resize(m.width, m.height)
	m.request.resize(m.width, m.height)
	m.settings.resize(m.width)
}

func (m *model) editing() bool {
	switch m.context {
	case contextDocumentation:
		return m.docs.editing()
	case contextRequest:
		return m.request.editing()
	case contextSettings:
		return m.settings.editing
	}
	return false
}

func (m *model) close() {
	if m.active != nil {
		if m.active.cancellable {
			m.active.cancel()
		}
	}
	m.commands.closeAndWait()
	m.active = nil
	m.closePendingSnapshots()
	m.docs.close()
}

func (m *model) registerSnapshot(id int64, snapshot *library.Snapshot) {
	if snapshot == nil {
		return
	}
	m.snapMu.Lock()
	m.pendingSnapshots[id] = snapshot
	m.snapMu.Unlock()
}

func (m *model) releaseSnapshot(id int64, snapshot *library.Snapshot, closeSnapshot bool) {
	m.snapMu.Lock()
	if m.pendingSnapshots[id] == snapshot {
		delete(m.pendingSnapshots, id)
	}
	m.snapMu.Unlock()
	if closeSnapshot && snapshot != nil {
		_ = snapshot.Close()
	}
}

func (m *model) closePendingSnapshots() {
	m.snapMu.Lock()
	snapshots := make([]*library.Snapshot, 0, len(m.pendingSnapshots))
	for id, snapshot := range m.pendingSnapshots {
		snapshots = append(snapshots, snapshot)
		delete(m.pendingSnapshots, id)
	}
	m.snapMu.Unlock()
	for _, snapshot := range snapshots {
		_ = snapshot.Close()
	}
}

func (m *model) updateFocusedComponent(message tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	switch m.context {
	case contextDocumentation:
		if m.docs.filtering {
			m.docs.filter, cmd = m.docs.filter.Update(message)
		} else if m.docs.importer != nil && m.docs.importer.focus < len(m.docs.importer.inputs) {
			index := m.docs.importer.focus
			m.docs.importer.inputs[index], cmd = m.docs.importer.inputs[index].Update(message)
		} else if len(m.docs.frames) > 0 && m.docs.currentFrame().searching {
			frame := m.docs.currentFrame()
			frame.search, cmd = frame.search.Update(message)
		}
	case contextRequest:
		r := m.request
		switch r.focus {
		case 0:
			r.method, cmd = r.method.Update(message)
		case 1:
			r.url, cmd = r.url.Update(message)
		case 3:
			r.headers, cmd = r.headers.Update(message)
		case 4:
			r.body, cmd = r.body.Update(message)
		case 5:
			r.timeout, cmd = r.timeout.Update(message)
		case 6:
			r.retries, cmd = r.retries.Update(message)
		case 7:
			r.jsonPath, cmd = r.jsonPath.Update(message)
		}
	case contextSettings:
		if m.settings.editing {
			m.settings.input, cmd = m.settings.input.Update(message)
		}
	}
	return cmd
}

func fullFrame(width, height int, title string, lines []string) string {
	interior := max(1, width-2)
	body := max(1, height-2)
	for len(lines) < body {
		lines = append(lines, "")
	}
	if len(lines) > body {
		lines = lines[:body]
	}
	label := " " + styleTitle.Render(title) + " "
	used := lipgloss.Width(label)
	var out strings.Builder
	out.WriteString(styleFrame.Render(border.TopLeft))
	out.WriteString(label)
	out.WriteString(styleFrame.Render(strings.Repeat(border.Top, max(0, interior-used)) + border.TopRight))
	out.WriteByte('\n')
	for _, line := range lines {
		out.WriteString(styleFrame.Render(border.Left))
		out.WriteString(fixed(line, interior))
		out.WriteString(styleFrame.Render(border.Right))
		out.WriteByte('\n')
	}
	out.WriteString(styleFrame.Render(border.BottomLeft + strings.Repeat(border.Bottom, interior) + border.BottomRight))
	return out.String()
}

func loadCatalog(ctx context.Context, snapshot *library.Snapshot) ([]library.Collection, []library.API, error) {
	if snapshot == nil {
		return nil, nil, errors.New("documentation library is unavailable")
	}
	var collections []library.Collection
	for page := 1; ; page++ {
		result, err := snapshot.Collections(ctx, library.CollectionsRequest{Page: page})
		if err != nil {
			return nil, nil, err
		}
		collections = append(collections, result.Collections...)
		if page >= result.TotalPages {
			break
		}
	}
	var apis []library.API
	for page := 1; ; page++ {
		result, err := snapshot.List(ctx, library.ListRequest{Page: page})
		if err != nil {
			return nil, nil, err
		}
		apis = append(apis, result.APIs...)
		if page >= result.TotalPages {
			break
		}
	}
	return collections, apis, nil
}

func libraryOptions(runtime *bootstrap.Runtime) library.Options {
	return runtime.LibraryOptions()
}

// RunInteractive starts one alternate-screen Bubble Tea application.
func RunInteractive(ctx context.Context, runtime *bootstrap.Runtime, options Options) error {
	if runtime == nil {
		return errors.New("runtime is required")
	}
	options = normalizeOptions(options)
	root := newModel(ctx, runtime, options)
	_, err := runTUIProgram(ctx, root, options, true)
	root.close()
	return err
}

func newTeaProgram(root *model, options Options) *tea.Program {
	return newTUIProgram(root.ctx, root, options, true)
}

func teaProgramOptions(root *model, options Options) []tea.ProgramOption {
	return tuiProgramOptions(root.ctx, options, true)
}

func runTUIProgram(ctx context.Context, root tea.Model, options Options, alternateScreen bool) (tea.Model, error) {
	return newTUIProgram(ctx, root, options, alternateScreen).Run()
}

func newTUIProgram(ctx context.Context, root tea.Model, options Options, alternateScreen bool) *tea.Program {
	return tea.NewProgram(root, tuiProgramOptions(ctx, options, alternateScreen)...)
}

func tuiProgramOptions(ctx context.Context, options Options, alternateScreen bool) []tea.ProgramOption {
	programOptions := []tea.ProgramOption{
		tea.WithContext(ctx),
		tea.WithInput(options.Stdin),
		tea.WithOutput(options.Stdout),
	}
	if alternateScreen {
		programOptions = append(programOptions, tea.WithAltScreen())
	}
	return programOptions
}

func displayPath(path string) string {
	if path == "" {
		return "root"
	}
	return path
}
