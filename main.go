package main

import (
	"fmt"
	"log"
	"os"

	"tiny-site-backend/initializers"
	"tiny-site-backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Printf("Failed to load environment variables: %v\n", err)
		return
	}

	initializers.ConnectDB(&config)

	Origin := os.Getenv("DOMAIN")
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     Origin,
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST",
		AllowCredentials: true,
	}))
	routes.SetupRoutes(app)

	port := ":8000"
	fmt.Printf("Server listening on port %s\n", port)
	if err := app.Listen(port); err != nil {
		log.Printf("Error while starting the server: %v\n", err)
	}
}
