package main

import (
	"fmt"
	"log"
	"shini/cmd/handlers"
	"shini/cmd/middlewares"
	"shini/storage/postgres"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const webPort = "3000"

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	//TO DO:
	//Get stats
	api := app.Group("api")
	api.Get("/user", handlers.User)
	api.Post("/register", handlers.Register)
	api.Post("/login", handlers.Login)
	api.Post("/logout", handlers.Logout)

	logged := api.Group("admin")
	logged.Use(middlewares.IsAuthenticated)
	logged.Post("/addchips/:id", handlers.AddChips)
	logged.Post("/win/:id", handlers.Win)
	logged.Post("/finishgame/:id", handlers.FinishGame)
	logged.Get("/getstats", handlers.GetStats)
	logged.Get("/users", handlers.GetUsers)
	logged.Post("/creategame", handlers.CreateGame)
	logged.Post("/addplayer/:id", handlers.AddPlayer)
	logged.Get("/getplayers/:id", handlers.GetPlayers)
	logged.Get("/combinations", handlers.GetCombinations)

	postgres.ConnectToDB()

	log.Fatal(app.Listen(fmt.Sprintf(":%s", webPort)))
}
