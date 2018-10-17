package elasticex

import (
	"bytes"
	json "encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeQuery(t *testing.T) {
	// Set test cases
	testCases := []struct {
		query    RangeQuery
		expected string
	}{
		{
			query: RangeQuery{
				Name:   "timestamp",
				Gte:    1527658861,
				Lte:    1527702061,
				Format: "epoch_second",
			},
			expected: `{
				"range": {
					"timestamp": {
						"format": "epoch_second",
						"from": 1527658861,
						"include_lower": true,
						"include_upper": true,
						"to": 1527702061
					}
				}
			}`,
		},
		{
			query: RangeQuery{
				Name:   "timestamp",
				Gte:    1527658861913,
				Lte:    1527702061913,
				Format: "epoch_millis",
			},
			expected: `{
				"range": {
					"timestamp": {
						"format": "epoch_millis",
						"from": 1527658861913,
						"include_lower": true,
						"include_upper": true,
						"to": 1527702061913
					}
				}
			}`,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			source, err := testCase.query.Build().Source()
			assert.Nil(t, err)
			data, err := jsonex.Marshal(source)
			assert.Nil(t, err)
			buffer := new(bytes.Buffer)
			err = json.Compact(buffer, []byte(testCase.expected))
			assert.Nil(t, err)
			assert.Equal(t, buffer.String(), string(data))
		})
	}
}
