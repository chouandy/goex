package elasticex

import "fmt"

// Result firehose result struct
type Result struct {
	Attempted  int `json:"attempted"`
	Successful int `json:"successful"`
	Failed     int `json:"failed"`
}

// Inline inline
func (r *Result) Inline() string {
	return fmt.Sprintf("attempted: %d, successful: %d, failed: %d",
		r.Attempted, r.Successful, r.Failed,
	)
}
