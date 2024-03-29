package server

import (
	"absurdlab.io/tigerd/cmd/server/internal"
	"absurdlab.io/tigerd/internal/buildinfo"
	"github.com/hellofresh/health-go/v5"
	"github.com/labstack/echo-contrib/prometheus"
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
	jwksHandler *jwksHandler,
	health *health.Health,
) {
	e.GET("/.well-known/openid-configuration", wellKnownHandler.GetConfiguration)
	e.GET("/jwks.json", jwksHandler.GetJSONWebKeySet)
	e.GET("/health", echo.WrapHandler(health.Handler()))
}

func setupMetrics(e *echo.Echo, metrics *internal.Metrics) {
	e.Use(metricsMiddleware(metrics))

	prom := prometheus.NewPrometheus(
		"tigerd",
		func(c echo.Context) bool {
			switch c.Path() {
			case "/health",
				"/metrics":
				return true
			default:
				return false
			}
		},
		metrics.ToPrometheusMetrics(),
	)

	// prometheus metrics endpoint /metrics is mounted implicitly on echo.
	prom.Use(e)
}
