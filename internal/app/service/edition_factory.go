package service

import (
	"context"
	"strings"
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

	return editionImpl{
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

	return editionImpl{
		editionEntity:     e,
		editionRepository: f.editionRepository,
	}, nil
}

func (f editionFactory) ListEditionsByStatus(ctx context.Context, status string) ([]entity.EditionInterface, error) {
	var editions []entity.EditionInterface

	if strings.ToLower(status) == "active" {
		e := &entity.Edition{}
		if err := f.editionRepository.GetCurrent(ctx, e); err != nil {
			log.Info().Msg("error getting current edition")
			return nil, entity.ErrNoCurrentEditionFound
		}

		editions = append(editions, editionImpl{
			editionEntity:     e,
			editionRepository: f.editionRepository,
		})

	} else {
		leditions, _ := f.editionRepository.ListEditions(ctx)
		for _, edition := range leditions {
			editions = append(editions, editionImpl{
				editionEntity:     edition,
				editionRepository: f.editionRepository,
			})
		}
	}

	return editions, nil
}
