package stringsex

// In check str is in strs or not
func In(str string, strs []string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}
	return false
}

// Include check strs include str or not
func Include(strs []string, str string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}
	return false
}

// Exclude check strs exclude str or not
func Exclude(strs []string, str string) bool {
	for _, s := range strs {
		if s == str {
			return false
		}
	}
	return true
}
