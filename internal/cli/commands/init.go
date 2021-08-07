package commands

import (
	"crypto-system/internal/context"
	"crypto-system/internal/crypto"
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

			// 查看全部环境变量
			// environ := os.Environ()
			// for i := range environ {
			// 	fmt.Println(environ[i])
			// }
			path := c.String("p")
			if path == "" {
				path = os.Getenv("HOMEPATH")
			}
			if AddConfigPathAndSetEnv(path) != nil {

				log.Fatalln("set configure path to \"" + path + "\" is error ")
				return nil
			}

			if context.GenerateConfig() != nil {
				log.Fatalln("generate configure file is error")
				return nil
			}

			_, err := crypto.GenerateKey()
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
	return os.Setenv("CRYPT_SYSTEM_CONFIG_PATH", path)
}
