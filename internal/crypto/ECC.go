package crypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/pem"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
)

func ReadFileECCKey(path string) (*Crypto, error) {

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

func GenerateECCKey(path string) (*Crypto, error) {
	// privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	// 应该换成以太坊的S256()，不然其他椭圆曲线方式会导致后续解析失败
	// 因为ecisa内部的byte与公私钥的转换都是默认由S256()实现的
	privateKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	publicBlock := &pem.Block{
		Type:  "公钥",
		Bytes: crypto.FromECDSAPub(&privateKey.PublicKey),
	}
	file, err := os.Create(filepath.Join(path, "public.pem"))
	if err != nil {
		return nil, err
	}
	err = pem.Encode(file, publicBlock)
	if err != nil {
		return nil, err
	}

	privateBlock := &pem.Block{
		Type:  "私钥",
		Bytes: crypto.FromECDSA(privateKey),
	}
	file, err = os.Create(filepath.Join(path, "private.pem"))
	if err != nil {
		return nil, err
	}
	err = pem.Encode(file, privateBlock)
	if err != nil {
		return nil, err
	}

	return &Crypto{
		PublicKey:  publicBlock.Bytes,
		PrivateKey: privateBlock.Bytes,
	}, nil
}

func ECCEncrypt(data []byte, publickKey []byte) ([]byte, error) {

	puk, err := crypto.UnmarshalPubkey(publickKey)

	if err != nil {
		return nil, err
	}

	ct, err := ecies.Encrypt(rand.Reader, ecies.ImportECDSAPublic(puk), data, nil, nil)
	return ct, err
}

func ECCDecrypt(data []byte, privateKey []byte) ([]byte, error) {
	privateEcdsa, err := crypto.ToECDSA(privateKey)
	if err != nil {
		return nil, err
	}
	pt, err := ecies.ImportECDSA(privateEcdsa).Decrypt(data, nil, nil)
	return pt, err
}
