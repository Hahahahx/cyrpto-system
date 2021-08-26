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
		Usage: "initialize configure file",
		Action: func(c *cli.Context) error {

			// 查看是否有配置环境变量
			// 没有的话则设置在当前目录下构建配置文件夹
			path := os.Getenv("CRYPT_SYSTEM_CONFIG_PATH")
			if path == "" {
				abs, _ := filepath.Abs(".")
				path = abs
			}
			if AddConfigPathAndSetEnv(path) != nil {
				log.Fatalln("set configure path to \"" + path + "\" is error ")
				return nil
			}

			// 创建config.yaml
			if context.GenerateConfig() != nil {
				log.Fatalln("generate configure file is error")
				return nil
			}

			// 创建密钥
			_, err := context.GenerateKey()
			if err != nil {
				log.Fatalln("load crypto is error")
				return nil
			}

			return nil
		},
	}
}

// 设置并生成配置文件夹，并添加到环境变量中
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
