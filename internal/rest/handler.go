package rest

import (
	"context"

	"github.com/dct-tournaments/gamer-stats-api/internal/leagueoflegends"
	pleagueoflegends "github.com/dct-tournaments/gamer-stats-api/pkg/leagueoflegends"
	"github.com/gin-gonic/gin"
)

const (
	statsV0Group = "/stats-api/v0"
	statsV1Group = "/stats-api/v1"

	leagueOfLegendsStatsPath = "/league-of-legends"
)

type LeagueOfLegendsService interface {
	GetPlayerStatsByNameAndTagLine(
		ctx context.Context,
		region string,
		name string,
		tagLine string,
		startAt *int64,
		queueType *pleagueoflegends.QueueID,
	) (*leagueoflegends.PlayerStats, error)
	GetPlayerStatsByPUUID(
		ctx context.Context,
		region string,
		puuid string,
		startAt *int64,
		queueType *pleagueoflegends.QueueID,
	) (*leagueoflegends.PlayerStats, error)
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
	v1 := router.Group(statsV1Group)

	v0.GET(leagueOfLegendsStatsPath, h.GetLeagueOfLegendsStats)
	v1.GET(leagueOfLegendsStatsPath, h.GetLeagueOfLegendsStatsV1)
}
