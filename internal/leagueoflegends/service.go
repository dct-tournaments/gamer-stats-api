package leagueoflegends

import (
	"context"
	"strings"

	"github.com/pkg/errors"

	"github.com/dct-tournaments/gamer-stats-api/pkg/leagueoflegends"
)

//go:generate mockgen -source=../../internal/leagueoflegends/service.go -destination=../../pkg/mocks/leagueoflegends.go -package=mocks
type LeagueOfLegendsAPIService interface {
	GetSummonerByName(
		ctx context.Context,
		platformRouting leagueoflegends.PlatformRouting,
		username string,
	) (*leagueoflegends.Summoner, error)
	GetMatchesByPUUID(
		ctx context.Context,
		platformRouting leagueoflegends.PlatformRouting,
		puuid string,
		startTime *int64,
		endTime *int64,
		start *int,
		count *int,
	) ([]string, error)
	GetMatchByID(
		ctx context.Context,
		platformRouting leagueoflegends.PlatformRouting,
		matchID string,
	) (*leagueoflegends.Match, error)
	GetAccountByRiotID(
		ctx context.Context,
		gameName string,
		tagLine string,
		regionalRounting leagueoflegends.PlatformRouting,
	) (*leagueoflegends.RiotAccount, error)
}

type service struct {
	leagueOfLegendsAPIService LeagueOfLegendsAPIService
}

type Service interface {
	GetPlayerStats(
		ctx context.Context,
		region string,
		name string,
		tagLine string,
		startAt *int64,
		queueType *leagueoflegends.QueueID,
	) (*PlayerStats, error)
	GetPlayerStatsByPUUID(
		ctx context.Context,
		region string,
		puuid string,
		startAt *int64,
		queueType *leagueoflegends.QueueID,
	) (*PlayerStats, error)
}

func NewService(lolservice LeagueOfLegendsAPIService) Service {
	return &service{
		leagueOfLegendsAPIService: lolservice,
	}
}

func (s *service) getPlayerPUUIDByRiotID(
	ctx context.Context,
	name string,
	tagLine string,
	region string,
) (string, error) {
	account, err := s.leagueOfLegendsAPIService.GetAccountByRiotID(
		ctx,
		name,
		tagLine,
		leagueoflegends.PlatformRouting(region),
	)
	if err != nil {
		return "", err
	}

	return account.Puuid, nil
}

func (s *service) GetPlayerStats(
	ctx context.Context,
	region string,
	name string,
	tagLine string,
	startAt *int64,
	queueType *leagueoflegends.QueueID,
) (*PlayerStats, error) {
	puuid, err := s.getPlayerPUUIDByRiotID(ctx, name, tagLine, region)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get player PUUID by riot id")
	}

	matchIDs, err := s.leagueOfLegendsAPIService.GetMatchesByPUUID(
		ctx,
		leagueoflegends.PlatformRouting(region),
		puuid,
		startAt,
		nil,
		nil,
		nil,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get player matches by PUUID")
	}

	playerKillsCount := 0
	playerDeathCount := 0
	playerAssistCount := 0
	playerWardsPlacedCount := 0

	for _, id := range matchIDs {
		match, err := s.leagueOfLegendsAPIService.GetMatchByID(ctx, leagueoflegends.PlatformRouting(region), id)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get match by ID")
		}

		if queueType != nil && match.Info.QueueID != *queueType {
			continue
		}

		for _, player := range match.Info.Participants {
			if strings.EqualFold(player.PUUID, puuid) {
				playerKillsCount += player.Kills
				playerDeathCount += player.Deaths
				playerAssistCount += player.Assists
				playerWardsPlacedCount += player.WardsPlaced
			}
		}
	}

	return &PlayerStats{
		KillCount:   playerKillsCount,
		DeathCount:  playerDeathCount,
		AssistCount: playerAssistCount,
		WardsPlaced: playerWardsPlacedCount,
	}, nil
}

func (s *service) GetPlayerStatsByPUUID(
	ctx context.Context,
	region string,
	puuid string,
	startAt *int64,
	queueType *leagueoflegends.QueueID,
) (*PlayerStats, error) {
	matchIDs, err := s.leagueOfLegendsAPIService.GetMatchesByPUUID(
		ctx,
		leagueoflegends.PlatformRouting(region),
		puuid,
		startAt,
		nil,
		nil,
		nil,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get player matches by PUUID")
	}

	playerKillsCount := 0
	playerDeathCount := 0
	playerAssistCount := 0
	playerWardsPlacedCount := 0

	for _, id := range matchIDs {
		match, err := s.leagueOfLegendsAPIService.GetMatchByID(ctx, leagueoflegends.PlatformRouting(region), id)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get match by ID")
		}

		if queueType != nil && match.Info.QueueID != *queueType {
			continue
		}

		for _, player := range match.Info.Participants {
			if strings.EqualFold(player.PUUID, puuid) {
				playerKillsCount += player.Kills
				playerDeathCount += player.Deaths
				playerAssistCount += player.Assists
				playerWardsPlacedCount += player.WardsPlaced
			}
		}
	}

	return &PlayerStats{
		KillCount:   playerKillsCount,
		DeathCount:  playerDeathCount,
		AssistCount: playerAssistCount,
		WardsPlaced: playerWardsPlacedCount,
	}, nil
}
