package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"schneider.vip/problem"
)

func NewReport(
	instance *echo.Echo,
	// retrieveReportsUseCase usecase.RetrieveReport
) Report {

	return &report{
		instance: instance,
		// retrieveReportUseCase: retrieveReportsUseCase,
	}
}

type Report interface {
	Router
}

type report struct {
	instance *echo.Echo
	// retrieveReportUseCase usecase.RetrieveReport
}

func (c report) RegisterRoutes() {
	c.instance.GET("/editions/:editionId/teams/:teamId/report", c.retrieveReport())
}

func (tc report) retrieveReport() func(c echo.Context) error {
	return func(c echo.Context) error {
		log.Info().Msg("GET /editions/:editionId/teams/:teamId/report")
		teamID, _ := strconv.ParseInt(c.Param("teamId"), 10, 64)
		editionID, _ := strconv.ParseInt(c.Param("editionId"), 10, 64)

		if teamID == 0 || editionID == 0 {
			log.Info().Msg("import request error. invalid params")
			status := http.StatusBadRequest
			prob := problem.New(
				problem.Title("INVALID_PARAMS_ERROR"),
				problem.Detail("Bad Request"),
				problem.Status(status),
			)

			return c.JSON(status, prob)
		}

		log.Info().Msgf("GET /editions/%v/teams/%v/report", teamID, editionID)

		// report, err := tc.createReportUseCase.Execute(c.Request().Context(), createReportReq.Name)
		// if err != nil {
		// 	log.Error().Interface("new report request", createReportReq).Msg("/edition/reports error")
		// 	return tc.handleErr(c, err)
		// }

		// response := response.NewCreatedReport(*report)
		response := ReportResponse{
			Name: "Edition Test",
			// Scores: ScoresResp{
			// 	TeamName: "Brasil",
			// 	Seiri:    "1",
			// 	Seiton:   "10",
			// 	Seiso:    "10",
			// 	Seiketsu: "10",
			// 	Shitsuke: "10",
			// 	Total:    "41",
			// },
		}

		log.Info().Msg("retrieving report success")
		return c.JSON(http.StatusOK, response)
	}
}

type ReportResponse struct {
	Name string `json:"edition_name"`
	// Scores ScoresResp `json:"scores"`
}
