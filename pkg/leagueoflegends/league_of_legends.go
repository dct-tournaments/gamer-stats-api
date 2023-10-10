package leagueoflegends

import (
	"net/http"
	"time"
)

type PlatformRouting string
type RegionalRouting string

const (
	timeout = 30 * time.Second

	EUW1PlatformRouting PlatformRouting = "euw1"

	EuropeRegionalRouting RegionalRouting = "europe"
)

type Config struct {
	APIKey string `json:"league_of_legends_api_key" env:"LEAGUE_OF_LEGENDS_API_KEY"`
}

type service struct {
	httpClient *http.Client
	config     Config
}

func NewService(config Config) *service {
	return &service{
		httpClient: &http.Client{Timeout: timeout},
		config:     config,
	}
}
