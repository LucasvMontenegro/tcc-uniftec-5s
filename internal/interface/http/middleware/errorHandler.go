package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"schneider.vip/problem"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	c.Logger().Warn(err)

	var problemJSON *problem.Problem
	status := http.StatusInternalServerError
	message := "internal server error"

	if he, ok := err.(*echo.HTTPError); ok {
		status = he.Code
		message = he.Message.(string)
	}

	problemJSON = problem.New(
		problem.Title(message),
		problem.Status(status),
	)

	c.JSON(status, problemJSON)
}
