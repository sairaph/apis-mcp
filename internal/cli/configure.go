package cli

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync/atomic"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/ansi"
	"github.com/sairaph/apis-mcp/internal/config"
	"github.com/sairaph/apis-mcp/internal/docpacks"
	"github.com/sairaph/apis-mcp/internal/install"
	"github.com/sairaph/apis-mcp/library"
)

type configSetting struct {
	key      string
	label    string
	boolean  bool
	readOnly bool
}

var configSettings = []configSetting{
	{key: "version", label: "Config version", readOnly: true},
	{key: "list_token_budget", label: "List token budget"},
	{key: "read_token_budget", label: "Read token budget"},
	{key: "response_size_limit", label: "Response size limit (bytes)"},
	{key: "allow_large_download", label: "Allow large downloads", boolean: true},
	{key: "free_disk_reserve", label: "Free disk reserve (bytes)"},
	{key: "maximum_redirects", label: "Maximum redirects"},
	{key: "maximum_retries", label: "Maximum retries"},
	{key: "maximum_header_timeout_seconds", label: "Header timeout (seconds)"},
	{key: "background_after_seconds", label: "Background after (seconds)"},
	{key: "stalled_download_seconds", label: "Stalled download (seconds)"},
	{key: "retention_hours", label: "Retention (hours)"},
	{key: "tls_verify", label: "Verify TLS certificates", boolean: true},
}

type setupStep int

const (
	setupClients setupStep = iota
	setupAPIs
	setupSummary
	setupSettings
	setupApplying
	setupDone
)

type setupPacksRefreshedMsg struct {
	id      int
	catalog docpacks.Catalog
	err     error
}

type setupPackApplyResult struct {
	attempted bool
	changed   bool
	selected  int
	bytes     int64
	err       error
}

type setupAppliedMsg struct {
	id      int
	results []clientApplyResult
	packs   setupPackApplyResult
	err     error
}

type setupApplyProgressMsg struct {
	id    int
	event docpacks.ApplyEvent
}

type setupApplyPhase int

const (
	setupApplyPreparing setupApplyPhase = iota
	setupApplyIndexing
	setupApplyPublishing
	setupApplySaving
	setupApplyRegistering
	setupApplyFinishing
)

type setupApplyPhaseMsg struct {
	id    int
	phase setupApplyPhase
}

type setupApplyClosedMsg struct{ id int }

type setupModel struct {
	ctx     context.Context
	options Options
	paths   config.Paths

	step             setupStep
	clients          []install.Status
	selected         []bool
	cursor           int
	showAll          bool
	settings         config.Config
	originalSettings config.Config
	settingCursor    int
	editing          bool
	input            string
	choice           bool
	width            int
	height           int

	packManager      *docpacks.Manager
	packActive       map[string]docpacks.Pack
	packSelected     map[string]bool
	packUnlisted     map[string]bool
	packCatalog      docpacks.Catalog
	packCatalogOK    bool
	packLoading      bool
	packError        error
	packMessage      string
	packRefreshID    int
	packRefreshStop  context.CancelFunc
	packCursor       int
	packScrollRow    int
	packApplyResult  setupPackApplyResult
	packApplyEvents  map[string]docpacks.ApplyEvent
	packApplyOrder   []string
	packApplyStage   docpacks.ApplyStage
	packApplyID      int
	packApplyStop    context.CancelFunc
	packApplyDone    <-chan struct{}
	packApplyUpdates <-chan tea.Msg
	packApplyCanStop *atomic.Bool
	packApplyPhase   setupApplyPhase
	packApplyChanged bool
	packRemovalOnly  bool

	results   []clientApplyResult
	message   string
	failure   error
	cancelled bool
	saved     bool
}

// RunConfigure uses the same unframed, one-step-at-a-time flow as the
// interactive-terminal-mcp installer.
func RunConfigure(ctx context.Context, options Options) error {
	options = normalizeOptions(options)
	paths, err := config.DefaultPaths()
	if err != nil {
		return err
	}
	cfg, err := config.Load(paths)
	if err != nil {
		return err
	}
	manager, err := docpacks.Open(paths.Packs, docpacks.Options{})
	if err != nil {
		return fmt.Errorf("open API packs: %w", err)
	}
	state, err := newSetupModelWithPacks(ctx, options, paths, cfg, install.Detect("", ""), manager)
	if err != nil {
		return fmt.Errorf("open active API packs: %w", err)
	}
	program := tea.NewProgram(state, tea.WithContext(ctx), tea.WithInput(options.Stdin), tea.WithOutput(options.Stdout))
	_, err = program.Run()
	return setupRunResult(ctx, state, err)
}

func setupRunResult(ctx context.Context, state *setupModel, programErr error) error {
	if programErr != nil && state.packApplyStop != nil {
		state.packApplyStop()
	}
	if state.packApplyDone != nil {
		<-state.packApplyDone
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	if programErr != nil {
		return programErr
	}
	if state.failure != nil {
		return state.failure
	}
	if state.cancelled || !state.saved {
		return errors.New("configuration cancelled before a successful save")
	}
	return nil
}

func newSetupModel(ctx context.Context, options Options, paths config.Paths, cfg config.Config, clients []install.Status) *setupModel {
	state, _ := newSetupModelWithPacks(ctx, options, paths, cfg, clients, nil)
	return state
}

func newSetupModelWithPacks(ctx context.Context, options Options, paths config.Paths, cfg config.Config, clients []install.Status, manager *docpacks.Manager) (*setupModel, error) {
	state := &setupModel{
		ctx: ctx, options: options, paths: paths,
		clients: clients, selected: make([]bool, len(clients)), settings: cfg, originalSettings: cfg,
		width: 100, height: 30, packManager: manager, packCursor: -1,
		packActive: make(map[string]docpacks.Pack), packSelected: make(map[string]bool), packUnlisted: make(map[string]bool),
	}
	for index, status := range clients {
		state.selected[index] = status.Configured || status.Detected && status.Err == nil
	}
	if manager != nil {
		active, err := manager.Active()
		if err != nil {
			return nil, err
		}
		for id, pack := range active.Packs {
			state.packActive[id] = pack
			state.packSelected[id] = true
		}
	}
	state.cursor = state.firstSelectable()
	return state, nil
}

func (m *setupModel) Init() tea.Cmd { return nil }

func (m *setupModel) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch message := message.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = message.Width, message.Height
		m.ensurePackVisible()
		return m, nil
	case setupPacksRefreshedMsg:
		if message.id != m.packRefreshID {
			return m, nil
		}
		if m.packRefreshStop != nil {
			m.packRefreshStop()
			m.packRefreshStop = nil
		}
		m.packLoading, m.packError = false, message.err
		if message.err == nil {
			m.replacePackCatalog(message.catalog)
		}
		return m, nil
	case setupAppliedMsg:
		if message.id != m.packApplyID {
			return m, nil
		}
		m.packApplyStop = nil
		m.packApplyUpdates = nil
		if m.cancelled {
			return m, tea.Quit
		}
		m.results, m.packApplyResult, m.failure, m.saved, m.step = message.results, message.packs, message.err, message.err == nil, setupDone
		return m, nil
	case setupApplyProgressMsg:
		if message.id != m.packApplyID || m.step != setupApplying {
			return m, nil
		}
		m.packApplyStage = message.event.Stage
		if message.event.Stage == docpacks.ApplyStageIndexing {
			m.packApplyPhase = setupApplyIndexing
		} else if message.event.Stage == docpacks.ApplyStageApplying {
			m.packApplyPhase = setupApplyPublishing
		}
		if message.event.PackID != "" {
			m.packApplyEvents[message.event.PackID] = message.event
		}
		return m, waitSetupApply(m.packApplyID, m.packApplyUpdates, m.ctx)
	case setupApplyPhaseMsg:
		if message.id != m.packApplyID || m.step != setupApplying {
			return m, nil
		}
		m.packApplyPhase = message.phase
		return m, waitSetupApply(m.packApplyID, m.packApplyUpdates, m.ctx)
	case setupApplyClosedMsg:
		if message.id != m.packApplyID {
			return m, nil
		}
		m.packApplyStop = nil
		m.packApplyUpdates = nil
		if m.cancelled || m.ctx.Err() != nil {
			return m, tea.Quit
		}
		m.failure = errors.New("configuration apply worker ended without a result")
		m.step = setupDone
		return m, nil
	case tea.KeyMsg:
		if message.String() == "ctrl+c" {
			if m.step == setupApplying && m.packApplyStop != nil {
				if m.packApplyCanStop != nil && m.packApplyCanStop.CompareAndSwap(true, false) {
					m.cancelled = true
					m.message = "Cancelling; waiting for safe cleanup…"
					m.packApplyStop()
				} else {
					m.message = "Finishing safely; this stage cannot be cancelled."
				}
				return m, nil
			}
			m.cancelled = true
			return m, tea.Quit
		}
		switch m.step {
		case setupClients:
			return m.updateSetupClients(message)
		case setupAPIs:
			return m.updateSetupAPIs(message)
		case setupSummary:
			return m.updateSetupSummary(message)
		case setupSettings:
			return m.updateSetupSettings(message)
		case setupDone:
			switch message.String() {
			case "enter", "q", "esc":
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m *setupModel) updateSetupClients(key tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch key.String() {
	case "q", "esc":
		m.cancelled = true
		return m, tea.Quit
	case "up", "k":
		m.moveClientCursor(-1)
	case "down", "j":
		m.moveClientCursor(1)
	case "v":
		m.showAll = !m.showAll
		m.cursor = m.firstVisibleSelectable()
	case " ":
		m.toggleClient()
	case "a":
		m.toggleAllClients()
	case "enter":
		if m.packManager == nil {
			m.step, m.cursor, m.message = setupSummary, 0, ""
			break
		}
		return m, m.beginPackRefresh()
	}
	return m, nil
}

func (m *setupModel) updateSetupAPIs(key tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch key.String() {
	case "esc":
		m.invalidatePackRefresh()
		m.step, m.cursor = setupClients, m.firstSelectable()
	case "r":
		if !m.packLoading {
			return m, m.beginPackRefresh()
		}
	case "left", "h":
		m.movePackCursor(-1, 0)
	case "right", "l":
		m.movePackCursor(1, 0)
	case "up", "k":
		m.movePackCursor(0, -1)
	case "down", "j":
		m.movePackCursor(0, 1)
	case " ":
		if !m.packLoading {
			pack := m.focusedPack()
			if pack == nil {
				break
			}
			m.packSelected[pack.ID] = !m.packSelected[pack.ID]
			m.packMessage = ""
		}
	case "a":
		if !m.packLoading {
			m.toggleAllPacks()
		}
	case "enter":
		if m.packLoading {
			m.invalidatePackRefresh()
			m.packMessage = "refresh skipped"
		}
		m.step, m.cursor, m.message = setupSummary, 0, ""
	}
	return m, nil
}

func (m *setupModel) beginPackRefresh() tea.Cmd {
	m.step = setupAPIs
	m.packLoading, m.packError, m.packMessage = true, nil, ""
	m.packRefreshID++
	ctx, cancel := context.WithCancel(m.ctx)
	m.packRefreshStop = cancel
	id, manager := m.packRefreshID, m.packManager
	return func() tea.Msg {
		catalog, err := manager.Refresh(ctx)
		return setupPacksRefreshedMsg{id: id, catalog: catalog, err: err}
	}
}

func (m *setupModel) invalidatePackRefresh() {
	if !m.packLoading {
		return
	}
	if m.packRefreshStop != nil {
		m.packRefreshStop()
		m.packRefreshStop = nil
	}
	m.packRefreshID++
	m.packLoading = false
}

func (m *setupModel) updateSetupSummary(key tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch key.String() {
	case "up", "k", "down", "j":
		m.cursor = 1 - m.cursor
	case "esc":
		if m.packManager != nil {
			m.step, m.packCursor = setupAPIs, max(0, m.packCursor)
			m.ensurePackVisible()
		} else {
			m.step, m.cursor = setupClients, m.firstSelectable()
		}
	case "r":
		m.settings, m.message = config.Default(), "Recommended defaults restored; they are saved when setup finishes."
	case "enter":
		if m.cursor == 1 {
			m.step, m.settingCursor, m.message = setupSettings, 0, ""
			return m, nil
		}
		m.step = setupApplying
		return m, m.applySetup()
	}
	return m, nil
}

func (m *setupModel) updateSetupSettings(key tea.KeyMsg) (tea.Model, tea.Cmd) {
	settings := editableConfigSettings()
	setting := settings[m.settingCursor]
	if m.editing {
		if setting.boolean {
			switch key.String() {
			case "up", "k", "down", "j", "left", "right", " ":
				m.choice = !m.choice
			case "enter":
				if err := setConfigValue(&m.settings, setting.key, strconv.FormatBool(m.choice)); err != nil {
					m.message = styleError.Render(err.Error())
					return m, nil
				}
				m.editing, m.message = false, ""
			case "esc":
				m.editing = false
			}
			return m, nil
		}
		switch key.Type {
		case tea.KeyEnter:
			if err := setConfigValue(&m.settings, setting.key, m.input); err != nil {
				m.message = styleError.Render(err.Error())
				return m, nil
			}
			m.editing, m.message = false, ""
		case tea.KeyEsc:
			m.editing = false
		case tea.KeyBackspace:
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}
		case tea.KeyRunes:
			m.input += string(key.Runes)
		}
		return m, nil
	}
	switch key.String() {
	case "esc", "q":
		m.step, m.cursor, m.message = setupSummary, 0, ""
	case "up", "k":
		m.settingCursor = (m.settingCursor - 1 + len(settings)) % len(settings)
	case "down", "j":
		m.settingCursor = (m.settingCursor + 1) % len(settings)
	case "r":
		m.settings, m.message = config.Default(), "Recommended defaults restored."
	case "enter":
		m.editing = true
		if setting.boolean {
			m.choice = configBoolValue(m.settings, setting.key)
		} else {
			m.input = configValue(m.settings, setting.key)
		}
	}
	return m, nil
}

func (m *setupModel) applySetup() tea.Cmd {
	paths, cfg := m.paths, m.settings
	originalSettings := m.originalSettings
	clients := append([]install.Status(nil), m.clients...)
	selected := append([]bool(nil), m.selected...)
	manager, catalog, catalogOK := m.packManager, m.packCatalog, m.packCatalogOK
	desired := m.selectedPackIDs()
	packSelected, _, packBytes, _ := m.packSelectionSummary()
	packChanged := m.packSelectionChanged()
	executable := m.options.Executable
	parentCtx := m.ctx
	m.packApplyID++
	applyID := m.packApplyID
	ctx, cancel := context.WithCancel(parentCtx)
	m.packApplyStop = cancel
	m.packApplyEvents = make(map[string]docpacks.ApplyEvent, len(desired))
	m.packApplyOrder = m.packApplyOrder[:0]
	if packChanged {
		m.packApplyOrder = append(m.packApplyOrder, desired...)
		sort.Strings(m.packApplyOrder)
	}
	var preparedTotal int64
	available := make(map[string]docpacks.Pack, len(catalog.Packs))
	for _, pack := range catalog.Packs {
		available[pack.ID] = pack
	}
	for _, id := range m.packApplyOrder {
		pack := available[id]
		preparedTotal += pack.Bytes
		m.packApplyEvents[id] = docpacks.ApplyEvent{
			Stage: docpacks.ApplyStageWaiting, PackID: id, PackName: pack.Name,
			PackBytesTotal: pack.Bytes,
		}
	}
	for id, event := range m.packApplyEvents {
		event.PreparedBytesTotal = preparedTotal
		m.packApplyEvents[id] = event
	}
	m.packApplyStage = docpacks.ApplyStageWaiting
	m.packApplyPhase = setupApplyPreparing
	m.packApplyChanged = packChanged
	m.packRemovalOnly = packChanged && len(desired) == 0
	m.message = ""
	canStop := &atomic.Bool{}
	canStop.Store(manager != nil && catalogOK && packChanged)
	m.packApplyCanStop = canStop
	updates := make(chan tea.Msg, 64)
	done := make(chan struct{})
	m.packApplyDone = done
	m.packApplyUpdates = updates
	go func() {
		defer close(done)
		defer close(updates)
		defer cancel()
		packResult := setupPackApplyResult{
			changed: packChanged, selected: packSelected, bytes: packBytes,
		}
		send := func(message tea.Msg) bool {
			if ctx.Err() != nil {
				return false
			}
			select {
			case updates <- message:
				return ctx.Err() == nil
			case <-ctx.Done():
				return false
			}
		}
		finish := func(message setupAppliedMsg) { message.id = applyID; send(message) }
		phase := func(value setupApplyPhase) bool {
			return send(setupApplyPhaseMsg{id: applyID, phase: value})
		}
		if manager != nil && catalogOK {
			packResult.attempted = true
			if packChanged {
				err := manager.Apply(ctx, catalog, desired, func(ctx context.Context, archives []string) error {
					return library.Rebuild(ctx, library.Options{
						UserRoot: paths.Library, IndexPath: filepath.Join(paths.Index, "library.sqlite"),
						PackArchives: archives, ListTokenBudget: cfg.ListTokenBudget, ReadTokenBudget: cfg.ReadTokenBudget,
					})
				}, func(event docpacks.ApplyEvent) {
					if event.Stage == docpacks.ApplyStageIndexing {
						canStop.Store(false)
					}
					send(setupApplyProgressMsg{id: applyID, event: event})
				})
				if err != nil {
					packResult.err = err
					finish(setupAppliedMsg{packs: packResult, err: fmt.Errorf("apply API packs: %w", err)})
					return
				}
			}
		}
		canStop.Store(false)
		if !phase(setupApplySaving) {
			return
		}
		if parentCtx.Err() != nil {
			return
		}
		if cfg != originalSettings {
			if err := config.Save(paths, cfg); err != nil {
				finish(setupAppliedMsg{packs: packResult, err: err})
				return
			}
		}
		if !phase(setupApplyRegistering) {
			return
		}
		var results []clientApplyResult
		var failures []error
		for index, status := range clients {
			if parentCtx.Err() != nil {
				return
			}
			if status.Err != nil || !selected[index] && !status.Configured {
				continue
			}
			action := "registered"
			var applied []install.Result
			var err error
			options := install.Options{Executable: executable, ClientIDs: []string{status.Client.ID}, Backup: true}
			if selected[index] {
				applied, err = install.Configure(options)
			} else {
				action = "unregistered"
				applied, err = install.Uninstall(options)
			}
			result := clientApplyResult{client: status.Client, action: action, path: status.Client.ConfigPath, err: err}
			if len(applied) > 0 {
				result.changed, result.path, result.backup = applied[0].Changed, applied[0].Path, applied[0].Backup
			}
			if err != nil {
				failures = append(failures, err)
			}
			results = append(results, result)
		}
		if !phase(setupApplyFinishing) {
			return
		}
		finish(setupAppliedMsg{results: results, packs: packResult, err: errors.Join(failures...)})
	}()
	return waitSetupApply(applyID, updates, parentCtx)
}

func waitSetupApply(id int, updates <-chan tea.Msg, parent context.Context) tea.Cmd {
	if updates == nil {
		return nil
	}
	return func() tea.Msg {
		select {
		case message, ok := <-updates:
			if !ok {
				return setupApplyClosedMsg{id: id}
			}
			return message
		case <-parent.Done():
			return setupApplyClosedMsg{id: id}
		}
	}
}

func (m *setupModel) View() string {
	switch m.step {
	case setupClients:
		return m.viewSetupClients()
	case setupAPIs:
		return m.viewSetupAPIs()
	case setupSummary:
		return m.viewSetupSummary()
	case setupSettings:
		return m.viewSetupSettings()
	case setupApplying:
		return m.viewSetupApplying()
	default:
		return m.viewSetupDone()
	}
}

func setupHeader() string { return styleTitle.Render("apis-mcp setup") }

func (m *setupModel) viewSetupApplying() string {
	lines := []string{setupHeader(), "", "Applying configuration"}
	var prepared, total int64
	for _, event := range m.packApplyEvents {
		prepared = max(prepared, event.PreparedBytesDone)
		total = max(total, event.PreparedBytesTotal)
	}
	if m.packManager == nil {
		lines = append(lines, "  API packs are not managed in this setup.")
	} else if !m.packCatalogOK {
		lines = append(lines, "  API packs unchanged because the catalog is unavailable.")
	} else if !m.packApplyChanged {
		lines = append(lines, "  API packs unchanged; skipping archive verification.")
	} else {
		lines = append(lines, "  API packs")
		packBudget := max(1, m.height-8)
		shown := min(len(m.packApplyOrder), packBudget)
		if len(m.packApplyOrder) > packBudget {
			shown = max(0, packBudget-1)
		}
		for _, id := range m.packApplyOrder[:shown] {
			event := m.packApplyEvents[id]
			name := event.PackName
			if name == "" {
				name = id
			}
			lines = append(lines, fmt.Sprintf("  %-24s %-16s %s / %s", safeLine(name), setupApplyStatus(event), formatSetupBytes(event.PackBytesDone), formatSetupBytes(event.PackBytesTotal)))
		}
		if hidden := len(m.packApplyOrder) - shown; hidden > 0 {
			lines = append(lines, fmt.Sprintf("  %d more packs", hidden))
		} else if len(m.packApplyOrder) == 0 {
			if m.packRemovalOnly {
				lines = append(lines, "  No API packs selected; preparing removal-only rebuild.")
			} else {
				lines = append(lines, "  No API packs selected; nothing to download.")
			}
		}
		lines = append(lines, fmt.Sprintf("  Prepared %s / %s  %s", formatSetupBytes(prepared), formatSetupBytes(total), setupProgressBar(prepared, total, 24)))
	}
	lines = append(lines, "  "+setupApplyPhaseText(m.packApplyPhase))
	if m.cancelled {
		lines = append(lines, "  "+m.message)
	} else if m.message == "Finishing safely; this stage cannot be cancelled." {
		lines = append(lines, "  "+m.message)
	} else if m.packApplyCanStop != nil && m.packApplyCanStop.Load() {
		lines = append(lines, styleDim.Render("  ctrl+c cancel safely"))
	} else {
		lines = append(lines, styleDim.Render("  Finishing safely; this stage cannot be cancelled."))
	}
	if m.height > 0 && len(lines) > m.height {
		if m.height == 1 {
			lines = lines[:1]
		} else {
			lines = append([]string{lines[0]}, lines[len(lines)-m.height+1:]...)
		}
	}
	return m.setupView(lines)
}

func setupApplyPhaseText(phase setupApplyPhase) string {
	switch phase {
	case setupApplyIndexing:
		return "Indexing documentation library…"
	case setupApplyPublishing:
		return "Publishing API pack selection…"
	case setupApplySaving:
		return "Saving settings…"
	case setupApplyRegistering:
		return "Registering clients…"
	case setupApplyFinishing:
		return "Finishing setup…"
	default:
		return "Preparing API packs…"
	}
}

func setupApplyStatus(event docpacks.ApplyEvent) string {
	switch event.Stage {
	case docpacks.ApplyStageCheckingCache:
		return "checking cache"
	case docpacks.ApplyStageDownloading:
		return "downloading"
	case docpacks.ApplyStageVerifying:
		return "verifying"
	case docpacks.ApplyStageReady:
		if event.Cached {
			return "cached"
		}
		return "ready"
	default:
		return "waiting"
	}
}

func setupProgressBar(done, total int64, width int) string {
	filled := 0
	if total > 0 {
		filled = int(min(total, max(int64(0), done)) * int64(width) / total)
	}
	return "[" + strings.Repeat("#", filled) + strings.Repeat("-", width-filled) + "]"
}

func (m *setupModel) viewSetupClients() string {
	var out strings.Builder
	out.WriteString(setupHeader())
	out.WriteString("\nAI clients — which should be able to use your APIs?\n\n")
	indices := m.visibleClients()
	if len(indices) == 0 {
		out.WriteString(styleDim.Render("  No AI clients detected. Install one, then run\n  `apis-mcp configure`.\n"))
	}
	for _, index := range indices {
		status := m.clients[index]
		cursor := " "
		if index == m.cursor && status.Err == nil {
			cursor = styleTitle.Render(">")
		}
		mark := styleDim.Render("○")
		if status.Err != nil {
			mark = styleDim.Render("·")
		} else if m.selected[index] {
			mark = styleSuccess.Render("●")
		}
		line := fmt.Sprintf("%-22s %s", status.Client.Name, styleDim.Render(setupClientStatus(status)))
		if status.Err != nil {
			line = styleDim.Render(fmt.Sprintf("%-22s ", status.Client.Name)) + styleWarning.Render("inspection failed")
		}
		fmt.Fprintf(&out, " %s %s %s\n", cursor, mark, line)
	}
	hidden := len(m.clients) - len(indices)
	if hidden > 0 && !m.showAll {
		out.WriteString("\n" + styleDim.Render(fmt.Sprintf("  press v to show %d client(s) that are not installed", hidden)))
	} else if m.showAll {
		out.WriteString("\n" + styleDim.Render("  press v to hide clients that are not installed"))
	}
	if m.message != "" {
		out.WriteString("\n\n  " + m.message)
	}
	out.WriteString("\n\n" + styleDim.Render("  ↑↓ move · space toggle · a all/none · v show all · enter continue · q cancel"))
	return out.String()
}

func (m *setupModel) viewSetupAPIs() string {
	if m.width < minimumWidth || m.height < minimumHeight {
		return fmt.Sprintf("\n  Terminal is %dx%d.\n  apis-mcp needs at least %dx%d.\n", m.width, m.height, minimumWidth, minimumHeight)
	}
	lines := []string{
		setupHeader(),
		"Manage APIs — choose downloadable documentation packs",
		m.packCatalogStatus(),
		"",
	}
	if !m.packCatalogOK {
		if m.packLoading {
			lines = append(lines, "  Refreshing the API pack catalog; the current selection is locked.", "", "  enter continue unchanged · esc back")
		} else {
			lines = append(lines,
				"  The catalog is unavailable. Existing API packs will be left unchanged.",
				"",
				"  r retry · enter continue · esc back",
			)
		}
		return m.setupView(lines)
	}

	columns := m.packColumns()
	visibleRows := m.packVisibleRows()
	totalRows := (len(m.packCatalog.Packs) + columns - 1) / columns
	startRow := min(m.packScrollRow, max(0, totalRows-visibleRows))
	endRow := min(totalRows, startRow+visibleRows)
	cellWidth := max(1, (m.width-2-(columns-1)*2)/columns)
	for row := startRow; row < endRow; row++ {
		cells := make([]string, 0, columns)
		for column := 0; column < columns; column++ {
			index := row*columns + column
			if index >= len(m.packCatalog.Packs) {
				cells = append(cells, fixed("", cellWidth))
				continue
			}
			pack := m.packCatalog.Packs[index]
			cursor := " "
			if index == m.packCursor {
				cursor = styleTitle.Render(">")
			}
			mark := styleDim.Render("○")
			if m.packSelected[pack.ID] {
				mark = styleSuccess.Render("●")
			}
			name := safeLine(pack.Name)
			if index == m.packCursor {
				name = styleTitle.Render(name)
			}
			cells = append(cells, fixed(fmt.Sprintf("%s %s %s · %s", cursor, mark, name, m.packStatus(pack)), cellWidth))
		}
		lines = append(lines, "  "+strings.Join(cells, "  "))
	}
	for row := endRow - startRow; row < visibleRows; row++ {
		lines = append(lines, "")
	}

	lines = append(lines, "")
	if pack := m.focusedPack(); pack != nil {
		lines = append(lines, "  "+styleTitle.Render(safeLine(pack.Name))+"  "+styleDim.Render(m.packStatus(*pack)))
		description := strings.TrimSpace(safeMultiline(pack.Description))
		if description == "" {
			description = "No description provided."
		}
		wrapped := wrapLines(description, max(10, m.width-4))
		for index := 0; index < 2; index++ {
			line := ""
			if index < len(wrapped) {
				line = "  " + wrapped[index]
			}
			lines = append(lines, line)
		}
		versions := safeLine(strings.Join(pack.Versions, ", "))
		sizeLabel := "download"
		if m.packUnlisted[pack.ID] {
			sizeLabel = "installed archive"
		}
		lines = append(lines, fmt.Sprintf("  Versions: %s · %d pages · %s %s", versions, pack.Pages, formatSetupBytes(pack.Bytes), sizeLabel))
	} else {
		lines = append(lines, "  No API packs are currently listed.", "", "", "")
	}
	lines = append(lines, "", "  Pending: "+m.packSummaryText())
	footer := "  arrows/hjkl move · space toggle · a all/none · r refresh · enter continue · esc back"
	if m.packLoading {
		footer = "  arrows/hjkl move · refreshing; selection locked · enter use current · esc back"
	}
	lines = append(lines, styleDim.Render(footer))
	return m.setupView(lines)
}

func (m *setupModel) setupView(lines []string) string {
	for index, line := range lines {
		lines[index] = ansi.Truncate(line, max(1, m.width), "…")
	}
	return strings.Join(lines, "\n")
}

func (m *setupModel) packCatalogStatus() string {
	if !m.packCatalogOK {
		if m.packLoading {
			return styleDim.Render("Catalog: loading…")
		}
		if m.packError != nil {
			detail := ansi.Truncate(safeLine(m.packError.Error()), max(10, m.width-25), "…")
			return styleWarning.Render("Catalog unavailable: " + detail)
		}
		return styleDim.Render("Catalog: not loaded")
	}
	unlisted := m.unlistedPackCount()
	listed := len(m.packCatalog.Packs) - unlisted
	status := fmt.Sprintf("Catalog: %d packs", listed)
	if unlisted > 0 {
		status += fmt.Sprintf(" · %d unlisted installed", unlisted)
	}
	if m.packLoading {
		status += " · refreshing…"
	} else if m.packError != nil {
		status += " · refresh failed; showing the previous catalog"
	}
	if m.packMessage != "" {
		status += " · " + m.packMessage
	}
	totalRows := (len(m.packCatalog.Packs) + m.packColumns() - 1) / m.packColumns()
	if totalRows > m.packVisibleRows() {
		start := min(m.packScrollRow, totalRows-1) + 1
		end := min(totalRows, m.packScrollRow+m.packVisibleRows())
		status += fmt.Sprintf(" · rows %d-%d of %d", start, end, totalRows)
	}
	return styleDim.Render(status)
}

func (m *setupModel) packColumns() int {
	switch {
	case m.width >= 120:
		return 3
	case m.width >= 80:
		return 2
	default:
		return 1
	}
}

func (m *setupModel) packVisibleRows() int { return max(1, m.height-12) }

func (m *setupModel) replacePackCatalog(catalog docpacks.Catalog) {
	focusedID, previous := "", m.packCursor
	if pack := m.focusedPack(); pack != nil {
		focusedID = pack.ID
	}
	refreshed := make([]docpacks.Pack, len(catalog.Packs))
	copy(refreshed, catalog.Packs)
	catalog.Packs = refreshed
	listed := make(map[string]bool, len(catalog.Packs))
	for _, pack := range catalog.Packs {
		listed[pack.ID] = true
	}
	missing := make([]string, 0, len(m.packActive))
	for id := range m.packActive {
		if !listed[id] {
			missing = append(missing, id)
		}
	}
	sort.Strings(missing)
	m.packUnlisted = make(map[string]bool, len(missing))
	for _, id := range missing {
		catalog.Packs = append(catalog.Packs, m.packActive[id])
		m.packUnlisted[id] = true
	}
	m.packCatalog, m.packCatalogOK = catalog, true
	if len(catalog.Packs) == 0 {
		m.packCursor, m.packScrollRow = -1, 0
		return
	}
	m.packCursor = min(max(0, previous), len(catalog.Packs)-1)
	for index, pack := range catalog.Packs {
		if pack.ID == focusedID {
			m.packCursor = index
			break
		}
	}
	m.ensurePackVisible()
}

func (m *setupModel) focusedPack() *docpacks.Pack {
	if !m.packCatalogOK || m.packCursor < 0 || m.packCursor >= len(m.packCatalog.Packs) {
		return nil
	}
	return &m.packCatalog.Packs[m.packCursor]
}

func (m *setupModel) movePackCursor(horizontal, vertical int) {
	count := len(m.packCatalog.Packs)
	if count == 0 || m.packCursor < 0 {
		return
	}
	columns := m.packColumns()
	row, column := m.packCursor/columns, m.packCursor%columns
	next := m.packCursor
	if horizontal < 0 && column > 0 {
		next--
	} else if horizontal > 0 && column+1 < columns && next+1 < count {
		next++
	} else if vertical < 0 && row > 0 {
		next -= columns
	} else if vertical > 0 {
		if next+columns < count {
			next += columns
		} else if row < (count-1)/columns {
			next = count - 1
		}
	}
	m.packCursor = next
	m.packMessage = ""
	m.ensurePackVisible()
}

func (m *setupModel) ensurePackVisible() {
	if m.packCursor < 0 {
		m.packScrollRow = 0
		return
	}
	row := m.packCursor / m.packColumns()
	visible := m.packVisibleRows()
	totalRows := (len(m.packCatalog.Packs) + m.packColumns() - 1) / m.packColumns()
	m.packScrollRow = min(m.packScrollRow, max(0, totalRows-visible))
	if row < m.packScrollRow {
		m.packScrollRow = row
	} else if row >= m.packScrollRow+visible {
		m.packScrollRow = row - visible + 1
	}
}

func (m *setupModel) toggleAllPacks() {
	selectAll := false
	for _, pack := range m.packCatalog.Packs {
		if !m.packSelected[pack.ID] {
			selectAll = true
			break
		}
	}
	for _, pack := range m.packCatalog.Packs {
		m.packSelected[pack.ID] = selectAll
	}
	if selectAll {
		m.packMessage = "selected all"
	} else {
		m.packMessage = "cleared all"
	}
}

func (m *setupModel) packStatus(pack docpacks.Pack) string {
	installed, ok := m.packActive[pack.ID]
	if !ok {
		return "available"
	}
	if m.packUnlisted[pack.ID] {
		return "installed · unlisted"
	}
	if installed.SHA256 != pack.SHA256 {
		return "update"
	}
	return "installed"
}

func (m *setupModel) selectedPackIDs() []string {
	ids := make([]string, 0, len(m.packCatalog.Packs))
	for _, pack := range m.packCatalog.Packs {
		if m.packSelected[pack.ID] {
			ids = append(ids, pack.ID)
		}
	}
	return ids
}

func (m *setupModel) packSelectionChanged() bool {
	if !m.packCatalogOK {
		return false
	}
	desired := make(map[string]docpacks.Pack)
	for _, pack := range m.packCatalog.Packs {
		if m.packSelected[pack.ID] {
			desired[pack.ID] = pack
		}
	}
	return !reflect.DeepEqual(m.packActive, desired)
}

func (m *setupModel) packSelectionSummary() (selected, downloads int, bytes int64, removals int) {
	if !m.packCatalogOK {
		for id := range m.packActive {
			if m.packSelected[id] {
				selected++
			}
		}
		return selected, 0, 0, 0
	}
	available := make(map[string]bool, len(m.packCatalog.Packs))
	for _, pack := range m.packCatalog.Packs {
		available[pack.ID] = true
		if !m.packSelected[pack.ID] {
			continue
		}
		selected++
		installed, ok := m.packActive[pack.ID]
		if !ok || installed.SHA256 != pack.SHA256 {
			downloads++
			bytes += pack.Bytes
		}
	}
	for id := range m.packActive {
		if !available[id] || !m.packSelected[id] {
			removals++
		}
	}
	return selected, downloads, bytes, removals
}

func (m *setupModel) packSummaryText() string {
	selected, downloads, bytes, removals := m.packSelectionSummary()
	text := fmt.Sprintf("%d selected · %d downloads · %s", selected, downloads, formatSetupBytes(bytes))
	unlisted := 0
	for id := range m.packUnlisted {
		if m.packSelected[id] {
			unlisted++
		}
	}
	if unlisted > 0 {
		text += fmt.Sprintf(" · %d unlisted retained", unlisted)
	}
	if removals > 0 {
		text += fmt.Sprintf(" · %d removals", removals)
	}
	if !m.packCatalogOK {
		text += " · unchanged"
	}
	return text
}

func (m *setupModel) unlistedPackCount() int { return len(m.packUnlisted) }

func formatSetupBytes(bytes int64) string {
	if bytes < 1024 {
		return fmt.Sprintf("%d B", bytes)
	}
	units := []string{"KiB", "MiB", "GiB", "TiB"}
	value := float64(bytes)
	unit := "B"
	for _, candidate := range units {
		value /= 1024
		unit = candidate
		if value < 1024 {
			break
		}
	}
	number := strings.TrimSuffix(strconv.FormatFloat(value, 'f', 1, 64), ".0")
	return number + " " + unit
}

func (m *setupModel) viewSetupSummary() string {
	var out strings.Builder
	out.WriteString(setupHeader())
	out.WriteString("\nAPI tool configuration\n")
	if m.packManager != nil {
		out.WriteString("  API packs: " + m.packSummaryText() + "\n")
	}
	if setupUsesDefaults(m.settings) {
		out.WriteString(styleDim.Render("Recommended defaults") + "\n")
	}
	out.WriteByte('\n')
	for _, setting := range editableConfigSettings() {
		fmt.Fprintf(&out, "  %-31s %s\n", setting.label, configValue(m.settings, setting.key))
	}
	out.WriteByte('\n')
	for index, option := range []string{"Continue", "Change settings"} {
		cursor := " "
		if index == m.cursor {
			cursor = styleTitle.Render(">")
		}
		fmt.Fprintf(&out, " %s %s\n", cursor, option)
	}
	if !setupUsesDefaults(m.settings) {
		out.WriteString("\n" + styleDim.Render("  r restore defaults"))
	}
	if m.message != "" {
		out.WriteString("\n\n  " + m.message)
	}
	out.WriteString("\n\n" + styleDim.Render("  Limits apply to retrieved API documentation and responses."))
	return out.String()
}

func (m *setupModel) viewSetupSettings() string {
	var out strings.Builder
	out.WriteString(setupHeader())
	out.WriteString("\nSettings\n\n")
	for index, setting := range editableConfigSettings() {
		cursor := " "
		if index == m.settingCursor {
			cursor = styleTitle.Render(">")
		}
		value := configValue(m.settings, setting.key)
		if m.editing && index == m.settingCursor {
			if setting.boolean {
				value = "< " + strconv.FormatBool(m.choice) + " >"
			} else {
				value = m.input + "_"
			}
		}
		fmt.Fprintf(&out, " %s %-31s %s\n", cursor, setting.label, value)
	}
	if m.message != "" {
		out.WriteString("\n  " + m.message)
	}
	footer := "  ↑↓ move · enter edit · r restore defaults · esc back"
	if m.editing {
		footer = "  enter confirm · esc cancel"
	}
	out.WriteString("\n\n" + styleDim.Render(footer))
	out.WriteString("\n" + styleDim.Render("  Limits apply to retrieved API documentation and responses."))
	return out.String()
}

func (m *setupModel) viewSetupDone() string {
	var out strings.Builder
	out.WriteString(setupHeader() + "\n\n")
	if m.failure != nil && len(m.results) == 0 {
		out.WriteString(styleError.Render("  "+safeLine(m.failure.Error())) + "\n\n")
	}
	if m.packApplyResult.attempted && m.packApplyResult.err == nil {
		if !m.packApplyResult.changed {
			fmt.Fprintf(&out, "  API packs: %d active; unchanged.\n", m.packApplyResult.selected)
		} else if m.packApplyResult.selected == 0 {
			out.WriteString("  API packs: none selected; managed packs removed and library rebuilt.\n")
		} else {
			fmt.Fprintf(&out, "  API packs: %d active; library rebuilt; download budget %s.\n", m.packApplyResult.selected, formatSetupBytes(m.packApplyResult.bytes))
		}
		out.WriteByte('\n')
	} else if m.packManager != nil && !m.packApplyResult.attempted && m.failure == nil {
		out.WriteString("  API packs: unchanged because the catalog was unavailable.\n\n")
	}
	if len(m.results) == 0 && m.failure == nil {
		out.WriteString("  No clients were selected, so nothing was registered.\n")
		out.WriteString("  Settings were saved.\n\n")
	} else if len(m.results) > 0 {
		for _, result := range m.results {
			state := result.action
			if result.err != nil {
				state = styleError.Render(strings.TrimPrefix(result.err.Error(), result.client.Name+": "))
			} else if !result.changed {
				state = "unchanged"
			}
			fmt.Fprintf(&out, "  %-22s %s\n", result.client.Name, state)
		}
		out.WriteByte('\n')
		changed := false
		for _, result := range m.results {
			changed = changed || result.err == nil && result.changed
		}
		if changed {
			out.WriteString("  Restart affected clients so they pick up the change.\n\n")
		}
	}
	if m.failure == nil {
		out.WriteString(styleDim.Render("  Run `apis-mcp` to browse and use your APIs,\n  or `apis-mcp configure` to change any of this later.") + "\n\n")
	}
	out.WriteString(styleDim.Render("  enter to finish"))
	return out.String()
}

func editableConfigSettings() []configSetting { return configSettings[1:] }

func setupUsesDefaults(cfg config.Config) bool {
	defaults := config.Default()
	for _, setting := range editableConfigSettings() {
		if configValue(cfg, setting.key) != configValue(defaults, setting.key) {
			return false
		}
	}
	return true
}

func setupClientStatus(status install.Status) string {
	switch {
	case status.Err != nil:
		return "inspection failed"
	case status.Configured:
		return "configured"
	case status.Detected:
		return "detected"
	default:
		return "not installed"
	}
}

func (m *setupModel) visibleClients() []int {
	var indices []int
	for index, status := range m.clients {
		if status.Detected || status.Configured || status.Err != nil || m.showAll {
			indices = append(indices, index)
		}
	}
	return indices
}

func (m *setupModel) firstSelectable() int {
	for index, status := range m.clients {
		if status.Err == nil && (status.Detected || status.Configured) {
			return index
		}
	}
	return m.firstVisibleSelectable()
}

func (m *setupModel) firstVisibleSelectable() int {
	for _, index := range m.visibleClients() {
		if m.clients[index].Err == nil {
			return index
		}
	}
	return -1
}

func (m *setupModel) moveClientCursor(direction int) {
	indices := m.visibleClients()
	selectable := indices[:0]
	for _, index := range indices {
		if m.clients[index].Err == nil {
			selectable = append(selectable, index)
		}
	}
	if len(selectable) == 0 {
		return
	}
	position := 0
	for index, value := range selectable {
		if value == m.cursor {
			position = index
			break
		}
	}
	position = (position + direction + len(selectable)) % len(selectable)
	m.cursor = selectable[position]
}

func (m *setupModel) toggleClient() {
	if m.cursor < 0 || m.cursor >= len(m.clients) || m.clients[m.cursor].Err != nil || !m.clientVisible(m.cursor) {
		return
	}
	m.selected[m.cursor], m.message = !m.selected[m.cursor], ""
}

func (m *setupModel) toggleAllClients() {
	selectAll := false
	for index, status := range m.clients {
		if status.Err == nil && m.clientVisible(index) && !m.selected[index] {
			selectAll = true
			break
		}
	}
	for index, status := range m.clients {
		if status.Err == nil && m.clientVisible(index) {
			m.selected[index] = selectAll
		}
	}
	if selectAll {
		m.message = "Selected every client."
	} else {
		m.message = "Cleared every client."
	}
}

func (m *setupModel) clientVisible(index int) bool {
	status := m.clients[index]
	return status.Detected || status.Configured || status.Err != nil || m.showAll
}

func configureRunResult(ctx context.Context, root *model, programErr error) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if programErr != nil {
		return programErr
	}
	if root.runErr != nil {
		return root.runErr
	}
	if !root.configureSaved || root.settings.dirty {
		return errors.New("configuration cancelled before a successful save")
	}
	return nil
}

func configValue(cfg config.Config, key string) string {
	switch key {
	case "version":
		return strconv.Itoa(cfg.Version)
	case "list_token_budget":
		return strconv.Itoa(cfg.ListTokenBudget)
	case "read_token_budget":
		return strconv.Itoa(cfg.ReadTokenBudget)
	case "response_size_limit":
		return strconv.FormatInt(cfg.ResponseSizeLimit, 10)
	case "allow_large_download":
		return strconv.FormatBool(cfg.AllowLargeDownload)
	case "free_disk_reserve":
		return strconv.FormatInt(cfg.FreeDiskReserve, 10)
	case "maximum_redirects":
		return strconv.Itoa(cfg.MaximumRedirects)
	case "maximum_retries":
		return strconv.Itoa(cfg.MaximumRetries)
	case "maximum_header_timeout_seconds":
		return strconv.Itoa(cfg.MaximumHeaderTimeout)
	case "background_after_seconds":
		return strconv.Itoa(cfg.BackgroundAfterSeconds)
	case "stalled_download_seconds":
		return strconv.Itoa(cfg.StalledDownloadSeconds)
	case "retention_hours":
		return strconv.Itoa(cfg.RetentionHours)
	case "tls_verify":
		return strconv.FormatBool(cfg.TLSVerify)
	default:
		return ""
	}
}

func configBoolValue(cfg config.Config, key string) bool {
	return key == "allow_large_download" && cfg.AllowLargeDownload || key == "tls_verify" && cfg.TLSVerify
}

func setConfigValue(cfg *config.Config, key, raw string) error {
	next := *cfg
	if key == "allow_large_download" || key == "tls_verify" {
		value, err := strconv.ParseBool(strings.TrimSpace(raw))
		if err != nil {
			return err
		}
		if key == "allow_large_download" {
			next.AllowLargeDownload = value
		} else {
			next.TLSVerify = value
		}
	} else if key == "response_size_limit" || key == "free_disk_reserve" {
		value, err := strconv.ParseInt(strings.TrimSpace(raw), 10, 64)
		if err != nil {
			return err
		}
		if key == "response_size_limit" {
			next.ResponseSizeLimit = value
		} else {
			next.FreeDiskReserve = value
		}
	} else {
		value, err := strconv.Atoi(strings.TrimSpace(raw))
		if err != nil {
			return err
		}
		switch key {
		case "list_token_budget":
			next.ListTokenBudget = value
		case "read_token_budget":
			next.ReadTokenBudget = value
		case "maximum_redirects":
			next.MaximumRedirects = value
		case "maximum_retries":
			next.MaximumRetries = value
		case "maximum_header_timeout_seconds":
			next.MaximumHeaderTimeout = value
		case "background_after_seconds":
			next.BackgroundAfterSeconds = value
		case "stalled_download_seconds":
			next.StalledDownloadSeconds = value
		case "retention_hours":
			next.RetentionHours = value
		default:
			return errors.New("setting is read-only")
		}
	}
	if err := config.Validate(next); err != nil {
		return err
	}
	*cfg = next
	return nil
}
