package apigwex

import (
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	// Set routers
	NewRouter()
	Router.Add("GET", "hello", func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		return OKResponse("world")
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
			expected: `{"message":"404 Not Found"}`,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			request := events.APIGatewayProxyRequest{}
			NewLogger(request)
			defer Logger.Log()
			handler := Router.Get(testCase.method, testCase.path)
			resp, err := handler(request)
			assert.Nil(t, err)
			assert.Equal(t, testCase.expected, resp.Body)
		})
	}
}
