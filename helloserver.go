package main

import (
	"helloserver/backend"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	log.Println("loading dotenv...")
	backend.LoadDotenv()

	app := &cli.App{
		Name:  "helloserver",
		Usage: "Operate the helloserver & related tasks",
		Commands: []*cli.Command{{
			Name:  "start",
			Usage: "Start the helloserver and listen for incoming requests",
			Action: func(c *cli.Context) error {
				b := backend.NewBackend()
				return b.Start()
			}},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
