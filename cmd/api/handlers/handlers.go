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
