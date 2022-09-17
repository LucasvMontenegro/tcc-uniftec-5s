package service

import (
	"context"

	"github.com/rs/zerolog/log"
	account_aggregate "github.com/tcc-uniftec-5s/internal/domain/accountAggregate"
)

type UserImpl struct {
	userEntity     *account_aggregate.UserEntity
	userRepository account_aggregate.UserRepository
}

func (s UserImpl) Self(ctx context.Context) *account_aggregate.UserEntity {
	return s.userEntity
}

func (s UserImpl) Create(ctx context.Context) (*account_aggregate.UserEntity, error) {
	log.Info().Msg("creating user")

	err := s.userRepository.Save(ctx, s.userEntity)
	if err != nil {
		log.Info().Msg("error creating user")
		return nil, err
	}

	return s.userEntity, nil
}
