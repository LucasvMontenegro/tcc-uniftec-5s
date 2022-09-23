package entity

import (
	"context"
	"time"
)

type Edition struct {
	ID          *int64
	Winner      *Team
	Name        string
	Description *string
	Status      *string
	StartDate   time.Time
	EndDate     time.Time
}

type EditionInterface interface {
	Self() *Edition
	Create(ctx context.Context) error
}

type EditionRepository interface {
	Save(ctx context.Context, edition *Edition) error
}

type EditionFactoryInterface interface {
	NewEdition(name string, description *string, startDate, endDate time.Time) EditionInterface
}
