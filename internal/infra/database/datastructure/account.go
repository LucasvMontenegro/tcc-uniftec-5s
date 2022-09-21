package datastructure

import (
	"time"
)

type Account struct {
	ID           *int64 `gorm:"primarykey"`
	CredentialID *int64
	UserID       *int64
	Email        *string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}
