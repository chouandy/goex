package osex

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 104 ns/op
func BenchmarkGetenv(b *testing.B) {
	os.Setenv("REGION", "us-east-1")
	for i := 0; i < b.N; i++ {
		os.Getenv("REGION")
	}
}

func TestGetenvParseInt(t *testing.T) {
	// Set test cases
	testCases := []struct {
		env      string
		expected int
	}{
		{
			env:      "1",
			expected: 1,
		},
		{
			env:      "10",
			expected: 10,
		},
		{
			env:      "100",
			expected: 100,
		},
		{
			env:      "aaa",
			expected: 0,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			os.Setenv("PARSE_INT", testCase.env)
			intValue := GetenvParseInt("PARSE_INT")
			assert.Equal(t, testCase.expected, intValue)
		})
	}
}
