package routes

import (
	"tiny-site-backend/controllers"
	"tiny-site-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func setupUserRoutes(router fiber.Router) {
	router.Get("/self", middleware.DeserializeUser, controllers.GetMe)
}
