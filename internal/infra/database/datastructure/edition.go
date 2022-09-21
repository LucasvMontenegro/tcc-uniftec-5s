package datastructure

import (
	"time"
)

type Edition struct {
	ID          *int64 `gorm:"primarykey"`
	WinnerID    *int64
	Name        *string
	Description *string
	StartDate   *time.Time
	EndDate     *time.Time
}
