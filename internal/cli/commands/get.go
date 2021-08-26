package commands

import (
	"crypto-system/action"
	"crypto-system/internal/context"
	"errors"

	"github.com/urfave/cli/v2"
)

func Get() *cli.Command {
	return &cli.Command{
		Name:  "get",
		Usage: "get file from ipfs",
		Action: func(c *cli.Context) error {
			context.Load()

			if c.Args().First() == "" {
				context.App.Logger.Error(errors.New("error Args"))
			}

			action.Get(&action.GetOptions{
				CID:     c.Args().First(),
				Decrypt: c.Bool("d"),
				Newname: c.String("n"),
			})

			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
				Usage:   "file name <file.key> `string`",
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
