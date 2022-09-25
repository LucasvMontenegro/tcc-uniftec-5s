package repository

import (
	"context"

	"github.com/tcc-uniftec-5s/internal/domain/entity"
	"github.com/tcc-uniftec-5s/internal/infra/database/datastructure"
	"gorm.io/gorm"
)

type team struct {
	db *gorm.DB
}

func NewTeam(db *gorm.DB) entity.TeamRepository {
	return &team{
		db: db,
	}
}

func (r team) Save(ctx context.Context, team *entity.Team, edition *entity.Edition) error {
	dbconn := r.db
	ctxValue, ok := ctx.Value(CtxKey{}).(CtxValue)
	if ok {
		dbconn = ctxValue.tx
	}

	teamDS := datastructure.Team{
		Name:      team.Name,
		EditionId: edition.ID,
	}

	err := dbconn.
		WithContext(ctx).
		Table("teams").
		Create(&teamDS).
		Error

	team.Id = teamDS.Id
	return err
}
