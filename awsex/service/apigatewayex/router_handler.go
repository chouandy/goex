package apigatewayex

// NotFoundHandler not found handler
func NotFoundHandler(ctx *Context) {
	ctx.NotFoundResponse()
}
