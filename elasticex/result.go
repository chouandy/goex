package elasticex

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/chouandy/goex/awsex/service/cloudwatchex"
)

// Result result struct
type Result struct {
	Attempted  int `json:"attempted"`
	Successful int `json:"successful"`
	Failed     int `json:"failed"`
}

// Inline inline
func (r *Result) Inline() string {
	return fmt.Sprintf("attempted: %d, successful: %d, failed: %d",
		r.Attempted, r.Successful, r.Failed,
	)
}

// PutMetricData put metric data
func (r *Result) PutMetricData(namespace, stage string) error {
	// Get timestamp
	timestamp := time.Now().UTC()
	// New input
	input := &cloudwatch.PutMetricDataInput{
		Namespace: aws.String(namespace),
		MetricData: []cloudwatch.MetricDatum{
			// Attempted metric
			cloudwatch.MetricDatum{
				MetricName: aws.String("AttemptedCount"),
				Timestamp:  aws.Time(timestamp),
				Unit:       cloudwatch.StandardUnitCount,
				Value:      aws.Float64(float64(r.Attempted)),
				Dimensions: []cloudwatch.Dimension{
					cloudwatch.Dimension{
						Name:  aws.String("Stage"),
						Value: aws.String(stage),
					},
				},
			},
			// Successful metric
			cloudwatch.MetricDatum{
				MetricName: aws.String("SuccessfulCount"),
				Timestamp:  aws.Time(timestamp),
				Unit:       cloudwatch.StandardUnitCount,
				Value:      aws.Float64(float64(r.Successful)),
				Dimensions: []cloudwatch.Dimension{
					cloudwatch.Dimension{
						Name:  aws.String("Stage"),
						Value: aws.String(stage),
					},
				},
			},
			// Failed metric
			cloudwatch.MetricDatum{
				MetricName: aws.String("FailedCount"),
				Timestamp:  aws.Time(timestamp),
				Unit:       cloudwatch.StandardUnitCount,
				Value:      aws.Float64(float64(r.Failed)),
				Dimensions: []cloudwatch.Dimension{
					cloudwatch.Dimension{
						Name:  aws.String("Stage"),
						Value: aws.String(stage),
					},
				},
			},
		},
	}
	// New request
	req := cloudwatchex.Client.PutMetricDataRequest(input)
	// Send request
	_, err := req.Send()

	return err
}
