package govalo

import (
	"github.com/itsbradn/govalo/pkg/api"
	"github.com/itsbradn/govalo/pkg/auth"
	"github.com/itsbradn/govalo/pkg/conf"
)

func Setup(region, username, password string) (string, error) {
	platform, err := conf.GetClientPlatformEncoded()
	if err != nil {
		return "", err
	}

	conf.AddAuthHeaders("X-Riot-ClientPlatform", platform)

	cookie, err := auth.GetAuthCookies()
	if err != nil {
		return "", err
	}

	conf.AddAuthHeaders("Cookie", cookie)

	tokens, cookie, err := auth.Login(username, password)
	if err != nil {
		return "", err
	}

	conf.AddAuthHeaders("Cookie", cookie)
	conf.AddAuthHeaders("Authorization", "Bearer " + tokens.AccessToken)
	
	token, err := auth.GetEntitlementToken()
	if err != nil {
		return "", err
	}

	conf.AddAuthHeaders("X-Riot-Entitlements-JWT", token)
	conf.AddAuthHeaders("X-Riot-ClientVersion", "release-08.02-shipping-9-2265102")

	entitlements, err := auth.GetEntitlementsFromToken(token)
	if err != nil {
		return "", err
	}

	name, err := api.GetNameService(region, entitlements.Subject)
	if err != nil {
		return "", err
	}

	return name.GameName + "#" + name.TagLine, nil
}