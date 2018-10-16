package firehosex

// Result firehose result struct
type Result struct {
	Attempted  int64 `json:"attempted"`
	Successful int64 `json:"successful"`
	Failed     int64 `json:"failed"`
}
