package awsex

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3manager"
	"github.com/chouandy/goex/awsex/apigwex"
	"github.com/chouandy/goex/httpex"
)

// S3Downloader s3 downloader
var S3Downloader *s3manager.Downloader

// InitS3Downloader init s3 downloader
func InitS3Downloader(region string) error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = region
	S3Downloader = s3manager.NewDownloader(cfg)

	return nil
}

// InitS3DownloaderMiddleware init s3 downloader middleware
func InitS3DownloaderMiddleware(ctx *apigwex.Context) error {
	if S3Downloader == nil {
		fmt.Print("[Middleware] Init S3 Downloader...")
		if err := InitS3Downloader(ctx.Region); err != nil {
			fmt.Println(err)
			return httpex.NewError(500, "", "Failed to init s3 downloader")
		}
		fmt.Println("done")
	}

	return nil
}
