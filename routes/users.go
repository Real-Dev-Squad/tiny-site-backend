package routes

import (
	"github.com/Real-Dev-Squad/tiny-site-backend/controllers"
	"github.com/Real-Dev-Squad/tiny-site-backend/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func UserRoutes(rg *gin.RouterGroup, db *bun.DB) {

	users := rg.Group("/users")
	user := rg.Group("/user")

	users.Use(middleware.AuthMiddleware())
	user.Use(middleware.AuthMiddleware())

	users.GET("", func(ctx *gin.Context) {
		controller.GetUserList(ctx, db)
	})

	users.GET("/:id", func(ctx *gin.Context) {
		controller.GetUserByID(ctx, db)
	})

	user.GET("", func(ctx *gin.Context) {
		controller.GetSelfUser(ctx, db)
	})

	user.GET("/:id/urls", func(ctx *gin.Context) {
		controller.GetAllURLs(ctx, db)
	})
}
