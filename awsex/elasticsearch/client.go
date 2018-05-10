package elasticsearch

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// BulkResponse bulk response
type BulkResponse struct {
	Took   uint64 `json:"took"`
	Errors bool   `json:"errors"`
	Items  []struct {
		Create struct {
			Index  string `json:"_index"`
			Type   string `json:"_type"`
			ID     string `json:"_id"`
			Status int    `json:"status"`
			Error  string `json:"error"`
		} `json:"create"`
		Index struct {
			Index   string `json:"_index"`
			Type    string `json:"_type"`
			ID      string `json:"_id"`
			Version int    `json:"_version"`
			Status  int    `json:"status"`
			Error   string `json:"error"`
		} `json:"index"`
	} `json:"items"`
}

// Client elasticsearch client struct
type Client struct {
	Host url.URL
}

// NewClient new elasticsearch client
func NewClient(endpoint string) *Client {
	u := url.URL{
		Scheme: "https",
		Host:   endpoint,
	}
	return &Client{Host: u}
}

// Bulk elasticsearch bulk api
func (c *Client) Bulk(data []byte) (*BulkResponse, error) {
	// Build api url
	url := c.Host.String() + "/_bulk"
	// Build request body
	reqBody := bytes.NewBuffer(data)
	// Send Request
	res, err := http.Post(url, "application/x-ndjson", reqBody)
	if err != nil {
		return nil, err
	}
	resBody, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	// Parse response body
	output := &BulkResponse{}
	err = json.Unmarshal(resBody, output)
	if err != nil {
		return nil, err
	}

	return output, nil
}
