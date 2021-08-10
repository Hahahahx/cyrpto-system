package commands

import (
	"crypto-system/action"
	"crypto-system/internal/context"
	"errors"
	"github.com/urfave/cli/v2" // imports as package "cli"
)

func Get() *cli.Command {
	return &cli.Command{
		Name: "get",
		// Aliases: []string{"a"},
		Usage: "get file",
		Action: func(c *cli.Context) error {
			context.Load()

			if c.Args().First() == "" {
				context.App.Logger.Error(errors.New("error Args"))
			}

			action.Get(&context.Request{
				App: context.App,
				Cli: c,
			})

			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Required: true,
				Aliases:  []string{"n"},
				Usage:    "file name <file.key> `string`",
			},
			&cli.BoolFlag{
				Name:    "decrypt",
				Value:   false,
				Aliases: []string{"d"},
				Usage:   "decrypt file `bool`",
			},
		},
	}
}
