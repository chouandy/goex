package elasticex

import (
	"encoding/json"

	"github.com/olivere/elastic"
)

// SearchDateHistogramBuckets search date histogram buckets struct
type SearchDateHistogramBuckets struct {
	Buckets []SearchDateHistogramItem `json:"buckets"`
}

// SearchDateHistogramItem search date histogram item struct
type SearchDateHistogramItem struct {
	Timestamp int64  `json:"timestamp"`
	Datetime  string `json:"datetime"`
	Value     int64  `json:"value"`
}

// SearchDateHistogramSumItem search date histogram sum item struct
type SearchDateHistogramSumItem struct {
	Value float64 `json:"value"`
}

// GetSearchDateHistogramBuckets get search date histogram buckets
func (c *SearchService) GetSearchDateHistogramBuckets(sr *elastic.SearchResult) (json.RawMessage, error) {
	// Init buckets
	buckets := SearchDateHistogramBuckets{
		Buckets: make([]SearchDateHistogramItem, 0),
	}
	// Get date histogram items
	if items, found := sr.Aggregations.DateHistogram(c.DateHistogramAggregation.Field); found {
		for _, bucket := range items.Buckets {
			// Init item
			item := SearchDateHistogramItem{
				Timestamp: (int64)(bucket.Key / 1000),
				Datetime:  *bucket.KeyAsString,
			}
			// Get item value
			if c.DateHistogramAggregation.SumAggregation == nil {
				item.Value = bucket.DocCount
			} else {
				sum := SearchDateHistogramSumItem{}
				sum.UnmarshalJSON(*bucket.Aggregations[c.DateHistogramAggregation.SumAggregation.Field])
				item.Value = (int64)(sum.Value)
			}
			// Append buckets item
			buckets.Buckets = append(buckets.Buckets, item)
		}
	}

	return buckets.MarshalJSON()
}
