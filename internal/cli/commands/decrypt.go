package commands

import (
	"crypto-system/action"
	"crypto-system/internal/context"
	"errors"
	"io/ioutil"

	"github.com/urfave/cli/v2"
)

func Decrypt() *cli.Command {
	return &cli.Command{
		Name:  "decrypt",
		Usage: "decrypt file on local",
		Action: func(c *cli.Context) error {
			context.Load()

			if c.Args().First() == "" {
				context.App.Logger.Error(errors.New("error Args"))
			}

			fileKey, err := ioutil.ReadFile(c.String("k"))
			context.App.Logger.Error(err, "load key error")

			action.Decrypt(&action.DecryptOptions{
				Filename: c.Args().First(),
				Key:      fileKey,
				Newname:  c.String("n"),
			})

			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
				Usage:   "encrypt name `string`",
			},
			&cli.StringFlag{
				Name:     "key",
				Required: true,
				Aliases:  []string{"k"},
				Usage:    "crypt key `string`",
			},
		},
	}
}
