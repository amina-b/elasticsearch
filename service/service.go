package service

import (
	"log"

	es "github.com/elastic/go-elasticsearch/v8"
)

var Client *es.Client

// Connect to the elastic cloud with the given cloud id and api key
func Init(cloudID, apiKey string) error {

	config := es.Config{
		CloudID:                 cloudID,
		APIKey:                  apiKey,
		EnableCompatibilityMode: true,
	}

	client, err := es.NewClient(config)

	if err != nil {
		log.Printf("failed to create elasticsearch client. Error: %v", err)
		return err
	}

	Client = client

	return nil
}
