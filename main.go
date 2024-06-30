package main

import (
	"flag"

	"github.com/Real-Dev-Squad/tiny-site-backend/config"
	"github.com/Real-Dev-Squad/tiny-site-backend/logger"
	"github.com/Real-Dev-Squad/tiny-site-backend/routes"
	"github.com/Real-Dev-Squad/tiny-site-backend/utils"
)

func main() {
	dsn := config.DbUrl
	db, err := utils.SetupDBConnection(dsn)
	if err != nil {
		logger.Fatal("failed to connect to the database:", err)
	}


	port := config.Port
	if port == "" {
		port = "8080"
	}

	flag.Parse()

	routes.Listen("0.0.0.0:"+port, db)
}
