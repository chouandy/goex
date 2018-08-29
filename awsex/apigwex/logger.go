package apigwex

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/chouandy/goex/httpex"
)

// Logger logger
var Logger *APIGWLogger

// NewLogger new Logger instance
func NewLogger(request events.APIGatewayProxyRequest) {
	Logger = &APIGWLogger{
		RequestTime: time.Now().UTC(),
		RequestID:   request.RequestContext.RequestID,
		Method:      request.HTTPMethod,
		Path:        request.Path,
		QueryStringParameters: request.QueryStringParameters,
		PathParameters:        request.PathParameters,
		Body:                  request.Body,
		Identity: &Identity{
			AccountID: request.RequestContext.Identity.AccountID,
			SourceIP:  request.RequestContext.Identity.SourceIP,
			UserArn:   request.RequestContext.Identity.UserArn,
			UserAgent: request.RequestContext.Identity.UserAgent,
		},
	}
}

// APIGWLogger apigw logger struct
type APIGWLogger struct {
	RequestTime           time.Time         `json:"-"`
	Timestamp             string            `json:"timestamp"`
	RequestID             string            `json:"request_id"`
	Level                 string            `json:"level"`
	Status                int               `json:"status"`
	Method                string            `json:"method"`
	Path                  string            `json:"path"`
	Latency               string            `json:"latency"`
	Identity              *Identity         `json:"identity,omitempty"`
	QueryStringParameters map[string]string `json:"query_string_parameters,omitempty"`
	PathParameters        map[string]string `json:"path_parameters,omitempty"`
	Body                  string            `json:"body,omitempty"`
	Error                 json.RawMessage   `json:"error,omitempty"`
	Metadata              json.RawMessage   `json:"metadata,omitempty"`
}

// Identity identity struct
type Identity struct {
	AccountID string `json:"account_id"`
	SourceIP  string `json:"source_ip"`
	UserArn   string `json:"user_arn"`
	UserAgent string `json:"user_agent"`
}

// SetStatus set status
func (l *APIGWLogger) SetStatus(status int) {
	l.Status = status
	l.Level = httpex.GetLogLevel(status)
}

// Log print logger
func (l *APIGWLogger) Log() {
	// Log timestamp and latency
	end := time.Now().UTC()
	l.Timestamp = end.Format(time.RFC3339)
	l.Latency = fmt.Sprintf("%v", end.Sub(l.RequestTime))
	// Log to json format
	if data, err := l.MarshalJSON(); err == nil {
		fmt.Println(string(data))
	}
}
