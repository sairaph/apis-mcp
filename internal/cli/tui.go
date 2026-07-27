package cli

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/sairaph/apis-mcp/internal/app"
	"github.com/sairaph/apis-mcp/internal/bootstrap"
	"github.com/sairaph/apis-mcp/internal/importer"
	"github.com/sairaph/apis-mcp/library"
)

type menuItem struct {
	title       string
	description string
	action      string
}

var menu = []menuItem{
	{title: "Browse documentation", description: "List APIs, versions, and document IDs", action: "library"},
	{title: "Search documentation", description: "Search within one document ID", action: "search"},
	{title: "Read documentation", description: "Read one page by document and page ID", action: "read"},
	{title: "HTTP call", description: "Send a local workspace HTTP request", action: "call"},
	{title: "Import documentation", description: "Import markdown, OpenAPI, llms.txt, HTML, or Docsify", action: "import"},
	{title: "Rebuild library", description: "Validate sources and publish a fresh index", action: "rebuild"},
	{title: "List sessions", description: "Review persisted cookie sessions", action: "sessions"},
	{title: "Inspect session", description: "Show cookies for a session ID", action: "session-show"},
	{title: "Delete session", description: "Delete a session after confirmation", action: "session-delete"},
	{title: "Clean cache", description: "Remove expired responses and incomplete entries", action: "cache"},
	{title: "Configure", description: "Edit settings and register detected MCP clients", action: "configure"},
	{title: "Doctor", description: "Run local runtime and installation checks", action: "doctor"},
	{title: "Quit", description: "Leave apis-mcp", action: "quit"},
}

type actionResult struct {
	action string
	text   string
	err    error
}

type formField struct {
	label string
	value string
}

type model struct {
	ctx        context.Context
	runtime    *bootstrap.Runtime
	options    Options
	cursor     int
	width      int
	height     int
	busy       bool
	status     string
	overview   string
	apiCount   int
	collection int
	sessions   int
	detail     string
	formAction string
	fields     []formField
	field      int
	confirm    bool
	handoff    bool
}

func newModel(ctx context.Context, runtime *bootstrap.Runtime, options Options) model {
	m := model{ctx: ctx, runtime: runtime, options: options, width: 80, height: 24, status: "Ready"}
	m.refreshOverview()
	return m
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch message := message.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = message.Width, message.Height
	case tea.KeyMsg:
		if m.busy {
			if message.String() == "ctrl+c" || message.String() == "q" {
				return m, tea.Quit
			}
			return m, nil
		}
		if m.formAction != "" {
			return m.updateForm(message)
		}
		switch message.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			m.cursor = (m.cursor - 1 + len(menu)) % len(menu)
		case "down", "j":
			m.cursor = (m.cursor + 1) % len(menu)
		case "enter", " ":
			item := menu[m.cursor]
			if item.action == "quit" {
				return m, tea.Quit
			}
			if item.action == "configure" {
				m.handoff = true
				return m, tea.Quit
			}
			if fields := fieldsFor(item.action); fields != nil {
				m.formAction, m.fields, m.field = item.action, fields, 0
				m.status = "Fill the form and press enter to run"
				return m, nil
			}
			m.busy, m.status = true, "Working: "+item.title
			return m, m.run(item.action)
		case "r":
			m.busy, m.status = true, "Refreshing overview"
			return m, m.run("library")
		}
	case actionResult:
		m.busy = false
		if message.err != nil {
			m.status = "Error: " + message.err.Error()
		} else {
			m.status = "Completed: " + menuTitle(message.action)
			m.detail = message.text
			m.formAction, m.fields, m.confirm = "", nil, false
		}
		if message.action == "library" || message.action == "sessions" {
			m.refreshOverview()
		}
	}
	return m, nil
}

func (m model) View() string {
	accent := lipgloss.NewStyle().Foreground(lipgloss.Color("42"))
	muted := lipgloss.NewStyle().Foreground(lipgloss.Color("244"))
	title := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("229")).Background(lipgloss.Color("24")).Padding(0, 1)
	selected := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("230")).Background(lipgloss.Color("57")).Padding(0, 1)
	panelWidth := max(12, min(70, m.width-4))
	var out strings.Builder
	out.WriteString(title.Render("APIS / MCP"))
	out.WriteString("\n")
	out.WriteString(muted.Render("Local API documentation and HTTP workspace"))
	out.WriteString("\n\n")
	out.WriteString(accent.Render(fmt.Sprintf("%d APIs", m.apiCount)))
	out.WriteString("  ")
	out.WriteString(accent.Render(fmt.Sprintf("%d collections", m.collection)))
	out.WriteString("  ")
	out.WriteString(accent.Render(fmt.Sprintf("%d sessions", m.sessions)))
	out.WriteString("\n")
	out.WriteString(muted.Render(shorten(m.overview, panelWidth)))
	out.WriteString("\n\n")
	if m.formAction != "" {
		out.WriteString(accent.Render(menuTitle(m.formAction)))
		out.WriteString("\n")
		for index, field := range m.fields {
			line := fmt.Sprintf("%-18s %s", field.label+":", field.value)
			if index == m.field {
				line = selected.Width(panelWidth).Render("> " + shorten(line, panelWidth-2))
			} else {
				line = "  " + shorten(line, panelWidth-2)
			}
			out.WriteString(line + "\n")
		}
		out.WriteString("\n")
		if m.confirm {
			out.WriteString(accent.Render("Delete this session permanently? (y/n)"))
			out.WriteString("\n")
		}
		out.WriteString(accent.Render(m.status))
		out.WriteString("\n")
		out.WriteString(muted.Render("type to edit  tab/up/down fields  ctrl+u clear  enter run  esc back"))
		return out.String()
	}
	for index, item := range menu {
		line := fmt.Sprintf("%-20s %s", item.title, item.description)
		line = shorten(line, panelWidth)
		if index == m.cursor {
			out.WriteString(selected.Width(panelWidth).Render("> " + line))
		} else {
			out.WriteString("  " + line)
		}
		out.WriteByte('\n')
	}
	out.WriteString("\n")
	status := m.status
	if m.busy {
		status += " ..."
	}
	out.WriteString(accent.Render(status))
	if m.detail != "" {
		out.WriteString("\n\n")
		out.WriteString(renderPanel(m.detail, panelWidth, max(3, m.height-len(menu)-11)))
	}
	out.WriteString("\n")
	out.WriteString(muted.Render("up/down navigate  enter run  r refresh  q quit"))
	return out.String()
}

func (m model) updateForm(message tea.KeyMsg) (tea.Model, tea.Cmd) {
	if m.confirm {
		switch message.String() {
		case "y", "Y":
			m.confirm, m.busy = false, true
			return m, m.runForm()
		case "n", "N", "esc":
			m.confirm = false
			m.status = "Deletion cancelled"
		}
		return m, nil
	}
	switch message.Type {
	case tea.KeyEsc:
		m.formAction, m.fields, m.confirm = "", nil, false
		m.status = "Ready"
	case tea.KeyCtrlC:
		return m, tea.Quit
	case tea.KeyTab, tea.KeyDown:
		m.field = (m.field + 1) % len(m.fields)
	case tea.KeyShiftTab, tea.KeyUp:
		m.field = (m.field - 1 + len(m.fields)) % len(m.fields)
	case tea.KeyCtrlU:
		m.fields[m.field].value = ""
	case tea.KeyBackspace, tea.KeyDelete:
		value := m.fields[m.field].value
		if len(value) > 0 {
			m.fields[m.field].value = value[:len(value)-1]
		}
	case tea.KeyRunes:
		m.fields[m.field].value += string(message.Runes)
	case tea.KeyEnter:
		if m.formAction == "session-delete" {
			m.confirm = true
			m.status = "Confirm session deletion"
			return m, nil
		}
		m.busy, m.status = true, "Working: "+menuTitle(m.formAction)
		return m, m.runForm()
	}
	return m, nil
}

func fieldsFor(action string) []formField {
	switch action {
	case "search":
		return []formField{{label: "Document ID"}, {label: "Query"}}
	case "read":
		return []formField{{label: "Document ID"}, {label: "Page ID"}}
	case "call":
		return []formField{{label: "Method", value: "GET"}, {label: "URL"}, {label: "Headers JSON"}, {label: "Payload JSON"}}
	case "import":
		return []formField{{label: "Kind", value: "markdown"}, {label: "API name"}, {label: "Version"}, {label: "Source"}}
	case "session-show", "session-delete":
		return []formField{{label: "Session ID"}}
	default:
		return nil
	}
}

func menuTitle(action string) string {
	for _, item := range menu {
		if item.action == action {
			return item.title
		}
	}
	return action
}

func (m model) run(action string) tea.Cmd {
	return func() tea.Msg {
		result := actionResult{action: action}
		switch action {
		case "library":
			listed, err := m.runtime.Library.List(m.ctx, library.ListRequest{Page: 1})
			result.err = err
			if err == nil {
				result.text, result.err = renderValue(listed)
			}
		case "rebuild":
			result.err = m.runtime.RebuildLibrary(m.ctx)
			result.text = "Library rebuilt. Restart clients to use the new generation."
		case "sessions":
			items, err := m.runtime.Sessions.List()
			result.err = err
			if err == nil {
				result.text, result.err = renderValue(items)
			}
		case "cache":
			cleaned, err := m.runtime.Cache.Cleanup()
			result.err = err
			if err == nil {
				result.text, result.err = renderValue(cleaned)
			}
		case "doctor":
			checks := diagnose(m.runtime, m.options)
			result.text, result.err = renderValue(checks)
		}
		return result
	}
}

func (m model) runForm() tea.Cmd {
	action := m.formAction
	fields := append([]formField(nil), m.fields...)
	return func() tea.Msg {
		result := actionResult{action: action}
		var value any
		switch action {
		case "search":
			value, result.err = m.runtime.Library.Search(m.ctx, library.SearchRequest{DocID: strings.TrimSpace(fields[0].value), Query: strings.TrimSpace(fields[1].value), Page: 1})
		case "read":
			value, result.err = m.runtime.Library.Read(m.ctx, library.ReadRequest{DocID: strings.TrimSpace(fields[0].value), PageID: strings.TrimSpace(fields[1].value)})
		case "call":
			headers, err := jsonOrPath(fields[2].value)
			if err != nil {
				result.err = fmt.Errorf("headers: %w", err)
				break
			}
			payload, err := jsonOrPath(fields[3].value)
			if err != nil {
				result.err = fmt.Errorf("payload: %w", err)
				break
			}
			value, result.err = m.runtime.HTTP.Call(m.ctx, app.CallInput{Method: strings.TrimSpace(fields[0].value), Endpoint: strings.TrimSpace(fields[1].value), Headers: headers, Payload: payload})
		case "import":
			kind := strings.ToLower(strings.TrimSpace(fields[0].value))
			name, version, source := strings.TrimSpace(fields[1].value), strings.TrimSpace(fields[2].value), strings.TrimSpace(fields[3].value)
			options := importer.Options{LibraryRoot: m.runtime.Paths.Library, Rebuild: m.runtime.RebuildLibrary}
			switch kind {
			case "markdown":
				value, result.err = importer.ImportMarkdown(m.ctx, source, options)
			case "openapi":
				value, result.err = importer.ImportOpenAPI(m.ctx, name, version, source, options)
			case "llms":
				value, result.err = importer.ImportLLMSTxt(m.ctx, name, version, source, options)
			case "html":
				value, result.err = importer.ImportHTML(m.ctx, name, version, source, options)
			case "docsify":
				options.HTMLLimitsSet, options.MaxHTMLPages, options.MaxHTMLDepth = true, -1, -1
				value, result.err = importer.ImportDocsify(m.ctx, name, version, source, options)
			default:
				result.err = fmt.Errorf("unknown import kind %q", kind)
			}
		case "session-show":
			value, result.err = m.runtime.Sessions.Inspect(strings.TrimSpace(fields[0].value))
		case "session-delete":
			id := strings.TrimSpace(fields[0].value)
			result.err = m.runtime.Sessions.Delete(id)
			value = "Session deleted: " + id
		}
		if result.err == nil {
			result.text, result.err = renderValue(value)
		}
		return result
	}
}

func renderValue(value any) (string, error) {
	var output bytes.Buffer
	if err := RenderHuman(&output, value); err != nil {
		return "", err
	}
	return strings.TrimSpace(output.String()), nil
}

func (m *model) refreshOverview() {
	if listed, err := m.runtime.Library.List(m.ctx, library.ListRequest{Page: 1}); err == nil {
		m.apiCount = listed.Total
	}
	if collections, err := m.runtime.Library.Collections(m.ctx, library.CollectionsRequest{Page: 1}); err == nil {
		m.collection = collections.Total
	}
	if sessions, err := m.runtime.Sessions.List(); err == nil {
		m.sessions = len(sessions)
	}
	m.overview = "Generation " + m.runtime.Library.Fingerprint()
}

// RunInteractive starts the full-screen Bubble Tea application.
func RunInteractive(ctx context.Context, runtime *bootstrap.Runtime, options Options) error {
	if runtime == nil {
		return fmt.Errorf("runtime is required")
	}
	options = normalizeOptions(options)
	for {
		program := tea.NewProgram(newModel(ctx, runtime, options), tea.WithContext(ctx), tea.WithAltScreen(), tea.WithInput(options.Stdin), tea.WithOutput(options.Stdout))
		final, err := program.Run()
		if err != nil {
			return err
		}
		result, ok := final.(model)
		if !ok || !result.handoff {
			return nil
		}
		if err := RunConfigure(ctx, options); err != nil {
			return err
		}
	}
}

func shorten(value string, width int) string {
	if width < 4 || len(value) <= width {
		return value
	}
	return value[:width-3] + "..."
}

func renderPanel(value string, width, height int) string {
	lines := strings.Split(value, "\n")
	if len(lines) > height {
		lines = append(lines[:height-1], fmt.Sprintf("... %d more lines", len(lines)-height+1))
	}
	for index := range lines {
		lines[index] = shorten(lines[index], width)
	}
	return strings.Join(lines, "\n")
}
