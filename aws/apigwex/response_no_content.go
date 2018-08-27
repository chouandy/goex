package apigwex

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// NoContentResponse no content response
func NoContentResponse() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusNoContent,
	}, nil
}