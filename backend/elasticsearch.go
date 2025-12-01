package backend

import (
	"context"
	"fmt"

	"socialai/constants"
	"socialai/util"

	"github.com/olivere/elastic/v7"
)

// a global ESClient
var ESBackend *ElasticsearchBackend

// Wrap ESClient
type ElasticsearchBackend struct {
	client *elastic.Client
}

func InitElasticsearchBackend(config *util.ElasticsearchInfo) {
	client, err := elastic.NewClient(
		elastic.SetURL(config.Address),
		elastic.SetBasicAuth(config.Username, config.Password))
	if err != nil {
		panic(err)
	}

	exists, err := client.IndexExists(constants.POST_INDEX).Do(context.Background())
	if err != nil {
		panic(err)
	}

	if !exists {
		mapping := `{
           "mappings": {
               "properties": {
                   "id":       { "type": "keyword" },
                   "user":     { "type": "keyword" },
                   "message":  { "type": "text" },
                   "url":      { "type": "keyword", "index": false },
                   "type":     { "type": "keyword", "index": false }
               }
           }
       }`
		_, err := client.CreateIndex(constants.POST_INDEX).Body(mapping).Do(context.Background())
		if err != nil {
			panic(err)
		}
	}

	exists, err = client.IndexExists(constants.USER_INDEX).Do(context.Background())
	if err != nil {
		panic(err)
	}

	if !exists {
		mapping := `{
                       "mappings": {
                               "properties": {
                                       "username": {"type": "keyword"},
                                       "password": {"type": "keyword"},
                                       "age":      {"type": "long", "index": false},
                                       "gender":   {"type": "keyword", "index": false}
                               }
                       }
               }`
		_, err = client.CreateIndex(constants.USER_INDEX).Body(mapping).Do(context.Background())
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Indexes are created.")

	ESBackend = &ElasticsearchBackend{client: client}
}

func (backend *ElasticsearchBackend) ReadFromES(query elastic.Query, index string) (*elastic.SearchResult, error) {
	// Refresh the index before search to ensure we get the latest data
	// This prevents stale results due to Elasticsearch's near real-time nature
	_, err := backend.client.Refresh(index).Do(context.Background())
	if err != nil {
		// Log but don't fail - continue with search even if refresh fails
		fmt.Printf("Warning: Failed to refresh index %s: %v\n", index, err)
	}

	// First, get the total count to determine how many results to fetch
	countResult, err := backend.client.Count(index).Query(query).Do(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to count documents: %v", err)
	}

	// Set size to get all matching documents (with a reasonable upper limit)
	size := int(countResult)
	if size > 10000 {
		size = 10000 // Limit to prevent memory issues
	}

	// Debug: log the size being used
	fmt.Printf("Search query: count=%d, size=%d\n", countResult, size)

	searchResult, err := backend.client.Search().
		Index(index).
		Query(query).
		Size(size).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		return nil, err
	}

	return searchResult, nil
}

func (backend *ElasticsearchBackend) SaveToES(i interface{}, index string, id string) error {
	_, err := backend.client.Index().
		Index(index).
		Id(id).
		BodyJson(i).
		Refresh("wait_for").
		Do(context.Background())
	return err
}

func (backend *ElasticsearchBackend) DeleteFromES(query elastic.Query, index string) error {
	// DeleteByQuery doesn't support Refresh("wait_for"), so we use WaitForCompletion
	// and then manually refresh the index
	result, err := backend.client.DeleteByQuery().
		Index(index).
		Query(query).
		WaitForCompletion(true).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		return err
	}

	// Refresh the index after deletion to make changes immediately searchable
	_, err = backend.client.Refresh(index).Do(context.Background())
	if err != nil {
		fmt.Printf("Warning: Failed to refresh index after deletion: %v\n", err)
		// Don't fail the deletion if refresh fails
	}

	// Log deletion result for debugging
	if result != nil {
		fmt.Printf("DeleteByQuery result: deleted %d documents\n", result.Deleted)
	}

	return nil
}
