package cli

import (
	"crypto-system/internal/cli/commands"
	"github.com/urfave/cli/v2" // imports as package "cli"
	"log"
	"os"
	"sort"
	"time"
)

func init() {

	app := &cli.App{
		EnableBashCompletion: true,
		Name:                 "ipfs-crypto",
		Version:              "v1.0.0",
		Compiled:             time.Now(),
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Uxxx",
				Email: "1219654535@qq.com",
			},
		},
		Copyright: "(c) 2021 金钱猫科技股份有限公司",
		Commands: []*cli.Command{

			commands.Add(),
			commands.Init(),
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
