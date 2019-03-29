package ginex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/chouandy/goex/httpex"
	"github.com/gin-gonic/gin"
)

// NotLoggedPaths not logged paths
var NotLoggedPaths map[string]struct{}

// Log log struct
type Log struct {
	Timestamp             string            `json:"timestamp"`
	RequestID             string            `json:"request_id,omitempty"`
	Level                 string            `json:"level"`
	Status                int               `json:"status"`
	Method                string            `json:"method"`
	Path                  string            `json:"path"`
	Latency               string            `json:"latency"`
	QueryStringParameters url.Values        `json:"query_string_parameters,omitempty"`
	PathParameters        map[string]string `json:"path_parameters,omitempty"`
	Body                  string            `json:"body,omitempty"`
	Error                 json.RawMessage   `json:"error,omitempty"`
	ClientIP              string            `json:"client_ip"`
	Location              string            `json:"location,omitempty"`
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
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

		// Replace context writer
		writer := &bodyLogWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = writer

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
			QueryStringParameters: make(url.Values),
			PathParameters:        make(map[string]string),
			ClientIP:              c.ClientIP(),
		}
		// Set request id
		if requestID, exists := c.Get("request_id"); exists {
			log.RequestID = requestID.(string)
		}
		// Set query string parameters
		log.QueryStringParameters = c.Request.URL.Query()
		// Set path parameters
		for _, param := range c.Params {
			log.PathParameters[param.Key] = param.Value
		}
		// Set request body
		if body, err := c.GetRawData(); err == nil {
			log.Body = string(body)
		}
		// Set error
		if log.Status >= http.StatusBadRequest {
			log.Error = json.RawMessage(writer.body.Bytes())
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
