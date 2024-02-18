package api

import (
	"encoding/json"
	"fmt"

	"github.com/itsbradn/govalo/pkg/http"
)

type OwnedItemsResponseBody struct {
	EntitlementsByTypes []struct {
		ItemTypeID   string `json:"ItemTypeID"`
		Entitlements []struct {
			TypeID     string `json:"TypeID"`
			ItemID     string `json:"ItemID"`
			InstanceID string `json:"InstanceID"`
		} `json:"Entitlements"`
	} `json:"EntitlementsByTypes"`
}

func GetOwnedItems(shard, puuid, itemType string) (*OwnedItemsResponseBody, error) {
	res, err := http.SendRequest("GET", fmt.Sprintf("https://pd.%s.a.pvp.net/store/v1/entitlements/%s/%s", shard, puuid, itemType), nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var body *OwnedItemsResponseBody
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
