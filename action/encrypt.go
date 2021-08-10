package action

import (
	"crypto-system/internal/context"
	"os"
)

func Encrypt(c *context.Request) {

	filename := c.Cli.Args().First()

	file, err := os.Open(filename)
	c.App.Logger.Error(err)

	fileInfo, err := file.Stat()
	if err != nil {
		c.App.Logger.Error(err)
	}
	buf := make([]byte, fileInfo.Size())

	file.Read(buf)

	encryptData, key := encryptFile(c, buf)

	key = localEncryptKey(c, key)
	c.App.Logger.Log(fileInfo.Name() + "加密完成 √")

	c.App.Logger.Log("保存好您的文件密钥，它将是唯一能解开该文件的重要凭借:\n\t", key)

	download(c, []byte(key), c.Cli.String("n")+".key")
	download(c, encryptData, c.Cli.String("n"))

}
