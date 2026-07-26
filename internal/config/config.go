package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/pelletier/go-toml/v2"
	"github.com/sairaph/apis-mcp/internal/fsx"
)

const currentVersion = 1

type Config struct {
	Version int `toml:"version"`

	ListTokenBudget int `toml:"list_token_budget"`
	ReadTokenBudget int `toml:"read_token_budget"`

	ResponseSizeLimit      int64         `toml:"response_size_limit"`
	AllowLargeDownload     bool          `toml:"allow_large_download"`
	FreeDiskReserve        int64         `toml:"free_disk_reserve"`
	MaximumRedirects       int           `toml:"maximum_redirects"`
	MaximumRetries         int           `toml:"maximum_retries"`
	MaximumHeaderTimeout   int           `toml:"maximum_header_timeout_seconds"`
	BackgroundAfter        time.Duration `toml:"-"`
	BackgroundAfterSeconds int           `toml:"background_after_seconds"`
	StalledDownloadAfter   time.Duration `toml:"-"`
	StalledDownloadSeconds int           `toml:"stalled_download_seconds"`
	Retention              time.Duration `toml:"-"`
	RetentionHours         int           `toml:"retention_hours"`
	TLSVerify              bool          `toml:"tls_verify"`
}

type Paths struct {
	Root        string
	Config      string
	Library     string
	Index       string
	Cache       string
	Sessions    string
	Diagnostics string
}

func Default() Config {
	return Config{
		Version:                currentVersion,
		ListTokenBudget:        2_000,
		ReadTokenBudget:        4_000,
		ResponseSizeLimit:      50 << 20,
		AllowLargeDownload:     true,
		FreeDiskReserve:        20 << 30,
		MaximumRedirects:       5,
		MaximumRetries:         30,
		MaximumHeaderTimeout:   600,
		BackgroundAfterSeconds: 60,
		StalledDownloadSeconds: 3600,
		RetentionHours:         24,
		TLSVerify:              true,
	}
}

func DefaultPaths() (Paths, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Paths{}, fmt.Errorf("resolve home directory: %w", err)
	}
	root := filepath.Join(home, ".apis-mcp")
	return Paths{
		Root:        root,
		Config:      filepath.Join(root, "config.toml"),
		Library:     filepath.Join(root, "library"),
		Index:       filepath.Join(root, "index"),
		Cache:       filepath.Join(root, "cache", "calls"),
		Sessions:    filepath.Join(root, "sessions"),
		Diagnostics: filepath.Join(root, "diagnostics.log"),
	}, nil
}

func Load(paths Paths) (Config, error) {
	cfg := Default()
	raw, err := os.ReadFile(paths.Config)
	if errors.Is(err, os.ErrNotExist) {
		if err := Save(paths, cfg); err != nil {
			return Config{}, err
		}
		return normalize(cfg)
	}
	if err != nil {
		return Config{}, fmt.Errorf("read config: %w", err)
	}
	if err := toml.Unmarshal(raw, &cfg); err != nil {
		return Config{}, fmt.Errorf("parse config: %w", err)
	}
	return normalize(cfg)
}

func Save(paths Paths, cfg Config) error {
	if err := Validate(cfg); err != nil {
		return err
	}
	directory := filepath.Dir(paths.Config)
	if err := os.MkdirAll(directory, 0o700); err != nil {
		return fmt.Errorf("create application directory: %w", err)
	}
	raw, err := toml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("encode config: %w", err)
	}
	tmp, err := os.CreateTemp(directory, ".config-*.tmp")
	if err != nil {
		return fmt.Errorf("create temporary config: %w", err)
	}
	tmpName := tmp.Name()
	defer os.Remove(tmpName)
	if err := tmp.Chmod(0o600); err != nil {
		tmp.Close()
		return err
	}
	if _, err := tmp.Write(raw); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Sync(); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	if err := fsx.Replace(tmpName, paths.Config); err != nil {
		return fmt.Errorf("publish config: %w", err)
	}
	if dir, err := os.Open(directory); err == nil {
		_ = dir.Sync()
		_ = dir.Close()
	}
	return nil
}

// Validate checks all persisted settings without writing them.
func Validate(cfg Config) error {
	_, err := normalize(cfg)
	return err
}

// RestoreDefaults atomically replaces the active configuration with defaults.
// Confirmation is intentionally a presentation-layer responsibility.
func RestoreDefaults(paths Paths) (Config, error) {
	cfg := Default()
	if err := Save(paths, cfg); err != nil {
		return Config{}, err
	}
	return normalize(cfg)
}

func normalize(cfg Config) (Config, error) {
	if cfg.Version != currentVersion {
		return Config{}, fmt.Errorf("unsupported config version %d", cfg.Version)
	}
	if cfg.ListTokenBudget <= 0 || cfg.ReadTokenBudget <= 0 {
		return Config{}, errors.New("token budgets must be positive")
	}
	if cfg.ResponseSizeLimit <= 0 || cfg.FreeDiskReserve < 0 {
		return Config{}, errors.New("response size must be positive and disk reserve cannot be negative")
	}
	if cfg.MaximumRedirects < 0 || cfg.MaximumRedirects > 20 {
		return Config{}, errors.New("maximum_redirects must be between 0 and 20")
	}
	if cfg.MaximumRetries < 0 || cfg.MaximumRetries > 30 {
		return Config{}, errors.New("maximum_retries must be between 0 and 30")
	}
	if cfg.MaximumHeaderTimeout < 30 || cfg.MaximumHeaderTimeout > 600 {
		return Config{}, errors.New("maximum_header_timeout_seconds must be between 30 and 600")
	}
	if cfg.BackgroundAfterSeconds <= 0 || cfg.StalledDownloadSeconds <= 0 || cfg.RetentionHours <= 0 {
		return Config{}, errors.New("background, stalled download, and retention values must be positive")
	}
	cfg.BackgroundAfter = time.Duration(cfg.BackgroundAfterSeconds) * time.Second
	cfg.StalledDownloadAfter = time.Duration(cfg.StalledDownloadSeconds) * time.Second
	cfg.Retention = time.Duration(cfg.RetentionHours) * time.Hour
	return cfg, nil
}
