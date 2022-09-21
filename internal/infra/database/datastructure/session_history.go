package datastructure

import (
	"time"
)

type SessionHistory struct {
	ID        *int64 `gorm:"primarykey"`
	AccountID *int64
	CreatedAt *time.Time
	ExpiresAt *time.Time
}
