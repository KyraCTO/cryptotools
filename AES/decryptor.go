package AES

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	mr "math/rand"
	"time"
)

func DecryptAESPayload(aesKey string, payload string) string {
	rawPayload, _ := hex.DecodeString(payload)

	cipherData, err := aes.NewCipher([]byte(aesKey))
	if err != nil {
		fmt.Println(err)
	}
	gcm, err := cipher.NewGCM(cipherData)
	if err != nil {
		fmt.Println(err)
	}
	nonceSize := gcm.NonceSize()
	if len(rawPayload) < nonceSize {
		fmt.Println(err)
	}
	nonce, ciphertext := rawPayload[:nonceSize], rawPayload[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Plain Text from payload: ")
	fmt.Println(string(plaintext))
	return string(plaintext)
}

func EncryptAESPayload(aesKey string, payload string) string {
	cipherData,err := aes.NewCipher([]byte(aesKey))
	if err != nil {
		fmt.Println(err)
	}
	gcm, err := cipher.NewGCM(cipherData)
	if err != nil {
		fmt.Println(err)
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}
	sealedData := gcm.Seal(nonce, nonce, []byte(payload), nil)
	fmt.Printf("%x \n", sealedData)
	return hex.EncodeToString(sealedData)
}

func GenerateAESKey() string {
	var symKey = passPhrase(24)
	fmt.Println("symKey raw: ", symKey)
	symKey = base64.StdEncoding.EncodeToString([]byte(symKey))
	fmt.Println(len(symKey))
	fmt.Println("symKey encoded: ", symKey)
	return symKey
}
func passPhrase(n int) string {
	mr.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	buff := make([]rune, n)
	for i := range buff {
		buff[i] = letters[mr.Intn(len(letters))]
	}
	return string(buff)
}
