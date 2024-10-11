package routes

import (
	controller "github.com/Real-Dev-Squad/tiny-site-backend/controllers"
	middleware "github.com/Real-Dev-Squad/tiny-site-backend/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func UserRoutes(rg *gin.RouterGroup, db *bun.DB) {

    users := rg.Group("/users")
    user := rg.Group("/user")
    users.Use(middleware.AuthMiddleware())
    user.Use(middleware.AuthMiddleware())

    users.GET("/:id", func(ctx *gin.Context) {
        controller.GetUserByID(ctx, db)
    })

    users.GET("/self", func(ctx *gin.Context) {
        controller.GetSelfUser(ctx, db)
    })
}
