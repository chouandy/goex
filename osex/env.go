package osex

import (
	"os"
	"strconv"
)

// Getenv get env with fallback
func Getenv(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

// GetenvParseInt get env parse int
func GetenvParseInt(key string) int {
	value := os.Getenv(key)
	intValue, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return 0
	}
	return int(intValue)
}
