package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	authGroup := app.Group("/api/auth")
	userGroup := app.Group("/api/users")

	setupAuthRoutes(authGroup)
	setupUserRoutes(userGroup)
}
