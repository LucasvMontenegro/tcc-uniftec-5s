package api

import (
	"github.com/labstack/echo/v4"

	"github.com/tcc-uniftec-5s/internal/domain/sample"
	"github.com/tcc-uniftec-5s/internal/interface/http/handler"
)

type serviceImpl struct {
	api  *echo.Echo
	Port string
}

func NewService(
	port string,
	applicationName string,
	sampleService sample.Service,
) *serviceImpl {
	echoAPI := echo.New()

	handler.MakeHealtHandler(echoAPI)
	handler.MakeSampleHandler(echoAPI, sampleService)

	return &serviceImpl{
		api:  echoAPI,
		Port: port,
	}
}

func (s serviceImpl) StartServer() error {
	return s.api.Start(s.Port)
}
