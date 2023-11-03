package routes

import (
	"net/http"

	"github.com/Real-Dev-Squad/tiny-site-backend/middlewares"
	"github.com/Real-Dev-Squad/tiny-site-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func UserRoutes(rg *gin.RouterGroup, db *bun.DB) {
	users := rg.Group("/users")
	users.Use(middleware.AuthMiddleware())

	users.GET("", func(ctx *gin.Context) {

		var users []models.User
		err := db.NewSelect().Model(&users).OrderExpr("id ASC").Limit(10).Scan(ctx)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "users fetched successfully",
			"data":    users,
		})
	})

	users.GET("/:id",func(ctx *gin.Context) {
		id := ctx.Param("id")

		var user models.User
		err := db.NewSelect().Model(&user).Where("id = ?", id).Scan(ctx)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "user fetched successfully",
			"data":    user,
		})
	})

	users.GET("/self", func(ctx *gin.Context) {
		userEmail, _ := ctx.Get("user")

		var user models.User
		err := db.NewSelect().Model(&user).Where("email = ?", userEmail).Scan(ctx)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "user fetched successfully",
			"data":    user,
		})
	})
}
