package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io"
	"os"
)

type Encrypter struct {
	key string
}

func NewEncrypter() (*Encrypter, error) {
	key := os.Getenv("ENCRYPT_KEY")
	if key == "" {
		return nil, errors.New("ENCRYPT_KEY is not set")
	}
	return &Encrypter{
		key: key,
	}, nil
}

func (e *Encrypter) Encrypt(plainString []byte) ([]byte, error) {
	// Используем SHA256 для получения ключа фиксированного размера (32 байта для AES-256)
	keyHash := sha256.Sum256([]byte(e.key))
	block, err := aes.NewCipher(keyHash[:])
	if err != nil {
		return nil, err
	}
	aesGSM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, aesGSM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}
	encrypted := aesGSM.Seal(nonce, nonce, plainString, nil)
	return encrypted, nil

}

func (e *Encrypter) Decrypt(encryptedString []byte) ([]byte, error) {
	// Используем SHA256 для получения ключа фиксированного размера (32 байта для AES-256)
	keyHash := sha256.Sum256([]byte(e.key))
	block, err := aes.NewCipher(keyHash[:])
	if err != nil {
		return nil, err
	}
	aesGSM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	NonceSize := aesGSM.NonceSize()

	nonce, encryptedString := encryptedString[:NonceSize], encryptedString[NonceSize:]
	decrypted, err := aesGSM.Open(nil, nonce, encryptedString, nil)
	if err != nil {
		return nil, err
	}
	return decrypted, nil
}
