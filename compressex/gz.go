package compressex

import (
	"bytes"
	"compress/gzip"
)

// GZipString gzip string
func GZipString(str string) *bytes.Buffer {
	b := new(bytes.Buffer)
	w := gzip.NewWriter(b)
	w.Write([]byte(str))
	w.Close()

	return b
}
