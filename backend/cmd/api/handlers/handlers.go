package handlers

import (
    "github.com/gofiber/fiber/v2"
    _ "github.com/jackc/pgconn"
    _ "github.com/jackc/pgx/v4"
    _ "github.com/jackc/pgx/v4/stdlib"
)


func Hi(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
        "message":"Hi, bro!",
    })
}

func Register(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
        "message":"This is registration",
    })
}

func Login(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
        "message": "This is login",
    })
}

func AddChips(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
        "message": "This is addchips",
        })
}
func Win(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
        "message": "This is win",
        })
}
func FinishGame(c *fiber.Ctx)error {
    return c.JSON(fiber.Map{
        "message":"This is finishgame",
    })
}

func GetStats(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
        "message": "This this is getstats",
        })
}