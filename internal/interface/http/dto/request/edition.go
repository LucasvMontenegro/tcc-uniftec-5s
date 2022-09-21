package request

import "time"

type CreateEdition struct {
	Edition struct {
		Name        string    `json:"name" validate:"required"`
		Description *string   `json:"description"`
		StartDate   time.Time `json:"start_date" validate:"required"`
		EndDate     time.Time `json:"end_date" validate:"required"`
	} `json:"edition" validate:"required"`

	Prize struct {
		Name        string  `json:"name" validate:"required"`
		Description *string `json:"description"`
	} `json:"prize" validate:"required"`
}
