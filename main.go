package main

import (
	"absurdlab.io/tigerd/cmd/server"
	"absurdlab.io/tigerd/internal/buildinfo"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	tigerd := &cli.App{
		Name:     "tigerd",
		Version:  buildinfo.Version,
		Compiled: buildinfo.CompileAtTime(),
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
