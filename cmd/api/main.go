package main

import (
	"fmt"
	"log"
	"shini/cmd/api/handlers"
	"shini/storage/postgres"

	"github.com/gofiber/fiber/v2"
)

const webPort = "3000"

func main() {
	app := fiber.New()
	//TO DO:
	//Register
	//Login
	//Create game
	//Add a player to a game
	//Register a win
	//Add chips
	//Finish the game and count money
	//Get stats
	//
	app.Get("/hi", handlers.Hi)
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)

	logged := app.Group("")
	logged.Post("/register", handlers.Register)
	logged.Post("/addchips/:id", handlers.AddChips)
	logged.Post("/win/:id", handlers.Win)
	logged.Post("/finishgame/:id", handlers.FinishGame)
	logged.Get("/getstats", handlers.GetStats)

	postgres.ConnectToDB()

	log.Fatal(app.Listen(fmt.Sprintf(":%s", webPort)))
}
