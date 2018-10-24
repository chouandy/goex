package stringsex

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatMACAddress(t *testing.T) {
	// Set test cases
	testCases := []struct {
		macAddress string
		expected   string
	}{
		{
			macAddress: "AABBCCDDEEFF",
			expected:   "AA:BB:CC:DD:EE:FF",
		},
		{
			macAddress: "AA:BBCCDDEEF",
			expected:   "AA:BBCCDDEEF",
		},
		{
			macAddress: "AABBCCDDEEFFGG",
			expected:   "AABBCCDDEEFFGG",
		},
		{
			macAddress: "AABBCCDDEE",
			expected:   "AABBCCDDEE",
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.expected, FormatMACAddress(testCase.macAddress))
		})
	}
}
