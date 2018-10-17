package elasticex

import (
	"bytes"
	json "encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumAggregation(t *testing.T) {
	// Set test cases
	testCases := []struct {
		field    string
		expected string
	}{
		{
			field: "aggregate-count",
			expected: `{
				"sum": {
					"field": "aggregate-count"
				}
			}`,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			agg := SumAggregation{
				Field: testCase.field,
			}
			source, err := agg.Build().Source()
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
