package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	usecase "github.com/tcc-uniftec-5s/internal/app/use_case"
	"github.com/tcc-uniftec-5s/internal/interface/http/dto/response"
)

func NewRanking(
	instance *echo.Echo,
	retrieveRankingUseCase usecase.RetrieveRanking,
) Ranking {

	return &ranking{
		instance:               instance,
		retrieveRankingUseCase: retrieveRankingUseCase,
	}
}

type Ranking interface {
	Router
}

type ranking struct {
	instance               *echo.Echo
	retrieveRankingUseCase usecase.RetrieveRanking
}

func (c ranking) RegisterRoutes() {
	c.instance.GET("/edition/ranking", c.retrieveRanking())
}

func (tc ranking) retrieveRanking() func(c echo.Context) error {
	return func(c echo.Context) error {
		log.Info().Msg("GET /edition/ranking")

		ranking, err := tc.retrieveRankingUseCase.Execute(c.Request().Context())
		if err != nil {
			return err // todo handle err
		}

		res := response.BuildRanking(*ranking)

		log.Info().Msg("retrieving ranking success")
		return c.JSON(http.StatusOK, res)
	}
}
