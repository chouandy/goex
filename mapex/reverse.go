package mapex

// Reverse reverse map key value
func Reverse(m map[string]string) map[string]string {
	reversed := make(map[string]string, len(m))
	for k, v := range m {
		reversed[v] = k
	}

	return reversed
}
