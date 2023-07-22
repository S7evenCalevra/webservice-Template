// cmd/main.go

package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/S7evenCalevra/webservice-Template/database"
)

func main() {

	database.ConnectDb()

	app := fiber.New()

	setupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Div Rhino Trivia App!")
	})

	app.Listen(":3000")
}
