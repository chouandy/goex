package httpex

import (
	"fmt"
)

// JSONErrorMessageFormat json message format
var JSONErrorMessageFormat = `{"code":"%s","message":"%s"}`

// JSONErrorResponse response for api json format
func JSONErrorResponse(code, message string) string {
	return fmt.Sprintf(JSONErrorMessageFormat, code, message)
}
