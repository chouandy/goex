package testingex

import (
	"io"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// AwsSignerV4 httpex testing aws signer v4
type AwsSignerV4 struct {
	Region      string
	ServiceName string
}

// Sign sign aws signature v4
func (a *AwsSignerV4) Sign(r *http.Request, body io.ReadSeeker) (http.Header, error) {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic(err)
	}
	signer := v4.NewSigner(cfg.Credentials)

	return signer.Sign(r, body, a.ServiceName, a.Region, time.Now())
}
