package context

import (
	"crypto-system/internal/crypto"
	"os"
)

func GenerateKey() (*crypto.Crypto, error) {
	configPath := os.Getenv("CRYPT_SYSTEM_CONFIG_PATH")
	crypto, err := crypto.GenerateRSAKey(configPath)
	if err != nil {
		return nil, err
	}
	return crypto, nil
}

func LoadKey(path string) (*crypto.Crypto, error) {
	crypto, err := crypto.ReadFileRSAKey(path)
	if err != nil {
		return nil, err
	}
	return crypto, nil
}
