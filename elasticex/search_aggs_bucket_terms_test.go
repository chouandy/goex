package elasticex

import (
	"bytes"
	json "encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTermsAggregation(t *testing.T) {
	// Set test cases
	testCases := []struct {
		agg      TermsAggregation
		expected struct {
			name   string
			source string
		}
	}{
		{
			agg: TermsAggregation{
				Fields: []string{
					"event-class",
				},
			},
			expected: struct {
				name   string
				source string
			}{
				name: "event-class",
				source: `{
					"terms": {
						"field": "event-class",
						"order": [{
							"_count": "desc"
						}]
					}
				}`,
			},
		},
		{
			agg: TermsAggregation{
				Fields: []string{
					"event-class",
				},
				SumAggregation: &SumAggregation{
					Field: "aggregate-count",
				},
			},
			expected: struct {
				name   string
				source string
			}{
				name: "event-class",
				source: `{
					"aggregations": {
						"aggregate-count": {
							"sum": {
								"field": "aggregate-count"
							}
						}
					},
					"terms": {
						"field": "event-class",
						"order": [{
							"aggregate-count": "desc"
						}]
					}
				}`,
			},
		},
		{
			agg: TermsAggregation{
				Fields: []string{
					"client-mac",
					"event-class",
				},
			},
			expected: struct {
				name   string
				source string
			}{
				name: "client-mac",
				source: `{
					"aggregations": {
						"event-class": {
							"terms": {
								"field": "event-class",
								"order": [{
									"_count": "desc"
								}]
							}
						}
					},
					"terms": {
						"field": "client-mac",
						"order": [{
							"_count": "desc"
						}]
					}
				}`,
			},
		},
		{
			agg: TermsAggregation{
				Fields: []string{
					"client-mac",
					"event-class",
				},
				RankingSize: 100,
			},
			expected: struct {
				name   string
				source string
			}{
				name: "client-mac",
				source: `{
					"aggregations": {
						"event-class": {
							"terms": {
								"field": "event-class",
								"order": [{
									"_count": "desc"
								}],
								"size": 100
							}
						}
					},
					"terms": {
						"field": "client-mac",
						"order": [{
							"_count": "desc"
						}],
						"size": 100
					}
				}`,
			},
		},
		{
			agg: TermsAggregation{
				Fields: []string{
					"client-mac",
					"event-class",
				},
				RankingSize: 100,
				SumAggregation: &SumAggregation{
					Field: "aggregate-count",
				},
			},
			expected: struct {
				name   string
				source string
			}{
				name: "client-mac",
				source: `{
					"aggregations": {
						"aggregate-count": {
							"sum": {
								"field": "aggregate-count"
							}
						},
						"event-class": {
							"aggregations": {
								"aggregate-count": {
									"sum": {
										"field": "aggregate-count"
									}
								}
							},
							"terms": {
								"field": "event-class",
								"order": [{
									"aggregate-count": "desc"
								}],
								"size": 100
							}
						}
					},
					"terms": {
						"field": "client-mac",
						"order": [{
							"aggregate-count": "desc"
						}],
						"size": 100
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
			data, err := json.Marshal(source)
			assert.Nil(t, err)
			buffer := new(bytes.Buffer)
			err = json.Compact(buffer, []byte(testCase.expected.source))
			assert.Nil(t, err)
			assert.Equal(t, string(buffer.Bytes()), string(data))
		})
	}
}
