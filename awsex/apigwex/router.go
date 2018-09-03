package apigwex

import (
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

// HandlerFunc handler func
type HandlerFunc func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

// Router router
var Router *APIGWRouter

// NewRouter new router
func NewRouter() {
	Router = &APIGWRouter{
		Routes: map[string]HandlerFunc{},
	}
}

// APIGWRouter router
type APIGWRouter struct {
	Routes map[string]HandlerFunc
}

// Add add route by method, path
func (r APIGWRouter) Add(method, path string, hander HandlerFunc) {
	key := strings.ToLower(method + ":" + path)
	r.Routes[key] = hander
}

// Get get route by method, path
func (r APIGWRouter) Get(method, path string) HandlerFunc {
	key := strings.ToLower(method + ":" + path)
	handler, ok := r.Routes[key]
	if !ok {
		return NotFoundHandler
	}
	return handler
}
