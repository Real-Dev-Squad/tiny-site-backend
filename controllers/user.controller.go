package controllers

import (
	"tiny-site-backend/models"

	"github.com/gin-gonic/gin"
)

func GetSelf(c *gin.Context, user models.UserResponse) {
	c.JSON(200, gin.H{"status": "success", "data": gin.H{"user": user}})
}
