package ginex

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/chouandy/goex/httpex"
	"github.com/gin-gonic/gin"
)

// NotLoggedPaths not logged paths
var NotLoggedPaths map[string]struct{}

// Log log struct
type Log struct {
	Timestamp             string            `json:"timestamp"`
	Level                 string            `json:"level"`
	Status                int               `json:"status"`
	Method                string            `json:"method"`
	Path                  string            `json:"path"`
	Latency               string            `json:"latency"`
	QueryStringParameters map[string]string `json:"query_string_parameters,omitempty"`
	PathParameters        map[string]string `json:"path_parameters,omitempty"`
	Body                  string            `json:"body,omitempty"`
	ClientIP              string            `json:"client_ip"`
	Location              string            `json:"location,omitempty"`
}

// Logger logger
func Logger() gin.HandlerFunc {
	return LoggerWithWriter(gin.DefaultWriter)
}

// LoggerWithNotLogged logger with not logged
func LoggerWithNotLogged(paths ...string) gin.HandlerFunc {
	// Set not logged
	if length := len(paths); length > 0 {
		NotLoggedPaths = make(map[string]struct{}, length)
		for _, path := range paths {
			NotLoggedPaths[path] = struct{}{}
		}
	}

	return LoggerWithWriter(gin.DefaultWriter)
}

// LoggerWithWriter logger with writer
func LoggerWithWriter(out io.Writer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Not logged
		if _, ok := NotLoggedPaths[c.Request.URL.Path]; ok {
			return
		}

		// New log
		log := &Log{
			Timestamp:             start.Format(time.RFC3339),
			Level:                 httpex.GetLogLevel(c.Writer.Status()),
			Status:                c.Writer.Status(),
			Method:                c.Request.Method,
			Path:                  c.Request.URL.Path,
			Latency:               fmt.Sprintf("%v", time.Now().UTC().Sub(start)),
			QueryStringParameters: make(map[string]string),
			PathParameters:        make(map[string]string),
			ClientIP:              c.ClientIP(),
		}
		// Set query string parameters
		for key := range c.Request.URL.Query() {
			log.QueryStringParameters[key] = c.Request.URL.Query().Get(key)
		}
		// Set path parameters
		for _, param := range c.Params {
			log.PathParameters[param.Key] = param.Value
		}
		// Set request body
		if body, err := c.GetRawData(); err == nil {
			log.Body = string(body)
		}
		// Set location
		if log.Status == http.StatusFound {
			log.Location = c.Writer.Header().Get("Location")
		}
		// Marshal log
		data, err := jsonex.Marshal(log)
		// Print log
		if err == nil {
			fmt.Fprintln(out, string(data))
		}
	}
}
