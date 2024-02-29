package api

import (
	"encoding/json"
	"fmt"

	"github.com/itsbradn/govalo/pkg/http"
)

type PlayerMMRResponseBody struct {
	Version                     uint16 `json:"Version"`
	PUUID                       string `json:"Subject"`
	NewPlayerExperienceFinished bool   `json:"NewPlayerExperienceFinished"`
	QueueSkills                 map[string]struct {
		TotalGamesNeededForRating         uint8 `json:"TotalGamesNeededForRating"`
		TotalGamesNeededForLeaderboard    uint8 `json:"TotalGamesNeededForLeaderboard"`
		CurrentSeasonGamesNeededForRating uint8 `json:"CurrentSeasonGamesNeededForRating"`
		SeasonalInfoBySeasonID            map[string]struct {
			SeasonID                   string            `json:"SeasonID"`
			NumberOfWins               uint32            `json:"NumberOfWins"`
			NumberOfWinsWithPlacements uint32            `json:"NumberOfWinsWithPlacements"`
			NumberOfGames              uint32            `json:"NumberOfGames"`
			Rank                       uint8             `json:"Rank"`
			CapstoneWins               uint32            `json:"CapstoneWins"`
			LeaderboardRank            uint32            `json:"LeaderboardRank"`
			CompetitiveTier            uint8             `json:"CompetitiveTier"`
			RankedRating               uint8             `json:"RankedRating"`
			WinsByTier                 map[string]uint32 `json:"WinsByTier"`
			GamesNeededForRating       uint8             `json:"GamesNeededForRating"`
			TotalWinsNeededForRank     uint8             `json:"TotalWinsNeededForRank"`
		} `json:"SeasonalInfoBySeasonID"`
	} `json:"QueueSkills"`
	LatestCompetitiveUpdate struct {
		MatchID                      string `json:"MatchID"`
		MapID                        string `json:"MapID"`
		SeasonID                     string `json:"SeasonID"`
		MatchStartTime               uint64 `json:"MatchStartTime"`
		TierAfterUpdate              uint8  `json:"TierAfterUpdate"`
		TierBeforeUpdate             uint8  `json:"TierBeforeUpdate"`
		RankedRatingAfterUpdate      uint8  `json:"RankedRatingAfterUpdate"`
		RankedRatingBeforeUpdate     uint8  `json:"RankedRatingBeforeUpdate"`
		RankedRatingEarned           uint8  `json:"RankedRatingEarned"`
		RankedRatingPerformanceBonus uint8  `json:"RankedRatingPerformanceBonus"`
		CompetitiveMovement          string `json:"CompetitiveMovement"`
		AFKPenalty                   uint8  `json:"AFKPenalty"`
	} `json:"LatestCompetitiveUpdate"`
	IsLeaderboardAnonymized bool `json:"IsLeaderboardAnonymized"`
	IsActRankBadgeHidden    bool `json:"IsActRankBadgeHidden"`
}

func GetPlayerMMR(shard, puuid string) (*PlayerMMRResponseBody, error) {
	res, err := http.SendRequest("GET", fmt.Sprintf("https://pd.%s.a.pvp.net/mmr/v1/players/%s", shard, puuid), nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var body *PlayerMMRResponseBody
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
