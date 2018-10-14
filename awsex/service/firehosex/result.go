package firehosex

// Result firehose result struct
type Result struct {
	Attempted  int `json:"attempted"`
	Successful int `json:"successful"`
	Failed     int `json:"failed"`
}
