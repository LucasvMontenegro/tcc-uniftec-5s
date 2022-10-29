package service

import (
	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

type fiveSImpl struct {
	fiveSEntity     *entity.FiveS
	fiveSRepository entity.FiveSRepository
}

func (s fiveSImpl) Self() *entity.FiveS {
	return s.fiveSEntity
}
