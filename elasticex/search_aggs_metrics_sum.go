package elasticex

import (
	"github.com/olivere/elastic"
)

// SumAggregation sum aggregation struct
type SumAggregation struct {
	Field string
}

// Build build query
func (c *SumAggregation) Build() *elastic.SumAggregation {
	return elastic.NewSumAggregation().Field(c.Field)
}
