package cryptoex

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/pbkdf2"
)

// FileEncrypter 使用 key 對 src 檔案加密, 並轉存為 dst 檔案
func FileEncrypter(src string, dst string, key []byte) error {
	// Check src file is exist or not
	if _, err := os.Stat(src); os.IsNotExist(err) {
		return err
	}

	// Read src file
	plaintext, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	nonce := make([]byte, 12)

	// Randomizing the nonce
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New)

	block, err := aes.NewCipher(dk)
	if err != nil {
		return err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	// Append the nonce to the end of file
	ciphertext = append(ciphertext, nonce...)

	f, err := os.Create(dst)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, bytes.NewReader(ciphertext))
	if err != nil {
		return err
	}

	return nil
}

// FileDecrypter 使用 key 對加密的 src 檔案進行解密, 並轉存為 dst 檔案
func FileDecrypter(src string, dst string, key []byte) error {
	// Check src file is exist or not
	if _, err := os.Stat(src); os.IsNotExist(err) {
		return err
	}

	// Read src file
	ciphertext, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	salt := ciphertext[len(ciphertext)-12:]
	str := hex.EncodeToString(salt)

	nonce, err := hex.DecodeString(str)
	if err != nil {
		return err
	}

	dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New)

	block, err := aes.NewCipher(dk)
	if err != nil {
		return err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext[:len(ciphertext)-12], nil)
	if err != nil {
		return err
	}

	f, err := os.Create(dst)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, bytes.NewReader(plaintext))
	if err != nil {
		return err
	}

	return nil
}
