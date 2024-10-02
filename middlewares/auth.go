package middleware

import (
	"net/http"
	"strconv"

	"github.com/Real-Dev-Squad/tiny-site-backend/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenCookie, err := ctx.Request.Cookie("token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			ctx.Abort()
			return
		}

		token := tokenCookie.Value

		claims, err := utils.VerifyToken(token)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			ctx.Abort()
			return
		}

		userIDStr := claims["userID"].(string)
		userID, err := strconv.ParseInt(userIDStr, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid UserID"})
			ctx.Abort()
			return
		}

		ctx.Set("user", claims["email"])
		ctx.Set("userID", userID)
		ctx.Next()
	}
}