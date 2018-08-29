package awsex

import (
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// STSClient sts client
var STSClient *sts.STS

// InitSTSClient init sts client
func InitSTSClient(region string) error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = region
	STSClient = sts.New(cfg)

	return nil
}
