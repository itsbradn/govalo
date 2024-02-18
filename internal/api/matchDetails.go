package api

import (
	"encoding/json"
	"fmt"

	"github.com/itsbradn/govalo/pkg/http"
)

type MatchDetailsResponseBody struct {
	MatchInfo struct {
		MatchID                     string            `json:"matchId"`
		MapID                       string            `json:"mapID"`
		GamePodID                   string            `json:"gamePodId"`
		GameLoopZone                string            `json:"gameLoopZone"`
		GameServerAddress           string            `json:"gameServerAddress"`
		GameVersion                 string            `json:"gameVersion"`
		GameLengthMillis            uint32            `json:"gameLengthMillis"`
		GameStartMillis             uint64            `json:"gameStartMillis"`
		ProvisioningFlowID          string            `json:"provisioningFlowID"`
		IsCompleted                 bool              `json:"isCompleted"`
		CustomGameName              string            `json:"customGameName"`
		ForcePostProcessing         bool              `json:"forcePostProcessing"`
		QueueID                     string            `json:"queueID"`
		GameMode                    string            `json:"gameMode"`
		IsRanked                    bool              `json:"isRanked"`
		IsSampled                   bool              `json:"isMatchSampled"`
		SeasonID                    string            `json:"seasonId"`
		CompletionState             string            `json:"completionState"`
		PlatformType                string            `json:"platformType"`
		PremierMatchInfo            struct{}          `json:"premierMatchInfo"`
		PartyRRPenalties            map[string]uint16 `json:"partyRRPenalties"`
		ShouldMatchDisablePenalties bool              `json:"shouldMatchDisablePenalties"`
	} `json:"matchInfo"`
	Players []struct {
		PUUID        string `json:"subject"`
		GameName     string `json:"gameName"`
		TagLine      string `json:"tagLine"`
		PlatformInfo struct {
			PlatformType      string `json:"platformType"`
			PlatformOS        string `json:"platformOS"`
			PlatformOSVersion string `json:"platformOSVersion"`
			PlatformChipset   string `json:"platformChipset"`
		} `json:"platformInfo"`
		TeamID      string `json:"teamId"`
		CharacterID string `json:"characterId"`
		Stats       struct {
			Score          uint16 `json:"score"`
			RoundsPlayed   uint8  `json:"roundsPlayed"`
			Kills          uint8  `json:"kills"`
			Deaths         uint8  `json:"deaths"`
			Assists        uint8  `json:"assists"`
			PlaytimeMillis uint32 `json:"playtimeMillis"`
			AbilityCasts   struct {
				GrenadeCasts  uint8 `json:"grenadeCasts"`
				Ability1Casts uint8 `json:"ability1Casts"`
				Ability2Casts uint8 `json:"ability2Casts"`
				UltimateCasts uint8 `json:"ultimateCasts"`
			} `json:"abilityCasts"`
		} `json:"stats"`
		RoundDamage []struct {
			Round    uint8  `json:"round"`
			Receiver string `json:"receiver"`
			Damage   uint16 `json:"damage"`
		} `json:"roundDamage"`
		CompetitiveTier        uint8  `json:"competitiveTier"`
		IsObserver             bool   `json:"isObserver"`
		PlayerCard             string `json:"playerCard"`
		PlayerTitle            string `json:"playerTitle"`
		PreferredLevelBorder   string `json:"preferredLevelBorder"`
		AccountLevel           uint16 `json:"accountLevel"`
		SessionPlaytimeMinutes uint16 `json:"sessionPlaytimeMinutes"`
		XpModifications        []struct {
			Value float32 `json:"Value"`
			ID    string  `json:"ID"`
		} `json:"xpModifications"`
		BehaviorFactors struct {
			AFKRounds                   uint8   `json:"afkRounds"`
			Collisions                  float32 `json:"collisions"`
			CommsRatingRecovery         uint32  `json:"commsRatingRecovery"`
			DamageParticipationOutgoing uint32  `json:"damageParticipationOutgoing"`
			FriendlyFireIncoming        uint32  `json:"friendlyFireIncoming"`
			FriendlyFireOutgoing        uint32  `json:"friendlyFireOutgoing"`
			MouseMovement               uint32  `json:"mouseMovement"`
			StayedInSpawnRounds         uint32  `json:"stayedInSpawnRounds"`
		} `json:"behaviorFactors"`
	} `json:"players"`
	Bots    []struct{} `json:"bots"`
	Coaches []struct {
		PUUID  string `json:"subject"`
		TeamID string `json:"teamId"`
	} `json:"coaches"`
	Teams []struct {
		TeamID       string `json:"teamId"`
		Won          bool   `json:"won"`
		RoundsPlayed uint8  `json:"roundsPlayed"`
		RoundsWon    uint8  `json:"roundsWon"`
		NumPoints    uint8  `json:"numPoints"`
	} `json:"teams"`
	Rounds []struct {
		Number               uint8  `json:"roundNum"`
		Result               string `json:"roundResult"`
		Ceremony             string `json:"roundCeremony"`
		WinningTeam          string `json:"winningTeam"`
		SpikePlanter         string `json:"bombPlanter"`
		SpikeDefuser         string `json:"bombDefuser"`
		PlantRoundTime       uint32 `json:"plantRoundTime"`
		PlantPlayerLocations []struct {
			PUUID       string  `json:"subject"`
			ViewRadians float32 `json:"viewRadians"`
			Location    struct {
				X int32 `json:"x"`
				Y int32 `json:"y"`
			} `json:"location"`
		} `json:"plantPlayerLocations"`
		PlantLocation struct {
			X int32 `json:"x"`
			Y int32 `json:"y"`
		} `json:"plantLocation"`
		PlantSite             string `json:"plantSite"`
		DefuseRoundTime       uint32 `json:"defuseRoundTime"`
		DefusePlayerLocations []struct {
			PUUID       string  `json:"subject"`
			ViewRadians float32 `json:"viewRadians"`
			Location    struct {
				X int32 `json:"x"`
				Y int32 `json:"y"`
			} `json:"location"`
		} `json:"defusePlayerLocations"`
		DefuseLocation struct {
			X int32 `json:"x"`
			Y int32 `json:"y"`
		} `json:"defuseLocation"`
		PlayerStats []struct {
			PUUID string `json:"subject"`
			Kills []struct {
				GameTime       uint32 `json:"gameTime"`
				RoundTime      uint32 `json:"roundTime"`
				Killer         string `json:"killer"`
				Victim         string `json:"victim"`
				VictimLocation struct {
					X int32 `json:"x"`
					Y int32 `json:"y"`
				} `json:"victimLocation"`
				Assistants      []string `json:"assistants"`
				PlayerLocations []struct {
					PUUID       string  `json:"subject"`
					ViewRadians float32 `json:"viewRadians"`
					Location    struct {
						X int32 `json:"x"`
						Y int32 `json:"y"`
					} `json:"location"`
				} `json:"playerLocations"`
				FinishingDamage struct {
					Type                  string `json:"damageType"`
					Item                  string `json:"damageItem"`
					UsedSecondaryFireMode bool   `json:"isSecondaryFireMode"`
				} `json:"finishingDamage"`
			} `json:"kills"`
			Damages []struct {
				Receiver  string `json:"receiver"`
				Damage    uint16 `json:"damage"`
				Legshots  uint8  `json:"legshots"`
				Bodyshots uint8  `json:"bodyshots"`
				Headshots uint8  `json:"headshots"`
			} `json:"damage"`
			Score   uint16 `json:"score"`
			Economy struct {
				LoadoutValue     uint16 `json:"loadoutValue"`
				Weapon           string `json:"weapon"`
				Armor            string `json:"armor"`
				RemainingCredits uint16 `json:"remaining"`
				CreditsSpent     uint16 `json:"spent"`
			} `json:"economy"`
			WasAfk        bool `json:"wasAfk"`
			WasPenalized  bool `json:"wasPenalized"`
			StayedInSpawn bool `json:"stayedInSpawn"`
		} `json:"playerStats"`
		ResultCode      string `json:"roundResultCode"`
		PlayerEconomies []struct {
			PUUID            string `json:"subject"`
			LoadoutValue     uint16 `json:"loadoutValue"`
			Weapon           string `json:"weapon"`
			Armor            string `json:"armor"`
			RemainingCredits uint16 `json:"remaining"`
			CreditsSpent     uint16 `json:"spent"`
		} `json:"playerEconomies"`
		PlayerScores []struct {
			PUUID string `json:"subject"`
			Score uint16 `json:"score"`
		} `json:"playerScores"`
	} `json:"roundResults"`

	Kills []struct {
		GameTime       uint32 `json:"gameTime"`
		RoundTime      uint32 `json:"roundTime"`
		Killer         string `json:"killer"`
		Victim         string `json:"victim"`
		VictimLocation struct {
			X int32 `json:"x"`
			Y int32 `json:"y"`
		} `json:"victimLocation"`
		Assistants      []string `json:"assistants"`
		PlayerLocations []struct {
			PUUID       string  `json:"subject"`
			ViewRadians float32 `json:"viewRadians"`
			Location    struct {
				X int32 `json:"x"`
				Y int32 `json:"y"`
			} `json:"location"`
		} `json:"playerLocations"`
		FinishingDamage struct {
			Type                  string `json:"damageType"`
			Item                  string `json:"damageItem"`
			UsedSecondaryFireMode bool   `json:"isSecondaryFireMode"`
		} `json:"finishingDamage"`
		Round uint8 `json:"round"`
	} `json:"kills"`
}

func GetMatchDetails(shard, uuid string) (*MatchDetailsResponseBody, error) {
	res, err := http.SendRequest("GET", fmt.Sprintf("https://pd.%s.a.pvp.net/match-details/v1/matches/%s", shard, uuid), nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var body *MatchDetailsResponseBody
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
