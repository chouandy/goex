package compressex

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

// TarFileHandlerFunc tar file handler function
type TarFileHandlerFunc func(tarReader *tar.Reader, header *tar.Header)

// UntarGZFile untar gzip file
func UntarGZFile(f io.Reader, h TarFileHandlerFunc) {
	// 先用 gzip 讀取
	gzf, err := gzip.NewReader(f)
	if err != nil {
		panic(err)
	}
	// 再用 tar 讀取
	tarReader := tar.NewReader(gzf)
	// 取得所有檔案
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		switch header.Typeflag {
		case tar.TypeDir: // = directory
		case tar.TypeReg: // = regular file
			h(tarReader, header)
		default:
		}
	}
}
