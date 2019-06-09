package gormex

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3manager"
	"github.com/chouandy/goex/awsex/service/s3ex"
	"github.com/chouandy/goex/awsex/service/s3managerex"
	"github.com/gobuffalo/flect"
)

// HashSecret hash secret
var HashSecret = []byte(os.Getenv("GORM_ATTACHMENT_HASH_SECRET"))

// S3BucketName s3 bucket name
var S3BucketName = os.Getenv("GORM_ATTACHMENT_S3_BUCKET_NAME")

// Attachment attachment struct
type Attachment struct {
	ResourceName string
	ResourceID   uint64
	Name         string
	Content      []byte
	FileName     string
	ContentType  string
	UpdatedAt    time.Time
}

// GetUpdatedAtTimestampString get updated at timestamp string
func (a *Attachment) GetUpdatedAtTimestampString() string {
	return strconv.FormatInt(a.UpdatedAt.Unix(), 10)
}

// GetFileNameHashKey get file name hash key
// format: ":resource_name(plural)/:attachment_name(plural)/:resource_id/:style/:updated_at"
func (a *Attachment) GetFileNameHashKey() string {
	return strings.Join([]string{
		flect.Pluralize(flect.Underscore(a.ResourceName)),
		flect.Pluralize(flect.Underscore(a.Name)),
		strconv.FormatInt(int64(a.ResourceID), 10),
		"original",
		a.GetUpdatedAtTimestampString(),
	}, "/")
}

// GetHashedFileName get hashed file name
func (a *Attachment) GetHashedFileName() string {
	hashKey := []byte(a.GetFileNameHashKey())
	mac := hmac.New(sha1.New, HashSecret)
	mac.Write(hashKey)

	return fmt.Sprintf("%x", mac.Sum(nil))
}

// GetFileExtension get file extension
func (a *Attachment) GetFileExtension() string {
	return filepath.Ext(a.FileName)
}

// GetPath get path
// format: ":resource_name(plural)/:attachment_name(plural)/:resource_id/:style/:hash:extension"
func (a *Attachment) GetPath() string {
	return strings.Join([]string{
		flect.Pluralize(flect.Underscore(a.ResourceName)),
		flect.Pluralize(flect.Underscore(a.Name)),
		strconv.FormatInt(int64(a.ResourceID), 10),
		"original",
		a.GetHashedFileName() + a.GetFileExtension(),
	}, "/")
}

// Upload upload attachment to s3
func (a *Attachment) Upload() error {
	// New input
	input := &s3manager.UploadInput{
		Bucket:      aws.String(S3BucketName),
		Key:         aws.String(a.GetPath()),
		Body:        bytes.NewReader(a.Content),
		ContentType: aws.String(a.ContentType),
	}
	// Upload to s3
	if _, err := s3managerex.Uploader.Upload(input); err != nil {
		return err
	}

	return nil
}

// GetURL get url
func (a *Attachment) GetURL() string {
	// New input
	input := &s3.GetObjectInput{
		Bucket: aws.String(S3BucketName),
		Key:    aws.String(a.GetPath()),
	}
	// New request
	req := s3ex.Client.GetObjectRequest(input)
	// presign url with 1 hours expiration
	url, _ := req.Presign(3600 * time.Second)

	return url
}

// Delete delete
func (a *Attachment) Delete() error {
	// New input
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(S3BucketName),
		Key:    aws.String(a.GetPath()),
	}
	// New request
	req := s3ex.Client.DeleteObjectRequest(input)
	// Send request
	if _, err := req.Send(context.Background()); err != nil {
		return err
	}

	return nil
}
