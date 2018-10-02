package awsex

import (
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3manager"
)

// S3Downloader s3 uploader
var S3Downloader *s3manager.Downloader

// InitS3Downloader init s3 uploader
func InitS3Downloader(region string) error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = region
	S3Downloader = s3manager.NewDownloader(cfg)

	return nil
}
