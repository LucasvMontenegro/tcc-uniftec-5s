package entity

import "context"

type Team struct {
	Id      *int64
	Name    string
	Edition *Edition
	Users   []*User
	Scores  *string // todo
}

type TeamInterface interface {
	Self() *Team
	Create(ctx context.Context, edition *Edition) error
}

type TeamRepository interface {
	Save(ctx context.Context, team *Team, edition *Edition) error
}

type TeamFactoryInterface interface {
	NewTeam(ctx context.Context, name string, edition *Edition) TeamInterface
}
