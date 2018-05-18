package httpex

import "fmt"

// JSONErrorMessageFormat json message format
var JSONErrorMessageFormat = `{"code":"%s","message":"%s"}`

// Error error struct
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ToJSON to JSON format
func (c *Error) ToJSON() string {
	return fmt.Sprintf(JSONErrorMessageFormat, c.Code, c.Message)
}
