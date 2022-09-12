package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	usecase "github.com/tcc-uniftec-5s/internal/app/use_case"
	"github.com/tcc-uniftec-5s/internal/interface/http/dto/request"
	"schneider.vip/problem"
)

func NewSignupController(
	instance *echo.Echo,
	signupUseCase usecase.Signup) SignupController {

	return &signupController{
		Instance:      instance,
		signupUseCase: signupUseCase,
	}
}

type SignupController interface {
	HTTPController
	Signup() func(c echo.Context) error
}

type signupController struct {
	Instance      *echo.Echo
	signupUseCase usecase.Signup
}

func (c signupController) RegisterRoutes() {
	c.Instance.POST("/signup", c.Signup())
}

func (sc signupController) Signup() func(c echo.Context) error {
	return func(c echo.Context) error {
		log.Info().Msg("/signup")

		var signupReq request.Signup

		if err := c.Bind(&signupReq); err != nil {
			log.Info().Interface("signup request", signupReq).Msg("deserialization error")
			status := http.StatusBadRequest
			prob := problem.New(
				problem.Title("REQUEST_DESERIALIZATION_ERROR"),
				problem.Detail("Bad Request"),
				problem.Status(status),
			)

			return c.JSON(status, prob)
		}

		// todo adicionar validacao de payload

		err := sc.signupUseCase.Signup(c.Request().Context(), signupReq.Email, signupReq.Password, signupReq.Name)
		if err != nil {
			log.Error().Interface("signup request", signupReq).Msg("/signup error")
			return c.NoContent(http.StatusInternalServerError) // todo handle error
		}

		log.Info().Msg("signup success")
		return c.NoContent(http.StatusNoContent)
	}
}
