package api

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/itsbradn/govalo/pkg/http"
)

type CompetitiveUpdatesResponseBody struct {
	Version uint16 `json:"Version"`
	PUUID   string `json:"Subject"`
	Matches []struct {
		MatchID                      string `json:"MatchID"`
		MapID                        string `json:"MapID"`
		SeasonID                     string `json:"SeasonID"`
		MatchStartTime               uint64 `json:"MatchStartTime"`
		TierAfterUpdate              uint8  `json:"TierAfterUpdate"`
		TierBeforeUpdate             uint8  `json:"TierBeforeUpdate"`
		RankedRatingAfterUpdate      uint8  `json:"RankedRatingAfterUpdate"`
		RankedRatingBeforeUpdate     uint8  `json:"RankedRatingBeforeUpdate"`
		RankedRatingEarned           int8   `json:"RankedRatingEarned"`
		RankedRatingPerformanceBonus uint8  `json:"RankedRatingPerformanceBonus"`
		CompetitiveMovement          string `json:"CompetitiveMovement"`
		AFKPenalty                   int8   `json:"AFKPenalty"`
	} `json:"Matches"`
}

type CompetitiveUpdatesOptions struct {
	StartIndex uint16
	EndIndex   uint16
}

func GetCompetitiveUpdates(shard, puuid string, options *CompetitiveUpdatesOptions) (*CompetitiveUpdatesResponseBody, error) {
	uri := fmt.Sprintf("https://pd.%v.a.pvp.net/mmr/v1/players/%v/competitiveupdates", shard, puuid)
	params := url.Values{}

	if options != nil {
		if options.StartIndex > 0 {
			params.Set("startIndex", strconv.FormatInt(int64(options.StartIndex), 10))
		}

		if options.EndIndex > 0 {
			params.Set("endIndex", strconv.FormatInt(int64(options.EndIndex), 10))
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

	var body *CompetitiveUpdatesResponseBody
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
