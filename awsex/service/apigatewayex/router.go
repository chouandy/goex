package apigatewayex

// HandlerFunc handler func
type HandlerFunc func(ctx *Context)

// MiddlewareFunc middleware func
type MiddlewareFunc func(ctx *Context) error

// Route route
type Route struct {
	Handler     HandlerFunc
	Middlewares []MiddlewareFunc
	Logging     bool
}

// Router router
type Router struct {
	Routes map[string]*Route
}

// NewRouter new router
func NewRouter() *Router {
	return &Router{
		Routes: map[string]*Route{},
	}
}

// Add add route by method, path
func (r *Router) Add(method, path string, route *Route) {
	r.Routes[method+":"+path] = route
}

// Get get route by method, path
func (r *Router) Get(method, path string) *Route {
	route, ok := r.Routes[method+":"+path]
	if !ok {
		return &Route{
			Handler: NotFoundHandler,
		}
	}
	return route
}
