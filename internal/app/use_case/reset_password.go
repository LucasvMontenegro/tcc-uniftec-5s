package usecase

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/tcc-uniftec-5s/internal/app/repository"
	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

type ResetPassword interface {
	Execute(ctx context.Context, email, password string) error
}

func NewResetPassword(
	txHandler repository.TxHandlerInterface,
	credentialFactory entity.CredentialFactoryInterface) ResetPassword {

	return resetPassword{
		txHandler:         txHandler,
		credentialFactory: credentialFactory,
	}
}

type resetPassword struct {
	txHandler         repository.TxHandlerInterface
	credentialFactory entity.CredentialFactoryInterface
}

func (rp resetPassword) Execute(ctx context.Context, email, password string) error {
	log.Info().Msg("starting reset password use case")

	ctx, err := rp.txHandler.NewContextWithTransaction(ctx)
	if err != nil {
		log.Info().Msg("tx handler failed to start transaction")
		return err
	}

	credential := rp.credentialFactory.NewCredential(email, password)

	if err := credential.UpdatePassword(ctx, password); err != nil {
		log.Info().Msg("rolling back tx")
		if txerr := rp.txHandler.Rollback(ctx); txerr != nil {
			return nil
		}

		return err
	}

	log.Info().Msg("committing tx")
	if err = rp.txHandler.Commit(ctx); err != nil {
		log.Info().Msg("rolling back tx")
		if txerr := rp.txHandler.Rollback(ctx); txerr != nil {
			return txerr
		}

		return err
	}

	log.Info().Msg("reset password use case done")
	return nil
}
