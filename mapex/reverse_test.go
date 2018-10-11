package mapex

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	// Set test cases
	testCases := []struct {
		m        map[string]string
		expected map[string]string
	}{
		{
			m:        map[string]string{"a": "b"},
			expected: map[string]string{"b": "a"},
		},
		{
			m:        map[string]string{"ab": "cd"},
			expected: map[string]string{"cd": "ab"},
		},
		{
			m:        map[string]string{"a": "b", "c": "d"},
			expected: map[string]string{"b": "a", "d": "c"},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			reversed := Reverse(testCase.m)
			assert.Equal(t, testCase.expected, reversed)
		})
	}
}
