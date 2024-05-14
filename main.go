package main

import (
	"flag"
	"log"
	"os"

	migration "github.com/Real-Dev-Squad/tiny-site-backend/cmd/bun"
	"github.com/Real-Dev-Squad/tiny-site-backend/routes"
	"github.com/Real-Dev-Squad/tiny-site-backend/utils"
)

func main() {
    utils.LoadEnv(".env")
    dsn := os.Getenv("DB_URL")
    db, err := utils.SetupDBConnection(dsn)
    if err != nil {
        log.Fatalf("failed to connect to the database: %v", err)
        os.Exit(1)
    }

    migration.RunMigrations()

    port := flag.String("port", os.Getenv("PORT"), "server address to listen on")
    flag.Parse()

    routes.Listen("0.0.0.0:"+*port, db)
}
