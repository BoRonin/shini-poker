package handlers

import (
	"fmt"
	"shini/cmd/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AddChips(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(err)
	}
	playerID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(err)
	}
	chips, _ := strconv.Atoi(data["chips"])
	gameId, _ := strconv.Atoi(data["game_id"])
	player := models.Player{
		Id: uint(playerID),
		Game: models.Game{
			Id: uint(gameId),
		},
	}
	if err := player.AddChips(chips); err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(err)
	}
	players, err := player.Game.GetPlayers()
	if err != nil {
		c.Status(fiber.StatusForbidden)
		return c.JSON(err)
	}
	return c.JSON(players)
}

func Win(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(err)
	}
	playerID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(err)
	}
	gameID, err := strconv.Atoi(data["game_id"])
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(err)
	}
	winId, _ := strconv.Atoi(data["combination"])
	player := models.Player{
		Id: uint(playerID),
		Game: models.Game{
			Id: uint(gameID),
		},
	}
	if err := player.Win(winId); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(err.Error())
	}
	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("A win for player with id %d and combination with id %d", player.Id, winId),
	})
}

func GetPlayers(c *fiber.Ctx) error {
	gameId, _ := strconv.Atoi(c.Params("id"))
	game := models.Game{
		Id: uint(gameId),
	}
	players, err := game.GetPlayers()
	if err != nil {
		c.Status(fiber.StatusForbidden)
		return c.JSON(err)
	}

	return c.JSON(players)
}

func AddPlayer(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(err)
	}
	gameId, _ := strconv.Atoi(c.Params("id"))
	id, _ := strconv.Atoi(data["user_id"])
	chips, _ := strconv.Atoi(data["chips"])
	newPlayer := models.Player{
		User: models.User{
			Id: uint(id),
		},
		Game: models.Game{
			Id: uint(gameId),
		},
		Chips: chips,
	}
	if err := newPlayer.Create(); err != nil {
		c.Status(fiber.StatusForbidden)
		return err
	}
	return c.JSON(newPlayer)
}
