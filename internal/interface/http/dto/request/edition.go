package request

type CreateEdition struct {
	Edition struct {
		Name        string  `json:"name" validate:"required"`
		Description *string `json:"description"`
		Status      *string `json:"status"`
		StartDate   string  `json:"start_date" validate:"required"`
		EndDate     string  `json:"end_date" validate:"required"`
	} `json:"edition" validate:"required"`

	Prize struct {
		Name        string  `json:"name" validate:"required"`
		Description *string `json:"description"`
	} `json:"prize" validate:"required"`
}
