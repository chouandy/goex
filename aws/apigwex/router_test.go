package apigwex

import (
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	// Set routers
	router := make(Router)
	router.Add("GET", "not_found", func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		return NotFoundResponse()
	})
	// Set test cases
	testCases := []struct {
		method   string
		path     string
		expected string
	}{
		{
			method:   "GET",
			path:     "not_found",
			expected: `{"message":"404 Not Found"}`,
		},
		{
			method:   "GET",
			path:     "invald",
			expected: `{"message":"404 Not Found"}`,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			handler, ok := router.Get(testCase.method, testCase.path)
			if !ok {
				assert.Nil(t, handler)
				return
			}
			request := events.APIGatewayProxyRequest{}
			resp, err := handler(request)
			assert.Nil(t, err)
			assert.Equal(t, testCase.expected, resp.Body)
		})
	}
}
