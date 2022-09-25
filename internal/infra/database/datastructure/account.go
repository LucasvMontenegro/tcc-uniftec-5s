package datastructure

import (
	"time"
)

type Account struct {
	ID           *int64 `gorm:"primarykey"`
	CredentialID *int64
	Credential   *Credential
	UserID       *int64
	User         *User
	Email        *string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}
