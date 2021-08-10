package action

import (
	"crypto-system/action/request"
	"crypto-system/internal/context"
	"crypto-system/internal/crypto"
	"crypto/md5"
	"encoding/hex"
	"os"
	"path/filepath"
)

type EncryptData struct {
	Key  string
	Data []byte
}

/**
如果已存在文件，则返回CID
否则返回MD5指纹
*/
func verifyMD5(c *context.Request, file []byte) request.ResultMD5 {

	md5Ctx := md5.New()
	md5Ctx.Write(file)
	cipherStr := md5Ctx.Sum(nil)
	frigerPrint := hex.EncodeToString(cipherStr)

	// http 鉴证是否重复
	return request.VerifyMD5(c, frigerPrint)
}

func encryptFile(c *context.Request, file []byte) ([]byte, string) {

	aesKey := crypto.GetRandomString(32)

	encryptData := crypto.AESEncrypt(file, aesKey)

	return encryptData, aesKey
}

func decryptFile(c *context.Request, file []byte, key string) []byte {

	return crypto.AESDecrypt(file, key)
}

func localEncryptKey(c *context.Request, key string) string {

	rsaKey, err := crypto.RSAEncrypt(key, c.App.Crypto.PublicKey)

	c.App.Logger.Error(err)

	return rsaKey

}

func remoteEncryptKey(c *context.Request, key string) string {

	// todo：http获取KMS公钥
	res := request.GetKMSKey(c)

	rsaKey, err := crypto.RSAEncrypt(key, []byte(res))

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

	rawKey, err := crypto.RSADecrypt(key, c.App.Crypto.PrivateKey)
	c.App.Logger.Error(err)
	return rawKey

}

func download(c *context.Request, data []byte, filename string) {

	err := os.MkdirAll(c.App.Config.Path.Download(), 0777)
	c.App.Logger.Error(err)

	download := filepath.Join(c.App.Config.Path.Download(), filename)

	file, err := os.OpenFile(download, os.O_CREATE|os.O_WRONLY, 0777)

	c.App.Logger.Error(err)

	defer file.Close()

	file.Write(data)

}
