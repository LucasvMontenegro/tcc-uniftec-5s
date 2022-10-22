package repository

import (
	"context"

	"github.com/tcc-uniftec-5s/internal/domain/entity"
	"github.com/tcc-uniftec-5s/internal/infra/database/datastructure"
	"gorm.io/gorm"
)

type edition struct {
	db *gorm.DB
}

func NewEdition(db *gorm.DB) entity.EditionRepository {
	return &edition{
		db: db,
	}
}

func (r edition) Save(ctx context.Context, edition *entity.Edition) error {
	dbconn := r.db
	ctxValue, ok := ctx.Value(CtxKey{}).(CtxValue)
	if ok {
		dbconn = ctxValue.tx
	}

	editionDS := datastructure.Edition{
		Name:        &edition.Name,
		Description: edition.Description,
		Status:      edition.Status,
		StartDate:   &edition.StartDate,
		EndDate:     &edition.EndDate,
	}

	err := dbconn.
		WithContext(ctx).
		Table("editions").
		Create(&editionDS).
		Error

	edition.ID = editionDS.ID
	return err
}

func (r edition) GetCurrent(ctx context.Context, edition *entity.Edition) error {
	dbconn := r.db
	ctxValue, ok := ctx.Value(CtxKey{}).(CtxValue)
	if ok {
		dbconn = ctxValue.tx
	}

	editionDS := datastructure.Edition{}

	err := dbconn.
		WithContext(ctx).
		Table("editions").
		Where("status = ?", "ACTIVE").
		First(&editionDS).
		Error

	if err == nil {
		edition.ID = editionDS.ID
		edition.Name = *editionDS.Name
		edition.Description = editionDS.Description
		edition.StartDate = *editionDS.StartDate
		edition.EndDate = *editionDS.EndDate
		edition.Status = editionDS.Status
	}

	return err
}

func (r edition) ListEditions(ctx context.Context) ([]*entity.Edition, error) {
	dbconn := r.db
	ctxValue, ok := ctx.Value(CtxKey{}).(CtxValue)
	if ok {
		dbconn = ctxValue.tx
	}

	editionDS := []datastructure.Edition{}
	err := dbconn.
		WithContext(ctx).
		Table("editions").
		Find(&editionDS).
		Error

	editions := []*entity.Edition{}
	for _, edition := range editionDS {
		t := entity.Edition{
			ID:          edition.ID,
			Name:        *edition.Name,
			Description: edition.Description,
			StartDate:   *edition.StartDate,
			EndDate:     *edition.EndDate,
			Status:      edition.Status,
		}

		editions = append(editions, &t)
	}

	return editions, err
}
