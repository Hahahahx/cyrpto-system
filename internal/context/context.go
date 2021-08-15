package context

import (
	"crypto-system/internal/crypto"
	"log"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/urfave/cli/v2" // imports as package "cli"
)

type Request struct {
	App *Context
	Cli *cli.Context
}

type Context struct {
	Cli    cli.App
	Crypto *crypto.Crypto
	Config Config
	Logger Logger
	Ipfs   *shell.Shell
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
	App.Ipfs = shell.NewShell(App.Config.Ipfs.Api())

	// res, err := App.Ipfs.ID()

	// if err != nil {
	// 	log.Println("connect ipfs is error")
	// 	return
	// } else {
	// 	fmt.Println("connected ipfs ", res.ID)
	// }

	// App.Logger.Log("ipfs id : " + res.ID)

}
