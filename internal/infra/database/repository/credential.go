package repository

import (
	"context"

	account_aggregate "github.com/tcc-uniftec-5s/internal/domain/accountAggregate"
	"github.com/tcc-uniftec-5s/internal/infra/database/datastructure"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type credentialRepository struct {
	db *gorm.DB
}

func NewCredentialRepository(db *gorm.DB) account_aggregate.CredentialRepository {
	return &credentialRepository{
		db: db,
	}
}

func (r credentialRepository) Save(ctx context.Context, credential *account_aggregate.CredentialEntity) error {
	dbconn := r.db
	ctxValue, ok := ctx.Value(CtxKey{}).(CtxValue)
	if ok {
		dbconn = ctxValue.tx
	}

	credentialDS := datastructure.Credential{
		Email:    credential.Email,
		Password: credential.Password,
	}

	err := dbconn.
		WithContext(ctx).
		Table("credentials").
		Create(&credentialDS).
		Error

	credential.ID = &credentialDS.ID.Int64
	return err
}

func (r credentialRepository) Update(ctx context.Context, credential *account_aggregate.CredentialEntity) error {
	dbconn := r.db
	ctxValue, ok := ctx.Value(CtxKey{}).(CtxValue)
	if ok {
		dbconn = ctxValue.tx
	}

	credentialDS := datastructure.Credential{
		ID:        null.IntFromPtr(credential.ID),
		AccountId: null.IntFromPtr(credential.Account.ID),
		Email:     credential.Email,
		Password:  credential.Password,
	}

	err := dbconn.
		WithContext(ctx).
		Table("credentials").
		Where("id = ?", credential.ID).
		Save(&credentialDS).
		Error

	credential.Account.ID = &credentialDS.AccountId.Int64
	return err
}
