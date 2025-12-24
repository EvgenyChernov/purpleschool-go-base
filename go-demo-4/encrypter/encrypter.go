package encrypter

import (
	"errors"
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

func (e *Encrypter) Encrypt(str string) (string, error) {

	return "", nil
}

func (e *Encrypter) Decrypt(str string) (string, error) {
	return "", nil
}
