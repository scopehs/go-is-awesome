package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/scopehs/tutorial/database"
	"github.com/scopehs/tutorial/models"
)

func AllPermissions(c *fiber.Ctx) error {
	var permissions []models.Permission

	database.DB.Find(&permissions)

	return c.JSON(permissions)
}
