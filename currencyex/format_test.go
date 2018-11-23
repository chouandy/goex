package currencyex

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSymbol(t *testing.T) {
	// Set test cases
	testCases := []struct {
		currency string
		expected string
	}{
		{
			currency: "USD",
			expected: "$",
		},
		{
			currency: "EUR",
			expected: "€",
		},
		{
			currency: "GBP",
			expected: "£",
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			symbol := Symbol(testCase.currency)
			assert.Equal(t, testCase.expected, symbol)
		})
	}
}
