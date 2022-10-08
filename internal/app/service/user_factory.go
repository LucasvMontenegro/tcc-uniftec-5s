package service

import (
	"context"

	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

type UserFactory struct {
	userRepository entity.UserRepository
}

func NewUserFactory(userRepository entity.UserRepository) entity.UserFactoryInterface {
	return UserFactory{
		userRepository: userRepository,
	}
}

func (f UserFactory) NewUser(account *entity.Account, name string) entity.UserInterface {
	user := entity.User{
		Account: account,
		Name:    name,
		Status:  "ACTIVE",
	}

	return UserImpl{
		userEntity:     &user,
		userRepository: f.userRepository,
	}
}

func (f UserFactory) GetTeamlessUsers(ctx context.Context) ([]entity.UserInterface, error) {
	users, err := f.userRepository.GetTeamlessUsers(ctx)
	if err != nil {
		return nil, err
	}

	uis := make([]entity.UserInterface, 0)
	for _, user := range users {
		impl := UserImpl{
			userEntity:     user,
			userRepository: f.userRepository,
		}

		uis = append(uis, impl)
	}

	return uis, nil
}

func (f UserFactory) ListUsers(ctx context.Context) ([]entity.UserInterface, error) {
	users, err := f.userRepository.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	uis := make([]entity.UserInterface, 0)
	for _, user := range users {
		impl := UserImpl{
			userEntity:     user,
			userRepository: f.userRepository,
		}

		uis = append(uis, impl)
	}

	return uis, nil
}
