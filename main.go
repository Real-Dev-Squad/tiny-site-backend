package main

import (
	"flag"
	"log"
	"os"

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

	migrationsDir := "./migrations" 
	if err := utils.ApplyMigrations(db, migrationsDir); err != nil {
		log.Fatalf("failed to apply migrations: %v", err)
		os.Exit(1)
	}


	port := flag.String("port", os.Getenv("PORT"), "server address to listen on")
	flag.Parse()

	routes.Listen("0.0.0.0:"+*port, db)
}
