package handlers

import (
    "github.com/gofiber/fiber/v2"
    _ "github.com/jackc/pgconn"
    _ "github.com/jackc/pgx/v4"
    _ "github.com/jackc/pgx/v4/stdlib"
    "shini/cmd/api/models"
    "strconv"
    "time"
)


func Hi(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
        "message":"Hi, bro!",
    })
}

func Register(c *fiber.Ctx) error {
    var data map[string]string
    if err := c.BodyParser(&data); err != nil {
        c.Status(fiber.StatusBadRequest)
        return c.JSON(err)
    }
    user := models.User{
        CreatedAt: time.Now(),
        Name: data["name"],
        Email: data["email"],
        Login: data["login"],
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
    return c.JSON(fiber.Map{
        "message": "This is login",
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
        CreatedAt: time.Now(),
        Multiplier: multiplier,
        Title: data["title"],
    }
    if err := game.Create(); err != nil {
        c.Status(fiber.StatusForbidden)
        return c.JSON(err)
    }

    return c.JSON(game)
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

func AddPlayer(c *fiber.Ctx) error{
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
     if err := newPlayer.Create(); err != nil{
         c.Status(fiber.StatusForbidden)
         return err
     }
    return c.JSON(newPlayer)
}