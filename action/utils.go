package action

import (
	"crypto-system/action/request"
	"crypto-system/internal/context"
	"crypto-system/internal/crypto"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"time"
)

type EncryptData struct {
	Key  string
	Data []byte
}

func md5sum(c *context.Request, file *os.File) string {

	start := time.Now() // 获取当前时间
	md5Ctx := md5.New()

	fileForEach(c, file, func(buf []byte) {
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
func verifyMD5(c *context.Request, file *os.File) (map[string]interface{}, string) {
	frigerPrint := md5sum(c, file)
	// http 鉴证是否重复
	return request.VerifyMD5(c, frigerPrint), frigerPrint
}

func encryptFile(c *context.Request, file []byte) ([]byte, string) {

	aesKey := crypto.GetRandomString(32)

	start := time.Now() // 获取当前时间
	encryptData := crypto.AesCTR_Encrypt(file, aesKey)
	elapsed := time.Since(start)
	fmt.Println("文件加密完成耗时：", elapsed)

	return encryptData, aesKey
}

func encryptFileCache(c *context.Request, file *os.File) (*os.File, string) {

	start := time.Now() // 获取当前时间

	// 创建一个临时缓存的文件
	cacheFile, err := os.OpenFile(filepath.Join(c.App.Config.Path.Cache(), file.Name()+".chache"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	c.App.Logger.Error(err)
	// defer cacheFile.Close()

	aesKey := crypto.GetRandomString(32)

	fileForEach(c, file, func(buf []byte) {
		encryptData := crypto.AesCTR_Encrypt(buf, aesKey)
		_, err = cacheFile.Write(encryptData)
		c.App.Logger.Error(err)
	})

	elapsed := time.Since(start)
	fmt.Println("文件加密完成耗时：", elapsed)
	return cacheFile, aesKey
}

func decryptFile(c *context.Request, file []byte, key string) []byte {

	start := time.Now() // 获取当前时间
	res := crypto.AesCTR_Encrypt(file, key)
	elapsed := time.Since(start)
	fmt.Println("文件解密完成耗时：", elapsed)

	return res
}

func decryptFileCache(c *context.Request, file *os.File, filename, key string) {
	start := time.Now() // 获取当前时间
	// 创建一个临时缓存的文件
	daownloadFile, err := os.OpenFile(filepath.Join(c.App.Config.Path.Cache(), file.Name()+".chache"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	c.App.Logger.Error(err)
	fileForEach(c, file, func(buf []byte) {
		encryptData := crypto.AesCTR_Decrypt(buf, key)
		_, err = daownloadFile.Write(encryptData)
		c.App.Logger.Error(err)
	})

	elapsed := time.Since(start)
	fmt.Println("文件加密完成耗时：", elapsed)
}

func localEncryptKey(c *context.Request, key string) string {

	rsaKey, err := crypto.RSAEncrypt(key, c.App.Crypto.PublicKey)

	c.App.Logger.Error(err)

	return rsaKey

}

func remoteEncryptKey(c *context.Request, key string) string {

	// todo：http获取KMS公钥
	res := request.GetKMSKey(c)

	start := time.Now() // 获取当前时间
	rsaKey, err := crypto.RSAEncrypt(key, []byte(res))
	elapsed := time.Since(start)
	fmt.Println("密钥RSA加密完成耗时：", elapsed)

	c.App.Logger.Error(err)

	return rsaKey
}

func localDecryptKey(c *context.Request, key string) string {

	rawKey, err := crypto.RSADecrypt(key, c.App.Crypto.PrivateKey)

	c.App.Logger.Error(err)

	return rawKey

}

func remoteDecryptKey(c *context.Request, cid string) string {

	// todo: http发送key和自己的公钥到服务端，得到新的密钥
	key := request.DecryptFile(c, cid)

	start := time.Now() // 获取当前时间
	rawKey, err := crypto.RSADecrypt(key, c.App.Crypto.PrivateKey)
	elapsed := time.Since(start)
	fmt.Println("密钥RSA解密完成耗时：", elapsed)
	c.App.Logger.Error(err)
	return rawKey

}

func download(c *context.Request, data []byte, filename string) {

	download := filepath.Join(c.App.Config.Path.Download(), filename)

	file, err := os.OpenFile(download, os.O_CREATE|os.O_WRONLY, 0777)

	c.App.Logger.Error(err)

	defer file.Close()

	file.Write(data)

}

func fileForEach(c *context.Request, file *os.File, fn func(buf []byte)) {
	info, _ := file.Stat()
	filesize := info.Size()
	filechunk := c.App.Config.File.Chunk

	blocks := uint64(math.Ceil(float64(filesize) / float64(filechunk)))

	for i := uint64(0); i < blocks; i++ {

		// 末尾数据小于chunk的话
		blocksize := int(math.Min(float64(filechunk), float64(filesize-int64(i*uint64(filechunk)))))

		if filesize < filechunk {
			blocksize = int(filesize)
		}

		buf := make([]byte, blocksize)
		file.Read(buf)
		fn(buf)
	}

}

func GetAbsPath(c *context.Request, file *os.File) string {

	absPath, err := filepath.Abs(filepath.Dir(file.Name()))
	c.App.Logger.Error(err)

	return filepath.Join(absPath, filepath.Base(file.Name()))
}
