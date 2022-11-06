package usecase

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/tcc-uniftec-5s/internal/app/repository"
	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

type Ranking struct{}

type RetrieveRanking interface {
	Execute(ctx context.Context) (*entity.Ranking, error)
}

func NewRetrieveRanking(
	rankingFactory entity.RankingFactoryInterface,
	editionFactory entity.EditionFactoryInterface) RetrieveRanking {

	return retrieveRanking{
		rankingFactory: rankingFactory,
		editionFactory: editionFactory,
	}
}

type retrieveRanking struct {
	txHandler      repository.TxHandlerInterface
	rankingFactory entity.RankingFactoryInterface
	editionFactory entity.EditionFactoryInterface
}

func (uc retrieveRanking) Execute(ctx context.Context) (*entity.Ranking, error) {
	log.Info().Msg("starting retrieve ranking use case")

	edition, err := uc.editionFactory.GetCurrent(ctx)
	if err != nil {
		return nil, err
	}

	ranking, err := uc.rankingFactory.Generate(ctx, edition.Self())

	log.Info().Msg("retrieve ranking use case done")
	return ranking.Self(), nil
}
