package govalo

import (
	"github.com/itsbradn/govalo/internal/api"
	"github.com/itsbradn/govalo/internal/auth"
	"github.com/itsbradn/govalo/pkg/conf"
)

const (
	VALORANT_REGION_NA    string = "na"
	VALORANT_REGION_LATAM string = "latam"
	VALORANT_REGION_BR    string = "br"
	VALORANT_REGION_EU    string = "eu"
	VALORANT_REGION_AP    string = "ap"
	VALORANT_REGION_KR    string = "kr"
)

const (
	VALORANT_SHARD_NA  string = "na"
	VALORANT_SHARD_PBE string = "pbe"
	VALORANT_SHARD_EU  string = "eu"
	VALORANT_SHARD_AP  string = "ap"
	VALORANT_SHARD_KR  string = "kr"
)

func regionToShard(region string) string {
	switch region {
	case "na":
		return VALORANT_SHARD_NA
	case "latam":
		return VALORANT_SHARD_NA
	case "br":
		return VALORANT_SHARD_NA
	case "eu":
		return VALORANT_SHARD_EU
	case "ap":
		return VALORANT_SHARD_AP
	case "kr":
		return VALORANT_SHARD_KR
	default:
		return VALORANT_SHARD_NA
	}
}

type GoValoAPI struct {
	Region string `json:"region"`
	Shard  string `json:"shard"`
	PUUID  string
}

func (vapi *GoValoAPI) GetUserInfo() (*api.UserInfoResponseBody, error) {
	return api.GetUserInfo()
}

func (vapi *GoValoAPI) GetSelfNameService() (*api.NameServiceResponseBody, error) {
	return api.GetNameService(vapi.Shard, vapi.PUUID)
}

func (vapi *GoValoAPI) GetNameService(uuid string) (*api.NameServiceResponseBody, error) {
	return api.GetNameService(vapi.Shard, uuid)
}

type MatchHistoryOptions = api.MatchHistoryOptions

func (vapi *GoValoAPI) GetMatchHistory(uuid string, options *MatchHistoryOptions) (*api.MatchHistoryResponseBody, error) {
	return api.GetMatchHistory(vapi.Shard, uuid, options)
}

func (vapi *GoValoAPI) GetMatchDetails(uuid string) (*api.MatchDetailsResponseBody, error) {
	return api.GetMatchDetails(vapi.Shard, uuid)
}

type CompetitiveUpdatesOptions = api.CompetitiveUpdatesOptions

func (vapi *GoValoAPI) GetCompetitiveUpdates(uuid string, options *CompetitiveUpdatesOptions) (*api.CompetitiveUpdatesResponseBody, error) {
	return api.GetCompetitiveUpdates(vapi.Shard, uuid, options)
}

func (vapi *GoValoAPI) GetStorefront() (*api.StoreFrontResponseBody, error) {
	return api.GetStorefront(vapi.Shard, vapi.PUUID)
}

func (vapi *GoValoAPI) GetWallet() (*api.WalletResponseBody, error) {
	return api.GetWallet(vapi.Shard, vapi.PUUID)
}

func (vapi *GoValoAPI) GetAccountXP() (*api.AccountExperienceResponseBody, error) {
	return api.GetAccountXP(vapi.Shard, vapi.PUUID)
}

const (
	OWNED_ITEM_TYPE_AGENTS        string = "01bb38e1-da47-4e6a-9b3d-945fe4655707"
	OWNED_ITEM_TYPE_CONTRACTS     string = "f85cb6f7-33e5-4dc8-b609-ec7212301948"
	OWNED_ITEM_TYPE_SPRAYS        string = "d5f120f8-ff8c-4aac-92ea-f2b5acbe9475"
	OWNED_ITEM_TYPE_GUN_BUDDIES   string = "dd3bf334-87f3-40bd-b043-682a57a8dc3a"
	OWNED_ITEM_TYPE_CARDS         string = "3f296c07-64c3-494c-923b-fe692a4fa1bd"
	OWNED_ITEM_TYPE_SKINS         string = "e7c63390-eda7-46e0-bb7a-a6abdacd2433"
	OWNED_ITEM_TYPE_SKIN_VARIANTS string = "3ad1b2b2-acdb-4524-852f-954a76ddae0a"
	OWNED_ITEM_TYPE_TITLES        string = "de7caa6b-adf7-4588-bbd1-143831e786c6"
)

func (vapi *GoValoAPI) GetOwnedItems(itemType string) (*api.OwnedItemsResponseBody, error) {
	return api.GetOwnedItems(vapi.Shard, vapi.PUUID, itemType)
}

func (vapi *GoValoAPI) GetStorePrices() (*api.PricesResponseBody, error) {
	return api.GetStorePrices(vapi.Shard)
}

type LeaderboardOptions = api.LeaderboardOptions

func (vapi *GoValoAPI) GetLeaderboard(seasonID string, options *LeaderboardOptions) (*api.LeaderboardResponseBody, error) {
	return api.GetLeaderboard(vapi.Shard, seasonID, options)
}

func (vapi *GoValoAPI) GetPlayerLoadout() (*api.PlayerLoadoutResponseBody, error) {
	return api.GetPlayerLoadout(vapi.Shard, vapi.PUUID)
}

func (vapi *GoValoAPI) GetPlayerMMR(puuid string) (*api.PlayerMMRResponseBody, error) {
	return api.GetPlayerMMR(vapi.Shard, puuid)
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
		Shard:  regionToShard(region),
	}, nil
}
