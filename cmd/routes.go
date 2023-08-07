package main

import (
	"github.com/S7evenCalevra/webservice-Template/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ReadRecords)

	app.Post("/fact", handlers.CreateRecord)
}
