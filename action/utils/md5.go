package utils

import (
	"crypto-system/action/request"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"time"
)

// 计算文件MD5
// 分块循环写入，不必当心大文件内存占用问题
func MD5sum(file *os.File) string {

	start := time.Now() // 获取当前时间
	md5Ctx := md5.New()

	FileForEach(file, func(buf []byte) {
		md5Ctx.Write(buf)
	})

	cipherStr := md5Ctx.Sum(nil)
	frigerPrint := hex.EncodeToString(cipherStr)
	elapsed := time.Since(start)
	fmt.Println("MD5计算完成耗时：", elapsed)
	return frigerPrint
}

// 如果已存在文件，则返回CID
// 否则返回MD5指纹
func VerifyMD5(file *os.File) (map[string]interface{}, string) {
	frigerPrint := MD5sum(file)
	// http 鉴证是否重复
	return request.VerifyMD5(frigerPrint), frigerPrint
}
