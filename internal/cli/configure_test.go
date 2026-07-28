package cli

import (
	"context"
	"errors"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/ansi"
	"github.com/sairaph/apis-mcp/internal/config"
	"github.com/sairaph/apis-mcp/internal/install"
)

func TestSetupFlowIsUnframedAndStepwise(t *testing.T) {
	state := newSetupModel(context.Background(), Options{}, config.Paths{}, config.Default(), []install.Status{
		{Client: install.Client{ID: "detected", Name: "Detected Client"}, Detected: true},
		{Client: install.Client{ID: "hidden", Name: "Hidden Client"}},
	})

	clients := ansi.Strip(state.View())
	for _, unwanted := range []string{"╭", "╮", "╰", "╯", "│", " Settings ", "Application settings", "Maintenance", "Ready"} {
		if strings.Contains(clients, unwanted) {
			t.Fatalf("setup client step contains main-application chrome %q:\n%s", unwanted, clients)
		}
	}
	for _, wanted := range []string{
		"apis-mcp setup",
		"AI clients — which should be able to use your APIs?",
		"Detected Client",
		"enter continue · q cancel",
	} {
		if !strings.Contains(clients, wanted) {
			t.Fatalf("setup client step is missing %q:\n%s", wanted, clients)
		}
	}

	state.Update(tea.KeyMsg{Type: tea.KeyEnter})
	summary := ansi.Strip(state.View())
	if state.step != setupSummary || !strings.Contains(summary, "API tool configuration") || strings.Contains(summary, "AI clients —") {
		t.Fatalf("setup did not advance to its separate summary step:\n%s", summary)
	}

	state.cursor = 1
	state.Update(tea.KeyMsg{Type: tea.KeyEnter})
	settings := ansi.Strip(state.View())
	if state.step != setupSettings || !strings.Contains(settings, "\nSettings\n") || strings.Contains(settings, "API tool configuration") {
		t.Fatalf("setup did not open settings as a separate step:\n%s", settings)
	}
}

func TestSetupCancellationCannotRenderCompletion(t *testing.T) {
	state := newSetupModel(context.Background(), Options{}, config.Paths{}, config.Default(), nil)
	_, command := state.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}})
	if !state.cancelled || command == nil {
		t.Fatal("setup cancellation did not quit the program")
	}
	if err := setupRunResult(context.Background(), state, nil); err == nil {
		t.Fatal("setup cancellation was reported as process success")
	}
	view := ansi.Strip(state.View())
	for _, misleading := range []string{"Settings were saved", "enter to finish", "Run `apis-mcp` to browse"} {
		if strings.Contains(view, misleading) {
			t.Fatalf("cancelled setup rendered completion text %q", misleading)
		}
	}
}

func TestSetupCannotSelectHiddenClients(t *testing.T) {
	state := newSetupModel(context.Background(), Options{}, config.Paths{}, config.Default(), []install.Status{
		{Client: install.Client{ID: "hidden", Name: "Hidden Client"}},
	})
	if state.cursor != -1 {
		t.Fatalf("empty visible client list cursor = %d", state.cursor)
	}
	state.Update(tea.KeyMsg{Type: tea.KeySpace})
	state.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'a'}})
	if state.selected[0] {
		t.Fatal("hidden undetected client was selected without being revealed")
	}
	state.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'v'}})
	state.Update(tea.KeyMsg{Type: tea.KeySpace})
	if !state.selected[0] {
		t.Fatal("revealed client could not be selected")
	}
}

func TestSetupDoneOnlyFollowsSuccessfulApply(t *testing.T) {
	state := newSetupModel(context.Background(), Options{}, config.Paths{}, config.Default(), nil)
	state.Update(setupAppliedMsg{})
	if state.step != setupDone || !state.saved || state.failure != nil {
		t.Fatalf("successful apply state = step %d saved %t failure %v", state.step, state.saved, state.failure)
	}
	view := ansi.Strip(state.View())
	if !strings.Contains(view, "Settings were saved") || !strings.Contains(view, "enter to finish") {
		t.Fatalf("successful setup summary is incomplete:\n%s", view)
	}
}

func TestSetupFailureDoesNotRenderSuccessGuidance(t *testing.T) {
	state := newSetupModel(context.Background(), Options{}, config.Paths{}, config.Default(), nil)
	state.Update(setupAppliedMsg{
		err: errors.New("save failed"),
		results: []clientApplyResult{{
			client: install.Client{Name: "Failed Client"}, changed: true, err: errors.New("client update failed"),
		}},
	})
	view := ansi.Strip(state.View())
	for _, misleading := range []string{"Settings were saved", "Run `apis-mcp` to browse", "Restart affected clients"} {
		if strings.Contains(view, misleading) {
			t.Fatalf("failed setup rendered success guidance %q:\n%s", misleading, view)
		}
	}
	if !strings.Contains(view, "save failed") {
		t.Fatalf("failed setup omitted its error:\n%s", view)
	}
}
