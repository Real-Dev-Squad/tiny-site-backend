package routes

import (
	controller "github.com/Real-Dev-Squad/tiny-site-backend/controllers"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func AuthRoutes(reg *gin.RouterGroup, db *bun.DB) {
	auth := reg.Group("/auth")
	googleAuth := auth.Group("/google")

	googleAuth.GET("/login", controller.GoogleLogin)

	googleAuth.GET("/callback", func(ctx *gin.Context) {
		controller.GoogleCallback(ctx, db)
	})

	auth.GET("/logout", controller.Logout)
}
