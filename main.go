package main

import (
	"fmt"
	"log"

	"tiny-site-backend/initializers"
	"tiny-site-backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load environment variables: %v\n", err)
	}

	err = initializers.ConnectDB(&config)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v\n", err)
	}

	r := gin.Default()

	r.Use(logger.SetLogger())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	corsConfig.AllowMethods = []string{"GET", "POST"}
	corsConfig.AllowCredentials = true
	r.Use(cors.New(corsConfig))

	routes.SetupRoutes(r)

	port := ":8000"
	fmt.Printf("Server listening on port %s\n", port)
	if err := r.Run(port); err != nil {
		log.Fatalf("Error while starting the server: %v\n", err)
	}
}
