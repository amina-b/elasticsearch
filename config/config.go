package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var AppConfig Config

func Load() {
	err := godotenv.Load()

	if err != nil {
		log.Printf(".env file not presented")
	}

	AppConfig = Config{
		ElasticSearch{
			CloudID: os.Getenv("ELASTICSEARCH_CLOUD_ID"),
			ApiKey:  os.Getenv("ELASTICSEARCH_API_KEY"),
			Index:   os.Getenv("ELASTICSEARCH_INDEX_NAME"),
		},
		Service{
			Port: os.Getenv("SERVICE_PORT"),
		},
	}

}
