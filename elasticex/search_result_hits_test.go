package elasticex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSearchHits(t *testing.T) {
	// Set test cases
	testCases := []struct {
		searchService SearchService
		searchResult  json.RawMessage
		expected      string
	}{
		{
			searchService: SearchService{
				Size: 5,
			},
			searchResult: []byte(`{
				"took": 1,
				"timed_out": false,
				"_shards": {
					"total": 6,
					"successful": 6,
					"skipped": 0,
					"failed": 0
				},
				"hits": {
					"total": 14,
					"max_score": null,
					"hits": [
						{
							"_index": "device-logs-2018.06.02",
							"_type": "doc",
							"_id": "jH8VwWMBxp_-mcF1MeMo",
							"_score": null,
							"_source": {
								"timestamp": 1527952724,
								"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
								"profile-id": 213,
								"client-mac": "AABBCCDDEEFF",
								"client-ip": "192.168.212.89",
								"event-class": "warning",
								"event-type": "IR",
								"aggregate-count": 80,
								"dir": "out",
								"remote-ip": "140.112.172.1",
								"category-index": 2,
								"country_iso_code": "TW"
							},
							"sort": [
								1527952724000
							]
						},
						{
							"_index": "device-logs-2018.06.02",
							"_type": "doc",
							"_id": "i38VwWMBxp_-mcF1MeMo",
							"_score": null,
							"_source": {
								"timestamp": 1527952723,
								"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
								"profile-id": 213,
								"client-mac": "AABBCCDDEEFF",
								"client-ip": "192.168.212.89",
								"event-class": "warning",
								"event-type": "WSB",
								"aggregate-count": 7,
								"dir": "out",
								"remote-ip": "111.222.111.2",
								"category-index": 4,
								"remote-url": "http://www.bot.com.tw/Pages/default.aspx",
								"country_iso_code": "CN",
								"remote-url_scheme": "http",
								"remote-url_host": "www.bot.com.tw",
								"remote-url_port": 80,
								"remote-url_path": "/Pages/default.aspx"
							},
							"sort": [
								1527952723000
							]
						},
						{
							"_index": "device-logs-2018.06.02",
							"_type": "doc",
							"_id": "gX8VwWMBxp_-mcF1MeMF",
							"_score": null,
							"_source": {
								"timestamp": 1527952664,
								"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
								"profile-id": 213,
								"client-mac": "AABBCCDDEEFF",
								"client-ip": "192.168.212.89",
								"event-class": "warning",
								"event-type": "IR",
								"aggregate-count": 62,
								"dir": "out",
								"remote-ip": "140.112.172.1",
								"category-index": 2,
								"country_iso_code": "TW"
							},
							"sort": [
								1527952664000
							]
						},
						{
							"_index": "device-logs-2018.06.02",
							"_type": "doc",
							"_id": "fX8VwWMBxp_-mcF1MeMF",
							"_score": null,
							"_source": {
								"timestamp": 1527952660,
								"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
								"profile-id": 213,
								"client-mac": "AABBCCDDEEFF",
								"client-ip": "192.168.212.89",
								"event-class": "threat",
								"event-type": "IR",
								"aggregate-count": 46,
								"dir": "out",
								"remote-ip": "140.112.172.1",
								"category-index": 2,
								"country_iso_code": "TW"
							},
							"sort": [
								1527952660000
							]
						},
						{
							"_index": "device-logs-2018.06.02",
							"_type": "doc",
							"_id": "bn8VwWMBxp_-mcF1MeMF",
							"_score": null,
							"_source": {
								"timestamp": 1527952597,
								"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
								"profile-id": 213,
								"client-mac": "AABBCCDDEEFF",
								"client-ip": "192.168.212.89",
								"event-class": "threat",
								"event-type": "IR",
								"aggregate-count": 77,
								"dir": "out",
								"remote-ip": "140.112.172.1",
								"category-index": 2,
								"country_iso_code": "TW"
							},
							"sort": [
								1527952597000
							]
						}
					]
				}
			}`),
			expected: `{
				"search_after": 1527952597000,
				"hits": [{
					"timestamp": 1527952724,
					"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
					"profile-id": 213,
					"client-mac": "AABBCCDDEEFF",
					"client-ip": "192.168.212.89",
					"event-class": "warning",
					"event-type": "IR",
					"aggregate-count": 80,
					"dir": "out",
					"remote-ip": "140.112.172.1",
					"category-index": 2,
					"country_iso_code": "TW"
				}, {
					"timestamp": 1527952723,
					"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
					"profile-id": 213,
					"client-mac": "AABBCCDDEEFF",
					"client-ip": "192.168.212.89",
					"event-class": "warning",
					"event-type": "WSB",
					"aggregate-count": 7,
					"dir": "out",
					"remote-ip": "111.222.111.2",
					"category-index": 4,
					"remote-url": "http://www.bot.com.tw/Pages/default.aspx",
					"country_iso_code": "CN",
					"remote-url_scheme": "http",
					"remote-url_host": "www.bot.com.tw",
					"remote-url_port": 80,
					"remote-url_path": "/Pages/default.aspx"
				}, {
					"timestamp": 1527952664,
					"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
					"profile-id": 213,
					"client-mac": "AABBCCDDEEFF",
					"client-ip": "192.168.212.89",
					"event-class": "warning",
					"event-type": "IR",
					"aggregate-count": 62,
					"dir": "out",
					"remote-ip": "140.112.172.1",
					"category-index": 2,
					"country_iso_code": "TW"
				}, {
					"timestamp": 1527952660,
					"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
					"profile-id": 213,
					"client-mac": "AABBCCDDEEFF",
					"client-ip": "192.168.212.89",
					"event-class": "threat",
					"event-type": "IR",
					"aggregate-count": 46,
					"dir": "out",
					"remote-ip": "140.112.172.1",
					"category-index": 2,
					"country_iso_code": "TW"
				}, {
					"timestamp": 1527952597,
					"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
					"profile-id": 213,
					"client-mac": "AABBCCDDEEFF",
					"client-ip": "192.168.212.89",
					"event-class": "threat",
					"event-type": "IR",
					"aggregate-count": 77,
					"dir": "out",
					"remote-ip": "140.112.172.1",
					"category-index": 2,
					"country_iso_code": "TW"
				}]
			}`,
		},
		{
			searchService: SearchService{
				Size: 5,
			},
			searchResult: []byte(`{
				"took": 3,
				"timed_out": false,
				"_shards": {
					"total": 6,
					"successful": 6,
					"skipped": 0,
					"failed": 0
				},
				"hits": {
					"total": 65,
					"max_score": null,
					"hits": [
						{
							"_index": "device-logs-2018.06.02",
							"_type": "doc",
							"_id": "bH8VwWMBxp_-mcF1MeMF",
							"_score": null,
							"_source": {
								"timestamp": 1527952595,
								"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
								"profile-id": 213,
								"client-mac": "EEFFAABBCCDD",
								"client-ip": "192.168.212.89",
								"event-class": "threat",
								"event-type": "WSB",
								"aggregate-count": 15,
								"dir": "out",
								"remote-ip": "111.222.111.2",
								"category-index": 4,
								"remote-url": "http://www.bot.com.tw/Pages/default.aspx",
								"country_iso_code": "CN",
								"remote-url_scheme": "http",
								"remote-url_host": "www.bot.com.tw",
								"remote-url_port": 80,
								"remote-url_path": "/Pages/default.aspx"
							},
							"sort": [
								1527952595000
							]
						},
						{
							"_index": "device-logs-2018.06.02",
							"_type": "doc",
							"_id": "aX8VwWMBxp_-mcF1MeMF",
							"_score": null,
							"_source": {
								"timestamp": 1527952543,
								"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
								"profile-id": 213,
								"client-mac": "DDEEFFAABBCC",
								"client-ip": "192.168.212.89",
								"event-class": "threat",
								"event-type": "IR",
								"aggregate-count": 1,
								"dir": "out",
								"remote-ip": "140.112.172.1",
								"category-index": 2,
								"country_iso_code": "TW"
							},
							"sort": [
								1527952543000
							]
						},
						{
							"_index": "device-logs-2018.06.02",
							"_type": "doc",
							"_id": "aH8VwWMBxp_-mcF1MeMF",
							"_score": null,
							"_source": {
								"timestamp": 1527952542,
								"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
								"profile-id": 213,
								"client-mac": "AABBCCDDEEFF",
								"client-ip": "192.168.212.89",
								"event-class": "threat",
								"event-type": "AV",
								"aggregate-count": 11,
								"dir": "in",
								"remote-ip": "140.112.172.1",
								"file-name": "virus.exe",
								"md5": "d41d8cd98f00b204e9800998ecf8427e",
								"malware-group": "Pua.Conduit",
								"country_iso_code": "TW"
							},
							"sort": [
								1527952542000
							]
						},
						{
							"_index": "device-logs-2018.06.02",
							"_type": "doc",
							"_id": "Z38VwWMBxp_-mcF1MeMF",
							"_score": null,
							"_source": {
								"timestamp": 1527952541,
								"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
								"profile-id": 213,
								"client-mac": "FFAABBCCDDEE",
								"client-ip": "192.168.212.89",
								"event-class": "threat",
								"event-type": "AV",
								"aggregate-count": 49,
								"dir": "in",
								"remote-ip": "140.112.172.1",
								"file-name": "virus.exe",
								"md5": "d41d8cd98f00b204e9800998ecf8427e",
								"malware-group": "Pua.Conduit",
								"country_iso_code": "TW"
							},
							"sort": [
								1527952541000
							]
						},
						{
							"_index": "device-logs-2018.06.02",
							"_type": "doc",
							"_id": "Zn8VwWMBxp_-mcF1MeMF",
							"_score": null,
							"_source": {
								"timestamp": 1527952540,
								"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
								"profile-id": 213,
								"client-mac": "DDEEFFAABBCC",
								"client-ip": "192.168.212.89",
								"event-class": "threat",
								"event-type": "IR",
								"aggregate-count": 92,
								"dir": "out",
								"remote-ip": "140.112.172.1",
								"category-index": 2,
								"country_iso_code": "TW"
							},
							"sort": [
								1527952540000
							]
						}
					]
				}
			}`),
			expected: `{
				"search_after": 1527952540000,
				"hits": [{
					"timestamp": 1527952595,
					"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
					"profile-id": 213,
					"client-mac": "EEFFAABBCCDD",
					"client-ip": "192.168.212.89",
					"event-class": "threat",
					"event-type": "WSB",
					"aggregate-count": 15,
					"dir": "out",
					"remote-ip": "111.222.111.2",
					"category-index": 4,
					"remote-url": "http://www.bot.com.tw/Pages/default.aspx",
					"country_iso_code": "CN",
					"remote-url_scheme": "http",
					"remote-url_host": "www.bot.com.tw",
					"remote-url_port": 80,
					"remote-url_path": "/Pages/default.aspx"
				}, {
					"timestamp": 1527952543,
					"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
					"profile-id": 213,
					"client-mac": "DDEEFFAABBCC",
					"client-ip": "192.168.212.89",
					"event-class": "threat",
					"event-type": "IR",
					"aggregate-count": 1,
					"dir": "out",
					"remote-ip": "140.112.172.1",
					"category-index": 2,
					"country_iso_code": "TW"
				}, {
					"timestamp": 1527952542,
					"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
					"profile-id": 213,
					"client-mac": "AABBCCDDEEFF",
					"client-ip": "192.168.212.89",
					"event-class": "threat",
					"event-type": "AV",
					"aggregate-count": 11,
					"dir": "in",
					"remote-ip": "140.112.172.1",
					"file-name": "virus.exe",
					"md5": "d41d8cd98f00b204e9800998ecf8427e",
					"malware-group": "Pua.Conduit",
					"country_iso_code": "TW"
				}, {
					"timestamp": 1527952541,
					"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
					"profile-id": 213,
					"client-mac": "FFAABBCCDDEE",
					"client-ip": "192.168.212.89",
					"event-class": "threat",
					"event-type": "AV",
					"aggregate-count": 49,
					"dir": "in",
					"remote-ip": "140.112.172.1",
					"file-name": "virus.exe",
					"md5": "d41d8cd98f00b204e9800998ecf8427e",
					"malware-group": "Pua.Conduit",
					"country_iso_code": "TW"
				}, {
					"timestamp": 1527952540,
					"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
					"profile-id": 213,
					"client-mac": "DDEEFFAABBCC",
					"client-ip": "192.168.212.89",
					"event-class": "threat",
					"event-type": "IR",
					"aggregate-count": 92,
					"dir": "out",
					"remote-ip": "140.112.172.1",
					"category-index": 2,
					"country_iso_code": "TW"
				}]
			}`,
		},
		{
			searchService: SearchService{
				Size: 5,
			},
			searchResult: []byte(`{
				"took": 6,
				"timed_out": false,
				"_shards": {
					"total": 21,
					"successful": 20,
					"skipped": 0,
					"failed": 1,
					"failures": [{
						"shard": 0,
						"index": ".kibana",
						"node": "Jc51I35_S2ONZEpm23Wzhg",
						"reason": {
							"type": "query_shard_exception",
							"reason": "No mapping found for [timestamp] in order to sort on",
							"index_uuid": "zltFbZTzSoir5wnFKt8RgQ",
							"index": ".kibana"
						}
					}]
				},
				"hits": {
					"total": 0,
					"max_score": null,
					"hits": []
				}
			}`),
			expected: `{
				"hits": []
			}`,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			buffer := new(bytes.Buffer)
			err := json.Compact(buffer, testCase.searchResult)
			assert.Nil(t, err)
			result, err := testCase.searchService.ParseSearchResult(buffer.Bytes())
			assert.Nil(t, err)
			buffer = new(bytes.Buffer)
			err = json.Compact(buffer, []byte(testCase.expected))
			assert.Nil(t, err)
			assert.Equal(t, string(buffer.Bytes()), string(result))
		})
	}
}
