package commands

import (
	"crypto-system/action"
	"crypto-system/internal/context"
	"errors"

	"github.com/urfave/cli/v2"
)

func Add() *cli.Command {
	return &cli.Command{
		Name:  "add",
		Usage: "add a task to the list",
		Action: func(c *cli.Context) error {
			context.Load()

			if c.Args().First() == "" {
				context.App.Logger.Error(errors.New("error Args"))
			}

			action.Add(&action.AddOptions{
				Filename: c.Args().First(),
				Encrypt:  c.Bool("e"),
			})

			return nil
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "encrypt",
				Value:   false,
				Aliases: []string{"e"},
				Usage:   "encrypt upload `bool`",
			},
		},
	}
}
