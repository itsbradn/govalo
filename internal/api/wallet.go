package api

import (
	"encoding/json"
	"fmt"

	"github.com/itsbradn/govalo/pkg/http"
)

type WalletResponseBody struct {
	Balances map[string]uint32 `json:"Balances"`
}

func GetWallet(shard, puuid string) (*WalletResponseBody, error) {
	res, err := http.SendRequest("GET", fmt.Sprintf("https://pd.%s.a.pvp.net/store/v1/wallet/%s", shard, puuid), nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var body *WalletResponseBody
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
