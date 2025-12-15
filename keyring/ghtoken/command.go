package ghtoken

import (
	"context"

	"github.com/urfave/cli/v3"
)

type InputSet struct {
	Stdin bool
}

func Command(actor *Actor) *cli.Command {
	setArgs := &InputSet{}

	return &cli.Command{
		Name:        "token",
		Usage:       "Manage GitHub Access token",
		Description: `Manage GitHub Access token by keyring.`,
		Commands: []*cli.Command{
			{
				Name:        "set",
				Usage:       "Set GitHub Access token",
				Description: `Set GitHub Access token to keyring.`,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:        "stdin",
						Usage:       "Read GitHub Access token from stdin",
						Destination: &setArgs.Stdin,
					},
				},
				Action: func(ctx context.Context, _ *cli.Command) error {
					return actor.Set(ctx, setArgs)
				},
			},
			{
				Name:        "remove",
				Aliases:     []string{"rm"},
				Usage:       "Remove GitHub Access token",
				Description: `Remove GitHub Access token from keyring.`,
				Action: func(_ context.Context, _ *cli.Command) error {
					return actor.Remove()
				},
			},
		},
	}
}
