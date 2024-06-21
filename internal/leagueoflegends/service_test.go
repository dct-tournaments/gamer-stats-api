package leagueoflegends_test

import (
	"context"
	"testing"

	"github.com/dct-tournaments/gamer-stats-api/internal/leagueoflegends"
	plol "github.com/dct-tournaments/gamer-stats-api/pkg/leagueoflegends"
	"github.com/dct-tournaments/gamer-stats-api/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_GetPlayerStats_RankedQueue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // will assert if all methods expeted to be called, were called

	mockLeagueOfLegendsAPIService := mocks.NewMockLeagueOfLegendsAPIService(ctrl)

	ctx := context.Background()

	// mock get account by riot id
	playerName := "player"
	playerPuuid := "123"
	tagLine := "euw"
	region := "EUW1"

	expectedRiotAccount := &plol.RiotAccount{Puuid: playerPuuid, GameName: playerName, TagLine: tagLine}

	mockLeagueOfLegendsAPIService.EXPECT().
		GetAccountByRiotID(ctx, playerName, tagLine, plol.PlatformRouting(region)).
		Return(expectedRiotAccount, nil).
		Times(1)

	// mock get matches by puuid
	expectedMatchesID := []string{"match1", "match2"}
	mockLeagueOfLegendsAPIService.EXPECT().
		GetMatchesByPUUID(ctx, plol.EUW1PlatformRouting, expectedRiotAccount.Puuid, nil, nil, nil, nil).
		Return(expectedMatchesID, nil).
		Times(1)

	// mock get match by id
	expectedMatch1 := &plol.Match{
		Info: plol.Info{
			Participants: []plol.Participants{
				{
					PUUID:   playerPuuid,
					Kills:   10,
					Deaths:  0,
					Assists: 5,
				},
			},
			QueueID: plol.RankedQueueID,
		},
	}

	expectedMatch2 := &plol.Match{
		Info: plol.Info{
			Participants: []plol.Participants{
				{
					PUUID:   playerPuuid,
					Kills:   20,
					Deaths:  3,
					Assists: 7,
				},
			},
			QueueID: plol.FlexQueueID,
		},
	}

	mockLeagueOfLegendsAPIService.EXPECT().
		GetMatchByID(ctx, plol.EUW1PlatformRouting, expectedMatchesID[0]).
		Return(expectedMatch1, nil).
		Times(1)

	mockLeagueOfLegendsAPIService.EXPECT().
		GetMatchByID(ctx, plol.EUW1PlatformRouting, expectedMatchesID[1]).
		Return(expectedMatch2, nil).
		Times(1)

	service := leagueoflegends.NewService(mockLeagueOfLegendsAPIService)

	rankedGames := plol.RankedQueueID

	stats, err := service.GetPlayerStats(ctx, region, playerName, tagLine, nil, &rankedGames)

	assert.Nil(t, err)
	assert.Equal(t, 10, stats.KillCount)
	assert.Equal(t, 0, stats.DeathCount)
	assert.Equal(t, 5, stats.AssistCount)
}
