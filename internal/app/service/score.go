package service

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

type scoreImpl struct {
	scoreEntity     *entity.Score
	scoreRepository entity.ScoreRepository
}

func (s scoreImpl) Self() *entity.Score {
	return s.scoreEntity
}

func (s scoreImpl) Create(ctx context.Context) error {
	log.Info().Msg("creating score")

	if err := s.scoreRepository.Save(ctx, s.scoreEntity); err != nil {
		log.Info().Msg("error saving score")
		return err
	}

	return nil
}
