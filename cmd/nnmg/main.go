package nnmg

import (
	"os"

	"log"

	"github.com/urfave/cli/v2"
)

// This is an entrypoint / wrapper for testability. ENSURE TO ALWAYS RETURN A VALUE FOR os.Exit() !!
func Main() int {
	app := &cli.App{
		Name:  "nnmg",
		Usage: "Nano-migrate (nnmg) is a tiny, env-aware, forward-only, postgres migrations manager.",
		Commands: []*cli.Command{
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "Initialise the migrations table",
				Action:  InitialiseMigrations,
			},
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "Create a new migration",
				Action:  CreateNewMigration,
			},
			{
				Name:    "migrate",
				Aliases: []string{"m"},
				Usage:   "Apply all pending (un-applied) migrations",
				Action:  ApplyMigrations,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Println(err)
		return 1
	}

	return 0
}
