package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/olivere/elastic"
)

const mapping = `
{
	"settings": {
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings": {
		"_doc": {
			"properties": {
				"name": {
					"type": "keyword"
				}
			}
		}
	}
}
`

func main() {
	ctx := context.Background()
	testLimitOfIndices(ctx)
}

func testLimitOfIndices(ctx context.Context) {
	time.Sleep(30 * time.Second)

	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://elasticsearch:9200"))
	if err != nil {
		fmt.Println("Pass 1")
		panic(err)
	}

	// Use the IndexExists service to check if a specified index exists.
	for i := 0; i <= 10100; i++ {
		exists, err := client.IndexExists("threads_" + strconv.Itoa(i)).Do(ctx)
		if err != nil {
			fmt.Println("Pass 2")
			panic(err)
		}
		if !exists {
			// Create a new index.
			createIndex(ctx, client, i)
		}
	}
}

func createIndex(ctx context.Context, client *elastic.Client, i int) {
	createIndex, err := client.CreateIndex("threads_" + strconv.Itoa(i)).BodyString(mapping).Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("create index : threads_" + strconv.Itoa(i))
	if !createIndex.Acknowledged {
		// Not acknowledged
	}
}
