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
				"took": 4,
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
					"total": 2238,
					"max_score": null,
					"hits": [{
							"_index": "device-logs-2018.05.30",
							"_type": "doc",
							"_id": "LH8osmMBxp_-mcF1BieZ",
							"_score": null,
							"_source": {
								"timestamp": "2018-05-30T17:40:59Z",
								"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
								"profile-id": 213,
								"client-mac": "CCDDEEFFAABB",
								"client-ip": "192.168.212.89",
								"event-class": "threat",
								"event-type": "WSB",
								"aggregate-count": 30,
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
								1527702059000
							]
						},
						{
							"_index": "device-logs-2018.05.30",
							"_type": "doc",
							"_id": "KX8osmMBxp_-mcF1BieZ",
							"_score": null,
							"_source": {
								"timestamp": "2018-05-30T17:40:56Z",
								"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
								"profile-id": 213,
								"client-mac": "BBCCDDEEFFAA",
								"client-ip": "192.168.212.89",
								"event-class": "warning",
								"event-type": "IR",
								"aggregate-count": 95,
								"dir": "out",
								"remote-ip": "140.112.172.1",
								"category-index": 2,
								"country_iso_code": "TW"
							},
							"sort": [
								1527702056000
							]
						},
						{
							"_index": "device-logs-2018.05.30",
							"_type": "doc",
							"_id": "JX8ismMBxp_-mcF12ieS",
							"_score": null,
							"_source": {
								"timestamp": "2018-05-30T17:40:02Z",
								"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
								"profile-id": 213,
								"client-mac": "AABBCCDDEEFF",
								"client-ip": "192.168.212.89",
								"event-class": "threat",
								"event-type": "IR",
								"aggregate-count": 67,
								"dir": "out",
								"remote-ip": "140.112.172.1",
								"category-index": 2,
								"country_iso_code": "TW"
							},
							"sort": [
								1527702002000
							]
						},
						{
							"_index": "device-logs-2018.05.30",
							"_type": "doc",
							"_id": "JH8ismMBxp_-mcF12ieS",
							"_score": null,
							"_source": {
								"timestamp": "2018-05-30T17:40:01Z",
								"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
								"profile-id": 213,
								"client-mac": "BBCCDDEEFFAA",
								"client-ip": "192.168.212.89",
								"event-class": "threat",
								"event-type": "IR",
								"aggregate-count": 25,
								"dir": "out",
								"remote-ip": "140.112.172.1",
								"category-index": 2,
								"country_iso_code": "TW"
							},
							"sort": [
								1527702001000
							]
						},
						{
							"_index": "device-logs-2018.05.30",
							"_type": "doc",
							"_id": "I38ismMBxp_-mcF12ieS",
							"_score": null,
							"_source": {
								"timestamp": "2018-05-30T17:40:00Z",
								"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
								"profile-id": 213,
								"client-mac": "BBCCDDEEFFAA",
								"client-ip": "192.168.212.89",
								"event-class": "threat",
								"event-type": "WSB",
								"aggregate-count": 22,
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
								1527702000000
							]
						}
					]
				}
			}`),
			expected: `{
				"search_after": 1527702000000,
				"hits": [{
					"timestamp": "2018-05-30T17:40:59Z",
					"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
					"profile-id": 213,
					"client-mac": "CCDDEEFFAABB",
					"client-ip": "192.168.212.89",
					"event-class": "threat",
					"event-type": "WSB",
					"aggregate-count": 30,
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
					"timestamp": "2018-05-30T17:40:56Z",
					"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
					"profile-id": 213,
					"client-mac": "BBCCDDEEFFAA",
					"client-ip": "192.168.212.89",
					"event-class": "warning",
					"event-type": "IR",
					"aggregate-count": 95,
					"dir": "out",
					"remote-ip": "140.112.172.1",
					"category-index": 2,
					"country_iso_code": "TW"
				}, {
					"timestamp": "2018-05-30T17:40:02Z",
					"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
					"profile-id": 213,
					"client-mac": "AABBCCDDEEFF",
					"client-ip": "192.168.212.89",
					"event-class": "threat",
					"event-type": "IR",
					"aggregate-count": 67,
					"dir": "out",
					"remote-ip": "140.112.172.1",
					"category-index": 2,
					"country_iso_code": "TW"
				}, {
					"timestamp": "2018-05-30T17:40:01Z",
					"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
					"profile-id": 213,
					"client-mac": "BBCCDDEEFFAA",
					"client-ip": "192.168.212.89",
					"event-class": "threat",
					"event-type": "IR",
					"aggregate-count": 25,
					"dir": "out",
					"remote-ip": "140.112.172.1",
					"category-index": 2,
					"country_iso_code": "TW"
				}, {
					"timestamp": "2018-05-30T17:40:00Z",
					"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
					"profile-id": 213,
					"client-mac": "BBCCDDEEFFAA",
					"client-ip": "192.168.212.89",
					"event-class": "threat",
					"event-type": "WSB",
					"aggregate-count": 22,
					"dir": "out",
					"remote-ip": "111.222.111.2",
					"category-index": 4,
					"remote-url": "http://www.bot.com.tw/Pages/default.aspx",
					"country_iso_code": "CN",
					"remote-url_scheme": "http",
					"remote-url_host": "www.bot.com.tw",
					"remote-url_port": 80,
					"remote-url_path": "/Pages/default.aspx"
				}]
			}`,
		},
		{
			searchService: SearchService{
				Size: 5,
			},
			searchResult: []byte(`{
				"took": 4,
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
					"total": 717,
					"max_score": null,
					"hits": [{
							"_index": "device-logs-2018.05.30",
							"_type": "doc",
							"_id": "JX8ismMBxp_-mcF12ieS",
							"_score": null,
							"_source": {
								"timestamp": "2018-05-30T17:40:02Z",
								"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
								"profile-id": 213,
								"client-mac": "AABBCCDDEEFF",
								"client-ip": "192.168.212.89",
								"event-class": "threat",
								"event-type": "IR",
								"aggregate-count": 67,
								"dir": "out",
								"remote-ip": "140.112.172.1",
								"category-index": 2,
								"country_iso_code": "TW"
							},
							"sort": [
								1527702002000
							]
						},
						{
							"_index": "device-logs-2018.05.30",
							"_type": "doc",
							"_id": "G38ismMBxp_-mcF12idX",
							"_score": null,
							"_source": {
								"timestamp": "2018-05-30T17:39:03Z",
								"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
								"profile-id": 213,
								"client-mac": "AABBCCDDEEFF",
								"client-ip": "192.168.212.89",
								"event-class": "warning",
								"event-type": "IPS",
								"aggregate-count": 97,
								"dir": "in",
								"remote-ip": "140.112.172.1",
								"sid": "801093",
								"category-index": 9,
								"message": "some description for this sid",
								"country_iso_code": "TW"
							},
							"sort": [
								1527701943000
							]
						},
						{
							"_index": "device-logs-2018.05.30",
							"_type": "doc",
							"_id": "F38ismMBxp_-mcF12idX",
							"_score": null,
							"_source": {
								"timestamp": "2018-05-30T17:38:59Z",
								"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
								"profile-id": 213,
								"client-mac": "AABBCCDDEEFF",
								"client-ip": "192.168.212.89",
								"event-class": "threat",
								"event-type": "AV",
								"aggregate-count": 7,
								"dir": "in",
								"remote-ip": "140.112.172.1",
								"file-name": "virus.exe",
								"md5": "d41d8cd98f00b204e9800998ecf8427e",
								"malware-group": "Pua.Conduit",
								"country_iso_code": "TW"
							},
							"sort": [
								1527701939000
							]
						},
						{
							"_index": "device-logs-2018.05.30",
							"_type": "doc",
							"_id": "D38ismMBxp_-mcF12idX",
							"_score": null,
							"_source": {
								"timestamp": "2018-05-30T17:38:03Z",
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
								1527701883000
							]
						},
						{
							"_index": "device-logs-2018.05.30",
							"_type": "doc",
							"_id": "Cn8ismMBxp_-mcF12idX",
							"_score": null,
							"_source": {
								"timestamp": "2018-05-30T17:37:58Z",
								"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
								"profile-id": 213,
								"client-mac": "AABBCCDDEEFF",
								"client-ip": "192.168.212.89",
								"event-class": "threat",
								"event-type": "AV",
								"aggregate-count": 70,
								"dir": "in",
								"remote-ip": "140.112.172.1",
								"file-name": "virus.exe",
								"md5": "d41d8cd98f00b204e9800998ecf8427e",
								"malware-group": "Pua.Conduit",
								"country_iso_code": "TW"
							},
							"sort": [
								1527701878000
							]
						}
					]
				}
			}`),
			expected: `{
				"search_after": 1527701878000,
				"hits": [{
					"timestamp": "2018-05-30T17:40:02Z",
					"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
					"profile-id": 213,
					"client-mac": "AABBCCDDEEFF",
					"client-ip": "192.168.212.89",
					"event-class": "threat",
					"event-type": "IR",
					"aggregate-count": 67,
					"dir": "out",
					"remote-ip": "140.112.172.1",
					"category-index": 2,
					"country_iso_code": "TW"
				}, {
					"timestamp": "2018-05-30T17:39:03Z",
					"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
					"profile-id": 213,
					"client-mac": "AABBCCDDEEFF",
					"client-ip": "192.168.212.89",
					"event-class": "warning",
					"event-type": "IPS",
					"aggregate-count": 97,
					"dir": "in",
					"remote-ip": "140.112.172.1",
					"sid": "801093",
					"category-index": 9,
					"message": "some description for this sid",
					"country_iso_code": "TW"
				}, {
					"timestamp": "2018-05-30T17:38:59Z",
					"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
					"profile-id": 213,
					"client-mac": "AABBCCDDEEFF",
					"client-ip": "192.168.212.89",
					"event-class": "threat",
					"event-type": "AV",
					"aggregate-count": 7,
					"dir": "in",
					"remote-ip": "140.112.172.1",
					"file-name": "virus.exe",
					"md5": "d41d8cd98f00b204e9800998ecf8427e",
					"malware-group": "Pua.Conduit",
					"country_iso_code": "TW"
				}, {
					"timestamp": "2018-05-30T17:38:03Z",
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
					"timestamp": "2018-05-30T17:37:58Z",
					"dev-uuid": "550e8400-e29b-41d4-a716-446655440000",
					"profile-id": 213,
					"client-mac": "AABBCCDDEEFF",
					"client-ip": "192.168.212.89",
					"event-class": "threat",
					"event-type": "AV",
					"aggregate-count": 70,
					"dir": "in",
					"remote-ip": "140.112.172.1",
					"file-name": "virus.exe",
					"md5": "d41d8cd98f00b204e9800998ecf8427e",
					"malware-group": "Pua.Conduit",
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
			assert.IsType(t, nil, err)
			result, err := testCase.searchService.ParseSearchResult(buffer.Bytes())
			assert.IsType(t, nil, err)
			buffer = new(bytes.Buffer)
			err = json.Compact(buffer, []byte(testCase.expected))
			assert.IsType(t, nil, err)
			assert.Equal(t, string(buffer.Bytes()), string(result))
		})
	}
}
