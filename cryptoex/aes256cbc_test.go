package cryptoex

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAes256CBCEncrypter(t *testing.T) {
	plaintext := "hello world"
	password := "password"
	iv := "0123456789abcdef"
	ciphertext, err := Aes256CBCEncrypter([]byte(plaintext), []byte(password), []byte(iv))
	assert.Nil(t, err)
	assert.Equal(t, "iasq5vF4kH1XQSuwdIbRIg==", base64.StdEncoding.EncodeToString(ciphertext))
}

func TestAes256CBCDecrypter(t *testing.T) {
	ciphertext := "iasq5vF4kH1XQSuwdIbRIg=="
	password := "password"
	iv := "0123456789abcdef"
	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	assert.Nil(t, err)
	plaintext, err := Aes256CBCDecrypter(decodedCiphertext, []byte(password), []byte(iv))
	assert.Nil(t, err)
	assert.Equal(t, "hello world", string(plaintext))
}
