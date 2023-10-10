package main

import (
	"context"

	"github.com/dct-tournaments/gamer-stats-api/pkg/leagueoflegends"
	"github.com/gin-gonic/gin"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	config, err := NewConfig()
	if err != nil {
		return
	}

	// services
	_ = leagueoflegends.NewService(config.LeagueOfLegends)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "healthy",
		})
	})

	r.Run()
}
