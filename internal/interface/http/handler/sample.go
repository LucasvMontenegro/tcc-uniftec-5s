package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/tcc-uniftec-5s/internal/domain/sample"
	"github.com/tcc-uniftec-5s/internal/infra/constants"
	"github.com/tcc-uniftec-5s/internal/interface/http/dto/request"
	"github.com/tcc-uniftec-5s/internal/interface/http/dto/response"
	"gorm.io/gorm"
	"schneider.vip/problem"
)

func MakeSampleHandler(api *echo.Echo, sampleService sample.Service) {
	api.GET("/v1/samples/:referenceUUID", GetSampleByReferenceUUID(sampleService))
	api.POST("/v1/samples", CreateSample(sampleService))
}

func CreateSample(sampleService sample.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		log.Info().Msg("creating Sample")

		var createSampleRequest request.CreateSample
		var createdSampleResponse response.Sample

		if err := c.Bind(&createSampleRequest); err != nil {
			log.Info().Interface("createSampleRequest", createSampleRequest).Msg("deserialization error")
			status := http.StatusBadRequest
			prob := problem.New(
				problem.Title("REQUEST_DESERIALIZATION_ERROR"),
				problem.Detail("Bad Request"),
				problem.Status(status),
			)

			return c.JSON(status, prob)
		}

		createDTO := sample.CreateDTO{
			ReferenceUUID: createSampleRequest.ReferenceUUID,
		}

		createdSample, err := sampleService.CreateSample(c.Request().Context(), createDTO)
		if err != nil {
			log.Error().Interface("createSampleRequest", createSampleRequest).Msg("error on Sample creation")
			status := http.StatusInternalServerError
			prob := problem.New(
				problem.Title("UNEXPECTED_ERROR"),
				problem.Detail("Internal Server Error"),
				problem.Status(status),
			)
			return c.JSON(status, prob)
		}

		log.Info().Msg("retrieving created sample")

		createdSampleResponse.FromSample(createdSample)
		c.Response().Header().Set("Location", fmt.Sprint(constants.SampleRessource, "/", createdSample.ReferenceUUID))
		return c.JSON(http.StatusCreated, createdSampleResponse)
	}
}

func GetSampleByReferenceUUID(sampleService sample.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		sampleResponse := response.Sample{}

		byReferenceUUID := c.Param("referenceUUID")

		fetchSample, err := sampleService.GetByReferenceUUID(c.Request().Context(), byReferenceUUID)

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				log.Warn().Msgf("GetSampleByReferenceUUID not found by %s", byReferenceUUID)

				errResponse := problem.New(
					problem.Title(sample.ErrSampleNotFound.Error()),
					problem.Detail("Sample not found or invalid"),
					problem.Status(http.StatusNotFound),
				)

				return c.JSON(http.StatusNotFound, errResponse)
			}

			log.Error().Msgf("GetSampleByReferenceUUID error when trying by %s", byReferenceUUID)

			return c.JSON(http.StatusInternalServerError, constants.ProblemInternalServerError)
		}

		sampleResponse.FromSample(fetchSample)

		return c.JSON(http.StatusOK, sampleResponse)
	}
}
