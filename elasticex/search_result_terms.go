package elasticex

import (
	"encoding/json"

	"github.com/olivere/elastic"
)

// SearchTermsSumItem search terms sum item struct
type SearchTermsSumItem struct {
	Value float64 `json:"value"`
}

// GetSearchTermBuckets get search term buckets
func (c *SearchService) GetSearchTermBuckets(sr *elastic.SearchResult) (json.RawMessage, error) {
	// Init buckets
	term := map[string]int64{}
	// Get term items
	if items, found := sr.Aggregations.Terms(c.TermsAggregation.Fields[0]); found {
		for _, bucket := range items.Buckets {
			// Get item value
			if c.TermsAggregation.SumAggregation == nil {
				term[bucket.Key.(string)] = bucket.DocCount
			} else {
				sum := SearchTermsSumItem{}
				sum.UnmarshalJSON(*bucket.Aggregations[c.TermsAggregation.SumAggregation.Field])
				term[bucket.Key.(string)] = (int64)(sum.Value)
			}
		}
	}

	return json.Marshal(term)
}

// GetSearchTermsBuckets get search terms buckets
func (c *SearchService) GetSearchTermsBuckets(sr *elastic.SearchResult) (json.RawMessage, error) {
	// Init terms
	terms := make([]map[string]interface{}, 0)
	// Get term items
	if items, found := sr.Aggregations.Terms(c.TermsAggregation.Fields[0]); found {
		for _, bucket := range items.Buckets {
			// Init term
			term := map[string]interface{}{
				c.TermsAggregation.Fields[0]: bucket.Key,
			}
			// Get sub term items
			subTerm := map[string]int64{}
			if subItems, found := bucket.Terms(c.TermsAggregation.Fields[1]); found {
				for _, subBucket := range subItems.Buckets {
					// Get item value
					if c.TermsAggregation.SumAggregation == nil {
						subTerm[subBucket.Key.(string)] = subBucket.DocCount
					} else {
						sum := SearchTermsSumItem{}
						sum.UnmarshalJSON(*subBucket.Aggregations[c.TermsAggregation.SumAggregation.Field])
						subTerm[subBucket.Key.(string)] = (int64)(sum.Value)
					}
				}
			}
			term[c.TermsAggregation.Fields[1]] = subTerm
			// Append term to terms
			terms = append(terms, term)
		}
	}

	return json.Marshal(terms)
}
