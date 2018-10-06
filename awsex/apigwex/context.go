package apigwex

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/chouandy/goex/httpex"
)

// DEBUG debug
var DEBUG = os.Getenv("DEBUG")

// Context context struct
type Context struct {
	RequestTime time.Time
	Request     events.APIGatewayProxyRequest
	Response    events.APIGatewayProxyResponse
	Logger      Logger

	// Extra
	Region string
}

// NewContext new context
func NewContext(request events.APIGatewayProxyRequest) (ctx *Context) {
	ctx = &Context{
		RequestTime: time.Now().UTC(),
		Request:     request,
		Logger: Logger{
			RequestID:             request.RequestContext.RequestID,
			Method:                request.HTTPMethod,
			Path:                  request.Path,
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

// Log log
func (c *Context) Log() {
	// Set logger field
	c.Logger.Timestamp = c.RequestTime.Format(time.RFC3339)
	c.Logger.Latency = fmt.Sprintf("%v", time.Now().UTC().Sub(c.RequestTime))
	c.Logger.Status = c.Response.StatusCode
	c.Logger.Level = httpex.GetLogLevel(c.Logger.Status)
	// Set location
	if c.Response.StatusCode == http.StatusFound {
		c.Logger.Location = c.Response.Headers["Location"]
	}
	// Set error
	if c.Response.StatusCode > http.StatusBadRequest {
		c.Logger.Error = json.RawMessage(c.Response.Body)
	}
	// Set response body for debug
	if DEBUG == "1" && c.Response.StatusCode == http.StatusOK {
		c.Logger.Response = json.RawMessage(c.Response.Body)
	}
	// Log
	go c.Logger.Log()
}

// WrapRoute wrap route
func (c *Context) WrapRoute(route *Route) {
	// Logging
	if route.Logging {
		defer c.Log()
	}
	// Middlewares
	for _, middleware := range route.Middlewares {
		if err := middleware(c); err != nil {
			c.ErrorResponse(err.(httpex.Error))
			return
		}
	}
	// Handler
	route.Handler(c)
}
