package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/scopehs/tutorial/database"
	"github.com/scopehs/tutorial/routes"
)

func main() {

	// Connect to Database
	database.Connect()

	// Make Fiber
	app := fiber.New()

	// Make routes
	routes.Setup(app)

	// Start Server
	app.Listen(":4444")
}
