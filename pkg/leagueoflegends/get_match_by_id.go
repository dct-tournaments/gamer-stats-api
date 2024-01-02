package leagueoflegends

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	getMatchByIDPath = "https://%s.api.riotgames.com/lol/match/v5/matches/%s"
)

type Match struct {
	Info Info `json:"info"`
}

type Info struct {
	Participants []Participants `json:"participants"`
}

type Participants struct {
	Kills  int    `json:"kills"`
	Deaths int    `json:"deaths"`
	PUUID  string `json:"puuid"`
}

func (s service) GetMatchByID(ctx context.Context, platformRouting PlatformRouting, matchID string) (*Match, error) {
	region := s.fromPlatformRoutingToRegionalRouting(platformRouting)

	baseURL := fmt.Sprintf(getMatchByIDPath, region, matchID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Riot-Token", s.config.APIKey)

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

	var match Match

	if err = json.NewDecoder(resp.Body).Decode(&match); err != nil {
		return nil, err
	}

	return &match, nil
}
