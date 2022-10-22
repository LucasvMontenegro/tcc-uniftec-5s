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
	Score   *float64
}

type ScoreInterface interface {
	Self() *Score
}

type ScoreRepository interface {
	ListScores(ctx context.Context, teamID int64) ([]*Score, error)
}

type ScoreFactoryInterface interface {
	ListScores(ctx context.Context, teamID int64) ([]ScoreInterface, error)
}
