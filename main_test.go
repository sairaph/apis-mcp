package main

import (
	"bytes"
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/sairaph/apis-mcp/internal/bootstrap"
	"github.com/sairaph/apis-mcp/internal/cli"
)

func TestBareTTYRoutesToLauncherWithoutOpeningRuntime(t *testing.T) {
	var opened, launched, ranCLI bool
	dependencies := runDependencies{
		isTerminal: func() bool { return true },
		openRuntime: func(context.Context) (*bootstrap.Runtime, error) {
			opened = true
			return nil, errors.New("runtime must not open")
		},
		runLauncher: func(_ context.Context, options cli.Options) error {
			launched = options.Version == "test"
			return nil
		},
		runCLI: func(context.Context, *bootstrap.Runtime, []string, cli.Options) int {
			ranCLI = true
			return 0
		},
	}
	options := cli.Options{Version: "test", Stderr: &bytes.Buffer{}}
	if code := runWithDependencies(context.Background(), nil, options, dependencies); code != 0 {
		t.Fatalf("exit code = %d, want 0", code)
	}
	if opened || ranCLI || !launched {
		t.Fatalf("opened=%t ranCLI=%t launched=%t", opened, ranCLI, launched)
	}
}

func TestBareNonTTYStillRoutesToMCP(t *testing.T) {
	runtime := &bootstrap.Runtime{}
	var events []string
	dependencies := runDependencies{
		isTerminal: func() bool { return false },
		openRuntime: func(context.Context) (*bootstrap.Runtime, error) {
			events = append(events, "open")
			return runtime, nil
		},
		serveMCP: func(_ context.Context, got *bootstrap.Runtime, version string) error {
			if got != runtime || version != "test" {
				t.Fatalf("ServeMCP runtime/version = %p/%q", got, version)
			}
			events = append(events, "serve")
			return nil
		},
		closeRuntime: func(got *bootstrap.Runtime) error {
			if got != runtime {
				t.Fatalf("closed runtime = %p, want %p", got, runtime)
			}
			events = append(events, "close")
			return nil
		},
		runLauncher: func(context.Context, cli.Options) error {
			t.Fatal("launcher ran for non-TTY input")
			return nil
		},
	}
	options := cli.Options{Version: "test", Stderr: &bytes.Buffer{}}
	if code := runWithDependencies(context.Background(), nil, options, dependencies); code != 0 {
		t.Fatalf("exit code = %d, want 0", code)
	}
	if want := []string{"open", "serve", "close"}; !reflect.DeepEqual(events, want) {
		t.Fatalf("events = %v, want %v", events, want)
	}
}

func TestExplicitConfigureDoesNotOpenRuntime(t *testing.T) {
	var gotRuntime *bootstrap.Runtime
	var gotArgs []string
	dependencies := runDependencies{
		isTerminal: func() bool { return true },
		openRuntime: func(context.Context) (*bootstrap.Runtime, error) {
			t.Fatal("runtime opened for explicit configure")
			return nil, nil
		},
		runCLI: func(_ context.Context, runtime *bootstrap.Runtime, args []string, _ cli.Options) int {
			gotRuntime, gotArgs = runtime, append([]string(nil), args...)
			return 0
		},
	}
	options := cli.Options{Stderr: &bytes.Buffer{}}
	if code := runWithDependencies(context.Background(), []string{"configure"}, options, dependencies); code != 0 {
		t.Fatalf("exit code = %d, want 0", code)
	}
	if gotRuntime != nil || !reflect.DeepEqual(gotArgs, []string{"configure"}) {
		t.Fatalf("runCLI runtime/args = %p/%v", gotRuntime, gotArgs)
	}
}

func TestLauncherFailureIsReported(t *testing.T) {
	want := errors.New("terminal failed")
	var stderr bytes.Buffer
	dependencies := runDependencies{
		isTerminal: func() bool { return true },
		runLauncher: func(context.Context, cli.Options) error {
			return want
		},
	}
	if code := runWithDependencies(context.Background(), nil, cli.Options{Stderr: &stderr}, dependencies); code != 1 {
		t.Fatalf("exit code = %d, want 1", code)
	}
	if got := stderr.String(); got != "apis-mcp: terminal failed\n" {
		t.Fatalf("stderr = %q", got)
	}
}
