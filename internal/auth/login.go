package auth

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/itsbradn/govalo/pkg/http"
)

type LoginRequestBody struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponseBody struct {
	Type     string                    `json:"type"`
	Response LoginResponseBodyResponse `json:"response"`
	Country  string                    `json:"country"`
}

type LoginResponseBodyResponse struct {
	Mode       string `json:"mode"`
	Parameters struct {
		Uri string `json:"uri"`
	} `json:"parameters"`
}
type LoginUriTokens struct {
	AccessToken string `json:"access_token"`
	IdToken     string `json:"id_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func getLoginTokens(uri string) (*LoginUriTokens, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	q, err := url.ParseQuery(u.Fragment)
	if err != nil {
		return nil, err
	}

	accessToken := q.Get("access_token")
	idToken := q.Get("id_token")

	expiresIn, err := strconv.Atoi(q.Get("expires_in"))
	if err != nil {
		return nil, err
	}

	return &LoginUriTokens{
		AccessToken: accessToken,
		IdToken:     idToken,
		ExpiresIn:   expiresIn,
	}, nil
}

func Login(username, password string) (*LoginUriTokens, string, error) {
	res, err := http.SendRequest("PUT", "https://auth.riotgames.com/api/v1/authorization", LoginRequestBody{
		Type: "auth",
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, "", err
	}
	defer res.Body.Close()

	body := new(LoginResponseBody)
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return nil, "", err
	}

	tokens, err := getLoginTokens(body.Response.Parameters.Uri)
	if err != nil {
		return nil, "", err
	}

	cookie, err := http.ParseCookies(res.Header["Set-Cookie"], "ssid")
	if err != nil {
		return nil, "", err
	}

	return tokens, cookie, nil
}