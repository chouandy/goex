package lambdaex

import (
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/chouandy/goex/awsex/service/apigatewayex"
	"github.com/chouandy/goex/awsex/service/sfnex"
	"github.com/chouandy/goex/httpex"
)

// Client lambda client
var Client *lambda.Client

// InitClient init lambda client
func InitClient() error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = os.Getenv("REGION")
	Client = lambda.New(cfg)

	return nil
}

// InitClientMiddleware init lambda client middleware
func InitClientMiddleware(ctx *apigatewayex.Context) error {
	if Client == nil {
		fmt.Print("[Middleware] Init Lambda Client...")
		if err := InitClient(); err != nil {
			fmt.Println(err)
			return httpex.NewError(500, "", "Failed to init lambda client")
		}
		fmt.Println("done")
	}

	return nil
}

// InitClientTaskMiddleware init lambda client task middleware
func InitClientTaskMiddleware(ctx *sfnex.Context) error {
	if Client == nil {
		fmt.Print("[Middleware] Init Lambda Client...")
		if err := InitClient(); err != nil {
			fmt.Println(err)
			return errors.New("Failed to init lambda client")
		}
		fmt.Println("done")
	}

	return nil
}
