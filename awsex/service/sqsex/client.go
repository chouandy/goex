package sqsex

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/chouandy/goex/awsex/service/apigatewayex"
	"github.com/chouandy/goex/httpex"
)

// Client sqs client
var Client *sqs.Client

// InitClient init sqs client
func InitClient() error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = os.Getenv("REGION")
	Client = sqs.New(cfg)

	return nil
}

// InitClientMiddleware init sqs client middleware
func InitClientMiddleware(ctx *apigatewayex.Context) error {
	if Client == nil {
		fmt.Print("[Middleware] Init SQS Client...")
		if err := InitClient(); err != nil {
			fmt.Println(err)
			return httpex.NewError(500, "", "Failed to init sqs client")
		}
		fmt.Println("done")
	}

	return nil
}
