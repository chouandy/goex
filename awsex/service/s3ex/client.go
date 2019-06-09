package s3ex

import (
	"errors"
	"fmt"
	"os"

	"github.com/chouandy/goex/awsex/service/sfnex"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/chouandy/goex/awsex/service/apigatewayex"
	"github.com/chouandy/goex/httpex"
)

// Client s3 client
var Client *s3.Client

// InitClient init s3 client
func InitClient() error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = os.Getenv("REGION")
	Client = s3.New(cfg)

	return nil
}

// InitClientMiddleware init s3 client middleware
func InitClientMiddleware(ctx *apigatewayex.Context) error {
	if Client == nil {
		fmt.Print("[Middleware] Init S3 Client...")
		if err := InitClient(); err != nil {
			fmt.Println(err)
			return httpex.NewError(500, "", "Failed to init s3 client")
		}
		fmt.Println("done")
	}

	return nil
}

// InitClientTaskMiddleware init lambda client task middleware
func InitClientTaskMiddleware(ctx *sfnex.Context) error {
	if Client == nil {
		fmt.Print("[Middleware] Init S3 Client...")
		if err := InitClient(); err != nil {
			fmt.Println(err)
			return errors.New("Failed to init s3 client")
		}
		fmt.Println("done")
	}

	return nil
}
