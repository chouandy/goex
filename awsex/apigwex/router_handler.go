package apigwex

// HandlerFunc handler func
type HandlerFunc func(ctx *Context)

// NotFoundHandler not found handler
func NotFoundHandler(ctx *Context) {
	ctx.NotFoundResponse()
}
