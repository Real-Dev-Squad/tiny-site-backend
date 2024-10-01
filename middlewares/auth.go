package middleware

import (
	"net/http"

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

		ctx.Set("user", claims["email"])
		ctx.Set("userID", claims["userID"])
		ctx.Next()
	}
}



