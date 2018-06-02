package elasticex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSearchTermBuckets(t *testing.T) {
	// Set test cases
	testCases := []struct {
		searchService SearchService
		searchResult  json.RawMessage
		expected      string
	}{
		{
			searchService: SearchService{
				TermsAggregation: &TermsAggregation{
					Fields: []string{
						"event-class",
					},
					SumAggregation: &SumAggregation{
						Field: "aggregate-count",
					},
				},
			},
			searchResult: []byte(`{
				"took": 6,
				"timed_out": false,
				"_shards": {
					"total": 21,
					"successful": 21,
					"skipped": 0,
					"failed": 0
				},
				"hits": {
					"total": 7066,
					"max_score": 0,
					"hits": []
				},
				"aggregations": {
					"event-class": {
						"doc_count_error_upper_bound": 0,
						"sum_other_doc_count": 0,
						"buckets": [{
								"key": "warning",
								"doc_count": 2403,
								"aggregate-count": {
									"value": 118390
								}
							},
							{
								"key": "threat",
								"doc_count": 2229,
								"aggregate-count": {
									"value": 108014
								}
							},
							{
								"key": "block",
								"doc_count": 2434,
								"aggregate-count": {
									"value": 118094
								}
							}
						]
					}
				}
			}`),
			expected: `{
				"event-class": {
					"block": 118094,
					"threat": 108014,
					"warning": 118390
				}
			}`,
		},
		{
			searchService: SearchService{
				TermsAggregation: &TermsAggregation{
					Fields: []string{
						"event-class",
					},
					SumAggregation: &SumAggregation{
						Field: "aggregate-count",
					},
				},
			},
			searchResult: []byte(`{
				"took": 2,
				"timed_out": false,
				"_shards": {
					"total": 21,
					"successful": 21,
					"skipped": 0,
					"failed": 0
				},
				"hits": {
					"total": 0,
					"max_score": 0,
					"hits": []
				},
				"aggregations": {
					"event-class": {
						"doc_count_error_upper_bound": 0,
						"sum_other_doc_count": 0,
						"buckets": []
					}
				}
			}`),
			expected: `{
				"event-class": {}
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

func TestGetSearchTermsBuckets(t *testing.T) {
	// Set test cases
	testCases := []struct {
		searchService SearchService
		searchResult  json.RawMessage
		expected      string
	}{
		{
			searchService: SearchService{
				TermsAggregation: &TermsAggregation{
					Fields: []string{
						"client-mac",
						"event-class",
					},
					SumAggregation: &SumAggregation{
						Field: "aggregate-count",
					},
				},
			},
			searchResult: []byte(`{
				"took": 13,
				"timed_out": false,
				"_shards": {
					"total": 21,
					"successful": 21,
					"skipped": 0,
					"failed": 0
				},
				"hits": {
					"total": 7066,
					"max_score": 0,
					"hits": []
				},
				"aggregations": {
					"client-mac": {
						"doc_count_error_upper_bound": 0,
						"sum_other_doc_count": 0,
						"buckets": [{
								"key": "FFAABBCCDDEE",
								"doc_count": 1022,
								"aggregate-count": {
									"value": 49915
								},
								"event-class": {
									"doc_count_error_upper_bound": 0,
									"sum_other_doc_count": 0,
									"buckets": [{
											"key": "warning",
											"doc_count": 410,
											"aggregate-count": {
												"value": 19119
											}
										},
										{
											"key": "threat",
											"doc_count": 293,
											"aggregate-count": {
												"value": 14784
											}
										},
										{
											"key": "block",
											"doc_count": 319,
											"aggregate-count": {
												"value": 16012
											}
										}
									]
								}
							},
							{
								"key": "EEFFAABBCCDD",
								"doc_count": 1217,
								"aggregate-count": {
									"value": 59217
								},
								"event-class": {
									"doc_count_error_upper_bound": 0,
									"sum_other_doc_count": 0,
									"buckets": [{
											"key": "warning",
											"doc_count": 406,
											"aggregate-count": {
												"value": 18682
											}
										},
										{
											"key": "threat",
											"doc_count": 395,
											"aggregate-count": {
												"value": 17506
											}
										},
										{
											"key": "block",
											"doc_count": 416,
											"aggregate-count": {
												"value": 23029
											}
										}
									]
								}
							},
							{
								"key": "DDEEFFAABBCC",
								"doc_count": 1251,
								"aggregate-count": {
									"value": 63023
								},
								"event-class": {
									"doc_count_error_upper_bound": 0,
									"sum_other_doc_count": 0,
									"buckets": [{
											"key": "warning",
											"doc_count": 469,
											"aggregate-count": {
												"value": 25474
											}
										},
										{
											"key": "threat",
											"doc_count": 421,
											"aggregate-count": {
												"value": 20510
											}
										},
										{
											"key": "block",
											"doc_count": 361,
											"aggregate-count": {
												"value": 17039
											}
										}
									]
								}
							},
							{
								"key": "CCDDEEFFAABB",
								"doc_count": 1143,
								"aggregate-count": {
									"value": 53745
								},
								"event-class": {
									"doc_count_error_upper_bound": 0,
									"sum_other_doc_count": 0,
									"buckets": [{
											"key": "warning",
											"doc_count": 341,
											"aggregate-count": {
												"value": 15592
											}
										},
										{
											"key": "threat",
											"doc_count": 334,
											"aggregate-count": {
												"value": 16748
											}
										},
										{
											"key": "block",
											"doc_count": 468,
											"aggregate-count": {
												"value": 21405
											}
										}
									]
								}
							},
							{
								"key": "BBCCDDEEFFAA",
								"doc_count": 1313,
								"aggregate-count": {
									"value": 64226
								},
								"event-class": {
									"doc_count_error_upper_bound": 0,
									"sum_other_doc_count": 0,
									"buckets": [{
											"key": "warning",
											"doc_count": 417,
											"aggregate-count": {
												"value": 20122
											}
										},
										{
											"key": "threat",
											"doc_count": 429,
											"aggregate-count": {
												"value": 21812
											}
										},
										{
											"key": "block",
											"doc_count": 467,
											"aggregate-count": {
												"value": 22292
											}
										}
									]
								}
							},
							{
								"key": "AABBCCDDEEFF",
								"doc_count": 1120,
								"aggregate-count": {
									"value": 54372
								},
								"event-class": {
									"doc_count_error_upper_bound": 0,
									"sum_other_doc_count": 0,
									"buckets": [{
											"key": "warning",
											"doc_count": 360,
											"aggregate-count": {
												"value": 19401
											}
										},
										{
											"key": "threat",
											"doc_count": 357,
											"aggregate-count": {
												"value": 16654
											}
										},
										{
											"key": "block",
											"doc_count": 403,
											"aggregate-count": {
												"value": 18317
											}
										}
									]
								}
							}
						]
					}
				}
			}`),
			expected: `[{
				"client-mac": "FFAABBCCDDEE",
				"event-class": {
					"block": 16012,
					"threat": 14784,
					"warning": 19119
				}
			}, {
				"client-mac": "EEFFAABBCCDD",
				"event-class": {
					"block": 23029,
					"threat": 17506,
					"warning": 18682
				}
			}, {
				"client-mac": "DDEEFFAABBCC",
				"event-class": {
					"block": 17039,
					"threat": 20510,
					"warning": 25474
				}
			}, {
				"client-mac": "CCDDEEFFAABB",
				"event-class": {
					"block": 21405,
					"threat": 16748,
					"warning": 15592
				}
			}, {
				"client-mac": "BBCCDDEEFFAA",
				"event-class": {
					"block": 22292,
					"threat": 21812,
					"warning": 20122
				}
			}, {
				"client-mac": "AABBCCDDEEFF",
				"event-class": {
					"block": 18317,
					"threat": 16654,
					"warning": 19401
				}
			}]`,
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
