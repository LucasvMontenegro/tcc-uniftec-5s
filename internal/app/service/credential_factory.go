package service

import "github.com/tcc-uniftec-5s/internal/domain/entity"

type credentialFactory struct {
	credentialRepository entity.CredentialRepository
}

func NewCredentialFactory(credentialRepository entity.CredentialRepository) entity.CredentialFactoryInterface {
	return credentialFactory{
		credentialRepository: credentialRepository,
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
		CredentialEntity:     &credential,
		CredentialRepository: f.credentialRepository,
	}
}
