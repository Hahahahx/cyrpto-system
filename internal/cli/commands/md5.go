package commands

import (
	"crypto-system/action"
	"crypto-system/internal/context"
	"errors"
	"github.com/urfave/cli/v2" // imports as package "cli"
)

func MD5() *cli.Command {
	return &cli.Command{
		Name: "md5",
		// Aliases: []string{"a"},
		Usage: "md5sum file",
		Action: func(c *cli.Context) error {
			context.Load()

			if c.Args().First() == "" {
				context.App.Logger.Error(errors.New("error Args"))
			}

			action.MD5(&context.Request{
				App: context.App,
				Cli: c,
			})

			return nil
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "verify",
				Value:   false,
				Aliases: []string{"v"},
				Usage:   "verify file from MDMS `bool`",
			},
		},
	}
}
