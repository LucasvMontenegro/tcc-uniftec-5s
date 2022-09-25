package usecase

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/tcc-uniftec-5s/internal/app/repository"
	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

type CreateTeam interface {
	Execute(ctx context.Context, name string) (*entity.Team, error)
}

func NewCreateTeam(
	txHandler repository.TxHandlerInterface,
	teamFactory entity.TeamFactoryInterface,
	editionFactory entity.EditionFactoryInterface) CreateTeam {

	return createTeam{
		txHandler:      txHandler,
		teamFactory:    teamFactory,
		editionFactory: editionFactory,
	}
}

type createTeam struct {
	txHandler      repository.TxHandlerInterface
	teamFactory    entity.TeamFactoryInterface
	editionFactory entity.EditionFactoryInterface
}

func (uc createTeam) Execute(ctx context.Context, name string) (*entity.Team, error) {
	log.Info().Msg("starting create team use case")

	ctx, err := uc.txHandler.NewContextWithTransaction(ctx)
	if err != nil {
		log.Info().Msg("tx handler failed to start transaction")
		return nil, err
	}

	edition, err := uc.editionFactory.GetCurrent(ctx)
	if err != nil {
		return nil, err
	}

	team := uc.teamFactory.NewTeam(ctx, name, edition.Self())
	if err := team.Create(ctx, edition.Self()); err != nil {
		log.Info().Msg("rolling back tx")
		if txerr := uc.txHandler.Rollback(ctx); txerr != nil {
			return nil, txerr
		}

		return nil, err
	}

	log.Info().Msg("committing tx")
	if err = uc.txHandler.Commit(ctx); err != nil {
		log.Info().Msg("rolling back tx")
		if txerr := uc.txHandler.Rollback(ctx); txerr != nil {
			return nil, txerr
		}

		return nil, err
	}

	log.Info().Msg("create team use case done")
	return team.Self(), nil
}
