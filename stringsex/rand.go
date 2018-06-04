package stringsex

import "math/rand"

// Alphanumeric alphanumeric characters
var Alphanumeric = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// Rand random generate string
func Rand(runes []byte, n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}

	return string(b)
}
