package awsex

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/chouandy/goex/awsex/apigwex"
	"github.com/chouandy/goex/httpex"
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

// InitFirehoseClientMiddleware init firehose client middleware
func InitFirehoseClientMiddleware(ctx *apigwex.Context) error {
	if FirehoseClient == nil {
		fmt.Print("[Middleware] Init Firehose Client...")
		if err := InitFirehoseClient(ctx.Region); err != nil {
			fmt.Println(err)
			return httpex.NewError(500, "", "Failed to init firehose client")
		}
		fmt.Println("done")
	}

	return nil
}
