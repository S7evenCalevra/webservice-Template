package handlers

import (
	"github.com/S7evenCalevra/webservice-Template/database"
	"github.com/S7evenCalevra/webservice-Template/models"
	"github.com/gofiber/fiber/v2"
)

func ReadRecords(c *fiber.Ctx) error {
	facts := []models.DataRecord{}
	database.DB.Db.Find(&facts)

	return c.Status(200).JSON(facts)
}

func CreateRecord(c *fiber.Ctx) error {
	fact := new(models.DataRecord)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}

func UpdateRecord(c *fiber.Ctx) error {
	fact := new(models.DataRecord)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}
