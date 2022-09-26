package service

import (
	"github.com/tcc-uniftec-5s/internal/domain/entity"
	"github.com/tcc-uniftec-5s/internal/token"
)

type credentialFactory struct {
	credentialRepository entity.CredentialRepository
	jwtMaker             token.Maker
}

func NewCredentialFactory(credentialRepository entity.CredentialRepository, maker token.Maker) entity.CredentialFactoryInterface {
	return credentialFactory{
		credentialRepository: credentialRepository,
		jwtMaker:             maker,
	}
}

func (f credentialFactory) NewCredential(email, password string) entity.CredentialInterface {
	account := entity.Account{}
	credential := entity.Credential{
		Email:    email,
		Password: password,
		Account:  &account,
	}

	return CredentialImpl{
		credentialEntity:     &credential,
		credentialRepository: f.credentialRepository,
		jwtMaker:             f.jwtMaker,
	}
}
