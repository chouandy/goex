package cloudwatchex

import (
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/chouandy/goex/awsex/service/apigatewayex"
	"github.com/chouandy/goex/awsex/service/cloudwatcheventsex"
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
		if err := InitClient(); err != nil {
			fmt.Printf("[Middleware] Init CloudWatch Client...%s\n", err)
			return httpex.NewError(500, "", "Failed to init cloudwatch client")
		}
	}

	return nil
}

// InitClientEventMiddleware init cloudwatch client task middleware
func InitClientEventMiddleware(ctx *cloudwatcheventsex.Context) error {
	if Client == nil {
		if err := InitClient(); err != nil {
			fmt.Printf("[Middleware] Init CloudWatch Client...%s\n", err)
			return errors.New("Failed to init cloudwatch client")
		}
	}

	return nil
}
