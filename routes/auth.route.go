package routes

import (
	"tiny-site-backend/controllers"
	"tiny-site-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func setupAuthRoutes(router fiber.Router) {
	router.Post("/register", middleware.ValidateSignUpInput, controllers.SignUpUser)
	router.Post("/login", middleware.ValidateSignInInput, controllers.SignInUser)
	router.Get("/logout", middleware.DeserializeUser, controllers.LogoutUser)
}
