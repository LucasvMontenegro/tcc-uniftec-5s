package datastructure

import (
	"time"
)

type Credential struct {
	ID        *int64 `gorm:"primarykey"`
	AccountId *int64
	Email     *string
	Password  *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
