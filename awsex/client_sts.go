package awsex

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/chouandy/goex/awsex/apigwex"
	"github.com/chouandy/goex/httpex"
)

// STSClient sts client
var STSClient *sts.STS

// InitSTSClient init sts client
func InitSTSClient(region string) error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = region
	STSClient = sts.New(cfg)

	return nil
}

// InitSTSClientMiddleware init sts client middleware
func InitSTSClientMiddleware(ctx *apigwex.Context) error {
	if STSClient == nil {
		fmt.Print("[Middleware] Init STS Client...")
		if err := InitSTSClient(ctx.Region); err != nil {
			fmt.Println(err)
			return httpex.NewError(500, "", "Failed to init sts client")
		}
		fmt.Println("done")
	}

	return nil
}
