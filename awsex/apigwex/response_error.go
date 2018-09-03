package apigwex

import (
	json "encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/chouandy/goex/httpex"
)

// ErrorResponse error response
func ErrorResponse(err httpex.Error) (events.APIGatewayProxyResponse, error) {
	Logger.SetStatus(err.StatusCode())
	Logger.Error = json.RawMessage(err.Error())
	Logger.Log()
	return events.APIGatewayProxyResponse{
		StatusCode: err.StatusCode(),
		Headers: map[string]string{
			"Content-Type": httpex.JSONContentType,
		},
		Body: err.Error(),
	}, nil
}
