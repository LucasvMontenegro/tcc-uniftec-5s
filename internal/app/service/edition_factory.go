package service

import (
	"time"

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
