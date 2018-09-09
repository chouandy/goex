package elasticex

import (
	"time"

	"github.com/olivere/elastic"
)

// ElasticClient elastic client
var ElasticClient *elastic.Client

// InitElasticClient init elastic client
func InitElasticClient(url string) (err error) {
	ElasticClient, err = elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetScheme("https"),
		elastic.SetSniff(false),
		elastic.SetRetrier(elastic.NewBackoffRetrier(
			elastic.NewExponentialBackoff(10*time.Millisecond, 8*time.Second),
		)),
	)

	return err
}
