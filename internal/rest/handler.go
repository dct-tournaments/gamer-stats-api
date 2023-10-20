package rest

import (
	"context"

	"github.com/dct-tournaments/gamer-stats-api/internal/leagueoflegends"
	"github.com/gin-gonic/gin"
)

const (
	statsV0Group = "/stats-api/v0"

	leagueOfLegendsStatsPath = "/league-of-legends"
)

type LeagueOfLegendsService interface {
	GetPlayerStats(ctx context.Context, region string, name string, startAt *int64) (*leagueoflegends.PlayerStats, error)
}

// Handler defines the REST API handlers.
type handler struct {
	leagueOfLegendsService LeagueOfLegendsService
}

// NewHandler returns a new instance of Handler.
func NewHandler(lolservice LeagueOfLegendsService) *handler {
	return &handler{
		leagueOfLegendsService: lolservice,
	}
}

func (h *handler) Register(router *gin.Engine) {
	v0 := router.Group(statsV0Group)

	v0.GET(leagueOfLegendsStatsPath, h.GetLeagueOfLegendsStats)
}
