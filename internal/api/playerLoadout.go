package api

import (
	"encoding/json"
	"fmt"

	"github.com/itsbradn/govalo/pkg/http"
)

type PlayerLoadoutResponseBody struct {
	PUUID string `json:"Subject"`
	Version uint16 `json:"Version"`
	Guns []struct{
		ID string `json:"ID"`
		CharmInstanceID string `json:"CharmInstanceID"`
		CharmID string `json:"CharmID"`
		CharmLevelID string `json:"CharmLevelID"`
		SkinID string `json:"SkinID"`
		SkinLevelID string `json:"SkinLevelID"`
		ChromaID string `json:"ChromaID"`
	} `json:"Guns"`
	Sprays []struct{
		EquipSlotID string `json:"EquipSlotID"`
		SprayID string `json:"SprayID"`
	} `json:"Sprays"`
	Identity struct{
		PlayerCardID string `json:"PlayerCardID"`
		PlayerTitleID string `json:"PlayerTitleID"`
		AccountLevel uint32 `json:"AccountLevel"`
		PreferredLevelBorderID string `json:"PreferredLevelBorderID"`
		HideAccountLevel bool `json:"HideAccountLevel"`
	} `json:"Identity"`
	Incognito bool `json:"Incognito"`
}

func GetPlayerLoadout(shard, puuid string) (*PlayerLoadoutResponseBody, error) {
	res, err := http.SendRequest("GET", fmt.Sprintf("https://pd.%s.a.pvp.net/personalization/v2/players/%s/playerloadout", shard, puuid), nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var body *PlayerLoadoutResponseBody
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return body, nil
}