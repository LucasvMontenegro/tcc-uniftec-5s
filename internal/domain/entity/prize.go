package entity

import "context"

type Prize struct {
	ID          *int64
	Edition     *Edition
	Name        string
	Description *string
}

type PrizeInterface interface {
	Create(ctx context.Context, edition *Edition) error
}

type PrizeRepository interface {
	Save(ctx context.Context, prize *Prize) error
}

type PrizeFactoryInterface interface {
	NewPrize(name string, description *string, edition *Edition) PrizeInterface
}
