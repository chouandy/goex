package testingex

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

// Request httpex testing request struct
type Request struct {
	URL                   string
	Endpoint              string
	Path                  string
	HTTPMethod            string
	Headers               map[string]string
	QueryStringParameters map[string]string
	PathParameters        map[string]string
	Body                  string
	AwsSignerV4           *AwsSignerV4
	CheckRedirect         bool
}

// GetURL get url
func (a *Request) GetURL() string {
	if len(a.URL) > 0 {
		return a.URL
	}
	return a.Endpoint + a.Path
}

// CompactRequestBody compact request body
func (a *Request) CompactRequestBody() []byte {
	buffer := new(bytes.Buffer)
	if err := json.Compact(buffer, []byte(a.Body)); err != nil {
		return []byte(a.Body)
	}

	return buffer.Bytes()
}

// SetHeaders set headers
func (a *Request) SetHeaders(req *http.Request) {
	// Interate request headers
	for k, v := range a.Headers {
		req.Header.Set(k, v)
	}
	// Set AWS Sigv4 Authorization header
	if a.AwsSignerV4 != nil {
		_, err := a.AwsSignerV4.Sign(req, bytes.NewReader(a.CompactRequestBody()))
		if err != nil {
			panic(err)
		}
	}
}

// Send send request
func (a *Request) Send() (*http.Response, error) {
	// New http client
	client := http.Client{
		Timeout: time.Duration(30 * time.Second),
	}
	// Check redirect
	if a.CheckRedirect {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}
	// New http request
	req, err := http.NewRequest(
		a.HTTPMethod,
		a.GetURL(),
		bytes.NewReader(a.CompactRequestBody()),
	)
	if err != nil {
		return nil, err
	}
	a.SetHeaders(req)

	return client.Do(req)
}
