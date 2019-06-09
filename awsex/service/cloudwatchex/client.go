package cloudwatchex

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/chouandy/goex/awsex/service/apigatewayex"
	"github.com/chouandy/goex/httpex"
)

// Client cloudwatch client
var Client *cloudwatch.Client

// InitClient init cloudwatch client
func InitClient() error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = os.Getenv("REGION")
	Client = cloudwatch.New(cfg)

	return nil
}

// InitClientMiddleware init cloudwatch client middleware
func InitClientMiddleware(ctx *apigatewayex.Context) error {
	if Client == nil {
		fmt.Print("[Middleware] Init CloudWatch Client...")
		if err := InitClient(); err != nil {
			fmt.Println(err)
			return httpex.NewError(500, "", "Failed to init cloudwatch client")
		}
		fmt.Println("done")
	}

	return nil
}
