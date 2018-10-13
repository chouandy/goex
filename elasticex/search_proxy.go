package elasticex

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/chouandy/goex/awsex"
)

// SearchProxyInput search proxy input struct
type SearchProxyInput struct {
	Warmup bool `json:"warmup"`

	Index  string          `json:"index"`
	Type   string          `json:"type"`
	Source json.RawMessage `json:"source"`
}

// SendSearchProxyRequest send proxy request
func SendSearchProxyRequest(i *SearchProxyInput) (json.RawMessage, error) {
	// New payload
	payload, err := jsonex.Marshal(i)
	if err != nil {
		return nil, err
	}
	// New input
	input := &lambda.InvokeInput{
		FunctionName: functionName,
		Payload:      payload,
	}
	// New request
	req := awsex.LambdaClient.InvokeRequest(input)
	// Send request
	resp, err := req.Send()
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
