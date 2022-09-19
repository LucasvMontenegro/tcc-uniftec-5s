package controller

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	usecase "github.com/tcc-uniftec-5s/internal/app/use_case"
	"github.com/tcc-uniftec-5s/internal/interface/http/dto/request"
	"github.com/tcc-uniftec-5s/internal/interface/http/dto/response"
	"schneider.vip/problem"
)

func NewLoginController(
	instance *echo.Echo,
	loginUseCase usecase.Login) LoginController {

	return &loginController{
		Instance:     instance,
		loginUseCase: loginUseCase,
	}
}

type LoginController interface {
	HTTPController
	Login() func(c echo.Context) error
}

type loginController struct {
	Instance      *echo.Echo
	signupUseCase usecase.Signup
	loginUseCase  usecase.Login
}

func (c loginController) RegisterRoutes() {
	c.Instance.POST("/login", c.Login())
}

func (lc loginController) Login() func(c echo.Context) error {
	return func(c echo.Context) error {
		log.Info().Msg("/login")

		var loginReq request.Login

		if err := c.Bind(&loginReq); err != nil {
			log.Info().Interface("login request", loginReq).Msg("deserialization error")
			status := http.StatusBadRequest
			prob := problem.New(
				problem.Title("REQUEST_DESERIALIZATION_ERROR"),
				problem.Detail("Bad Request"),
				problem.Status(status),
			)

			return c.JSON(status, prob)
		}

		validate := validator.New()
		if err := validate.Struct(loginReq); err != nil {
			log.Error().Interface("login request", loginReq).Msg("/login validation error")
			status := http.StatusBadRequest
			prob := problem.New(
				problem.Title("REQUEST_VALIDATION_ERROR"),
				problem.Detail("Bad Request"),
				problem.Status(status),
			)

			return c.JSON(status, prob)
		}

		jwt, err := lc.loginUseCase.Execute(c.Request().Context(), loginReq.Email, loginReq.Password)
		if err != nil {
			log.Error().Interface("login request", loginReq).Msg("/login error")
			return c.NoContent(http.StatusInternalServerError) // todo handle error
		}

		log.Info().Msg("login success")
		return c.JSON(200, response.Login{JWT: *jwt})
	}
}
