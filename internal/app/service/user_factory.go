package service

import "github.com/tcc-uniftec-5s/internal/domain/entity"

type UserFactory struct {
	userRepository entity.UserRepository
}

func NewUserFactory(userRepository entity.UserRepository) entity.UserFactoryInterface {
	return UserFactory{
		userRepository: userRepository,
	}
}

func (f UserFactory) NewUser(account *entity.AccountEntity, name string) entity.UserInterface {
	user := entity.UserEntity{
		Account: account,
		Name:    name,
		Status:  "ACTIVE",
	}

	return UserImpl{
		userEntity:     &user,
		userRepository: f.userRepository,
	}
}
