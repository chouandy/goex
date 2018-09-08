package firehosex

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/chouandy/goex/awsex"
)

// PutRecords put records to firehose
func PutRecords(streamName string, records []firehose.Record) error {
	// Build request input
	input := &firehose.PutRecordBatchInput{
		DeliveryStreamName: aws.String(streamName),
		Records:            records,
	}
	// New request
	req := awsex.FirehoseClient.PutRecordBatchRequest(input)
	// Send request
	if _, err := req.Send(); err != nil {
		return err
	}

	return nil
}
