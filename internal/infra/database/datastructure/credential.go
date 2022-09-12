package datastructure

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type Credential struct {
	ID        null.Int `gorm:"primarykey"`
	AccountId null.Int
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
