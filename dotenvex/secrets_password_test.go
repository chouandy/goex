package dotenvex

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSecretsPassword(t *testing.T) {
	// Set test cases
	testCases := []struct {
		expected string
	}{
		{
			expected: "string",
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			password := GetSecretsPassword()
			fmt.Println(password)
			assert.IsType(t, testCase.expected, password)
		})
	}
}
