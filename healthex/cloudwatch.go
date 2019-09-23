package healthex

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

var cloudwatchClient *cloudwatch.Client
var metricName = "HealthyCount"

// CloudWatchMonitor cloudwatch config struct
type CloudWatchMonitor struct {
	Enable      bool
	Region      string
	Namespace   string
	ServiceName string        `mapstructure:"service_name"`
	Interval    time.Duration // Second
}

// Init init cloudwatch monitor
func (c *CloudWatchMonitor) Init() error {
	// Check enable or not
	if !c.Enable {
		return nil
	}

	// New client
	if err := c.InitClient(); err != nil {
		return err
	}

	go func() {
		for {
			// Put metric data
			c.PutMetricData()
			// Wait for interval
			time.Sleep(c.Interval * time.Second)
		}
	}()

	return nil
}

// InitClient init monitor
func (c *CloudWatchMonitor) InitClient() error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = c.Region
	cloudwatchClient = cloudwatch.New(cfg)

	return nil
}

// PutMetricData put metric data
func (c *CloudWatchMonitor) PutMetricData() error {
	// New timestamp
	timestamp := time.Now().UTC()

	// New input
	input := &cloudwatch.PutMetricDataInput{
		Namespace: aws.String(c.Namespace),
		MetricData: []cloudwatch.MetricDatum{
			cloudwatch.MetricDatum{
				MetricName: aws.String(metricName),
				Timestamp:  aws.Time(timestamp),
				Unit:       cloudwatch.StandardUnitCount,
				Value:      aws.Float64(1),
				Dimensions: []cloudwatch.Dimension{
					cloudwatch.Dimension{
						Name:  aws.String("Service"),
						Value: aws.String(c.ServiceName),
					},
				},
			},
		},
	}

	// New request
	req := cloudwatchClient.PutMetricDataRequest(input)
	// Send request
	if _, err := req.Send(context.Background()); err != nil {
		fmt.Printf("[Health][CloudWatchMonitor][PutMetricData] %s\n", err)
	}

	return nil
}
