package fasthttpex

import "github.com/valyala/fasthttp"

// NewRouter new router
func NewRouter() *Router {
	return &Router{
		Routes: map[string]fasthttp.RequestHandler{},
	}
}

// Router router
type Router struct {
	Routes map[string]fasthttp.RequestHandler
}

// Add add route by path
func (r *Router) Add(path string, hander fasthttp.RequestHandler) {
	r.Routes[path] = hander
}

// Get get route by path
func (r *Router) Get(path string) fasthttp.RequestHandler {
	handler, ok := r.Routes[path]
	if !ok {
		return NotFoundHandler
	}
	return handler
}
