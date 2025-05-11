package urfave

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/urfave-cli-v3-util/helpall"
	"github.com/suzuki-shunsuke/urfave-cli-v3-util/vcmd"
	"github.com/urfave/cli/v3"
)

type LDFlags struct {
	Version string
	Commit  string
	Date    string
}

type Runner struct {
	Stdin   io.Reader
	Stdout  io.Writer
	Stderr  io.Writer
	LDFlags *LDFlags
	LogE    *logrus.Entry
}

func Command(logE *logrus.Entry, ldflags *LDFlags, cmd *cli.Command) *cli.Command {
	r := &Runner{
		Stdin:   os.Stdin,
		Stdout:  os.Stdout,
		Stderr:  os.Stderr,
		LDFlags: ldflags,
		LogE:    logE,
	}
	cmd.Version = ldflags.Version
	r.Command(cmd)
	return cmd
}

func (r *Runner) Command(cmd *cli.Command) *cli.Command {
	cmd.ConfigureShellCompletionCommand = func(cmd *cli.Command) {
		cmd.Hidden = false
	}
	return helpall.With(vcmd.With(cmd, r.LDFlags.Commit), nil)
}
