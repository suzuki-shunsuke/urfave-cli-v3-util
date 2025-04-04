package vcmd

import (
	"context"
	"encoding/json"
	"errors"
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

func (cmd *Command) New() (*cli.Command, error) {
	if cmd == nil {
		return nil, errors.New("Command is nil")
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
			},
		},
	}, nil
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
	fmt.Println(cmd.Stdout, cmd.version())
	return nil
}

func (cmd *Command) version() string {
	if cmd.Name == "" {
		if cmd.Version == "" {
			if cmd.SHA == "" {
				return "unknown"
			}
			return cmd.SHA
		}
		if cmd.SHA == "" {
			return cmd.Version
		}
		return fmt.Sprintf("%s (%s)", cmd.Version, cmd.SHA)
	}
	if cmd.Version == "" {
		if cmd.SHA == "" {
			return "unknown"
		}
		return fmt.Sprintf("%s (%s)", cmd.Name, cmd.SHA)
	}
	if cmd.SHA == "" {
		return fmt.Sprintf("%s %s", cmd.Name, cmd.Version)
	}
	return fmt.Sprintf("%s %s (%s)", cmd.Name, cmd.Version, cmd.SHA)
}
