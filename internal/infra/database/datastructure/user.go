package datastructure

import (
	"time"
)

type User struct {
	ID        *int64 `gorm:"primarykey"`
	AccountID *int64
	Account   *Account
	Name      *string
	Status    *string
	IsAdmin   *bool
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
