package timex

import (
	"time"
)

// Efficiency efficiency struct
type Efficiency struct {
	StartAt time.Time
}

// Calculate calculate
func (e *Efficiency) Calculate(nums float64) float64 {
	return nums / time.Now().UTC().Sub(e.StartAt).Seconds()
}
