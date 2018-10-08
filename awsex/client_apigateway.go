package awsex

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/chouandy/goex/awsex/apigwex"
	"github.com/chouandy/goex/httpex"
)

// APIGatewayClient apigateway client
var APIGatewayClient *apigateway.APIGateway

// InitAPIGatewayClient init apigateway client
func InitAPIGatewayClient(region string) error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = region
	APIGatewayClient = apigateway.New(cfg)

	return nil
}

// InitAPIGatewayClientMiddleware init apigateway client middleware
func InitAPIGatewayClientMiddleware(ctx *apigwex.Context) error {
	if APIGatewayClient == nil {
		fmt.Print("[Middleware] Init APIGateway Client...")
		if err := InitAPIGatewayClient(ctx.Region); err != nil {
			fmt.Println(err)
			return httpex.NewError(500, "", "Failed to init apigateway client")
		}
		fmt.Println("done")
	}

	return nil
}
