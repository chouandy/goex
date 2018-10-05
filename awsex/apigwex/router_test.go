package apigwex

import (
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	// Set routers
	router := NewRouter()
	router.Add("GET", "hello", func(ctx *Context) {
		ctx.OKResponse("world")
	})

	// Set test cases
	testCases := []struct {
		method   string
		path     string
		expected string
	}{
		{
			method:   "GET",
			path:     "hello",
			expected: "world",
		},
		{
			method:   "GET",
			path:     "not_found",
			expected: `{"message":"Not Found"}`,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			request := events.APIGatewayProxyRequest{}
			ctx := NewContext(request)
			router.Get(testCase.method, testCase.path)(ctx)
			assert.Equal(t, testCase.expected, ctx.Response.Body)
		})
	}
}
