package handlers

import (
	"shini/cmd/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateGame(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(err)
	}
	multiplier, _ := strconv.Atoi(data["multiplier"])
	game := models.Game{
		CreatedAt:  time.Now(),
		Multiplier: multiplier,
		Title:      data["title"],
	}
	if err := game.Create(); err != nil {
		c.Status(fiber.StatusForbidden)
		return c.JSON(fiber.Map{
			"message": "The game already exists",
		})
	}

	return c.JSON(game)
}

func FinishGame(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(err)
	}
	gameid, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "can't parse the game id",
		})
	}
	game := models.Game{
		Id: uint(gameid),
	}
	var playerList []models.Player
	for k, v := range data {
		id, err := strconv.Atoi(k)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(err)
		}
		final_chips, err := strconv.Atoi(v)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(err)
		}
		playerList = append(playerList, models.Player{
			Id:         uint(id),
			ChipsFinal: final_chips,
			Game:       game,
		})
	}

	for _, v := range playerList {
		score, err := v.SetFinalChips()
		if err != nil {
			return c.JSON(err.Error())
		}
		data[strconv.Itoa(int(v.Id))] = strconv.Itoa(score)
	}

	err = game.Finilize()
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		c.JSON(err)
	}
	return c.JSON(data)
}
