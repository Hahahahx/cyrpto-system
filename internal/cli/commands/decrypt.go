package commands

import (
	"crypto-system/action"
	"crypto-system/internal/context"
	"errors"
	"github.com/urfave/cli/v2" // imports as package "cli"
)

func Decrypt() *cli.Command {
	return &cli.Command{
		Name: "decrypt",
		// Aliases: []string{"a"},
		Usage: "decrypt file",
		Action: func(c *cli.Context) error {
			context.Load()

			if c.Args().First() == "" {
				context.App.Logger.Error(errors.New("error Args"))
			}

			action.Decrypt(&context.Request{
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
				Usage:    "encrypt name `string`",
			},
			&cli.StringFlag{
				Name:     "key",
				Required: true,
				Aliases:  []string{"key"},
				Usage:    "crypt key `string`",
			},
		},
	}
}
