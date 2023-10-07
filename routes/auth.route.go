package routes

import (
	"tiny-site-backend/controllers"
	"tiny-site-backend/middleware"

	"github.com/gin-gonic/gin"
)

func setupAuthRoutes(router *gin.RouterGroup) {
	router.POST("/register", middleware.ValidateSignUpInput(), controllers.SignUpUser)
	router.POST("/login", middleware.ValidateSignInInput(), controllers.SignInUser)
	router.GET("/logout", middleware.DeserializeUser(), controllers.LogoutUser)
}
