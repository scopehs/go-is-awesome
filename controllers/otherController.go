// Other Controller
// Duno what i'm gona do in here yet!
package controllers

import "github.com/gofiber/fiber/v2"

// Well hello there Other...!
func Other(c *fiber.Ctx) error {
	return c.SendString("Other Controller 👋!")
}
