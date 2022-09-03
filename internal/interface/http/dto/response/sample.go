package response

import (
	"time"

	"github.com/tcc-uniftec-5s/internal/infra/database/entity"
	"gopkg.in/guregu/null.v4"
)

type Sample struct {
	ReferenceUUID string    `json:"reference_uuid"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     null.Time `json:"deleted_at" extensions:"x-nullable"`
}

func (c *Sample) FromSample(sample entity.Sample) {
	c.ReferenceUUID = sample.ReferenceUUID
	c.CreatedAt = sample.CreatedAt
	c.UpdatedAt = sample.UpdatedAt
	// c.DeletedAt = sample.DeletedAt
}