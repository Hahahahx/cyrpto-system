package commands

import (
	"crypto-system/action"
	"crypto-system/internal/context"

	"github.com/urfave/cli/v2"
)

func List() *cli.Command {
	return &cli.Command{
		Name:  "ls",
		Usage: "show all encrypt file",
		Action: func(c *cli.Context) error {
			context.Load()

			action.List()

			return nil
		},
	}
}
