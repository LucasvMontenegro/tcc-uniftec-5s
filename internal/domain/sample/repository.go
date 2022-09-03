package sample

import (
	"context"

	"github.com/tcc-uniftec-5s/internal/infra/database/entity"
)

type Repository interface {
	Create(ctx context.Context, sample *entity.Sample) error
	GetByReferenceUUID(ctx context.Context, sample *entity.Sample, reference string) error
}
