package cryptoex

import (
	"encoding/base64"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyGenerator(t *testing.T) {
	g := &KeyGenerator{
		SecretKey:  []byte("b455266004a3942ef427b6926032e78cf2aca6a1af1261c78aa18d9fd79aecfe3befb254b0d3165f78e653791fea292b818baec530cb10911dde5afbcf7a1b59"),
		Iterations: int(math.Pow(2, 16)),
	}

	// Set test cases
	testCases := []struct {
		salt     []byte
		keyLen   int
		expected string
	}{
		{
			salt:     []byte("Devise reset_password_token"),
			keyLen:   64,
			expected: "_uSfxp9ouVWOqonn-bw2x5X35rNNLG0iStWLv9v3-BoQXZng1Wcz3GQ8g4NznbwRQSv0MZXZCRDFglx2V8lV8Q==",
		},
		{
			salt:     []byte("Devise confirmation_token"),
			keyLen:   64,
			expected: "DMOBWzsXyQaCH0jEE1mgbUILDZuwhJCEPFZ9f7ubQCR02tB-5Me5Xo7C2zI2WedbaozNL5z0BdNrJiksM_MCeQ==",
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			key := g.GenerateKey(testCase.salt, testCase.keyLen)
			encodedKey := base64.URLEncoding.EncodeToString(key)
			assert.Equal(t, testCase.expected, encodedKey)
		})
	}
}
