package elasticex

import (
	"github.com/olivere/elastic"
)

// DateHistogramAggregation terms aggregation struct
type DateHistogramAggregation struct {
	Field             string
	Interval          string
	Format            string
	ExtendedBoundsMin int
	ExtendedBoundsMax int
	SumAggregation    *SumAggregation
}

// Build build query
func (c *DateHistogramAggregation) Build() (string, *elastic.DateHistogramAggregation) {
	agg := elastic.NewDateHistogramAggregation()
	// Set filed, interval
	agg = agg.Field(c.Field).Interval(c.Interval)
	// Set format
	if len(c.Format) == 0 {
		c.Format = "date_time_no_millis"
	}
	agg = agg.Format(c.Format)
	// Set extended bounds
	if c.ExtendedBoundsMin > 0 && c.ExtendedBoundsMax > 0 {
		agg = agg.ExtendedBounds(c.ExtendedBoundsMin, c.ExtendedBoundsMax)
	}
	// Set sum aggregation
	if c.SumAggregation != nil {
		agg = agg.SubAggregation(c.SumAggregation.Field, c.SumAggregation.Build())
	}

	return c.Field, agg
}
