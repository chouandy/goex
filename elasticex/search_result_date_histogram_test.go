package elasticex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSearchDateHistogramBuckets(t *testing.T) {
	// Set test cases
	testCases := []struct {
		searchService SearchService
		searchResult  json.RawMessage
		expected      string
	}{
		{
			searchService: SearchService{
				DateHistogramAggregation: &DateHistogramAggregation{
					Field:             "timestamp",
					Interval:          "1h",
					ExtendedBoundsMin: 1527552000000,
					ExtendedBoundsMax: 1527638400000,
				},
			},
			searchResult: []byte(`{
				"took": 4,
				"timed_out": false,
				"_shards": {
					"total": 21,
					"successful": 21,
					"skipped": 0,
					"failed": 0
				},
				"hits": {
					"total": 3477,
					"max_score": 0,
					"hits": []
				},
				"aggregations": {
					"timestamp": {
						"buckets": [{
								"key_as_string": "2018-05-30T00:00:00.000Z",
								"key": 1527638400000,
								"doc_count": 0
							},
							{
								"key_as_string": "2018-05-30T01:00:00.000Z",
								"key": 1527642000000,
								"doc_count": 0
							},
							{
								"key_as_string": "2018-05-30T02:00:00.000Z",
								"key": 1527645600000,
								"doc_count": 0
							},
							{
								"key_as_string": "2018-05-30T03:00:00.000Z",
								"key": 1527649200000,
								"doc_count": 0
							},
							{
								"key_as_string": "2018-05-30T04:00:00.000Z",
								"key": 1527652800000,
								"doc_count": 0
							},
							{
								"key_as_string": "2018-05-30T05:00:00.000Z",
								"key": 1527656400000,
								"doc_count": 0
							},
							{
								"key_as_string": "2018-05-30T06:00:00.000Z",
								"key": 1527660000000,
								"doc_count": 0
							},
							{
								"key_as_string": "2018-05-30T07:00:00.000Z",
								"key": 1527663600000,
								"doc_count": 0
							},
							{
								"key_as_string": "2018-05-30T08:00:00.000Z",
								"key": 1527667200000,
								"doc_count": 0
							},
							{
								"key_as_string": "2018-05-30T09:00:00.000Z",
								"key": 1527670800000,
								"doc_count": 0
							},
							{
								"key_as_string": "2018-05-30T10:00:00.000Z",
								"key": 1527674400000,
								"doc_count": 477
							},
							{
								"key_as_string": "2018-05-30T11:00:00.000Z",
								"key": 1527678000000,
								"doc_count": 632
							},
							{
								"key_as_string": "2018-05-30T12:00:00.000Z",
								"key": 1527681600000,
								"doc_count": 197
							},
							{
								"key_as_string": "2018-05-30T13:00:00.000Z",
								"key": 1527685200000,
								"doc_count": 209
							},
							{
								"key_as_string": "2018-05-30T14:00:00.000Z",
								"key": 1527688800000,
								"doc_count": 186
							},
							{
								"key_as_string": "2018-05-30T15:00:00.000Z",
								"key": 1527692400000,
								"doc_count": 189
							},
							{
								"key_as_string": "2018-05-30T16:00:00.000Z",
								"key": 1527696000000,
								"doc_count": 204
							},
							{
								"key_as_string": "2018-05-30T17:00:00.000Z",
								"key": 1527699600000,
								"doc_count": 197
							},
							{
								"key_as_string": "2018-05-30T18:00:00.000Z",
								"key": 1527703200000,
								"doc_count": 211
							},
							{
								"key_as_string": "2018-05-30T19:00:00.000Z",
								"key": 1527706800000,
								"doc_count": 182
							},
							{
								"key_as_string": "2018-05-30T20:00:00.000Z",
								"key": 1527710400000,
								"doc_count": 194
							},
							{
								"key_as_string": "2018-05-30T21:00:00.000Z",
								"key": 1527714000000,
								"doc_count": 208
							},
							{
								"key_as_string": "2018-05-30T22:00:00.000Z",
								"key": 1527717600000,
								"doc_count": 191
							},
							{
								"key_as_string": "2018-05-30T23:00:00.000Z",
								"key": 1527721200000,
								"doc_count": 200
							},
							{
								"key_as_string": "2018-05-31T00:00:00.000Z",
								"key": 1527724800000,
								"doc_count": 0
							}
						]
					}
				}
			}`),
			expected: `{
				"buckets": [{
					"timestamp": 1527638400,
					"datetime": "2018-05-30T00:00:00.000Z",
					"value": 0
				}, {
					"timestamp": 1527642000,
					"datetime": "2018-05-30T01:00:00.000Z",
					"value": 0
				}, {
					"timestamp": 1527645600,
					"datetime": "2018-05-30T02:00:00.000Z",
					"value": 0
				}, {
					"timestamp": 1527649200,
					"datetime": "2018-05-30T03:00:00.000Z",
					"value": 0
				}, {
					"timestamp": 1527652800,
					"datetime": "2018-05-30T04:00:00.000Z",
					"value": 0
				}, {
					"timestamp": 1527656400,
					"datetime": "2018-05-30T05:00:00.000Z",
					"value": 0
				}, {
					"timestamp": 1527660000,
					"datetime": "2018-05-30T06:00:00.000Z",
					"value": 0
				}, {
					"timestamp": 1527663600,
					"datetime": "2018-05-30T07:00:00.000Z",
					"value": 0
				}, {
					"timestamp": 1527667200,
					"datetime": "2018-05-30T08:00:00.000Z",
					"value": 0
				}, {
					"timestamp": 1527670800,
					"datetime": "2018-05-30T09:00:00.000Z",
					"value": 0
				}, {
					"timestamp": 1527674400,
					"datetime": "2018-05-30T10:00:00.000Z",
					"value": 477
				}, {
					"timestamp": 1527678000,
					"datetime": "2018-05-30T11:00:00.000Z",
					"value": 632
				}, {
					"timestamp": 1527681600,
					"datetime": "2018-05-30T12:00:00.000Z",
					"value": 197
				}, {
					"timestamp": 1527685200,
					"datetime": "2018-05-30T13:00:00.000Z",
					"value": 209
				}, {
					"timestamp": 1527688800,
					"datetime": "2018-05-30T14:00:00.000Z",
					"value": 186
				}, {
					"timestamp": 1527692400,
					"datetime": "2018-05-30T15:00:00.000Z",
					"value": 189
				}, {
					"timestamp": 1527696000,
					"datetime": "2018-05-30T16:00:00.000Z",
					"value": 204
				}, {
					"timestamp": 1527699600,
					"datetime": "2018-05-30T17:00:00.000Z",
					"value": 197
				}, {
					"timestamp": 1527703200,
					"datetime": "2018-05-30T18:00:00.000Z",
					"value": 211
				}, {
					"timestamp": 1527706800,
					"datetime": "2018-05-30T19:00:00.000Z",
					"value": 182
				}, {
					"timestamp": 1527710400,
					"datetime": "2018-05-30T20:00:00.000Z",
					"value": 194
				}, {
					"timestamp": 1527714000,
					"datetime": "2018-05-30T21:00:00.000Z",
					"value": 208
				}, {
					"timestamp": 1527717600,
					"datetime": "2018-05-30T22:00:00.000Z",
					"value": 191
				}, {
					"timestamp": 1527721200,
					"datetime": "2018-05-30T23:00:00.000Z",
					"value": 200
				}, {
					"timestamp": 1527724800,
					"datetime": "2018-05-31T00:00:00.000Z",
					"value": 0
				}]
			}`,
		},
		{
			searchService: SearchService{
				DateHistogramAggregation: &DateHistogramAggregation{
					Field:             "timestamp",
					Interval:          "1h",
					ExtendedBoundsMin: 1527552000000,
					ExtendedBoundsMax: 1527638400000,
					SumAggregation: &SumAggregation{
						Field: "aggregate-count",
					},
				},
			},
			searchResult: []byte(`{
				"took": 4,
				"timed_out": false,
				"_shards": {
					"total": 21,
					"successful": 21,
					"skipped": 0,
					"failed": 0
				},
				"hits": {
					"total": 3477,
					"max_score": 0,
					"hits": []
				},
				"aggregations": {
					"timestamp": {
						"buckets": [{
								"key_as_string": "2018-05-30T00:00:00.000Z",
								"key": 1527638400000,
								"doc_count": 0,
								"aggregate-count": {
									"value": 0
								}
							},
							{
								"key_as_string": "2018-05-30T01:00:00.000Z",
								"key": 1527642000000,
								"doc_count": 0,
								"aggregate-count": {
									"value": 0
								}
							},
							{
								"key_as_string": "2018-05-30T02:00:00.000Z",
								"key": 1527645600000,
								"doc_count": 0,
								"aggregate-count": {
									"value": 0
								}
							},
							{
								"key_as_string": "2018-05-30T03:00:00.000Z",
								"key": 1527649200000,
								"doc_count": 0,
								"aggregate-count": {
									"value": 0
								}
							},
							{
								"key_as_string": "2018-05-30T04:00:00.000Z",
								"key": 1527652800000,
								"doc_count": 0,
								"aggregate-count": {
									"value": 0
								}
							},
							{
								"key_as_string": "2018-05-30T05:00:00.000Z",
								"key": 1527656400000,
								"doc_count": 0,
								"aggregate-count": {
									"value": 0
								}
							},
							{
								"key_as_string": "2018-05-30T06:00:00.000Z",
								"key": 1527660000000,
								"doc_count": 0,
								"aggregate-count": {
									"value": 0
								}
							},
							{
								"key_as_string": "2018-05-30T07:00:00.000Z",
								"key": 1527663600000,
								"doc_count": 0,
								"aggregate-count": {
									"value": 0
								}
							},
							{
								"key_as_string": "2018-05-30T08:00:00.000Z",
								"key": 1527667200000,
								"doc_count": 0,
								"aggregate-count": {
									"value": 0
								}
							},
							{
								"key_as_string": "2018-05-30T09:00:00.000Z",
								"key": 1527670800000,
								"doc_count": 0,
								"aggregate-count": {
									"value": 0
								}
							},
							{
								"key_as_string": "2018-05-30T10:00:00.000Z",
								"key": 1527674400000,
								"doc_count": 477,
								"aggregate-count": {
									"value": 24885
								}
							},
							{
								"key_as_string": "2018-05-30T11:00:00.000Z",
								"key": 1527678000000,
								"doc_count": 632,
								"aggregate-count": {
									"value": 30851
								}
							},
							{
								"key_as_string": "2018-05-30T12:00:00.000Z",
								"key": 1527681600000,
								"doc_count": 197,
								"aggregate-count": {
									"value": 8991
								}
							},
							{
								"key_as_string": "2018-05-30T13:00:00.000Z",
								"key": 1527685200000,
								"doc_count": 209,
								"aggregate-count": {
									"value": 10268
								}
							},
							{
								"key_as_string": "2018-05-30T14:00:00.000Z",
								"key": 1527688800000,
								"doc_count": 186,
								"aggregate-count": {
									"value": 8761
								}
							},
							{
								"key_as_string": "2018-05-30T15:00:00.000Z",
								"key": 1527692400000,
								"doc_count": 189,
								"aggregate-count": {
									"value": 8251
								}
							},
							{
								"key_as_string": "2018-05-30T16:00:00.000Z",
								"key": 1527696000000,
								"doc_count": 204,
								"aggregate-count": {
									"value": 9642
								}
							},
							{
								"key_as_string": "2018-05-30T17:00:00.000Z",
								"key": 1527699600000,
								"doc_count": 197,
								"aggregate-count": {
									"value": 9057
								}
							},
							{
								"key_as_string": "2018-05-30T18:00:00.000Z",
								"key": 1527703200000,
								"doc_count": 211,
								"aggregate-count": {
									"value": 10673
								}
							},
							{
								"key_as_string": "2018-05-30T19:00:00.000Z",
								"key": 1527706800000,
								"doc_count": 182,
								"aggregate-count": {
									"value": 8816
								}
							},
							{
								"key_as_string": "2018-05-30T20:00:00.000Z",
								"key": 1527710400000,
								"doc_count": 194,
								"aggregate-count": {
									"value": 10586
								}
							},
							{
								"key_as_string": "2018-05-30T21:00:00.000Z",
								"key": 1527714000000,
								"doc_count": 208,
								"aggregate-count": {
									"value": 9853
								}
							},
							{
								"key_as_string": "2018-05-30T22:00:00.000Z",
								"key": 1527717600000,
								"doc_count": 191,
								"aggregate-count": {
									"value": 9415
								}
							},
							{
								"key_as_string": "2018-05-30T23:00:00.000Z",
								"key": 1527721200000,
								"doc_count": 200,
								"aggregate-count": {
									"value": 9410
								}
							},
							{
								"key_as_string": "2018-05-31T00:00:00.000Z",
								"key": 1527724800000,
								"doc_count": 0,
								"aggregate-count": {
									"value": 0
								}
							}
						]
					}
				}
			}`),
			expected: `{
				"buckets": [{
					"timestamp": 1527638400,
					"datetime": "2018-05-30T00:00:00.000Z",
					"value": 0
				}, {
					"timestamp": 1527642000,
					"datetime": "2018-05-30T01:00:00.000Z",
					"value": 0
				}, {
					"timestamp": 1527645600,
					"datetime": "2018-05-30T02:00:00.000Z",
					"value": 0
				}, {
					"timestamp": 1527649200,
					"datetime": "2018-05-30T03:00:00.000Z",
					"value": 0
				}, {
					"timestamp": 1527652800,
					"datetime": "2018-05-30T04:00:00.000Z",
					"value": 0
				}, {
					"timestamp": 1527656400,
					"datetime": "2018-05-30T05:00:00.000Z",
					"value": 0
				}, {
					"timestamp": 1527660000,
					"datetime": "2018-05-30T06:00:00.000Z",
					"value": 0
				}, {
					"timestamp": 1527663600,
					"datetime": "2018-05-30T07:00:00.000Z",
					"value": 0
				}, {
					"timestamp": 1527667200,
					"datetime": "2018-05-30T08:00:00.000Z",
					"value": 0
				}, {
					"timestamp": 1527670800,
					"datetime": "2018-05-30T09:00:00.000Z",
					"value": 0
				}, {
					"timestamp": 1527674400,
					"datetime": "2018-05-30T10:00:00.000Z",
					"value": 24885
				}, {
					"timestamp": 1527678000,
					"datetime": "2018-05-30T11:00:00.000Z",
					"value": 30851
				}, {
					"timestamp": 1527681600,
					"datetime": "2018-05-30T12:00:00.000Z",
					"value": 8991
				}, {
					"timestamp": 1527685200,
					"datetime": "2018-05-30T13:00:00.000Z",
					"value": 10268
				}, {
					"timestamp": 1527688800,
					"datetime": "2018-05-30T14:00:00.000Z",
					"value": 8761
				}, {
					"timestamp": 1527692400,
					"datetime": "2018-05-30T15:00:00.000Z",
					"value": 8251
				}, {
					"timestamp": 1527696000,
					"datetime": "2018-05-30T16:00:00.000Z",
					"value": 9642
				}, {
					"timestamp": 1527699600,
					"datetime": "2018-05-30T17:00:00.000Z",
					"value": 9057
				}, {
					"timestamp": 1527703200,
					"datetime": "2018-05-30T18:00:00.000Z",
					"value": 10673
				}, {
					"timestamp": 1527706800,
					"datetime": "2018-05-30T19:00:00.000Z",
					"value": 8816
				}, {
					"timestamp": 1527710400,
					"datetime": "2018-05-30T20:00:00.000Z",
					"value": 10586
				}, {
					"timestamp": 1527714000,
					"datetime": "2018-05-30T21:00:00.000Z",
					"value": 9853
				}, {
					"timestamp": 1527717600,
					"datetime": "2018-05-30T22:00:00.000Z",
					"value": 9415
				}, {
					"timestamp": 1527721200,
					"datetime": "2018-05-30T23:00:00.000Z",
					"value": 9410
				}, {
					"timestamp": 1527724800,
					"datetime": "2018-05-31T00:00:00.000Z",
					"value": 0
				}]
			}`,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			buffer := new(bytes.Buffer)
			err := json.Compact(buffer, testCase.searchResult)
			assert.IsType(t, nil, err)
			result, err := testCase.searchService.ParseSearchResult(buffer.Bytes())
			assert.IsType(t, nil, err)
			buffer = new(bytes.Buffer)
			err = json.Compact(buffer, []byte(testCase.expected))
			assert.IsType(t, nil, err)
			assert.Equal(t, string(buffer.Bytes()), string(result))
		})
	}
}
