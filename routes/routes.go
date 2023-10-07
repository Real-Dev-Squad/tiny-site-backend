package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	authGroup := r.Group("/api/auth")
	userGroup := r.Group("/api/users")

	setupAuthRoutes(authGroup)
	setupUserRoutes(userGroup)
}
