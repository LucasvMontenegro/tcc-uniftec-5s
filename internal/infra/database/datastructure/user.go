package datastructure

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type User struct {
	ID        null.Int `gorm:"primarykey"`
	AccountID null.Int
	Name      string
	Status    string
	IsAdmin   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
