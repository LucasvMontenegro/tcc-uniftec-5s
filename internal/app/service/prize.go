package service

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

type prize struct {
	prizeEntity     *entity.Prize
	prizeRepository entity.PrizeRepository
}

func (s prize) Self() *entity.Prize {
	return s.prizeEntity
}

func (s prize) Create(ctx context.Context, edition *entity.Edition) error {
	log.Info().Msg("creating prize")

	s.setEdition(ctx, edition)

	if err := s.prizeRepository.Save(ctx, s.Self()); err != nil {
		log.Info().Msg("error saving prize")
		return err
	}

	return nil
}

func (s prize) setEdition(ctx context.Context, edition *entity.Edition) {
	s.prizeEntity.Edition = edition
}
