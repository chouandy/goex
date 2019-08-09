package ec2ex

import (
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/chouandy/goex/awsex/service/apigatewayex"
	"github.com/chouandy/goex/awsex/service/cloudwatcheventsex"
	"github.com/chouandy/goex/httpex"
)

// Client ec2 client
var Client *ec2.Client

// InitClient init ec2 client
func InitClient() error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = os.Getenv("REGION")
	Client = ec2.New(cfg)

	return nil
}

// InitClientMiddleware init ec2 client middleware
func InitClientMiddleware(ctx *apigatewayex.Context) error {
	if Client == nil {
		if err := InitClient(); err != nil {
			fmt.Printf("[Middleware] Init EC2 Client...%s\n", err)
			return httpex.NewError(500, "", "Failed to init ec2 client")
		}
	}

	return nil
}

// InitClientEventMiddleware init ec2 client task middleware
func InitClientEventMiddleware(ctx *cloudwatcheventsex.Context) error {
	if Client == nil {
		if err := InitClient(); err != nil {
			fmt.Printf("[Middleware] Init EC2 Client...%s\n", err)
			return errors.New("Failed to init ec2 client")
		}
	}

	return nil
}
