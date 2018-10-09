package cryptoex

import (
	"bytes"
	"errors"
)

// Pad returns the byte array passed as a parameter padded with bytes such that
// the new byte array will be an exact multiple of the expected block size.
// For example, if the expected block size is 8 bytes (e.g. PKCS #5) and that
// the initial byte array is:
// 	[]byte{0x0A, 0x0B, 0x0C, 0x0D}
// the returned array will be:
// 	[]byte{0x0A, 0x0B, 0x0C, 0x0D, 0x04, 0x04, 0x04, 0x04}
// The value of each octet of the padding is the size of the padding. If the
// array passed as a parameter is already an exact multiple of the block size,
// the original array will be padded with a full block.
func Pad(buf []byte, blockSize int) ([]byte, error) {
	bufLen := len(buf)
	padLen := blockSize - (bufLen % blockSize)
	padText := bytes.Repeat([]byte{byte(padLen)}, padLen)
	return append(buf, padText...), nil
}

// Unpad removes the padding of a given byte array, according to the same rules
// as described in the Pad function. For example if the byte array passed as a
// parameter is:
// 	[]byte{0x0A, 0x0B, 0x0C, 0x0D, 0x04, 0x04, 0x04, 0x04}
// the returned array will be:
// 	[]byte{0x0A, 0x0B, 0x0C, 0x0D}
func Unpad(buf []byte, blockSize int) ([]byte, error) {
	bufLen := len(buf)
	if bufLen == 0 {
		return nil, errors.New("cryptgo/padding: invalid padding size")
	}

	pad := buf[bufLen-1]
	padLen := int(pad)
	if padLen > bufLen || padLen > blockSize {
		return nil, errors.New("cryptgo/padding: invalid padding size")
	}

	for _, v := range buf[bufLen-padLen : bufLen-1] {
		if v != pad {
			return nil, errors.New("cryptgo/padding: invalid padding")
		}
	}

	return buf[:bufLen-padLen], nil
}
