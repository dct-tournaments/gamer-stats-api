package leagueoflegends

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const (
	getMatchesByPUUIDPath = "https://%s.api.riotgames.com/lol/match/v5/matches/by-puuid/%s/ids"
)

func (s service) GetMatchesByPUUID(
	ctx context.Context,
	region RegionalRouting,
	puuid string,
	startTime *int64,
	endTime *int64,
	start *int,
	count *int,
) ([]string, error) {
	baseURL := fmt.Sprintf(getMatchesByPUUIDPath, region, puuid)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Riot-Token", s.config.APIKey)

	query := req.URL.Query()

	if startTime != nil && *startTime > 0 {
		strStartTime := strconv.FormatInt(*startTime, 10)
		query.Add("startTime", strStartTime)
	}

	if endTime != nil && *endTime > 0 {
		strEndTime := strconv.FormatInt(*endTime, 10)
		query.Add("endTime", strEndTime)
	}

	if start != nil && *start > 0 {
		strStart := strconv.Itoa(*start)
		query.Add("start", strStart)
	}

	if count != nil && *count > 0 {
		strCount := strconv.Itoa(*count)
		query.Add("count", strCount)
	}

	req.URL.RawQuery = query.Encode()

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		buf := new(strings.Builder)

		if _, err := io.Copy(buf, resp.Body); err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, buf.String())
	}

	var matchIDs []string

	if err = json.NewDecoder(resp.Body).Decode(&matchIDs); err != nil {
		return nil, err
	}

	return matchIDs, nil
}
