package middleware

import (
	"fmt"

	"tiny-site-backend/initializers"
	"tiny-site-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func DeserializeUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := extractToken(c)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		config, _ := initializers.LoadConfig(".")
		secret := []byte(config.JwtSecret)

		token, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
			if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
			}
			return secret, nil
		})
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"status": "fail", "message": fmt.Sprintf("invalid token: %v", err)})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"status": "fail", "message": "invalid token claim"})
			return
		}

		userID, ok := claims["sub"].(string)
		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"status": "fail", "message": "missing user ID in token"})
			return
		}

		var user models.User
		result := initializers.DB.First(&user, "id = ?", userID)
		if result.Error != nil {
			c.AbortWithStatusJSON(403, gin.H{"status": "error", "message": "the user belonging to this token no longer exists"})
			return
		}

		c.Set("user", models.FilterUserRecord(&user))
		c.Next()
	}
}

func extractToken(c *gin.Context) (string, error) {
	authorization := c.GetHeader("Authorization")
	if authorization != "" && len(authorization) > 7 && authorization[:7] == "Bearer " {
		return authorization[7:], nil
	}
	token, err := c.Cookie("token")
	if err == nil {
		return token, nil
	}
	return "", fmt.Errorf("you are not logged in")
}
