package commands

import (
	"crypto-system/action"
	"crypto-system/internal/context"
	"errors"

	"github.com/urfave/cli/v2"
)

func Encrypt() *cli.Command {
	return &cli.Command{
		Name:  "encrypt",
		Usage: "encrypt file on local",
		Action: func(c *cli.Context) error {
			context.Load()

			if c.Args().First() == "" {
				context.App.Logger.Error(errors.New("error Args"))
			}

			action.Encrypt(&action.EncryptOptions{
				Filename: c.Args().First(),
			})

			return nil
		},
	}
}
