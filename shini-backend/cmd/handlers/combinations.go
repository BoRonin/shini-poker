package handlers

import (
	"shini/cmd/models"

	"github.com/gofiber/fiber/v2"
)

func GetCombinations(c *fiber.Ctx) error {
	var combinations []models.Combination
	if err := models.GetAllCombinations(&combinations); err != nil {
		c.Status(fiber.StatusInternalServerError)
		c.JSON(err)
	}
	return c.JSON(combinations)
}
