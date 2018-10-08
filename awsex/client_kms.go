package awsex

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/chouandy/goex/awsex/apigwex"
	"github.com/chouandy/goex/httpex"
)

// KMSClient kms client
var KMSClient *kms.KMS

// InitKMSClient init kms client
func InitKMSClient(region string) error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = region
	KMSClient = kms.New(cfg)

	return nil
}

// InitKMSClientMiddleware init kms client middleware
func InitKMSClientMiddleware(ctx *apigwex.Context) error {
	if KMSClient == nil {
		fmt.Print("[Middleware] Init KMS Client...")
		if err := InitKMSClient(ctx.Region); err != nil {
			fmt.Println(err)
			return httpex.NewError(500, "", "Failed to init kms client")
		}
		fmt.Println("done")
	}

	return nil
}
