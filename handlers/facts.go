package handlers

import (
	"github.com/S7evenCalevra/webservice-Template/handlers"
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Sending Trivia App!")
}
