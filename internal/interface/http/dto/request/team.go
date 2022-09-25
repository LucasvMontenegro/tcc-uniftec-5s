package request

type CreateTeam struct {
	Name string `json:"name" validate:"required"`
}
