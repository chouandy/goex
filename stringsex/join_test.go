package stringsex

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoin(t *testing.T) {
	// Set test cases
	testCases := []struct {
		a        []int
		sep      string
		expected string
	}{
		{
			a:        []int{},
			sep:      ",",
			expected: "",
		},
		{
			a:        []int{1},
			sep:      ",",
			expected: "1",
		},
		{
			a:        []int{1, 3, 5, 7, 9},
			sep:      ",",
			expected: "1,3,5,7,9",
		},
		{
			a:        []int{1, 3, 5, 7, 9},
			sep:      "-",
			expected: "1-3-5-7-9",
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			s := JoinInt(testCase.a, testCase.sep)
			assert.Equal(t, testCase.expected, s)
		})
	}
}
