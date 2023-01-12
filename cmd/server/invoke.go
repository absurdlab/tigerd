package server

import (
	"absurdlab.io/tigerd/internal/buildinfo"
	"github.com/hellofresh/health-go/v5"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"golang.org/x/net/http2"
)

func start(cfg *Config, e *echo.Echo, logger *zerolog.Logger) error {
	logger.Info().
		Int("port", cfg.Port).
		Str("version", buildinfo.Version).
		Msg("Listening for requests.")

	return e.StartH2CServer(cfg.address(), &http2.Server{})
}

func mountEndpoints(
	e *echo.Echo,
	wellKnownHandler *wellKnownHandler,
	health *health.Health,
) {
	e.GET("/.well-known/openid-configuration", wellKnownHandler.GetConfiguration)
	e.GET("/health", echo.WrapHandler(health.Handler()))
}
