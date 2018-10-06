package apigwex

import "github.com/aws/aws-lambda-go/events"

// GetAPIKey Get api key
func GetAPIKey(request events.APIGatewayProxyRequest) string {
	// Get api key from header X-Api-Key
	if len(request.Headers["X-Api-Key"]) > 0 {
		return request.Headers["X-Api-Key"]
	}
	// Get api key from header x-api-key
	return request.Headers["x-api-key"]
}
