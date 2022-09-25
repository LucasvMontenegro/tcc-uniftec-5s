package service

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/tcc-uniftec-5s/internal/domain/entity"
	"github.com/tcc-uniftec-5s/internal/token"
)

type CredentialImpl struct {
	CredentialEntity     *entity.Credential
	CredentialRepository entity.CredentialRepository
}

func (s CredentialImpl) Self(ctx context.Context) *entity.Credential {
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

func (s CredentialImpl) AddAccount(ctx context.Context, account *entity.Account) error {
	log.Info().Msg("adding account to credential")

	s.CredentialEntity.Account = account

	if err := s.CredentialRepository.SetAccount(ctx, s.CredentialEntity); err != nil {
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
		return err
	}

	if err := s.generateJWT(ctx); err != nil {
		return err
	}

	return nil
}

type Claims struct {
	username string
	is_admin bool
}

func (Claims) Valid() error {
	return nil
}

func (s CredentialImpl) generateJWT(ctx context.Context) error {
	log.Info().Msg("generating JWT")

	maker, err := token.NewJWTMaker("fsdjkfhsdhfsdhkfjkdshjfkskdjfjksdjkjkfjkshkfjs")
	if err != nil {
		return err
	}
	token, err := maker.CreateToken(
		s.CredentialEntity.Account.User.Name,
		s.CredentialEntity.Account.User.IsAdmin,
	)
	if err != nil {
		return err
	}

	s.CredentialEntity.JWT = &token
	return nil
}

func (s CredentialImpl) UpdatePassword(ctx context.Context, password string) error {
	log.Info().Msg("updating credential password")

	s.CredentialEntity.Password = password

	if err := s.CredentialRepository.UpdatePassword(ctx, s.CredentialEntity); err != nil {
		// IF NOT FOUND log.Info().Msg("credential not found")
		log.Info().Msg("error updating credential password")
		return nil
	}

	return nil
}
