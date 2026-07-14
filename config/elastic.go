package config

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
)

var ES *elasticsearch.Client

func InitElastic() {

	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	})

	if err != nil {
		panic(err)
	}

	ES = client

	info, err := ES.Info()
	if err != nil {
		panic(err)
	}
	defer info.Body.Close()

	fmt.Println("Connected to Elasticsearch")
}