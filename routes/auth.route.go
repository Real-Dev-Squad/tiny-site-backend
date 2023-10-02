package routes

import (
	"tiny-site-backend/controllers"
	"tiny-site-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func setupAuthRoutes(router fiber.Router) {
	router.Post("/register", controllers.SignUpUser)
	router.Post("/login", controllers.SignInUser)
	router.Get("/logout", middleware.DeserializeUser, controllers.LogoutUser)
}
