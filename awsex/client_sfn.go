package awsex

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/chouandy/goex/awsex/apigwex"
	"github.com/chouandy/goex/httpex"
)

// SFNClient sfn client
var SFNClient *sfn.SFN

// InitSFNClient init sfn client
func InitSFNClient(region string) error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = region
	SFNClient = sfn.New(cfg)

	return nil
}

// InitSFNClientMiddleware init sfn client middleware
func InitSFNClientMiddleware(ctx *apigwex.Context) error {
	if SFNClient == nil {
		fmt.Print("[Middleware] Init SFN Client...")
		if err := InitSFNClient(ctx.Region); err != nil {
			fmt.Println(err)
			return httpex.NewError(500, "", "Failed to init sfn client")
		}
		fmt.Println("done")
	}

	return nil
}
