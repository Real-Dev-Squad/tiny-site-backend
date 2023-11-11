package routes

import (
	controller "github.com/Real-Dev-Squad/tiny-site-backend/controllers"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func TinyURLRoutes(rg *gin.RouterGroup, db *bun.DB) {
	tinyURL := rg.Group("/tinyurl")
	urls := rg.Group("/urls")

	tinyURL.POST("", func(ctx *gin.Context) {
		controller.CreateTinyURL(ctx, db)
	})

	tinyURL.GET("/:shortURL", func(ctx *gin.Context) {
		controller.RedirectShortURL(ctx, db)
	})

	urls.GET("/:shortURL", func(ctx *gin.Context) {
		controller.GetURLDetails(ctx, db)
	})
}
