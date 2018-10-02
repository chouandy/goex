package fasthttpex

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func TestRouter(t *testing.T) {
	// Set routers
	router := NewRouter()
	router.Add("hello", func(ctx *fasthttp.RequestCtx) {
		ctx.SetBodyString("world")
	})

	// Set test cases
	testCases := []struct {
		path     string
		expected string
	}{
		{
			path:     "hello",
			expected: "world",
		},
		{
			path:     "not_found",
			expected: "404 Page not found",
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			ctx := &fasthttp.RequestCtx{}
			router.Get(testCase.path)(ctx)
			assert.Equal(t, testCase.expected, string(ctx.Response.Body()))
		})
	}
}
