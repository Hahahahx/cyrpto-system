package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func main() {
	orig := "http://c.biancheng.net/golang/"
	key := "123456781234567812345678"
	fmt.Println("原文：", orig)
	encryptCode := AESEncrypt([]byte(orig), key)
	fmt.Println("密文：", encryptCode)
	decryptCode := AESDecrypt(encryptCode, key)
	fmt.Println("解密结果：", decryptCode)
}

func AESEncrypt(orig []byte, key string) []byte {
	// 转成字节数组
	k := []byte(key)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	orig = PKCS7Padding(orig, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryted := make([]byte, len(orig))
	// 加密
	blockMode.CryptBlocks(cryted, orig)
	return cryted
}

func AESDecrypt(cryted []byte, key string) []byte {
	k := []byte(key)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(cryted))
	// 解密
	blockMode.CryptBlocks(orig, cryted)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return orig
}

//补码
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
