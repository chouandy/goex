package cryptoex

import "testing"

func TestPadUnpad(t *testing.T) {
	testdata := []byte("This is the test data")
	paddata, _ := Pad(testdata, 32)
	unpaddata, _ := Unpad(paddata, 32)
	if string(testdata) != string(unpaddata) {
		t.Fatal("Padding/Unpadding failed! Error in Pad/Unpad functions")
	}
}
