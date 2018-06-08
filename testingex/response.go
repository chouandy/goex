package testingex

import (
	"bytes"
	"encoding/json"
)

// Response api testing request struct
type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       string
}

// CompactResponseBody compact response body
func (a *Response) CompactResponseBody() []byte {
	buffer := new(bytes.Buffer)
	if err := json.Compact(buffer, []byte(a.Body)); err != nil {
		return []byte(a.Body)
	}

	return buffer.Bytes()
}
