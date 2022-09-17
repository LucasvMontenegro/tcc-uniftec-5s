package datastructure

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type Account struct {
	ID           null.Int `gorm:"primarykey"`
	CredentialID null.Int
	UserID       null.Int
	Email        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
