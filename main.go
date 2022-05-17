package main

import (
	"log"
	"net/http"

	"github.com/amina-b/elasticsearch/config"
	e "github.com/amina-b/elasticsearch/endpoints"
	"github.com/amina-b/elasticsearch/service"
)

func main() {

	// Load configuration from env variables
	config.Load()

	// Initialize elasticsearch client
	service.Init(config.AppConfig.ElasticSearch.CloudID, config.AppConfig.ElasticSearch.ApiKey)

	// Add, update, delete documents
	e.Init(config.AppConfig.ElasticSearch.Index)

	log.Printf("listening on port %v", config.AppConfig.Service.Port)

	// Server
	err := http.ListenAndServe(config.AppConfig.Service.Port, nil)

	if err != nil {
		log.Fatalf("failed to start server. Error: %v", err)
	}

}
