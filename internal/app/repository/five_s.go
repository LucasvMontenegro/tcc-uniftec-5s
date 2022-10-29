package repository

import (
	"context"

	"github.com/tcc-uniftec-5s/internal/domain/entity"
	"github.com/tcc-uniftec-5s/internal/infra/database/datastructure"
	"gorm.io/gorm"
)

type fiveS struct {
	db *gorm.DB
}

func NewFiveS(db *gorm.DB) entity.FiveSRepository {
	return &fiveS{
		db: db,
	}
}

func (r fiveS) GetByName(ctx context.Context, fiveS *entity.FiveS) error {
	dbconn := r.db
	ctxValue, ok := ctx.Value(CtxKey{}).(CtxValue)
	if ok {
		dbconn = ctxValue.tx
	}

	fiveSDS := datastructure.FiveS{}

	err := dbconn.
		WithContext(ctx).
		Table("five_s").
		Where("name = ?", fiveS.Name).
		First(&fiveSDS).
		Error

	if err == nil {
		fiveS.ID = fiveSDS.ID
		fiveS.Name = fiveSDS.Name
	}

	return err
}
