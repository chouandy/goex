package ginex

import "github.com/gin-gonic/gin"

// NotFoundResponseBody not found response body
var NotFoundResponseBody = gin.H{
	"message": "Not Found",
}

// InternalServerErrorResponseBody internal server error response body
var InternalServerErrorResponseBody = gin.H{
	"message": "Internal Server Error",
}
