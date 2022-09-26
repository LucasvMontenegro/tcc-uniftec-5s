package controller

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func NewAccessValidator() AccessValidator {
	return &accessValidator{}
}

type AccessValidator interface {
	Restrict(c echo.Context) error
}

type accessValidator struct{}

func (accessValidator) Restrict(c echo.Context) error {
	user, ok := c.Get("user").(*jwt.Token)
	if !ok {
		log.Error().Msg("error parsing user token")
		return echo.ErrUnauthorized
	}

	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		log.Error().Msg("error parsing token claims")
		return echo.ErrUnauthorized
	}

	isAdmin, ok := claims["is_admin"].(bool)
	if !ok {
		log.Error().Msg("error parsing admin claim")
		return echo.ErrUnauthorized
	}

	if !isAdmin {
		log.Warn().Msg("invalid access attempt. user is not admin")
		return echo.ErrUnauthorized
	}

	return nil
}
