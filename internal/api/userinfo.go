package api

import (
	"encoding/json"

	"github.com/itsbradn/govalo/pkg/http"
)

type UserInfoResponseBody struct {
	Country       string `json:"country"`
	PlayerUUID    string `json:"sub"`
	EmailVerified bool   `json:"email_verified"`
	CountryAt     uint64 `json:"country_at"`
	Password      struct {
		ChangedAt uint64 `json:"cng_at"`
		BeenReset bool   `json:"reset"`
		MustReset bool   `json:"must_reset"`
	} `json:"pw"`
	PhoneNumberVerified        bool     `json:"phone_number_verified"`
	AccountVerified            bool     `json:"account_verified"`
	FederatedIdentityProviders []string `json:"federated_identity_providers"`
	Locale                     string   `json:"player_locale"`
	Account                    struct {
		Type      uint16 `json:"type"`
		State     string `json:"state"`
		Adm       bool   `json:"adm"`
		GameName  string `json:"game_name"`
		TagLine   string `json:"tag_line"`
		CreatedAt uint64 `json:"created_at"`
	} `json:"acct"`
	Age uint8 `json:"age"`
	JTI string `json:"jti"`
}

func GetUserInfo() (*UserInfoResponseBody, error) {
	res, err := http.SendRequest("GET", "https://auth.riotgames.com/userinfo", nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var body *UserInfoResponseBody
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
