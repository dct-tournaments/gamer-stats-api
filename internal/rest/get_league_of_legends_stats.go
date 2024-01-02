package rest

import (
	"net/http"
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
// @Param        username	query   string  true  "league of legends username"
// @Param        start_time	query   string  false "start time in RFC3339 format"
// @Param        region		query   string  true  "league of legends region: br1,eun1,euw1,jp1,kr,la1,la2,na1,oc1,ph2,ru,sg2,th2,tr1,tw2,vn2"
// @Success      200  {object}     GetLeagueOfLegendsStatsResponse
// @Failure      400
// @Failure      500
// @Router       /stats-api/v0/league-of-legends [get].
func (h *handler) GetLeagueOfLegendsStats(c *gin.Context) {
	username := c.Query("username")
	startTimeStr := c.Query("start_time")
	region := c.Query("region")

	if username == "" || region == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})

		return
	}

	var startTimePtr *int64

	if startTimeStr != "" {
		startTime, err := time.Parse(time.RFC3339, startTimeStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid start time"})

			return
		}

		startTimeInt := startTime.Unix()
		startTimePtr = &startTimeInt
	}

	stats, err := h.leagueOfLegendsService.GetPlayerStats(c.Request.Context(), region, username, startTimePtr)
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
