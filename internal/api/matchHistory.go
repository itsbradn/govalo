package api

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/itsbradn/govalo/pkg/http"
)

type MatchHistoryResponseBody struct {
	PUUID         string `json:"Subject"`
	StartingIndex uint16 `json:"BeginIndex"`
	EndingIndex   uint16 `json:"EndIndex"`
	Total         uint16 `json:"Total"`
	History       []struct {
		MatchID       string `json:"MatchID"`
		GameStartTime uint64 `json:"GameStartTime"`
		QueueID       string `json:"QueueID"`
	} `json:"History"`
}

type MatchHistoryOptions struct {
	StartIndex uint16
	EndIndex   uint16
	Queue      string
}

func GetMatchHistory(shard, puuid string, options *MatchHistoryOptions) (*MatchHistoryResponseBody, error) {
	uri := fmt.Sprintf("https://pd.%v.a.pvp.net/match-history/v1/history/%v", shard, puuid)
	params := url.Values{}

	if options != nil {
		if options.StartIndex > 0 {
			params.Set("startIndex", strconv.FormatInt(int64(options.StartIndex), 10))
		}

		if options.EndIndex > 0 {
			params.Set("endIndex", strconv.FormatInt(int64(options.EndIndex), 10))
		}

		if options.Queue != "" && options.Queue != "ANY" {
			params.Set("queue", options.Queue)
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

	var body *MatchHistoryResponseBody
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
