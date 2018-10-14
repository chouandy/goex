package s3managerex

import (
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3manager"
	"github.com/chouandy/goex/awsex/service/apigatewayex"
	"github.com/chouandy/goex/awsex/service/sfnex"
	"github.com/chouandy/goex/httpex"
)

// Uploader s3 uploader
var Uploader *s3manager.Uploader

// InitUploader init s3 uploader
func InitUploader() error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = os.Getenv("REGION")
	Uploader = s3manager.NewUploader(cfg)

	return nil
}

// InitUploaderMiddleware init s3 uploader middleware
func InitUploaderMiddleware(ctx *apigatewayex.Context) error {
	if Uploader == nil {
		fmt.Print("[Middleware] Init S3 Uploader...")
		if err := InitUploader(); err != nil {
			fmt.Println(err)
			return httpex.NewError(500, "", "Failed to init s3 uploader")
		}
		fmt.Println("done")
	}

	return nil
}

// InitUploaderTaskMiddleware init s3 uploader task middleware
func InitUploaderTaskMiddleware(ctx *sfnex.Context) error {
	if Uploader == nil {
		fmt.Print("[Middleware] Init S3 Uploader...")
		if err := InitUploader(); err != nil {
			fmt.Println(err)
			return errors.New("Failed to init s3 uploader")
		}
		fmt.Println("done")
	}

	return nil
}
