package httpex

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	// Set test cases
	testCases := []struct {
		err      error
		expected struct {
			statusCode int
			code       string
			message    string
		}
	}{
		{
			err: NewError(400, "400.1", "user_id can't be blank"),
			expected: struct {
				statusCode int
				code       string
				message    string
			}{
				statusCode: 400,
				code:       "400.1",
				message:    "user_id can't be blank",
			},
		},
		{
			err: NewError(500, "500.1", "Failed to load default aws config"),
			expected: struct {
				statusCode int
				code       string
				message    string
			}{
				statusCode: 500,
				code:       "500.1",
				message:    "Failed to load default aws config",
			},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			httperr, ok := testCase.err.(Error)
			assert.Equal(t, true, ok)
			assert.Equal(t, testCase.expected.statusCode, httperr.StatusCode())
			assert.Equal(t, testCase.expected.code, httperr.Code())
			assert.Equal(t, testCase.expected.message, httperr.Message())
			msg := fmt.Sprintf(JSONErrorMessageFormat,
				testCase.expected.code,
				testCase.expected.message,
			)
			assert.Equal(t, msg, httperr.Error())
		})
	}
}
