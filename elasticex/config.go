package elasticex

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
)

var (
	debug        = len(os.Getenv("DEBUG")) > 0
	endpoint     = os.Getenv("ELASTICSEARCH_ENDPOINT")
	functionName = aws.String(os.Getenv("SEARCH_PROXY_FUNCTION_NAME"))
	indexPrefix  = os.Getenv("ELASTICSEARCH_INDEX_PREFIX")
)
