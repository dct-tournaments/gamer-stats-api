package rest

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// @ID           statsAPIV1GetLeagueOfLegendsStats
// @Summary      Get League of Legends stats for a player using PUUID.
// @Description  Get League of Legends stats for a player using PUUID.
// @Tags         stats-api/v0
// @Accept       json
// @Produce      json
// @query        puuid	query   string  true  "league of legends player puuid"
// @Param        start_time	query   string  false "start time epoch format"
// @Param        region		query   string  true  "league of legends region: br1,eun1,euw1,jp1,kr,la1,la2,na1,oc1,ph2,ru,sg2,th2,tr1,tw2,vn2"
// @Param        queue_type	query   string  true  "queue type: all,ranked"
// @Success      200  {object}     GetLeagueOfLegendsStatsResponse
// @Failure      400
// @Failure      500
// @Router       /stats-api/v1/league-of-legends [get].
func (h *handler) GetLeagueOfLegendsStatsV1(c *gin.Context) {
	puuid := c.Query("puuid")
	region := c.Query("region")
	startTimeStr := c.Query("start_time")
	queueTypeStr := c.Query("queue_type")

	if puuid == "" || queueTypeStr == "" || region == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})

		return
	}

	var startTimePtr *int64

	if startTimeStr != "" {
		// convert a string to int64
		startTime, err := strconv.ParseInt(startTimeStr, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid start time"})

			return
		}

		startTimePtr = &startTime
	}

	queueID := toQueueTypeParamToQueueID(QueueType(queueTypeStr))

	stats, err := h.leagueOfLegendsService.GetPlayerStatsByPUUID(
		c.Request.Context(),
		region,
		puuid,
		startTimePtr,
		queueID,
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	response := GetLeagueOfLegendsStatsResponse{
		Stats: Stats{
			Kills:       stats.KillCount,
			Deaths:      stats.DeathCount,
			Assists:     stats.AssistCount,
			WardsPlaced: stats.WardsPlaced,
		},
		QueryExecutedAt: time.Now().Unix(),
	}

	c.JSON(200, response)
}
