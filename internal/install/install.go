// Package install configures local MCP clients to start apis-mcp over stdio.
package install

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/gofrs/flock"
	"github.com/pelletier/go-toml/v2"
	"github.com/sairaph/apis-mcp/internal/fsx"
	"github.com/tailscale/hujson"
	"gopkg.in/yaml.v3"
)

const serverName = "apis-mcp"

// ErrConcurrentChange means a client config changed after it was read. The
// caller must retry so the other writer's update can be merged rather than
// overwritten.
var ErrConcurrentChange = errors.New("client configuration changed concurrently")

// Format identifies a client's configuration representation.
type Format string

const (
	FormatJSON          Format = "json"
	FormatVSCode        Format = "vscode-json"
	FormatOpenCode      Format = "opencode-json"
	FormatOpenCodeJSONC Format = "opencode-jsonc"
	FormatZed           Format = "zed-json"
	FormatTOML          Format = "toml"
	FormatYAMLList      Format = "yaml-list"
	FormatYAMLMap       Format = "yaml-map"
)

// Client is one supported MCP client and its per-user configuration file.
type Client struct {
	ID         string
	Name       string
	Format     Format
	ConfigPath string
}

// Status reports whether a client configuration exists and contains apis-mcp.
type Status struct {
	Client     Client
	Detected   bool
	Configured bool
	Err        error
}

// Options controls configuration updates. With no ClientIDs, only clients with
// an existing config file are changed. All creates config files for all clients.
type Options struct {
	Executable string
	Home       string
	GOOS       string
	ClientIDs  []string
	All        bool
	DryRun     bool
	Backup     bool
}

// Result describes one requested client update.
type Result struct {
	Client  Client
	Path    string
	Changed bool
	DryRun  bool
	Backup  string
	Removed bool
}

// SupportedClients returns the registry resolved for the current user.
func SupportedClients() []Client {
	home, _ := os.UserHomeDir()
	return Registry(home, runtime.GOOS)
}

// Registry returns the data-driven client registry for a home and platform.
func Registry(home, goos string) []Client {
	config := filepath.Join(home, ".config")
	claudeDesktop := filepath.Join(config, "Claude", "claude_desktop_config.json")
	vsCode := filepath.Join(config, "Code", "User", "mcp.json")
	zedConfig := filepath.Join(config, "zed", "settings.json")
	gooseConfig := filepath.Join(config, "goose", "config.yaml")
	if goos == "darwin" {
		claudeDesktop = filepath.Join(home, "Library", "Application Support", "Claude", "claude_desktop_config.json")
		vsCode = filepath.Join(home, "Library", "Application Support", "Code", "User", "mcp.json")
	} else if goos == "windows" {
		roaming := filepath.Join(home, "AppData", "Roaming")
		claudeDesktop = filepath.Join(roaming, "Claude", "claude_desktop_config.json")
		vsCode = filepath.Join(roaming, "Code", "User", "mcp.json")
		zedConfig = filepath.Join(roaming, "Zed", "settings.json")
		gooseConfig = filepath.Join(roaming, "Block", "goose", "config", "config.yaml")
	}
	codeUser := filepath.Dir(vsCode)
	openCodePath := filepath.Join(config, "opencode", "opencode.json")
	openCodeFormat := FormatOpenCode
	openCodeJSONC := filepath.Join(config, "opencode", "opencode.jsonc")
	if _, err := os.Stat(openCodePath); errors.Is(err, os.ErrNotExist) {
		if _, jsoncErr := os.Stat(openCodeJSONC); jsoncErr == nil {
			openCodePath, openCodeFormat = openCodeJSONC, FormatOpenCodeJSONC
		}
	}
	return []Client{
		{ID: "claude-desktop", Name: "Claude Desktop", Format: FormatJSON, ConfigPath: claudeDesktop},
		{ID: "claude-code", Name: "Claude Code", Format: FormatJSON, ConfigPath: filepath.Join(home, ".claude.json")},
		{ID: "cursor", Name: "Cursor", Format: FormatJSON, ConfigPath: filepath.Join(home, ".cursor", "mcp.json")},
		{ID: "windsurf", Name: "Windsurf", Format: FormatJSON, ConfigPath: filepath.Join(home, ".codeium", "windsurf", "mcp_config.json")},
		{ID: "vscode", Name: "Visual Studio Code", Format: FormatVSCode, ConfigPath: vsCode},
		{ID: "cline", Name: "Cline", Format: FormatJSON, ConfigPath: filepath.Join(codeUser, "globalStorage", "saoudrizwan.claude-dev", "settings", "cline_mcp_settings.json")},
		{ID: "roo-code", Name: "Roo Code", Format: FormatJSON, ConfigPath: filepath.Join(codeUser, "globalStorage", "rooveterinaryinc.roo-cline", "settings", "mcp_settings.json")},
		{ID: "zed", Name: "Zed", Format: FormatZed, ConfigPath: zedConfig},
		{ID: "gemini", Name: "Gemini CLI", Format: FormatJSON, ConfigPath: filepath.Join(home, ".gemini", "settings.json")},
		{ID: "codex", Name: "Codex CLI", Format: FormatTOML, ConfigPath: filepath.Join(home, ".codex", "config.toml")},
		{ID: "opencode", Name: "OpenCode", Format: openCodeFormat, ConfigPath: openCodePath},
		{ID: "continue", Name: "Continue", Format: FormatYAMLList, ConfigPath: filepath.Join(home, ".continue", "config.yaml")},
		{ID: "goose", Name: "Goose", Format: FormatYAMLMap, ConfigPath: gooseConfig},
	}
}

// Detect inspects all supported client configuration files.
func Detect(home, goos string) []Status {
	if home == "" {
		home, _ = os.UserHomeDir()
	}
	if goos == "" {
		goos = runtime.GOOS
	}
	clients := Registry(home, goos)
	statuses := make([]Status, 0, len(clients))
	for _, client := range clients {
		raw, err := os.ReadFile(client.ConfigPath)
		status := Status{Client: client}
		if errors.Is(err, os.ErrNotExist) {
			statuses = append(statuses, status)
			continue
		}
		if err != nil {
			status.Err = err
			statuses = append(statuses, status)
			continue
		}
		status.Detected = true
		_, configured, err := transform(client.Format, raw, "", false, true)
		status.Configured, status.Err = configured, err
		statuses = append(statuses, status)
	}
	return statuses
}

// Configure adds or replaces the apis-mcp entry in selected client configs.
func Configure(options Options) ([]Result, error) { return update(options, false) }

// Uninstall removes only the apis-mcp entry from selected client configs.
func Uninstall(options Options) ([]Result, error) { return update(options, true) }

func update(options Options, remove bool) ([]Result, error) {
	home := options.Home
	if home == "" {
		home, _ = os.UserHomeDir()
	}
	goos := options.GOOS
	if goos == "" {
		goos = runtime.GOOS
	}
	executable := options.Executable
	if !remove {
		if executable == "" {
			var err error
			executable, err = os.Executable()
			if err != nil {
				return nil, fmt.Errorf("resolve executable: %w", err)
			}
		}
		absolute, err := filepath.Abs(executable)
		if err != nil {
			return nil, fmt.Errorf("resolve executable: %w", err)
		}
		executable = absolute
	}
	clients, err := selectClients(Registry(home, goos), options)
	if err != nil {
		return nil, err
	}
	var results []Result
	var failures []error
	for _, client := range clients {
		result, err := updateClient(client, executable, remove, options)
		results = append(results, result)
		if err != nil {
			failures = append(failures, fmt.Errorf("%s: %w", client.Name, err))
		}
	}
	return results, errors.Join(failures...)
}

func selectClients(registry []Client, options Options) ([]Client, error) {
	byID := make(map[string]Client, len(registry))
	for _, client := range registry {
		byID[client.ID] = client
	}
	if len(options.ClientIDs) > 0 {
		seen := make(map[string]bool)
		var selected []Client
		for _, id := range options.ClientIDs {
			client, ok := byID[id]
			if !ok {
				return nil, fmt.Errorf("unsupported client %q", id)
			}
			if !seen[id] {
				selected, seen[id] = append(selected, client), true
			}
		}
		return selected, nil
	}
	if options.All {
		return registry, nil
	}
	var selected []Client
	for _, client := range registry {
		if _, err := os.Stat(client.ConfigPath); err == nil {
			selected = append(selected, client)
		}
	}
	return selected, nil
}

func updateClient(client Client, executable string, remove bool, options Options) (Result, error) {
	result := Result{Client: client, Path: client.ConfigPath, DryRun: options.DryRun, Removed: remove}
	raw, err := os.ReadFile(client.ConfigPath)
	missing := errors.Is(err, os.ErrNotExist)
	if err != nil && !missing {
		return result, err
	}
	if missing {
		if remove {
			return result, nil
		}
		raw = nil
	}
	next, present, err := transform(client.Format, raw, executable, remove, false)
	if err != nil {
		return result, err
	}
	if remove && !present || bytes.Equal(raw, next) {
		return result, nil
	}
	result.Changed = true
	if options.DryRun {
		return result, nil
	}
	if err := os.MkdirAll(filepath.Dir(client.ConfigPath), 0o700); err != nil {
		return result, err
	}
	lock := flock.New(filepath.Join(filepath.Dir(client.ConfigPath), "."+filepath.Base(client.ConfigPath)+".apis-mcp.lock"))
	if err := lock.Lock(); err != nil {
		return result, fmt.Errorf("lock config: %w", err)
	}
	defer lock.Close()
	if matches, err := sourceMatches(client.ConfigPath, raw, missing); err != nil {
		return result, err
	} else if !matches {
		return result, ErrConcurrentChange
	}
	if !missing && options.Backup {
		result.Backup, err = backupFile(client.ConfigPath, raw)
		if err != nil {
			return result, fmt.Errorf("create backup: %w", err)
		}
	}
	mode := os.FileMode(0o600)
	if info, err := os.Stat(client.ConfigPath); err == nil {
		mode = info.Mode().Perm()
	}
	if err := writeAtomic(client.ConfigPath, next, mode, raw, missing); err != nil {
		return result, err
	}
	return result, nil
}

func transform(format Format, raw []byte, executable string, remove, inspect bool) ([]byte, bool, error) {
	switch format {
	case FormatJSON, FormatVSCode, FormatOpenCode:
		return transformJSON(format, raw, executable, remove, inspect)
	case FormatZed:
		return transformZedJSONC(raw, executable, remove, inspect)
	case FormatOpenCodeJSONC:
		return transformOpenCodeJSONC(raw, executable, remove, inspect)
	case FormatTOML:
		return transformTOML(raw, executable, remove, inspect)
	case FormatYAMLList, FormatYAMLMap:
		return transformYAML(format, raw, executable, remove, inspect)
	default:
		return nil, false, fmt.Errorf("unsupported config format %q", format)
	}
}

func transformOpenCodeJSONC(raw []byte, executable string, remove, inspect bool) ([]byte, bool, error) {
	desired := map[string]any{"type": "local", "command": []string{executable, "mcp"}, "enabled": true}
	return transformJSONCEntry(raw, "mcp", desired, remove, inspect)
}

func transformZedJSONC(raw []byte, executable string, remove, inspect bool) ([]byte, bool, error) {
	desired := map[string]any{"command": executable, "args": []string{"mcp"}, "env": map[string]string{}}
	return transformJSONCEntry(raw, "context_servers", desired, remove, inspect)
}

func transformJSONCEntry(raw []byte, key string, desired map[string]any, remove, inspect bool) ([]byte, bool, error) {
	if len(bytes.TrimSpace(raw)) == 0 {
		raw = []byte("{}\n")
	}
	value, err := hujson.Parse(raw)
	if err != nil {
		return nil, false, fmt.Errorf("parse JSONC: %w", err)
	}
	standard := value.Clone()
	standard.Standardize()
	root, err := decodeJSONObject(standard.Pack())
	if err != nil {
		return nil, false, err
	}
	servers, err := mapField(root, key)
	if err != nil {
		return nil, false, err
	}
	_, present := servers[serverName]
	if inspect {
		return raw, present, nil
	}
	operations := make([]map[string]any, 0, 2)
	if remove {
		if !present {
			return raw, false, nil
		}
		operations = append(operations, map[string]any{"op": "remove", "path": "/" + key + "/" + serverName})
	} else {
		if present && sameJSONValue(servers[serverName], desired) {
			return raw, true, nil
		}
		if _, exists := root[key]; !exists {
			operations = append(operations, map[string]any{"op": "add", "path": "/" + key, "value": map[string]any{}})
		}
		op := "add"
		if present {
			op = "replace"
		}
		operations = append(operations, map[string]any{
			"op": op, "path": "/" + key + "/" + serverName,
			"value": desired,
		})
	}
	patch, err := json.Marshal(operations)
	if err != nil {
		return nil, present, err
	}
	if err := value.Patch(patch); err != nil {
		return nil, present, fmt.Errorf("update JSONC: %w", err)
	}
	value.Format()
	return value.Pack(), present, nil
}

func sameJSONValue(left, right any) bool {
	leftJSON, leftErr := json.Marshal(left)
	rightJSON, rightErr := json.Marshal(right)
	return leftErr == nil && rightErr == nil && bytes.Equal(leftJSON, rightJSON)
}

func decodeJSONObject(raw []byte) (map[string]any, error) {
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.UseNumber()
	var decoded any
	if err := decoder.Decode(&decoded); err != nil {
		return nil, fmt.Errorf("parse JSON: %w", err)
	}
	root, ok := decoded.(map[string]any)
	if !ok || root == nil {
		return nil, errors.New("JSON config must be a top-level object")
	}
	return root, nil
}

func transformJSON(format Format, raw []byte, executable string, remove, inspect bool) ([]byte, bool, error) {
	root := map[string]any{}
	if len(bytes.TrimSpace(raw)) > 0 {
		decoder := json.NewDecoder(bytes.NewReader(raw))
		decoder.UseNumber()
		var decoded any
		if err := decoder.Decode(&decoded); err != nil {
			return nil, false, fmt.Errorf("parse JSON: %w", err)
		}
		if err := requireEOF(decoder.Decode); err != nil {
			return nil, false, fmt.Errorf("parse JSON: %w", err)
		}
		var ok bool
		root, ok = decoded.(map[string]any)
		if !ok || root == nil {
			return nil, false, errors.New("JSON config must be a top-level object")
		}
	}
	key := "mcpServers"
	if format == FormatVSCode {
		key = "servers"
	} else if format == FormatOpenCode {
		key = "mcp"
	}
	servers, err := mapField(root, key)
	if err != nil {
		return nil, false, err
	}
	_, present := servers[serverName]
	if inspect {
		return raw, present, nil
	}
	if remove {
		delete(servers, serverName)
	} else {
		switch format {
		case FormatVSCode:
			servers[serverName] = map[string]any{"type": "stdio", "command": executable, "args": []string{"mcp"}}
		case FormatOpenCode:
			servers[serverName] = map[string]any{"type": "local", "command": []string{executable, "mcp"}, "enabled": true}
		default:
			servers[serverName] = map[string]any{"command": executable, "args": []string{"mcp"}}
		}
	}
	if len(servers) == 0 && remove {
		delete(root, key)
	} else {
		root[key] = servers
	}
	next, err := json.MarshalIndent(root, "", "  ")
	if err != nil {
		return nil, present, err
	}
	return append(next, '\n'), present, nil
}

func transformTOML(raw []byte, executable string, remove, inspect bool) ([]byte, bool, error) {
	root := map[string]any{}
	if len(bytes.TrimSpace(raw)) > 0 {
		var decoded any
		if err := toml.Unmarshal(raw, &decoded); err != nil {
			return nil, false, fmt.Errorf("parse TOML: %w", err)
		}
		var ok bool
		root, ok = decoded.(map[string]any)
		if !ok || root == nil {
			return nil, false, errors.New("TOML config must be a top-level object")
		}
	}
	servers, err := mapField(root, "mcp_servers")
	if err != nil {
		return nil, false, err
	}
	_, present := servers[serverName]
	if inspect {
		return raw, present, nil
	}
	if remove {
		delete(servers, serverName)
	} else {
		servers[serverName] = map[string]any{"command": executable, "args": []string{"mcp"}}
	}
	if len(servers) == 0 && remove {
		delete(root, "mcp_servers")
	} else {
		root["mcp_servers"] = servers
	}
	next, err := toml.Marshal(root)
	return next, present, err
}

func transformYAML(format Format, raw []byte, executable string, remove, inspect bool) ([]byte, bool, error) {
	root := map[string]any{}
	fresh := len(bytes.TrimSpace(raw)) == 0
	if len(bytes.TrimSpace(raw)) > 0 {
		decoder := yaml.NewDecoder(bytes.NewReader(raw))
		var decoded any
		if err := decoder.Decode(&decoded); err != nil {
			return nil, false, fmt.Errorf("parse YAML: %w", err)
		}
		if err := requireEOF(decoder.Decode); err != nil {
			return nil, false, fmt.Errorf("parse YAML: %w", err)
		}
		var ok bool
		root, ok = decoded.(map[string]any)
		if !ok || root == nil {
			return nil, false, errors.New("YAML config must be a top-level object")
		}
	}
	present := false
	if format == FormatYAMLList {
		if fresh && !remove && !inspect {
			root["name"] = "Local Config"
			root["version"] = "1.0.0"
			root["schema"] = "v1"
		}
		items, err := listField(root, "mcpServers")
		if err != nil {
			return nil, false, err
		}
		kept := make([]any, 0, len(items)+1)
		for _, item := range items {
			entry := object(item)
			if entry["name"] == serverName {
				present = true
				continue
			}
			kept = append(kept, item)
		}
		if inspect {
			return raw, present, nil
		}
		if !remove {
			kept = append(kept, map[string]any{"name": serverName, "command": executable, "args": []string{"mcp"}})
		}
		if len(kept) == 0 && remove {
			delete(root, "mcpServers")
		} else {
			root["mcpServers"] = kept
		}
	} else {
		extensions, err := mapField(root, "extensions")
		if err != nil {
			return nil, false, err
		}
		_, present = extensions[serverName]
		if inspect {
			return raw, present, nil
		}
		if remove {
			delete(extensions, serverName)
		} else {
			extensions[serverName] = map[string]any{
				"name": serverName, "type": "stdio", "cmd": executable,
				"args": []string{"mcp"}, "enabled": true,
			}
		}
		if len(extensions) == 0 && remove {
			delete(root, "extensions")
		} else {
			root["extensions"] = extensions
		}
	}
	next, err := yaml.Marshal(root)
	return next, present, err
}

func object(value any) map[string]any {
	if value == nil {
		return map[string]any{}
	}
	if result, ok := value.(map[string]any); ok {
		return result
	}
	return map[string]any{}
}

func slice(value any) []any {
	if result, ok := value.([]any); ok {
		return result
	}
	return nil
}

func mapField(root map[string]any, key string) (map[string]any, error) {
	value, exists := root[key]
	if !exists || value == nil {
		return map[string]any{}, nil
	}
	result, ok := value.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("%s must be an object", key)
	}
	return result, nil
}

func listField(root map[string]any, key string) ([]any, error) {
	value, exists := root[key]
	if !exists || value == nil {
		return nil, nil
	}
	result, ok := value.([]any)
	if !ok {
		return nil, fmt.Errorf("%s must be a list", key)
	}
	return result, nil
}

func requireEOF(decode func(any) error) error {
	var extra any
	if err := decode(&extra); !errors.Is(err, io.EOF) {
		if err == nil {
			return errors.New("multiple documents are not supported")
		}
		return err
	}
	return nil
}

func writeAtomic(path string, raw []byte, mode os.FileMode, source []byte, sourceMissing bool) error {
	temporary, err := os.CreateTemp(filepath.Dir(path), ".apis-mcp-*.tmp")
	if err != nil {
		return err
	}
	name := temporary.Name()
	defer os.Remove(name)
	if err := temporary.Chmod(mode); err != nil {
		temporary.Close()
		return err
	}
	if _, err := temporary.Write(raw); err != nil {
		temporary.Close()
		return err
	}
	if err := temporary.Sync(); err != nil {
		temporary.Close()
		return err
	}
	if err := temporary.Close(); err != nil {
		return err
	}
	if matches, err := sourceMatches(path, source, sourceMissing); err != nil {
		return err
	} else if !matches {
		return ErrConcurrentChange
	}
	if err := fsx.Replace(name, path); err != nil {
		return fmt.Errorf("publish config: %w", err)
	}
	if directory, err := os.Open(filepath.Dir(path)); err == nil {
		_ = directory.Sync()
		_ = directory.Close()
	}
	return nil
}

func sourceMatches(path string, expected []byte, expectedMissing bool) (bool, error) {
	raw, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return expectedMissing, nil
	}
	if err != nil {
		return false, fmt.Errorf("recheck config: %w", err)
	}
	return !expectedMissing && bytes.Equal(raw, expected), nil
}

func backupFile(source string, raw []byte) (string, error) {
	backup, err := os.CreateTemp(filepath.Dir(source), filepath.Base(source)+".bak-"+time.Now().UTC().Format("20060102T150405Z")+"-*")
	if err != nil {
		return "", err
	}
	name := backup.Name()
	if err := backup.Chmod(0o600); err != nil {
		backup.Close()
		os.Remove(name)
		return "", err
	}
	if _, err := backup.Write(raw); err != nil {
		backup.Close()
		os.Remove(name)
		return "", err
	}
	if err := backup.Sync(); err != nil {
		backup.Close()
		os.Remove(name)
		return "", err
	}
	if err := backup.Close(); err != nil {
		os.Remove(name)
		return "", err
	}
	return name, nil
}

// ClientIDs returns stable supported IDs for help and validation.
func ClientIDs() []string {
	clients := SupportedClients()
	ids := make([]string, 0, len(clients))
	for _, client := range clients {
		ids = append(ids, client.ID)
	}
	sort.Strings(ids)
	return ids
}

// ClientNames renders the stable ID/name registry for human-facing help.
func ClientNames() string {
	clients := SupportedClients()
	parts := make([]string, 0, len(clients))
	for _, client := range clients {
		parts = append(parts, client.ID+" ("+client.Name+")")
	}
	return strings.Join(parts, ", ")
}
