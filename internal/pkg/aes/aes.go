package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/global"
	log "github.com/sirupsen/logrus"
	"io"
)

var key []byte

func init() {
	if len(global.Config.AesKey) != 32 {
		log.Fatalln("AES key 长度不为 32")
	}
	key = []byte(global.Config.AesKey)
}

func Encrypt(str []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(str))
	iv := ciphertext[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], str)

	return ciphertext, nil
}

func Decrypt(str []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(str) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	iv := str[:aes.BlockSize]
	str = str[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(str, str)

	return str, nil
}
