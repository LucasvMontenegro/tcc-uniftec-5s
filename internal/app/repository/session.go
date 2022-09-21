package repository

import (
	"context"
	"time"

	"github.com/tcc-uniftec-5s/internal/domain/entity"
	"github.com/tcc-uniftec-5s/internal/infra/database/datastructure"
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

func (r sessionRepository) Save(ctx context.Context, session *entity.Session) error {
	// redis
	return nil
}

func (r sessionRepository) SaveHistory(ctx context.Context, session *entity.Session) error {
	dbconn := r.db
	ctxValue, ok := ctx.Value(CtxKey{}).(CtxValue)
	if ok {
		dbconn = ctxValue.tx
	}

	createdAt := time.Now()
	expiresAt := createdAt.Add(2 * time.Hour)
	sessionHistDS := datastructure.SessionHistory{
		AccountID: session.AccountEntity.ID,
		CreatedAt: &createdAt,
		ExpiresAt: &expiresAt,
	}

	err := dbconn.
		WithContext(ctx).
		Table("session_history").
		Create(&sessionHistDS).
		Error

	return err
}
