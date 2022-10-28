package repository

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgconn"
	"github.com/tcc-uniftec-5s/internal/domain/entity"
	"github.com/tcc-uniftec-5s/internal/infra/database"
	"github.com/tcc-uniftec-5s/internal/infra/database/datastructure"
	"gorm.io/gorm"
)

type credentialRepository struct {
	db *gorm.DB
}

func NewCredentialRepository(db *gorm.DB) entity.CredentialRepository {
	return &credentialRepository{
		db: db,
	}
}

func (r credentialRepository) Save(ctx context.Context, credential *entity.Credential) error {
	dbconn := r.db
	ctxValue, ok := ctx.Value(CtxKey{}).(CtxValue)
	if ok {
		dbconn = ctxValue.tx
	}

	createdAt := time.Now()
	credentialDS := datastructure.Credential{
		Email:     &credential.Email,
		Password:  &credential.Password,
		CreatedAt: &createdAt,
	}

	err := dbconn.
		WithContext(ctx).
		Table("credentials").
		Create(&credentialDS).
		Error

	if err != nil {
		var pgErr *pgconn.PgError
		ok := errors.As(err, &pgErr)
		if ok && pgErr.Code == database.ErrCodes[database.ErrUniqueViolation] {
			return entity.ErrCredentialAlreadyExists
		}

		return err
	} else {
		credential.ID = credentialDS.ID
	}

	return nil
}

func (r credentialRepository) SetAccount(ctx context.Context, credential *entity.Credential) error {
	err := r.updates(ctx, credential.ID, map[string]interface{}{
		"account_id": credential.Account.ID,
		"updated_at": time.Now(),
	})

	return err
}

func (r credentialRepository) Identify(ctx context.Context, credential *entity.Credential) error {
	dbconn := r.db
	ctxValue, ok := ctx.Value(CtxKey{}).(CtxValue)
	if ok {
		dbconn = ctxValue.tx
	}

	credentialDS := datastructure.Credential{}

	err := dbconn.
		WithContext(ctx).
		Table("credentials").
		Where("email = ? and password = ?", credential.Email, credential.Password).
		Preload("Account").
		Preload("Account.User").
		First(&credentialDS).
		Error

	if err == gorm.ErrRecordNotFound {
		err = entity.ErrCredentialNotFound
	}

	if err == nil {
		userDS := credentialDS.Account.User

		credential.ID = credentialDS.ID
		credential.Account = &entity.Account{
			ID: credentialDS.Account.ID,
			User: &entity.User{
				ID:      userDS.ID,
				Name:    *userDS.Name,
				IsAdmin: *userDS.IsAdmin,
			},
		}
	}

	return err
}

func (r credentialRepository) UpdatePassword(ctx context.Context, credential *entity.Credential) error {
	dbconn := r.db
	ctxValue, ok := ctx.Value(CtxKey{}).(CtxValue)
	if ok {
		dbconn = ctxValue.tx
	}

	err := dbconn.
		WithContext(ctx).
		Table("credentials").
		Where("email = ?", credential.Email).
		Updates(map[string]interface{}{"password": credential.Password, "updated_at": time.Now()}).
		Error

	return err
}

func (r credentialRepository) updates(ctx context.Context, id interface{}, fields map[string]interface{}) error {
	dbconn := r.db
	ctxValue, ok := ctx.Value(CtxKey{}).(CtxValue)
	if ok {
		dbconn = ctxValue.tx
	}

	err := dbconn.
		WithContext(ctx).
		Table("credentials").
		Where("id = ?", id).
		Updates(fields).
		Error

	return err
}
