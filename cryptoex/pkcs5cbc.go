package cryptoex

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
)

// Pkcs5CBCEncrypter 對文字進行 AES-256-CBC Pkcs5 加密
func Pkcs5CBCEncrypter(plaintext []byte, password []byte) ([]byte, error) {
	// 根據 password 產生 key, iv
	key, iv := Pkcs5Keyivgen(password, 32, 16)
	// 進行加密
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

// Pkcs5CBCDecrypter 對文字進行 AES-256-CBC Pkcs5 解密
func Pkcs5CBCDecrypter(ciphertext []byte, password []byte) ([]byte, error) {
	// 根據 password 產生 key, iv
	key, iv := Pkcs5Keyivgen(password, 32, 16)
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

// Pkcs5Keyivgen 將 password 轉換成 Pkcs5 用的 key, iv
func Pkcs5Keyivgen(password []byte, keylen int, ivlen int) ([]byte, []byte) {
	totalLen := keylen + ivlen
	d := make([][]byte, totalLen/16)
	d[0] = password
	for i := 0; i < totalLen/16; i++ {
		if i > 0 {
			d[i] = append(d[i-1], password...)
		}
		for j := 1; j <= 2048; j++ {
			h := md5.New()
			h.Write(d[i])
			d[i] = h.Sum(nil)
		}
	}
	pbkey := d[0]
	for _, i := range d[1:] {
		pbkey = append(pbkey, i...)
	}
	key := pbkey[0:keylen]
	iv := pbkey[keylen:totalLen]

	return key, iv
}
