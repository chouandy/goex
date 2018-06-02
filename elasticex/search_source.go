package elasticex

import (
	"encoding/json"

	"github.com/olivere/elastic"
)

// BuildSearchSource build search source
func (c *SearchService) BuildSearchSource() (json.RawMessage, error) {
	// New search source
	ss := elastic.NewSearchSource()
	// Set size
	ss = ss.Size(c.Size)
	// Set sorter
	if c.Sorter != nil {
		ss = ss.Sort(c.Sorter.Field, c.Sorter.Ascending)
	}
	// Set search after
	if c.SearchAfter != nil {
		ss = ss.SearchAfter(c.SearchAfter)
	}
	// Set terms aggregation
	if c.TermsAggregation != nil {
		ss = ss.Aggregation(c.TermsAggregation.Build())
	}
	// Set date histogram aggregation
	if c.DateHistogramAggregation != nil {
		ss = ss.Aggregation(c.DateHistogramAggregation.Build())
	}
	// Set bool query
	if c.BoolQuery != nil {
		ss = ss.Query(c.BoolQuery.Build())
	}
	// Generate search source
	src, err := ss.Source()
	if err != nil {
		return nil, err
	}

	return json.Marshal(src)
}
