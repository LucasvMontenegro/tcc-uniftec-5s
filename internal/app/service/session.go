package service

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

type SessionImpl struct {
	sessionEntity     *entity.Session
	sessionRepository entity.SessionRepository
}

func (s SessionImpl) Self(ctx context.Context) *entity.Session {
	return s.sessionEntity
}

func (s SessionImpl) Save(ctx context.Context) error {
	log.Info().Msg("saving session")

	if err := s.sessionRepository.Save(ctx, s.Self(ctx)); err != nil {
		log.Info().Msg("error saving session")
		return err
	}

	log.Info().Msg("saving session history")
	if err := s.sessionRepository.SaveHistory(ctx, s.Self(ctx)); err != nil {
		log.Info().Msg("error saving session history")
		return err
	}

	return nil
}
