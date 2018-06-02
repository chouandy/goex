package elasticex

import (
	"github.com/olivere/elastic"
)

// RangeQuery range query struct
type RangeQuery struct {
	Name   string
	Gte    int
	Lte    int
	Format string
}

// Build build query
func (c *RangeQuery) Build() *elastic.RangeQuery {
	return elastic.NewRangeQuery(c.Name).Gte(c.Gte).Lte(c.Lte).Format(c.Format)
}
