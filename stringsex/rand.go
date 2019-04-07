package stringsex

import (
	"math/rand"
	"time"
	"unicode"
)

// Alphanumeric alphanumeric characters
var Alphanumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Rand random generate string
func Rand(characters string, n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, n)
	for i := range b {
		b[i] = characters[r.Intn(len(characters))]
	}

	return string(b)
}

// PasswordCharacters password characters
var PasswordCharacters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"

// RandPassword generate randpassword
func RandPassword(n int, number, upper, lower, special bool) string {
	password := ""
	for !CheckPassword(password, n, number, upper, lower, special) {
		password = Rand(PasswordCharacters, n)
	}
	return password
}

// CheckPassword check password
func CheckPassword(password string, n int, number, lower, upper, special bool) bool {
	// Check length
	if len(password) < n {
		return false
	}

	// Check Number, Lowercase character, Uppercase character, Special character
	var isNumber, isLower, isUpper, isSpecial bool

	// Ignore number check
	if !number {
		isNumber = true
	}
	// Ignore lower check
	if !lower {
		isLower = true
	}
	// Ignore upper check
	if !upper {
		isUpper = true
	}
	// Ignore special check
	if !special {
		isSpecial = true
	}

	// Check password
	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			isNumber = true
		case unicode.IsLower(c):
			isLower = true
		case unicode.IsUpper(c):
			isUpper = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			isSpecial = true
		}
	}

	return isNumber && isLower && isUpper && isSpecial
}
