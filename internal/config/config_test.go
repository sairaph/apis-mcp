package config

import (
	"os"
	"path/filepath"
	"testing"
)

func testPaths(t *testing.T) Paths {
	t.Helper()
	root := t.TempDir()
	return Paths{Root: root, Config: filepath.Join(root, "nested", "config.toml")}
}

func TestSaveAndLoadConfigAtomically(t *testing.T) {
	paths := testPaths(t)
	cfg := Default()
	cfg.ListTokenBudget = 3210
	if err := Save(paths, cfg); err != nil {
		t.Fatal(err)
	}
	loaded, err := Load(paths)
	if err != nil {
		t.Fatal(err)
	}
	if loaded.ListTokenBudget != 3210 || loaded.BackgroundAfter.Seconds() != float64(cfg.BackgroundAfterSeconds) {
		t.Fatalf("unexpected loaded config: %+v", loaded)
	}
	info, err := os.Stat(paths.Config)
	if err != nil {
		t.Fatal(err)
	}
	if info.Mode().Perm() != 0o600 {
		t.Fatalf("config mode is %o", info.Mode().Perm())
	}
	entries, err := os.ReadDir(filepath.Dir(paths.Config))
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		if entry.Name() != filepath.Base(paths.Config) {
			t.Fatalf("unexpected staging file left behind: %s", entry.Name())
		}
	}
}

func TestInvalidConfigDoesNotReplaceSavedConfig(t *testing.T) {
	paths := testPaths(t)
	valid := Default()
	if err := Save(paths, valid); err != nil {
		t.Fatal(err)
	}
	invalid := valid
	invalid.MaximumRedirects = 21
	if err := Validate(invalid); err == nil {
		t.Fatal("Validate accepted invalid redirects")
	}
	if err := Save(paths, invalid); err == nil {
		t.Fatal("Save accepted invalid config")
	}
	loaded, err := Load(paths)
	if err != nil {
		t.Fatal(err)
	}
	if loaded.MaximumRedirects != valid.MaximumRedirects {
		t.Fatalf("invalid save replaced config: %+v", loaded)
	}
}

func TestRestoreDefaults(t *testing.T) {
	paths := testPaths(t)
	changed := Default()
	changed.RetentionHours = 99
	if err := Save(paths, changed); err != nil {
		t.Fatal(err)
	}
	restored, err := RestoreDefaults(paths)
	if err != nil {
		t.Fatal(err)
	}
	if restored.RetentionHours != Default().RetentionHours {
		t.Fatalf("restore returned non-default config: %+v", restored)
	}
	loaded, err := Load(paths)
	if err != nil {
		t.Fatal(err)
	}
	if loaded.RetentionHours != Default().RetentionHours {
		t.Fatalf("restore did not persist defaults: %+v", loaded)
	}
}

func TestDefaultPathsIncludesPacks(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	paths, err := DefaultPaths()
	if err != nil {
		t.Fatal(err)
	}
	if paths.Packs != filepath.Join(home, ".apis-mcp", "packs") {
		t.Fatalf("packs path = %q", paths.Packs)
	}
}
