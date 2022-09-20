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

func NewResetPasswordController(
	instance *echo.Echo,
	resetPasswordUseCase usecase.ResetPassword) ResetPassword {

	return &resetPassword{
		Instance:             instance,
		resetPasswordUseCase: resetPasswordUseCase,
	}
}

type ResetPassword interface {
	Router
	ResetPassword() func(c echo.Context) error
}

type resetPassword struct {
	Instance             *echo.Echo
	resetPasswordUseCase usecase.ResetPassword
}

func (c resetPassword) RegisterRoutes() {
	c.Instance.POST("/reset-password", c.ResetPassword())
}

func (lc resetPassword) ResetPassword() func(c echo.Context) error {
	return func(c echo.Context) error {
		log.Info().Msg("/reset-password")

		var resetPasswordReq request.ResetPassword

		if err := c.Bind(&resetPasswordReq); err != nil {
			log.Info().Interface("login request", resetPasswordReq).Msg("deserialization error")
			status := http.StatusBadRequest
			prob := problem.New(
				problem.Title("REQUEST_DESERIALIZATION_ERROR"),
				problem.Detail("Bad Request"),
				problem.Status(status),
			)

			return c.JSON(status, prob)
		}

		validate := validator.New()
		if err := validate.Struct(resetPasswordReq); err != nil {
			log.Error().Interface("login request", resetPasswordReq).Msg("/login validation error")
			status := http.StatusBadRequest
			prob := problem.New(
				problem.Title("REQUEST_VALIDATION_ERROR"),
				problem.Detail("Bad Request"),
				problem.Status(status),
			)

			return c.JSON(status, prob)
		}

		err := lc.resetPasswordUseCase.Execute(c.Request().Context(), resetPasswordReq.Email, resetPasswordReq.Password)
		if err != nil {
			log.Error().Interface("reset-password request", resetPasswordReq).Msg("/login error")
			return c.NoContent(http.StatusInternalServerError) // todo handle error
		}

		log.Info().Msg("reset-password success")
		return c.NoContent(http.StatusNoContent)
	}
}
