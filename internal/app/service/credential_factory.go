package service

import account_aggregate "github.com/tcc-uniftec-5s/internal/domain/accountAggregate"

type CredentialFactory struct {
	credentialRepository account_aggregate.CredentialRepository
}

type CredentialFactoryInterface interface {
	NewCredential(email, password string) account_aggregate.CredentialInterface
}

func NewCredentialFactory(credentialRepository account_aggregate.CredentialRepository) CredentialFactoryInterface {
	return CredentialFactory{
		credentialRepository: credentialRepository,
	}
}

func (f CredentialFactory) NewCredential(email, password string) account_aggregate.CredentialInterface {
	credential := account_aggregate.CredentialEntity{
		Email:    email,
		Password: password,
	}

	return CredentialImpl{
		CredentialEntity:     &credential,
		CredentialRepository: f.credentialRepository,
	}
}
