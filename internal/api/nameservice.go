package api

import (
	"encoding/json"
	"fmt"

	"github.com/itsbradn/govalo/pkg/http"
)

type NameServiceRequestBody []string

type NameServiceResponseBody struct {
	DisplayName string `json:"DisplayName"`
	Subject     string `json:"Subject"`
	GameName    string `json:"GameName"`
	TagLine     string `json:"TagLine"`
}

func GetNameService(shard, subject string) (*NameServiceResponseBody, error) {
	res, err := http.SendRequest("PUT", fmt.Sprintf("https://pd.%s.a.pvp.net/name-service/v2/players", shard), NameServiceRequestBody{subject})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var body []*NameServiceResponseBody
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return body[0], nil
}
