package middleware

import (
	"fmt"
	"strings"

	"tiny-site-backend/initializers"
	"tiny-site-backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func DeserializeUser(c *fiber.Ctx) error {
	var tokenString string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("token") != "" {
		tokenString = c.Cookies("token")
	}

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
	}

	config, _ := initializers.LoadConfig(".")

	token, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		return []byte(config.JwtSecret), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("invalidate token: %v", err)})
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "invalid token claim"})
	}

	subClaim, ok := claims["sub"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "missing user ID in token"})
	}

	userID := subClaim
	fmt.Printf("User ID from token: %s\n", userID)

	var user models.User
	result := initializers.DB.First(&user, "id = ?", userID)
	if result.Error != nil {
		fmt.Printf("DB Query Error: %v\n", result.Error)
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "error", "message": "the user belonging to this token no longer exists"})
	}

	c.Locals("user", models.FilterUserRecord(&user))

	return c.Next()
}
