package elasticex

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertIntervalToSeconds(t *testing.T) {
	// Set test cases
	testCases := []struct {
		interval string
		expected int
	}{
		{
			interval: "1h",
			expected: 3600,
		},
		{
			interval: "6h",
			expected: 21600,
		},
		{
			interval: "1d",
			expected: 86400,
		},
		{
			interval: "7d",
			expected: 604800,
		},
		{
			interval: "1M",
			expected: 2592000,
		},
		{
			interval: "6M",
			expected: 15552000,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			seconds, err := ConvertIntervalToSeconds(testCase.interval)
			assert.IsType(t, nil, err)
			assert.Equal(t, testCase.expected, seconds)
		})
	}
}
