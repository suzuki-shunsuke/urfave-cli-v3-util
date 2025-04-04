package helpall

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/urfave/cli/v3"
)

type Options struct{}

// With appends a new command to show the help of all commands to the given command and returns the given command.
func With(rootCmd *cli.Command, opts *Options) *cli.Command {
	rootCmd.Commands = append(rootCmd.Commands, New(rootCmd, opts))
	return rootCmd
}

// New returns a new command to show the help of all commands.
func New(rootCmd *cli.Command, _ *Options) *cli.Command {
	return &cli.Command{
		Name:   "help-all",
		Hidden: true,
		Usage:  "show all help",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Fprintln(cmd.Writer, "```console")
			fmt.Fprintf(cmd.Writer, "$ %s --help\n", rootCmd.Name)
			if err := cli.ShowAppHelp(rootCmd); err != nil {
				return err
			}
			fmt.Fprintln(cmd.Writer, "```")

			cmdName := "help-all"
			if cmd.Name != "" {
				cmdName = cmd.Name
			}

			for _, c := range rootCmd.Commands {
				if c.Name == cmdName {
					continue
				}
				if err := showCommandHelp(ctx, rootCmd.Writer, c, rootCmd, 2); err != nil { //nolint:mnd
					return err
				}
			}
			return nil
		},
	}
}

func showCommandHelp(ctx context.Context, w io.Writer, cmd, parentCommand *cli.Command, level int) error {
	if cmd.Hidden || cmd.Name == "help" {
		return nil
	}
	command := parentCommand.Name + " " + cmd.Name
	fmt.Fprintf(w, "\n%s %s\n\n", strings.Repeat("#", level), command)
	fmt.Fprintln(w, "```console")
	fmt.Fprintf(w, "$ %s --help\n", command)

	if err := cli.ShowCommandHelp(ctx, parentCommand, cmd.Name); err != nil {
		return err
	}
	fmt.Fprintln(w, "```")

	level2 := level + 1
	for _, subcmd := range cmd.Commands {
		if err := showCommandHelp(ctx, w, subcmd, cmd, level2); err != nil {
			return err
		}
	}
	return nil
}
