package cli

import (
	"context"
	"errors"
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/sairaph/apis-mcp/internal/config"
	"github.com/sairaph/apis-mcp/internal/install"
)

var maintenanceItems = []string{
	"Save settings and client registrations",
	"Restore recommended defaults",
	"Clean response cache",
	"Rebuild and reopen documentation",
	"Run diagnostics",
}

type settingsState struct {
	cfg      config.Config
	clients  []install.Status
	selected []bool
	original []bool

	panel             int
	clientCursor      int
	settingCursor     int
	maintenanceCursor int
	maintenanceOffset int
	editing           bool
	input             textinput.Model

	results     []clientApplyResult
	diagnostics []Diagnostic
	dirty       bool
}

type clientApplyResult struct {
	client  install.Client
	action  string
	changed bool
	path    string
	backup  string
	err     error
}

type settingsSavePayload struct {
	clients []install.Status
	results []clientApplyResult
	message string
	failed  int
	err     error
}

func newSettingsState(cfg config.Config) *settingsState {
	return &settingsState{cfg: cfg, input: newTextInput("setting value", "")}
}

func (s *settingsState) setClients(clients []install.Status, selectDetected bool) {
	s.clients = clients
	s.selected = make([]bool, len(clients))
	s.original = make([]bool, len(clients))
	for index, status := range clients {
		s.selected[index] = status.Configured || selectDetected && status.Detected && status.Err == nil
		s.original[index] = status.Configured
		if s.selected[index] != s.original[index] {
			s.dirty = true
		}
	}
	if len(clients) == 0 {
		s.clientCursor = 0
	} else {
		s.clientCursor = min(s.clientCursor, len(clients)-1)
	}
}

func (s *settingsState) clientsDirty() bool {
	for index := range s.selected {
		if index >= len(s.original) || s.selected[index] != s.original[index] {
			return true
		}
	}
	return false
}

func (s *settingsState) resize(width int) { s.input.Width = max(10, width/2-22) }

func (m *model) updateSettings(key tea.KeyMsg) (tea.Model, tea.Cmd) {
	s := m.settings
	if s.editing {
		switch key.String() {
		case "esc":
			s.editing = false
			s.input.Blur()
			return m, nil
		case "enter":
			setting := configSettings[s.settingCursor]
			if err := setConfigValue(&s.cfg, setting.key, s.input.Value()); err != nil {
				m.status, m.severity = "Invalid value: "+err.Error(), severityError
				return m, nil
			}
			s.editing = false
			s.input.Blur()
			s.dirty, m.configureSaved = true, false
			m.status, m.severity = "Setting changed; save to apply it", severityWarning
			return m, nil
		}
		var cmd tea.Cmd
		s.input, cmd = s.input.Update(key)
		return m, cmd
	}
	if s.panel == 2 {
		switch key.String() {
		case "pgup":
			s.maintenanceOffset = max(0, s.maintenanceOffset-5)
			return m, nil
		case "pgdown":
			s.maintenanceOffset += 5
			return m, nil
		}
	}
	switch key.String() {
	case "tab", "right", "l":
		s.panel = (s.panel + 1) % 3
	case "shift+tab", "left", "h":
		s.panel = (s.panel + 2) % 3
	case "up", "k":
		s.move(-1)
	case "down", "j":
		s.move(1)
	case "enter", " ":
		switch s.panel {
		case 0:
			s.toggleClient(m)
		case 1:
			return m, s.editSetting(m)
		case 2:
			return m, m.runMaintenance()
		}
	case "s":
		return m, m.saveSettingsCmd()
	case "d":
		m.confirmRestoreDefaults()
	case "r":
		return m, m.refreshClientsCmd()
	}
	return m, nil
}

func (s *settingsState) move(direction int) {
	switch s.panel {
	case 0:
		if len(s.clients) > 0 {
			s.clientCursor = (s.clientCursor + direction + len(s.clients)) % len(s.clients)
		}
	case 1:
		s.settingCursor = (s.settingCursor + direction + len(configSettings)) % len(configSettings)
	case 2:
		s.maintenanceCursor = (s.maintenanceCursor + direction + len(maintenanceItems)) % len(maintenanceItems)
		s.maintenanceOffset = 0
	}
}

func (s *settingsState) toggleClient(m *model) {
	if len(s.clients) == 0 || s.clientCursor >= len(s.clients) {
		return
	}
	status := s.clients[s.clientCursor]
	if status.Err != nil {
		m.status, m.severity = status.Client.Name+" cannot be changed: "+status.Err.Error(), severityError
		return
	}
	s.selected[s.clientCursor] = !s.selected[s.clientCursor]
	s.dirty, m.configureSaved = true, false
	action := "register"
	if !s.selected[s.clientCursor] {
		action = "unregister"
	}
	m.status, m.severity = status.Client.Name+" will "+action+" when settings are saved", severityWarning
}

func (s *settingsState) editSetting(m *model) tea.Cmd {
	setting := configSettings[s.settingCursor]
	if setting.readOnly {
		m.status, m.severity = "Config version is read-only", severityWarning
		return nil
	}
	if setting.boolean {
		_ = setConfigValue(&s.cfg, setting.key, fmt.Sprintf("%t", !configBoolValue(s.cfg, setting.key)))
		s.dirty, m.configureSaved = true, false
		m.status, m.severity = "Setting changed; save to apply it", severityWarning
		return nil
	}
	s.editing = true
	s.input.SetValue(configValue(s.cfg, setting.key))
	s.input.CursorEnd()
	return s.input.Focus()
}

func (m *model) saveSettingsCmd() tea.Cmd {
	s := m.settings
	if err := config.Validate(s.cfg); err != nil {
		m.status, m.severity = err.Error(), severityError
		return nil
	}
	cfg := s.cfg
	clients := append([]install.Status(nil), s.clients...)
	selected := append([]bool(nil), s.selected...)
	original := append([]bool(nil), s.original...)
	paths, executable := m.paths, m.options.Executable
	return m.startOperation("settings-save", "Saving settings and client registrations", false, func(_ context.Context, id int64) tea.Msg {
		if err := config.Save(paths, cfg); err != nil {
			return asyncMsg{id: id, kind: "settings-save", err: err}
		}
		results := make([]clientApplyResult, 0, len(clients))
		var failures []error
		changed := 0
		failed := 0
		for index, status := range clients {
			if status.Err != nil || selected[index] == original[index] && !selected[index] {
				continue
			}
			action := "registered"
			var applied []install.Result
			var err error
			options := install.Options{Executable: executable, ClientIDs: []string{status.Client.ID}, Backup: true}
			if !selected[index] && original[index] {
				action = "unregistered"
				applied, err = install.Uninstall(options)
			} else if selected[index] {
				applied, err = install.Configure(options)
			} else {
				continue
			}
			result := clientApplyResult{client: status.Client, action: action, err: err, path: status.Client.ConfigPath}
			if err != nil {
				failed++
				failures = append(failures, fmt.Errorf("%s: %w", status.Client.Name, err))
			}
			if len(applied) > 0 {
				result.changed, result.path, result.backup = applied[0].Changed, applied[0].Path, applied[0].Backup
			}
			if result.changed {
				changed++
			}
			results = append(results, result)
		}
		refreshed := install.Detect("", "")
		message := "Configuration saved. Restart active apis-mcp client connections to apply application settings."
		if changed > 0 {
			message = fmt.Sprintf("Configuration saved; reload or restart %d changed AI client(s), and restart active connections for application settings.", changed)
		}
		if failed > 0 {
			message = fmt.Sprintf("Configuration saved, but %d client registration update(s) failed; review per-client results.", failed)
		}
		return asyncMsg{id: id, kind: "settings-save", value: settingsSavePayload{clients: refreshed, results: results, message: message, failed: failed, err: errors.Join(failures...)}}
	})
}

func (s *settingsState) applySave(payload settingsSavePayload) {
	s.results = payload.results
	s.dirty = false
	s.setClients(payload.clients, false)
	for _, result := range payload.results {
		if result.err == nil {
			continue
		}
		for index, status := range s.clients {
			if status.Client.ID == result.client.ID {
				s.selected[index] = result.action == "registered"
				s.dirty = true
				break
			}
		}
	}
}

func (m *model) refreshClientsCmd() tea.Cmd {
	return m.startOperation("settings-clients", "Inspecting AI client registrations", false, func(_ context.Context, id int64) tea.Msg {
		return asyncMsg{id: id, kind: "settings-clients", value: install.Detect("", "")}
	})
}

func (m *model) confirmRestoreDefaults() {
	m.confirm("Restore recommended defaults?", []string{
		"Every application setting is replaced and saved atomically.",
		"AI client registrations, documentation, sessions, and cache entries are not removed.",
	}, "Restore defaults", func() tea.Cmd {
		m.configureSaved = false
		paths := m.paths
		return m.startOperation("restore", "Restoring recommended defaults", false, func(_ context.Context, id int64) tea.Msg {
			cfg, err := config.RestoreDefaults(paths)
			return asyncMsg{id: id, kind: "restore", value: cfg, err: err}
		})
	})
}

func (m *model) runMaintenance() tea.Cmd {
	switch m.settings.maintenanceCursor {
	case 0:
		return m.saveSettingsCmd()
	case 1:
		m.confirmRestoreDefaults()
		return nil
	case 2:
		if m.runtime == nil || m.runtime.Cache == nil {
			m.status, m.severity = "Cache maintenance requires the full application runtime", severityWarning
			return nil
		}
		store := m.runtime.Cache
		return m.startOperation("cache-clean", "Cleaning response cache", false, func(_ context.Context, id int64) tea.Msg {
			result, err := store.Cleanup()
			text := fmt.Sprintf("Removed %d expired cache entries and %d orphans", result.RemovedEntries, result.RemovedOrphans)
			return asyncMsg{id: id, kind: "cache-clean", value: text, err: err}
		})
	case 3:
		return m.reloadLibraryCmd()
	case 4:
		if m.runtime == nil {
			m.status, m.severity = "Diagnostics require the full application runtime", severityWarning
			return nil
		}
		runtime, options := m.runtime, m.options
		return m.startOperation("doctor", "Running diagnostics", false, func(_ context.Context, id int64) tea.Msg {
			return asyncMsg{id: id, kind: "doctor", value: diagnose(runtime, options)}
		})
	}
	return nil
}

func (m *model) viewSettings(width, height int) []string {
	s := m.settings
	clients := s.clientRows(height - 1)
	settings := s.settingRows(height - 1)
	maintenance := s.maintenanceRows(height - 1)
	if m.width >= 120 {
		left, middle := 40, 47
		return joinPanes(height,
			pane("AI client registrations", s.panel == 0, left, height, clients),
			pane("Application settings", s.panel == 1, middle, height, settings),
			pane("Maintenance", s.panel == 2, width-left-middle-2, height, maintenance),
		)
	}
	if m.width >= 80 {
		left := width * 2 / 5
		rightRows, rightTitle := settings, "Application settings"
		if s.panel == 2 {
			rightRows, rightTitle = maintenance, "Maintenance"
		}
		return joinPanes(height,
			pane("AI client registrations", s.panel == 0, left, height, clients),
			pane(rightTitle, s.panel != 0, width-left-1, height, rightRows),
		)
	}
	switch s.panel {
	case 0:
		return pane("AI client registrations", true, width, height, clients)
	case 1:
		return pane("Application settings", true, width, height, settings)
	default:
		return pane("Maintenance", true, width, height, maintenance)
	}
}

func (s *settingsState) clientRows(height int) []string {
	if len(s.clients) == 0 {
		return []string{"  No supported client configurations found."}
	}
	start, end := visibleWindow(s.clientCursor, len(s.clients), height)
	rows := make([]string, 0, end-start)
	for index := start; index < end; index++ {
		status := s.clients[index]
		state := "not detected"
		if status.Detected {
			state = "detected"
		}
		if status.Configured {
			state = "configured"
		}
		if status.Err != nil {
			state = "inspection error"
		}
		line := fmt.Sprintf("  [%s] %-20s %s", checkMark(s.selected[index]), safeLine(status.Client.Name), state)
		rows = append(rows, selectedLine(line, 120, s.panel == 0 && index == s.clientCursor))
	}
	return rows
}

func (s *settingsState) settingRows(height int) []string {
	start, end := visibleWindow(s.settingCursor, len(configSettings), height)
	rows := make([]string, 0, end-start)
	for index := start; index < end; index++ {
		setting := configSettings[index]
		value := configValue(s.cfg, setting.key)
		if s.editing && index == s.settingCursor {
			value = s.input.View()
		}
		line := fmt.Sprintf("  %-31s %s", setting.label, value)
		rows = append(rows, selectedLine(line, 120, s.panel == 1 && index == s.settingCursor))
	}
	return rows
}

func (s *settingsState) maintenanceRows(height int) []string {
	rows := make([]string, 0, height)
	for index, item := range maintenanceItems {
		rows = append(rows, selectedLine("  "+item, 120, s.panel == 2 && index == s.maintenanceCursor))
	}
	if len(s.results) > 0 {
		rows = append(rows, "", "  "+styleDim.Render("Last client apply"))
		for _, result := range s.results {
			state := result.action
			if result.err != nil {
				state = styleError.Render(safeLine(result.err.Error()))
			} else if !result.changed {
				state = "unchanged"
			}
			rows = append(rows, fmt.Sprintf("  %-20s %s", safeLine(result.client.Name), state))
			if result.changed {
				rows = append(rows, "    "+styleWarning.Render("reload or restart this client"))
			}
		}
	}
	if len(s.diagnostics) > 0 {
		rows = append(rows, "", "  "+styleDim.Render("Diagnostics"))
		for _, item := range s.diagnostics {
			state := styleSuccess.Render("OK")
			if !item.OK {
				state = styleWarning.Render("WARN")
			}
			rows = append(rows, fmt.Sprintf("  %-4s %-16s %s", state, safeLine(item.Name), safeLine(item.Detail)))
		}
	}
	if len(rows) > height {
		s.maintenanceOffset = max(0, min(s.maintenanceOffset, len(rows)-height))
		return rows[s.maintenanceOffset : s.maintenanceOffset+height]
	}
	s.maintenanceOffset = 0
	return rows
}

func (m *model) settingsFooter() string {
	if m.settings.editing {
		return "type to edit  arrows move cursor  enter validate  esc cancel"
	}
	if m.configureOnly {
		return "tab panels  j/k navigate  enter edit/toggle/run  pgup/pgdn results  s save  d defaults  ? help  q finish"
	}
	return "tab panels  j/k navigate  enter edit/toggle/run  pgup/pgdn results  s save  d defaults  r inspect clients  ? help  q quit"
}
