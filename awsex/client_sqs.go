package awsex

import (
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

// SQSClient sqs client
var SQSClient *sqs.SQS

// InitSQSClient init sqs client
func InitSQSClient(region string) error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = region
	SQSClient = sqs.New(cfg)

	return nil
}
