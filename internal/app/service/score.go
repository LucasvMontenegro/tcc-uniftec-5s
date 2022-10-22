package service

import (
	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

type scoreImpl struct {
	scoreEntity     *entity.Score
	scoreRepository entity.ScoreRepository
}

func (s scoreImpl) Self() *entity.Score {
	return s.scoreEntity
}
