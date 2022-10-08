package usecase

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/tcc-uniftec-5s/internal/app/repository"
	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

type ListUsers interface {
	Execute(ctx context.Context) ([]entity.UserInterface, error)
}

func NewListUsers(
	txHandler repository.TxHandlerInterface,
	userFactory entity.UserFactoryInterface) ListUsers {

	return listUsers{
		txHandler:   txHandler,
		userFactory: userFactory,
	}
}

type listUsers struct {
	txHandler   repository.TxHandlerInterface
	userFactory entity.UserFactoryInterface
}

func (uc listUsers) Execute(ctx context.Context) ([]entity.UserInterface, error) {
	log.Info().Msg("starting list users use case")

	users, err := uc.userFactory.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	log.Info().Msg("list users use case done")
	return users, nil
}
