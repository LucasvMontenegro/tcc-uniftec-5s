package service

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

type fiveSFactory struct {
	fiveSRepository entity.FiveSRepository
}

func NewFiveSFactory(fiveSRepository entity.FiveSRepository) entity.FiveSFactoryInterface {
	return fiveSFactory{
		fiveSRepository: fiveSRepository,
	}
}

func (f fiveSFactory) GetByName(ctx context.Context, name string) (entity.FiveSInterface, error) {
	log.Info().Msg("getting fiveS by name")

	e := &entity.FiveS{
		Name: &name,
	}

	if err := f.fiveSRepository.GetByName(ctx, e); err != nil {
		log.Info().Msg("error getting fiveS by name")
		return nil, err
	}

	return fiveSImpl{
		fiveSEntity:     e,
		fiveSRepository: f.fiveSRepository,
	}, nil
}
