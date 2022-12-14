package service

import (
	"context"

	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

type scoreFactory struct {
	scoreRepository entity.ScoreRepository
}

func NewScoreFactory(scoreRepository entity.ScoreRepository) entity.ScoreFactoryInterface {
	return scoreFactory{
		scoreRepository: scoreRepository,
	}
}

func (f scoreFactory) ListScores(ctx context.Context, teamID int64) ([]entity.ScoreInterface, error) {
	var scores []entity.ScoreInterface

	lscores, _ := f.scoreRepository.ListScores(ctx, teamID)
	for _, score := range lscores {
		scores = append(scores, scoreImpl{
			scoreEntity:     score,
			scoreRepository: f.scoreRepository,
		})
	}

	return scores, nil
}

func (f scoreFactory) New(ctx context.Context, team *entity.Team, fives *entity.FiveS, score int) entity.ScoreInterface {
	entity := &entity.Score{
		FiveSID: fives.ID,
		TeamID:  team.ID,
		Score:   &score,
	}

	return scoreImpl{
		scoreEntity:     entity,
		scoreRepository: f.scoreRepository,
	}
}
