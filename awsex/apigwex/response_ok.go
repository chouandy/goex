package apigwex

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/chouandy/goex/httpex"
)

// OKResponse ok response
func OKResponse(body string) (events.APIGatewayProxyResponse, error) {
	Logger.SetStatus(http.StatusOK)
	Logger.Log()
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": httpex.JSONContentType,
		},
		Body: body,
	}, nil
}
