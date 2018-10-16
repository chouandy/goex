package firehosex

import "fmt"

// Result firehose result struct
type Result struct {
	Attempted  int64 `json:"attempted"`
	Successful int64 `json:"successful"`
	Failed     int64 `json:"failed"`
}

// Inline inline
func (r *Result) Inline() string {
	return fmt.Sprintf("attempted: %d, successful: %d, failed: %d",
		r.Attempted, r.Successful, r.Failed,
	)
}
