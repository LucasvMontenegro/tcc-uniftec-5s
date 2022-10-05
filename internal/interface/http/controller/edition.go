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

func NewEdition(
	instance *echo.Echo,
	restricted *echo.Group,
	accessValidator AccessValidator,
	createEditionUseCase usecase.CreateEdition) Edition {

	return &edition{
		Instance:             instance,
		restricted:           restricted,
		accessValidator:      accessValidator,
		createEditionUseCase: createEditionUseCase,
	}
}

type Edition interface {
	Router
	CreateEdition() func(c echo.Context) error
}

type edition struct {
	Instance             *echo.Echo
	restricted           *echo.Group
	accessValidator      AccessValidator
	createEditionUseCase usecase.CreateEdition
}

func (c edition) RegisterRoutes() {
	c.Instance.POST("/editions", c.CreateEdition())
}

func (ec edition) CreateEdition() func(c echo.Context) error {
	return func(c echo.Context) error {
		log.Info().Msg("/editions")

		var createEditionReq request.CreateEdition

		if err := c.Bind(&createEditionReq); err != nil {
			log.Info().Interface("new edition request", createEditionReq).Msg("deserialization error")
			status := http.StatusBadRequest
			prob := problem.New(
				problem.Title("REQUEST_DESERIALIZATION_ERROR"),
				problem.Detail("Bad Request"),
				problem.Status(status),
			)

			return c.JSON(status, prob)
		}

		validate := validator.New()
		if err := validate.Struct(createEditionReq); err != nil {
			log.Error().Interface("new edition request", createEditionReq).Msg("/editions validation error")
			status := http.StatusBadRequest
			prob := problem.New(
				problem.Title("REQUEST_VALIDATION_ERROR"),
				problem.Detail("Bad Request"),
				problem.Status(status),
			)

			return c.JSON(status, prob)
		}

		dto := usecase.CreateEditionDTO{
			EditionDTO: usecase.EditionDTO{
				Name:        createEditionReq.Edition.Name,
				Description: createEditionReq.Edition.Description,
				StartDate:   createEditionReq.Edition.StartDate,
				EndDate:     createEditionReq.Edition.EndDate,
			},
			PrizeDTO: usecase.PrizeDTO{
				Name:        createEditionReq.Prize.Name,
				Description: createEditionReq.Prize.Description,
			},
		}

		err := ec.createEditionUseCase.Execute(c.Request().Context(), dto)
		if err != nil {
			log.Error().Interface("new edition request", createEditionReq).Msg("/editions error")
			return c.NoContent(http.StatusInternalServerError) // todo handle error
		}

		log.Info().Msg("new edition success")
		return c.NoContent(http.StatusNoContent)
	}
}
