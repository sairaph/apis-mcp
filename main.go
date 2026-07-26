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

	if len(args) == 0 && !isTerminal() || len(args) > 0 && args[0] == "mcp" {
		if len(args) > 1 {
			fmt.Fprintln(os.Stderr, "apis-mcp: mcp does not accept arguments")
			return 2
		}
		runtime, err := bootstrap.Open(ctx)
		if err != nil {
			fmt.Fprintln(os.Stderr, "apis-mcp:", err)
			return 1
		}
		serveErr := mcpserver.ServeStdio(ctx, runtime, version)
		closeErr := closeRuntime(runtime)
		if serveErr != nil {
			fmt.Fprintln(os.Stderr, "apis-mcp:", serveErr)
			return 1
		}
		if closeErr != nil {
			fmt.Fprintln(os.Stderr, "apis-mcp: shutdown:", closeErr)
			return 1
		}
		return 0
	}

	command, err := cli.Parse(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "apis-mcp:", err)
		return 2
	}
	if command.Name == "help" || command.Name == "version" || command.Name == "install" || command.Name == "configure" || command.Name == "uninstall" {
		return cli.Run(ctx, nil, args, options)
	}
	runtime, err := bootstrap.Open(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, "apis-mcp:", err)
		return 1
	}
	code := cli.Run(ctx, runtime, args, options)
	if err := closeRuntime(runtime); err != nil {
		fmt.Fprintln(os.Stderr, "apis-mcp: shutdown:", err)
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
