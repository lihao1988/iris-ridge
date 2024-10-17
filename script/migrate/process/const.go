package process

import (
	"os"
	"strconv"
	"text/template"
)

// evn data
var (
	// GOOSENOCOLOR GOOSE NO COLOR
	GOOSENOCOLOR = envOr("NO_COLOR", "false")
)

// An DirVar is an directory variable Name=Value.
type DirVar struct {
	Name  string
	Value string
}

func checkNoColorFromEnv() bool {
	ok, _ := strconv.ParseBool(GOOSENOCOLOR)
	return ok
}

// envOr returns os.Getenv(key) if set, or else default.
func envOr(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		val = def
	}
	return val
}

var (
	usagePrefix = `Usage: ./goose [OPTIONS] COMMAND

Examples:
    ./goose status
    ./goose create init sql
    ./goose create add_some_column sql
    ./goose create fetch_user_data go
    ./goose up

Options:
`

	usageCommands = `
Commands:
    up                   Migration the DB to the most recent version available
    up-by-one            Migration the DB up by 1
    up-to VERSION        Migration the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    reset                Roll back all migrations
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with the current timestamp
    fix                  Apply sequential ordering to migrations
    validate             Check migration files without running them
`
)

var sqlMigrationTemplate = template.Must(template.New("goose.sql-migration").Parse(`-- Thank you for giving goose a try!
-- 
-- This file was automatically created running goose init. If you're familiar with goose
-- feel free to remove/rename this file, write some SQL and goose up. Briefly,
-- 
-- Documentation can be found here: https://pressly.github.io/goose
--
-- A single goose .sql file holds both Up and Down migrations.
-- 
-- All goose .sql files are expected to have a -- +goose Up annotation.
-- The -- +goose Down annotation is optional, but recommended, and must come after the Up annotation.
-- 
-- The -- +goose NO TRANSACTION annotation may be added to the top of the file to run statements 
-- outside a transaction. Both Up and Down migrations within this file will be run without a transaction.
-- 
-- More complex statements that have semicolons within them must be annotated with 
-- the -- +goose StatementBegin and -- +goose StatementEnd annotations to be properly recognized.
-- 
-- Use GitHub issues for reporting bugs and requesting features, enjoy!

-- +goose Up
SELECT 'up SQL query';

-- +goose Down
SELECT 'down SQL query';
`))
