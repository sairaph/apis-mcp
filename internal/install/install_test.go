package install

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pelletier/go-toml/v2"
	"github.com/tailscale/hujson"
	"gopkg.in/yaml.v3"
)

func TestRegistryHasThirteenClients(t *testing.T) {
	clients := Registry("/home/test", "linux")
	if len(clients) != 13 {
		t.Fatalf("got %d clients", len(clients))
	}
	seen := map[string]bool{}
	for _, client := range clients {
		if seen[client.ID] || client.ConfigPath == "" {
			t.Fatalf("invalid client: %+v", client)
		}
		seen[client.ID] = true
	}
}

func TestJSONLargeIntegerIsPreserved(t *testing.T) {
	const large = "9007199254740993123456789"
	next, _, err := transform(FormatJSON, []byte(`{"unrelated":`+large+`}`), "/bin/apis-mcp", false, false)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(next), large) {
		t.Fatalf("large integer was corrupted: %s", next)
	}
}

func TestRejectsNullAndUnsafeTopLevelShapes(t *testing.T) {
	tests := []struct {
		name   string
		format Format
		raw    string
	}{
		{name: "JSON null", format: FormatJSON, raw: `null`},
		{name: "JSON list", format: FormatJSON, raw: `[]`},
		{name: "YAML null", format: FormatYAMLMap, raw: `null`},
		{name: "YAML list", format: FormatYAMLList, raw: "- item\n"},
		{name: "TOML server array", format: FormatTOML, raw: "[[mcp_servers]]\ncommand = 'unsafe'\n"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if _, _, err := transform(test.format, []byte(test.raw), "/bin/apis-mcp", false, false); err == nil {
				t.Fatal("expected unsafe shape to be rejected")
			}
		})
	}
}

func TestAtomicPublishRejectsConcurrentSourceChange(t *testing.T) {
	directory := t.TempDir()
	path := filepath.Join(directory, "client.json")
	original := []byte(`{"owner":"first"}`)
	concurrent := []byte(`{"owner":"second"}`)
	if err := os.WriteFile(path, original, 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, concurrent, 0o600); err != nil {
		t.Fatal(err)
	}
	err := writeAtomic(path, []byte(`{"mcpServers":{}}`), 0o600, original, false)
	if !errors.Is(err, ErrConcurrentChange) {
		t.Fatalf("got %v, want ErrConcurrentChange", err)
	}
	raw, readErr := os.ReadFile(path)
	if readErr != nil || string(raw) != string(concurrent) {
		t.Fatalf("concurrent content was overwritten: %s (%v)", raw, readErr)
	}
	entries, err := os.ReadDir(directory)
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), ".apis-mcp-") {
			t.Fatalf("temporary file was not removed: %s", entry.Name())
		}
	}
}

func TestConcurrentClientUpdateRetriesAndMergesLatestSource(t *testing.T) {
	directory := t.TempDir()
	path := filepath.Join(directory, "opencode.json")
	if err := os.WriteFile(path, []byte(`{"theme":"dark"}`), 0o600); err != nil {
		t.Fatal(err)
	}
	staleBackup := filepath.Join(directory, "stale.backup")
	client := Client{ID: "opencode", Name: "OpenCode", Format: FormatOpenCode, ConfigPath: path}
	attempts := 0
	result, err := retryClientUpdate(func() (Result, error) {
		attempts++
		if attempts == 1 {
			if err := os.WriteFile(staleBackup, []byte("unused"), 0o600); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(path, []byte(`{"theme":"light","unrelated":true}`), 0o600); err != nil {
				t.Fatal(err)
			}
			return Result{Client: client, Backup: staleBackup}, ErrConcurrentChange
		}
		return updateClientOnce(client, "/bin/apis-mcp", false, Options{Backup: true})
	})
	if err != nil {
		t.Fatal(err)
	}
	if attempts != 2 || !result.Changed {
		t.Fatalf("retry result = attempts %d, %+v", attempts, result)
	}
	if _, err := os.Stat(staleBackup); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("unused retry backup was not removed: %v", err)
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(raw), `"theme": "light"`) || !strings.Contains(string(raw), `"unrelated": true`) || !strings.Contains(string(raw), `"apis-mcp"`) {
		t.Fatalf("latest OpenCode config was not merged: %s", raw)
	}
}

func TestExhaustedConcurrentRetriesAreNotReportedChanged(t *testing.T) {
	attempts := 0
	result, err := retryClientUpdate(func() (Result, error) {
		attempts++
		return Result{Changed: true}, ErrConcurrentChange
	})
	if !errors.Is(err, ErrConcurrentChange) || result.Changed || attempts != clientUpdateAttempts {
		t.Fatalf("exhausted retry = attempts %d result %+v err %v", attempts, result, err)
	}
}

func TestConfigureJSONPreservesServersAndIsIdempotent(t *testing.T) {
	home := t.TempDir()
	path := filepath.Join(home, ".cursor", "mcp.json")
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatal(err)
	}
	original := `{"theme":"dark","mcpServers":{"other":{"command":"other"}}}`
	if err := os.WriteFile(path, []byte(original), 0o640); err != nil {
		t.Fatal(err)
	}
	options := Options{Home: home, GOOS: "linux", Executable: "bin/apis-mcp", ClientIDs: []string{"cursor"}, Backup: true}
	results, err := Configure(options)
	if err != nil {
		t.Fatal(err)
	}
	if !results[0].Changed || results[0].Backup == "" {
		t.Fatalf("unexpected result: %+v", results[0])
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var root map[string]any
	if err := json.Unmarshal(raw, &root); err != nil {
		t.Fatal(err)
	}
	servers := root["mcpServers"].(map[string]any)
	if servers["other"] == nil || root["theme"] != "dark" {
		t.Fatalf("unrelated config lost: %s", raw)
	}
	entry := servers[serverName].(map[string]any)
	if !filepath.IsAbs(entry["command"].(string)) || entry["args"].([]any)[0] != "mcp" {
		t.Fatalf("wrong MCP entry: %#v", entry)
	}
	results, err = Configure(options)
	if err != nil || results[0].Changed {
		t.Fatalf("second configure was not idempotent: %+v, %v", results, err)
	}
}

func TestRepresentativeFormatWriters(t *testing.T) {
	tests := []struct {
		format Format
		raw    string
		check  func(*testing.T, []byte)
	}{
		{FormatVSCode, `{"servers":{"keep":{"command":"x"}}}`, func(t *testing.T, raw []byte) {
			var root map[string]any
			_ = json.Unmarshal(raw, &root)
			servers := root["servers"].(map[string]any)
			if servers["keep"] == nil || servers[serverName].(map[string]any)["type"] != "stdio" {
				t.Fatalf("bad VS Code JSON: %s", raw)
			}
		}},
		{FormatOpenCode, `{"mcp":{"keep":{"type":"remote"}}}`, func(t *testing.T, raw []byte) {
			if !strings.Contains(string(raw), `"command": [`) || !strings.Contains(string(raw), `"mcp"`) {
				t.Fatalf("bad OpenCode JSON: %s", raw)
			}
		}},
		{FormatTOML, "model = 'test'\n[mcp_servers.keep]\ncommand = 'x'\n", func(t *testing.T, raw []byte) {
			var root map[string]any
			if err := toml.Unmarshal(raw, &root); err != nil || root["model"] != "test" {
				t.Fatalf("bad TOML: %s (%v)", raw, err)
			}
			if object(root["mcp_servers"])["keep"] == nil || object(root["mcp_servers"])[serverName] == nil {
				t.Fatalf("TOML servers lost: %s", raw)
			}
		}},
		{FormatYAMLList, "name: config\nmcpServers:\n  - name: keep\n    command: x\n", func(t *testing.T, raw []byte) {
			var root map[string]any
			if err := yaml.Unmarshal(raw, &root); err != nil || root["name"] != "config" || len(slice(root["mcpServers"])) != 2 {
				t.Fatalf("bad YAML: %s (%v)", raw, err)
			}
		}},
	}
	for _, test := range tests {
		next, _, err := transform(test.format, []byte(test.raw), "/opt/apis-mcp", false, false)
		if err != nil {
			t.Errorf("%s: %v", test.format, err)
			continue
		}
		t.Run(string(test.format), func(t *testing.T) { test.check(t, next) })
	}
}

func TestFreshContinueConfigIncludesRequiredMetadata(t *testing.T) {
	next, _, err := transform(FormatYAMLList, nil, "/opt/apis-mcp", false, false)
	if err != nil {
		t.Fatal(err)
	}
	var root map[string]any
	if err := yaml.Unmarshal(next, &root); err != nil {
		t.Fatal(err)
	}
	if root["name"] != "Local Config" || root["version"] != "1.0.0" || root["schema"] != "v1" {
		t.Fatalf("fresh Continue metadata missing: %s", next)
	}
	if len(slice(root["mcpServers"])) != 1 {
		t.Fatalf("fresh Continue MCP entry missing: %s", next)
	}
}

func TestOpenCodeJSONCPreservesCommentsAndIsIdempotent(t *testing.T) {
	raw := []byte("{\n  // keep this comment\n  \"theme\": \"dark\",\n}\n")
	next, _, err := transform(FormatOpenCodeJSONC, raw, "/opt/apis-mcp", false, false)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(next), "// keep this comment") || !strings.Contains(string(next), `"theme": "dark"`) {
		t.Fatalf("JSONC content was not preserved:\n%s", next)
	}
	standard, err := hujson.Standardize(next)
	if err != nil {
		t.Fatal(err)
	}
	var root map[string]any
	if err := json.Unmarshal(standard, &root); err != nil {
		t.Fatal(err)
	}
	entry := object(object(root["mcp"])[serverName])
	if entry["type"] != "local" {
		t.Fatalf("OpenCode entry missing: %s", next)
	}
	again, _, err := transform(FormatOpenCodeJSONC, next, "/opt/apis-mcp", false, false)
	if err != nil || string(again) != string(next) {
		t.Fatalf("JSONC update is not idempotent:\n%s\n%s\n%v", next, again, err)
	}
}

func TestRegistryUsesExistingOpenCodeJSONC(t *testing.T) {
	home := t.TempDir()
	path := filepath.Join(home, ".config", "opencode", "opencode.jsonc")
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte("{}\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	for _, client := range Registry(home, "linux") {
		if client.ID == "opencode" {
			if client.ConfigPath != path || client.Format != FormatOpenCodeJSONC {
				t.Fatalf("OpenCode client = %+v", client)
			}
			return
		}
	}
	t.Fatal("OpenCode client missing")
}

func TestZedUsesCurrentJSONCShape(t *testing.T) {
	raw := []byte("{\n  // Zed settings comment\n  \"context_servers\": {\"keep\": {\"command\": \"other\"}},\n}\n")
	next, _, err := transform(FormatZed, raw, "/opt/apis-mcp", false, false)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(next), "// Zed settings comment") {
		t.Fatalf("Zed comment was not preserved:\n%s", next)
	}
	standard, err := hujson.Standardize(next)
	if err != nil {
		t.Fatal(err)
	}
	var root map[string]any
	if err := json.Unmarshal(standard, &root); err != nil {
		t.Fatal(err)
	}
	servers := object(root["context_servers"])
	entry := object(servers[serverName])
	if servers["keep"] == nil || entry["command"] != "/opt/apis-mcp" || slice(entry["args"])[0] != "mcp" {
		t.Fatalf("bad Zed entry: %s", next)
	}
	if _, nested := entry["command"].(map[string]any); nested {
		t.Fatalf("Zed command used obsolete nested shape: %s", next)
	}
}

func TestGooseEntryAndWindowsPaths(t *testing.T) {
	next, _, err := transform(FormatYAMLMap, nil, "/opt/apis-mcp", false, false)
	if err != nil {
		t.Fatal(err)
	}
	var root map[string]any
	if err := yaml.Unmarshal(next, &root); err != nil {
		t.Fatal(err)
	}
	entry := object(object(root["extensions"])[serverName])
	if entry["name"] != serverName || entry["type"] != "stdio" || entry["cmd"] != "/opt/apis-mcp" {
		t.Fatalf("bad Goose entry: %s", next)
	}

	home := filepath.Join("C:", "Users", "test")
	clients := Registry(home, "windows")
	paths := make(map[string]string)
	for _, client := range clients {
		paths[client.ID] = client.ConfigPath
	}
	if want := filepath.Join(home, "AppData", "Roaming", "Zed", "settings.json"); paths["zed"] != want {
		t.Fatalf("Zed Windows path = %q, want %q", paths["zed"], want)
	}
	if want := filepath.Join(home, "AppData", "Roaming", "Block", "goose", "config", "config.yaml"); paths["goose"] != want {
		t.Fatalf("Goose Windows path = %q, want %q", paths["goose"], want)
	}
}

func TestDryRunAndUninstall(t *testing.T) {
	home := t.TempDir()
	path := filepath.Join(home, ".gemini", "settings.json")
	options := Options{Home: home, GOOS: "linux", Executable: "/bin/apis-mcp", ClientIDs: []string{"gemini"}, DryRun: true}
	results, err := Configure(options)
	if err != nil || !results[0].Changed {
		t.Fatalf("dry-run failed: %+v %v", results, err)
	}
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Fatalf("dry-run created config: %v", err)
	}
	options.DryRun = false
	if _, err := Configure(options); err != nil {
		t.Fatal(err)
	}
	results, err = Uninstall(options)
	if err != nil || !results[0].Changed {
		t.Fatalf("uninstall failed: %+v %v", results, err)
	}
	raw, _ := os.ReadFile(path)
	if strings.Contains(string(raw), serverName) {
		t.Fatalf("server was not removed: %s", raw)
	}
}

func TestWrongServerContainerDoesNotOverwriteConfig(t *testing.T) {
	home := t.TempDir()
	path := filepath.Join(home, ".cursor", "mcp.json")
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatal(err)
	}
	original := []byte(`{"mcpServers":"managed elsewhere"}`)
	if err := os.WriteFile(path, original, 0o600); err != nil {
		t.Fatal(err)
	}
	_, err := Configure(Options{Home: home, GOOS: "linux", Executable: "/bin/apis-mcp", ClientIDs: []string{"cursor"}})
	if err == nil {
		t.Fatal("expected incompatible config to be rejected")
	}
	raw, readErr := os.ReadFile(path)
	if readErr != nil || string(raw) != string(original) {
		t.Fatalf("config changed after rejection: %s (%v)", raw, readErr)
	}
}
