package apigatewayex

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/chouandy/goex/httpex"
)

// OKResponse ok response
func (c *Context) OKResponse(body string) {
	c.Response = events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": httpex.JSONContentType,
		},
		Body: body,
	}
}

// NoContentResponse no content response
func (c *Context) NoContentResponse() {
	c.Response = events.APIGatewayProxyResponse{
		StatusCode: http.StatusNoContent,
	}
}

// FoundResponse redirect response
func (c *Context) FoundResponse(u string) {
	c.Response = events.APIGatewayProxyResponse{
		StatusCode: http.StatusFound,
		Headers: map[string]string{
			"Location": u,
		},
	}
}

// NotFoundResponse not found response
func (c *Context) NotFoundResponse() {
	c.Response = events.APIGatewayProxyResponse{
		StatusCode: http.StatusNotFound,
		Headers: map[string]string{
			"Content-Type": httpex.JSONContentType,
		},
		Body: `{"message":"Not Found"}`,
	}
}

// ErrorResponse build error response
func (c *Context) ErrorResponse(err httpex.Error) {
	c.Response = events.APIGatewayProxyResponse{
		StatusCode: err.StatusCode(),
		Headers: map[string]string{
			"Content-Type": httpex.JSONContentType,
		},
		Body: err.Error(),
	}
}
