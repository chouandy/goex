package cryptoex

import (
	"crypto/sha1"

	"golang.org/x/crypto/pbkdf2"
)

// KeyGenerator key generator struct
type KeyGenerator struct {
	SecretKey  []byte
	Iterations int
}

// GenerateKey generate key
func (g *KeyGenerator) GenerateKey(salt []byte, keyLen int) []byte {
	return pbkdf2.Key(g.SecretKey, salt, g.Iterations, keyLen, sha1.New)
}
