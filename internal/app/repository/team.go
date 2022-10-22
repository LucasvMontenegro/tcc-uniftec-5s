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
		EditionID: edition.ID,
	}

	err := dbconn.
		WithContext(ctx).
		Table("teams").
		Create(&teamDS).
		Error

	team.ID = teamDS.ID
	return err
}

func (r team) ListTeamsByEdition(ctx context.Context, edition *entity.Edition) ([]*entity.Team, error) {
	dbconn := r.db
	ctxValue, ok := ctx.Value(CtxKey{}).(CtxValue)
	if ok {
		dbconn = ctxValue.tx
	}

	teamDS := []datastructure.Team{}
	err := dbconn.
		WithContext(ctx).
		Table("teams").
		Where("edition_id = ?", edition.ID).
		Find(&teamDS).
		Error

	teams := []*entity.Team{}
	for _, team := range teamDS {
		t := entity.Team{
			ID:   team.ID,
			Name: team.Name,
		}

		teams = append(teams, &t)
	}

	return teams, err
}
