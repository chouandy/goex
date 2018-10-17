package stringsex

import (
	"encoding/json"
	"strings"
)

// ParseMultipleLineToStringArray parse multiple line data to string array, and remove empty line
func ParseMultipleLineToStringArray(data string) []string {
	// Split data with \n
	raw := strings.Split(data, "\n")
	// Remove empty line
	var new []string
	for _, str := range raw {
		if str != "" {
			new = append(new, str)
		}
	}

	return new
}

// ConvertStringArrayToJSONArray convert string array to json array
func ConvertStringArrayToJSONArray(data []string) []map[string]interface{} {
	var new []map[string]interface{}
	for _, str := range data {
		var v map[string]interface{}
		if err := json.Unmarshal([]byte(str), &v); err == nil {
			new = append(new, v)
		}
	}

	return new
}
