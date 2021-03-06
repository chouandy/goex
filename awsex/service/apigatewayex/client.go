package apigatewayex

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/chouandy/goex/httpex"
)

// Client apigateway client
var Client *apigateway.Client

// InitClient init apigateway client
func InitClient() error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = os.Getenv("REGION")
	Client = apigateway.New(cfg)

	return nil
}

// InitClientMiddleware init apigateway client middleware
func InitClientMiddleware(ctx *Context) error {
	if Client == nil {
		if err := InitClient(); err != nil {
			fmt.Printf("[Middleware] Init APIGateway Client...%s\n", err)
			return httpex.NewError(500, "", "Failed to init apigateway client")
		}
	}

	return nil
}
