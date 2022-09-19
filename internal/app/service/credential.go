package service

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

type CredentialImpl struct {
	CredentialEntity     *entity.CredentialEntity
	CredentialRepository entity.CredentialRepository
}

func (s CredentialImpl) Self(ctx context.Context) *entity.CredentialEntity {
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

func (s CredentialImpl) AddAccount(ctx context.Context, account *entity.AccountEntity) error {
	log.Info().Msg("adding account to credential")

	s.CredentialEntity.Account = account

	if err := s.CredentialRepository.Update(ctx, s.CredentialEntity); err != nil {
		log.Info().Msg("error saving credential")
		return err
	}

	return nil
}

func (s CredentialImpl) Identify(ctx context.Context) error {
	log.Info().Msg("identifying credential")
	if err := s.CredentialRepository.Identify(ctx, s.CredentialEntity); err != nil {
		// IF NOT FOUND log.Info().Msg("credential not found")
		log.Info().Msg("error identifying credential")
		return nil
	}

	s.generateJWT(ctx)
	return nil
}

func (s CredentialImpl) generateJWT(ctx context.Context) {
	log.Info().Msg("generating JWT")
	s.CredentialEntity.JWT = "JWT"
}
