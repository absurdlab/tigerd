package main

import (
	"absurdlab.io/tigerd/cmd/server"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

func main() {
	tigerd := cli.App{
		Name:     "tigerd",
		Version:  "",
		Compiled: time.Time{},
		Commands: cli.Commands{
			server.Command(),
		},
		Authors: []*cli.Author{
			{Name: "Weinan Qiu", Email: "davidiamyou@gmail.com"},
		},
		Copyright: "MIT",
	}

	if err := tigerd.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("failed to run tigerd")
	}
}
