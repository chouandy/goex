package httpex

import (
	"fmt"
)

// JSONMessageFormat json message format
var JSONMessageFormat = `{"code":"%s","message":"%s"}`

// APIResponse response for api json format
func APIResponse(code, message string) string {
	return fmt.Sprintf(JSONMessageFormat, code, message)
}
