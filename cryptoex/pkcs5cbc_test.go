package cryptoex

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPkcs5CBCEncrypter(t *testing.T) {
	plaintext := "hello world"
	password := "password"
	ciphertext, err := Pkcs5CBCEncrypter([]byte(plaintext), []byte(password))
	assert.Nil(t, err)
	assert.Equal(t, "gOuFXWarFTUClEJbi6rsVQ==", base64.StdEncoding.EncodeToString(ciphertext))
}

func TestPkcs5CBCDecrypter(t *testing.T) {
	ciphertext := "gOuFXWarFTUClEJbi6rsVQ=="
	password := "password"
	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	assert.Nil(t, err)
	plaintext, err := Pkcs5CBCDecrypter(decodedCiphertext, []byte(password))
	assert.Nil(t, err)
	assert.Equal(t, "hello world", string(plaintext))
}

func TestPkcs5Keyivgen(t *testing.T) {
	password := "password"
	key, iv := Pkcs5Keyivgen([]byte(password), 32, 16)
	assert.Equal(t, "nXW2O7tivQkXlatopXa4YMtSswqjmtNfLty75wluPEE=", base64.StdEncoding.EncodeToString(key))
	assert.Equal(t, "UAPs4DrvzLUEuY7xNJF3Ng==", base64.StdEncoding.EncodeToString(iv))
}
