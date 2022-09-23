package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	usecase "github.com/tcc-uniftec-5s/internal/app/use_case"
	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

func NewUser(
	instance *echo.Echo,
	listTeamlessUsersUseCase usecase.ListTeamlessUsers) User {

	return &user{
		Instance:                 instance,
		listTeamlessUsersUseCase: listTeamlessUsersUseCase,
	}
}

type User interface {
	Router
	ListTeamlessUsers() func(c echo.Context) error
}

type user struct {
	Instance                 *echo.Echo
	listTeamlessUsersUseCase usecase.ListTeamlessUsers
}

func (c user) RegisterRoutes() {
	c.Instance.GET("/users", c.ListTeamlessUsers())
}

func (uc user) ListTeamlessUsers() func(c echo.Context) error {
	return func(c echo.Context) error {
		log.Info().Msg("/users")

		u, err := uc.listTeamlessUsersUseCase.Execute(c.Request().Context())
		if err != nil {
			log.Error().Msg("/users error")
			return c.NoContent(http.StatusInternalServerError) // todo handle error
		}

		response := []entity.User{}
		for _, e := range u {
			response = append(response, *e.Self())
		}

		log.Info().Msg("listing teamless users success")
		return c.JSON(http.StatusOK, response)
	}
}
