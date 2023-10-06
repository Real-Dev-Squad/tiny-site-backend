package controllers

import (
	"tiny-site-backend/models"

	"github.com/gofiber/fiber/v2"
)

func GetSelf(c *fiber.Ctx) error {
	// Retrieve the user data from the context. This data is typically set in previous middleware user authentication.
	user := c.Locals("user").(models.UserResponse)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": user}})
}
