package awsex

import (
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
)

// FirehoseClient firehose client
var FirehoseClient *firehose.Firehose

// InitFirehoseClient init firehose client
func InitFirehoseClient(region string) error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = region
	FirehoseClient = firehose.New(cfg)

	return nil
}
