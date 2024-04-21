package controllers

import (
	"apilokdu/models"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	var ant []models.Loketdua
	models.DB.Db.Find(&ant)
	return c.Status(fiber.StatusOK).JSON(ant)
}

func Show(c *fiber.Ctx) error {
	ant := &models.Loketdua{}
	id := c.Params("id")
	if err := models.DB.Db.First(ant, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(ant)
}

func Update(c *fiber.Ctx) error {
	ant := &models.Loketdua{}
	id := c.Params("id")
	if err := models.DB.Db.First(ant, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}
	if err := c.BodyParser(ant); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	models.DB.Db.Save(ant)
	return c.Status(fiber.StatusOK).JSON(ant)
}

func Reset(c *fiber.Ctx) error {
	ant := &models.Loketdua{}
	id := c.Params("id")
	if err := models.DB.Db.First(ant, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}
	if err := c.BodyParser(ant); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	ant.Seq = 0
	models.DB.Db.Save(ant)
	return c.Status(fiber.StatusOK).JSON(ant)
}

func Create(c *fiber.Ctx) error {
	ant := new(models.Loketdua)
	if err := c.BodyParser(ant); err != nil {
	 return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	  "message": err.Error(),
	 })
	}
	models.DB.Db.Create(&ant)
	return c.Status(fiber.StatusCreated).JSON(ant)
   }
