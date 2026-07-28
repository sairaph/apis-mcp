package cli

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sairaph/apis-mcp/internal/config"
	"github.com/sairaph/apis-mcp/internal/install"
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
	setupSummary
	setupSettings
	setupApplying
	setupDone
)

type setupAppliedMsg struct {
	results []clientApplyResult
	err     error
}

type setupModel struct {
	ctx     context.Context
	options Options
	paths   config.Paths

	step          setupStep
	clients       []install.Status
	selected      []bool
	cursor        int
	showAll       bool
	settings      config.Config
	settingCursor int
	editing       bool
	input         string
	choice        bool

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
	state := newSetupModel(ctx, options, paths, cfg, install.Detect("", ""))
	program := tea.NewProgram(state, tea.WithContext(ctx), tea.WithInput(options.Stdin), tea.WithOutput(options.Stdout))
	_, err = program.Run()
	return setupRunResult(ctx, state, err)
}

func setupRunResult(ctx context.Context, state *setupModel, programErr error) error {
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
	state := &setupModel{
		ctx: ctx, options: options, paths: paths,
		clients: clients, selected: make([]bool, len(clients)), settings: cfg,
	}
	for index, status := range clients {
		state.selected[index] = status.Configured || status.Detected && status.Err == nil
	}
	state.cursor = state.firstSelectable()
	return state
}

func (m *setupModel) Init() tea.Cmd { return nil }

func (m *setupModel) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch message := message.(type) {
	case setupAppliedMsg:
		m.results, m.failure, m.saved, m.step = message.results, message.err, message.err == nil, setupDone
		return m, nil
	case tea.KeyMsg:
		if message.String() == "ctrl+c" {
			m.cancelled = true
			return m, tea.Quit
		}
		switch m.step {
		case setupClients:
			return m.updateSetupClients(message)
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
		m.step, m.cursor, m.message = setupSummary, 0, ""
	}
	return m, nil
}

func (m *setupModel) updateSetupSummary(key tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch key.String() {
	case "up", "k", "down", "j":
		m.cursor = 1 - m.cursor
	case "esc":
		m.step, m.cursor = setupClients, m.firstSelectable()
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
	clients := append([]install.Status(nil), m.clients...)
	selected := append([]bool(nil), m.selected...)
	executable := m.options.Executable
	return func() tea.Msg {
		if err := config.Save(paths, cfg); err != nil {
			return setupAppliedMsg{err: err}
		}
		var results []clientApplyResult
		var failures []error
		for index, status := range clients {
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
		return setupAppliedMsg{results: results, err: errors.Join(failures...)}
	}
}

func (m *setupModel) View() string {
	switch m.step {
	case setupClients:
		return m.viewSetupClients()
	case setupSummary:
		return m.viewSetupSummary()
	case setupSettings:
		return m.viewSetupSettings()
	case setupApplying:
		return setupHeader() + "\n\nSaving settings and registering clients…\n"
	default:
		return m.viewSetupDone()
	}
}

func setupHeader() string { return styleTitle.Render("apis-mcp setup") }

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

func (m *setupModel) viewSetupSummary() string {
	var out strings.Builder
	out.WriteString(setupHeader())
	out.WriteString("\nAPI tool configuration\n")
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
		out.WriteString(styleError.Render("  "+m.failure.Error()) + "\n\n")
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
