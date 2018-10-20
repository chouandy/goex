package ginex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 959 ns/op
func BenchmarkEmptyStringValidation1(b *testing.B) {
	var str string
	for i := 0; i < b.N; i++ {
		assert.False(b, str != "")
	}
}

// 956 ns/op
func BenchmarkEmptyStringValidation2(b *testing.B) {
	var str string
	for i := 0; i < b.N; i++ {
		assert.False(b, len(str) > 0)
	}
}
