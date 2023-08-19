package nnmg

import (
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/urfave/cli/v2"

	"github.com/netterminalmachine/nano-migrate/internal/core"
	"github.com/netterminalmachine/nano-migrate/internal/helpers"
)

func ApplyMigrations(ctx *cli.Context) error {
	if ctx == nil {
		return errors.New("Parent context is missing.")
	}

	config := helpers.LoadConfig()

	// need a pool for multi-statement queries
	pool, err := pgxpool.New(ctx.Context, config.PgUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	core.RunMigration(ctx.Context, config, pool)
	return nil
}
