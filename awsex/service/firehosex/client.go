package firehosex

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/chouandy/goex/awsex/service/apigatewayex"
	"github.com/chouandy/goex/httpex"
)

// Client firehose client
var Client *firehose.Client

// InitClient init firehose client
func InitClient() error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = os.Getenv("REGION")
	Client = firehose.New(cfg)

	return nil
}

// InitClientMiddleware init firehose client middleware
func InitClientMiddleware(ctx *apigatewayex.Context) error {
	if Client == nil {
		if err := InitClient(); err != nil {
			fmt.Printf("[Middleware] Init Firehose Client...%s\n", err)
			return httpex.NewError(500, "", "Failed to init firehose client")
		}
	}

	return nil
}
