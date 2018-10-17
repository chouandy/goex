package elasticex

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/olivere/elastic"
)

// LoggingTask logging task struct
type LoggingTask struct {
	// CloudwatchLogs
	Event     events.CloudwatchLogsEvent
	LogEvents []events.CloudwatchLogsLogEvent

	// Elasticsearch
	Stage            string
	Region           string
	BulkableRequests []elastic.BulkableRequest
	Result           *Result
}

// GetLogEvents get log events
func (t *LoggingTask) GetLogEvents() error {
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
func (t *LoggingTask) NewBulkableRequests() {
	// New bulkable requests
	t.BulkableRequests = make([]elastic.BulkableRequest, 0)

	/* Handle log events concurrently */
	// New wg
	var wg sync.WaitGroup
	// wg add
	wg.Add(len(t.LogEvents))
	// New texts queue
	queue := make(chan elastic.BulkableRequest, 1)

	// Interate log events
	for i := range t.LogEvents {
		go func(event events.CloudwatchLogsLogEvent) {
			// Unmarshal event message
			var log map[string]interface{}
			if err := jsonex.Unmarshal([]byte(event.Message), &log); err != nil {
				return
			}
			// Set stage, region
			log["stage"] = t.Stage
			log["region"] = t.Region
			// Get index by timestamp
			timestamp, err := time.Parse(time.RFC3339, log["timestamp"].(string))
			if err != nil {
				return
			}
			index := indexPrefix + "-" + timestamp.Format("2006.01.02")
			// New bulkable request
			bulkableRequest := elastic.NewBulkIndexRequest().Index(index).Type("doc").Id(event.ID).Doc(log)
			// Send bulkable request to queue
			queue <- bulkableRequest
		}(t.LogEvents[i])
	}

	// Receive bulkable request from queue
	go func() {
		for bulkableRequest := range queue {
			// Append to bulkable requests
			t.BulkableRequests = append(t.BulkableRequests, bulkableRequest)
			// wg done
			wg.Done()
		}
	}()

	// wg wait
	wg.Wait()
}

// SendBulkableRequests send bulkable requests
func (t *LoggingTask) SendBulkableRequests() error {
	// Set attempted
	t.Result = &Result{
		Attempted: len(t.BulkableRequests),
	}

	// New ticker
	ticker := time.NewTicker(time.Second / 10)
	// Chunk bulkable requests and send
	for i := 0; i < t.Result.Attempted; i += 50 {
		<-ticker.C

		// New chunk end
		end := i + 50
		if end > t.Result.Attempted {
			end = t.Result.Attempted
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
				t.Result.Failed++
			} else {
				t.Result.Successful++
			}
		}
	}
	// Stop ticker
	ticker.Stop()

	return nil
}
