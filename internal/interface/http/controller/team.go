package controller

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	usecase "github.com/tcc-uniftec-5s/internal/app/use_case"
	"github.com/tcc-uniftec-5s/internal/interface/http/dto/request"
	"schneider.vip/problem"
)

func NewTeam(
	instance *echo.Echo,
	createTeamUseCase usecase.CreateTeam) Team {

	return &team{
		Instance:          instance,
		createTeamUseCase: createTeamUseCase,
	}
}

type Team interface {
	Router
	CreateTeam() func(c echo.Context) error
}

type team struct {
	Instance          *echo.Echo
	createTeamUseCase usecase.CreateTeam
}

func (c team) RegisterRoutes() {
	c.Instance.POST("/edition/teams", c.CreateTeam())
}

func (ec team) CreateTeam() func(c echo.Context) error {
	return func(c echo.Context) error {
		log.Info().Msg("/edition/teams")

		var createTeamReq request.CreateTeam

		if err := c.Bind(&createTeamReq); err != nil {
			log.Info().Interface("new team request", createTeamReq).Msg("deserialization error")
			status := http.StatusBadRequest
			prob := problem.New(
				problem.Title("REQUEST_DESERIALIZATION_ERROR"),
				problem.Detail("Bad Request"),
				problem.Status(status),
			)

			return c.JSON(status, prob)
		}

		validate := validator.New()
		if err := validate.Struct(createTeamReq); err != nil {
			log.Error().Interface("new team request", createTeamReq).Msg("/edition/teams validation error")
			status := http.StatusBadRequest
			prob := problem.New(
				problem.Title("REQUEST_VALIDATION_ERROR"),
				problem.Detail("Bad Request"),
				problem.Status(status),
			)

			return c.JSON(status, prob)
		}

		team, err := ec.createTeamUseCase.Execute(c.Request().Context(), createTeamReq.Name)
		if err != nil {
			log.Error().Interface("new team request", createTeamReq).Msg("/edition/teams error")
			return c.NoContent(http.StatusInternalServerError) // todo handle error
		}

		log.Info().Msg("new team success")
		return c.JSON(http.StatusOK, team)
	}
}
