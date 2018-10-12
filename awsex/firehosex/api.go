package firehosex

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/chouandy/goex/awsex"
)

// PutRecords put records to firehose
func PutRecords(deliveryStreamName string, records []firehose.Record) error {
	// Build request input
	input := &firehose.PutRecordBatchInput{
		DeliveryStreamName: aws.String(deliveryStreamName),
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

// PutChunkRecords put chunk records to firehose
func PutChunkRecords(deliveryStreamName string, records []firehose.Record, chunkSize int) error {
	// Get records num
	recordsNum := len(records)
	// Chunk records and upload to firehose
	for i := 0; i < recordsNum; i += chunkSize {
		// New chunk end
		end := i + chunkSize
		if end > recordsNum {
			end = recordsNum
		}
		// Get chunk
		chunk := records[i:end]
		// Upload to firehose
		if err := PutRecords(deliveryStreamName, chunk); err != nil {
			return err
		}
	}

	return nil
}
