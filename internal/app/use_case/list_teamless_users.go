package usecase

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/tcc-uniftec-5s/internal/app/repository"
	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

type ListTeamlessUsers interface {
	Execute(ctx context.Context) ([]entity.UserInterface, error)
}

func NewListTeamlessUsers(
	txHandler repository.TxHandlerInterface,
	userFactory entity.UserFactoryInterface) ListTeamlessUsers {

	return listTeamlessUsers{
		txHandler:   txHandler,
		userFactory: userFactory,
	}
}

type listTeamlessUsers struct {
	txHandler   repository.TxHandlerInterface
	userFactory entity.UserFactoryInterface
}

func (uc listTeamlessUsers) Execute(ctx context.Context) ([]entity.UserInterface, error) {
	log.Info().Msg("starting list teamless users use case")

	users, err := uc.userFactory.GetTeamlessUsers(ctx)
	if err != nil {
		return nil, err
	}

	log.Info().Msg("list teamless users use case done")
	return users, nil
}
