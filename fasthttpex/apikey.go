package fasthttpex

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3manager"
	"github.com/chouandy/goex/awsex"
	"gopkg.in/yaml.v2"
)

// APIKey api key struct
type APIKey struct {
	Whitelist map[string]string
	S3Region  string
	S3Bucket  string
	S3Prefix  string
	S3Key     string
}

// Upload upload apikey whitelist to s3
func (c *APIKey) Upload(file *os.File) error {
	// New input
	input := &s3manager.UploadInput{
		Bucket: aws.String(c.S3Bucket),
		Key:    aws.String(fmt.Sprintf("%s/%s", c.S3Prefix, c.S3Key)),
		Body:   file,
	}
	// Upload to s3
	_, err := awsex.S3Uploader.Upload(input)

	return err
}

// Load load apikey whitelist from s3
func (c *APIKey) Load() error {
	// New input
	input := &s3.GetObjectInput{
		Bucket: aws.String(c.S3Bucket),
		Key:    aws.String(fmt.Sprintf("%s/%s", c.S3Prefix, c.S3Key)),
	}
	// New request
	req := awsex.S3Client.GetObjectRequest(input)
	// Send request
	result, err := req.Send()
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return err
	}
	// Parse body
	var whitelist []map[string]string
	if err := yaml.Unmarshal(body, &whitelist); err != nil {
		return err
	}
	// Set whitelist
	for _, v := range whitelist {
		c.Whitelist[v["apikey"]] = v["vendor"]
	}

	return nil
}
