package elasticex

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/olivere/elastic"
)

// HousekeepTask housekeep task struct
type HousekeepTask struct {
	Time      time.Time
	Retention int
	IndexName string
}

// Run run housekeep
func (t *HousekeepTask) Run() error {
	// New housekeep date
	housekeepDate := t.Time.AddDate(0, 0, -(t.Retention + 1)).Format("2006.01.02")
	// New index name
	t.IndexName = indexPrefix + "-" + housekeepDate
	// Send request
	resp, err := Client.DeleteIndex(t.IndexName).Do(context.Background())
	if err != nil {
		// 404 not found
		if elastic.IsNotFound(err) {
			return fmt.Errorf("index (%s) not found", t.IndexName)
		}
		return err
	}
	// Check acknowledged or not
	if !resp.Acknowledged {
		return errors.New("not acknowledged")
	}

	return nil
}

// ResultInline result inline
func (t *HousekeepTask) ResultInline() string {
	return fmt.Sprintf("index (%s) has been deleted", t.IndexName)
}
