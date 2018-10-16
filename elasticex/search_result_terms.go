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
	// Init term
	term := map[string]interface{}{}
	// Get sub term items
	subTerm := map[string]int64{}
	if items, found := sr.Aggregations.Terms(c.TermsAggregation.Fields[0]); found {
		for _, bucket := range items.Buckets {
			// Get item value
			if c.TermsAggregation.SumAggregation == nil {
				subTerm[bucket.Key.(string)] = bucket.DocCount
			} else {
				sum := SearchTermsSumItem{}
				jsonex.Unmarshal(*bucket.Aggregations[c.TermsAggregation.SumAggregation.Field], &sum)
				subTerm[bucket.Key.(string)] = int64(sum.Value)
			}
		}
	}
	term[c.TermsAggregation.Fields[0]] = subTerm

	return jsonex.Marshal(term)
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
						jsonex.Unmarshal(*subBucket.Aggregations[c.TermsAggregation.SumAggregation.Field], &sum)
						subTerm[subBucket.Key.(string)] = int64(sum.Value)
					}
				}
			}
			term[c.TermsAggregation.Fields[1]] = subTerm
			// Append term to terms
			terms = append(terms, term)
		}
	}

	return jsonex.Marshal(terms)
}
