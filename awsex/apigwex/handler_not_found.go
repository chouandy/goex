package apigwex

import "github.com/aws/aws-lambda-go/events"

// NotFoundHandler not found handler
func NotFoundHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return NotFoundResponse()
}
