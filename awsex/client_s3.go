package awsex

import (
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3Client s3 client
var S3Client *s3.S3

// InitS3Client init s3 client
func InitS3Client(region string) error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = region
	S3Client = s3.New(cfg)

	return nil
}
