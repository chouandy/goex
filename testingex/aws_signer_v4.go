package testingex

import (
	"io"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// AwsSignerV4 httpex testing aws signer v4
type AwsSignerV4 struct {
	Region          string
	ServiceName     string
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
}

// Sign sign aws signature v4
func (a *AwsSignerV4) Sign(r *http.Request, body io.ReadSeeker) (http.Header, error) {
	var provider aws.CredentialsProvider
	if len(a.AccessKeyID) > 0 && len(a.SecretAccessKey) > 0 {
		// New credentials provider with access_key_id, secret_access_key
		provider = aws.NewStaticCredentialsProvider(
			a.AccessKeyID, a.SecretAccessKey, a.SessionToken,
		)
	} else {
		// Load credentials provider form env
		cfg, err := external.LoadDefaultAWSConfig()
		if err != nil {
			panic(err)
		}
		provider = cfg.Credentials
	}
	// New signer
	signer := v4.NewSigner(provider)

	return signer.Sign(r, body, a.ServiceName, a.Region, time.Now())
}
