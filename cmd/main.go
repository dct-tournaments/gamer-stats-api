package main

import (
	"context"

	"github.com/dct-tournaments/gamer-stats-api/internal/leagueoflegends"
	"github.com/dct-tournaments/gamer-stats-api/internal/rest"
	lolAPI "github.com/dct-tournaments/gamer-stats-api/pkg/leagueoflegends"
	"github.com/gin-gonic/gin"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	config, err := NewConfig()
	if err != nil {
		return
	}

	// external API services
	lolAPIService := lolAPI.NewService(config.LeagueOfLegends)

	// internal services
	lolservice := leagueoflegends.NewService(lolAPIService)

	// rest handlers
	lolRestHandler := rest.NewHandler(lolservice)

	router := gin.Default()

	lolRestHandler.Register(router)

	router.Run(":8080")
}
