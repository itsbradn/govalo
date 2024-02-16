package auth

import (
	"encoding/json"

	"github.com/itsbradn/govalo/pkg/http"
)

type AuthCookiesRequestBody struct {
	ClientID     string `json:"client_id"`
	Nonce        string `json:"nonce"`
	RedirectUri  string `json:"redirect_uri"`
	ResponseType string `json:"response_type"`
	Scope        string `json:"scope"`
}

type AuthCookiesResponseBody struct {
	Type    string `json:"type"`
	Country string `json:"country"`
}

func GetAuthCookies() (string, error) {
	res, err := http.SendRequest("POST", "https://auth.riotgames.com/api/v1/authorization", AuthCookiesRequestBody{
		ClientID:     "play-valorant-web-prod",
		Nonce:        "1",
		RedirectUri:  "https://playvalorant.com/opt_in",
		ResponseType: "token id_token",
		Scope:        "openid account",
	})

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body := new(AuthCookiesResponseBody)
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return "", err
	}

	cookie, err := http.ParseCookies(res.Header["Set-Cookie"], "asid")
	if err != nil {
		return "", err
	}

	return cookie, nil
}
