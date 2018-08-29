package awsex

import (
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3manager"
)

// S3Uploader s3 uploader
var S3Uploader *s3manager.Uploader

// InitS3Uploader init s3 uploader
func InitS3Uploader(region string) error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = region
	S3Uploader = s3manager.NewUploader(cfg)

	return nil
}
