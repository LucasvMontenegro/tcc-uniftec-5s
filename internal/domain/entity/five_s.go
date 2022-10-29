package entity

import "context"

type FiveS struct {
	ID          *int64
	Name        *string
	Description *string
}

type FiveSInterface interface {
	Self() *FiveS
}

type FiveSRepository interface {
	GetByName(ctx context.Context, fiveS *FiveS) error
}

type FiveSFactoryInterface interface {
	GetByName(ctx context.Context, name string) (FiveSInterface, error)
}
