package awsex

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3manager"
	"github.com/chouandy/goex/awsex/apigwex"
	"github.com/chouandy/goex/httpex"
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

// InitS3UploaderMiddleware init s3 uploader middleware
func InitS3UploaderMiddleware(ctx *apigwex.Context) error {
	if S3Uploader == nil {
		fmt.Print("[Middleware] Init S3 Uploader...")
		if err := InitS3Uploader(ctx.Region); err != nil {
			fmt.Println(err)
			return httpex.NewError(500, "", "Failed to init s3 uploader")
		}
		fmt.Println("done")
	}

	return nil
}
