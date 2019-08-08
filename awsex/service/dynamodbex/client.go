package dynamodbex

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/chouandy/goex/awsex/service/apigatewayex"
	"github.com/chouandy/goex/httpex"
)

// Client dynamodb client
var Client *dynamodb.Client

// InitClient init dynamodb client
func InitClient() error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = os.Getenv("REGION")
	Client = dynamodb.New(cfg)

	return nil
}

// InitClientMiddleware init dynamodb client middleware
func InitClientMiddleware(ctx *apigatewayex.Context) error {
	if Client == nil {
		if err := InitClient(); err != nil {
			fmt.Printf("[Middleware] Init DynamoDB Client...%s\n", err)
			return httpex.NewError(500, "", "Failed to init dynamodb client")
		}
	}

	return nil
}
