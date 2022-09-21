package repository

import (
	"context"

	"github.com/tcc-uniftec-5s/internal/domain/entity"
	"github.com/tcc-uniftec-5s/internal/infra/database/datastructure"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) entity.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r userRepository) Save(ctx context.Context, user *entity.User) error {
	dbconn := r.db
	ctxValue, ok := ctx.Value(CtxKey{}).(CtxValue)
	if ok {
		dbconn = ctxValue.tx
	}

	status := string(user.Status)
	userDS := datastructure.User{
		AccountID: user.Account.ID,
		Name:      &user.Name,
		Status:    &status,
		IsAdmin:   &user.IsAdmin,
	}

	err := dbconn.
		WithContext(ctx).
		Table("users").
		Create(&userDS).
		Error

	user.ID = userDS.ID
	return err
}
