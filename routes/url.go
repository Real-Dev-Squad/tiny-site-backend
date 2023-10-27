package routes

import (
	controller "github.com/Real-Dev-Squad/tiny-site-backend/controllers"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func TinyURLRoutes(rg *gin.RouterGroup, db *bun.DB) {
	tinyURL := rg.Group("/tinyurl")

	tinyURL.POST("", func(ctx *gin.Context) {
		controller.CreateTinyURL(ctx, db)
	})
}
