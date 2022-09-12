package service

import (
	"context"

	"github.com/rs/zerolog/log"
	account_aggregate "github.com/tcc-uniftec-5s/internal/domain/accountAggregate"
)

type AccountImpl struct {
	accountEntity     *account_aggregate.AccountEntity
	accountRepository account_aggregate.AccountRepository
}

func (s AccountImpl) Self(ctx context.Context) *account_aggregate.AccountEntity {
	return s.accountEntity
}

func (s AccountImpl) Create(ctx context.Context) (*account_aggregate.AccountEntity, error) {
	log.Info().Msg("creating account")

	if err := s.accountRepository.Save(ctx, s.Self(ctx)); err != nil {
		log.Info().Msg("error saving account")
		return nil, err
	}

	return s.Self(ctx), nil
}

func (s AccountImpl) AddUser(ctx context.Context, user *account_aggregate.UserEntity) error {
	log.Info().Msg("adding user to account")

	s.accountEntity.User = user

	if err := s.accountRepository.Update(ctx, s.Self(ctx)); err != nil {
		log.Info().Msg("error updating account")
		return err
	}

	return nil
}
