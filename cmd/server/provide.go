package server

import (
	"absurdlab.io/tigerd/internal/buildinfo"
	"absurdlab.io/tigerd/oidc"
	"encoding/json"
	"github.com/hellofresh/health-go/v5"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/ziflex/lecho/v3"
	"os"
	"strings"
	"time"
)

func newEcho(logger *zerolog.Logger) *echo.Echo {
	e := echo.New()
	e.HidePort = true
	e.HideBanner = true
	e.Logger = lecho.New(logger)
	//e.HTTPErrorHandler
	return e
}

func newOpenIDConnectProvider(cfg *Config) *oidc.Provider {
	return oidc.NewProvider(cfg.ExternalURL)
}

func newWellKnownHandler(p *oidc.Provider) (*wellKnownHandler, error) {
	providerBytes, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return &wellKnownHandler{providerBytes: providerBytes}, nil
}

func newBaseLogger(cfg *Config) (*zerolog.Logger, error) {
	level, err := zerolog.ParseLevel(strings.ToLower(cfg.LogLevel))
	if err != nil {
		return nil, err
	}
	logger := zerolog.New(os.Stdout).Level(level)

	if cfg.LogJSON {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	} else {
		logger = logger.Output(&zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		})
	}

	logger = logger.With().Timestamp().Logger()

	return &logger, nil
}

func newHealth() (*health.Health, error) {
	return health.New(
		health.WithComponent(health.Component{
			Name:    "tigerd",
			Version: buildinfo.Version,
		}),
		health.WithSystemInfo(),
	)
}
