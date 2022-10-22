package repository

import (
	"context"

	"github.com/tcc-uniftec-5s/internal/domain/entity"
	"github.com/tcc-uniftec-5s/internal/infra/database/datastructure"
	"gorm.io/gorm"
)

type score struct {
	db *gorm.DB
}

func NewScore(db *gorm.DB) entity.ScoreRepository {
	return &score{
		db: db,
	}
}

func (r score) ListScores(ctx context.Context, teamID int64) ([]*entity.Score, error) {
	dbconn := r.db
	ctxValue, ok := ctx.Value(CtxKey{}).(CtxValue)
	if ok {
		dbconn = ctxValue.tx
	}

	scoreDS := []datastructure.Score{}
	err := dbconn.
		WithContext(ctx).
		Table("scores").
		Where("team_id = ?", teamID).
		Preload("Team").
		Preload("FiveS").
		Find(&scoreDS).
		Error

	scores := []*entity.Score{}
	for _, score := range scoreDS {
		s := entity.Score{
			ID:      score.ID,
			FiveSID: score.FiveSID,
			FiveS: &entity.FiveS{
				ID:   score.FiveS.ID,
				Name: *score.FiveS.Name,
			},
			TeamID: score.TeamID,
			Team: &entity.Team{
				ID:   score.TeamID,
				Name: score.Team.Name,
			},
			Score: score.Score,
		}

		scores = append(scores, &s)
	}

	return scores, err
}
