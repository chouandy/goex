package osex

import (
	"os"
	"testing"
)

// 104 ns/op
func BenchmarkGetenv(b *testing.B) {
	os.Setenv("REGION", "us-east-1")
	for i := 0; i < b.N; i++ {
		os.Getenv("REGION")
	}
}
