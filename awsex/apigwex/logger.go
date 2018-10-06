package apigwex

import (
	"encoding/json"
	"fmt"
)

// Logger apigw logger struct
type Logger struct {
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
	Location              string            `json:"location,omitempty"`
	Metadata              json.RawMessage   `json:"metadata,omitempty"`
	Response              json.RawMessage   `json:"response,omitempty"`
}

// Identity identity struct
type Identity struct {
	AccountID string `json:"account_id"`
	SourceIP  string `json:"source_ip"`
	UserArn   string `json:"user_arn"`
	UserAgent string `json:"user_agent"`
}

// Log log
func (l *Logger) Log() {
	if data, err := jsonex.Marshal(l); err == nil {
		fmt.Println(string(data))
	}
}
