package cloudwatcheventsex

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
	"github.com/chouandy/goex/awsex/service/apigatewayex"
	"github.com/chouandy/goex/httpex"
)

// Client cloudwatchevents client
var Client *cloudwatchevents.Client

// InitClient init cloudwatchevents client
func InitClient() error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = os.Getenv("REGION")
	Client = cloudwatchevents.New(cfg)

	return nil
}

// InitClientMiddleware init cloudwatchevents client middleware
func InitClientMiddleware(ctx *apigatewayex.Context) error {
	if Client == nil {
		if err := InitClient(); err != nil {
			fmt.Printf("[Middleware] Init CloudWatch Events Client...%s\n", err)
			return httpex.NewError(500, "", "Failed to init cloudwatchevents client")
		}
	}

	return nil
}
