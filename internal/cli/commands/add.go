package commands

import (
	"fmt"
	"github.com/urfave/cli/v2" // imports as package "cli"
)

func Add() *cli.Command {
	return &cli.Command{
		Name: "add",
		// Aliases: []string{"a"},
		Usage: "add a task to the list",
		Action: func(c *cli.Context) error {
			filename := c.Args().First()
			if c.Bool("e") {
				fmt.Println("added task: ", filename)
			} else {
				fmt.Println("added task: ", filename)
			}
			return nil
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "encryt",
				Value:   false,
				Aliases: []string{"e"},
				Usage:   "encrypt upload `bool`",
			},
		},
	}
}
