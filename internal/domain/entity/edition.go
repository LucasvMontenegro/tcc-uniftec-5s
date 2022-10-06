package entity

import (
	"context"
	"errors"
	"time"
)

var ErrNoCurrentEditionFound = errors.New("no current edition found")
var ErrInvalidEditionDate = errors.New("start date must be before end date")

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
	GetCurrent(ctx context.Context, edition *Edition) error
}

type EditionFactoryInterface interface {
	NewEdition(name string, description *string, startDate, endDate time.Time) EditionInterface
	GetCurrent(ctx context.Context) (EditionInterface, error)
}
