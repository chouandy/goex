package awsex

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/chouandy/goex/awsex/apigwex"
	"github.com/chouandy/goex/httpex"
)

// SSMClient sts client
var SSMClient *ssm.SSM

// InitSSMClient init sts client
func InitSSMClient(region string) error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = region
	SSMClient = ssm.New(cfg)

	return nil
}

// InitSSMClientMiddleware init ssm client middleware
func InitSSMClientMiddleware(ctx *apigwex.Context) error {
	if SSMClient == nil {
		fmt.Print("[Middleware] Init SSM Client...")
		if err := InitSSMClient(ctx.Region); err != nil {
			fmt.Println(err)
			return httpex.NewError(500, "", "Failed to init ssm client")
		}
		fmt.Println("done")
	}

	return nil
}
