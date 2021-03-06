package sfnex

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/chouandy/goex/awsex/service/apigatewayex"
	"github.com/chouandy/goex/httpex"
)

// Client sfn client
var Client *sfn.Client

// InitClient init sfn client
func InitClient() error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = os.Getenv("REGION")
	Client = sfn.New(cfg)

	return nil
}

// InitClientMiddleware init sfn client middleware
func InitClientMiddleware(ctx *apigatewayex.Context) error {
	if Client == nil {
		if err := InitClient(); err != nil {
			fmt.Printf("[Middleware] Init SFN Client...%s\n", err)
			return httpex.NewError(500, "", "Failed to init sfn client")
		}
	}

	return nil
}
