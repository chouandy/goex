package elasticex

import (
	"github.com/olivere/elastic"
)

// BoolQuery bool query struct
type BoolQuery struct {
	mustClauses   []elastic.Query
	shouldClauses []elastic.Query
	RangeQuery    *RangeQuery
}

// NewBoolQuery new bool query
func NewBoolQuery() *BoolQuery {
	return &BoolQuery{
		mustClauses:   make([]elastic.Query, 0),
		shouldClauses: make([]elastic.Query, 0),
	}
}

// Build build query
func (c *BoolQuery) Build() *elastic.BoolQuery {
	q := elastic.NewBoolQuery()
	if len(c.mustClauses) > 0 {
		q = q.Must(c.mustClauses...)
	}
	if len(c.shouldClauses) > 0 {
		q = q.Must(c.shouldClauses...)
	}
	if c.RangeQuery != nil {
		q = q.Must(c.RangeQuery.Build())
	}

	return q
}

// SetClauses set clauses
func (c *BoolQuery) SetClauses(clauses map[string]interface{}) *BoolQuery {
	for name, value := range clauses {
		switch v := value.(type) {
		case string, int, int32, int64, float32, float64:
			c = c.Must(elastic.NewTermQuery(name, v))
		case []string:
			queries := make([]elastic.Query, 0)
			for _, v2 := range v {
				queries = append(queries, elastic.NewTermQuery(name, v2))
			}
			c = c.Should(queries...)
		case []int:
			queries := make([]elastic.Query, 0)
			for _, v2 := range v {
				queries = append(queries, elastic.NewTermQuery(name, v2))
			}
			c = c.Should(queries...)
		case []interface{}:
			queries := make([]elastic.Query, 0)
			for _, v2 := range v {
				queries = append(queries, elastic.NewTermQuery(name, v2))
			}
			c = c.Should(queries...)
		}
	}

	return c
}

// Must add must clauses
func (c *BoolQuery) Must(queries ...elastic.Query) *BoolQuery {
	c.mustClauses = append(c.mustClauses, queries...)
	return c
}

// Should add should clauses
func (c *BoolQuery) Should(queries ...elastic.Query) *BoolQuery {
	q := elastic.NewBoolQuery()
	q = q.MinimumNumberShouldMatch(1)
	q = q.Should(queries...)
	c.shouldClauses = append(c.shouldClauses, q)

	return c
}
