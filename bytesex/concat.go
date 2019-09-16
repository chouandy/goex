package bytesex

// Concat concat bytes slices
func Concat(slices [][]byte) []byte {
	// Get total length
	var totalLen int
	for i := range slices {
		totalLen += len(slices[i])
	}

	// Append slices
	tmp := make([]byte, totalLen)
	var index int
	for i := range slices {
		index += copy(tmp[index:], slices[i])
	}

	return tmp
}
