package commands

import (
	"crypto-system/action"
	"crypto-system/internal/context"
	"errors"
	"github.com/urfave/cli/v2" // imports as package "cli"
)

func Encrypt() *cli.Command {
	return &cli.Command{
		Name: "encrypt",
		// Aliases: []string{"a"},
		Usage: "encrypt file",
		Action: func(c *cli.Context) error {
			context.Load()

			if c.Args().First() == "" {
				context.App.Logger.Error(errors.New("error Args"))
			}

			action.Encrypt(&context.Request{
				App: context.App,
				Cli: c,
			})

			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Required: true,
				Value:    "<null>",
				Aliases:  []string{"n"},
				Usage:    "encrypt name <file.key> `string`",
			},
		},
	}
}
