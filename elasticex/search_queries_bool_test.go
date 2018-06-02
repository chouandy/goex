package elasticex

import (
	"bytes"
	json "encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoolQuery(t *testing.T) {
	// Set test cases
	testCases := []struct {
		query      map[string]interface{}
		rangeQuery *RangeQuery
		expected   string
	}{
		{
			query: map[string]interface{}{
				"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
				"event-class": []string{
					"threat",
					"warning",
				},
			},
			expected: `{
				"bool": {
					"must": [{
						"match_phrase": {
							"dev-uuid": {
								"query": "550e8400-e29b-41d4-a716-446655440000"
							}
						}
					}, {
						"bool": {
							"minimum_should_match": "1",
							"should": [{
								"match_phrase": {
									"event-class": {
										"query": "threat"
									}
								}
							}, {
								"match_phrase": {
									"event-class": {
										"query": "warning"
									}
								}
							}]
						}
					}]
				}
			}`,
		},
		{
			query: map[string]interface{}{
				"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
				"event-class": []string{
					"threat",
					"warning",
				},
				"client-mac": []string{
					"AABBCCDDEEFF",
					"BBCCDDEEFFAA",
					"CCDDEEFFAABB",
				},
			},
			expected: `{
				"bool": {
					"must": [{
						"match_phrase": {
							"dev-uuid": {
								"query": "550e8400-e29b-41d4-a716-446655440000"
							}
						}
					}, {
						"bool": {
							"minimum_should_match": "1",
							"should": [{
								"match_phrase": {
									"event-class": {
										"query": "threat"
									}
								}
							}, {
								"match_phrase": {
									"event-class": {
										"query": "warning"
									}
								}
							}]
						}
					}, {
						"bool": {
							"minimum_should_match": "1",
							"should": [{
								"match_phrase": {
									"client-mac": {
										"query": "AABBCCDDEEFF"
									}
								}
							}, {
								"match_phrase": {
									"client-mac": {
										"query": "BBCCDDEEFFAA"
									}
								}
							}, {
								"match_phrase": {
									"client-mac": {
										"query": "CCDDEEFFAABB"
									}
								}
							}]
						}
					}]
				}
			}`,
		},
		{
			query: map[string]interface{}{
				"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
				"event-class": []string{
					"threat",
					"warning",
				},
				"client-mac": []string{
					"AABBCCDDEEFF",
					"BBCCDDEEFFAA",
					"CCDDEEFFAABB",
				},
			},
			rangeQuery: &RangeQuery{
				Name:   "timestamp",
				Gte:    1527658861,
				Lte:    1527702061,
				Format: "epoch_second",
			},
			expected: `{
				"bool": {
					"must": [{
						"match_phrase": {
							"dev-uuid": {
								"query": "550e8400-e29b-41d4-a716-446655440000"
							}
						}
					}, {
						"bool": {
							"minimum_should_match": "1",
							"should": [{
								"match_phrase": {
									"event-class": {
										"query": "threat"
									}
								}
							}, {
								"match_phrase": {
									"event-class": {
										"query": "warning"
									}
								}
							}]
						}
					}, {
						"bool": {
							"minimum_should_match": "1",
							"should": [{
								"match_phrase": {
									"client-mac": {
										"query": "AABBCCDDEEFF"
									}
								}
							}, {
								"match_phrase": {
									"client-mac": {
										"query": "BBCCDDEEFFAA"
									}
								}
							}, {
								"match_phrase": {
									"client-mac": {
										"query": "CCDDEEFFAABB"
									}
								}
							}]
						}
					}, {
						"range": {
							"timestamp": {
								"format": "epoch_second",
								"from": 1527658861,
								"include_lower": true,
								"include_upper": true,
								"to": 1527702061
							}
						}
					}]
				}
			}`,
		},
		{
			query: map[string]interface{}{
				"dev-uuid":    "550e8400-e29b-41d4-a716-446655440000",
				"profile-id":  213,
				"event-class": "block",
			},
			rangeQuery: &RangeQuery{
				Name:   "timestamp",
				Gte:    1527658861,
				Lte:    1527702061,
				Format: "epoch_second",
			},
			expected: `{
				"bool": {
					"must": [{
						"match_phrase": {
							"dev-uuid": {
								"query": "550e8400-e29b-41d4-a716-446655440000"
							}
						}
					}, {
						"match_phrase": {
							"profile-id": {
								"query": 213
							}
						}
					}, {
						"match_phrase": {
							"event-class": {
								"query": "block"
							}
						}
					}, {
						"range": {
							"timestamp": {
								"format": "epoch_second",
								"from": 1527658861,
								"include_lower": true,
								"include_upper": true,
								"to": 1527702061
							}
						}
					}]
				}
			}`,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			query := NewBoolQuery()
			query = query.SetClauses(testCase.query)
			query.RangeQuery = testCase.rangeQuery
			source, err := query.Build().Source()
			assert.IsType(t, nil, err)
			data, err := json.Marshal(source)
			assert.IsType(t, nil, err)
			buffer := new(bytes.Buffer)
			err = json.Compact(buffer, []byte(testCase.expected))
			assert.IsType(t, nil, err)
			assert.Equal(t, string(buffer.Bytes()), string(data))
		})
	}
}
