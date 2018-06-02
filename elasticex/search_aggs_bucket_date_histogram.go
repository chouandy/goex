package elasticex

import (
	"github.com/olivere/elastic"
)

// DateHistogramAggregation terms aggregation struct
type DateHistogramAggregation struct {
	Field             string
	Interval          string
	ExtendedBoundsMin int
	ExtendedBoundsMax int
	SumAggregation    *SumAggregation
}

// Build build query
func (c *DateHistogramAggregation) Build() (string, *elastic.DateHistogramAggregation) {
	agg := elastic.NewDateHistogramAggregation()
	agg = agg.Field(c.Field).Interval(c.Interval)
	if c.ExtendedBoundsMin > 0 && c.ExtendedBoundsMax > 0 {
		agg = agg.ExtendedBounds(c.ExtendedBoundsMin, c.ExtendedBoundsMax)
	}
	if c.SumAggregation != nil {
		agg = agg.SubAggregation(c.SumAggregation.Field, c.SumAggregation.Build())
	}

	return c.Field, agg
}
