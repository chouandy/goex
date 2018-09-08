package apigwex

import (
	"strings"
)

// NewRouter new router
func NewRouter() *Router {
	return &Router{
		Routes: map[string]HandlerFunc{},
	}
}

// Router router
type Router struct {
	Routes map[string]HandlerFunc
}

// Add add route by method, path
func (r Router) Add(method, path string, hander HandlerFunc) {
	key := strings.ToLower(method + ":" + path)
	r.Routes[key] = hander
}

// Get get route by method, path
func (r Router) Get(method, path string) HandlerFunc {
	key := strings.ToLower(method + ":" + path)
	handler, ok := r.Routes[key]
	if !ok {
		return NotFoundHandler
	}
	return handler
}
