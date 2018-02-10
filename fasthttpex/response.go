package fasthttpex

import (
	"fmt"

	"github.com/chouandy/goex/httpex"
	"github.com/valyala/fasthttp"
)

// JSONErrorMessageFormat json error message format
var JSONErrorMessageFormat = `{"message":"%s"}`

// XMLErrorMessageFormat xml error message format
// <?xml version="1.0" encoding="UTF-8"?>
// <service name="%s">
// 	<entry name="result">
// 		<error>1</error>
// 		<msg>%s</msg>
// 	</entry>
// </service>
var XMLErrorMessageFormat = `<?xml version="1.0" encoding="UTF-8"?><service name="%s"><entry name="result"><error>1</error><msg>%s</msg></entry></service>`

// JSONError JSON format error response
func JSONError(ctx *fasthttp.RequestCtx, statusCode int, msg string) {
	ctx.Response.Reset()
	ctx.SetStatusCode(statusCode)
	ctx.SetContentTypeBytes(httpex.JSONContentTypeBytes)
	ctx.SetBodyString(fmt.Sprintf(JSONErrorMessageFormat, msg))
}

// XMLError XML format error response
func XMLError(ctx *fasthttp.RequestCtx, statusCode int, service string, msg string) {
	ctx.Response.Reset()
	ctx.SetStatusCode(statusCode)
	ctx.SetContentTypeBytes(httpex.XMLContentTypeBytes)
	ctx.SetBodyString(fmt.Sprintf(XMLErrorMessageFormat, service, msg))
}
