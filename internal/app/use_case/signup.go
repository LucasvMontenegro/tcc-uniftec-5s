package usecase

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/tcc-uniftec-5s/internal/app/repository"
	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

type Signup interface {
	Signup(ctx context.Context, email, password, name string) error
}

func NewSignup(
	txHandler repository.TxHandlerInterface,
	credentialFactory entity.CredentialFactoryInterface,
	accountFactory entity.AccountFactoryInterface,
	userFactory entity.UserFactoryInterface) Signup {

	return signup{
		txHandler:         txHandler,
		credentialFactory: credentialFactory,
		accountFactory:    accountFactory,
		userFactory:       userFactory,
	}
}

type signup struct {
	txHandler         repository.TxHandlerInterface
	credentialFactory entity.CredentialFactoryInterface
	accountFactory    entity.AccountFactoryInterface
	userFactory       entity.UserFactoryInterface
}

func (uc signup) Signup(ctx context.Context, email, password, name string) error {
	log.Info().Msg("starting signup use case")

	ctx, err := uc.txHandler.NewContextWithTransaction(ctx)
	if err != nil {
		log.Info().Msg("tx handler failed to start transaction")
		return err
	}

	credential := uc.credentialFactory.NewCredential(email, password)
	if err := credential.Signup(ctx); err != nil {
		log.Info().Msg("rolling back tx")
		if txerr := uc.txHandler.Rollback(ctx); txerr != nil {
			return txerr
		}

		return err
	}

	account := uc.accountFactory.NewAccount(credential.Self(ctx))
	if _, err := account.Create(ctx); err != nil {
		log.Info().Msg("rolling back tx")
		if txerr := uc.txHandler.Rollback(ctx); txerr != nil {
			return txerr
		}

		return err
	}

	if err := credential.AddAccount(ctx, account.Self(ctx)); err != nil {
		log.Info().Msg("rolling back tx")
		if txerr := uc.txHandler.Rollback(ctx); txerr != nil {
			return txerr
		}

		return err
	}

	user := uc.userFactory.NewUser(account.Self(ctx), name)
	_, err = user.Create(ctx)
	if err != nil {
		log.Info().Msg("rolling back tx")
		if txerr := uc.txHandler.Rollback(ctx); txerr != nil {
			return txerr
		}

		return err
	}

	if err := account.AddUser(ctx, user.Self()); err != nil {
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

	log.Info().Msg("signup use case done")
	return nil
}
