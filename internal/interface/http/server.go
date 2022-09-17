package server

import (
	"github.com/labstack/echo/v4"
	custom_middleware "github.com/tcc-uniftec-5s/internal/interface/http/middleware"
)

type server struct {
	Instance *echo.Echo
	Port     string
}

func New(
	port string,
	applicationName string,
) *server {
	instance := echo.New()

	instance.HTTPErrorHandler = custom_middleware.CustomHTTPErrorHandler
	return &server{
		Instance: instance,
		Port:     port,
	}
}

func (s server) Start() error {
	return s.Instance.Start(s.Port)
}
