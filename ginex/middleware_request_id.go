package ginex

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// RequestIDMiddleware request id middleware
func RequestIDMiddleware(c *gin.Context) {
	// Generate request id
	requestID := uuid.Must(uuid.NewV4()).String()
	// Set request id
	c.Set("request_id", requestID)
	// Set request id to response
	c.Writer.Header().Set("X-Request-Id", requestID)

	c.Next()
}
