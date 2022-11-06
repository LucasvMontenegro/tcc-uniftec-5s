package response

import (
	"strconv"

	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

func BuildRanking(ranking entity.Ranking) Ranking {
	description := ""
	if ranking.Edition.Description != nil {
		description = *ranking.Edition.Description
	}

	res := Ranking{
		RankingEdition: RankingEdition{
			Name:        ranking.Edition.Name,
			Description: description,
			Status:      *ranking.Edition.Status,
			StartDate:   ranking.Edition.StartDate.String(),
			EndDate:     ranking.Edition.EndDate.String(),
		},
	}

	for _, teamScore := range ranking.TeamScores {
		seiso := "none"
		seiri := "none"
		seiton := "none"
		seiketsu := "none"
		shitsuke := "none"
		total := "none"

		t := sumScores(teamScore.Scores)
		if t != 0 {
			total = strconv.Itoa(t)
		}

		for _, score := range teamScore.Scores {
			sensoName := *score.FiveS.Name
			switch sensoName {
			case "SEISO":
				seiso = strconv.Itoa(*score.Score)
			case "SEIRI":
				seiri = strconv.Itoa(*score.Score)
			case "SEITON":
				seiton = strconv.Itoa(*score.Score)
			case "SEIKETSU":
				seiketsu = strconv.Itoa(*score.Score)
			case "SHITSUKE":
				shitsuke = strconv.Itoa(*score.Score)
			}
		}

		scores := RankingScores{
			TeamName: teamScore.TeamName,
			Seiso:    seiso,
			Seiri:    seiri,
			Seiton:   seiton,
			Seiketsu: seiketsu,
			Shitsuke: shitsuke,
			Total:    total,
		}

		res.Scores = append(res.Scores, scores)

	}

	return res
}

type Ranking struct {
	RankingEdition
	// RankingPrize
	Scores []RankingScores
}

type RankingEdition struct {
	Name        string `json:"edition_name"`
	Description string `json:"edition_description"`
	Status      string `json:"status"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

type RankingScores struct {
	TeamName string `json:"team_name"`
	Seiri    string `json:"seiri"`
	Seiton   string `json:"seiton"`
	Seiso    string `json:"seiso"`
	Seiketsu string `json:"seiketsu"`
	Shitsuke string `json:"shitsuke"`
	Total    string `json:"total"`
}

// type RankingPrize struct {
// 	Name        string `json:"prize_name"`
// 	Description string `json:"prize_description"`
// }

func sumScores(scores []*entity.Score) int {
	total := 0
	for _, score := range scores {
		total += *score.Score
	}

	return total
}
