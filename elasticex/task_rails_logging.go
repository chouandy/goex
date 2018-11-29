package elasticex

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"sync"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/olivere/elastic"
)

// RailsLoggingTask rails logging task struct
type RailsLoggingTask struct {
	// CloudwatchLogs
	Event events.CloudwatchLogsEvent
	Data  events.CloudwatchLogsData

	// Elasticsearch
	Project          string
	Service          string
	Stage            string
	Region           string
	BulkableRequests []elastic.BulkableRequest
	Result           *Result

	// Extra
	TransformFuncs []RailsLoggingTransformFunc
}

// GetCloudwatchLogsData get cloudwatch logs data
func (t *RailsLoggingTask) GetCloudwatchLogsData() error {
	// Get cloudwatch logs data
	data, err := t.Event.AWSLogs.Parse()
	if err != nil {
		return err
	}

	// Ignore control message
	if data.MessageType == "CONTROL_MESSAGE" {
		return errors.New("Ignore control message")
	}

	// Set data
	t.Data = data

	return nil
}

// GetService get service
func (t *RailsLoggingTask) GetService() error {
	// New regexp
	re := regexp.MustCompile(
		fmt.Sprintf(`/%s-%s/%s-(.*)/rails-app-log`, t.Project, t.Stage, t.Project),
	)
	// Set service
	t.Service = re.FindStringSubmatch(t.Data.LogGroup)[1]
	// Check service
	if len(t.Service) == 0 {
		return errors.New("Failed to get service")
	}

	return nil
}

// NewBulkableRequests new bulkable requests
func (t *RailsLoggingTask) NewBulkableRequests() {
	// New bulkable requests
	t.BulkableRequests = make([]elastic.BulkableRequest, 0)

	/* Handle log events concurrently */
	// New wg
	var wg sync.WaitGroup
	// wg add
	wg.Add(len(t.Data.LogEvents))
	// New texts queue
	queue := make(chan elastic.BulkableRequest, 1)

	// Interate log events
	for _, event := range t.Data.LogEvents {
		event := event
		go func() {
			// Parse origin rails app log
			origin := new(OriginRailsAppLog)
			if err := jsonex.Unmarshal([]byte(event.Message), &origin); err != nil {
				wg.Done()
				return
			}
			// Check origin rails app log is valid or not
			if !origin.IsValid() {
				wg.Done()
				return
			}
			// Convert origin to flatten rails app log
			flatten := origin.Flatten()
			// Convert message.params
			if err := flatten.ConvertParamsToString(); err != nil {
				wg.Done()
				return
			}

			// Transform log
			for _, f := range t.TransformFuncs {
				f(flatten)
			}
			// Set stage, region
			flatten.Stage = t.Stage
			flatten.Region = t.Region
			// New index
			index := indexPrefix + "-" + t.Service + "-" + flatten.Timestamp.Format("2006.01.02")
			// New bulkable request
			bulkableRequest := elastic.NewBulkIndexRequest().Index(index).Type("doc").Id(event.ID).Doc(flatten)
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
func (t *RailsLoggingTask) SendBulkableRequests() error {
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

/* --------------------------------- */
/*           Rails App Log           */
/* --------------------------------- */

// OriginRailsAppLog origin rails app log
type OriginRailsAppLog struct {
	Timestamp *time.Time                `json:"timestamp"`
	Level     string                    `json:"level"`
	Message   *OriginRailsAppLogMessage `json:"message,omitempty"`
}

// OriginRailsAppLogMessage  origin rails app log message
type OriginRailsAppLogMessage struct {
	RequestID string            `json:"request_id"`
	RemoteIP  string            `json:"remote_ip"`
	Method    string            `json:"method"`
	Path      string            `json:"path"`
	Headers   map[string]string `json:"headers,omitempty"`
	Session   map[string]string `json:"session,omitempty"`
	Operator  map[string]string `json:"operator,omitempty"`
	Params    json.RawMessage   `json:"params,omitempty"`
	Status    int               `json:"status"`
	Duration  float64           `json:"duration"`
	View      float64           `json:"view"`
	DB        float64           `json:"db"`
	Location  string            `json:"location,omitempty"`
}

// Flatten flatten origin rails app log
func (o *OriginRailsAppLog) Flatten() *FlattenRailsAppLog {
	flatten := &FlattenRailsAppLog{
		Timestamp: o.Timestamp,
		Level:     o.Level,
		RequestID: o.Message.RequestID,
		RemoteIP:  o.Message.RemoteIP,
		Method:    o.Message.Method,
		Path:      o.Message.Path,
		Headers:   o.Message.Headers,
		Session:   o.Message.Session,
		Operator:  o.Message.Operator,
		Params:    o.Message.Params,
		Status:    o.Message.Status,
		Duration:  o.Message.Duration,
		View:      o.Message.View,
		DB:        o.Message.DB,
		Location:  o.Message.Location,
	}

	return flatten
}

// IsValid check origin rails app log is valid or not
func (o *OriginRailsAppLog) IsValid() bool {
	if o.Message == nil {
		return false
	}
	if o.Message.Status == 0 {
		return false
	}

	return true
}

// FlattenRailsAppLog flatten rails app log
type FlattenRailsAppLog struct {
	Stage         string            `json:"stage,omitempty"`
	Region        string            `json:"region,omitempty"`
	Timestamp     *time.Time        `json:"timestamp,omitempty"`
	Level         string            `json:"level"`
	RequestID     string            `json:"request_id"`
	RemoteIP      string            `json:"remote_ip"`
	Method        string            `json:"method"`
	Path          string            `json:"path"`
	Headers       map[string]string `json:"headers,omitempty"`
	Session       map[string]string `json:"session,omitempty"`
	Operator      map[string]string `json:"operator,omitempty"`
	Params        json.RawMessage   `json:"params,omitempty"`
	Status        int               `json:"status"`
	Duration      float64           `json:"duration"`
	View          float64           `json:"view"`
	DB            float64           `json:"db"`
	Location      string            `json:"location,omitempty"`
	RemoteGeoInfo json.RawMessage   `json:"remote_geo_info,omitempty"`
}

// ConvertParamsToString convert
func (f *FlattenRailsAppLog) ConvertParamsToString() (err error) {
	f.Params, err = jsonex.Marshal(string(f.Params))
	return
}
