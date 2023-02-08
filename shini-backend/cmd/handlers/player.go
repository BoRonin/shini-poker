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
	player := models.Player{
		Id: uint(playerID),
	}
	if err := player.AddChips(chips); err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(err)

	}
	return c.JSON(fiber.Map{
		"message": "Success",
		"chips":   player.Chips,
	})
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

type playersForList struct {
	Id    uint   `json:"user_id"`
	Name  string `json:"user_name"`
	Chips int    `json:"chips"`
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
	var pfl []playersForList

	for _, v := range players {
		pfl = append(pfl, playersForList{
			Id:    v.User.Id,
			Name:  v.User.Name,
			Chips: v.Chips,
		})
	}
	return c.JSON(pfl)
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
