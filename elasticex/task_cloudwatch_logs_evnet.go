package elasticex

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/olivere/elastic"
)

// CloudwatchLogsEventTask cloudwatch logs event task struct
type CloudwatchLogsEventTask struct {
	// CloudwatchLogs
	Event     events.CloudwatchLogsEvent
	LogEvents []events.CloudwatchLogsLogEvent

	// Elasticsearch
	Stage            string
	Region           string
	BulkableRequests []elastic.BulkableRequest
	Attempted        int
	Successful       int
	Failed           int
}

// GetLogEvents get log events
func (t *CloudwatchLogsEventTask) GetLogEvents() error {
	// Get cloudwatch logs data
	data, err := t.Event.AWSLogs.Parse()
	if err != nil {
		return err
	}

	// Ignore control message
	if data.MessageType == "CONTROL_MESSAGE" {
		return errors.New("Ignore control message")
	}

	// Set log events
	t.LogEvents = data.LogEvents

	return nil
}

// NewBulkableRequests new bulkable requests
func (t *CloudwatchLogsEventTask) NewBulkableRequests() {
	// New bulkable requests
	t.BulkableRequests = make([]elastic.BulkableRequest, 0)

	// Interate log events
	for _, event := range t.LogEvents {
		// Unmarshal event message
		var log map[string]interface{}
		if err := jsonex.Unmarshal([]byte(event.Message), &log); err != nil {
			continue
		}
		// Set stage, region
		log["stage"] = t.Stage
		log["region"] = t.Region
		// Get index by timestamp
		timestamp, err := time.Parse(time.RFC3339, log["timestamp"].(string))
		if err != nil {
			continue
		}
		index := indexPrefix + "-" + timestamp.Format("2006.01.02")
		// New bulkable request
		bulkableRequest := elastic.NewBulkIndexRequest().Index(index).Type("doc").Id(event.ID).Doc(log)
		// Append to bulkable requests
		t.BulkableRequests = append(t.BulkableRequests, bulkableRequest)
	}
}

// SendBulkableRequests send bulkable requests
func (t *CloudwatchLogsEventTask) SendBulkableRequests() error {
	// Set attempted
	t.Attempted = len(t.BulkableRequests)

	// Chunk bulkable requests and send
	for i := 0; i < t.Attempted; i += 50 {
		// New chunk end
		end := i + 50
		if end > t.Attempted {
			end = t.Attempted
		}
		// Get chunk
		chunk := t.BulkableRequests[i:end]
		// New bulk service
		bs := Client.Bulk()
		bs.Add(chunk...)
		// Send request
		resp, err := bs.Do(context.Background())
		if err != nil {
			return err
		}
		// Update successful, failed
		for _, item := range resp.Indexed() {
			if item.Status >= 300 {
				if debug {
					fmt.Println(item.Error)
				}
				t.Failed++
			} else {
				t.Successful++
			}
		}
	}

	return nil
}

// ResultInline result inline
func (t *CloudwatchLogsEventTask) ResultInline() string {
	return fmt.Sprintf("attempted: %d, successful: %d, failed: %d",
		t.Attempted, t.Successful, t.Failed,
	)
}
