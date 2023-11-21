package utils

import "github.com/gin-gonic/gin"

func HandleError(ctx *gin.Context, status int, message string, err error) {

	ctx.JSON(status, gin.H{
		"message": message,
		"error":   err.Error(),
	})
}
