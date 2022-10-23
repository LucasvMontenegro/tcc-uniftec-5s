package response

import (
	"github.com/tcc-uniftec-5s/internal/domain/entity"
	"github.com/tcc-uniftec-5s/internal/infra/utils"
)

func NewCreatedTeam(team entity.Team) CreatedTeam {
	sd := utils.FormatStringDate(team.Edition.StartDate)
	ed := utils.FormatStringDate(team.Edition.EndDate)

	return CreatedTeam{
		ID:   team.ID,
		Name: &team.Name,
		Edition: &Edition{
			ID:        team.Edition.ID,
			Name:      &team.Edition.Name,
			Status:    team.Edition.Status,
			StartDate: &sd,
			EndDate:   &ed,
		},
	}

}

func NewListedTeams(teams []entity.Team) []CreatedTeam {
	var response []CreatedTeam
	for _, team := range teams {
		tname := team.Name
		response = append(response, CreatedTeam{
			ID:   team.ID,
			Name: &tname,
		})
	}

	return response
}

type CreatedTeam struct {
	ID      *int64   `json:"id"`
	Name    *string  `json:"name"`
	Edition *Edition `json:"edition,omitempty"`
}
