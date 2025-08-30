package urfave

import (
	"github.com/suzuki-shunsuke/go-stdutil"
	"github.com/suzuki-shunsuke/urfave-cli-v3-util/helpall"
	"github.com/suzuki-shunsuke/urfave-cli-v3-util/vcmd"
	"github.com/urfave/cli/v3"
)

func Command(ldflags *stdutil.LDFlags, cmd *cli.Command) *cli.Command {
	cmd.Version = ldflags.Version
	cmd.EnableShellCompletion = true
	cmd.ConfigureShellCompletionCommand = func(cmd *cli.Command) {
		cmd.Hidden = false
	}
	return helpall.With(vcmd.With(cmd, ldflags.Commit), nil)
}
