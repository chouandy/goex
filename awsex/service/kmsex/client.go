package kmsex

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/chouandy/goex/awsex/service/apigatewayex"
	"github.com/chouandy/goex/httpex"
)

// Client kms client
var Client *kms.Client

// InitClient init kms client
func InitClient() error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = os.Getenv("REGION")
	Client = kms.New(cfg)

	return nil
}

// InitClientMiddleware init kms client middleware
func InitClientMiddleware(ctx *apigatewayex.Context) error {
	if Client == nil {
		if err := InitClient(); err != nil {
			fmt.Printf("[Middleware] Init KMS Client...%s\n", err)
			return httpex.NewError(500, "", "Failed to init kms client")
		}
	}

	return nil
}
