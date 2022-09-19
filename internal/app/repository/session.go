package repository

import (
	"context"
	"time"

	"github.com/tcc-uniftec-5s/internal/domain/entity"
	"github.com/tcc-uniftec-5s/internal/infra/database/datastructure"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type sessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) entity.SessionRepository {
	return &sessionRepository{
		db: db,
	}
}

func (r sessionRepository) Save(ctx context.Context, session *entity.SessionEntity) error {
	// redis
	return nil
}

func (r sessionRepository) SaveHistory(ctx context.Context, session *entity.SessionEntity) error {
	dbconn := r.db
	ctxValue, ok := ctx.Value(CtxKey{}).(CtxValue)
	if ok {
		dbconn = ctxValue.tx
	}

	sessionHistDS := datastructure.SessionHistory{
		AccountID: null.IntFromPtr(session.AccountEntity.ID),
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(2 * time.Hour),
	}

	err := dbconn.
		WithContext(ctx).
		Table("session_history").
		Create(&sessionHistDS).
		Error

	return err
}
