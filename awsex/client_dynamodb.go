package awsex

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/chouandy/goex/awsex/apigwex"
	"github.com/chouandy/goex/httpex"
)

// DynamoDBClient dynamodb client
var DynamoDBClient *dynamodb.DynamoDB

// InitDynamoDBClient init dynamodb client
func InitDynamoDBClient(region string) error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = region
	DynamoDBClient = dynamodb.New(cfg)

	return nil
}

// InitDynamoDBClientMiddleware init dynamodb client middleware
func InitDynamoDBClientMiddleware(ctx *apigwex.Context) error {
	if DynamoDBClient == nil {
		fmt.Print("[Middleware] Init DynamoDB Client...")
		if err := InitDynamoDBClient(ctx.Region); err != nil {
			fmt.Println(err)
			return httpex.NewError(500, "", "Failed to init dynamodb client")
		}
		fmt.Println("done")
	}

	return nil
}
