package httpex

// JSONContentType JSON Content-Type string format
var JSONContentType = "application/json; charset=utf-8"

// JSONContentTypeBytes JSON Content-Type bytes format
var JSONContentTypeBytes = []byte(JSONContentType)

// XMLContentType XML Content-Type string format
var XMLContentType = "application/xml; charset=utf-8"

// XMLContentTypeBytes XML Content-Type bytes format
var XMLContentTypeBytes = []byte(XMLContentType)

// PlainContentType Plain Content-Type string format
var PlainContentType = "text/plain; charset=utf-8"

// PlainContentTypeBytes Plain Content-Type bytes format
var PlainContentTypeBytes = []byte(PlainContentType)

// OKMessage ok message
var OKMessage = `{"message":"ok"}`

// OKMessageBytes ok message bytes format
var OKMessageBytes = []byte(OKMessage)

// NotFoundMessage not found message
var NotFoundMessage = `{"message":"Not Found"}`

// NotFoundMessageBytes not found message bytes format
var NotFoundMessageBytes = []byte(NotFoundMessage)

// InternalServerErrorMessage not found message
var InternalServerErrorMessage = `{"message":"Internal Server Error"}`

// InternalServerErrorMessageBytes not found message bytes format
var InternalServerErrorMessageBytes = []byte(InternalServerErrorMessage)
