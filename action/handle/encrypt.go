package handle

import (
	"crypto-system/action/request"
	"crypto-system/internal/context"
	"crypto-system/internal/crypto"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// 加密文件,通过内存读取全部数据
// ！弃用
func EncryptFile(file []byte) ([]byte, string) {

	aesKey := crypto.GetRandomString(32)

	start := time.Now() // 获取当前时间
	encryptData := crypto.AesCTR_crypter(file, aesKey)
	elapsed := time.Since(start)
	fmt.Println("文件加密完成耗时：", elapsed)

	return encryptData, aesKey
}

// 加密文件,通过chunk大小读取文件
// 避免了高内存的消耗,但是需要确保本地有足够的空间
// 在本地赋值了一份加密的文件,然后再通过该文件去上传
func EncryptFileCache(file *os.File) (*os.File, string) {

	start := time.Now() // 获取当前时间

	// 创建一个临时缓存的文件
	cacheFile, err := os.OpenFile(filepath.Join(context.App.Config.Path.Cache(), filepath.Base(file.Name())+".cache"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	context.App.Logger.Error(err)
	// defer cacheFile.Close()

	aesKey := crypto.GetRandomString(16)

	FileForEach(file, func(buf []byte) {
		encryptData := crypto.AesCTR_crypter(buf, aesKey)
		_, err = cacheFile.Write(encryptData)
		context.App.Logger.Error(err)
	})

	elapsed := time.Since(start)
	fmt.Println("文件加密完成耗时：", elapsed)
	return cacheFile, aesKey
}

// 加密文件密钥,通过本地的公钥
func EncryptByLocalKey(key string) string {

	rsaKey, err := crypto.RSAEncrypt(key, context.App.Crypto.PublicKey)

	context.App.Logger.Error(err)

	return rsaKey

}

// 加密文件密钥,通过KMS的公钥
func EncryptByRemoteKey(key string) string {

	// http获取KMS公钥
	res := request.GetKMSKey()

	start := time.Now() // 获取当前时间
	rsaKey, err := crypto.RSAEncrypt(key, []byte(res))
	elapsed := time.Since(start)
	fmt.Println("密钥RSA加密完成耗时：", elapsed)

	context.App.Logger.Error(err)

	return rsaKey
}
