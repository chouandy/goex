package healthex

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

var metricName = "HealthyCount"

// CloudWatchMonitorCfg cloudwatch config struct
type CloudWatchMonitorCfg struct {
	Enable      bool
	Region      string
	Namespace   string
	ServiceName string        `mapstructure:"service_name"`
	Interval    time.Duration // Second
}

// CloudWatchMonitor cloudwatch monitor
type CloudWatchMonitor struct {
	Client *cloudwatch.Client
	Cfg    *CloudWatchMonitorCfg
}

// InitClient init client
func (c *CloudWatchMonitor) InitClient() error {
	// New aws config
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = c.Cfg.Region

	// New client
	c.Client = cloudwatch.New(cfg)

	return nil
}

// PutMetricData put metric data
func (c *CloudWatchMonitor) PutMetricData() error {
	// New timestamp
	timestamp := time.Now().UTC()

	// New input
	input := &cloudwatch.PutMetricDataInput{
		Namespace: aws.String(c.Cfg.Namespace),
		MetricData: []cloudwatch.MetricDatum{
			cloudwatch.MetricDatum{
				MetricName: aws.String(metricName),
				Timestamp:  aws.Time(timestamp),
				Unit:       cloudwatch.StandardUnitCount,
				Value:      aws.Float64(1),
				Dimensions: []cloudwatch.Dimension{
					cloudwatch.Dimension{
						Name:  aws.String("Service"),
						Value: aws.String(c.Cfg.ServiceName),
					},
				},
			},
		},
	}

	// New request
	req := c.Client.PutMetricDataRequest(input)
	// Send request
	if _, err := req.Send(context.Background()); err != nil {
		fmt.Printf("[Health][CloudWatchMonitor][PutMetricData] %s\n", err)
	}

	return nil
}

// Run run cloudwatch monitor
func (c *CloudWatchMonitor) Run() {
	go func() {
		for {
			// Put metric data
			c.PutMetricData()
			// Wait for interval
			time.Sleep(c.Cfg.Interval * time.Second)
		}
	}()
}

// RunCloudWatchMonitor run cloudwatch monitor
func RunCloudWatchMonitor(cfg *CloudWatchMonitorCfg) error {
	// Check enable or not
	if !cfg.Enable {
		return nil
	}

	// New monitor
	monitor := CloudWatchMonitor{Cfg: cfg}

	// Init client
	if err := monitor.InitClient(); err != nil {
		return err
	}

	// Run monitor
	monitor.Run()

	return nil
}
