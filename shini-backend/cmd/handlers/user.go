package handlers

import (
	"shini/cmd/models"
	"shini/cmd/utils"
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
		return c.JSON(err.Error())
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

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	if err := models.GetAllUsers(&users); err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(err.Error())
	}
	c.Status(200)
	return c.JSON(users)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{

		"message": "success",
	})
}
