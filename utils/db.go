package utils

import (
	"database/sql"
	"fmt"

	"github.com/Real-Dev-Squad/tiny-site-backend/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func SetupDBConnection(dsn string) (*bun.DB, error) {
	maxOpenConnections := config.DbMaxOpenConnections
	fmt.Println("max open connections: ", maxOpenConnections)
	pgDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	pgDB.SetMaxOpenConns(maxOpenConnections)

	db := bun.NewDB(pgDB, pgdialect.New())

	dbConnectionError := db.Ping()
	if dbConnectionError != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", dbConnectionError)
	}

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	return db, nil
}
