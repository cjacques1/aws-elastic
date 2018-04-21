package elastic

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// ElasticSearch stores the basic parameters for es
type ElasticSearch struct {
	RootURL string
	Index   string
	Type    string
}

func (es *ElasticSearch) DeleteRecord(id string) (*DeleteResponse, error) {
	url := es.RootURL + "/" + es.Index + "/" + es.Type + "/" + id

	result, err := MakeDeleteRequest(url)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateRecord saves a record to ES
func (es *ElasticSearch) CreateRecord(record interface{}) (*CreateResponse, error) {
	url := es.RootURL + "/" + es.Index + "/" + es.Type + "/"

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(record)

	result, err := MakePostRequest(url, reqBodyBytes, CreateResponse{})
	if err != nil {
		return nil, err
	}
	typedResult, _ := result.(CreateResponse)
	return &typedResult, nil
}

// GetRecords returns all records that match the search query
func (es *ElasticSearch) GetRecords(query []byte) (*SearchResult, error) {
	url := es.RootURL + "/" + es.Index + "/" + es.Type + "/_search"

	w := bytes.Buffer{}
	w.Write(query)

	result, err := MakePostRequest(url, &w, SearchResult{})
	if err != nil {
		return nil, err
	}
	typedResult, _ := result.(SearchResult)
	return &typedResult, nil
}

// GetAllRecords returns all the stored records in the es type
func (es *ElasticSearch) GetAllRecords() (*SearchResult, error) {
	url := es.RootURL + "/" + es.Index + "/" + es.Type + "/_search"
	result, err := MakeGetRequest(url)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (es *ElasticSearch) UpdateRecord(record interface{}, id string) error {
	url := es.RootURL + "/" + es.Index + "/" + es.Type + "/" + id

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(record)

	_, err := http.Post(url, "application/json; charset=utf-8", reqBodyBytes)
	if err != nil {
		log.Println("Error: " + err.Error())
		return err
	}

	return nil
}
