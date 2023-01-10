package server

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"golang.org/x/net/http2"
)

func start(cfg *Config, e *echo.Echo, logger *zerolog.Logger) error {
	logger.Info().Int("port", cfg.Port).Msg("Listening for requests.")
	return e.StartH2CServer(cfg.address(), &http2.Server{})
}

func mountEndpoints(
	e *echo.Echo,
	wellKnownHandler *wellKnownHandler,
) {
	e.GET("/.well-known/openid-configuration", wellKnownHandler.GetConfiguration)
}
