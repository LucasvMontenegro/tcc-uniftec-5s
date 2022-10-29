package response

import (
	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

func NewListedScores(scores []entity.Score) []*Score {
	var response []*Score
	if len(scores) == 0 {
		return []*Score{}
	}

	for _, score := range scores {
		response = append(response, &Score{
			ID:        score.ID,
			FiveSName: score.FiveS.Name,
			TeamName:  &score.Team.Name,
			Score:     score.Score,
		})
	}

	return response
}

type Score struct {
	ID        *int64 `json:"id"`
	FiveSName *string
	TeamName  *string
	Score     *int
}
