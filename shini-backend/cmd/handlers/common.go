package handlers

import "github.com/gofiber/fiber/v2"

func GetStats(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "This this is getstats",
	})
}
