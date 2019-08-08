package cloudwatcheventsex

import "errors"

// NotFoundHandler not found handler
func NotFoundHandler(ctx *Context) {
	ctx.Exception = errors.New("event not found")
}
