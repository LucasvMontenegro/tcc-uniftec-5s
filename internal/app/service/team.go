package service

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"
	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

var ErrInvalidTeamDate = errors.New("start date must be before end date")

type teamImpl struct {
	teamEntity     *entity.Team
	teamRepository entity.TeamRepository
}

func (s teamImpl) Self() *entity.Team {
	return s.teamEntity
}

func (s teamImpl) Create(ctx context.Context, edition *entity.Edition) error {
	log.Info().Msg("creating team")

	if err := s.teamRepository.Save(ctx, s.Self(), edition); err != nil {
		log.Info().Msg("error saving team")
		return err
	}

	return nil
}
