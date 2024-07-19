package main

import (
	"fmt"
	"helloserver/config"
	"helloserver/log"
	"helloserver/server"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(fmt.Errorf("reading config: %w", err))
	}
	log.Init(cfg.LogLevel)

	app := &cli.App{
		Name:  "helloserver",
		Usage: "Operate the helloserver & related tasks",
		Commands: []*cli.Command{{
			Name:  "start",
			Usage: "Start the helloserver and listen for incoming requests",
			Action: func(c *cli.Context) error {
				return server.Start(cfg)
			}},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(fmt.Errorf("helloserver: %w", err))
	}
}
