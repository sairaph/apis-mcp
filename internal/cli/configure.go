package cli

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

type configureModel struct {
	ctx      context.Context
	options  Options
	paths    config.Paths
	cfg      config.Config
	clients  []install.Status
	selected []bool
	panel    int
	cursor   int
	width    int
	editing  bool
	input    string
	confirm  bool
	busy     bool
	status   string
}

type configureResult struct {
	text string
	err  error
}

func newConfigureModel(ctx context.Context, options Options, paths config.Paths, cfg config.Config, clients []install.Status) configureModel {
	selected := make([]bool, len(clients))
	for index := range clients {
		selected[index] = clients[index].Detected
	}
	return configureModel{
		ctx: ctx, options: options, paths: paths, cfg: cfg, clients: clients,
		selected: selected, width: 90, status: "Select detected clients and review settings",
	}
}

func (m configureModel) Init() tea.Cmd { return nil }

func (m configureModel) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch message := message.(type) {
	case tea.WindowSizeMsg:
		m.width = message.Width
	case configureResult:
		m.busy = false
		if message.err != nil {
			m.status = "Error: " + message.err.Error()
		} else {
			m.status = message.text
		}
	case restoredConfig:
		m.cfg, m.busy, m.status = message.cfg, false, "Defaults restored and saved atomically"
	case tea.KeyMsg:
		if m.busy {
			if message.String() == "ctrl+c" {
				return m, tea.Quit
			}
			return m, nil
		}
		if m.confirm {
			switch message.String() {
			case "y", "Y":
				m.confirm, m.busy = false, true
				return m, m.restoreDefaults()
			case "n", "N", "esc":
				m.confirm = false
				m.status = "Restore cancelled"
			}
			return m, nil
		}
		if m.editing {
			switch message.Type {
			case tea.KeyEnter:
				if err := setConfigValue(&m.cfg, configSettings[m.cursor].key, m.input); err != nil {
					m.status = "Invalid value: " + err.Error()
					return m, nil
				}
				m.editing = false
				m.status = "Setting updated; press s to save"
			case tea.KeyEsc:
				m.editing = false
			case tea.KeyBackspace, tea.KeyDelete:
				if len(m.input) > 0 {
					m.input = m.input[:len(m.input)-1]
				}
			case tea.KeyRunes:
				m.input += string(message.Runes)
			}
			return m, nil
		}
		switch message.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab", "left", "right", "h", "l":
			m.panel = (m.panel + 1) % 2
			m.cursor = 0
		case "up", "k":
			m.cursor = (m.cursor - 1 + m.itemCount()) % m.itemCount()
		case "down", "j":
			m.cursor = (m.cursor + 1) % m.itemCount()
		case "enter", " ":
			if m.panel == 0 {
				if m.clients[m.cursor].Detected {
					m.selected[m.cursor] = !m.selected[m.cursor]
				} else {
					m.status = "Only detected clients can be selected here; use --client to create one"
				}
			} else {
				setting := configSettings[m.cursor]
				if setting.readOnly {
					m.status = "Config version is read-only"
				} else if setting.boolean {
					_ = setConfigValue(&m.cfg, setting.key, strconv.FormatBool(!configBoolValue(m.cfg, setting.key)))
					m.status = "Setting updated; press s to save"
				} else {
					m.editing, m.input = true, configValue(m.cfg, setting.key)
				}
			}
		case "s":
			if err := config.Validate(m.cfg); err != nil {
				m.status = "Invalid configuration: " + err.Error()
				break
			}
			m.busy, m.status = true, "Saving configuration"
			return m, m.save()
		case "d":
			m.confirm = true
			m.status = "Restore every setting to defaults? (y/n)"
		}
	}
	return m, nil
}

func (m configureModel) itemCount() int {
	if m.panel == 0 {
		return len(m.clients)
	}
	return len(configSettings)
}

func (m configureModel) save() tea.Cmd {
	return func() tea.Msg {
		if err := config.Save(m.paths, m.cfg); err != nil {
			return configureResult{err: err}
		}
		ids := make([]string, 0, len(m.clients))
		for index, status := range m.clients {
			if m.selected[index] {
				ids = append(ids, status.Client.ID)
			}
		}
		var output bytes.Buffer
		if len(ids) > 0 {
			results, err := install.Configure(install.Options{Executable: m.options.Executable, ClientIDs: ids, Backup: true})
			if err != nil {
				return configureResult{err: err}
			}
			_ = RenderHuman(&output, results)
		}
		text := "Configuration saved atomically to " + m.paths.Config
		if rendered := strings.TrimSpace(output.String()); rendered != "" {
			text += " | " + strings.ReplaceAll(rendered, "\n", "; ")
		}
		return configureResult{text: text}
	}
}

func (m configureModel) restoreDefaults() tea.Cmd {
	return func() tea.Msg {
		cfg, err := config.RestoreDefaults(m.paths)
		if err != nil {
			return configureResult{err: err}
		}
		return restoredConfig{cfg: cfg}
	}
}

type restoredConfig struct{ cfg config.Config }

func (m configureModel) View() string {
	accent := lipgloss.NewStyle().Foreground(lipgloss.Color("42"))
	muted := lipgloss.NewStyle().Foreground(lipgloss.Color("244"))
	selectedStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("230")).Background(lipgloss.Color("57"))
	var out strings.Builder
	out.WriteString(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("229")).Background(lipgloss.Color("24")).Padding(0, 1).Render("APIS / MCP SETUP"))
	out.WriteString("\n")
	out.WriteString(muted.Render("Clients write their own config files; application settings are saved to " + m.paths.Config))
	out.WriteString("\n\n")
	clientHeader, settingHeader := "Clients", "MCP settings"
	if m.panel == 0 {
		clientHeader = accent.Render("Clients")
	} else {
		settingHeader = accent.Render("MCP settings")
	}
	out.WriteString(clientHeader + "\n")
	for index, status := range m.clients {
		mark, state := "[ ]", "not detected"
		if m.selected[index] {
			mark = "[x]"
		}
		if status.Detected {
			state = "detected"
		}
		if status.Configured {
			state = "configured"
		}
		line := fmt.Sprintf("  %s %-22s %s", mark, status.Client.Name, state)
		if m.panel == 0 && index == m.cursor {
			line = selectedStyle.Render(">" + line[1:])
		}
		out.WriteString(line + "\n")
	}
	out.WriteString("\n" + settingHeader + "\n")
	for index, setting := range configSettings {
		value := configValue(m.cfg, setting.key)
		if m.editing && m.panel == 1 && index == m.cursor {
			value = m.input + "_"
		}
		line := fmt.Sprintf("  %-34s %s", setting.label, value)
		if m.panel == 1 && index == m.cursor {
			line = selectedStyle.Render(">" + line[1:])
		}
		out.WriteString(line + "\n")
	}
	out.WriteString("\n" + accent.Render(m.status))
	out.WriteString("\n" + muted.Render("tab panels  up/down select  enter edit/toggle  s save  d restore defaults  q quit"))
	return out.String()
}

// RunConfigure starts the interactive setup used by a bare configure command.
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
	model := newConfigureModel(ctx, options, paths, cfg, install.Detect("", ""))
	program := tea.NewProgram(model, tea.WithContext(ctx), tea.WithAltScreen(), tea.WithInput(options.Stdin), tea.WithOutput(options.Stdout))
	_, err = program.Run()
	return err
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
