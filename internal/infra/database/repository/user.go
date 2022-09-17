package repository

import (
	"context"

	account_aggregate "github.com/tcc-uniftec-5s/internal/domain/accountAggregate"
	"github.com/tcc-uniftec-5s/internal/infra/database/datastructure"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) account_aggregate.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r userRepository) Save(ctx context.Context, user *account_aggregate.UserEntity) error {
	dbconn := r.db
	ctxValue, ok := ctx.Value(CtxKey{}).(CtxValue)
	if ok {
		dbconn = ctxValue.tx
	}

	userDS := datastructure.User{
		AccountID: null.IntFromPtr(user.Account.ID),
		Name:      user.Name,
		Status:    string(user.Status),
		IsAdmin:   user.IsAdmin,
	}

	err := dbconn.
		WithContext(ctx).
		Table("users").
		Create(&userDS).
		Error

	user.ID = &userDS.ID.Int64
	return err
}
