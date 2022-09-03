package sample

import (
	"context"

	"github.com/tcc-uniftec-5s/internal/infra/database/entity"
)

type Service interface {
	CreateSample(ctx context.Context, createDTO CreateDTO) (entity.Sample, error)
	GetByReferenceUUID(ctx context.Context, referenceUUID string) (entity.Sample, error)
}

type CreateDTO struct {
	ReferenceUUID string
}
