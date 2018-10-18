package compressex

import (
	"bytes"
	"compress/gzip"
)

// GZip gzip
func GZip(b []byte) *bytes.Buffer {
	buffer := new(bytes.Buffer)
	w := gzip.NewWriter(buffer)
	w.Write(b)
	w.Close()

	return buffer
}

// GZipString gzip string
func GZipString(str string) *bytes.Buffer {
	buffer := new(bytes.Buffer)
	w := gzip.NewWriter(buffer)
	w.Write([]byte(str))
	w.Close()

	return buffer
}
