package nnmg

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/urfave/cli/v2"

	"github.com/netterminalmachine/nano-migrate/internal/core"
	"github.com/netterminalmachine/nano-migrate/internal/helpers"
)

func CreateNewMigration(ctx *cli.Context) error {
	if ctx == nil {
		return errors.New("Parent context is missing.")
	}

	config := helpers.LoadConfig()

	if ctx.Args().Len() < 3 {
		return errors.New("You must provide a name for your migration, which must be at least three words. E.g. 'create foo table'.")
	}

	nameArgs := strings.Join(ctx.Args().Slice(), "-")
	fmt.Printf("Proposed naming: [%s]\n", nameArgs)

	// need a pool for multi-statement queries
	pool, err := pgxpool.New(ctx.Context, config.PgUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	core.CreateMigration(ctx.Context, config, pool, nameArgs, "core")
	return nil
}
