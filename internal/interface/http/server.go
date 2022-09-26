package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	custom_middleware "github.com/tcc-uniftec-5s/internal/interface/http/middleware"
)

type server struct {
	Instance   *echo.Echo
	Port       string
	Restricted *echo.Group
}

func New(
	port string,
	applicationName string,
	signingKey string,
) *server {
	instance := echo.New()
	skey := []byte(signingKey)

	instance.HTTPErrorHandler = custom_middleware.CustomHTTPErrorHandler
	restricted := instance.Group("", middleware.JWT(skey))
	return &server{
		Instance:   instance,
		Port:       port,
		Restricted: restricted,
	}
}

func (s server) Start() error {
	return s.Instance.Start(s.Port)
}
