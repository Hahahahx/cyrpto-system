package action

import (
	"crypto-system/internal/context"
	"fmt"
	"os"
	"time"
)

func MD5(c *context.Request) {

	start := time.Now() // 获取当前时间
	filename := c.Cli.Args().First()

	isVerify := c.Cli.Bool("v")
	file, err := os.Open(filename)
	c.App.Logger.Error(err)
	defer file.Close()

	absPath := GetAbsPath(c, file)

	if isVerify {

		res, frigerPrint := verifyMD5(c, file)
		c.App.Logger.Log(absPath, ":", frigerPrint)

		if res["hasFile"].(bool) {
			c.App.Logger.Log("文件已经存在,CID: ", res["CID"])
		}

		elapsed := time.Since(start)
		fmt.Println("该命令执行完成耗时：", elapsed)
		return

	}
	frigerPrint := md5sum(c, file)
	c.App.Logger.Log(absPath, ":", frigerPrint)
	elapsed := time.Since(start)
	fmt.Println("该命令执行完成耗时：", elapsed)
}
