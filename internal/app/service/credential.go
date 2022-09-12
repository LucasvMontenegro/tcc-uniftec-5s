package service

import (
	"context"

	"github.com/rs/zerolog/log"
	account_aggregate "github.com/tcc-uniftec-5s/internal/domain/accountAggregate"
)

type CredentialImpl struct {
	CredentialEntity     *account_aggregate.CredentialEntity
	CredentialRepository account_aggregate.CredentialRepository
}

func (s CredentialImpl) Self(ctx context.Context) *account_aggregate.CredentialEntity {
	return s.CredentialEntity
}

func (s CredentialImpl) Signup(ctx context.Context) error {
	log.Info().Msg("creating credential")

	if err := s.CredentialRepository.Save(ctx, s.CredentialEntity); err != nil {
		log.Info().Msg("error saving credential")
		return err
	}

	return nil
}

func (s CredentialImpl) AddAccount(ctx context.Context, account *account_aggregate.AccountEntity) error {
	log.Info().Msg("adding account to credential")

	s.CredentialEntity.Account = account

	if err := s.CredentialRepository.Update(ctx, s.CredentialEntity); err != nil {
		log.Info().Msg("error saving credential")
		return err
	}

	return nil
}
