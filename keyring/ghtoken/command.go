package ghtoken

import (
	"context"

	"github.com/urfave/cli/v3"
)

type Actor interface {
	Remove(ctx context.Context, c *cli.Command) error
	Set(ctx context.Context, c *cli.Command) error
}

func Command(actor Actor) *cli.Command {
	return &cli.Command{
		Name:        "token",
		Usage:       "Manage GitHub Access token",
		Description: `Manage GitHub Access token by keyring.`,
		Commands: []*cli.Command{
			{
				Name:        "set",
				Usage:       "Set GitHub Access token",
				Description: `Set GitHub Access token to keyring.`,
				Action:      actor.Set,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "stdin",
						Usage: "Read GitHub Access token from stdin",
					},
				},
			},
			{
				Name:        "remove",
				Aliases:     []string{"rm"},
				Usage:       "Remove GitHub Access token",
				Description: `Remove GitHub Access token from keyring.`,
				Action:      actor.Remove,
			},
		},
	}
}
