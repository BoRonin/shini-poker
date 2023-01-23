package handlers

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/gofiber/fiber/v2"
    _ "github.com/jackc/pgconn"
    _ "github.com/jackc/pgx/v4"
    _ "github.com/jackc/pgx/v4/stdlib"
    "io"
    "log"
    "nasaapi/storage/postgres"
    "net/http"
    "strings"
    "sync"
    "time"
)


func Hi(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
        "message":"Hi, bro!",
    })
}
