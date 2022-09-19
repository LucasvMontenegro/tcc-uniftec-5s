package datastructure

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type SessionHistory struct {
	ID        null.Int `gorm:"primarykey"`
	AccountID null.Int
	CreatedAt time.Time
	ExpiresAt time.Time
}
