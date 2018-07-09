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
			assert.Nil(t, err)
			assert.Equal(t, testCase.expected, testCase.url)
		})
	}
}

func TestAppendQueryStrings(t *testing.T) {
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
			expected: `https://www.google.com?a=b&c=d&c=f&user+name=andy+chou`,
		},
		{
			url: "https://192.168.212.1/gui/#/main/login/sso?checkForGrantCode=1",
			qs: map[string]string{
				"code":  "e0cb7d55bad918599408d673804b6c28a4b0c9a5bf863ac2a3d707e0e7734151",
				"token": "token",
			},
			expected: `https://192.168.212.1/gui/#/main/login/sso?checkForGrantCode=1&code=e0cb7d55bad918599408d673804b6c28a4b0c9a5bf863ac2a3d707e0e7734151&token=token`,
		},
		{
			url: "https://192.168.212.1/gui/",
			qs: map[string]string{
				"code":  "e0cb7d55bad918599408d673804b6c28a4b0c9a5bf863ac2a3d707e0e7734151",
				"token": "token",
			},
			expected: `https://192.168.212.1/gui/?code=e0cb7d55bad918599408d673804b6c28a4b0c9a5bf863ac2a3d707e0e7734151&token=token`,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			AppendQueryStrings(&testCase.url, testCase.qs)
			assert.Equal(t, testCase.expected, testCase.url)
		})
	}
}
