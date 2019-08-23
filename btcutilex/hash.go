package btcutilex

import (
	"crypto/sha256"
	"hash"

	"golang.org/x/crypto/ripemd160"
)

// CalcHash calculate the hash of hasher over buf.
func CalcHash(buf []byte, hasher hash.Hash) []byte {
	hasher.Write(buf)
	return hasher.Sum(nil)
}

// Sha256 calculates the hash sha256(b).
func Sha256(buf []byte) []byte {
	return CalcHash(buf, sha256.New())
}

// Ripemd160 calculates the hash ripemd160(b).
func Ripemd160(buf []byte) []byte {
	return CalcHash(buf, ripemd160.New())
}
