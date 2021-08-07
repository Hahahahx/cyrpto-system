package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ReadFileRSAKey(path string) (*Crypto, error) {

	publicKey, err := ioutil.ReadFile(filepath.Join(path, "public.pem"))
	if err != nil {
		return nil, err
	}
	privateKey, err := ioutil.ReadFile(filepath.Join(path, "private.pem"))
	if err != nil {
		return nil, err
	}
	return &Crypto{
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}, nil
}

func GenerateRSAKey(path string) (*Crypto, error) {
	var bits int
	flag.IntVar(&bits, "b", 2048, "密钥长度，默认为1024位")
	// 生成密钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}

	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	privateBlock := &pem.Block{
		Type:  "私钥",
		Bytes: derStream,
	}

	file, err := os.Create(filepath.Join(path, "private.pem"))
	if err != nil {
		return nil, err
	}

	err = pem.Encode(file, privateBlock)
	if err != nil {
		return nil, err
	}

	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}

	publicBlock := &pem.Block{
		Type:  "公钥",
		Bytes: derPkix,
	}
	file, err = os.Create(filepath.Join(path, "public.pem"))
	if err != nil {
		return nil, err
	}
	err = pem.Encode(file, publicBlock)
	if err != nil {
		return nil, err
	}
	return &Crypto{
		PublicKey:  publicBlock.Bytes,
		PrivateKey: privateBlock.Bytes,
	}, nil
}

// 加密
func RSAEncrypt(origData, publicKey []byte) (string, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return "", errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	pub := pubInterface.(*rsa.PublicKey)

	res, err := rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(res), nil
}

// 解密
func RSADecrypt(ciphertext string, privateKey []byte) ([]byte, error) {

	block, _ := pem.Decode(privateKey)
	if block == nil {
		fmt.Println("block error")
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("priv error")
		return nil, err
	}

	decodeBytes, _ := base64.StdEncoding.DecodeString(ciphertext)

	// cipherBytes := make([]byte, len(decodeBytes))

	// hex.Decode(cipherBytes, decodeBytes)

	res, err := rsa.DecryptPKCS1v15(rand.Reader, priv, decodeBytes)
	if err != nil {

		fmt.Println("DecryptPKCS1v15 error")
		return nil, err
	}
	return res, nil
}
