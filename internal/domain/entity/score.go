package entity

import (
	"context"
)

type Score struct {
	ID      *int64
	FiveS   *FiveS
	FiveSID *int64
	Team    *Team
	TeamID  *int64
	Score   *int
}

type ScoreInterface interface {
	Self() *Score
	Create(ctx context.Context) error
}

type ScoreRepository interface {
	ListScores(ctx context.Context, teamID int64) ([]*Score, error)
	Save(ctx context.Context, score *Score) error
}

type ScoreFactoryInterface interface {
	ListScores(ctx context.Context, teamID int64) ([]ScoreInterface, error)
	New(ctx context.Context, team *Team, fives *FiveS, score int) ScoreInterface
}
