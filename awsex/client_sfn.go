package awsex

import (
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
)

// SFNClient sfn client
var SFNClient *sfn.SFN

// InitSFNClient init sfn client
func InitSFNClient(region string) error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = region
	SFNClient = sfn.New(cfg)

	return nil
}
