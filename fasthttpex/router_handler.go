package fasthttpex

import "github.com/valyala/fasthttp"

// NotFoundHandler not found handler
func NotFoundHandler(ctx *fasthttp.RequestCtx) {
	ctx.NotFound()
}
