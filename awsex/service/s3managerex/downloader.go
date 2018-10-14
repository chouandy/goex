package s3managerex

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3manager"
	"github.com/chouandy/goex/awsex/service/apigatewayex"
	"github.com/chouandy/goex/httpex"
)

// Downloader s3 downloader
var Downloader *s3manager.Downloader

// InitDownloader init s3 downloader
func InitDownloader() error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = os.Getenv("REGION")
	Downloader = s3manager.NewDownloader(cfg)

	return nil
}

// InitDownloaderMiddleware init s3 downloader middleware
func InitDownloaderMiddleware(ctx *apigatewayex.Context) error {
	if Downloader == nil {
		fmt.Print("[Middleware] Init S3 Downloader...")
		if err := InitDownloader(); err != nil {
			fmt.Println(err)
			return httpex.NewError(500, "", "Failed to init s3 downloader")
		}
		fmt.Println("done")
	}

	return nil
}
