package httpex

// GetLogLevel get log level by status code
func GetLogLevel(code int) string {
	switch {
	case code >= 200 && code < 300:
		return "INFO"
	case code >= 300 && code < 400:
		return "INFO"
	case code >= 400 && code < 500:
		return "ERROR"
	default:
		return "FATAL"
	}
}
