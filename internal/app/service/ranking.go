package service

import (
	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

type rankingImpl struct {
	rankingEntity *entity.Ranking
}

func (s rankingImpl) Self() *entity.Ranking {
	return s.rankingEntity
}
