package stringsex

import (
	"fmt"
	"strings"
)

// FormatMACAddress format mac address
func FormatMACAddress(s string) string {
	if !strings.ContainsAny(s, ":") && len(s) == 12 {
		s = fmt.Sprintf("%02s:%02s:%02s:%02s:%02s:%02s",
			s[:2], s[2:4], s[4:6], s[6:8], s[8:10], s[10:],
		)
	}

	return s
}
