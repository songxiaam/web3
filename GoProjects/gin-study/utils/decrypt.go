package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
)

func MD5Hash(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

const (
	ivParameter = "1234567890abcdef" // IV长度必须是16字节
)

const SecretKey = "43&%5=6!*&21"

func Encrypt(key string, data string) (string, error) {
	if len(key) < 8 {
		return "", errors.New("加密失败，key不能小于8位")
	}
	if data == "" {
		return "", nil
	}

	// key 处理成 16/24/32 长度（AES-128/192/256）
	keyBytes := []byte(padKey(key))

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	// 填充数据
	origData := pkcs7Padding([]byte(data), aes.BlockSize)

	blockMode := cipher.NewCBCEncrypter(block, []byte(ivParameter))
	encrypted := make([]byte, len(origData))
	blockMode.CryptBlocks(encrypted, origData)

	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// 补齐 key（到 16/24/32 字节）
func padKey(key string) string {
	for len(key)%16 != 0 {
		key += "0"
	}
	return key[:16] // 只取前 16 字节（AES-128）
}

// PKCS7 补码
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}
