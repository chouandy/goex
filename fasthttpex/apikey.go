package fasthttpex

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	yaml "gopkg.in/yaml.v2"
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
	// Initialize S3 uploader
	awsConfig := &aws.Config{Region: aws.String(c.S3Region)}
	uploader := s3manager.NewUploader(session.New(awsConfig))
	// Init input variable
	input := &s3manager.UploadInput{
		Bucket: aws.String(c.S3Bucket),
		Key:    aws.String(fmt.Sprintf("%s/%s", c.S3Prefix, c.S3Key)),
		Body:   file,
	}
	// Upload to s3
	_, err := uploader.Upload(input)

	return err
}

// Load load apikey whitelist from s3
func (c *APIKey) Load() error {
	// Initialize S3 svc
	awsConfig := &aws.Config{Region: aws.String(c.S3Region)}
	svc := s3.New(session.New(awsConfig))
	// 初始化 input 參數
	input := &s3.GetObjectInput{
		Bucket: aws.String(c.S3Bucket),
		Key:    aws.String(fmt.Sprintf("%s/%s", c.S3Prefix, c.S3Key)),
	}
	// 取得 bucket 中的 object
	result, err := svc.GetObject(input)
	if err != nil {
		return err
	}
	content, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return err
	}
	// Parse content
	var whitelist []map[string]string
	if err := yaml.Unmarshal(content, &whitelist); err != nil {
		return err
	}
	// Set whitelist
	for _, v := range whitelist {
		c.Whitelist[v["apikey"]] = v["vendor"]
	}

	return nil
}
