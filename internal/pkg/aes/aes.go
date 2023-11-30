package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"github.com/ncuhome/GeniusAuthoritarianGate/internal/global"
	log "github.com/sirupsen/logrus"
	"unsafe"
)

var key []byte

func init() {
	if len(global.Config.AesKey) != 32 {
		log.Fatalln("AES key 长度不为 32")
	}
	key = []byte(global.Config.AesKey)
}

func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}
func pkcs7UnPadding(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("input data is empty")
	}

	padding := int(data[len(data)-1])
	if padding > len(data) || padding == 0 {
		return nil, errors.New("invalid padding")
	}

	for i := len(data) - padding; i < len(data); i++ {
		if data[i] != byte(padding) {
			return nil, errors.New("invalid padding")
		}
	}

	return data[:len(data)-padding], nil
}

func Encrypt(str []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	str = pkcs7Padding(str, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	encrypted := make([]byte, len(str))
	blockMode.CryptBlocks(encrypted, str)

	output := make([]byte, base64.URLEncoding.EncodedLen(len(encrypted)))
	base64.URLEncoding.Encode(output, encrypted)

	return output, nil
}

func EncryptString(str string) (string, error) {
	output, err := Encrypt(unsafe.Slice(unsafe.StringData(str), len(str)))
	if err != nil {
		return "", err
	}
	return unsafe.String(unsafe.SliceData(output), len(output)), nil
}

func Decrypt(str []byte) ([]byte, error) {
	ciphertext := make([]byte, base64.URLEncoding.DecodedLen(len(str)))
	n, err := base64.URLEncoding.Decode(ciphertext, str)
	if err != nil {
		return nil, err
	}
	ciphertext = ciphertext[:n]

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < block.BlockSize() || len(ciphertext)%block.BlockSize() != 0 {
		return nil, errors.New("invalid ciphertext")
	}

	blockMode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	decrypted := make([]byte, len(ciphertext))
	blockMode.CryptBlocks(decrypted, ciphertext)

	return pkcs7UnPadding(decrypted)
}

func DecryptString(str string) (string, error) {
	output, err := Decrypt(unsafe.Slice(unsafe.StringData(str), len(str)))
	if err != nil {
		return "", err
	}
	return unsafe.String(unsafe.SliceData(output), len(output)), nil
}
