package response

import (
	"github.com/tcc-uniftec-5s/internal/domain/entity"
	"github.com/tcc-uniftec-5s/internal/infra/utils"
)

func NewListedEditions(editions []entity.Edition) []Edition {
	var response []Edition
	for _, edition := range editions {
		sd := utils.FormatStringDate(edition.StartDate)
		ed := utils.FormatStringDate(edition.EndDate)

		response = append(response, Edition{
			ID:          edition.ID,
			Name:        &edition.Name,
			Description: edition.Description,
			Status:      edition.Status,
			StartDate:   &sd,
			EndDate:     &ed,
		})
	}

	return response
}

type Edition struct {
	ID          *int64  `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
	StartDate   *string `json:"start_date"`
	EndDate     *string `json:"end_date"`
}
