package utils

import (
	"database/sql"
	"os"
	"strconv"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func SetupDBConnection(dsn string) *bun.DB {
	maxOpenConnectionsStr := os.Getenv("DB_MAX_OPEN_CONNECTIONS")
	maxOpenConnections, err := strconv.Atoi(maxOpenConnectionsStr)

    if err != nil || maxOpenConnectionsStr == "" {
        maxOpenConnections = 10
    }

	pgDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	pgDB.SetMaxOpenConns(maxOpenConnections)

	db := bun.NewDB(pgDB, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	return db
}
