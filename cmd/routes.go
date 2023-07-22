// cmd/main.go

package main

import (
	"github.com/S7evenCalevra/webservice-Template/database"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)

	app.Post("/fact", handlers.CreateFact)
}
