package main

import (
	"github.com/gofiber/fiber/v2"
	"Go-Project-Template/internals"
)

func main() {
	app := fiber.New()

	app.Get("/," func(c *fiber.Ctx) error{
		return c.SendString("Hell, Testing")
	})
	app.Listen(":3000")
}
