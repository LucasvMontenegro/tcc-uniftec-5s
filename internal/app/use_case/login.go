package usecase

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/tcc-uniftec-5s/internal/app/repository"
	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

type Login interface {
	Execute(ctx context.Context, email, password string) (*string, error)
}

func NewLogin(
	txHandler repository.TxHandlerInterface,
	credentialFactory entity.CredentialFactoryInterface,
	sessionFactory entity.SessionFactoryInterface) Login {

	return login{
		txHandler:         txHandler,
		credentialFactory: credentialFactory,
		sessionFactory:    sessionFactory,
	}
}

type login struct {
	txHandler         repository.TxHandlerInterface
	credentialFactory entity.CredentialFactoryInterface
	sessionFactory    entity.SessionFactoryInterface
}

func (uc login) Execute(ctx context.Context, email, password string) (*string, error) {
	log.Info().Msg("starting login use case")

	ctx, err := uc.txHandler.NewContextWithTransaction(ctx)
	if err != nil {
		log.Info().Msg("tx handler failed to start transaction")
		return nil, err
	}

	credential := uc.credentialFactory.NewCredential(email, password)
	if err := credential.Identify(ctx); err != nil {
		log.Info().Msg("rolling back tx")
		if txerr := uc.txHandler.Rollback(ctx); txerr != nil {
			return nil, txerr
		}

		return nil, err
	}

	session := uc.sessionFactory.NewSession(credential.Self(ctx))
	if err := session.Save(ctx); err != nil {
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

	log.Info().Msg("login use case done")
	return session.Self(ctx).JWT, nil
}
