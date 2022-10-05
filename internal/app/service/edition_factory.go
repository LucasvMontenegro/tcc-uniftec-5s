package service

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

type editionFactory struct {
	editionRepository entity.EditionRepository
}

func NewEditionFactory(editionRepository entity.EditionRepository) entity.EditionFactoryInterface {
	return editionFactory{
		editionRepository: editionRepository,
	}
}

func (f editionFactory) NewEdition(name string, description *string, startDate time.Time, endDate time.Time) entity.EditionInterface {
	entity := entity.Edition{
		Name:        name,
		Description: description,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	return edition{
		editionEntity:     &entity,
		editionRepository: f.editionRepository,
	}
}

func (f editionFactory) GetCurrent(ctx context.Context) (entity.EditionInterface, error) {
	log.Info().Msg("getting current edition")
	e := &entity.Edition{}
	if err := f.editionRepository.GetCurrent(ctx, e); err != nil {
		log.Info().Msg("error getting current edition")
		return nil, entity.ErrNoCurrentEditionFound
	}

	return edition{
		editionEntity:     e,
		editionRepository: f.editionRepository,
	}, nil
}
