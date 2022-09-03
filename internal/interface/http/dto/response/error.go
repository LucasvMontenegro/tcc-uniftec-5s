package response

import (
	"github.com/labstack/echo/v4"
)

type InvalidParam struct {
	Name   string `json:"name"`
	Reason string `json:"reason"`
}

type Error struct {
	Title         string         `json:"title"`
	Detail        string         `json:"detail"`
	Status        int            `json:"status"`
	Type          string         `json:"type"`
	InvalidParams []InvalidParam `json:"invalid_params"`
}

func (res Error) Write(c echo.Context) error {
	return c.JSON(res.Status, res)
}
