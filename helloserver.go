package main

import (
	"fmt"
	"helloserver/config"
	"helloserver/log"
	"helloserver/server"
	"helloserver/sigctx"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	ctx := sigctx.WithSignal(os.Interrupt)
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
				return server.Start(c.Context, cfg)
			}},
		},
	}

	if err := app.RunContext(ctx, os.Args); err != nil {
		panic(fmt.Errorf("helloserver: %w", err))
	}
}
