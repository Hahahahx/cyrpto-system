package action

import (
	"crypto-system/internal/context"
	"crypto-system/internal/crypto"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Get(c *context.Request) {

	start := time.Now() // 获取当前时间
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

		ases := time.Now() // 获取当前时间
		data = crypto.AesCTR_Decrypt(data, key)
		elapsed := time.Since(ases)
		fmt.Println("文件AES解密完成耗时：", elapsed)

	}

	download(c, data, c.Cli.String("n"))
	elapsed := time.Since(start)
	fmt.Println("该命令执行完成耗时：", elapsed)

}
