package stringsex

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRand(t *testing.T) {
	// Set test cases
	testCases := []struct {
		n        int
		expected int
	}{
		{
			n:        7,
			expected: 7,
		},
		{
			n:        10,
			expected: 10,
		},
		{
			n:        15,
			expected: 15,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			s := Rand(Alphanumeric, testCase.n)
			assert.Equal(t, testCase.expected, len(s))
		})
	}
}

func TestRandPassword(t *testing.T) {
	// Set test cases
	testCases := []struct {
		n int
	}{
		{
			n: 8,
		},
		{
			n: 12,
		},
		{
			n: 16,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			password := RandPassword(testCase.n, true, true, true, true)
			assert.True(t, CheckPassword(password, true, true, true, true))
		})
	}
}

func TestCheckPassword(t *testing.T) {
	// Set test cases
	testCases := []struct {
		password string
		number   bool
		lower    bool
		upper    bool
		special  bool
		expected bool
	}{
		{
			password: "0Aa~",
			number:   true,
			lower:    true,
			upper:    true,
			special:  true,
			expected: true,
		},
		{
			password: "Aa~",
			number:   true,
			lower:    true,
			upper:    true,
			special:  true,
			expected: false,
		},
		{
			password: "Aa~",
			number:   false,
			lower:    true,
			upper:    true,
			special:  true,
			expected: true,
		},
		{
			password: "0a~",
			number:   true,
			lower:    true,
			upper:    true,
			special:  true,
			expected: false,
		},
		{
			password: "0a~",
			number:   true,
			lower:    false,
			upper:    true,
			special:  true,
			expected: true,
		},
		{
			password: "0A~",
			number:   true,
			lower:    true,
			upper:    true,
			special:  true,
			expected: false,
		},
		{
			password: "0A~",
			number:   true,
			lower:    true,
			upper:    false,
			special:  true,
			expected: true,
		},
		{
			password: "0Aa",
			number:   true,
			lower:    true,
			upper:    true,
			special:  true,
			expected: false,
		},
		{
			password: "0Aa",
			number:   true,
			lower:    true,
			upper:    true,
			special:  false,
			expected: true,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.expected, CheckPassword(
				testCase.password,
				testCase.number,
				testCase.upper,
				testCase.lower,
				testCase.special,
			))
		})
	}
}
