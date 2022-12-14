package controller

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	usecase "github.com/tcc-uniftec-5s/internal/app/use_case"
	"github.com/tcc-uniftec-5s/internal/domain/entity"
	"github.com/tcc-uniftec-5s/internal/interface/http/dto/request"
	"github.com/tcc-uniftec-5s/internal/interface/http/dto/response"
	"schneider.vip/problem"
)

func NewTeam(
	instance *echo.Echo,
	restricted *echo.Group,
	accessValidator AccessValidator,
	createTeamUseCase usecase.CreateTeam,
	listTeamsUseCase usecase.ListTeams) Team {

	return &team{
		instance:          instance,
		restricted:        restricted,
		accessValidator:   accessValidator,
		createTeamUseCase: createTeamUseCase,
		listTeamsUseCase:  listTeamsUseCase,
	}
}

type Team interface {
	Router
}

type team struct {
	instance          *echo.Echo
	restricted        *echo.Group
	accessValidator   AccessValidator
	createTeamUseCase usecase.CreateTeam
	listTeamsUseCase  usecase.ListTeams
}

func (c team) RegisterRoutes() {
	c.restricted.POST("/edition/teams", c.CreateTeam())
	c.instance.GET("/edition/teams", c.listTeams())
}

func (tc team) CreateTeam() func(c echo.Context) error {
	return func(c echo.Context) error {
		log.Info().Msg("POST /edition/teams")

		if err := tc.accessValidator.Restrict(c); err != nil {
			c.Error(err)
			return nil
		}

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

		team, err := tc.createTeamUseCase.Execute(c.Request().Context(), createTeamReq.Name)
		if err != nil {
			log.Error().Interface("new team request", createTeamReq).Msg("/edition/teams error")
			return tc.handleErr(c, err)
		}

		response := response.NewCreatedTeam(*team)
		log.Info().Msg("new team success")
		return c.JSON(http.StatusOK, response)
	}
}

func (tc team) listTeams() func(c echo.Context) error {
	return func(c echo.Context) error {
		log.Info().Msg("GET /edition/teams")

		teams, err := tc.listTeamsUseCase.Execute(c.Request().Context())
		if err != nil {
			log.Error().Msg("GET /edition/teams error")
			return tc.handleErr(c, err)
		}

		response := response.NewListedTeams(teams)
		return c.JSON(http.StatusOK, response)
	}
}

func (tc team) handleErr(c echo.Context, err error) error {
	var problemJSON *problem.Problem

	status := http.StatusInternalServerError
	detail := "internal server error"
	title := "internal server error"

	switch err {
	case entity.ErrNoCurrentEditionFound:
		status = http.StatusNotFound
		title = "not found"
		detail = "current edition not found"
	}

	problemJSON = problem.New(
		problem.Status(status),
		problem.Title(title),
		problem.Detail(detail),
	)

	return c.JSON(status, problemJSON)
}
