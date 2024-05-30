package rest

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Stats struct {
	Kills   int `json:"kills"`
	Deaths  int `json:"deaths"`
	Assists int `json:"assists"`
}

type GetLeagueOfLegendsStatsResponse struct {
	Stats           Stats `json:"stats"`
	QueryExecutedAt int64 `json:"query_executed_at"`
}

// @ID           statsAPIV0GetLeagueOfLegendsStats
// @Summary      Get League of Legends stats for a player.
// @Description  Get League of Legends stats for a player.
// @Tags         stats-api/v0
// @Accept       json
// @Produce      json
// @query        username	query   string  true  "league of legends username"
// @query		 tagline	query   string  true "league of legends tagline"
// @Param        start_time	query   string  false "start time epoch format"
// @Param        region		query   string  true  "league of legends region: br1,eun1,euw1,jp1,kr,la1,la2,na1,oc1,ph2,ru,sg2,th2,tr1,tw2,vn2"
// @Success      200  {object}     GetLeagueOfLegendsStatsResponse
// @Failure      400
// @Failure      500
// @Router       /stats-api/v0/league-of-legends [get].
func (h *handler) GetLeagueOfLegendsStats(c *gin.Context) {
	username := c.Query("username")
	tagline := c.Query("tagline")
	startTimeStr := c.Query("start_time")
	region := c.Query("region")

	if username == "" || region == "" {
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

	stats, err := h.leagueOfLegendsService.GetPlayerStats(c.Request.Context(), username, tagline, region, startTimePtr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	response := GetLeagueOfLegendsStatsResponse{
		Stats:           Stats{Kills: stats.KillCount, Deaths: stats.DeathCount, Assists: stats.AssistCount},
		QueryExecutedAt: time.Now().Unix(),
	}

	c.JSON(200, response)
}
