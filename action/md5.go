package action

import (
	"crypto-system/action/utils"
	"crypto-system/internal/context"
	"fmt"
	"os"
	"time"
)

type MD5Options struct {
	Filename string
	Verify   bool
}

func MD5(opts *MD5Options) {

	start := time.Now() // 获取当前时间
	file, err := os.Open(opts.Filename)
	context.App.Logger.Error(err)
	defer file.Close()

	absPath := utils.FileAbsPath(file)

	if opts.Verify {
		res, md5 := utils.VerifyMD5(file)
		context.App.Logger.Log(absPath, ":", md5)
		if res["hasFile"].(bool) {
			context.App.Logger.Log("文件已经存在,CID: ", res["CID"])
		}

	} else {
		md5 := utils.MD5sum(file)
		context.App.Logger.Log(absPath, ":", md5)
	}

	elapsed := time.Since(start)
	fmt.Println("该命令执行完成耗时：", elapsed)
}
