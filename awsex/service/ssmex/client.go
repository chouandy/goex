package ssmex

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/chouandy/goex/awsex/service/apigatewayex"
	"github.com/chouandy/goex/httpex"
)

// Client sts client
var Client *ssm.Client

// InitClient init sts client
func InitClient() error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = os.Getenv("REGION")
	Client = ssm.New(cfg)

	return nil
}

// InitClientMiddleware init ssm client middleware
func InitClientMiddleware(ctx *apigatewayex.Context) error {
	if Client == nil {
		fmt.Print("[Middleware] Init SSM Client...")
		if err := InitClient(); err != nil {
			fmt.Println(err)
			return httpex.NewError(500, "", "Failed to init ssm client")
		}
		fmt.Println("done")
	}

	return nil
}
