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
	riotAccountV1Path = "https://%s.api.riotgames.com/riot/account/v1/accounts/by-riot-id/%s/%s"
)

type RiotAccount struct {
	Puuid    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

func (s service) GetAccountByRiotID(
	ctx context.Context,
	gameName string,
	tagLine string,
	platformRouting PlatformRouting,
) (*RiotAccount, error) {
	regionalRouting := s.fromPlatformRoutingToRegionalRouting(platformRouting)

	reqPath := fmt.Sprintf(riotAccountV1Path, regionalRouting, gameName, tagLine)

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

	var account RiotAccount

	if err = json.NewDecoder(resp.Body).Decode(&account); err != nil {
		return nil, err
	}

	return &account, nil
}
