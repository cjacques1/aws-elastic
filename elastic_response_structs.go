package elastic

import (
	"encoding/json"
)

type SearchResult struct {
	TookInMillis int64         `json:"took"`
	TimedOut     bool          `json:"timed_out"`
	Shards       *ShardsInfo   `json:"_shards"`
	Hits         *SearchHits   `json:"hits"`
	Error        *ErrorDetails `json:"error,omitempty"`
	Status       int64         `json:"status,omitempty"`
}

// SearchHits specifies the list of search hits.
type SearchHits struct {
	TotalHits int64        `json:"total"`     // total number of hits found
	MaxScore  *float64     `json:"max_score"` // maximum score of all hits
	Hits      []*SearchHit `json:"hits"`      // the actual hits returned
}

// SearchHit is a single hit.
type SearchHit struct {
	Index  string           `json:"_index"`  // index name
	Type   string           `json:"_type"`   // type meta field
	ID     string           `json:"_id"`     // external or internal
	Score  *float64         `json:"_score"`  // computed score
	Source *json.RawMessage `json:"_source"` // stored document source
}

// ErrorDetails encapsulate error details from Elasticsearch.
// It is used in e.g. elastic.Error and elastic.BulkResponseItem.
type ErrorDetails struct {
	RootCause    []*ErrorDetails          `json:"root_cause,omitempty"`
	Type         string                   `json:"type"`
	Reason       string                   `json:"reason"`
	Phase        string                   `json:"phase,omitempty"`
	Grouped      bool                     `json:"grouped,omitempty"`
	FailedShards []map[string]interface{} `json:"failed_shards,omitempty"`
	CausedBy     map[string]interface{}   `json:"caused_by,omitempty"`
}

// shardsInfo represents information from a shard.
type ShardsInfo struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Failed     int `json:"failed"`
}

// CreateResponse is the struct for handling the response to creating an item
type CreateResponse struct {
	Index   string      `json:"_index,omitempty"`
	Type    string      `json:"_type,omitempty"`
	ID      string      `json:"_id,omitempty"`
	UID     string      `json:"_uid,omitempty"`
	Version *int64      `json:"_version,omitempty"`
	Result  string      `json:"result,omitempty"`
	Shards  *ShardsInfo `json:"_shards,omitempty"`
	Created bool        `json:"created,omitempty"`
}

// DeleteResponse is the struct for handling the response to deleting an item
type DeleteResponse struct {
	Found   bool        `json:"found,omitempty"`
	Index   string      `json:"_index,omitempty"`
	Type    string      `json:"_type,omitempty"`
	ID      string      `json:"_id,omitempty"`
	Version *int64      `json:"_version,omitempty"`
	Result  string      `json:"result,omitempty"`
	Shards  *ShardsInfo `json:"_shards,omitempty"`
}
