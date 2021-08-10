package commands

import (
	"crypto-system/action"
	"crypto-system/internal/context"
	"errors"
	"github.com/urfave/cli/v2" // imports as package "cli"
)

func Add() *cli.Command {
	return &cli.Command{
		Name: "add",
		// Aliases: []string{"a"},
		Usage: "add a task to the list",
		Action: func(c *cli.Context) error {
			context.Load()

			if c.Args().First() == "" {
				context.App.Logger.Error(errors.New("error Args"))
			}

			action.Add(&context.Request{
				App: context.App,
				Cli: c,
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
