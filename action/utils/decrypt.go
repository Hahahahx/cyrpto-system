package utils

import (
	"crypto-system/action/request"
	"crypto-system/internal/context"
	"crypto-system/internal/crypto"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// 解密文件,通过内存读取全部数据
func DecryptFile(c *context.Request, file []byte, key string) []byte {

	start := time.Now() // 获取当前时间
	res := crypto.AesCTR_Encrypt(file, key)
	elapsed := time.Since(start)
	fmt.Println("文件解密完成耗时：", elapsed)

	return res
}

// 解密文件,通过chunk读取文件,避免高内存消耗
// 读取的文件流直接解密并写出
func DecryptFileCache(c *context.Request, file *os.File, filename, key string) {
	start := time.Now() // 获取当前时间
	// 创建一个临时缓存的文件
	daownloadFile, err := os.OpenFile(filepath.Join(c.App.Config.Path.Cache(), file.Name()+".chache"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	c.App.Logger.Error(err)
	FileForEach(c, file, func(buf []byte) {
		encryptData := crypto.AesCTR_Decrypt(buf, key)
		_, err = daownloadFile.Write(encryptData)
		c.App.Logger.Error(err)
	})

	elapsed := time.Since(start)
	fmt.Println("文件加密完成耗时：", elapsed)
}

// 解密文件密钥,通过本地的私钥
func DecryptByLocalKey(c *context.Request, key string) string {

	rawKey, err := crypto.RSADecrypt(key, c.App.Crypto.PrivateKey)

	c.App.Logger.Error(err)

	return rawKey

}

// 解密文件密钥,通过远程KMS的私钥
// 发送自己的公钥给KMS
// 确保解密后的文件密钥也不应该是明文的状态在网络中流通
func DecryptByRemoteKey(c *context.Request, cid string) string {
	// todo: http发送key和自己的公钥到服务端，得到新的密钥
	key := request.DecryptFile(c, cid)
	start := time.Now() // 获取当前时间
	rawKey, err := crypto.RSADecrypt(key, c.App.Crypto.PrivateKey)
	elapsed := time.Since(start)
	fmt.Println("密钥RSA解密完成耗时：", elapsed)
	c.App.Logger.Error(err)
	return rawKey

}
