package cli

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/sairaph/apis-mcp/internal/config"
)

type configSetting struct {
	key      string
	label    string
	boolean  bool
	readOnly bool
}

var configSettings = []configSetting{
	{key: "version", label: "Config version", readOnly: true},
	{key: "list_token_budget", label: "List token budget"},
	{key: "read_token_budget", label: "Read token budget"},
	{key: "response_size_limit", label: "Response size limit (bytes)"},
	{key: "allow_large_download", label: "Allow large downloads", boolean: true},
	{key: "free_disk_reserve", label: "Free disk reserve (bytes)"},
	{key: "maximum_redirects", label: "Maximum redirects"},
	{key: "maximum_retries", label: "Maximum retries"},
	{key: "maximum_header_timeout_seconds", label: "Header timeout (seconds)"},
	{key: "background_after_seconds", label: "Background after (seconds)"},
	{key: "stalled_download_seconds", label: "Stalled download (seconds)"},
	{key: "retention_hours", label: "Retention (hours)"},
	{key: "tls_verify", label: "Verify TLS certificates", boolean: true},
}

// RunConfigure opens the shared application directly in its Settings context.
func RunConfigure(ctx context.Context, options Options) error {
	options = normalizeOptions(options)
	paths, err := config.DefaultPaths()
	if err != nil {
		return err
	}
	cfg, err := config.Load(paths)
	if err != nil {
		return err
	}
	root := newRootModel(ctx, nil, options, contextSettings)
	root.paths, root.settings.cfg, root.configureOnly = paths, cfg, true
	_, err = newTeaProgram(root, options).Run()
	root.close()
	return configureRunResult(ctx, root, err)
}

func configureRunResult(ctx context.Context, root *model, programErr error) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if programErr != nil {
		return programErr
	}
	if root.runErr != nil {
		return root.runErr
	}
	if !root.configureSaved || root.settings.dirty {
		return errors.New("configuration cancelled before a successful save")
	}
	return nil
}

func configValue(cfg config.Config, key string) string {
	switch key {
	case "version":
		return strconv.Itoa(cfg.Version)
	case "list_token_budget":
		return strconv.Itoa(cfg.ListTokenBudget)
	case "read_token_budget":
		return strconv.Itoa(cfg.ReadTokenBudget)
	case "response_size_limit":
		return strconv.FormatInt(cfg.ResponseSizeLimit, 10)
	case "allow_large_download":
		return strconv.FormatBool(cfg.AllowLargeDownload)
	case "free_disk_reserve":
		return strconv.FormatInt(cfg.FreeDiskReserve, 10)
	case "maximum_redirects":
		return strconv.Itoa(cfg.MaximumRedirects)
	case "maximum_retries":
		return strconv.Itoa(cfg.MaximumRetries)
	case "maximum_header_timeout_seconds":
		return strconv.Itoa(cfg.MaximumHeaderTimeout)
	case "background_after_seconds":
		return strconv.Itoa(cfg.BackgroundAfterSeconds)
	case "stalled_download_seconds":
		return strconv.Itoa(cfg.StalledDownloadSeconds)
	case "retention_hours":
		return strconv.Itoa(cfg.RetentionHours)
	case "tls_verify":
		return strconv.FormatBool(cfg.TLSVerify)
	default:
		return ""
	}
}

func configBoolValue(cfg config.Config, key string) bool {
	return key == "allow_large_download" && cfg.AllowLargeDownload || key == "tls_verify" && cfg.TLSVerify
}

func setConfigValue(cfg *config.Config, key, raw string) error {
	next := *cfg
	if key == "allow_large_download" || key == "tls_verify" {
		value, err := strconv.ParseBool(strings.TrimSpace(raw))
		if err != nil {
			return err
		}
		if key == "allow_large_download" {
			next.AllowLargeDownload = value
		} else {
			next.TLSVerify = value
		}
	} else if key == "response_size_limit" || key == "free_disk_reserve" {
		value, err := strconv.ParseInt(strings.TrimSpace(raw), 10, 64)
		if err != nil {
			return err
		}
		if key == "response_size_limit" {
			next.ResponseSizeLimit = value
		} else {
			next.FreeDiskReserve = value
		}
	} else {
		value, err := strconv.Atoi(strings.TrimSpace(raw))
		if err != nil {
			return err
		}
		switch key {
		case "list_token_budget":
			next.ListTokenBudget = value
		case "read_token_budget":
			next.ReadTokenBudget = value
		case "maximum_redirects":
			next.MaximumRedirects = value
		case "maximum_retries":
			next.MaximumRetries = value
		case "maximum_header_timeout_seconds":
			next.MaximumHeaderTimeout = value
		case "background_after_seconds":
			next.BackgroundAfterSeconds = value
		case "stalled_download_seconds":
			next.StalledDownloadSeconds = value
		case "retention_hours":
			next.RetentionHours = value
		default:
			return errors.New("setting is read-only")
		}
	}
	if err := config.Validate(next); err != nil {
		return err
	}
	*cfg = next
	return nil
}
