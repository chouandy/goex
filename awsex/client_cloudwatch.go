package awsex

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/chouandy/goex/awsex/apigwex"
	"github.com/chouandy/goex/httpex"
)

// CloudWatchClient cloudwatch client
var CloudWatchClient *cloudwatch.CloudWatch

// InitCloudWatchClient init cloudwatch client
func InitCloudWatchClient(region string) error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = region
	CloudWatchClient = cloudwatch.New(cfg)

	return nil
}

// InitCloudWatchClientMiddleware init cloudwatch client middleware
func InitCloudWatchClientMiddleware(ctx *apigwex.Context) error {
	if CloudWatchClient == nil {
		fmt.Print("[Middleware] Init CloudWatch Client...")
		if err := InitCloudWatchClient(ctx.Region); err != nil {
			fmt.Println(err)
			return httpex.NewError(500, "", "Failed to init cloudwatch client")
		}
		fmt.Println("done")
	}

	return nil
}
