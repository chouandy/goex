package sfnex

import "errors"

// NotFoundHandler not found handler
func NotFoundHandler(ctx *Context) {
	ctx.Exception = errors.New("task not found")
}
