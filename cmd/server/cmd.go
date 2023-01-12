package server

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
	"go.uber.org/fx"
	"regexp"
)

func Command() *cli.Command {
	var (
		cfg   = new(Config)
		flags = []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Usage:   "Path to a yaml configuration file.",
				EnvVars: []string{"TIGERD_CONFIG"},
			},
			altsrc.NewIntFlag(cfg.flagPort()),
			altsrc.NewStringFlag(cfg.flagExternalURL()),
			altsrc.NewStringFlag(cfg.flagLogLevel()),
			altsrc.NewBoolFlag(cfg.flagLogJSON()),
		}
	)

	return &cli.Command{
		Name:        "server",
		Description: "Serves the OpenID Connect server APIs",
		Before:      altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("config")),
		Flags:       flags,
		Action: func(cc *cli.Context) error {
			if err := cfg.validate(); err != nil {
				return fmt.Errorf("config: %s", err)
			}

			return fx.New(
				fx.NopLogger,
				fx.Supply(cfg),
				fx.Provide(
					newEcho,
					newBaseLogger,
					newHealth,
					newOpenIDConnectProvider,
					newWellKnownHandler,
				),
				fx.Invoke(
					mountEndpoints,
					start,
				),
			).Start(cc.Context)
		},
	}
}

type Config struct {
	Port        int    `yaml:"port"`
	ExternalURL string `yaml:"external_url"`
	LogLevel    string `yaml:"log_level"`
	LogJSON     bool   `yaml:"log_json"`
}

func (c *Config) address() string {
	return fmt.Sprintf(":%d", c.Port)
}

func (c *Config) validate() (err error) {
	err = validation.Errors{
		"port": validation.Validate(c.Port,
			validation.Required,
			validation.Min(1024),
			validation.Max(65535),
		),
		"external_url": validation.Validate(c.ExternalURL,
			validation.Required,
			is.URL,
			validation.Match(regexp.MustCompile("^https?://.*")),
		),
		"log_level": validation.Validate(c.LogLevel,
			validation.In("INFO", "DEBUG", "ERROR", "TRACE", "WARN"),
		),
	}.Filter()
	return
}

func (c *Config) flagPort() *cli.IntFlag {
	return &cli.IntFlag{
		Name:        "port",
		Usage:       "Port where server listens for requests.",
		Value:       4904,
		Destination: &c.Port,
		EnvVars:     []string{"TIGERD_PORT"},
	}
}

func (c *Config) flagExternalURL() *cli.StringFlag {
	return &cli.StringFlag{
		Name:        "external_url",
		Usage:       "Externally accessible URL for the server deployment.",
		Required:    true,
		Destination: &c.ExternalURL,
		EnvVars:     []string{"TIGERD_EXTERNAL_URL"},
	}
}

func (c *Config) flagLogLevel() *cli.StringFlag {
	return &cli.StringFlag{
		Name:        "log_level",
		Usage:       "Logging level. [INFO|ERROR|WARN|DEBUG|TRACE]",
		Value:       "INFO",
		Destination: &c.LogLevel,
		EnvVars:     []string{"TIGERD_LOG_LEVEL"},
	}
}

func (c *Config) flagLogJSON() *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:        "log_json",
		Usage:       "Enable logging with JSON format",
		Destination: &c.LogJSON,
		EnvVars:     []string{"TIGERD_LOG_JSON"},
	}
}
