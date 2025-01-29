package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"ladipage_server/common/configs"
)

func EncryptAes(plainText string) (string, error) {
	key := configs.Get().KeyAes
	keyBytes := make([]byte, 32)
	copy(keyBytes, []byte(key))

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	plainTextBytes := []byte(plainText)
	padding := aes.BlockSize - (len(plainTextBytes) % aes.BlockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	plainTextBytes = append(plainTextBytes, padText...)

	cipherText := make([]byte, aes.BlockSize+len(plainTextBytes))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCBCEncrypter(block, iv)
	stream.CryptBlocks(cipherText[aes.BlockSize:], plainTextBytes)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func DecryptAes(encryptedText string) (string, error) {
	key := configs.Get().KeyAes
	cipherText, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	keyBytes := make([]byte, 32)
	copy(keyBytes, []byte(key))

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCBCDecrypter(block, iv)
	stream.CryptBlocks(cipherText, cipherText)

	padding := int(cipherText[len(cipherText)-1])
	if padding > aes.BlockSize || padding == 0 {
		return "", fmt.Errorf("invalid padding")
	}
	for i := len(cipherText) - padding; i < len(cipherText); i++ {
		if cipherText[i] != byte(padding) {
			return "", fmt.Errorf("invalid padding")
		}
	}

	return string(cipherText[:len(cipherText)-padding]), nil
}
