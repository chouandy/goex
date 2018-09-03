package apigwex

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/chouandy/goex/httpex"
)

var notFoundBody = `{"message":"404 Not Found"}`

// NotFoundResponse not found response
func NotFoundResponse() (events.APIGatewayProxyResponse, error) {
	Logger.SetStatus(http.StatusNotFound)
	Logger.Log()
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusNotFound,
		Headers: map[string]string{
			"Content-Type": httpex.JSONContentType,
		},
		Body: notFoundBody,
	}, nil
}
