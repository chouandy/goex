package elasticex

import (
	"time"

	"github.com/olivere/elastic"
)

// Client client
var Client *elastic.Client

// InitClient init client
func InitClient() (err error) {
	Client, err = elastic.NewClient(
		elastic.SetURL("https://"+endpoint),
		elastic.SetScheme("https"),
		elastic.SetSniff(false),
		elastic.SetRetrier(elastic.NewBackoffRetrier(
			elastic.NewExponentialBackoff(10*time.Millisecond, 8*time.Second),
		)),
	)

	return
}
