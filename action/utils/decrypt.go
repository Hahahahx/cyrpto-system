package utils

import (
	"crypto-system/action/request"
	"crypto-system/internal/context"
	"crypto-system/internal/crypto"
	"fmt"
	"os"
	"time"
)

// 解密文件,通过内存读取全部数据
// ！弃用
func DecryptFile(file []byte, key string) []byte {

	start := time.Now() // 获取当前时间
	res := crypto.AesCTR_crypter(file, key)
	elapsed := time.Since(start)
	fmt.Println("文件解密完成耗时：", elapsed)

	return res
}

// 解密文件,通过chunk读取文件,避免高内存消耗
// 读取的文件流直接解密并写出
func DecryptFileCache(file *os.File, filename, key string) {
	start := time.Now() // 获取当前时间

	daownloadFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	context.App.Logger.Error(err)

	// 读取流同时解密并写入到新的文件中
	FileForEach(file, func(buf []byte) {
		encryptData := crypto.AesCTR_crypter(buf, key)
		_, err = daownloadFile.Write(encryptData)
		context.App.Logger.Error(err)
	})

	elapsed := time.Since(start)
	fmt.Println("文件解密完成耗时：", elapsed)
}

// 解密文件密钥,通过本地的私钥
func DecryptByLocalKey(key string) string {

	rawKey, err := crypto.RSADecrypt(key, context.App.Crypto.PrivateKey)

	context.App.Logger.Error(err)

	return rawKey

}

// 解密文件密钥,通过远程KMS的私钥
// 发送自己的公钥给KMS
// 确保解密后的文件密钥也不应该是明文的状态在网络中流通
func DecryptByRemoteKey(cid string) string {
	// todo: http发送key和自己的公钥到服务端，得到新的密钥
	key := request.DecryptFile(cid)
	start := time.Now() // 获取当前时间
	rawKey, err := crypto.RSADecrypt(key, context.App.Crypto.PrivateKey)
	elapsed := time.Since(start)
	fmt.Println("密钥RSA解密完成耗时：", elapsed)
	context.App.Logger.Error(err)
	return rawKey

}
