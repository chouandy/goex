package ginex

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// LatencyHandler latency handler
func LatencyHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
