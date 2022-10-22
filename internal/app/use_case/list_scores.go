package usecase

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

type ListScores interface {
	Execute(ctx context.Context, teamID int64) ([]entity.Score, error)
}

func NewListScores(scoreFactory entity.ScoreFactoryInterface) ListScores {
	return listScores{
		scoreFactory: scoreFactory,
	}
}

type listScores struct {
	scoreFactory entity.ScoreFactoryInterface
}

func (uc listScores) Execute(ctx context.Context, teamID int64) ([]entity.Score, error) {
	log.Info().Msg("starting list scores use case")

	lscores, _ := uc.scoreFactory.ListScores(ctx, teamID)

	var scores []entity.Score
	for _, score := range lscores {
		scores = append(scores, *score.Self())
	}

	log.Info().Msg("list scores use case done")
	return scores, nil
}
