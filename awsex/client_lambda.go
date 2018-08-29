package awsex

import (
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

// LambdaClient lambda client
var LambdaClient *lambda.Lambda

// InitLambdaClient init lambda client
func InitLambdaClient(region string) error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = region
	LambdaClient = lambda.New(cfg)

	return nil
}
