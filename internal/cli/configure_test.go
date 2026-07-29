package cli

import (
	"archive/zip"
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
	"github.com/sairaph/apis-mcp/internal/config"
	"github.com/sairaph/apis-mcp/internal/docpacks"
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

func TestSetupRunResultWaitsForApplyWorker(t *testing.T) {
	state := newSetupModel(context.Background(), Options{}, config.Paths{}, config.Default(), nil)
	done := make(chan struct{})
	state.packApplyDone = done
	returned := make(chan error, 1)
	go func() { returned <- setupRunResult(context.Background(), state, nil) }()
	select {
	case <-returned:
		t.Fatal("setup returned before its apply worker stopped")
	case <-time.After(25 * time.Millisecond):
	}
	close(done)
	select {
	case <-returned:
	case <-time.After(time.Second):
		t.Fatal("setup did not return after its apply worker stopped")
	}
}

func TestSetupRunResultCancelsWorkerAfterProgramError(t *testing.T) {
	state := newSetupModel(context.Background(), Options{}, config.Paths{}, config.Default(), nil)
	workerCtx, cancel := context.WithCancel(context.Background())
	state.packApplyStop = cancel
	done := make(chan struct{})
	state.packApplyDone = done
	go func() {
		<-workerCtx.Done()
		close(done)
	}()
	programErr := errors.New("terminal input failed")
	if err := setupRunResult(context.Background(), state, programErr); !errors.Is(err, programErr) {
		t.Fatalf("setup program error = %v, want %v", err, programErr)
	}
	if workerCtx.Err() == nil {
		t.Fatal("setup program error did not cancel its apply worker")
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
		err: errors.New("Failed Client: client update failed"),
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
	if !strings.Contains(view, "client update failed") {
		t.Fatalf("failed setup omitted its error:\n%s", view)
	}
	if strings.Count(view, "Failed Client") != 1 {
		t.Fatalf("failed setup duplicated its client error:\n%s", view)
	}
}

func TestSetupPackStepRefreshesAndFreshInstallSelectsNone(t *testing.T) {
	catalog := setupTestCatalog(4)
	var requests atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		requests.Add(1)
		_ = json.NewEncoder(writer).Encode(catalog)
	}))
	defer server.Close()
	manager, err := docpacks.Open(t.TempDir(), docpacks.Options{CatalogURL: server.URL + "/catalog.json"})
	if err != nil {
		t.Fatal(err)
	}
	state, err := newSetupModelWithPacks(context.Background(), Options{}, config.Paths{}, config.Default(), nil, manager)
	if err != nil {
		t.Fatal(err)
	}
	if requests.Load() != 0 {
		t.Fatal("opening the setup model performed network I/O")
	}
	_, command := state.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if state.step != setupAPIs || !state.packLoading || command == nil {
		t.Fatalf("client step did not start API refresh: step=%d loading=%t command=%v", state.step, state.packLoading, command)
	}
	if requests.Load() != 0 {
		t.Fatal("API refresh ran synchronously")
	}
	state.Update(command())
	if requests.Load() != 1 || !state.packCatalogOK || len(state.packCatalog.Packs) != len(catalog.Packs) {
		t.Fatalf("catalog refresh mismatch: requests=%d catalog=%+v error=%v", requests.Load(), state.packCatalog, state.packError)
	}
	for _, pack := range catalog.Packs {
		if state.packSelected[pack.ID] {
			t.Fatalf("fresh setup preselected %q", pack.ID)
		}
	}
	state.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if state.step != setupSummary {
		t.Fatalf("API step did not continue to summary: %d", state.step)
	}
}

func TestSetupPreselectsInstalledPacksAndKeepsFocusOnRefresh(t *testing.T) {
	root := t.TempDir()
	manager, err := docpacks.Open(root, docpacks.Options{})
	if err != nil {
		t.Fatal(err)
	}
	catalog := setupTestCatalog(3)
	installed := catalog.Packs[1]
	active := docpacks.ActiveState{SchemaVersion: 1, Packs: map[string]docpacks.Pack{installed.ID: installed}}
	raw, err := json.Marshal(active)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "active.json"), raw, 0o600); err != nil {
		t.Fatal(err)
	}
	state, err := newSetupModelWithPacks(context.Background(), Options{}, config.Paths{}, config.Default(), nil, manager)
	if err != nil {
		t.Fatal(err)
	}
	state.packRefreshID = 1
	state.Update(setupPacksRefreshedMsg{id: 1, catalog: catalog})
	if !state.packSelected[installed.ID] || state.packStatus(installed) != "installed" {
		t.Fatalf("installed pack was not preselected: selected=%v status=%s", state.packSelected, state.packStatus(installed))
	}
	state.packCursor = 1
	reordered := catalog
	reordered.Packs = []docpacks.Pack{catalog.Packs[1], catalog.Packs[2], catalog.Packs[0]}
	state.packRefreshID = 2
	state.Update(setupPacksRefreshedMsg{id: 2, catalog: reordered})
	if state.packCursor != 0 || state.focusedPack() == nil || state.focusedPack().ID != installed.ID {
		t.Fatalf("refresh did not retain focused pack ID: cursor=%d focused=%+v", state.packCursor, state.focusedPack())
	}
}

func TestSetupAppendsUnlistedActivePacksInStableOrder(t *testing.T) {
	root := t.TempDir()
	manager, err := docpacks.Open(root, docpacks.Options{})
	if err != nil {
		t.Fatal(err)
	}
	all := setupTestCatalog(4)
	active := docpacks.ActiveState{SchemaVersion: 1, Packs: map[string]docpacks.Pack{
		all.Packs[3].ID: all.Packs[3],
		all.Packs[1].ID: all.Packs[1],
	}}
	raw, err := json.Marshal(active)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "active.json"), raw, 0o600); err != nil {
		t.Fatal(err)
	}
	state, err := newSetupModelWithPacks(context.Background(), Options{}, config.Paths{}, config.Default(), nil, manager)
	if err != nil {
		t.Fatal(err)
	}
	state.replacePackCatalog(docpacks.Catalog{SchemaVersion: 1, Packs: []docpacks.Pack{all.Packs[0]}})
	got := []string{state.packCatalog.Packs[0].ID, state.packCatalog.Packs[1].ID, state.packCatalog.Packs[2].ID}
	want := []string{all.Packs[0].ID, all.Packs[1].ID, all.Packs[3].ID}
	if !reflect.DeepEqual(got, want) || !state.packSelected[all.Packs[1].ID] || !state.packSelected[all.Packs[3].ID] {
		t.Fatalf("unlisted merge order/selection = %v/%v, want %v selected", got, state.packSelected, want)
	}
}

func TestSetupPackGridResponsiveColumnsAndRenderWidth(t *testing.T) {
	state := newSetupModel(context.Background(), Options{}, config.Paths{}, config.Default(), nil)
	state.step = setupAPIs
	state.replacePackCatalog(setupTestCatalog(8))
	for _, test := range []struct {
		width   int
		columns int
	}{{120, 3}, {119, 2}, {80, 2}, {79, 1}, {60, 1}} {
		state.Update(tea.WindowSizeMsg{Width: test.width, Height: 20})
		if got := state.packColumns(); got != test.columns {
			t.Fatalf("width %d columns = %d, want %d", test.width, got, test.columns)
		}
		for _, line := range strings.Split(state.View(), "\n") {
			if got := lipgloss.Width(line); got > test.width {
				t.Fatalf("width %d rendered line width %d: %q", test.width, got, ansi.Strip(line))
			}
		}
	}
}

func TestSetupPackGridSpatialMovementAndRowScrolling(t *testing.T) {
	state := newSetupModel(context.Background(), Options{}, config.Paths{}, config.Default(), nil)
	state.step = setupAPIs
	state.replacePackCatalog(setupTestCatalog(31))
	state.Update(tea.WindowSizeMsg{Width: 120, Height: 20})
	state.Update(tea.KeyMsg{Type: tea.KeyRight})
	state.Update(setupRuneKey("j"))
	if state.packCursor != 4 {
		t.Fatalf("right/down cursor = %d, want 4", state.packCursor)
	}
	state.Update(setupRuneKey("h"))
	state.Update(tea.KeyMsg{Type: tea.KeyUp})
	if state.packCursor != 0 {
		t.Fatalf("left/up cursor = %d, want 0", state.packCursor)
	}
	state.packCursor = 2
	state.Update(tea.KeyMsg{Type: tea.KeyRight})
	if state.packCursor != 2 {
		t.Fatalf("right crossed a row boundary: %d", state.packCursor)
	}
	for range 10 {
		state.Update(setupRuneKey("j"))
	}
	if state.packCursor != 30 || state.packScrollRow == 0 {
		t.Fatalf("vertical navigation did not scroll complete rows: cursor=%d scroll=%d", state.packCursor, state.packScrollRow)
	}
	state.replacePackCatalog(setupTestCatalog(5))
	state.packCursor = 2
	state.Update(tea.KeyMsg{Type: tea.KeyDown})
	if state.packCursor != 4 {
		t.Fatalf("down did not clamp to ragged final row: cursor=%d, want 4", state.packCursor)
	}
}

func TestSetupPackToggleAllIncludesOffscreenCatalog(t *testing.T) {
	state := newSetupModel(context.Background(), Options{}, config.Paths{}, config.Default(), nil)
	state.step = setupAPIs
	state.replacePackCatalog(setupTestCatalog(30))
	state.Update(tea.WindowSizeMsg{Width: 120, Height: 20})
	state.Update(setupRuneKey("a"))
	for _, pack := range state.packCatalog.Packs {
		if !state.packSelected[pack.ID] {
			t.Fatalf("offscreen pack %q was not selected", pack.ID)
		}
	}
	state.Update(setupRuneKey("a"))
	for _, pack := range state.packCatalog.Packs {
		if state.packSelected[pack.ID] {
			t.Fatalf("pack %q was not cleared", pack.ID)
		}
	}
}

func TestSetupPackRefreshErrorCanRetryContinueOrGoBack(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		http.Error(writer, "offline", http.StatusServiceUnavailable)
	}))
	defer server.Close()
	manager, err := docpacks.Open(t.TempDir(), docpacks.Options{CatalogURL: server.URL + "/catalog.json"})
	if err != nil {
		t.Fatal(err)
	}
	state, err := newSetupModelWithPacks(context.Background(), Options{}, config.Paths{}, config.Default(), nil, manager)
	if err != nil {
		t.Fatal(err)
	}
	_, command := state.Update(tea.KeyMsg{Type: tea.KeyEnter})
	state.Update(command())
	view := ansi.Strip(state.View())
	for _, wanted := range []string{"Catalog unavailable", "r retry", "enter continue", "esc back"} {
		if !strings.Contains(view, wanted) {
			t.Fatalf("unavailable catalog view missing %q:\n%s", wanted, view)
		}
	}
	state.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if state.step != setupSummary {
		t.Fatalf("catalog error blocked continue: step=%d", state.step)
	}
	state.Update(tea.KeyMsg{Type: tea.KeyEsc})
	_, retry := state.Update(setupRuneKey("r"))
	if state.step != setupAPIs || !state.packLoading || retry == nil {
		t.Fatal("catalog error state could not retry")
	}
	state.Update(tea.KeyMsg{Type: tea.KeyEsc})
	if state.step != setupClients {
		t.Fatalf("API step did not return to clients: %d", state.step)
	}
}

func TestSetupStalledPackRefreshLocksSelectionAndCanBeSkipped(t *testing.T) {
	current := setupTestCatalog(2)
	refreshed := setupTestCatalog(3)
	started := make(chan struct{})
	release := make(chan struct{})
	var releaseOnce sync.Once
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		close(started)
		select {
		case <-release:
			_ = json.NewEncoder(writer).Encode(refreshed)
		case <-request.Context().Done():
		}
	}))
	t.Cleanup(func() {
		releaseOnce.Do(func() { close(release) })
		server.Close()
	})
	manager, err := docpacks.Open(t.TempDir(), docpacks.Options{CatalogURL: server.URL + "/catalog.json"})
	if err != nil {
		t.Fatal(err)
	}
	state, err := newSetupModelWithPacks(context.Background(), Options{}, config.Paths{}, config.Default(), nil, manager)
	if err != nil {
		t.Fatal(err)
	}
	state.step = setupAPIs
	state.replacePackCatalog(current)
	state.packSelected[current.Packs[0].ID] = true
	command := state.beginPackRefresh()
	result := make(chan tea.Msg, 1)
	go func() { result <- command() }()
	select {
	case <-started:
	case <-time.After(5 * time.Second):
		t.Fatal("pack refresh request did not start")
	}

	state.Update(tea.KeyMsg{Type: tea.KeySpace})
	if !state.packSelected[current.Packs[0].ID] {
		t.Fatal("space changed selection while refresh was loading")
	}
	state.Update(setupRuneKey("a"))
	if !state.packSelected[current.Packs[0].ID] || state.packSelected[current.Packs[1].ID] {
		t.Fatalf("all changed selection while refresh was loading: %v", state.packSelected)
	}
	refreshID := state.packRefreshID
	state.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if state.step != setupSummary || state.packLoading || state.packRefreshID == refreshID {
		t.Fatalf("enter did not skip stalled refresh: step=%d loading=%t refreshID=%d", state.step, state.packLoading, state.packRefreshID)
	}

	select {
	case message := <-result:
		state.Update(message)
	case <-time.After(5 * time.Second):
		t.Fatal("released pack refresh did not return")
	}
	if len(state.packCatalog.Packs) != len(current.Packs) || !state.packSelected[current.Packs[0].ID] || state.packSelected[current.Packs[1].ID] {
		t.Fatalf("stale refresh changed current selection/catalog: catalog=%+v selected=%v", state.packCatalog, state.packSelected)
	}
}

func TestSetupSummaryIncludesPackCountAndDownloadBytes(t *testing.T) {
	state := newSetupModel(context.Background(), Options{}, config.Paths{}, config.Default(), nil)
	state.packManager, _ = docpacks.Open(t.TempDir(), docpacks.Options{})
	catalog := setupTestCatalog(2)
	catalog.Packs[0].Bytes = 1024
	catalog.Packs[1].Bytes = 2048
	state.replacePackCatalog(catalog)
	state.packActive[catalog.Packs[0].ID] = catalog.Packs[0]
	state.packSelected[catalog.Packs[0].ID] = true
	state.packSelected[catalog.Packs[1].ID] = true
	state.step = setupSummary
	view := ansi.Strip(state.View())
	for _, wanted := range []string{"API packs: 2 selected", "1 downloads", "2 KiB"} {
		if !strings.Contains(view, wanted) {
			t.Fatalf("summary missing %q:\n%s", wanted, view)
		}
	}
}

func TestSetupApplyingRendersLivePackAndIndexProgress(t *testing.T) {
	manager, err := docpacks.Open(t.TempDir(), docpacks.Options{})
	if err != nil {
		t.Fatal(err)
	}
	pack := setupTestCatalog(1).Packs[0]
	state := newSetupModel(context.Background(), Options{}, config.Paths{}, config.Default(), nil)
	state.packManager = manager
	state.packCatalogOK = true
	state.step = setupApplying
	state.packApplyChanged = true
	state.packApplyID = 7
	state.packApplyOrder = []string{pack.ID}
	state.packApplyEvents = map[string]docpacks.ApplyEvent{pack.ID: {
		Stage: docpacks.ApplyStageWaiting, PackID: pack.ID, PackName: pack.Name,
		PackBytesTotal: 1024, PreparedBytesTotal: 1024,
	}}
	if view := ansi.Strip(state.View()); !strings.Contains(view, "waiting") {
		t.Fatalf("initial apply view did not show waiting pack:\n%s", view)
	}
	state.Update(setupApplyProgressMsg{id: 7, event: docpacks.ApplyEvent{
		Stage: docpacks.ApplyStageReady, PackID: pack.ID, PackName: pack.Name, Cached: true,
		PackBytesDone: 1024, PackBytesTotal: 1024, PreparedBytesDone: 1024, PreparedBytesTotal: 1024,
	}})
	if view := ansi.Strip(state.View()); !strings.Contains(view, "cached") {
		t.Fatalf("cached apply view did not identify cache reuse:\n%s", view)
	}

	for _, test := range []struct {
		stage  docpacks.ApplyStage
		done   int64
		wanted string
	}{
		{stage: docpacks.ApplyStageDownloading, done: 512, wanted: "downloading"},
		{stage: docpacks.ApplyStageVerifying, done: 1024, wanted: "verifying"},
	} {
		state.Update(setupApplyProgressMsg{id: 7, event: docpacks.ApplyEvent{
			Stage: test.stage, PackID: pack.ID, PackName: pack.Name,
			PackBytesDone: test.done, PackBytesTotal: 1024,
			PreparedBytesDone: test.done, PreparedBytesTotal: 1024,
		}})
		view := ansi.Strip(state.View())
		for _, wanted := range []string{pack.Name, test.wanted, formatSetupBytes(test.done) + " / 1 KiB", "Prepared", "["} {
			if !strings.Contains(view, wanted) {
				t.Fatalf("%s apply view missing %q:\n%s", test.stage, wanted, view)
			}
		}
	}

	state.Update(setupApplyProgressMsg{id: 7, event: docpacks.ApplyEvent{
		Stage: docpacks.ApplyStageReady, PackID: pack.ID, PackName: pack.Name,
		PackBytesDone: 1024, PackBytesTotal: 1024, PreparedBytesDone: 1024, PreparedBytesTotal: 1024,
	}})
	state.Update(setupApplyProgressMsg{id: 7, event: docpacks.ApplyEvent{
		Stage: docpacks.ApplyStageIndexing, PreparedBytesDone: 1024, PreparedBytesTotal: 1024,
	}})
	if view := ansi.Strip(state.View()); !strings.Contains(view, "Indexing documentation library") || !strings.Contains(view, "ready") {
		t.Fatalf("indexing view did not separate indexing from downloads:\n%s", view)
	}
	state.packApplyOrder = nil
	state.packApplyEvents = map[string]docpacks.ApplyEvent{}
	state.packRemovalOnly = true
	if view := ansi.Strip(state.View()); !strings.Contains(view, "removal-only rebuild") || !strings.Contains(view, "0 B / 0 B") {
		t.Fatalf("removal-only apply view is unclear:\n%s", view)
	}
}

func TestSetupApplyingRejectsStaleEventsAndResults(t *testing.T) {
	state := newSetupModel(context.Background(), Options{}, config.Paths{}, config.Default(), nil)
	state.step = setupApplying
	state.packApplyID = 2
	state.packApplyEvents = map[string]docpacks.ApplyEvent{"current": {
		Stage: docpacks.ApplyStageDownloading, PackID: "current", PackName: "Current",
		PackBytesDone: 10, PackBytesTotal: 100,
	}}
	state.Update(setupApplyProgressMsg{id: 1, event: docpacks.ApplyEvent{
		Stage: docpacks.ApplyStageReady, PackID: "current", PackBytesDone: 100, PackBytesTotal: 100,
	}})
	state.Update(setupAppliedMsg{id: 1})
	event := state.packApplyEvents["current"]
	if state.step != setupApplying || state.saved || event.Stage != docpacks.ApplyStageDownloading || event.PackBytesDone != 10 {
		t.Fatalf("stale apply messages changed current state: step=%d saved=%t event=%+v", state.step, state.saved, event)
	}
}

func TestSetupUnchangedPacksSkipManagerApply(t *testing.T) {
	pack, _ := setupDownloadablePack(t)
	catalog := docpacks.Catalog{SchemaVersion: 1, Packs: []docpacks.Pack{pack}}
	var requests atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		requests.Add(1)
		http.Error(writer, "must not request", http.StatusInternalServerError)
	}))
	defer server.Close()
	root := t.TempDir()
	manager, err := docpacks.Open(root, docpacks.Options{CatalogURL: server.URL + "/catalog.json"})
	if err != nil {
		t.Fatal(err)
	}
	raw, err := json.Marshal(docpacks.ActiveState{SchemaVersion: 1, Packs: map[string]docpacks.Pack{pack.ID: pack}})
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "active.json"), raw, 0o600); err != nil {
		t.Fatal(err)
	}
	state, err := newSetupModelWithPacks(context.Background(), Options{}, config.Paths{}, config.Default(), nil, manager)
	if err != nil {
		t.Fatal(err)
	}
	state.replacePackCatalog(catalog)
	state.step = setupApplying
	command := state.applySetup()
	message := command()
	phase, ok := message.(setupApplyPhaseMsg)
	if !ok || phase.phase != setupApplySaving {
		t.Fatalf("unchanged apply first message = %#v, want saving phase", message)
	}
	_, command = state.Update(message)
	result := setupRunApplyCommand(t, state, command)
	if result.err != nil || result.packs.changed || !result.packs.attempted || requests.Load() != 0 {
		t.Fatalf("unchanged fast path result=%+v requests=%d", result, requests.Load())
	}
	if view := ansi.Strip(state.View()); !strings.Contains(view, "skipping archive verification") {
		t.Fatalf("unchanged fast path view is unclear:\n%s", view)
	}
}

func TestSetupApplyCancellationBoundary(t *testing.T) {
	before := newSetupModel(context.Background(), Options{}, config.Paths{}, config.Default(), nil)
	before.step, before.packApplyID = setupApplying, 1
	before.packApplyCanStop = &atomic.Bool{}
	before.packApplyCanStop.Store(true)
	var beforeStops atomic.Int32
	before.packApplyStop = func() { beforeStops.Add(1) }
	before.Update(tea.KeyMsg{Type: tea.KeyCtrlC})
	if !before.cancelled || beforeStops.Load() != 1 || !strings.Contains(before.message, "Cancelling") {
		t.Fatalf("pre-index cancellation = cancelled %t stops %d message %q", before.cancelled, beforeStops.Load(), before.message)
	}

	after := newSetupModel(context.Background(), Options{}, config.Paths{}, config.Default(), nil)
	after.step, after.packApplyID = setupApplying, 2
	after.packApplyPhase = setupApplyPublishing
	after.packApplyCanStop = &atomic.Bool{}
	var afterStops atomic.Int32
	after.packApplyStop = func() { afterStops.Add(1) }
	after.Update(tea.KeyMsg{Type: tea.KeyCtrlC})
	if after.cancelled || afterStops.Load() != 0 || after.message != "Finishing safely; this stage cannot be cancelled." {
		t.Fatalf("post-index ctrl+c = cancelled %t stops %d message %q", after.cancelled, afterStops.Load(), after.message)
	}
	if view := ansi.Strip(after.View()); !strings.Contains(view, after.message) || !strings.Contains(view, "Publishing API pack selection") {
		t.Fatalf("safe finishing state not rendered:\n%s", view)
	}
	after.Update(setupAppliedMsg{id: 2})
	if after.cancelled || !after.saved || after.step != setupDone {
		t.Fatalf("published apply reported false cancellation: cancelled=%t saved=%t step=%d", after.cancelled, after.saved, after.step)
	}
}

func TestSetupApplyWaitHandlesClosedChannelAndParentCancellation(t *testing.T) {
	closed := make(chan tea.Msg)
	close(closed)
	if message := waitSetupApply(4, closed, context.Background())(); message != (setupApplyClosedMsg{id: 4}) {
		t.Fatalf("closed apply channel message = %#v", message)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	open := make(chan tea.Msg)
	if message := waitSetupApply(5, open, ctx)(); message != (setupApplyClosedMsg{id: 5}) {
		t.Fatalf("cancelled parent wait message = %#v", message)
	}
	state := newSetupModel(context.Background(), Options{}, config.Paths{}, config.Default(), nil)
	state.step, state.packApplyID = setupApplying, 4
	state.Update(setupApplyClosedMsg{id: 4})
	if state.step != setupDone || state.failure == nil {
		t.Fatalf("unexpected worker closure was not handled: step=%d failure=%v", state.step, state.failure)
	}
}

func TestSetupApplyingBoundsPackRowsAndRendersAllPhases(t *testing.T) {
	manager, err := docpacks.Open(t.TempDir(), docpacks.Options{})
	if err != nil {
		t.Fatal(err)
	}
	state := newSetupModel(context.Background(), Options{}, config.Paths{}, config.Default(), nil)
	state.packManager, state.packCatalogOK, state.packApplyChanged = manager, true, true
	state.step, state.width, state.height = setupApplying, 50, 10
	state.packApplyCanStop = &atomic.Bool{}
	state.packApplyEvents = make(map[string]docpacks.ApplyEvent)
	for index, pack := range setupTestCatalog(20).Packs {
		state.packApplyOrder = append(state.packApplyOrder, pack.ID)
		state.packApplyEvents[pack.ID] = docpacks.ApplyEvent{
			Stage: docpacks.ApplyStageReady, PackID: pack.ID, PackName: pack.Name,
			PackBytesDone: 1024, PackBytesTotal: 1024,
			PreparedBytesDone: int64(index+1) * 1024, PreparedBytesTotal: 20 * 1024,
		}
	}
	for phase, wanted := range map[setupApplyPhase]string{
		setupApplyIndexing:    "Indexing documentation library",
		setupApplyPublishing:  "Publishing API pack selection",
		setupApplySaving:      "Saving settings",
		setupApplyRegistering: "Registering clients",
		setupApplyFinishing:   "Finishing setup",
	} {
		state.packApplyPhase = phase
		view := ansi.Strip(state.View())
		if len(strings.Split(view, "\n")) > state.height || !strings.Contains(view, "19 more packs") || !strings.Contains(view, "Prepared 20 KiB / 20 KiB") || !strings.Contains(view, wanted) {
			t.Fatalf("bounded %d phase view is incomplete:\n%s", phase, view)
		}
		for _, line := range strings.Split(state.View(), "\n") {
			if lipgloss.Width(line) > state.width {
				t.Fatalf("applying line exceeds width %d: %q", state.width, ansi.Strip(line))
			}
		}
	}
}

func TestSetupControlCDuringApplyWaitsForWorkerCleanup(t *testing.T) {
	pack, _ := setupDownloadablePack(t)
	catalog := docpacks.Catalog{SchemaVersion: 1, Packs: []docpacks.Pack{pack}}
	requestStarted := make(chan struct{})
	requestDone := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		close(requestStarted)
		writer.WriteHeader(http.StatusOK)
		if flusher, ok := writer.(http.Flusher); ok {
			flusher.Flush()
		}
		<-request.Context().Done()
		close(requestDone)
	}))
	defer server.Close()
	manager, err := docpacks.Open(t.TempDir(), docpacks.Options{CatalogURL: server.URL + "/catalog.json"})
	if err != nil {
		t.Fatal(err)
	}
	root := t.TempDir()
	paths := config.Paths{Root: root, Config: filepath.Join(root, "config.toml")}
	state, err := newSetupModelWithPacks(context.Background(), Options{}, paths, config.Default(), nil, manager)
	if err != nil {
		t.Fatal(err)
	}
	state.replacePackCatalog(catalog)
	state.packSelected[pack.ID] = true
	state.settings.ListTokenBudget++
	state.step = setupApplying
	command := state.applySetup()
	message := command()
	_, command = state.Update(message)
	select {
	case <-requestStarted:
	case <-time.After(5 * time.Second):
		t.Fatal("apply request did not start")
	}
	_, quit := state.Update(tea.KeyMsg{Type: tea.KeyCtrlC})
	if quit != nil || state.step != setupApplying || !state.cancelled {
		t.Fatalf("ctrl+c quit before cleanup: command=%v step=%d cancelled=%t", quit, state.step, state.cancelled)
	}
	for command != nil {
		result := make(chan tea.Msg, 1)
		go func(current tea.Cmd) { result <- current() }(command)
		select {
		case message = <-result:
		case <-time.After(5 * time.Second):
			t.Fatal("apply worker did not finish after cancellation")
		}
		_, command = state.Update(message)
	}
	select {
	case <-requestDone:
	case <-time.After(5 * time.Second):
		t.Fatal("HTTP request remained active after setup quit")
	}
	if state.saved || state.step != setupApplying {
		t.Fatalf("cancelled apply rendered completion: saved=%t step=%d", state.saved, state.step)
	}
	if _, err := os.Stat(paths.Config); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("settings were saved before cancellable pack preparation finished: %v", err)
	}
}

func TestSetupPackApplyDownloadsRebuildsAndAllowsEmptySelection(t *testing.T) {
	pack, archive := setupDownloadablePack(t)
	catalog := docpacks.Catalog{SchemaVersion: 1, Packs: []docpacks.Pack{pack}}
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/catalog.json":
			_ = json.NewEncoder(writer).Encode(catalog)
		case "/" + pack.Asset:
			_, _ = writer.Write(archive)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()
	root := t.TempDir()
	paths := config.Paths{
		Root: root, Config: filepath.Join(root, "config.toml"), Library: filepath.Join(root, "library"),
		Packs: filepath.Join(root, "packs"), Index: filepath.Join(root, "index"),
	}
	userPage := filepath.Join(paths.Library, "user-api", "v1", "overview.md")
	setupWriteFile(t, filepath.Join(paths.Library, "user-api", "v1", "_index.md"), "---\nname: User API\nversion: v1\n---\n")
	setupWriteFile(t, userPage, "---\ntitle: User overview\n---\n\n# User overview\n")
	manager, err := docpacks.Open(paths.Packs, docpacks.Options{CatalogURL: server.URL + "/catalog.json"})
	if err != nil {
		t.Fatal(err)
	}
	state, err := newSetupModelWithPacks(context.Background(), Options{}, paths, config.Default(), nil, manager)
	if err != nil {
		t.Fatal(err)
	}
	state.replacePackCatalog(catalog)
	state.packSelected[pack.ID] = true
	message := setupRunApplyCommand(t, state, state.applySetup())
	if message.err != nil || !message.packs.attempted || message.packs.selected != 1 {
		t.Fatalf("pack apply result = %+v, error %v", message.packs, message.err)
	}
	active, err := manager.Active()
	if err != nil || len(active.Packs) != 1 {
		t.Fatalf("active packs after apply = %+v, %v", active, err)
	}
	entries, err := os.ReadDir(paths.Index)
	if err != nil || len(entries) == 0 {
		t.Fatalf("library rebuild did not publish an index: entries=%v error=%v", entries, err)
	}

	preserved, err := newSetupModelWithPacks(context.Background(), Options{}, paths, config.Default(), nil, manager)
	if err != nil {
		t.Fatal(err)
	}
	partial := setupTestCatalog(1)
	preserved.replacePackCatalog(partial)
	if len(preserved.packCatalog.Packs) != 2 || preserved.packCatalog.Packs[1].ID != pack.ID || !preserved.packUnlisted[pack.ID] || !preserved.packSelected[pack.ID] {
		t.Fatalf("partial catalog did not preserve active pack: catalog=%+v unlisted=%v selected=%v", preserved.packCatalog, preserved.packUnlisted, preserved.packSelected)
	}
	preserved.step, preserved.packCursor = setupAPIs, 1
	view := ansi.Strip(preserved.View())
	for _, wanted := range []string{"installed · unlisted", "1 unlisted retained"} {
		if !strings.Contains(view, wanted) {
			t.Fatalf("unlisted pack view missing %q:\n%s", wanted, view)
		}
	}
	message = setupRunApplyCommand(t, preserved, preserved.applySetup())
	if message.err != nil || message.packs.selected != 1 || message.packs.changed {
		t.Fatalf("partial catalog preservation apply = %+v, error %v", message.packs, message.err)
	}
	preserved.Update(message)
	if done := ansi.Strip(preserved.View()); !strings.Contains(done, "API packs: 1 active; unchanged.") {
		t.Fatalf("unchanged pack apply rendered the wrong result:\n%s", done)
	}
	active, err = manager.Active()
	if err != nil || len(active.Packs) != 1 || active.Packs[pack.ID].ID != pack.ID {
		t.Fatalf("partial catalog silently removed active pack: %+v, %v", active, err)
	}

	preserved.packCursor = 1
	preserved.replacePackCatalog(docpacks.Catalog{SchemaVersion: 1, Packs: []docpacks.Pack{}})
	if preserved.packCursor != 0 || preserved.focusedPack() == nil || preserved.focusedPack().ID != pack.ID {
		t.Fatalf("empty catalog did not keep unlisted focus stable: cursor=%d focused=%+v", preserved.packCursor, preserved.focusedPack())
	}
	message = setupRunApplyCommand(t, preserved, preserved.applySetup())
	active, err = manager.Active()
	if message.err != nil || message.packs.changed || err != nil || len(active.Packs) != 1 || active.Packs[pack.ID].ID != pack.ID {
		t.Fatalf("empty catalog silently changed active packs: result=%+v active=%+v errors=%v/%v", message.packs, active, message.err, err)
	}
	preserved.packSelected[pack.ID] = false
	message = setupRunApplyCommand(t, preserved, preserved.applySetup())
	if message.err != nil || !message.packs.attempted || message.packs.selected != 0 {
		t.Fatalf("empty pack apply result = %+v, error %v", message.packs, message.err)
	}
	active, err = manager.Active()
	if err != nil || len(active.Packs) != 0 {
		t.Fatalf("empty selection did not remove managed packs: %+v, %v", active, err)
	}
	if raw, err := os.ReadFile(userPage); err != nil || !strings.Contains(string(raw), "User overview") {
		t.Fatalf("pack apply changed user documentation: %q, %v", raw, err)
	}
}

func TestSetupPackApplyFailureStopsClientRegistration(t *testing.T) {
	root := t.TempDir()
	paths := config.Paths{Config: filepath.Join(root, "config.toml"), Packs: filepath.Join(root, "packs"), Index: filepath.Join(root, "index")}
	manager, err := docpacks.Open(paths.Packs, docpacks.Options{})
	if err != nil {
		t.Fatal(err)
	}
	clients := []install.Status{{Client: install.Client{ID: "not-a-client", Name: "Must Not Run"}, Detected: true}}
	state, err := newSetupModelWithPacks(context.Background(), Options{Executable: "/bin/apis-mcp"}, paths, config.Default(), clients, manager)
	if err != nil {
		t.Fatal(err)
	}
	state.packCatalog = setupTestCatalog(1)
	state.packCatalog.SchemaVersion = 0
	state.packCatalogOK = true
	state.packSelected[state.packCatalog.Packs[0].ID] = true
	message := setupRunApplyCommand(t, state, state.applySetup())
	if message.err == nil || message.packs.err == nil {
		t.Fatalf("invalid pack catalog unexpectedly applied: %+v", message)
	}
	if len(message.results) != 0 {
		t.Fatalf("client registration ran after pack failure: %+v", message.results)
	}
}

func setupTestCatalog(count int) docpacks.Catalog {
	packs := make([]docpacks.Pack, count)
	for index := range packs {
		id := fmt.Sprintf("api-%02d", index)
		hash := strings.Repeat(fmt.Sprintf("%x", index%16), 64)
		packs[index] = docpacks.Pack{
			ID: id, Name: fmt.Sprintf("API %02d", index), Description: fmt.Sprintf("Documentation for API %02d.", index),
			Asset: id + "-" + hash + ".zip", SHA256: hash, Bytes: int64(index+1) * 1024,
			UncompressedBytes: 2048, Files: 2, Pages: 1, Versions: []string{"v1"}, Collections: []string{"examples"},
		}
	}
	return docpacks.Catalog{SchemaVersion: 1, Packs: packs}
}

func setupDownloadablePack(t *testing.T) (docpacks.Pack, []byte) {
	t.Helper()
	manifest := "---\nname: Alpha\nversion: v1\ndescription: Alpha documentation.\ncollections: [examples]\n---\n"
	page := "---\ntitle: Overview\n---\n\n# Overview\n"
	var buffer bytes.Buffer
	writer := zip.NewWriter(&buffer)
	for name, content := range map[string]string{"alpha/v1/_index.md": manifest, "alpha/v1/overview.md": page} {
		file, err := writer.Create(name)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := file.Write([]byte(content)); err != nil {
			t.Fatal(err)
		}
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	archive := buffer.Bytes()
	hash := sha256.Sum256(archive)
	digest := hex.EncodeToString(hash[:])
	return docpacks.Pack{
		ID: "alpha", Name: "Alpha", Description: "Alpha documentation.", Asset: "alpha-" + digest + ".zip",
		SHA256: digest, Bytes: int64(len(archive)), UncompressedBytes: int64(len(manifest) + len(page)),
		Files: 2, Pages: 1, Versions: []string{"v1"}, Collections: []string{"examples"},
	}, archive
}

func setupWriteFile(t *testing.T, name, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(name), 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(name, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
}

func setupRuneKey(value string) tea.KeyMsg {
	return tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(value)}
}

func setupRunApplyCommand(t *testing.T, state *setupModel, command tea.Cmd) setupAppliedMsg {
	t.Helper()
	state.step = setupApplying
	for command != nil {
		message := command()
		if result, ok := message.(setupAppliedMsg); ok {
			return result
		}
		_, command = state.Update(message)
	}
	t.Fatal("apply command ended without a result")
	return setupAppliedMsg{}
}
