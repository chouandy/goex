package awsex

import (
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// DynamodbClient dynamodb client
var DynamodbClient *dynamodb.DynamoDB

// InitDynamodbClient init dynamodb client
func InitDynamodbClient(region string) error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = region
	DynamodbClient = dynamodb.New(cfg)

	return nil
}
