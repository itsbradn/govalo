package govalo

import (
	"github.com/itsbradn/govalo/internal/api"
	"github.com/itsbradn/govalo/internal/auth"
	"github.com/itsbradn/govalo/pkg/conf"
)

type GoValoAPI struct {
	Region string `json:"region"`
	PUUID  string
}

func (vapi *GoValoAPI) GetUserInfo() (*api.UserInfoResponseBody, error) {
	return api.GetUserInfo()
}

func (vapi *GoValoAPI) GetSelfNameService() (*api.NameServiceResponseBody, error) {
	return api.GetNameService(vapi.Region, vapi.PUUID)
}

func (vapi *GoValoAPI) GetNameService(uuid string) (*api.NameServiceResponseBody, error) {
	return api.GetNameService(vapi.Region, uuid)
}

type MatchHistoryOptions = api.MatchHistoryOptions

func (vapi *GoValoAPI) GetMatchHistory(uuid string, options *MatchHistoryOptions) (*api.MatchHistoryResponseBody, error) {
	return api.GetMatchHistory(vapi.Region, uuid, options)
}

func (vapi *GoValoAPI) GetMatchDetails(uuid string) (*api.MatchDetailsResponseBody, error) {
	return api.GetMatchDetails(vapi.Region, uuid)
}

type CompetitiveUpdatesOptions = api.CompetitiveUpdatesOptions

func (vapi *GoValoAPI) GetCompetitiveUpdates(uuid string, options *CompetitiveUpdatesOptions) (*api.CompetitiveUpdatesResponseBody, error) {
	return api.GetCompetitiveUpdates(vapi.Region, uuid, options)
}

func Setup(region, username, password string) (*GoValoAPI, error) {
	platform, err := conf.GetClientPlatformEncoded()
	if err != nil {
		return nil, err
	}

	conf.AddAuthHeaders("X-Riot-ClientPlatform", platform)

	cookie, err := auth.GetAuthCookies()
	if err != nil {
		return nil, err
	}

	conf.AddAuthHeaders("Cookie", cookie)

	tokens, cookie, err := auth.Login(username, password)
	if err != nil {
		return nil, err
	}

	conf.AddAuthHeaders("Cookie", cookie)
	conf.AddAuthHeaders("Authorization", "Bearer "+tokens.AccessToken)

	token, err := auth.GetEntitlementToken()
	if err != nil {
		return nil, err
	}

	conf.AddAuthHeaders("X-Riot-Entitlements-JWT", token)
	conf.AddAuthHeaders("X-Riot-ClientVersion", "release-08.02-shipping-9-2265102")

	userinfo, err := api.GetUserInfo()
	if err != nil {
		return nil, err
	}

	return &GoValoAPI{
		PUUID:  userinfo.PlayerUUID,
		Region: region,
	}, nil
}
