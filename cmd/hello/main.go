package main

import (
	"context"
	"log"
	"os"

	"github.com/suzuki-shunsuke/urfave-cli-v3-util/helpall"
	"github.com/urfave/cli/v3"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	return helpall.With(&cli.Command{
		Name: "hello",
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
	}, nil).Run(context.Background(), os.Args)
}
