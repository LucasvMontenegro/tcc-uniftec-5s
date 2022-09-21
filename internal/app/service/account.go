package service

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

type AccountImpl struct {
	accountEntity     *entity.Account
	accountRepository entity.AccountRepository
}

func (s AccountImpl) Self(ctx context.Context) *entity.Account {
	return s.accountEntity
}

func (s AccountImpl) Create(ctx context.Context) (*entity.Account, error) {
	log.Info().Msg("creating account")

	if err := s.accountRepository.Save(ctx, s.Self(ctx)); err != nil {
		log.Info().Msg("error saving account")
		return nil, err
	}

	return s.Self(ctx), nil
}

func (s AccountImpl) AddUser(ctx context.Context, user *entity.User) error {
	log.Info().Msg("adding user to account")

	s.accountEntity.User = user

	if err := s.accountRepository.SetUser(ctx, s.Self(ctx)); err != nil {
		log.Info().Msg("error updating account")
		return err
	}

	return nil
}
