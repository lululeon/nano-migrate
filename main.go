package main

import (
	"log"
	"os"

	"nano-migrate/cmd"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "nnmg",
		Usage: "Nano-migrate (nnmg) is a tiny, env-aware, forward-only, postgres migrations manager.",
		Commands: []*cli.Command{
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "Initialise the migrations table",
				Action:  cmd.InitialiseMigrations,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
