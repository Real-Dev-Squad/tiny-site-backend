package middleware

import (
	"tiny-site-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateSignUpInput() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload models.SignUpInput

		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(400, gin.H{"status": "error", "message": err.Error()})
			c.Abort()
			return
		}

		if err := validate.Struct(payload); err != nil {
			c.JSON(400, gin.H{"status": "error", "message": err.Error()})
			c.Abort()
			return
		}

		c.Set("validatedPayload", &payload)

		c.Next()
	}
}

func ValidateSignInInput() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload models.SignInInput

		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(400, gin.H{"status": "error", "message": err.Error()})
			c.Abort()
			return
		}

		if err := validate.Struct(payload); err != nil {
			c.JSON(400, gin.H{"status": "error", "message": err.Error()})
			c.Abort()
			return
		}

		c.Set("validatedPayload", &payload)

		c.Next()
	}
}
