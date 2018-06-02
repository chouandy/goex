package elasticex

// SearchService elasticsearch service struct
type SearchService struct {
	TermsAggregation         *TermsAggregation
	DateHistogramAggregation *DateHistogramAggregation
	BoolQuery                *BoolQuery
	Size                     int
	Sorter                   *Sorter
	SearchAfter              interface{}
}
