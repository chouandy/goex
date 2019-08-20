package ethex

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromWei(t *testing.T) {
	// Set test cases
	testCases := []struct {
		v        *big.Int
		d        int
		expected string
	}{
		{
			v:        big.NewInt(1),
			d:        18,
			expected: "0.000000000000000001",
		},
		{
			v:        big.NewInt(1000000000000000000),
			d:        18,
			expected: "1",
		},
		{
			v:        big.NewInt(1000000),
			d:        6,
			expected: "1",
		},
		{
			v:        new(big.Int).,
			d:        6,
			expected: "1",
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			amount := FromWei(testCase.v, testCase.d)
			assert.Equal(t, testCase.expected, amount.String())
		})
	}
}
