package cli

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sairaph/apis-mcp/internal/bootstrap"
)

type launcherAction int

const (
	launcherNoAction launcherAction = iota
	launcherAPIs
	launcherConfigure
	launcherExit
)

type launcherItem struct {
	label       string
	description string
	action      launcherAction
}

type launcherOpenWorkspaceMsg struct{}

const launcherOpeningDelay = 50 * time.Millisecond

var launcherItems = []launcherItem{
	{label: "APIs", description: "Open the API workspace", action: launcherAPIs},
	{label: "Configure", description: "Set up API packs, clients, and settings", action: launcherConfigure},
}

type launcherModel struct {
	items   []launcherItem
	cursor  int
	action  launcherAction
	status  string
	opening bool
}

func newLauncherModel(status string) *launcherModel {
	return &launcherModel{items: append([]launcherItem(nil), launcherItems...), status: status}
}

func (m *launcherModel) Init() tea.Cmd { return nil }

func (m *launcherModel) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	if _, ok := message.(launcherOpenWorkspaceMsg); ok && m.opening {
		m.action = launcherAPIs
		return m, tea.Quit
	}
	if m.opening {
		return m, nil
	}
	key, ok := message.(tea.KeyMsg)
	if !ok {
		return m, nil
	}
	switch key.String() {
	case "q", "esc", "ctrl+c":
		m.action = launcherExit
		return m, tea.Quit
	case "up", "k":
		m.cursor = (m.cursor + len(m.items) - 1) % len(m.items)
	case "down", "j":
		m.cursor = (m.cursor + 1) % len(m.items)
	case "enter":
		action := m.items[m.cursor].action
		if action == launcherAPIs {
			m.opening = true
			return m, tea.Tick(launcherOpeningDelay, func(time.Time) tea.Msg {
				return launcherOpenWorkspaceMsg{}
			})
		}
		m.action = action
		return m, tea.Quit
	}
	return m, nil
}

func (m *launcherModel) View() string {
	var out strings.Builder
	out.WriteString(styleTitle.Render("apis-mcp"))
	out.WriteString("\n")
	out.WriteString(styleDim.Render("Your API workspace"))
	out.WriteString("\n\n")
	for index, item := range m.items {
		cursor := " "
		label := fmt.Sprintf("%-12s", item.label)
		if index == m.cursor {
			cursor = styleTitle.Render(">")
			label = styleTitle.Render(label)
		}
		fmt.Fprintf(&out, " %s %s %s\n", cursor, label, styleDim.Render(item.description))
	}
	if m.opening {
		out.WriteString("\n")
		out.WriteString(styleTitle.Render("Opening API workspace..."))
		out.WriteString("\n")
	} else if m.status != "" {
		out.WriteString("\n")
		out.WriteString(styleError.Render(safeLine(m.status)))
		out.WriteString("\n")
	}
	out.WriteString("\n")
	if m.opening {
		out.WriteString(styleDim.Render("Verifying packs and opening the documentation index"))
	} else {
		out.WriteString(styleDim.Render("up/down or j/k move  enter select  q quit"))
	}
	out.WriteString("\n")
	return out.String()
}

type launcherDependencies struct {
	runMenu      func(context.Context, Options, string) (launcherAction, error)
	openRuntime  func(context.Context) (*bootstrap.Runtime, error)
	runWorkspace func(context.Context, *bootstrap.Runtime, Options) error
	closeRuntime func(*bootstrap.Runtime) error
	runConfigure func(context.Context, Options) error
}

func defaultLauncherDependencies() launcherDependencies {
	return launcherDependencies{
		runMenu: runLauncherMenu, openRuntime: bootstrap.Open, runWorkspace: RunInteractive,
		closeRuntime: closeLauncherRuntime, runConfigure: RunConfigure,
	}
}

// RunLauncher presents the bare-terminal menu and coordinates its independent
// normal-screen and alternate-screen programs.
func RunLauncher(ctx context.Context, options Options) error {
	return runLauncher(ctx, normalizeOptions(options), defaultLauncherDependencies())
}

func runLauncher(ctx context.Context, options Options, dependencies launcherDependencies) error {
	status := ""
	for {
		action, err := dependencies.runMenu(ctx, options, status)
		if err != nil {
			return err
		}
		switch action {
		case launcherExit:
			return nil
		case launcherConfigure:
			err = dependencies.runConfigure(ctx, options)
			if ctxErr := ctx.Err(); ctxErr != nil {
				return ctxErr
			}
			if err == nil || isConfigureCancellation(err) {
				status = ""
				continue
			}
			status = "Configure: " + err.Error()
		case launcherAPIs:
			status = runLauncherWorkspace(ctx, options, dependencies)
			if ctxErr := ctx.Err(); ctxErr != nil {
				return ctxErr
			}
		default:
			return errors.New("launcher closed without a selection")
		}
	}
}

func runLauncherMenu(ctx context.Context, options Options, status string) (launcherAction, error) {
	menu := newLauncherModel(status)
	if _, err := runTUIProgram(ctx, menu, options, false); err != nil {
		if ctxErr := ctx.Err(); ctxErr != nil {
			return launcherNoAction, ctxErr
		}
		return launcherNoAction, err
	}
	return menu.action, nil
}

func runLauncherWorkspace(ctx context.Context, options Options, dependencies launcherDependencies) string {
	runtime, err := dependencies.openRuntime(ctx)
	if err != nil {
		return "APIs: " + err.Error()
	}
	runErr := dependencies.runWorkspace(ctx, runtime, options)
	closeErr := dependencies.closeRuntime(runtime)
	switch {
	case runErr != nil && closeErr != nil:
		return fmt.Sprintf("APIs: %v; shutdown: %v", runErr, closeErr)
	case runErr != nil:
		return "APIs: " + runErr.Error()
	case closeErr != nil:
		return "APIs shutdown: " + closeErr.Error()
	default:
		return ""
	}
}

func closeLauncherRuntime(runtime *bootstrap.Runtime) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	return runtime.Close(ctx)
}

func isConfigureCancellation(err error) bool {
	return err != nil && err.Error() == "configuration cancelled before a successful save"
}
