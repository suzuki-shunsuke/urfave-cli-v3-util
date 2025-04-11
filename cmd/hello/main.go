package main

import (
	"context"
	"log"
	"os"

	"github.com/suzuki-shunsuke/urfave-cli-v3-util/helpall"
	"github.com/suzuki-shunsuke/urfave-cli-v3-util/vcmd"
	"github.com/urfave/cli/v3"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	return helpall.With(vcmd.With(&cli.Command{
		Name:    "hello",
		Version: "1.0.0",
		Commands: []*cli.Command{
			{
				Name:        "foo",
				Usage:       "foo command",
				Description: "This is a foo command",
			},
			{
				Name:        "bar",
				Usage:       "bar command",
				Description: "This is a bar command",
			},
		},
	}, "abc123"), nil).Run(context.Background(), os.Args)
}
