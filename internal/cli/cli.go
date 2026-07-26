// Package cli provides the human-facing command line and terminal interfaces.
package cli

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/sairaph/apis-mcp/internal/app"
	"github.com/sairaph/apis-mcp/internal/bootstrap"
	"github.com/sairaph/apis-mcp/internal/importer"
	"github.com/sairaph/apis-mcp/internal/install"
	"github.com/sairaph/apis-mcp/library"
	"golang.org/x/term"
)

// Options supplies process-level values without coupling the CLI to main.
type Options struct {
	Version    string
	Executable string
	Stdin      io.Reader
	Stdout     io.Writer
	Stderr     io.Writer
}

// Command is the parsed, transport-neutral representation of one invocation.
type Command struct {
	Name       string
	Action     string
	Help       bool
	Page       int
	NameFilter string
	Version    string
	Collection string
	DocID      string
	Path       string
	Query      string
	PageID     string
	Lines      []int

	Method             string
	Endpoint           string
	Headers            any
	Payload            any
	Timeout            int
	Retries            *int
	JSONPath           string
	Session            string
	AllowLargeDownload bool

	ID        string
	ClientIDs []string
	All       bool
	DryRun    bool
	Backup    bool
	Explicit  bool

	ImportKind string
	APIName    string
	Source     string
	MaxPages   int
	MaxDepth   int
}

// Parse parses one-shot CLI arguments. An empty argument list selects the TUI.
func Parse(args []string) (Command, error) {
	if len(args) == 0 {
		return Command{Name: "interactive"}, nil
	}
	if args[0] == "-h" || args[0] == "--help" || args[0] == "help" {
		return Command{Name: "help"}, nil
	}
	if args[0] == "-v" || args[0] == "--version" || args[0] == "version" {
		return Command{Name: "version"}, nil
	}
	command := Command{Name: args[0], Page: 1, Backup: true}
	rest := args[1:]
	switch command.Name {
	case "rebuild", "doctor":
		if len(rest) != 0 {
			return Command{}, unexpected(command.Name, rest)
		}
	case "collections":
		fs := flags(command.Name)
		fs.IntVar(&command.Page, "page", 1, "result page")
		if err := parseFlags(fs, rest, nil); err != nil {
			return Command{}, err
		}
	case "list":
		fs := flags(command.Name)
		fs.StringVar(&command.NameFilter, "name", "", "API name filter")
		fs.StringVar(&command.Version, "version", "", "version filter")
		fs.StringVar(&command.Collection, "collection", "", "collection filter")
		fs.IntVar(&command.Page, "page", 1, "result page")
		if err := parseFlags(fs, rest, nil); err != nil {
			return Command{}, err
		}
	case "pages":
		fs := flags(command.Name)
		fs.StringVar(&command.Path, "path", "", "documentation path")
		fs.IntVar(&command.Page, "page", 1, "result page")
		positionals, err := parsePositionals(fs, rest, nil)
		if err != nil {
			return Command{}, err
		}
		if len(positionals) != 1 {
			return Command{}, errors.New("usage: apis-mcp pages <doc-id> [--path PATH] [--page N]")
		}
		command.DocID = positionals[0]
	case "search":
		fs := flags(command.Name)
		fs.StringVar(&command.Path, "path", "", "documentation path")
		fs.IntVar(&command.Page, "page", 1, "result page")
		positionals, err := parsePositionals(fs, rest, nil)
		if err != nil {
			return Command{}, err
		}
		if len(positionals) < 2 {
			return Command{}, errors.New("usage: apis-mcp search <doc-id> <query> [--path PATH] [--page N]")
		}
		command.DocID, command.Query = positionals[0], strings.Join(positionals[1:], " ")
	case "read":
		var lineRange string
		fs := flags(command.Name)
		fs.StringVar(&lineRange, "lines", "", "inclusive START:END range")
		positionals, err := parsePositionals(fs, rest, nil)
		if err != nil {
			return Command{}, err
		}
		if len(positionals) != 2 {
			return Command{}, errors.New("usage: apis-mcp read <doc-id> <page-id> [--lines START:END]")
		}
		command.DocID, command.PageID = positionals[0], positionals[1]
		if lineRange != "" {
			command.Lines, err = parseLineRange(lineRange)
			if err != nil {
				return Command{}, err
			}
		}
	case "call":
		var headers, payload string
		var retries optionalInt
		fs := flags(command.Name)
		fs.StringVar(&headers, "headers", "", "JSON object or file path")
		fs.StringVar(&payload, "payload", "", "JSON value or file path")
		fs.IntVar(&command.Timeout, "timeout", 0, "header timeout in seconds")
		fs.Var(&retries, "retries", "retry count")
		fs.StringVar(&command.JSONPath, "json-path", "", "JSONPath preview selector")
		fs.StringVar(&command.Session, "session", "", "cookie session ID")
		fs.BoolVar(&command.AllowLargeDownload, "allow-large-download", false, "bypass response size cap")
		positionals, err := parsePositionals(fs, rest, map[string]bool{"allow-large-download": true})
		if err != nil {
			return Command{}, err
		}
		if len(positionals) != 2 {
			return Command{}, errors.New("usage: apis-mcp call <method> <url> [options]")
		}
		command.Method, command.Endpoint = positionals[0], positionals[1]
		if retries.set {
			value := retries.value
			command.Retries = &value
		}
		if command.Headers, err = jsonOrPath(headers); err != nil {
			return Command{}, fmt.Errorf("headers: %w", err)
		}
		if command.Payload, err = jsonOrPath(payload); err != nil {
			return Command{}, fmt.Errorf("payload: %w", err)
		}
	case "sessions":
		if len(rest) == 0 {
			command.Action = "list"
			break
		}
		command.Action = rest[0]
		switch command.Action {
		case "list", "create", "cleanup":
			if len(rest) != 1 {
				return Command{}, unexpected(command.Name, rest[1:])
			}
		case "show", "delete":
			if len(rest) != 2 {
				return Command{}, fmt.Errorf("usage: apis-mcp sessions %s <id>", command.Action)
			}
			command.ID = rest[1]
		default:
			return Command{}, fmt.Errorf("unknown sessions action %q", command.Action)
		}
	case "cache":
		if len(rest) != 1 || rest[0] != "cleanup" {
			return Command{}, errors.New("usage: apis-mcp cache cleanup")
		}
		command.Action = "cleanup"
	case "config":
		if len(rest) > 1 || len(rest) == 1 && rest[0] != "path" {
			return Command{}, errors.New("usage: apis-mcp config [path]")
		}
		command.Action = "path"
	case "install", "configure", "uninstall":
		command.Explicit = len(rest) > 0
		var clients stringList
		fs := flags(command.Name)
		fs.Var(&clients, "client", "client ID (repeatable)")
		fs.BoolVar(&command.All, "all", false, "configure every supported client")
		fs.BoolVar(&command.DryRun, "dry-run", false, "report changes without writing")
		fs.BoolVar(&command.Backup, "backup", true, "back up existing changed files")
		if err := parseFlags(fs, rest, map[string]bool{"all": true, "dry-run": true, "backup": true}); err != nil {
			return Command{}, err
		}
		command.ClientIDs = clients
	case "import":
		if len(rest) == 0 {
			return Command{}, errors.New("usage: apis-mcp import <markdown|openapi|llms|html> ...")
		}
		command.ImportKind = rest[0]
		switch command.ImportKind {
		case "markdown":
			if len(rest) != 2 {
				return Command{}, errors.New("usage: apis-mcp import markdown <directory>")
			}
			command.Source = rest[1]
		case "openapi", "llms":
			if len(rest) != 4 {
				return Command{}, fmt.Errorf("usage: apis-mcp import %s <api-name> <version> <file-or-url>", command.ImportKind)
			}
			command.APIName, command.Version, command.Source = rest[1], rest[2], rest[3]
		case "html":
			command.MaxPages, command.MaxDepth = importer.DefaultMaxHTMLPages, importer.DefaultMaxHTMLDepth
			fs := flags("import html")
			fs.IntVar(&command.MaxPages, "max-pages", importer.DefaultMaxHTMLPages, "maximum HTML pages")
			fs.IntVar(&command.MaxDepth, "max-depth", importer.DefaultMaxHTMLDepth, "maximum link depth")
			positionals, err := parsePositionals(fs, rest[1:], nil)
			if err != nil {
				return Command{}, err
			}
			if len(positionals) != 3 {
				return Command{}, errors.New("usage: apis-mcp import html <api-name> <version> <url> [--max-pages N] [--max-depth N]")
			}
			if command.MaxPages < 1 || command.MaxDepth < 1 {
				return Command{}, errors.New("HTML import limits must be positive")
			}
			command.APIName, command.Version, command.Source = positionals[0], positionals[1], positionals[2]
		default:
			return Command{}, fmt.Errorf("unknown import type %q; expected markdown, openapi, llms, or html", command.ImportKind)
		}
	default:
		return Command{}, fmt.Errorf("unknown command %q; run apis-mcp help", command.Name)
	}
	return command, nil
}

// Execute runs a parsed one-shot command or the full-screen app.
func Execute(ctx context.Context, runtime *bootstrap.Runtime, args []string, options Options) error {
	command, err := Parse(args)
	if err != nil {
		return err
	}
	options = normalizeOptions(options)
	if command.Name == "help" {
		_, err = io.WriteString(options.Stdout, Help(options.Version))
		return err
	}
	if command.Name == "version" {
		_, err = fmt.Fprintln(options.Stdout, displayVersion(options.Version))
		return err
	}
	if command.Name == "configure" && !command.Explicit && terminalIO(options.Stdin) && terminalIO(options.Stdout) {
		return RunConfigure(ctx, options)
	}
	if command.Name == "install" || command.Name == "configure" || command.Name == "uninstall" {
		return executeInstall(command, options)
	}
	if command.Name == "interactive" {
		if runtime == nil {
			return errors.New("runtime is required for interactive mode")
		}
		return RunInteractive(ctx, runtime, options)
	}
	if runtime == nil {
		return errors.New("runtime is required for this command")
	}
	var value any
	switch command.Name {
	case "rebuild":
		err = runtime.RebuildLibrary(ctx)
		value = "Library index rebuilt. New processes will use the new generation."
	case "collections":
		value, err = runtime.Library.Collections(ctx, library.CollectionsRequest{Page: command.Page})
	case "list":
		value, err = runtime.Library.List(ctx, library.ListRequest{Name: command.NameFilter, Version: command.Version, Collection: command.Collection, Page: command.Page})
	case "pages":
		value, err = runtime.Library.Pages(ctx, library.PagesRequest{DocID: command.DocID, Path: command.Path, Page: command.Page})
	case "search":
		value, err = runtime.Library.Search(ctx, library.SearchRequest{DocID: command.DocID, Query: command.Query, Path: command.Path, Page: command.Page})
	case "read":
		value, err = runtime.Library.Read(ctx, library.ReadRequest{DocID: command.DocID, PageID: command.PageID, Lines: command.Lines})
	case "call":
		value, err = runtime.HTTP.Call(ctx, command.callInput())
	case "sessions":
		value, err = executeSessions(ctx, runtime, command)
	case "cache":
		value, err = runtime.Cache.Cleanup()
	case "config":
		value = runtime.Paths.Config
	case "doctor":
		value = diagnose(runtime, options)
	case "import":
		importOptions := importer.Options{LibraryRoot: runtime.Paths.Library, Rebuild: runtime.RebuildLibrary}
		switch command.ImportKind {
		case "markdown":
			value, err = importer.ImportMarkdown(ctx, command.Source, importOptions)
		case "openapi":
			value, err = importer.ImportOpenAPI(ctx, command.APIName, command.Version, command.Source, importOptions)
		case "llms":
			value, err = importer.ImportLLMSTxt(ctx, command.APIName, command.Version, command.Source, importOptions)
		case "html":
			importOptions.MaxHTMLPages, importOptions.MaxHTMLDepth = command.MaxPages, command.MaxDepth
			value, err = importer.ImportHTML(ctx, command.APIName, command.Version, command.Source, importOptions)
		}
	}
	if err != nil {
		return err
	}
	return RenderHuman(options.Stdout, value)
}

// Run executes and reports errors, returning a process-friendly exit status.
func Run(ctx context.Context, runtime *bootstrap.Runtime, args []string, options Options) int {
	options = normalizeOptions(options)
	if err := Execute(ctx, runtime, args, options); err != nil {
		fmt.Fprintln(options.Stderr, "apis-mcp:", err)
		return 1
	}
	return 0
}

func executeInstall(command Command, options Options) error {
	installerOptions := install.Options{Executable: options.Executable, ClientIDs: command.ClientIDs, All: command.All, DryRun: command.DryRun, Backup: command.Backup}
	var results []install.Result
	var err error
	if command.Name == "uninstall" {
		results, err = install.Uninstall(installerOptions)
	} else {
		results, err = install.Configure(installerOptions)
	}
	if renderErr := RenderHuman(options.Stdout, results); err == nil {
		err = renderErr
	}
	return err
}

func executeSessions(ctx context.Context, runtime *bootstrap.Runtime, command Command) (any, error) {
	switch command.Action {
	case "list":
		return runtime.Sessions.List()
	case "show":
		return runtime.Sessions.Inspect(command.ID)
	case "delete":
		if err := runtime.Sessions.Delete(command.ID); err != nil {
			return nil, err
		}
		return "Session deleted: " + command.ID, nil
	case "cleanup":
		return runtime.Sessions.Cleanup()
	case "create":
		handle, err := runtime.Sessions.Create(ctx)
		if err != nil {
			return nil, err
		}
		id := handle.ID()
		if err := handle.Close(); err != nil {
			return nil, err
		}
		return "Session created: " + id, nil
	default:
		return nil, fmt.Errorf("unknown sessions action %q", command.Action)
	}
}

func (command Command) callInput() app.CallInput {
	return app.CallInput{
		Method: command.Method, Endpoint: command.Endpoint, Headers: command.Headers,
		Payload: command.Payload, Timeout: command.Timeout, Retries: command.Retries,
		JSONPath: command.JSONPath, Session: command.Session,
		AllowLargeDownload: command.AllowLargeDownload,
	}
}

// Diagnostic is one human-oriented local health check.
type Diagnostic struct {
	Name   string
	OK     bool
	Detail string
}

func diagnose(runtime *bootstrap.Runtime, options Options) []Diagnostic {
	checks := []Diagnostic{
		{Name: "configuration", OK: fileExists(runtime.Paths.Config), Detail: runtime.Paths.Config},
		{Name: "library index", OK: runtime.Library != nil && runtime.Library.Fingerprint() != "", Detail: runtime.Library.Fingerprint()},
		{Name: "cache", OK: directoryExists(runtime.Paths.Cache), Detail: runtime.Paths.Cache},
		{Name: "sessions", OK: directoryExists(runtime.Paths.Sessions), Detail: runtime.Paths.Sessions},
	}
	executable := options.Executable
	if executable == "" {
		executable, _ = os.Executable()
	}
	absolute, err := filepath.Abs(executable)
	checks = append(checks, Diagnostic{Name: "executable", OK: err == nil && filepath.IsAbs(absolute), Detail: absolute})
	configured := 0
	for _, status := range install.Detect("", "") {
		if status.Configured {
			configured++
		}
	}
	checks = append(checks, Diagnostic{Name: "MCP clients", OK: configured > 0, Detail: fmt.Sprintf("%d configured", configured)})
	return checks
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func directoryExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func normalizeOptions(options Options) Options {
	if options.Stdin == nil {
		options.Stdin = os.Stdin
	}
	if options.Stdout == nil {
		options.Stdout = os.Stdout
	}
	if options.Stderr == nil {
		options.Stderr = os.Stderr
	}
	if options.Executable == "" {
		options.Executable, _ = os.Executable()
	}
	return options
}

func terminalIO(value any) bool {
	file, ok := value.(interface{ Fd() uintptr })
	return ok && term.IsTerminal(int(file.Fd()))
}

func displayVersion(version string) string {
	if strings.TrimSpace(version) == "" {
		version = "development"
	}
	return "apis-mcp " + version
}

// Help returns CLI help independently of the MCP renderer.
func Help(version string) string {
	return displayVersion(version) + `

Usage:
  apis-mcp                         open the full-screen terminal app
  apis-mcp collections [--page N]
  apis-mcp list [--name TEXT] [--version VERSION] [--collection ID] [--page N]
  apis-mcp pages DOC_ID [--path PATH] [--page N]
  apis-mcp search DOC_ID QUERY [--path PATH] [--page N]
  apis-mcp read DOC_ID PAGE_ID [--lines START:END]
  apis-mcp call METHOD URL [--headers JSON|FILE] [--payload JSON|FILE] [options]
  apis-mcp rebuild
  apis-mcp import markdown DIRECTORY
  apis-mcp import openapi API_NAME VERSION FILE_OR_URL
  apis-mcp import llms API_NAME VERSION FILE_OR_URL
  apis-mcp import html API_NAME VERSION URL [--max-pages N] [--max-depth N]
  apis-mcp sessions [list|create|show ID|delete ID|cleanup]
  apis-mcp cache cleanup
  apis-mcp config [path]
  apis-mcp doctor
  apis-mcp install|configure [--client ID ...|--all] [--dry-run]
  apis-mcp uninstall [--client ID ...|--all] [--dry-run]
  apis-mcp version

Installer client IDs:
  ` + strings.Join(install.ClientIDs(), ", ") + `
`
}

func flags(name string) *flag.FlagSet {
	set := flag.NewFlagSet(name, flag.ContinueOnError)
	set.SetOutput(io.Discard)
	return set
}

func parseFlags(set *flag.FlagSet, args []string, boolFlags map[string]bool) error {
	positionals, err := parsePositionals(set, args, boolFlags)
	if err != nil {
		return err
	}
	if len(positionals) > 0 {
		return unexpected(set.Name(), positionals)
	}
	return nil
}

func parsePositionals(set *flag.FlagSet, args []string, boolFlags map[string]bool) ([]string, error) {
	ordered := make([]string, 0, len(args))
	var positional []string
	for index := 0; index < len(args); index++ {
		argument := args[index]
		if argument == "--" {
			positional = append(positional, args[index+1:]...)
			break
		}
		if !strings.HasPrefix(argument, "-") || argument == "-" {
			positional = append(positional, argument)
			continue
		}
		ordered = append(ordered, argument)
		name := strings.TrimLeft(strings.SplitN(argument, "=", 2)[0], "-")
		if strings.Contains(argument, "=") || boolFlags[name] {
			continue
		}
		if index+1 >= len(args) {
			return nil, fmt.Errorf("flag %s requires a value", argument)
		}
		index++
		ordered = append(ordered, args[index])
	}
	if err := set.Parse(ordered); err != nil {
		return nil, fmt.Errorf("%s: %w", set.Name(), err)
	}
	return positional, nil
}

func unexpected(name string, args []string) error {
	return fmt.Errorf("%s: unexpected arguments: %s", name, strings.Join(args, " "))
}

func parseLineRange(value string) ([]int, error) {
	parts := strings.Split(value, ":")
	if len(parts) != 2 {
		parts = strings.Split(value, "-")
	}
	if len(parts) != 2 {
		return nil, errors.New("lines must be START:END")
	}
	start, firstErr := strconv.Atoi(parts[0])
	end, secondErr := strconv.Atoi(parts[1])
	if firstErr != nil || secondErr != nil || start < 1 || end < start {
		return nil, errors.New("lines must be a positive inclusive START:END range")
	}
	return []int{start, end}, nil
}

func jsonOrPath(value string) (any, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return nil, nil
	}
	if strings.HasPrefix(value, "{") || strings.HasPrefix(value, "[") {
		var parsed any
		if err := json.Unmarshal([]byte(value), &parsed); err != nil {
			return nil, err
		}
		return parsed, nil
	}
	return value, nil
}

type optionalInt struct {
	value int
	set   bool
}

func (value *optionalInt) String() string { return strconv.Itoa(value.value) }
func (value *optionalInt) Set(raw string) error {
	parsed, err := strconv.Atoi(raw)
	if err != nil {
		return err
	}
	value.value, value.set = parsed, true
	return nil
}

type stringList []string

func (values *stringList) String() string { return strings.Join(*values, ",") }
func (values *stringList) Set(value string) error {
	if strings.TrimSpace(value) == "" {
		return errors.New("client ID cannot be empty")
	}
	*values = append(*values, value)
	return nil
}
