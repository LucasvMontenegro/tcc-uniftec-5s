package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func NewRanking(
	instance *echo.Echo,
	// retrieveRankingsUseCase usecase.RetrieveRanking
) Ranking {

	return &ranking{
		instance: instance,
		// retrieveRankingUseCase: retrieveRankingsUseCase,
	}
}

type Ranking interface {
	Router
}

type ranking struct {
	instance *echo.Echo
	// retrieveRankingUseCase usecase.RetrieveRanking
}

func (c ranking) RegisterRoutes() {
	c.instance.GET("/edition/ranking", c.retrieveRanking())
}

func (tc ranking) retrieveRanking() func(c echo.Context) error {
	return func(c echo.Context) error {
		log.Info().Msg("GET /edition/ranking")

		// ranking, err := tc.createRankingUseCase.Execute(c.Request().Context(), createRankingReq.Name)
		// if err != nil {
		// 	log.Error().Interface("new ranking request", createRankingReq).Msg("/edition/rankings error")
		// 	return tc.handleErr(c, err)
		// }

		// response := response.NewCreatedRanking(*ranking)
		response := RankingResponse{
			EditionResp: EditionResp{
				Name:        "Edition Test",
				Description: "Description Edition Test",
				Status:      "ACTIVE",
				StartDate:   "2022-10-01",
				EndDate:     "2022-11-30",
			},
			PrizeResp: PrizeResp{
				Name:        "Prize Test",
				Description: "Description Prize Test",
			},
			Scores: []ScoresResp{
				ScoresResp{
					TeamName: "Brasil",
					Seiri:    "1",
					Seiton:   "10",
					Seiso:    "10",
					Seiketsu: "10",
					Shitsuke: "10",
					Total:    "41",
				},
				ScoresResp{
					TeamName: "Argentina",
					Seiri:    "5",
					Seiton:   "5",
					Seiso:    "5",
					Seiketsu: "5",
					Shitsuke: "5",
					Total:    "25",
				},
				ScoresResp{
					TeamName: "Alemanha",
					Seiri:    "7",
					Seiton:   "10",
					Seiso:    "10",
					Seiketsu: "10",
					Shitsuke: "10",
					Total:    "47",
				},
			},
		}

		log.Info().Msg("retrieving ranking success")
		return c.JSON(http.StatusOK, response)
	}
}

type RankingResponse struct {
	EditionResp
	PrizeResp
	Scores []ScoresResp
}

type EditionResp struct {
	Name        string `json:"edition_name"`
	Description string `json:"edition_description"`
	Status      string `json:"status"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

type PrizeResp struct {
	Name        string `json:"prize_name"`
	Description string `json:"prize_description"`
}

type ScoresResp struct {
	TeamName string `json:"team_name"`
	Seiri    string `json:"seiri"`
	Seiton   string `json:"seiton"`
	Seiso    string `json:"seiso"`
	Seiketsu string `json:"seiketsu"`
	Shitsuke string `json:"shitsuke"`
	Total    string `json:"total"`
}
