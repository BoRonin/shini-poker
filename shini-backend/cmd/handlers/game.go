package handlers

import (
	"log"
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

type PlayerScore struct {
	models.Player
	Score int `json:"score"`
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
	var playerList []PlayerScore
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
		playerList = append(playerList, PlayerScore{
			Player: models.Player{
				Id:         uint(id),
				ChipsFinal: final_chips,
				Game:       game,
			},
		})
	}

	for i, v := range playerList {
		score, err := v.Player.SetFinalChips()
		if err != nil {
			return c.JSON(err.Error())
		}
		log.Println(score)
		playerList[i].Score = score
	}

	err = game.Finilize()
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		c.JSON(err)
	}
	return c.JSON(playerList)
}

func GetGames(c *fiber.Ctx) error {
	games, err := models.GetGames()
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return err
	}
	return c.JSON(games)
}

func IsGameFinished(c *fiber.Ctx) error {
	gameId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "bad request",
		})
	}
	game := new(models.Game)
	game.Id = uint(gameId)
	if err = game.IsFinished(); err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(err.Error())
	}
	return c.JSON(game.Finished)
}
