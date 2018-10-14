package apigatewayex

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/chouandy/goex/httpex"
)

// Client apigateway client
var Client *apigateway.APIGateway

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
		fmt.Print("[Middleware] Init APIGateway Client...")
		if err := InitClient(); err != nil {
			fmt.Println(err)
			return httpex.NewError(500, "", "Failed to init apigateway client")
		}
		fmt.Println("done")
	}

	return nil
}
