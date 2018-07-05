package elasticex

// SearchAfterQuery search after query struct
type SearchAfterQuery struct {
	SearchAfter interface{}
	Gte         float64
}

// IsLastHit is last hit
func (c *SearchAfterQuery) IsLastHit(sort float64) bool {
	if c.Gte == 0 {
		return false
	}
	if c.Gte > sort {
		return true
	}

	return false
}
