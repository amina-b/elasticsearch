package config

type Config struct {
	ElasticSearch ElasticSearch
	Service       Service
}

type ElasticSearch struct {
	CloudID string
	ApiKey  string
	Index   string
}

type Service struct {
	Port string
}
