package service

import (
	"context"

	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

type rankingFactory struct {
	teamFactory  entity.TeamFactoryInterface
	scoreFactory entity.ScoreFactoryInterface
}

func NewRankingFactory(
	teamFactory entity.TeamFactoryInterface,
	scoreFactory entity.ScoreFactoryInterface,
) entity.RankingFactoryInterface {

	return rankingFactory{
		teamFactory,
		scoreFactory,
	}
}

func (f rankingFactory) Generate(ctx context.Context, currentEdition *entity.Edition) (entity.RankingInterface, error) {
	e := entity.Ranking{
		Edition: currentEdition,
	}

	teams := f.teamFactory.ListTeamsByEdition(ctx, currentEdition)
	// e.TeamScores = make([]*entity.TeamScore, len(teams))

	for _, team := range teams {
		scores, err := f.scoreFactory.ListScores(ctx, *team.Self().ID)
		if err != nil {
			return nil, err
		}

		ts := &entity.TeamScore{
			TeamName: team.Self().Name,
		}

		for _, score := range scores {
			ts.Scores = append(ts.Scores, score.Self())
		}

		e.TeamScores = append(e.TeamScores, ts)
	}

	return rankingImpl{
		rankingEntity: &e,
	}, nil
}
