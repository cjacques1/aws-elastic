package elastic

// ElasticQuery stores the search query you want to make for elasticsearch
type ElasticQuery struct {
	Query  Query               `json:"query"`
	Fields []string            `json:"_source,omitempty"`
	From   string              `json:"from,omitempty"`
	Size   string              `json:"size,omitempty"`
	Sort   []map[string]string `json:"sort,omitempty"`
}

type Query struct {
	Bool BoolQuery `json:"bool"`
}

type BoolQuery struct {
	Must               []MySearch `json:"must,omitempty"`
	MustNot            []MySearch `json:"must_not,omitempty"`
	Should             []MySearch `json:"should,omitempty"`
	MinimumShouldMatch string     `json:"minimum_should_match,omitempty"`
}

// MySearch supports "match_phrase", "match", "term/s", "range", and "bool" type requests (see elasticsearch api for what these queries do).
type MySearch struct {
	Match       map[string]string             `json:"match_phrase,omitempty"`
	MatchObj    interface{}                   `json:"match,omitempty"`
	QueryString map[string]string             `json:"query_string,omitempty"`
	Terms       map[string][]string           `json:"terms,omitempty"`
	Term        map[string]string             `json:"term,omitempty"`
	Range       map[string]map[string]float64 `json:"range,omitempty"`
	Bool        map[string][]MySearch         `json:"bool,omitempty"`
}

type Match struct {
	Cleantext map[string]string `json:"cleantext,omitempty"`
}
