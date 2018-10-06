package awsex

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/chouandy/goex/awsex/apigwex"
	"github.com/chouandy/goex/httpex"
)

// SQSClient sqs client
var SQSClient *sqs.SQS

// InitSQSClient init sqs client
func InitSQSClient(region string) error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = region
	SQSClient = sqs.New(cfg)

	return nil
}

// InitSQSClientMiddleware init sqs client middleware
func InitSQSClientMiddleware(ctx *apigwex.Context) error {
	if SQSClient == nil {
		fmt.Print("[Middleware] Init SQS Client...")
		if err := InitSQSClient(ctx.Region); err != nil {
			fmt.Println(err)
			return httpex.NewError(500, "", "Failed to init sqs client")
		}
		fmt.Println("done")
	}

	return nil
}
