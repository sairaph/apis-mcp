package cli

import (
	"bytes"
	"context"
	"errors"
	"io"
	"reflect"
	"strings"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/ansi"
	"github.com/sairaph/apis-mcp/internal/bootstrap"
)

func TestLauncherMenuOrderAndNavigation(t *testing.T) {
	menu := newLauncherModel("")
	var labels []string
	for _, item := range menu.items {
		labels = append(labels, item.label)
	}
	if want := []string{"APIs", "Configure"}; !reflect.DeepEqual(labels, want) {
		t.Fatalf("menu labels = %v, want %v", labels, want)
	}
	view := ansi.Strip(menu.View())
	if strings.Index(view, "APIs") > strings.Index(view, "Configure") {
		t.Fatalf("menu order is wrong:\n%s", view)
	}

	menu.Update(tea.KeyMsg{Type: tea.KeyDown})
	if menu.cursor != 1 {
		t.Fatalf("down cursor = %d, want 1", menu.cursor)
	}
	menu.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'k'}})
	if menu.cursor != 0 {
		t.Fatalf("k cursor = %d, want 0", menu.cursor)
	}
	menu.Update(tea.KeyMsg{Type: tea.KeyUp})
	_, command := menu.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if menu.action != launcherConfigure || command == nil {
		t.Fatalf("wrapped selection action/command = %v/%v", menu.action, command)
	}
}

func TestLauncherMenuRendersAndLocksOpeningStateBeforeExit(t *testing.T) {
	menu := newLauncherModel("")
	_, command := menu.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if command == nil || !menu.opening || menu.action != launcherNoAction {
		t.Fatalf("opening state command/opening/action = %v/%t/%v", command, menu.opening, menu.action)
	}
	view := menu.View()
	if !strings.Contains(view, styleTitle.Render("Opening API workspace...")) {
		t.Fatalf("opening status is not rendered as information:\n%s", view)
	}
	if reflect.DeepEqual(styleTitle.GetForeground(), styleError.GetForeground()) {
		t.Fatal("informational and error statuses use the same foreground color")
	}

	menu.Update(tea.KeyMsg{Type: tea.KeyDown})
	menu.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}})
	if menu.cursor != 0 || menu.action != launcherNoAction {
		t.Fatalf("input changed opening menu cursor/action = %d/%v", menu.cursor, menu.action)
	}

	message := command()
	if _, ok := message.(launcherOpenWorkspaceMsg); !ok {
		t.Fatalf("opening tick returned %T", message)
	}
	_, quit := menu.Update(message)
	if menu.action != launcherAPIs || quit == nil {
		t.Fatalf("opening completion action/command = %v/%v", menu.action, quit)
	}
}

func TestLauncherProgramWritesOpeningStatusBeforeReturningAPIs(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	reader, writer := io.Pipe()
	defer reader.Close()
	go func() {
		_, _ = writer.Write([]byte("\r"))
		_ = writer.Close()
	}()
	var output bytes.Buffer
	action, err := runLauncherMenu(ctx, Options{Stdin: reader, Stdout: &output}, "")
	if err != nil {
		t.Fatal(err)
	}
	if action != launcherAPIs {
		t.Fatalf("action = %v, want APIs", action)
	}
	if rendered := ansi.Strip(output.String()); !strings.Contains(rendered, "Opening API workspace...") {
		t.Fatalf("program did not render opening status before returning:\n%s", rendered)
	}
}

func TestLauncherErrorStatusUsesErrorStyle(t *testing.T) {
	menu := newLauncherModel("APIs: open failed")
	view := menu.View()
	if !strings.Contains(view, styleError.Render("APIs: open failed")) {
		t.Fatalf("error status is not rendered with the error style:\n%s", view)
	}
}

func TestLauncherMenuExitKeys(t *testing.T) {
	keys := []tea.KeyMsg{
		{Type: tea.KeyRunes, Runes: []rune{'q'}},
		{Type: tea.KeyEsc},
		{Type: tea.KeyCtrlC},
	}
	for _, key := range keys {
		menu := newLauncherModel("")
		_, command := menu.Update(key)
		if menu.action != launcherExit || command == nil {
			t.Errorf("key %q action/command = %v/%v", key.String(), menu.action, command)
		}
	}
}

func TestLauncherConfiguresThenOpensFreshRuntime(t *testing.T) {
	runtime := &bootstrap.Runtime{}
	actions := []launcherAction{launcherConfigure, launcherAPIs, launcherExit}
	var events []string
	dependencies := launcherDependencies{
		runMenu: func(_ context.Context, _ Options, status string) (launcherAction, error) {
			events = append(events, "menu:"+status)
			action := actions[0]
			actions = actions[1:]
			return action, nil
		},
		runConfigure: func(context.Context, Options) error {
			events = append(events, "configure")
			return nil
		},
		openRuntime: func(context.Context) (*bootstrap.Runtime, error) {
			events = append(events, "open")
			return runtime, nil
		},
		runWorkspace: func(_ context.Context, got *bootstrap.Runtime, _ Options) error {
			if got != runtime {
				t.Fatalf("workspace runtime = %p, want %p", got, runtime)
			}
			events = append(events, "workspace")
			return nil
		},
		closeRuntime: func(got *bootstrap.Runtime) error {
			if got != runtime {
				t.Fatalf("closed runtime = %p, want %p", got, runtime)
			}
			events = append(events, "close")
			return nil
		},
	}
	if err := runLauncher(context.Background(), Options{}, dependencies); err != nil {
		t.Fatal(err)
	}
	want := []string{"menu:", "configure", "menu:", "open", "workspace", "close", "menu:"}
	if !reflect.DeepEqual(events, want) {
		t.Fatalf("events = %v, want %v", events, want)
	}
}

func TestLauncherOpensAndClosesFreshRuntimeForEachWorkspace(t *testing.T) {
	actions := []launcherAction{launcherAPIs, launcherAPIs, launcherExit}
	runtimes := []*bootstrap.Runtime{{}, {}}
	var opened, run, closed []*bootstrap.Runtime
	dependencies := launcherDependencies{
		runMenu: func(context.Context, Options, string) (launcherAction, error) {
			action := actions[0]
			actions = actions[1:]
			return action, nil
		},
		openRuntime: func(context.Context) (*bootstrap.Runtime, error) {
			runtime := runtimes[len(opened)]
			opened = append(opened, runtime)
			return runtime, nil
		},
		runWorkspace: func(_ context.Context, runtime *bootstrap.Runtime, _ Options) error {
			run = append(run, runtime)
			return nil
		},
		closeRuntime: func(runtime *bootstrap.Runtime) error {
			closed = append(closed, runtime)
			return nil
		},
		runConfigure: func(context.Context, Options) error { return nil },
	}
	if err := runLauncher(context.Background(), Options{}, dependencies); err != nil {
		t.Fatal(err)
	}
	if len(opened) != 2 || len(run) != 2 || len(closed) != 2 ||
		opened[0] != runtimes[0] || opened[1] != runtimes[1] ||
		run[0] != runtimes[0] || run[1] != runtimes[1] ||
		closed[0] != runtimes[0] || closed[1] != runtimes[1] {
		t.Fatalf("opened/run/closed = %v/%v/%v", opened, run, closed)
	}
}

func TestLauncherTreatsConfigureCancellationAsReturn(t *testing.T) {
	actions := []launcherAction{launcherConfigure, launcherExit}
	var statuses []string
	dependencies := launcherDependencies{
		runMenu: func(_ context.Context, _ Options, status string) (launcherAction, error) {
			statuses = append(statuses, status)
			action := actions[0]
			actions = actions[1:]
			return action, nil
		},
		runConfigure: func(context.Context, Options) error {
			return errors.New("configuration cancelled before a successful save")
		},
	}
	if err := runLauncher(context.Background(), Options{}, dependencies); err != nil {
		t.Fatal(err)
	}
	if want := []string{"", ""}; !reflect.DeepEqual(statuses, want) {
		t.Fatalf("statuses = %v, want %v", statuses, want)
	}
}

func TestLauncherSurfacesRunnerAndShutdownErrors(t *testing.T) {
	actions := []launcherAction{launcherAPIs, launcherConfigure, launcherExit}
	var statuses []string
	dependencies := launcherDependencies{
		runMenu: func(_ context.Context, _ Options, status string) (launcherAction, error) {
			statuses = append(statuses, status)
			action := actions[0]
			actions = actions[1:]
			return action, nil
		},
		openRuntime: func(context.Context) (*bootstrap.Runtime, error) {
			return &bootstrap.Runtime{}, nil
		},
		runWorkspace: func(context.Context, *bootstrap.Runtime, Options) error {
			return errors.New("workspace failed")
		},
		closeRuntime: func(*bootstrap.Runtime) error { return errors.New("close failed") },
		runConfigure: func(context.Context, Options) error { return errors.New("save failed") },
	}
	if err := runLauncher(context.Background(), Options{}, dependencies); err != nil {
		t.Fatal(err)
	}
	want := []string{"", "APIs: workspace failed; shutdown: close failed", "Configure: save failed"}
	if !reflect.DeepEqual(statuses, want) {
		t.Fatalf("statuses = %v, want %v", statuses, want)
	}
}

func TestLauncherDoesNotRunOrCloseWorkspaceWhenOpenFails(t *testing.T) {
	actions := []launcherAction{launcherAPIs, launcherExit}
	var status string
	dependencies := launcherDependencies{
		runMenu: func(_ context.Context, _ Options, got string) (launcherAction, error) {
			status = got
			action := actions[0]
			actions = actions[1:]
			return action, nil
		},
		openRuntime: func(context.Context) (*bootstrap.Runtime, error) {
			return nil, errors.New("open failed")
		},
		runWorkspace: func(context.Context, *bootstrap.Runtime, Options) error {
			t.Fatal("workspace ran after open failure")
			return nil
		},
		closeRuntime: func(*bootstrap.Runtime) error {
			t.Fatal("runtime closed after open failure")
			return nil
		},
	}
	if err := runLauncher(context.Background(), Options{}, dependencies); err != nil {
		t.Fatal(err)
	}
	if status != "APIs: open failed" {
		t.Fatalf("status = %q", status)
	}
}

func TestLauncherMenuUsesNormalScreenOptions(t *testing.T) {
	options := Options{}
	if got := len(tuiProgramOptions(context.Background(), options, false)); got != 3 {
		t.Fatalf("normal-screen option count = %d, want context, input, output", got)
	}
	if got := len(tuiProgramOptions(context.Background(), options, true)); got != 4 {
		t.Fatalf("alternate-screen option count = %d, want context, input, output, alternate screen", got)
	}
}
