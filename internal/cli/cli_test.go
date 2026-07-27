package cli

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sairaph/apis-mcp/internal/bootstrap"
	"github.com/sairaph/apis-mcp/internal/config"
	"github.com/sairaph/apis-mcp/internal/httpcall"
	"github.com/sairaph/apis-mcp/library"
)

func TestParseCommands(t *testing.T) {
	tests := []struct {
		args []string
		want func(Command) bool
	}{
		{nil, func(c Command) bool { return c.Name == "interactive" }},
		{[]string{"list", "--name", "stripe", "--page", "2"}, func(c Command) bool { return c.NameFilter == "stripe" && c.Page == 2 }},
		{[]string{"pages", "stripe-v1", "--path", "payments", "--page=3"}, func(c Command) bool { return c.DocID == "stripe-v1" && c.Path == "payments" && c.Page == 3 }},
		{[]string{"search", "stripe-v1", "create payment", "--page", "2"}, func(c Command) bool { return c.Query == "create payment" && c.Page == 2 }},
		{[]string{"read", "stripe-v1", "charges", "--lines", "10:20"}, func(c Command) bool { return len(c.Lines) == 2 && c.Lines[0] == 10 && c.Lines[1] == 20 }},
		{[]string{"call", "POST", "https://example.test", "--payload", `{"ok":true}`, "--retries", "0", "--allow-large-download"}, func(c Command) bool {
			return c.Method == "POST" && c.Retries != nil && *c.Retries == 0 && c.AllowLargeDownload
		}},
		{[]string{"configure", "--client", "cursor", "--client", "opencode", "--dry-run"}, func(c Command) bool { return len(c.ClientIDs) == 2 && c.DryRun && c.Backup }},
		{[]string{"import", "markdown", "/tmp/docs"}, func(c Command) bool { return c.ImportKind == "markdown" && c.Source == "/tmp/docs" }},
		{[]string{"import", "openapi", "Pet API", "v1", "https://example.test/openapi.yaml"}, func(c Command) bool {
			return c.ImportKind == "openapi" && c.APIName == "Pet API" && c.Version == "v1"
		}},
		{[]string{"import", "llms", "Pet API", "v1", "llms.txt"}, func(c Command) bool { return c.ImportKind == "llms" && c.Source == "llms.txt" }},
		{[]string{"import", "docsify", "Docs", "v1", "https://example.test/#/"}, func(c Command) bool {
			return c.ImportKind == "docsify" && c.APIName == "Docs" && c.Version == "v1" && c.Source == "https://example.test/#/"
		}},
		{[]string{"import", "html", "Pet API", "v1", "https://example.test/docs", "--max-pages", "12", "--max-depth=2"}, func(c Command) bool {
			return c.ImportKind == "html" && c.Source == "https://example.test/docs" && c.MaxPages == 12 && c.MaxDepth == 2
		}},
	}
	for _, test := range tests {
		command, err := Parse(test.args)
		if err != nil {
			t.Errorf("Parse(%v): %v", test.args, err)
			continue
		}
		if !test.want(command) {
			t.Errorf("Parse(%v) = %+v", test.args, command)
		}
	}
}

func TestConfigureOnlyTreatsProvidedFlagsAsExplicit(t *testing.T) {
	bare, err := Parse([]string{"configure"})
	if err != nil || bare.Explicit {
		t.Fatalf("bare configure parsed as explicit: %+v (%v)", bare, err)
	}
	flagged, err := Parse([]string{"configure", "--backup=true"})
	if err != nil || !flagged.Explicit {
		t.Fatalf("flagged configure was not explicit: %+v (%v)", flagged, err)
	}
}

func TestConfigureModelValidatesEdits(t *testing.T) {
	cfg := config.Default()
	if err := setConfigValue(&cfg, "maximum_redirects", "21"); err == nil {
		t.Fatal("invalid edit was accepted")
	}
	if cfg.MaximumRedirects != config.Default().MaximumRedirects {
		t.Fatal("invalid edit mutated config")
	}
	if err := setConfigValue(&cfg, "maximum_redirects", "12"); err != nil {
		t.Fatal(err)
	}
	if cfg.MaximumRedirects != 12 {
		t.Fatalf("valid edit was not applied: %+v", cfg)
	}
}

func TestConfigureModelConfirmsAndRestoresDefaults(t *testing.T) {
	root := t.TempDir()
	paths := config.Paths{Root: root, Config: filepath.Join(root, "config.toml")}
	cfg := config.Default()
	cfg.RetentionHours = 99
	m := newConfigureModel(context.Background(), Options{Executable: "/bin/apis-mcp"}, paths, cfg, nil)

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'d'}})
	m = updated.(configureModel)
	if !m.confirm {
		t.Fatal("restore did not request confirmation")
	}
	updated, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'n'}})
	m = updated.(configureModel)
	if m.confirm || m.cfg.RetentionHours != 99 {
		t.Fatal("cancelled restore changed config")
	}

	updated, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'d'}})
	m = updated.(configureModel)
	updated, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'y'}})
	m = updated.(configureModel)
	if cmd == nil {
		t.Fatal("confirmed restore did not schedule a save")
	}
	updated, _ = m.Update(cmd())
	m = updated.(configureModel)
	if m.cfg.RetentionHours != config.Default().RetentionHours {
		t.Fatalf("defaults were not restored: %+v", m.cfg)
	}
	loaded, err := config.Load(paths)
	if err != nil {
		t.Fatal(err)
	}
	if loaded.RetentionHours != config.Default().RetentionHours {
		t.Fatalf("restored defaults were not saved: %+v", loaded)
	}
}

func TestParseRejectsInvalidInput(t *testing.T) {
	for _, args := range [][]string{{"unknown"}, {"read", "doc"}, {"read", "doc", "page", "--lines", "5:2"}, {"cache", "clear"}, {"import"}, {"import", "browser", "docs"}, {"import", "openapi", "name"}, {"import", "html", "name", "v1", "https://example.test", "--max-pages", "0"}} {
		if command, err := Parse(args); err == nil {
			t.Errorf("Parse(%v) unexpectedly succeeded: %+v", args, command)
		}
	}
}

func TestHumanRenderingIsNotMCPDocument(t *testing.T) {
	result := library.ListResult{Pagination: library.Pagination{Page: 1, Total: 1, TotalPages: 1}, APIs: []library.API{{Name: "Stripe", Versions: []library.APIVersion{{Version: "v1", DocID: "stripe-v1", Pages: 12}}}}}
	var output bytes.Buffer
	if err := RenderHuman(&output, result); err != nil {
		t.Fatal(err)
	}
	got := output.String()
	if strings.HasPrefix(got, "---") || !strings.Contains(got, "API library") || !strings.Contains(got, "stripe-v1") {
		t.Fatalf("unexpected human rendering:\n%s", got)
	}
}

func TestHumanHTTPRenderingUsesPublishedOrExpectedBodyPath(t *testing.T) {
	for _, test := range []struct {
		name   string
		result httpcall.Result
		want   string
	}{
		{name: "complete", result: httpcall.Result{Cache: httpcall.CacheResult{BodyPath: "/cache/body.json", FinalPath: "/cache/stale.json"}}, want: "/cache/body.json"},
		{name: "background", result: httpcall.Result{Cache: httpcall.CacheResult{FinalPath: "/cache/future.json"}}, want: "/cache/future.json"},
	} {
		t.Run(test.name, func(t *testing.T) {
			var output bytes.Buffer
			if err := RenderHuman(&output, test.result); err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(output.String(), "body: "+test.want) {
				t.Fatalf("HTTP output = %q", output.String())
			}
		})
	}
}

func TestHelpListsCoreCommandsAndClients(t *testing.T) {
	help := Help("v1.2.3")
	for _, wanted := range []string{"apis-mcp v1.2.3", "collections", "sessions", "configure", "import markdown", "import openapi", "import llms", "import html", "import docsify", "claude-desktop", "opencode"} {
		if !strings.Contains(help, wanted) {
			t.Errorf("help missing %q", wanted)
		}
	}
}

func TestExecuteMarkdownImportRendersHumanResult(t *testing.T) {
	source := t.TempDir()
	if err := os.WriteFile(filepath.Join(source, "_index.md"), []byte("---\nname: CLI API\nversion: v1\n---\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(source, "overview.md"), []byte("---\ntitle: Overview\n---\n\n# Overview\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	root := t.TempDir()
	runtime := &bootstrap.Runtime{
		Paths:  config.Paths{Library: filepath.Join(root, "library"), Index: filepath.Join(root, "index")},
		Config: config.Default(),
	}
	var output bytes.Buffer
	if err := Execute(context.Background(), runtime, []string{"import", "markdown", source}, Options{Stdout: &output}); err != nil {
		t.Fatal(err)
	}
	for _, wanted := range []string{"Imported CLI API v1 (markdown): 1 pages", filepath.Join(root, "library", "cli-api", "v1")} {
		if !strings.Contains(output.String(), wanted) {
			t.Fatalf("output missing %q:\n%s", wanted, output.String())
		}
	}
}
