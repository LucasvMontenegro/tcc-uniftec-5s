package controller

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	usecase "github.com/tcc-uniftec-5s/internal/app/use_case"
	"github.com/tcc-uniftec-5s/internal/interface/http/dto/request"
	"github.com/tcc-uniftec-5s/internal/interface/http/dto/response"
	"schneider.vip/problem"
)

func NewScore(
	instance *echo.Echo,
	restricted *echo.Group,
	accessValidator AccessValidator,
	listScoresUseCase usecase.ListScores) Score {

	return &score{
		Instance:          instance,
		restricted:        restricted,
		accessValidator:   accessValidator,
		listScoresUseCase: listScoresUseCase,
	}
}

type Score interface {
	Router
}

type score struct {
	Instance          *echo.Echo
	restricted        *echo.Group
	accessValidator   AccessValidator
	listScoresUseCase usecase.ListScores
}

func (c score) RegisterRoutes() {
	c.Instance.GET("/teams/:teamid/scores", c.listScores())
	c.restricted.POST("/teams/:teamid/scores", c.score())
}

func (sc score) listScores() func(c echo.Context) error {
	return func(c echo.Context) error {
		log.Info().Msg("GET /teams/:teamid/scores")
		teamID, _ := strconv.ParseInt(c.Param("teamid"), 10, 64)

		scores, err := sc.listScoresUseCase.Execute(c.Request().Context(), teamID)
		if err != nil {
			log.Error().Msg("GET /teams/:teamid/scores")
			return sc.handleErr(c, err)
		}

		response := response.NewListedScores(scores)
		return c.JSON(http.StatusOK, response)
	}
}

func (sc score) score() func(c echo.Context) error {
	return func(c echo.Context) error {
		log.Info().Msg("POST /teams/:teamid/scores")
		teamID, _ := strconv.ParseInt(c.Param("teamid"), 10, 64)
		log.Info().Msgf("POST /teams/%v/scores", teamID)

		if err := sc.accessValidator.Restrict(c); err != nil {
			c.Error(err)
			return nil
		}

		var req request.Scores

		if err := c.Bind(&req); err != nil {
			log.Info().Interface("new edition request", req).Msg("deserialization error")
			status := http.StatusBadRequest
			prob := problem.New(
				problem.Title("REQUEST_DESERIALIZATION_ERROR"),
				problem.Detail("Bad Request"),
				problem.Status(status),
			)

			return c.JSON(status, prob)
		}

		validate := validator.New()
		if err := validate.Struct(req); err != nil {
			log.Error().Interface("score request", req).Msg("/teams/id/scores validation error")
			status := http.StatusBadRequest
			prob := problem.New(
				problem.Title("REQUEST_VALIDATION_ERROR"),
				problem.Detail("Bad Request"),
				problem.Status(status),
			)

			return c.JSON(status, prob)
		}
		// err := sc.scoreUseCase.Execute(c.Request().Context(), teamID)
		// if err != nil {
		// 	log.Error().Msg("POST /teams/:teamid/scores")
		// 	return sc.handleErr(c, err)
		// }

		return c.NoContent(http.StatusNoContent)
	}
}

func (sc score) handleErr(c echo.Context, err error) error {
	var problemJSON *problem.Problem

	status := http.StatusInternalServerError
	detail := "internal server error"
	title := "internal server error"

	problemJSON = problem.New(
		problem.Status(status),
		problem.Title(title),
		problem.Detail(detail),
	)

	return c.JSON(status, problemJSON)
}
