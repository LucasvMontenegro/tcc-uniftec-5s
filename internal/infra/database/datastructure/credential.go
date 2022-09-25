package datastructure

import (
	"time"
)

type Credential struct {
	ID        *int64 `gorm:"primarykey"`
	AccountID *int64
	Account   *Account
	Email     *string
	Password  *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
