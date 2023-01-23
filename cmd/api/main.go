package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "log"
    "shini/cmd/api/handlers"
    "shini/storage/postgres"
)

const webPort = "3000"


func main() {
    app := fiber.New()
    api.Get("/hi", handlers.Hi)

    postgres.ConnectToDB()

    log.Fatal(app.Listen(fmt.Sprintf(":%s", webPort)))
}


