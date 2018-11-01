package randex

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateFriendlyToken(t *testing.T) {
	// Set test cases
	testCases := []struct {
		length   int
		expected int
	}{
		{
			length:   15,
			expected: 15,
		},
		{
			length:   20,
			expected: 20,
		},
		{
			length:   32,
			expected: 32,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			token := GenerateFriendlyToken(testCase.length)
			assert.Equal(t, testCase.expected, len(token))
		})
	}
}
