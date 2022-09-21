package service

import (
	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

type prizeFactory struct {
	prizeRepository entity.PrizeRepository
}

func NewPrizeFactory(prizeRepository entity.PrizeRepository) entity.PrizeFactoryInterface {
	return prizeFactory{
		prizeRepository: prizeRepository,
	}
}

func (f prizeFactory) NewPrize(name string, description *string, edition *entity.Edition) entity.PrizeInterface {
	entity := entity.Prize{
		Name:        name,
		Description: description,
		Edition:     edition,
	}

	return prize{
		prizeEntity:     &entity,
		prizeRepository: f.prizeRepository,
	}
}
