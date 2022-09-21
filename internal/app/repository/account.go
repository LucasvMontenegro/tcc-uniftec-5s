package repository

import (
	"context"
	"time"

	"github.com/tcc-uniftec-5s/internal/domain/entity"
	"github.com/tcc-uniftec-5s/internal/infra/database/datastructure"
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

	createdAt := time.Now()
	accountDS := datastructure.Account{
		ID:           account.ID,
		CredentialID: account.Credential.ID,
		Email:        &account.Credential.Email,
		CreatedAt:    &createdAt,
	}

	err := dbconn.
		WithContext(ctx).
		Table("accounts").
		Create(&accountDS).
		Error

	account.ID = accountDS.ID
	return err
}

func (r accountRepository) SetUser(ctx context.Context, account *entity.Account) error {
	err := r.updates(ctx, account.ID, map[string]interface{}{
		"user_id":    account.User.ID,
		"updated_at": time.Now(),
	})

	return err
}

func (r accountRepository) updates(ctx context.Context, id interface{}, fields map[string]interface{}) error {
	dbconn := r.db
	ctxValue, ok := ctx.Value(CtxKey{}).(CtxValue)
	if ok {
		dbconn = ctxValue.tx
	}

	err := dbconn.
		WithContext(ctx).
		Table("accounts").
		Where("id = ?", id).
		Updates(fields).
		Error

	return err
}
