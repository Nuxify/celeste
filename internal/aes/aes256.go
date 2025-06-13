package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"

	apiError "celeste/internal/errors"
)

// Encrypt performs AES-256-GCM encryption on the given text using the provided key.
func Encrypt(text, keyString string) (string, error) {
	key := []byte(keyString)
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", errors.New(apiError.AES256GCMInvalidKey)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", errors.New(apiError.AES256GCMEncryptionFailed)
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", errors.New(apiError.AES256GCMEncryptionFailed)
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt performs AES-256-GCM decryption on the given encrypted text using the provided key.
func Decrypt(encryptedText, keyString string) (string, error) {
	key := []byte(keyString)

	// decode the base64 string
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", errors.New(apiError.AES256GCMDecryptionFailed)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", errors.New(apiError.AES256GCMInvalidKey)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", errors.New(apiError.AES256GCMDecryptionFailed)
	}

	// get the nonce size
	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", errors.New(apiError.AES256GCMDecryptionFailed)
	}

	// extract the nonce from the ciphertext
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", errors.New(apiError.AES256GCMDecryptionFailed)
	}

	return string(plaintext), nil
}
