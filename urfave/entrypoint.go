package urfave

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"github.com/suzuki-shunsuke/go-error-with-exit-code/ecerror"
	"github.com/suzuki-shunsuke/slog-error/slogerr"
	"github.com/suzuki-shunsuke/slog-util/slogutil"
	"github.com/urfave/cli/v3"
)

type Env struct {
	Program string
	Version string
	Stdin   *os.File
	Stdout  *os.File
	Stderr  *os.File
	Getenv  func(string) string
	Args    []string
}

func Main(name, version string, run Run, args ...any) {
	if code := core(name, version, run, args...); code != 0 {
		os.Exit(code)
	}
}

type Run func(ctx context.Context, logger *slogutil.Logger, env *Env) error

var ErrSilent = errors.New("")

// getVersion returns version if it isn't empty.
// Otherwise, it falls back to the module version embedded in the binary by the Go toolchain,
// so that the version is available even if the binary is built without -ldflags (e.g. go install).
func getVersion(version string) string {
	if version != "" {
		return version
	}
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "" {
		return info.Main.Version
	}
	return "unknown"
}

func core(name, version string, run Run, args ...any) int {
	version = getVersion(version)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	logger := slogutil.New(&slogutil.InputNew{
		Name:    name,
		Version: version,
		Out:     os.Stderr,
		Attrs:   args,
	})
	env := &Env{
		Program: name,
		Version: version,
		Stdin:   os.Stdin,
		Stdout:  os.Stdout,
		Stderr:  os.Stderr,
		Getenv:  os.Getenv,
		Args:    os.Args,
	}
	if err := run(ctx, logger, env); err != nil {
		if err.Error() != "" {
			slogerr.WithError(logger.Logger, err).Error(name + " failed")
		}
		return ecerror.GetExitCode(err)
	}
	return 0
}

type ActionFunc func(ctx context.Context, c *cli.Command, logger *slogutil.Logger) error

func Action(fn ActionFunc, logger *slogutil.Logger) cli.ActionFunc {
	return func(ctx context.Context, c *cli.Command) error {
		return fn(ctx, c, logger)
	}
}
