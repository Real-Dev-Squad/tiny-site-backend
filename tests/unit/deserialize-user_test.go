package controllers

import (
	"fmt"
	"tiny-site-backend/initializers"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var (
	MockUser = "5b8e88c2-955a-4e33-b228-632de1f38f45"
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

		_, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"status": "fail", "message": "invalid token claim"})
			return
		}
		c.Set("user", MockUser)
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
