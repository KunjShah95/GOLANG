package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func encrypt(plaintext, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// Create a new GCM (Galois/Counter Mode) cipher
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Generate a nonce
	nonce := make([]byte, aesGCM.NonceSize())

	// Encrypt the data
	ciphertext := aesGCM.Seal(nil, nonce, []byte(plaintext), nil)

	// Return base64 encoded ciphertext
	return base64.StdEncoding.EncodeToString(append(nonce, ciphertext...)), nil
}

func decrypt(encrypted, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// Create a new GCM cipher
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Decode the base64 encoded input
	data, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	// Extract nonce and ciphertext
	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	// Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func main() {
	key := "thisis32bitlongpassphrase1234567" // AES-256 requires 32 bytes
	plaintext := "Hello, Go encryption!"

	// Encrypt the plaintext
	encrypted, err := encrypt(plaintext, key)
	if err != nil {
		fmt.Println("Error encrypting:", err)
		return
	}
	fmt.Println("Encrypted:", encrypted)

	// Decrypt the ciphertext
	decrypted, err := decrypt(encrypted, key)
	if err != nil {
		fmt.Println("Error decrypting:", err)
		return
	}
	fmt.Println("Decrypted:", decrypted)
}
