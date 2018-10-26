package elasticex

import (
	"encoding/json"
	"errors"

	"github.com/olivere/elastic"
)

// ParseSearchResult parse search result
func (c *SearchService) ParseSearchResult(result json.RawMessage) (json.RawMessage, error) {
	// Decode search result
	var sr elastic.SearchResult
	if err := jsonex.Unmarshal(result, &sr); err != nil {
		return nil, err
	}

	// Get search terms buckets
	if c.TermsAggregation != nil && len(c.TermsAggregation.Fields) == 1 {
		return c.GetSearchTermBuckets(&sr)
	}
	if c.TermsAggregation != nil && len(c.TermsAggregation.Fields) == 2 {
		return c.GetSearchTermsBuckets(&sr)
	}
	// Get search date histogram buckets
	if c.DateHistogramAggregation != nil {
		return c.GetSearchDateHistogramBuckets(&sr)
	}
	// Get search hits
	if c.Size > 0 {
		return c.GetSearchHits(&sr)
	}

	return nil, errors.New("Unsupported search result")
}
