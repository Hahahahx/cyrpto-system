package context

import (
	"crypto-system/internal/crypto"
	"log"

	"github.com/urfave/cli/v2" // imports as package "cli"
)

type Context struct {
	Cli    cli.App
	Crypto *crypto.Crypto
	Config Config
	Logger Logger
}

var App *Context

func Load() {

	App = &Context{}

	if LoadConfig() != nil {
		log.Fatalln("load configure file is error, maybe you sould to execute init first")
		return
	}
	if LoadLogger() != nil {
		log.Fatalln("load logger is error")
		return
	}

	crypto, err := LoadKey(App.Config.Path.Config)
	if err != nil {
		log.Fatalln("load crypto is error")
		return
	}

	App.Crypto = crypto

}
