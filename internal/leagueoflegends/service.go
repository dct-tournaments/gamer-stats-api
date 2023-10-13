package leagueoflegends

import (
	"context"
	"strings"

	"github.com/pkg/errors"

	"github.com/dct-tournaments/gamer-stats-api/pkg/leagueoflegends"
)

type LeagueOfLegendsAPIService interface {
	GetSummonerByName(ctx context.Context, region leagueoflegends.PlatformRouting, username string) (*leagueoflegends.Summoner, error)
	GetMatchesByPUUID(
		ctx context.Context,
		region leagueoflegends.RegionalRouting,
		puuid string,
		startTime *int64,
		endTime *int64,
		start *int,
		count *int,
	) ([]string, error)
	GetMatchByID(ctx context.Context, region leagueoflegends.RegionalRouting, matchID string) (*leagueoflegends.Match, error)
}

type service struct {
	leagueOfLegendsAPIService LeagueOfLegendsAPIService
}

type Service interface {
	GetPlayerStats(ctx context.Context, region string, name string, startAt *int64) (*PlayerStats, error)
}

func NewService(lolservice LeagueOfLegendsAPIService) Service {
	return &service{
		leagueOfLegendsAPIService: lolservice,
	}
}

func (s *service) getPlayerPUUIDByName(ctx context.Context, region string, name string) (string, error) {
	summoner, err := s.leagueOfLegendsAPIService.GetSummonerByName(ctx, leagueoflegends.PlatformRouting(region), name)
	if err != nil {
		return "", err
	}

	return summoner.Puuid, nil
}

func (s *service) GetPlayerStats(ctx context.Context, region string, name string, startAt *int64) (*PlayerStats, error) {
	puuid, err := s.getPlayerPUUIDByName(ctx, region, name)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get player PUUID by name")
	}

	matchIDs, err := s.leagueOfLegendsAPIService.GetMatchesByPUUID(ctx, leagueoflegends.RegionalRouting(region), puuid, startAt, nil, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get player matches by PUUID")
	}

	playerKillsCount := 0

	for _, id := range matchIDs {
		match, err := s.leagueOfLegendsAPIService.GetMatchByID(ctx, leagueoflegends.RegionalRouting(region), id)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get match by ID")
		}

		for _, player := range match.Info.Participants {
			if strings.EqualFold(player.PUUID, puuid) {
				playerKillsCount += player.Kills
			}
		}
	}

	return &PlayerStats{KillCount: playerKillsCount}, nil
}
