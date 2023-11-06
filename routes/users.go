package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/uptrace/bun"
    "github.com/Real-Dev-Squad/tiny-site-backend/controllers"
    "github.com/Real-Dev-Squad/tiny-site-backend/middlewares"
)

func UserRoutes(rg *gin.RouterGroup, db *bun.DB) {
    users := rg.Group("/users")
    users.Use(middleware.AuthMiddleware())

    users.GET("", func(ctx *gin.Context) {
        controllers.GetUserList(ctx, db)
    })

    users.GET("/:id", func(ctx *gin.Context) {
        controllers.GetUserByID(ctx, db)
    })

    users.GET("/self", func(ctx *gin.Context) {
        controllers.GetSelfUser(ctx, db)
    })
}
