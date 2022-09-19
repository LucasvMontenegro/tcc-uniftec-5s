package service

import "github.com/tcc-uniftec-5s/internal/domain/entity"

type accountFactory struct {
	accountRepository entity.AccountRepository
}

func NewAccountFactory(accountRepository entity.AccountRepository) entity.AccountFactoryInterface {
	return accountFactory{
		accountRepository: accountRepository,
	}
}

func (f accountFactory) NewAccount(credential *entity.CredentialEntity) entity.AccountInterface {
	account := entity.AccountEntity{
		Credential: credential,
	}

	return AccountImpl{
		accountEntity:     &account,
		accountRepository: f.accountRepository,
	}
}
