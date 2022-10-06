package service

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

type edition struct {
	editionEntity     *entity.Edition
	editionRepository entity.EditionRepository
}

func (s edition) Self() *entity.Edition {
	return s.editionEntity
}

func (s edition) Create(ctx context.Context) error {
	log.Info().Msg("creating edition")

	if err := s.validateDates(ctx); err != nil {
		return err
	}

	if err := s.validateStatus(ctx); err != nil {
		return err
	}

	if err := s.editionRepository.Save(ctx, s.Self()); err != nil {
		log.Info().Msg("error saving edition")
		return err
	}

	return nil
}

func (s edition) validateDates(ctx context.Context) error {
	log.Info().Msg("validating dates")

	if s.editionEntity.EndDate.Before(s.editionEntity.StartDate) ||
		s.editionEntity.EndDate.Equal(s.editionEntity.StartDate) {
		log.Warn().Msg("start date must be before end date")
		return entity.ErrInvalidEditionDate
	}

	return nil
}

func (s edition) validateStatus(ctx context.Context) error {
	log.Info().Msg("validating status")

	if s.editionEntity.Status == nil {
		status := "WAITING"
		s.editionEntity.Status = &status
	}

	return nil
}
