package entity

import (
	"context"
)

type TeamScore struct {
	TeamName string
	Scores   []*Score
}

type Ranking struct {
	*Edition
	*Prize
	TeamScores []*TeamScore
}

type RankingInterface interface {
	Self() *Ranking
}

type RankingFactoryInterface interface {
	Generate(ctx context.Context, currentEdition *Edition) (RankingInterface, error)
}
