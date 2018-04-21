package elastic

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// MakePostRequest takes the url to post to, the request body as a bytes Buffer, and the struct to package the response in.
func MakePostRequest(url string, reqBodyBytes *bytes.Buffer, result interface{}) (interface{}, error) {
	resp, err := http.Post(url, "application/json; charset=utf-8", reqBodyBytes)
	if err != nil {
		log.Println("Error: " + err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	switch t := result.(type) {
	default:
		log.Println(t)
		return nil, errors.New("type not recognized")
	case CreateResponse:
		resultType, _ := result.(CreateResponse)

		json.Unmarshal(body, &resultType)
		if err != nil {
			log.Println("Error: " + err.Error())
			return nil, err
		}

		return resultType, nil
	case SearchResult:
		resultType, _ := result.(SearchResult)

		json.Unmarshal(body, &resultType)
		if err != nil {
			log.Println("Error: " + err.Error())
			return nil, err
		}

		return resultType, nil
	}
}

// MakeGetRequest takes a url and returns the ES result
func MakeGetRequest(url string) (*SearchResult, error) {
	result := new(SearchResult)

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error: " + err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &result)
	if err != nil {
		log.Println("Error: " + err.Error())
		return nil, err
	}

	return result, nil
}

// MakeDeleteRequest deletes the record from ES
func MakeDeleteRequest(myURLString string) (*DeleteResponse, error) {
	result := new(DeleteResponse)

	myURL, _ := url.Parse(myURLString)
	resp, err := http.DefaultClient.Do(&http.Request{
		Method: "DELETE",
		URL:    myURL,
	})
	if err != nil {
		log.Println("Error: " + err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &result)
	if err != nil {
		log.Println("Error: " + err.Error())
		return nil, err
	}

	return result, nil
}
