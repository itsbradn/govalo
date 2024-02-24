package api

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/itsbradn/govalo/pkg/http"
)

type LeaderboardResponseBody struct {
	Deployment string `json:"Deployment"`
	QueueID    string `json:"QueueID"`
	SeasonID   string `json:"SeasonID"`
	Players    struct {
		PlayerCardID    string `json:"PlayerCardID"`
		TitleID         string `json:"TitleID"`
		IsBanned        bool   `json:"IsBanned"`
		IsAnonymized    bool   `json:"IsAnonymized"`
		PUUID           string `json:"puuid"`
		GameName        string `json:"gameName"`
		TagLine         string `json:"tagLine"`
		LeaderboardRank uint16 `json:"leaderboardRank"`
		RankedRating    uint16 `json:"rankedRating"`
		Wins            uint16 `json:"numberOfWins"`
		Tier            uint8  `json:"competitiveTier"`
	} `json:"Players"`
	TotalPlayers          uint16 `json:"totalPlayers"`
	ImmortalStartingPage  uint16 `json:"immortalStartingPage"`
	ImmortalStartingIndex uint16 `json:"immortalStartingIndex"`
	TopTierRRThreshold    uint16 `json:"topTierRRThreshold"`
	TierDetails           map[string]struct {
		RankedRatingThreshold uint16 `json:"rankedRatingThreshold"`
		StartingIndex         uint16 `json:"startingIndex"`
	} `json:"tierDetails"`
	StartIndex uint16 `json:"startIndex"`
	Query      string `json:"query"`
}

type LeaderboardOptions struct {
	StartIndex uint16
	Size       uint16
	Query      string
}

func GetLeaderboard(shard, seasonID string, options *LeaderboardOptions) (*LeaderboardResponseBody, error) {
	uri := fmt.Sprintf("https://pd.%s.a.pvp.net/mmr/v1/leaderboards/affinity/na/queue/competitive/season/%s", shard, seasonID)
	params := url.Values{}

	if options != nil {
		if options.StartIndex > 0 {
			params.Set("startIndex", strconv.FormatInt(int64(options.StartIndex), 10))
		}

		if options.Size > 0 {
			params.Set("size", strconv.FormatInt(int64(options.Size), 10))
		}

		if options.Query != "" {
			params.Set("query", options.Query)
		}
	}
	encodedParams := params.Encode()
	if encodedParams != "" {
		uri = fmt.Sprintf("%v?%v", uri, encodedParams)
	}
	res, err := http.SendRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var body *LeaderboardResponseBody
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
