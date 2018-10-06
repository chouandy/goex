package awsex

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/chouandy/goex/awsex/apigwex"
	"github.com/chouandy/goex/httpex"
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

// InitS3ClientMiddleware init s3 client middleware
func InitS3ClientMiddleware(ctx *apigwex.Context) error {
	if S3Client == nil {
		fmt.Print("[Middleware] Init S3 Client...")
		if err := InitS3Client(ctx.Region); err != nil {
			fmt.Println(err)
			return httpex.NewError(500, "", "Failed to init s3 client")
		}
		fmt.Println("done")
	}

	return nil
}
