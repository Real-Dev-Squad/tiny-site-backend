package routes

import (
	"github.com/Real-Dev-Squad/tiny-site-backend/controllers"
	"github.com/Real-Dev-Squad/tiny-site-backend/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func TinyURLRoutes(rg *gin.RouterGroup, db *bun.DB) {
	shorten := rg.Group("/shorten")
	redirect := rg.Group("/redirect")
	urls := rg.Group("/urls")

	shorten.Use(middleware.AuthMiddleware())

	shorten.POST("", func(ctx *gin.Context) {
		controller.CreateTinyURL(ctx, db)
	})

	redirect.GET("/:shortURL", func(ctx *gin.Context) {
		controller.RedirectShortURL(ctx, db)
	})

	urls.GET("/:shortURL", func(ctx *gin.Context) {
		controller.GetURLDetails(ctx, db)
	})
}
