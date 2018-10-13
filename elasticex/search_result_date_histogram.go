package elasticex

import (
	"encoding/json"

	"github.com/olivere/elastic"
)

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
	buckets := make([]SearchDateHistogramItem, 0)
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
				jsonex.Unmarshal(*bucket.Aggregations[c.DateHistogramAggregation.SumAggregation.Field], &sum)
				item.Value = (int64)(sum.Value)
			}
			// Append buckets item
			buckets = append(buckets, item)
		}
	}

	return jsonex.Marshal(buckets)
}
