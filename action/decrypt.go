package action

import (
	"crypto-system/internal/context"
	"fmt"
	"io/ioutil"
	"os"
)

func Decrypt(c *context.Request) {

	filename := c.Cli.Args().First()

	fileKey, err := ioutil.ReadFile(c.Cli.String("k"))
	c.App.Logger.Error(err)

	file, err := os.Open(filename)
	c.App.Logger.Error(err)

	fileInfo, err := file.Stat()
	if err != nil {
		c.App.Logger.Error(err)
	}

	fmt.Println("fileSize : ", fileInfo.Size())

	// buf := make([]byte, 50)
	// file.Read(buf)

	key := localDecryptKey(c, string(fileKey))

	decryptFileCache(c, file, filename, key)

	// decryptData := decryptFile(c, buf, key)
	c.App.Logger.Log(fileInfo.Name() + "解密完成 √")

	// download(c, decryptData, c.Cli.String("n"))
}
