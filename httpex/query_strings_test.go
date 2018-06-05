package httpex

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeQueryStrings(t *testing.T) {
	// Set test cases
	testCases := []struct {
		url      string
		qs       map[string]string
		expected string
	}{
		{
			url: "https://www.google.com",
			qs: map[string]string{
				"a": "b",
				"c": "d",
			},
			expected: `https://www.google.com?a=b&c=d`,
		},
		{
			url: "https://www.google.com?a=b&c=d",
			qs: map[string]string{
				"user name": "andy chou",
				"c":         "f",
			},
			expected: `https://www.google.com?a=b&c=f&user+name=andy+chou`,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			fmt.Println(testCase.url)
			err := MergeQueryStrings(&testCase.url, testCase.qs)
			fmt.Println(testCase.url)
			assert.IsType(t, nil, err)
			assert.Equal(t, testCase.expected, testCase.url)
		})
	}
}
