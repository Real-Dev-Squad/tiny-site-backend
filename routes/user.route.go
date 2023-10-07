package routes

import (
	"tiny-site-backend/controllers"
	"tiny-site-backend/middleware"
	"tiny-site-backend/models"

	"github.com/gin-gonic/gin"
)

func setupUserRoutes(userGroup *gin.RouterGroup) {
	userGroup.GET("/self", middleware.DeserializeUser(), func(c *gin.Context) {
		user := c.MustGet("user").(models.UserResponse)
		controllers.GetSelf(c, user)
	})
}
