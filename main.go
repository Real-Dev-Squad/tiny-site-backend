package main

import (
	"flag"
	"os"

	"github.com/Real-Dev-Squad/tiny-site-backend/routes"
	"github.com/Real-Dev-Squad/tiny-site-backend/utils"
)

func main() {
	utils.LoadEnv(".env")
	dsn := os.Getenv("DB_URL")
	db := utils.SetupDBConnection(dsn)

	port := flag.String("port", ":8080", "server address to listen on")
	flag.Parse()

	routes.Listen("127.0.0.1"+*port, db)
}
