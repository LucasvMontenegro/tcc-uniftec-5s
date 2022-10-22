package usecase

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

type ListEditions interface {
	Execute(ctx context.Context, status string) ([]entity.Edition, error)
}

func NewListEditions(editionFactory entity.EditionFactoryInterface) ListEditions {

	return listEditions{
		editionFactory: editionFactory,
	}
}

type listEditions struct {
	editionFactory entity.EditionFactoryInterface
}

func (uc listEditions) Execute(ctx context.Context, status string) ([]entity.Edition, error) {
	log.Info().Msg("starting list editions use case")

	leditions, _ := uc.editionFactory.ListEditionsByStatus(ctx, status)

	var editions []entity.Edition
	for _, edition := range leditions {
		editions = append(editions, *edition.Self())
	}

	log.Info().Msg("list editions use case done")
	return editions, nil
}
