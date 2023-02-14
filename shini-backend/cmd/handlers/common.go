package handlers

import (
	"shini/cmd/models"

	"github.com/gofiber/fiber/v2"
)

func GetStats(c *fiber.Ctx) error {
	stats, err := models.GetStats()
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		c.JSON(err.Error())
	}
	return c.JSON(stats)
}
