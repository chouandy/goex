package elasticex

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/olivere/elastic"
)

// SNSEventTask sns event struct
type SNSEventTask struct {
	Event            events.SNSEvent
	BulkableRequests []elastic.BulkableRequest
	Result           *Result
}

// NewBulkableRequests new bulkable requests
func (t *SNSEventTask) NewBulkableRequests() {
	// New bulkable requests
	t.BulkableRequests = make([]elastic.BulkableRequest, 0)

	/* Handle log events concurrently */
	// New wg
	var wg sync.WaitGroup
	// wg add
	wg.Add(len(t.Event.Records))
	// New texts queue
	queue := make(chan elastic.BulkableRequest, 1)

	// Interate log events
	for _, record := range t.Event.Records {
		record := record
		go func() {
			// New index
			index := indexPrefix + "-" + record.SNS.Timestamp.Format("2006.01.02")
			// New bulkable request
			bulkableRequest := elastic.NewBulkIndexRequest().Index(index).Type("doc").Id(record.SNS.MessageID).Doc(record.SNS.Message)
			// Send bulkable request to queue
			queue <- bulkableRequest
		}()
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
func (t *SNSEventTask) SendBulkableRequests() error {
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
