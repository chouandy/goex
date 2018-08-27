package apigwex

import (
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

// RequestHandler request handler
type RequestHandler func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

// Router router
type Router map[string]RequestHandler

// Add add route by method, path
func (r Router) Add(method, path string, hander RequestHandler) {
	key := strings.ToLower(method + "/" + path)
	r[key] = hander
}

// Get get route by method, path
func (r Router) Get(method, path string) (RequestHandler, bool) {
	key := strings.ToLower(method + "/" + path)
	handler, ok := r[key]
	return handler, ok
}
