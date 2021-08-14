package utils

import (
	"crypto-system/action/request"
	"crypto-system/internal/context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"time"
)

func MD5sum(c *context.Request, file *os.File) string {

	start := time.Now() // 获取当前时间
	md5Ctx := md5.New()

	FileForEach(c, file, func(buf []byte) {
		md5Ctx.Write(buf)
	})

	cipherStr := md5Ctx.Sum(nil)
	frigerPrint := hex.EncodeToString(cipherStr)
	elapsed := time.Since(start)
	fmt.Println("MD5计算完成耗时：", elapsed)
	return frigerPrint
}

/**
如果已存在文件，则返回CID
否则返回MD5指纹
*/
func VerifyMD5(c *context.Request, file *os.File) (map[string]interface{}, string) {
	frigerPrint := MD5sum(c, file)
	// http 鉴证是否重复
	return request.VerifyMD5(c, frigerPrint), frigerPrint
}
