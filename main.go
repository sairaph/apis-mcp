package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sairaph/apis-mcp/internal/bootstrap"
	"github.com/sairaph/apis-mcp/internal/cli"
	"github.com/sairaph/apis-mcp/internal/mcpserver"
	"golang.org/x/term"
)

var version = "dev"

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	os.Exit(run(ctx, os.Args[1:]))
}

func run(ctx context.Context, args []string) int {
	executable, _ := os.Executable()
	options := cli.Options{
		Version: version, Executable: executable,
		Stdin: os.Stdin, Stdout: os.Stdout, Stderr: os.Stderr,
	}
	return runWithDependencies(ctx, args, options, defaultRunDependencies())
}

type runDependencies struct {
	isTerminal   func() bool
	openRuntime  func(context.Context) (*bootstrap.Runtime, error)
	closeRuntime func(*bootstrap.Runtime) error
	serveMCP     func(context.Context, *bootstrap.Runtime, string) error
	runCLI       func(context.Context, *bootstrap.Runtime, []string, cli.Options) int
	runLauncher  func(context.Context, cli.Options) error
}

func defaultRunDependencies() runDependencies {
	return runDependencies{
		isTerminal: isTerminal, openRuntime: bootstrap.Open, closeRuntime: closeRuntime,
		serveMCP: func(ctx context.Context, runtime *bootstrap.Runtime, version string) error {
			return mcpserver.ServeStdio(ctx, runtime, version)
		},
		runCLI: cli.Run, runLauncher: cli.RunLauncher,
	}
}

func runWithDependencies(ctx context.Context, args []string, options cli.Options, dependencies runDependencies) int {
	if len(args) == 0 && dependencies.isTerminal() {
		if err := dependencies.runLauncher(ctx, options); err != nil {
			fmt.Fprintln(options.Stderr, "apis-mcp:", err)
			return 1
		}
		return 0
	}

	if len(args) == 0 || len(args) > 0 && args[0] == "mcp" {
		if len(args) > 1 {
			fmt.Fprintln(options.Stderr, "apis-mcp: mcp does not accept arguments")
			return 2
		}
		runtime, err := dependencies.openRuntime(ctx)
		if err != nil {
			fmt.Fprintln(options.Stderr, "apis-mcp:", err)
			return 1
		}
		serveErr := dependencies.serveMCP(ctx, runtime, options.Version)
		closeErr := dependencies.closeRuntime(runtime)
		if serveErr != nil {
			fmt.Fprintln(options.Stderr, "apis-mcp:", serveErr)
			return 1
		}
		if closeErr != nil {
			fmt.Fprintln(options.Stderr, "apis-mcp: shutdown:", closeErr)
			return 1
		}
		return 0
	}

	command, err := cli.Parse(args)
	if err != nil {
		fmt.Fprintln(options.Stderr, "apis-mcp:", err)
		return 2
	}
	if command.Name == "help" || command.Name == "version" || command.Name == "install" || command.Name == "configure" || command.Name == "uninstall" {
		return dependencies.runCLI(ctx, nil, args, options)
	}
	runtime, err := dependencies.openRuntime(ctx)
	if err != nil {
		fmt.Fprintln(options.Stderr, "apis-mcp:", err)
		return 1
	}
	code := dependencies.runCLI(ctx, runtime, args, options)
	if err := dependencies.closeRuntime(runtime); err != nil {
		fmt.Fprintln(options.Stderr, "apis-mcp: shutdown:", err)
		return 1
	}
	return code
}

func closeRuntime(runtime *bootstrap.Runtime) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	return runtime.Close(ctx)
}

func isTerminal() bool {
	return term.IsTerminal(int(os.Stdin.Fd())) && term.IsTerminal(int(os.Stdout.Fd()))
}
