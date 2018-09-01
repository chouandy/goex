package awsex

import (
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

// SSMClient sts client
var SSMClient *ssm.SSM

// InitSSMClient init sts client
func InitSSMClient(region string) error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = region
	SSMClient = ssm.New(cfg)

	return nil
}
