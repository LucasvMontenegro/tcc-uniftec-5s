package service

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/tcc-uniftec-5s/internal/domain/entity"
	"github.com/tcc-uniftec-5s/internal/token"
)

type CredentialImpl struct {
	credentialEntity     *entity.Credential
	credentialRepository entity.CredentialRepository
	jwtMaker             token.Maker
}

func (s CredentialImpl) Self(ctx context.Context) *entity.Credential {
	return s.credentialEntity
}

func (s CredentialImpl) Signup(ctx context.Context) error {
	log.Info().Msg("creating credential")

	if err := s.credentialRepository.Save(ctx, s.credentialEntity); err != nil {
		log.Info().Msg("error saving credential")
		return err
	}

	return nil
}

func (s CredentialImpl) AddAccount(ctx context.Context, account *entity.Account) error {
	log.Info().Msg("adding account to credential")

	s.credentialEntity.Account = account

	if err := s.credentialRepository.SetAccount(ctx, s.credentialEntity); err != nil {
		log.Info().Msg("error saving credential")
		return err
	}

	return nil
}

func (s CredentialImpl) Identify(ctx context.Context) error {
	log.Info().Msg("identifying credential")
	if err := s.credentialRepository.Identify(ctx, s.credentialEntity); err != nil {
		// IF NOT FOUND log.Info().Msg("credential not found")
		log.Info().Msg("error identifying credential")
		return err
	}

	if err := s.generateJWT(ctx); err != nil {
		return err
	}

	return nil
}

func (s CredentialImpl) UpdatePassword(ctx context.Context, password string) error {
	log.Info().Msg("updating credential password")

	s.credentialEntity.Password = password

	if err := s.credentialRepository.UpdatePassword(ctx, s.credentialEntity); err != nil {
		// IF NOT FOUND log.Info().Msg("credential not found")
		log.Info().Msg("error updating credential password")
		return nil
	}

	return nil
}

func (s CredentialImpl) generateJWT(ctx context.Context) error {
	log.Info().Msg("generating JWT")

	token, err := s.jwtMaker.CreateToken(
		s.credentialEntity.Account.User.Name,
		s.credentialEntity.Account.User.IsAdmin,
	)
	if err != nil {
		return err
	}

	s.credentialEntity.JWT = &token
	return nil
}
