package elasticex

import (
	"encoding/json"

	"github.com/olivere/elastic"
)

// SearchHits search hits struct
type SearchHits struct {
	SearchAfter interface{}       `json:"search_after,omitempty"`
	Hits        []json.RawMessage `json:"hits"`
}

// GetSearchHits get search hits
func (c *SearchService) GetSearchHits(sr *elastic.SearchResult) (json.RawMessage, error) {
	// Init hits
	hits := SearchHits{
		Hits: make([]json.RawMessage, 0),
	}
	// Check hits count
	hitsCount := len(sr.Hits.Hits)
	if hitsCount > 0 {
		// Iterate through hits
		for i, hit := range sr.Hits.Hits {
			// Get source
			hits.Hits = append(hits.Hits, *hit.Source)
			// Get search after
			if i+1 == hitsCount {
				hits.SearchAfter = hit.Sort[0]
			}
		}
	}

	return hits.MarshalJSON()
}
