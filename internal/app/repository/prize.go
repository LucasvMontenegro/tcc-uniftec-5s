package repository

import (
	"context"

	"github.com/tcc-uniftec-5s/internal/domain/entity"
	"github.com/tcc-uniftec-5s/internal/infra/database/datastructure"
	"gorm.io/gorm"
)

type prize struct {
	db *gorm.DB
}

func NewPrize(db *gorm.DB) entity.PrizeRepository {
	return &prize{
		db: db,
	}
}

func (r prize) Save(ctx context.Context, prize *entity.Prize) error {
	dbconn := r.db
	ctxValue, ok := ctx.Value(CtxKey{}).(CtxValue)
	if ok {
		dbconn = ctxValue.tx
	}

	prizeDS := datastructure.Prize{
		Name:        &prize.Name,
		Description: prize.Description,
		EditionID:   prize.Edition.ID,
	}

	err := dbconn.
		WithContext(ctx).
		Table("prizes").
		Create(&prizeDS).
		Error

	prize.ID = prizeDS.ID
	return err
}
