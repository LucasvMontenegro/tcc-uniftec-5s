package service

import (
	account_aggregate "github.com/tcc-uniftec-5s/internal/domain/accountAggregate"
)

type AccountFactory struct {
	accountRepository account_aggregate.AccountRepository
}

type AccountFactoryInterface interface {
	NewAccount(credential *account_aggregate.CredentialEntity) account_aggregate.AccountInterface
}

func NewAccountFactory(accountRepository account_aggregate.AccountRepository) AccountFactoryInterface {
	return AccountFactory{
		accountRepository: accountRepository,
	}
}

func (f AccountFactory) NewAccount(credential *account_aggregate.CredentialEntity) account_aggregate.AccountInterface {
	account := account_aggregate.AccountEntity{
		Credential: credential,
	}

	return AccountImpl{
		accountEntity:     &account,
		accountRepository: f.accountRepository,
	}
}
