package service

import (
	account_aggregate "github.com/tcc-uniftec-5s/internal/domain/accountAggregate"
)

type UserFactory struct {
	userRepository account_aggregate.UserRepository
}

type UserFactoryInterface interface {
	NewUser(account *account_aggregate.AccountEntity, name string) account_aggregate.UserInterface
}

func NewUserFactory(userRepository account_aggregate.UserRepository) UserFactoryInterface {
	return UserFactory{
		userRepository: userRepository,
	}
}

func (f UserFactory) NewUser(account *account_aggregate.AccountEntity, name string) account_aggregate.UserInterface {
	user := account_aggregate.UserEntity{
		Account: account,
		Name:    name,
		Status:  "ACTIVE",
	}

	return UserImpl{
		userEntity:     &user,
		userRepository: f.userRepository,
	}
}
