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
	summonerV4Path = "https://%s.api.riotgames.com/lol/summoner/v4/summoners/by-name/%s"
)

type Summoner struct {
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
	Puuid         string `json:"puuid"`
	Name          string `json:"name"`
	ProfileIconID int64  `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	SummonerLevel int64  `json:"summonerLevel"`
}

func (s service) GetSummonerByName(ctx context.Context, region PlatformRouting, username string) (*Summoner, error) {
	reqPath := fmt.Sprintf(summonerV4Path, region, username)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqPath, nil)
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

	var summoner Summoner
	if err = json.NewDecoder(resp.Body).Decode(&summoner); err != nil {
		return nil, err
	}

	return &summoner, nil
}
