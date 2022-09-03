package entity

import (
	"gorm.io/gorm"
)

type Sample struct {
	gorm.Model
	ReferenceUUID string
}
