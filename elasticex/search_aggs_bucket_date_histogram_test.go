package elasticex

import (
	"bytes"
	json "encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDateHistogramAggregation(t *testing.T) {
	// Set test cases
	testCases := []struct {
		agg      DateHistogramAggregation
		expected struct {
			name   string
			source string
		}
	}{
		{
			agg: DateHistogramAggregation{
				Field:    "timestamp",
				Interval: "1h",
			},
			expected: struct {
				name   string
				source string
			}{
				name: "timestamp",
				source: `{
					"date_histogram": {
						"field": "timestamp",
						"format": "date_time_no_millis",
						"interval": "1h"
					}
				}`,
			},
		},
		{
			agg: DateHistogramAggregation{
				Field:             "timestamp",
				Interval:          "1h",
				ExtendedBoundsMin: 1526017600000,
			},
			expected: struct {
				name   string
				source string
			}{
				name: "timestamp",
				source: `{
					"date_histogram": {
						"field": "timestamp",
						"format": "date_time_no_millis",
						"interval": "1h"
					}
				}`,
			},
		},
		{
			agg: DateHistogramAggregation{
				Field:             "timestamp",
				Interval:          "1h",
				ExtendedBoundsMin: 1526017600000,
				ExtendedBoundsMax: 1526050000000,
			},
			expected: struct {
				name   string
				source string
			}{
				name: "timestamp",
				source: `{
					"date_histogram": {
						"extended_bounds": {
							"max": 1526050000000,
							"min": 1526017600000
						},
						"field": "timestamp",
						"format": "date_time_no_millis",
						"interval": "1h"
					}
				}`,
			},
		},
		{
			agg: DateHistogramAggregation{
				Field:             "timestamp",
				Interval:          "1h",
				ExtendedBoundsMin: 1526017600000,
				ExtendedBoundsMax: 1526050000000,
				SumAggregation: &SumAggregation{
					Field: "aggregate-count",
				},
			},
			expected: struct {
				name   string
				source string
			}{
				name: "timestamp",
				source: `{
					"aggregations": {
						"aggregate-count": {
							"sum": {
								"field": "aggregate-count"
							}
						}
					},
					"date_histogram": {
						"extended_bounds": {
							"max": 1526050000000,
							"min": 1526017600000
						},
						"field": "timestamp",
						"format": "date_time_no_millis",
						"interval": "1h"
					}
				}`,
			},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			name, agg := testCase.agg.Build()
			assert.Equal(t, testCase.expected.name, name)
			source, err := agg.Source()
			assert.Nil(t, err)
			data, err := jsonex.Marshal(source)
			buffer := new(bytes.Buffer)
			err = json.Compact(buffer, []byte(testCase.expected.source))
			assert.Nil(t, err)
			assert.Equal(t, string(buffer.Bytes()), string(data))
		})
	}
}
