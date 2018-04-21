# aws-elastic
This repo will allow you to communicate with aws elasticsearch in golang.

# get the library
run ```go get github.com/cjacques1/aws-elastic```

# Calls supported
```
elasticSearch := ElasticSearch{
  RootURL: "<url to aws elastic instance",
  Index:   "<index name>",
  Type:    "<type name>",
}

elasticSearch.DeleteRecord(id string) (*DeleteResponse, error)
elasticSearch.CreateRecord(record interface{}) (*CreateResponse, error)
elasticSearch.GetRecords(query []byte) (*SearchResult, error)
elasticSearch.GetAllRecords() (*SearchResult, error)
elasticSearch.UpdateRecord(record interface{}, id string) error
```

# example using library

```
import (
  elastic "github.com/cjacques1/aws-elastic"
  
  "bytes"
  "encoding/json"
)

elasticSearch := ElasticSearch{
  RootURL: "<url to aws elastic instance",
  Index:   "<index name>",
  Type:    "<type name>",
}

elasticQuery := ElasticQuery{
	Query: Query{
		Bool: BoolQuery{
			Must: []MySearch{
				MySearch{Term: map[string]string{"lang": "en"}},
			},
			Should: []MySearch{
				MySearch{
					Bool: map[string][]MySearch{
						"must": []MySearch{
							MySearch{
								Match: map[string]string{
									"cleantext": "machine learning",
								},
							},
						},
					},
				},
			},
			MinimumShouldMatch: "1",
		},
	},
	From: "0",
	Size: "5000",
	Sort: []map[string]string{
		map[string]string{"unix_timestamp": "desc"},
		map[string]string{"_score": "desc"},
	},
}

reqBodyBytes := new(bytes.Buffer)
json.NewEncoder(reqBodyBytes).Encode(elasticQuery)

results, err := elasticSearch.GetRecords(reqBodyBytes.Bytes())
