//go:build dev

package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"

	"github.com/sairaph/apis-mcp/internal/importer"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	os.Exit(run(ctx, os.Args[1:], os.Stdout, os.Stderr, nil))
}

func run(ctx context.Context, args []string, stdout, stderr io.Writer, client *http.Client) int {
	if len(args) == 0 {
		printUsage(stderr)
		return 2
	}
	switch args[0] {
	case "start":
		return runStart(args[1:], stdout, stderr)
	case "status":
		return runJobLookup("status", args[1:], stdout, stderr, func(store *jobStore, id string) (any, error) { return store.get(id) })
	case "cancel":
		return runJobLookup("cancel", args[1:], stdout, stderr, func(store *jobStore, id string) (any, error) { return store.cancel(id) })
	case "list":
		return runList(args[1:], stdout, stderr)
	case "watch":
		return runWatch(ctx, args[1:], stdout, stderr)
	case "__worker":
		return runWorkerCommand(ctx, args[1:], stderr, client)
	default:
		fmt.Fprintf(stderr, "ingest: unknown command %q\n", args[0])
		printUsage(stderr)
		return 2
	}
}

func runStart(args []string, stdout, stderr io.Writer) int {
	flags := newFlagSet("start", stderr)
	var name, version, output, scope string
	var maxPages, maxDepth int
	var maxSourceBytes, maxTotalBytes int64
	var collections collectionFlag
	flags.StringVar(&name, "name", "", "API name (defaults to the URL host)")
	flags.StringVar(&version, "version", "latest", "API version")
	flags.StringVar(&output, "out", "", "canonical Markdown library root")
	flags.StringVar(&scope, "scope", "path", "crawl scope: path or domain")
	flags.IntVar(&maxPages, "max-pages", -1, "optional maximum HTML pages (-1 is unlimited)")
	flags.IntVar(&maxDepth, "max-depth", -1, "optional maximum HTML link depth (-1 is unlimited)")
	flags.Int64Var(&maxSourceBytes, "max-source-bytes", importer.DefaultMaxSourceBytes, "maximum bytes downloaded from one source")
	flags.Int64Var(&maxTotalBytes, "max-total-bytes", importer.DefaultMaxTotalBytes, "maximum bytes downloaded across all sources")
	flags.Var(&collections, "collections", "comma-separated collection IDs (repeatable)")
	if err := flags.Parse(args); err != nil {
		return 2
	}
	if flags.NArg() != 1 || strings.TrimSpace(output) == "" {
		fmt.Fprintln(stderr, "usage: ingest start -out DIR [options] URL")
		return 2
	}
	source := strings.TrimSpace(flags.Arg(0))
	parsed, err := url.Parse(source)
	if err != nil || parsed.Host == "" || parsed.Scheme != "http" && parsed.Scheme != "https" {
		fmt.Fprintln(stderr, "ingest: source must be an HTTP(S) URL")
		return 2
	}
	if name == "" {
		name = parsed.Hostname()
	}
	if name == "" || strings.TrimSpace(version) == "" || scope != "path" && scope != "domain" || maxPages == 0 || maxPages < -1 || maxDepth < -1 || maxSourceBytes < 1 || maxTotalBytes < maxSourceBytes {
		fmt.Fprintln(stderr, "ingest: invalid name, version, scope, or crawl limit")
		return 2
	}
	store, err := openJobStore(output, true)
	if err != nil {
		return reportError(stderr, err)
	}
	job, err := startDetachedJob(store, ingestRequest{
		Output: store.output, Source: source, Name: name, Version: version, Collections: collections,
		Scope: scope, MaxPages: maxPages, MaxDepth: maxDepth, MaxSourceBytes: maxSourceBytes, MaxTotalBytes: maxTotalBytes,
	})
	if err != nil {
		return reportError(stderr, err)
	}
	return writeJSON(stdout, stderr, job)
}

func runJobLookup(command string, args []string, stdout, stderr io.Writer, operation func(*jobStore, string) (any, error)) int {
	flags := newFlagSet(command, stderr)
	var output string
	flags.StringVar(&output, "out", "", "canonical Markdown library root")
	if err := flags.Parse(args); err != nil {
		return 2
	}
	if flags.NArg() != 1 || output == "" {
		fmt.Fprintf(stderr, "usage: ingest %s -out DIR JOB_ID\n", command)
		return 2
	}
	store, err := openJobStore(output, false)
	if err != nil {
		return reportError(stderr, err)
	}
	value, err := operation(store, flags.Arg(0))
	if err != nil {
		return reportError(stderr, err)
	}
	return writeJSON(stdout, stderr, value)
}

func runList(args []string, stdout, stderr io.Writer) int {
	flags := newFlagSet("list", stderr)
	var output string
	flags.StringVar(&output, "out", "", "canonical Markdown library root")
	if err := flags.Parse(args); err != nil || flags.NArg() != 0 || output == "" {
		fmt.Fprintln(stderr, "usage: ingest list -out DIR")
		return 2
	}
	store, err := openJobStore(output, false)
	if err != nil {
		return reportError(stderr, err)
	}
	jobs, err := store.list()
	if err != nil {
		return reportError(stderr, err)
	}
	return writeJSON(stdout, stderr, jobs)
}

func runWatch(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	flags := newFlagSet("watch", stderr)
	var output string
	flags.StringVar(&output, "out", "", "canonical Markdown library root")
	if err := flags.Parse(args); err != nil || flags.NArg() != 1 || output == "" {
		fmt.Fprintln(stderr, "usage: ingest watch -out DIR JOB_ID")
		return 2
	}
	store, err := openJobStore(output, false)
	if err != nil {
		return reportError(stderr, err)
	}
	job, err := watchJob(ctx, store, flags.Arg(0), stdout)
	if err != nil {
		return reportError(stderr, err)
	}
	if job.State == jobFailed || job.State == jobCanceled {
		return 1
	}
	return 0
}

func runWorkerCommand(ctx context.Context, args []string, stderr io.Writer, client *http.Client) int {
	flags := newFlagSet("__worker", stderr)
	var output, id string
	flags.StringVar(&output, "out", "", "")
	flags.StringVar(&id, "id", "", "")
	if err := flags.Parse(args); err != nil || flags.NArg() != 0 || output == "" || id == "" {
		return 2
	}
	store, err := openJobStore(output, false)
	if err == nil {
		err = runWorker(ctx, store, id, client)
	}
	if err != nil {
		fmt.Fprintln(stderr, "ingest worker:", err)
		return 1
	}
	return 0
}

func newFlagSet(name string, stderr io.Writer) *flag.FlagSet {
	flags := flag.NewFlagSet(name, flag.ContinueOnError)
	flags.SetOutput(stderr)
	return flags
}

func writeJSON(stdout, stderr io.Writer, value any) int {
	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(value); err != nil {
		return reportError(stderr, err)
	}
	return 0
}

func reportError(stderr io.Writer, err error) int {
	fmt.Fprintln(stderr, "ingest:", err)
	return 1
}

func printUsage(writer io.Writer) {
	fmt.Fprintln(writer, "usage: ingest <start|status|watch|cancel|list> [options]")
}

type collectionFlag []string

func (values *collectionFlag) String() string { return strings.Join(*values, ",") }
func (values *collectionFlag) Set(raw string) error {
	seen := make(map[string]bool, len(*values))
	for _, value := range *values {
		seen[value] = true
	}
	for _, value := range strings.Split(raw, ",") {
		value = strings.TrimSpace(value)
		if !validCollection(value) {
			return fmt.Errorf("invalid collection %q; use lowercase letters, digits, and single underscores", value)
		}
		if !seen[value] {
			*values = append(*values, value)
			seen[value] = true
		}
	}
	return nil
}

func validCollection(value string) bool {
	if value == "" || strings.HasPrefix(value, "_") || strings.HasSuffix(value, "_") {
		return false
	}
	previousUnderscore := false
	for _, character := range value {
		if character == '_' {
			if previousUnderscore {
				return false
			}
			previousUnderscore = true
			continue
		}
		if character < 'a' || character > 'z' {
			if character < '0' || character > '9' {
				return false
			}
		}
		previousUnderscore = false
	}
	return true
}
