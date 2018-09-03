package apigwex

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// FoundResponse redirect response
func FoundResponse(u string) (events.APIGatewayProxyResponse, error) {
	Logger.SetStatus(http.StatusFound)
	Logger.Log()
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusFound,
		Headers: map[string]string{
			"Location": u,
		},
	}, nil
}
