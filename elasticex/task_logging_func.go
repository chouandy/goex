package elasticex

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"gitlab.ecoworkinc.com/zyxel/sdk-go/service/gisapi"
)

// LoggingTransformFunc logging transform func
type LoggingTransformFunc func(*map[string]interface{})

// GetAPIGatewaySourceIPGeoInfo get api gateway source ip geo info
func GetAPIGatewaySourceIPGeoInfo(log *map[string]interface{}) {
	// Get identity
	identity, found := (*log)["identity"]
	if !found {
		return
	}

	// Get source ip
	sourceIP, found := identity.(map[string]interface{})["source_ip"]
	if !found {
		return
	}

	// New input
	input := gisapi.V2GetGeoInfoInput{
		IP: aws.String(sourceIP.(string)),
	}
	// New request
	req := gisapi.Client.V2GetGeoInfoRequest(&input)
	resp, err := req.Send()
	if err != nil {
		return
	}

	// Marshal response body
	data, err := jsonex.Marshal(resp)
	if err != nil {
		return
	}

	// Set source geo info
	(*log)["source_geo_info"] = json.RawMessage(data)
}
