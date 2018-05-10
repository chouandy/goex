package elasticsearch

import (
	"bytes"
	"encoding/json"
)

// IndexBulk elasticsearch index bulk struct
type IndexBulk struct {
	Action IndexBulkMetaData `json:"index"`
	Source []byte            `json:"-"`
}

// IndexBulkMetaData elasticsearch index bulk action struct
type IndexBulkMetaData struct {
	ID    string `json:"_id,omitempty"`
	Index string `json:"_index"`
	Type  string `json:"_type"`
}

// GetActionJSON get index bulk action json string
// example: { "index" : { "_index" : "test", "_type" : "type1", "_id" : "1" } }
func (c *IndexBulk) GetActionJSON() []byte {
	action, _ := json.Marshal(c)
	return action
}

// ToJSON generate elasticsearch bulk json
func (c *IndexBulk) ToJSON() []byte {
	bulk := make([][]byte, 2)
	bulk[0] = c.GetActionJSON()
	bulk[1] = c.Source

	return append(bytes.Join(bulk, []byte("\n")), '\n')
}
