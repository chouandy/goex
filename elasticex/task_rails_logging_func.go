package elasticex

import "gitlab.ecoworkinc.com/zyxel/sdk-go/service/gis"

// RailsLoggingTransformFunc logging transform func
type RailsLoggingTransformFunc func(*FlattenRailsAppLog)

// GetRailsAppLogRemoteIPGeoInfo get api gateway source ip geo info
func GetRailsAppLogRemoteIPGeoInfo(flatten *FlattenRailsAppLog) {
	// Check remote ip
	if len(flatten.RemoteIP) == 0 {
		return
	}

	// New input
	input := &gis.GetGeoInfoInput{
		IP: flatten.RemoteIP,
	}
	// Send request
	resp, err := gis.Client.SendV1GeoInfoJSONRequest(input)
	if err != nil {
		return
	}

	// Set remote ip geo info
	flatten.RemoteGeoInfo = resp
}
