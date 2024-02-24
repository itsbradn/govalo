package api

import (
	"encoding/json"
	"fmt"

	"github.com/itsbradn/govalo/pkg/http"
)

type AccountExperienceResponseBody struct {
	Version  uint16 `json:"Version"`
	PUUID    string `json:"Subject"`
	Progress struct {
		Level uint32 `json:"Level"`
		XP    uint64 `json:"XP"`
	} `json:"Progress"`
	History []struct {
		MatchID       string `json:"ID"`
		MatchStart    string `json:"MatchStart"` // ISO 8601 Date
		StartProgress struct {
			Level uint32 `json:"Level"`
			XP    uint64 `json:"XP"`
		} `json:"StartProgress"`
		EndProgress struct {
			Level uint32 `json:"Level"`
			XP    uint64 `json:"XP"`
		} `json:"EndProgress"`
		XPDelta   uint32 `json:"XPDelta"`
		XPSources struct {
			ID     string `json:"ID"`
			Amount uint32 `json:"Amount"`
		} `json:"XPSources"`
	} `json:"History"`
	LastGrantedFirstWin   string `json:"LastTimeGrantedFirstWin"`   // ISO 8601 Date
	NextFirstWinAvailable string `json:"NextTimeFirstWinAvailable"` // ISO 8601 Date
}

func GetAccountXP(shard, puuid string) (*AccountExperienceResponseBody, error) {
	res, err := http.SendRequest("GET", fmt.Sprintf("https://pd.%s.a.pvp.net/account-xp/v1/players/%s", shard, puuid), nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var body *AccountExperienceResponseBody
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
