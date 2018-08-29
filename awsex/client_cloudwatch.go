package awsex

import (
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

// CloudWatchClient cloudwatch client
var CloudWatchClient *cloudwatch.CloudWatch

// InitCloudWatchClient init cloudwatch client
func InitCloudWatchClient(region string) error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = region
	CloudWatchClient = cloudwatch.New(cfg)

	return nil
}
