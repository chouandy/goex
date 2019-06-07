package timex

import "time"

// NowToSecond now to second
func NowToSecond() time.Time {
	return time.Unix(time.Now().Unix(), 0)
}
