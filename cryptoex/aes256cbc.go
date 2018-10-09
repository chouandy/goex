package cryptoex

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
)

// Aes256CBCEncrypter 對資料進行 AES-256-CBC 加密
func Aes256CBCEncrypter(plaintext []byte, key []byte, iv []byte) ([]byte, error) {
	// 如果 key 長度不是 32 時
	if len(key) != 32 {
		h := sha256.New()
		h.Write(key)
		key = h.Sum(nil)
	}
	// 進行解密
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 如果 plaintext 除以 block size 的餘數不為 0 時，需擴充 size
	if len(plaintext)%block.BlockSize() != 0 {
		plaintext, _ = Pad(plaintext, block.BlockSize())
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(plaintext, plaintext)

	return plaintext, nil
}

// Aes256CBCDecrypter 對用 AES-256-CBC 加密的文字進行解密
func Aes256CBCDecrypter(ciphertext []byte, key []byte, iv []byte) ([]byte, error) {
	// 如果 key 長度不是 32 時
	if len(key) != 32 {
		h := sha256.New()
		h.Write(key)
		key = h.Sum(nil)
	}
	// 進行解密
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	// 將 ciphertext unpad
	ciphertext, err = Unpad(ciphertext, block.BlockSize())
	if err != nil {
		return nil, err
	}

	return ciphertext, nil
}
