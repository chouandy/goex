package api

import "testing"

// 30.7 ns/op
func BenchmarkSetURL1(b *testing.B) {
	scheme := "http"
	endpoint := "test.example.com"

	setURL := func() string {
		// Set url
		var url string
		if len(scheme) > 0 {
			url = scheme + "://" + endpoint
		} else {
			url = "https://" + endpoint
		}

		return url
	}

	for n := 0; n < b.N; n++ {
		setURL()
	}
}

// 56.0 ns/op
func BenchmarkSetURL2(b *testing.B) {
	scheme := "http"
	endpoint := "test.example.com"

	setURL := func() string {
		// Set url
		url := "https://" + endpoint
		if len(scheme) > 0 {
			url = scheme + "://" + endpoint
		}

		return url
	}

	for n := 0; n < b.N; n++ {
		setURL()
	}
}
