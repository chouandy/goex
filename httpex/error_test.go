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
			json       string
			inline     string
			grpcInline string
		}
	}{
		{
			err: NewError(400, "400.1", "user_id can't be blank"),
			expected: struct {
				statusCode int
				json       string
				inline     string
				grpcInline string
			}{
				statusCode: 400,
				json:       `{"code":"400.1","message":"user_id can't be blank"}`,
				inline:     `code: 400.1, message: user_id can't be blank`,
				grpcInline: `statusCode: 400, code: 400.1, message: user_id can't be blank`,
			},
		},
		{
			err: NewError(500, "500.1", "Failed to load default aws config"),
			expected: struct {
				statusCode int
				json       string
				inline     string
				grpcInline string
			}{
				statusCode: 500,
				json:       `{"code":"500.1","message":"Failed to load default aws config"}`,
				inline:     `code: 500.1, message: Failed to load default aws config`,
				grpcInline: `statusCode: 500, code: 500.1, message: Failed to load default aws config`,
			},
		},
		{
			err: NewError(403, "", "Forbidden"),
			expected: struct {
				statusCode int
				json       string
				inline     string
				grpcInline string
			}{
				statusCode: 403,
				json:       `{"message":"Forbidden"}`,
				inline:     `Forbidden`,
				grpcInline: `statusCode: 403, message: Forbidden`,
			},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			httpErr, ok := testCase.err.(Error)
			assert.Equal(t, true, ok)
			assert.Equal(t, testCase.expected.statusCode, httpErr.StatusCode())
			assert.Equal(t, testCase.expected.json, httpErr.Error())
			assert.Equal(t, testCase.expected.inline, httpErr.ErrorInline())
			assert.Equal(t, testCase.expected.grpcInline, httpErr.GrpcErrorInline())
			httpErr2, ok := ParseGrpcErrorInline(httpErr.GrpcErrorInline())
			assert.True(t, ok)
			assert.Equal(t, httpErr, httpErr2)
		})
	}
}
