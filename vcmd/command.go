package vcmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/urfave/cli/v3"
)

type Command struct {
	Name    string
	Version string
	SHA     string
	Stdout  io.Writer
}

func New(cmd *Command) *cli.Command {
	return cmd.New()
}

func With(cmd *cli.Command, sha string) *cli.Command {
	cmd.Commands = append(cmd.Commands, New(&Command{
		Name:    cmd.Name,
		Version: cmd.Version,
		SHA:     sha,
		Stdout:  cmd.Writer,
	}))
	return cmd
}

func (cmd *Command) New() *cli.Command {
	if cmd == nil {
		cmd = &Command{}
	}
	if cmd.Stdout == nil {
		cmd.Stdout = os.Stdout
	}
	return &cli.Command{
		Name:   "version",
		Usage:  "Show version",
		Action: cmd.action,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "json",
				Aliases: []string{"j"},
				Usage:   "Output version in JSON format",
			},
		},
	}
}

func (cmd *Command) action(_ context.Context, c *cli.Command) error {
	if c.Bool("json") {
		encoder := json.NewEncoder(cmd.Stdout)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(map[string]string{
			"name":    cmd.Name,
			"version": cmd.Version,
			"sha":     cmd.SHA,
		}); err != nil {
			return fmt.Errorf("encode JSON: %w", err)
		}
		return nil
	}
	version := cmd.Version
	if version == "" {
		version = "unknown"
	}
	fmt.Fprintln(cmd.Stdout, version)
	return nil
}
