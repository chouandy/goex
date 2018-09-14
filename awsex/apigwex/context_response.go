package apigwex

import (
	json "encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/chouandy/goex/httpex"
)

// OKResponse ok response
func (c *Context) OKResponse(body string) {
	// Build response
	c.Response = events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": httpex.JSONContentType,
		},
		Body: body,
	}
	// Logger
	c.Logger.SetStatus(http.StatusOK)
	if DEBUG == "1" {
		c.Logger.Response = json.RawMessage(c.Response.Body)
	}
	c.Logger.Log()
}

// NoContentResponse no content response
func (c *Context) NoContentResponse() {
	// Build response
	c.Response = events.APIGatewayProxyResponse{
		StatusCode: http.StatusNoContent,
	}
	// Logger
	c.Logger.SetStatus(http.StatusNoContent)
	c.Logger.Log()
}

// FoundResponse redirect response
func (c *Context) FoundResponse(u string) {
	// Build response
	c.Response = events.APIGatewayProxyResponse{
		StatusCode: http.StatusFound,
		Headers: map[string]string{
			"Location": u,
		},
	}
	// Logger
	c.Logger.SetStatus(http.StatusFound)
	c.Logger.Location = u
	c.Logger.Log()
}

// NotFoundResponse not found response
func (c *Context) NotFoundResponse() {
	// Build response
	c.Response = events.APIGatewayProxyResponse{
		StatusCode: http.StatusNotFound,
		Headers: map[string]string{
			"Content-Type": httpex.JSONContentType,
		},
		Body: `{"message":"404 Not Found"}`,
	}
	// Logger
	c.Logger.SetStatus(http.StatusNotFound)
	c.Logger.Log()
}

// ErrorResponse build error response
func (c *Context) ErrorResponse(err httpex.Error) {
	// Build response
	c.Response = events.APIGatewayProxyResponse{
		StatusCode: err.StatusCode(),
		Headers: map[string]string{
			"Content-Type": httpex.JSONContentType,
		},
		Body: err.Error(),
	}
	// Logger
	c.Logger.SetStatus(err.StatusCode())
	c.Logger.Error = json.RawMessage(err.Error())
	c.Logger.Log()
}
