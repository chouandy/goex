package awsex

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/chouandy/goex/awsex/apigwex"
	"github.com/chouandy/goex/httpex"
)

// LambdaClient lambda client
var LambdaClient *lambda.Lambda

// InitLambdaClient init lambda client
func InitLambdaClient(region string) error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = region
	LambdaClient = lambda.New(cfg)

	return nil
}

// InitLambdaClientMiddleware init lambda client middleware
func InitLambdaClientMiddleware(ctx *apigwex.Context) error {
	if LambdaClient == nil {
		fmt.Print("[Middleware] Init Lambda Client...")
		if err := InitLambdaClient(ctx.Region); err != nil {
			fmt.Println(err)
			return httpex.NewError(500, "", "Failed to init lambda client")
		}
		fmt.Println("done")
	}

	return nil
}
