package main

import (
	"flag"
	"log"
	"os"

	"github.com/Real-Dev-Squad/tiny-site-backend/config"
	"github.com/Real-Dev-Squad/tiny-site-backend/routes"
	"github.com/Real-Dev-Squad/tiny-site-backend/utils"
)

func main() {
	dsn := config.DbUrl
	db, err := utils.SetupDBConnection(dsn)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}


	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	flag.Parse()

	routes.Listen("0.0.0.0:"+port, db)
}
