package apigwex

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// NotFoundHandler not found handler
func NotFoundHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	Logger.SetStatus(http.StatusNotFound)
	return NotFoundResponse()
}
