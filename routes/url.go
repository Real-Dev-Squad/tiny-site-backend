package routes

import (
	controller "github.com/Real-Dev-Squad/tiny-site-backend/controllers"
	"github.com/Real-Dev-Squad/tiny-site-backend/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func TinyURLRoutes(rg *gin.RouterGroup, db *bun.DB) {
	redirect := rg.Group("/redirect")
	tinyURL := rg.Group("/tinyurl")
	urls := rg.Group("/urls")

	tinyURL.Use(middleware.AuthMiddleware())

	tinyURL.POST("", func(ctx *gin.Context) {
		controller.CreateTinyURL(ctx, db)
	})

	urls.Group("/self", middleware.AuthMiddleware()).GET("", func(ctx *gin.Context) {
		controller.GetAllURLs(ctx, db)
	})

	urls.GET("/:shortURL", func(ctx *gin.Context) {
		controller.GetURLDetails(ctx, db)
	})

	redirect.GET("/:shortURL", func(ctx *gin.Context) {
		controller.RedirectShortURL(ctx, db)
	})
}
