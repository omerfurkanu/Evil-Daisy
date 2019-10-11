package Security

import (
	"io"

	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
)

// Encrypt ...
func Encrypt(text string, key []byte) []byte {
	text = "kocacizmelimehmetagakocacizmelimehmetagakocacizmelimehmetaga:" + text
	plainText := []byte(text)
	if block, err := aes.NewCipher(key); err == nil {
		cipherText := make([]byte, aes.BlockSize+len(plainText))
		iv := cipherText[:aes.BlockSize]
		if _, err := io.ReadFull(rand.Reader, iv); err == nil {
			stream := cipher.NewCFBEncrypter(block, iv)
			stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)
			return []byte(base64.URLEncoding.EncodeToString(cipherText))
		} else {
			return nil
		}
	} else {
		return nil
	}
}

// Decrypt ...
func Decrypt(cryptoText string, key []byte) []byte {
	if cipherText, err := base64.URLEncoding.DecodeString(cryptoText); err == nil {
		if block, err := aes.NewCipher(key); err == nil {
			if len(cipherText) < aes.BlockSize {

			}

			iv := cipherText[:aes.BlockSize]
			cipherText = cipherText[aes.BlockSize:]
			stream := cipher.NewCFBDecrypter(block, iv)
			stream.XORKeyStream(cipherText, cipherText)
			s := string(cipherText[:])
			return []byte(s[61:])
		} else {
			return nil
		}
	} else {
		return nil
	}
}
