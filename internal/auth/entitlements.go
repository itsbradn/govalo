package auth

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/itsbradn/govalo/pkg/http"
)

type Entitlements struct {
	Entitlements []string `json:"entitlements"`
	Hash         string   `json:"at_hash"`
	Subject      string   `json:"sub"`
	Issuer       string   `json:"iss"`
	IssuedAt     int      `json:"iat"`
	JTI          string   `json:"jti"`
}

type EntitlementsResponseBody struct {
	Token string `json:"entitlements_token"`
}

func GetEntitlementToken() (string, error) {
	res, err := http.SendRequest("POST", "https://entitlements.auth.riotgames.com/api/token/v1", nil)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body := new(EntitlementsResponseBody)
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return "", err
	}

	return body.Token, nil
}

func GetEntitlementsFromToken(entitlement string) (*Entitlements, error) {
	token := strings.Split(entitlement, ".")[1]

	decoded, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}

	body := new(Entitlements)
	err = json.Unmarshal(decoded, &body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
