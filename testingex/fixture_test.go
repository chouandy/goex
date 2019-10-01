package testingex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFixture(t *testing.T) {
	// Set test cases
	testCases := []struct {
		filename string
		expected string
	}{
		{
			filename: "case1.json",
			expected: `{
				"address": "n1UvNvZjJUsQUbfCvuxSHnSkY9v1RxzApp",
				"txid": "ccfd5fb6c4108b3438c3ed563a9fc35f52543fd2159e9626aab00ce58ec4977a",
				"vout": 1,
				"scriptPubKey": "76a914db001596bffae0a05f2db0ebe65b52a7a86b760988ac",
				"amount": 0.03320652,
				"satoshis": 3320652,
				"height": 1580149,
				"confirmations": 32
			}`,
		},
		{
			filename: "case2.json",
			expected: `{
				"txid": "f0d5f66ef152ab0d5e4e38b448a9e597d79268215bbe98b585692a82a487e3bb",
				"version": 1,
				"locktime": 0,
				"blockhash": "000000009929404e821ff142e3b98cb5ef58535ecea7f316e82dca5a9b0124c4",
				"blockheight": 1579784,
				"confirmations": 398,
				"time": 1569570207,
				"blocktime": 1569570207,
				"valueOut": 0.07226009,
				"size": 615,
				"valueIn": 0.07231301,
				"fees": 0.00005292
			}
			`,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			var bufferTestCase bytes.Buffer
			err := json.Compact(&bufferTestCase, LoadFixture(testCase.filename))
			assert.Nil(t, err)

			var bufferExpected bytes.Buffer
			err = json.Compact(&bufferExpected, []byte(testCase.expected))
			assert.Nil(t, err)

			assert.Equal(t, bufferExpected.Bytes(), bufferTestCase.Bytes())
		})
	}
}
