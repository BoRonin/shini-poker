package handlers

import (
	"fmt"
	"shini/cmd/api/models"
	"shini/cmd/api/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(err)
	}
	user := models.User{
		CreatedAt: time.Now(),
		Name:      data["name"],
		Email:     data["email"],
		Login:     data["login"],
	}
	user.SetPassword(data["password"])
	if err := user.Create(); err != nil {
		c.Status(fiber.StatusForbidden)
		return c.JSON(fiber.Map{
			"message": "No, dude, can't create",
		})
	}
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	info := new(models.LoginInfo)
	if err := c.BodyParser(&info); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(err)
	}
	user := models.User{}
	err := info.GetUserByLogin(&user)
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(err)
	}
	if err := user.ComparePasswords(info.Password); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}
	token, err := utils.GenerateJwt(int(user.Id))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(user)
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	userId, err := utils.ParseJwt(cookie)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "not logged in",
		})
	}
	id, _ := strconv.Atoi(userId)
	user := models.User{
		Id: uint(id),
	}
	if err := user.GetUserById(); err != nil {
		c.Status(fiber.StatusNotFound)
		c.JSON(err)
	}
	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}
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
			"message": "Unauthorized",
		})
	}

	return c.JSON(game)
}

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
	winId, _ := strconv.Atoi(data["combination"])
	player := models.Player{
		Id: uint(playerID),
	}
	if err := player.Win(winId); err != nil {
		c.Status(fiber.StatusBadRequest)
		c.JSON(err)
	}
	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("A win for player with id %d and combination with id %d", player.Id, winId),
	})
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

func GetStats(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "This this is getstats",
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
	fmt.Println(game)
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
