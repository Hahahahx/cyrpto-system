package action

import (
	"crypto-system/internal/context"
	"crypto-system/internal/crypto"
	"io/ioutil"
	"net/http"
)

func Get(c *context.Request) {

	cid := c.Cli.Args().First()

	url := c.App.Config.Ipfs.GetFileURL(cid)

	resp, err := http.Get(url)

	c.App.Logger.Error(err, "get url error")
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	c.App.Logger.Error(err, "read data error")

	isEncrypt := c.Cli.Bool("d")

	if isEncrypt {

		key := remoteDecryptKey(c, cid)

		data = crypto.AESDecrypt(data, key)

	}

	download(c, data, c.Cli.String("n"))

}
