package apigwex

import (
	"time"

	"github.com/aws/aws-lambda-go/events"
)

// Context context struct
type Context struct {
	Request  events.APIGatewayProxyRequest
	Response events.APIGatewayProxyResponse
	Logger   Logger
}

// NewContext new context
func NewContext(request events.APIGatewayProxyRequest) (ctx *Context) {
	ctx = &Context{
		Request: request,
		Logger: Logger{
			RequestTime: time.Now().UTC(),
			RequestID:   request.RequestContext.RequestID,
			Method:      request.HTTPMethod,
			Path:        request.Path,
			QueryStringParameters: request.QueryStringParameters,
			PathParameters:        request.PathParameters,
			Body:                  request.Body,
			Identity: &Identity{
				AccountID: request.RequestContext.Identity.AccountID,
				SourceIP:  request.RequestContext.Identity.SourceIP,
				UserArn:   request.RequestContext.Identity.UserArn,
				UserAgent: request.RequestContext.Identity.UserAgent,
			},
		},
	}

	return
}
