package usecase

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

type ScoreDTO struct {
	FiveSName string
	Score     int
}

type Score interface {
	Execute(ctx context.Context, teamID int64, dto []ScoreDTO) error
}

func NewScore(scoreFactory entity.ScoreFactoryInterface,
	fiveSFactory entity.FiveSFactoryInterface,
	teamFactory entity.TeamFactoryInterface) Score {
	return score{
		scoreFactory: scoreFactory,
		teamFactory:  teamFactory,
		fiveSFactory: fiveSFactory,
	}
}

type score struct {
	scoreFactory entity.ScoreFactoryInterface
	fiveSFactory entity.FiveSFactoryInterface
	teamFactory  entity.TeamFactoryInterface
}

func (uc score) Execute(ctx context.Context, teamID int64, dto []ScoreDTO) error {
	log.Info().Msg("starting score use case")

	team, err := uc.teamFactory.GetByID(ctx, teamID)
	if err != nil {
		return err
	}

	for _, d := range dto {
		fiveS, err := uc.fiveSFactory.GetByName(ctx, d.FiveSName)
		if err != nil {
			return err
		}

		score := uc.scoreFactory.New(ctx, team.Self(), fiveS.Self(), d.Score)

		if err := score.Create(ctx); err != nil {
			return err
		}
	}

	log.Info().Msg("score use case done")
	return nil
}
