package commands

import (
	"crypto-system/internal/context"
	"fmt"
	"github.com/urfave/cli/v2" // imports as package "cli"
	"log"
	"os"
	"path/filepath"
)

func Init() *cli.Command {
	return &cli.Command{
		Name: "init",
		// Aliases: []string{"a"},
		Usage: "initialize crypto-system configure file",
		Action: func(c *cli.Context) error {

			path := c.String("p")
			if path == "" {
				abs, _ := filepath.Abs(".")
				path = abs
			}
			if AddConfigPathAndSetEnv(path) != nil {

				log.Fatalln("set configure path to \"" + path + "\" is error ")
				return nil
			}

			if context.GenerateConfig() != nil {
				log.Fatalln("generate configure file is error")
				return nil
			}

			_, err := context.GenerateKey()
			if err != nil {
				log.Fatalln("load crypto is error")
				return nil
			}

			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "path",
				Aliases: []string{"p"},
				Usage:   "set configure file path `string`",
			},
		},
	}
}

func AddConfigPathAndSetEnv(path string) error {

	base := filepath.Base(path)

	if base != ".crypto-system" {
		path = filepath.Join(path, ".crypto-system")
	}

	if err := os.MkdirAll(path, 0777); err != nil {
		return err
	}

	fmt.Println("init to path ", path)

	return os.Setenv("CRYPT_SYSTEM_CONFIG_PATH", path)
}
