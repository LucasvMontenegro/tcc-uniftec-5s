package repository

import (
	"context"
	"time"

	"github.com/tcc-uniftec-5s/internal/domain/entity"
	"github.com/tcc-uniftec-5s/internal/infra/database/datastructure"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) entity.AccountRepository {
	return &accountRepository{
		db: db,
	}
}

func (r accountRepository) Save(ctx context.Context, account *entity.Account) error {
	dbconn := r.db
	ctxValue, ok := ctx.Value(CtxKey{}).(CtxValue)
	if ok {
		dbconn = ctxValue.tx
	}

	accountDS := datastructure.Account{
		ID:           null.IntFromPtr(account.ID),
		CredentialID: null.IntFromPtr(account.Credential.ID),
		Email:        account.Credential.Email,
		CreatedAt:    time.Now(),
	}

	err := dbconn.
		WithContext(ctx).
		Table("accounts").
		Create(&accountDS).
		Error

	account.ID = &accountDS.ID.Int64
	return err
}

func (r accountRepository) Update(ctx context.Context, account *entity.Account) error {
	dbconn := r.db
	ctxValue, ok := ctx.Value(CtxKey{}).(CtxValue)
	if ok {
		dbconn = ctxValue.tx
	}

	accountDS := datastructure.Account{
		ID:           null.IntFromPtr(account.ID),
		CredentialID: null.IntFromPtr(account.Credential.ID),
		Email:        account.Credential.Email,
		UserID:       null.IntFromPtr(account.User.ID),
	}

	err := dbconn.
		WithContext(ctx).
		Table("accounts").
		Where("id = ?", account.ID).
		Save(&accountDS).
		Error

	return err
}
