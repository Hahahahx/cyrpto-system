package commands

import (
	"crypto-system/action"
	"crypto-system/internal/context"
	"errors"

	"github.com/urfave/cli/v2"
)

func MD5() *cli.Command {
	return &cli.Command{
		Name:  "md5",
		Usage: "md5 file",
		Action: func(c *cli.Context) error {
			context.Load()

			if c.Args().First() == "" {
				context.App.Logger.Error(errors.New("error Args"))
			}

			action.MD5(&action.MD5Options{
				Filename: c.Args().First(),
				Verify:   c.Bool("v"),
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
