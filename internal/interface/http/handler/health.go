package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateHealthHandler() func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	}
}

func MakeHealtHandler(api *echo.Echo) {
	api.Any("health", CreateHealthHandler())
}
