package usecase

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/tcc-uniftec-5s/internal/app/repository"
	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

type CreateEditionDTO struct {
	PrizeDTO
	EditionDTO
}

type PrizeDTO struct {
	Name        string
	Description *string
}

type EditionDTO struct {
	Name        string
	Description *string
	Status      *string
	StartDate   time.Time
	EndDate     time.Time
}

type CreateEdition interface {
	Execute(ctx context.Context, dto CreateEditionDTO) error
}

func NewCreateEdition(
	txHandler repository.TxHandlerInterface,
	editionFactory entity.EditionFactoryInterface,
	prizeFactory entity.PrizeFactoryInterface) CreateEdition {

	return createEdition{
		txHandler:      txHandler,
		editionFactory: editionFactory,
		prizeFactory:   prizeFactory,
	}
}

type createEdition struct {
	txHandler      repository.TxHandlerInterface
	editionFactory entity.EditionFactoryInterface
	prizeFactory   entity.PrizeFactoryInterface
}

func (uc createEdition) Execute(ctx context.Context, dto CreateEditionDTO) error {
	log.Info().Msg("starting create edition use case")

	ctx, err := uc.txHandler.NewContextWithTransaction(ctx)
	if err != nil {
		log.Info().Msg("tx handler failed to start transaction")
		return err
	}

	edition := uc.editionFactory.NewEdition(dto.EditionDTO.Name, dto.EditionDTO.Description, dto.EditionDTO.Status, dto.EditionDTO.StartDate, dto.EditionDTO.EndDate)
	if err := edition.Create(ctx); err != nil {
		log.Info().Msg("rolling back tx")
		if txerr := uc.txHandler.Rollback(ctx); txerr != nil {
			return txerr
		}

		return err
	}

	prize := uc.prizeFactory.NewPrize(dto.PrizeDTO.Name, dto.PrizeDTO.Description, edition.Self())
	if err := prize.Create(ctx, edition.Self()); err != nil {
		log.Info().Msg("rolling back tx")
		if txerr := uc.txHandler.Rollback(ctx); txerr != nil {
			return txerr
		}

		return err
	}

	log.Info().Msg("committing tx")
	if err = uc.txHandler.Commit(ctx); err != nil {
		log.Info().Msg("rolling back tx")
		if txerr := uc.txHandler.Rollback(ctx); txerr != nil {
			return txerr
		}

		return err
	}

	log.Info().Msg("create edition use case done")
	return nil
}
