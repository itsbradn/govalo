package api

import (
	"encoding/json"
	"fmt"

	"github.com/itsbradn/govalo/pkg/http"
)

type PricesResponseBody struct {
	Offers []struct {
		OfferID          string            `json:"OfferID"`
		IsDirectPurchase bool              `json:"IsDirectPurchase"`
		StartDate        string            `json:"StartDate"` // ISO 1806 Date
		Cost             map[string]uint32 `json:"Cost"`
		Rewards          []struct {
			ItemTypeID string `json:"ItemTypeID"`
			ItemID     string `json:"ItemID"`
			Quantity   uint8  `json:"Quantity"`
		} `json:"Rewards"`
	} `json:"Offers"`
}

func GetStorePrices(shard string) (*PricesResponseBody, error) {
	res, err := http.SendRequest("GET", fmt.Sprintf("https://pd.%s.a.pvp.net/store/v1/offers", shard), nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var body *PricesResponseBody
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
