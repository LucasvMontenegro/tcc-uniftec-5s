package service

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

type UserImpl struct {
	userEntity     *entity.User
	userRepository entity.UserRepository
}

func (s UserImpl) Self() *entity.User {
	return s.userEntity
}

func (s UserImpl) Create(ctx context.Context) (*entity.User, error) {
	log.Info().Msg("creating user")

	err := s.userRepository.Save(ctx, s.userEntity)
	if err != nil {
		log.Info().Msg("error creating user")
		return nil, err
	}

	return s.userEntity, nil
}
