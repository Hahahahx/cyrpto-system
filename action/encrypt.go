package action

import (
	"crypto-system/internal/context"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func Encrypt(c *context.Request) {

	start := time.Now() // 获取当前时间
	filename := c.Cli.Args().First()

	file, err := os.Open(filename)
	c.App.Logger.Error(err)

	fileInfo, err := file.Stat()
	if err != nil {
		c.App.Logger.Error(err)
	}
	// buf := make([]byte, fileInfo.Size())

	// file.Read(buf)

	cache, key := encryptFileCache(c, file)
	cache.Close()

	// encryptData, key := encryptFile(c, buf)

	key = localEncryptKey(c, key)
	c.App.Logger.Log(fileInfo.Name() + "加密完成 √")

	c.App.Logger.Log("保存好您的文件密钥，它将是唯一能解开该文件的重要凭借:\n\t", key)

	download(c, []byte(key), c.Cli.String("n")+".key")
	// download(c, encryptData, c.Cli.String("n"))

	downloadFile := filepath.Join(c.App.Config.Path.Download(), file.Name())

	if filename = c.Cli.String("n"); filename != "" {
		downloadFile = filepath.Join(c.App.Config.Path.Download(), filename)
	}

	fmt.Println(GetAbsPath(c, cache), "\t======>\t", downloadFile)

	os.Rename(GetAbsPath(c, cache), downloadFile)
	elapsed := time.Since(start)
	fmt.Println("该命令执行完成耗时：", elapsed)

}
